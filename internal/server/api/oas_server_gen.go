// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateNote implements createNote operation.
	//
	// Create new note from draft with (optional) attachments.
	// Returns public URL and unique ID.
	// Consumer should not make any assumptions about ID and treat it as
	// arbitrary string with variable reasonable length.
	// Attachments with name index.html will be ignored.
	// Note can use relative reference to attachments as-is.
	//
	// POST /notes
	CreateNote(ctx context.Context, req *DraftMultipart) (*Note, error)
	// DeleteNote implements deleteNote operation.
	//
	// Remove existent note and all attachments.
	//
	// DELETE /note/{id}
	DeleteNote(ctx context.Context, params DeleteNoteParams) error
	// UpdateNote implements updateNote operation.
	//
	// Update existent note by ID.
	// Old attachments may not be removed, but could be replaced.
	//
	// PUT /note/{id}
	UpdateNote(ctx context.Context, req *DraftMultipart, params UpdateNoteParams) error
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
