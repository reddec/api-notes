package local

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func New(dir string) *Local {
	d, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	return &Local{dir: d}
}

type Local struct {
	dir string
}

func (l *Local) Get(_ context.Context, uid string) (io.ReadCloser, error) {
	return os.Open(filepath.Join(l.dir, uid))
}

func (l *Local) Set(_ context.Context, uid string, content io.Reader) error {
	p := filepath.Join(l.dir, uid)
	dir := filepath.Dir(p)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("create dir: %w", err)
	}

	tmp, err := os.CreateTemp(dir, "")
	if err != nil {
		return fmt.Errorf("create temp: %w", err)
	}
	defer os.RemoveAll(tmp.Name())
	defer tmp.Close()

	if _, err := io.Copy(tmp, content); err != nil {
		return fmt.Errorf("copy content: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("close temp: %w", err)
	}
	return os.Rename(tmp.Name(), p)
}

func (l *Local) Delete(_ context.Context, uid string) error {
	return os.RemoveAll(filepath.Join(l.dir, uid))
}
