package storage

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)

var ErrNotFound = os.ErrNotExist

type Storage interface {
	Get(ctx context.Context, uid string) (io.ReadCloser, error)
	Set(ctx context.Context, uid string, content io.Reader) error
	Delete(ctx context.Context, uid string) error
}

const idSize = 16

func GenID() string {
	var uid [idSize]byte
	_, err := io.ReadFull(rand.Reader, uid[:])
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(uid[:])
}

func IsValidID(value string) bool {
	v, err := hex.DecodeString(value)
	return err == nil && len(v) == idSize
}

func ShardedPath(value string) string {
	return filepath.Join(value[0:2], value[2:4], value[4:6], value[6:])
}
