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

type Server struct {
	Storage  storage.Storage
	Renderer *render.Renderer
	BaseURL  string
}

func (srv *Server) CreateNote(ctx context.Context, req api.OptDraftMultipart) (*api.Note, error) {
	id := storage.GenID()
	if err := srv.storeNote(ctx, req, id); err != nil {
		return nil, err
	}
	return &api.Note{
		ID:        id,
		PublicURL: path.Join(srv.BaseURL, id),
	}, nil
}

func (srv *Server) storeNote(ctx context.Context, req api.OptDraftMultipart, id string) error {
	html, err := srv.Renderer.Render(req.Value.Title, req.Value.Text)
	if err != nil {
		return fmt.Errorf("render HTML: %w", err)
	}

	if err := srv.Storage.Set(ctx, id, bytes.NewReader(html)); err != nil {
		return fmt.Errorf("store HTML: %w", err)
	}

	for _, attachment := range req.Value.Attachment {
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

func (srv *Server) UpdateNote(ctx context.Context, req api.OptDraftMultipart, params api.UpdateNoteParams) error {
	return srv.storeNote(ctx, req, params.ID)
}
