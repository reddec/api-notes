package commands

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/reddec/api-notes/internal/render"
	"github.com/reddec/api-notes/internal/server/api"
	"github.com/reddec/api-notes/internal/server/impl"
	"github.com/reddec/api-notes/internal/storage/local"
)

type ServeCMD struct {
	PublicURL string `short:"u" long:"public-url" env:"PUBLIC_URL" description:"Public URL for redirects" default:"http://127.0.0.1:8080"`
	Bind      string `short:"b" long:"bind" env:"BIND" description:"API binding address" default:"127.0.0.1:8080"`
	Dir       string `short:"d" long:"dir" env:"DIR" description:"Directory to store notes" default:"notes"`
	Token     string `short:"t" long:"token" env:"TOKEN" description:"Authorization token, empty means any token can be usedauth is disabled"`
}

func (cmd *ServeCMD) Execute([]string) error {
	view := render.New()
	store := local.New(cmd.Dir)

	var security api.SecurityHandler = &impl.StaticToken{Token: cmd.Token}
	if cmd.Token == "" {
		security = &impl.AnyToken{}
	}

	handler, err := api.NewServer(&impl.Server{
		Storage:  store,
		Renderer: view,
		BaseURL:  cmd.PublicURL,
	}, security)

	if err != nil {
		return fmt.Errorf("create server: %w", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	server := &http.Server{Addr: cmd.Bind, Handler: handler}
	go func() {
		<-ctx.Done()
		_ = server.Close()
	}()
	log.Println("ready on", cmd.Bind)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("serve: %w", err)
	}
	return nil
}
