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

func HandleEditFile(w http.ResponseWriter, r *http.Request) {
	path := utils.SanitizePath(r.URL.Path[6:]) // Remove "/edit/" prefix
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

	tmpl, err := template.ParseFiles(filepath.Join(models.TemplateDir, "edit.html"))
	if err != nil {
		http.Error(w, utils.UNABLE_TO_LOAD_FILE, http.StatusInternalServerError)
		return
	}

	data := struct {
		Path    string
		Content string
	}{
		Path:    path,
		Content: string(content),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, utils.UNABLE_TO_RENDER_TEMPLATE, http.StatusInternalServerError)
	}
}
