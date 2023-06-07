package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jessevdk/go-flags"

	"github.com/reddec/api-notes/internal/render"
	"github.com/reddec/api-notes/internal/storage"
	"github.com/reddec/api-notes/internal/storage/local"
)

//nolint:gochecknoglobals
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

type Config struct {
	PublicURL string `short:"u" long:"public-url" env:"PUBLIC_URL" description:"Public URL for redirects" default:"http://127.0.0.1:8080"`
	Bind      string `short:"b" long:"bind" env:"BIND" description:"API binding address" default:"127.0.0.1:8080"`
	Dir       string `short:"d" long:"dir" env:"DIR" description:"Directory to store notes" default:"notes"`
}

func main() {
	var config Config
	parser := flags.NewParser(&config, flags.Default)
	parser.ShortDescription = "API-Notes"
	parser.LongDescription = fmt.Sprintf("Serve notes with API\napit-notes %s, commit %s, built at %s by %s\nAuthor: Aleksandr Baryshnikov <owner@reddec.net>", version, commit, date, builtBy)

	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}
	view := render.New()
	store := local.New(config.Dir)
	internal := chi.NewRouter()
	// create note
	internal.Post("/note", func(writer http.ResponseWriter, request *http.Request) {
		id := storage.GenID()
		if err := storeHandler(id, request, view, store); err != nil {
			_ = store.Delete(request.Context(), id)
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		writer.Header().Set("Location", config.PublicURL+"/"+id)
		writer.Header().Set("X-Correlation-ID", id)
		writer.WriteHeader(http.StatusSeeOther)
	})
	// update note
	internal.Put("/note/*", func(writer http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, "*")
		if err := storeHandler(id, request, view, store); err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		writer.Header().Set("Location", config.PublicURL+"/"+id)
		writer.WriteHeader(http.StatusSeeOther)
	})
	// delete note
	internal.Delete("/note/*", func(writer http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, "*")
		if err := store.Delete(request.Context(), id); err != nil {
			http.Error(writer, err.Error(), http.StatusBadGateway)
			return
		}
		writer.WriteHeader(http.StatusNoContent)
	})

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	server := &http.Server{Addr: config.Bind, Handler: internal}
	go func() {
		<-ctx.Done()
		_ = server.Close()
	}()
	log.Println("ready on", config.Bind)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Panicln(err)
	}
}

func storeHandler(id string, req *http.Request, view *render.Renderer, store storage.Storage) error {
	switch strings.Split(req.Header.Get("Content-Type"), ";")[0] {
	case "application/json":
		return fromJSON(id, req, view, store)
	case "multipart/form-data":
		return fromMultipart(id, req, view, store)
	default:
		return fromForm(id, req, view, store)
	}
}

func fromJSON(id string, req *http.Request, view *render.Renderer, store storage.Storage) error {
	type Note struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	}

	var note Note
	if err := json.NewDecoder(req.Body).Decode(&note); err != nil {
		return fmt.Errorf("decode JSON: %w", err)
	}

	content, err := view.Render(note.Title, note.Text)
	if err != nil {
		return fmt.Errorf("render markdown: %w", err)
	}
	if err := store.Set(req.Context(), path.Join(id, "index.html"), bytes.NewReader(content)); err != nil {
		return fmt.Errorf("store html: %w", err)
	}
	return nil
}

func fromForm(id string, req *http.Request, view *render.Renderer, store storage.Storage) error {
	title := req.FormValue("title")
	markdown := req.FormValue("text")
	content, err := view.Render(title, markdown)
	if err != nil {
		return fmt.Errorf("render markdown: %w", err)
	}
	if err := store.Set(req.Context(), path.Join(id, "index.html"), bytes.NewReader(content)); err != nil {
		return fmt.Errorf("store html: %w", err)
	}
	return nil
}

func fromMultipart(id string, req *http.Request, view *render.Renderer, store storage.Storage) error {
	ctx := req.Context()
	reader, err := req.MultipartReader()
	if err != nil {
		return fmt.Errorf("parse form: %w", err)
	}
	var title string
	for {
		r, err := reader.NextPart()
		if errors.Is(err, io.EOF) {
			break
		}
		switch r.FormName() {
		case "title":
			data, err := io.ReadAll(r)
			if err != nil {
				return fmt.Errorf("read title: %w", err)
			}
			title = string(data)
		case "text":
			markdown, err := io.ReadAll(r)
			if err != nil {
				return fmt.Errorf("read markdown: %w", err)
			}
			content, err := view.Render(title, string(markdown))
			if err != nil {
				return fmt.Errorf("render markdown: %w", err)
			}
			if err := store.Set(ctx, path.Join(id, "index.html"), bytes.NewReader(content)); err != nil {
				return fmt.Errorf("store html: %w", err)
			}
		default:
			fname := r.FileName()
			if fname == "" {
				continue
			}
			if err := store.Set(ctx, path.Join(id, fname), r); err != nil {
				return fmt.Errorf("store attachment %s: %w", fname, err)
			}
		}
	}
	return nil
}
