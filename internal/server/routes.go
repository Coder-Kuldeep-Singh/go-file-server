package server

import (
	"net/http"

	"github.com/Coder-Kuldeep-Singh/go-file-server/internal/controllers"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.HandleListFiles)
	mux.HandleFunc("/view/", controllers.HandleViewFile)
	mux.HandleFunc("/edit/", controllers.HandleEditFile)
	mux.HandleFunc("/save/", controllers.HandleSaveFile)
	mux.HandleFunc("/delete/", controllers.HandleDeleteFile)
	mux.HandleFunc("/rename/", controllers.HandleRename)
	mux.HandleFunc("/create/", controllers.HandleCreate)

	return mux
}
