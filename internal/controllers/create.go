package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Coder-Kuldeep-Singh/go-file-server/internal/utils"

	"github.com/Coder-Kuldeep-Singh/go-file-server/internal/models"
)

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, utils.METHOD_NOT_ALLOWED, http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	path := utils.SanitizePath(r.FormValue("path"))
	createType := r.FormValue("type")

	fullPath := filepath.Join(models.BasePath, path, name)

	if !strings.HasPrefix(fullPath, models.BasePath) {
		http.Error(w, utils.ACCESS_DENIED, http.StatusInternalServerError)
		return
	}

	if createType == "file" {
		_, err := os.Create(fullPath)
		if err != nil {
			http.Error(w, utils.UNABLE_TO_CREATE_FILE, http.StatusInternalServerError)
			return
		}
	} else if createType == "folder" {
		err := os.Mkdir(fullPath, 0755)
		if err != nil {
			http.Error(w, utils.UNABLE_TO_CREATE_FOLDER, http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, utils.INVALID_CREATE_TYPE, http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/"+strings.TrimPrefix(path, "./"), http.StatusSeeOther)
}
