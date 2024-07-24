package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Coder-Kuldeep-Singh/go-file-server/internal/models"
	"github.com/Coder-Kuldeep-Singh/go-file-server/internal/utils"
)

func HandleViewFile(w http.ResponseWriter, r *http.Request) {
	path := utils.SanitizePath(r.URL.Path[6:]) // Remove "/view/" prefix
	fullPath := filepath.Join(models.BasePath, path)
	if !strings.HasPrefix(fullPath, models.BasePath) {
		http.Error(w, utils.ACCESS_DENIED, http.StatusInternalServerError)
		return
	}

	content, err := os.ReadFile(fullPath)
	if err != nil {
		http.Error(w, utils.UNABLE_TO_READ_FILE, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(path))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(content)
}
