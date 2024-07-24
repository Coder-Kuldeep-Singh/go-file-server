package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Coder-Kuldeep-Singh/go-file-server/internal/models"
	"github.com/Coder-Kuldeep-Singh/go-file-server/internal/utils"
)

func HandleRename(w http.ResponseWriter, r *http.Request) {
	path := utils.SanitizePath(r.URL.Path[8:]) // Remove "/rename/" prefix
	fullPath := filepath.Join(models.BasePath, path)
	if !strings.HasPrefix(fullPath, models.BasePath) {
		http.Error(w, utils.ACCESS_DENIED, http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles(filepath.Join(models.TemplateDir, "rename.html"))
		if err != nil {
			http.Error(w, utils.UNABLE_TO_LOAD_TEMPLATE, http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, path)
		if err != nil {
			http.Error(w, utils.UNABLE_TO_RENDER_TEMPLATE, http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodPost {
		newName := r.FormValue("newName")
		newPath := filepath.Join(filepath.Dir(fullPath), newName)

		err := os.Rename(fullPath, newPath)
		if err != nil {
			http.Error(w, utils.UNABLE_TO_RENAME_FILE, http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/"+filepath.Dir(path), http.StatusSeeOther)
	} else {
		http.Error(w, utils.METHOD_NOT_ALLOWED, http.StatusMethodNotAllowed)
	}
}
