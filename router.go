package main

import (
	"changeme/frontend/components"
	"io/fs"
	"net/http"

	"github.com/a-h/templ"
)



func Router(assets fs.FS) http.Handler {
	static, err := fs.Sub(assets, "frontend/static")
	if err != nil {
		panic("Not static files found in frontend/static")
	}
	mux := http.NewServeMux()

	mux.Handle("/", templ.Handler(components.Layout(components.HelloWorld())))
	fileServer := http.FileServer(http.FS(static))
	mux.Handle("/files", http.StripPrefix("/files", fileServer))

	return mux
} 

