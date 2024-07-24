package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Coder-Kuldeep-Singh/go-file-server/internal/models"
	"github.com/Coder-Kuldeep-Singh/go-file-server/internal/utils"
)

func HandleSaveFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, utils.METHOD_NOT_ALLOWED, http.StatusMethodNotAllowed)
		return
	}

	path := utils.SanitizePath(r.URL.Path[6:]) // Remove "/save/" prefix
	fullPath := filepath.Join(models.BasePath, path)
	if !strings.HasPrefix(fullPath, models.BasePath) {
		http.Error(w, utils.ACCESS_DENIED, http.StatusInternalServerError)
		return
	}
	content := r.FormValue("content")
	err := os.WriteFile(fullPath, []byte(content), 0644)
	if err != nil {
		http.Error(w, utils.UNABLE_TO_SAVE_FILE, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/"+filepath.Dir(path), http.StatusSeeOther)
}
