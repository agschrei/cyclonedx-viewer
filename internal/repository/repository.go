package repository

import (
	"embed"
	"io/fs"
)

//go:embed templates/*
var templates embed.FS

//go:embed static/*
var static embed.FS

var FileRepository = buildRepository()

type Repository struct {
	Templates fs.FS
	Static    fs.FS
}

func buildRepository() Repository {
	templates, _ := fs.Sub(templates, "templates")

	return Repository{
		Templates: templates,
		Static:    static,
	}
}
