package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Coder-Kuldeep-Singh/go-file-server/internal/models"
	"github.com/Coder-Kuldeep-Singh/go-file-server/internal/utils"
)

func HandleDeleteFile(w http.ResponseWriter, r *http.Request) {
	path := utils.SanitizePath(r.URL.Path[8:]) // Remove "/delete/" prefix
	fullPath := filepath.Join(models.BasePath, path)
	if !strings.HasPrefix(fullPath, models.BasePath) {
		http.Error(w, utils.ACCESS_DENIED, http.StatusInternalServerError)
		return
	}
	err := os.Remove(fullPath)
	if err != nil {
		http.Error(w, utils.UNABLE_TO_REMOVE_FILE, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/"+filepath.Dir(path), http.StatusSeeOther)
}
