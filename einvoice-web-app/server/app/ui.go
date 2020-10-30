package app

import (
	"net/http"
	"os"
	"path/filepath"
)

type UiHandler struct {
	StaticPath string
	IndexPath  string
}

// Let client do the routing.
// If static file does not exist on requested path return file on IndexPath
func (h UiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepend the path with the path to the static directory
	path = filepath.Join(h.StaticPath, path)

	// Check if file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// File does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.StaticPath, h.IndexPath))
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		http.FileServer(http.Dir(h.StaticPath)).ServeHTTP(w, r)
	}
}
