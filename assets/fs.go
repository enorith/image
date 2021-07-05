package assets

import (
	"embed"
	"io/fs"
)

//go:embed **/*
var FS embed.FS

func ReadFile(name string) ([]byte, error) {
	return fs.ReadFile(FS, name)
}
