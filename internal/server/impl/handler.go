package impl

import (
	"bytes"
	"context"
	"fmt"
	"path"
	"path/filepath"

	"github.com/reddec/api-notes/internal/render"
	"github.com/reddec/api-notes/internal/server/api"
	"github.com/reddec/api-notes/internal/storage"
)

const indexPage = "index.html"

type Server struct {
	Storage  storage.Storage
	Renderer *render.Renderer
	BaseURL  string
}

func (srv *Server) CreateNote(ctx context.Context, req *api.DraftMultipart) (*api.Note, error) {
	id := storage.GenID()
	if err := srv.storeNote(ctx, req, id); err != nil {
		return nil, err
	}
	return &api.Note{
		ID:        id,
		PublicURL: path.Join(srv.BaseURL, id),
	}, nil
}

func (srv *Server) storeNote(ctx context.Context, req *api.DraftMultipart, id string) error {
	var attachments []string
	if !req.HideAttachments.Value {
		attachments = make([]string, 0, len(req.Attachment))
		for _, name := range req.Attachment {
			attachments = append(attachments, name.Name)
		}
	}
	html, err := srv.Renderer.Render(req.Title, req.Text, req.Author.Value, attachments)
	if err != nil {
		return fmt.Errorf("render HTML: %w", err)
	}

	if err := srv.Storage.Set(ctx, filepath.Join(id, indexPage), bytes.NewReader(html)); err != nil {
		return fmt.Errorf("store HTML: %w", err)
	}

	for _, attachment := range req.Attachment {
		if attachment.Name == indexPage {
			continue
		}
		subID := filepath.Join(id, attachment.Name)
		if err := srv.Storage.Set(ctx, subID, attachment.File); err != nil {
			return fmt.Errorf("store attachment %s: %w", attachment.Name, err)
		}
	}
	return nil
}

func (srv *Server) DeleteNote(ctx context.Context, params api.DeleteNoteParams) error {
	return srv.Storage.Delete(ctx, params.ID)
}

func (srv *Server) UpdateNote(ctx context.Context, req *api.DraftMultipart, params api.UpdateNoteParams) error {
	return srv.storeNote(ctx, req, params.ID)
}
