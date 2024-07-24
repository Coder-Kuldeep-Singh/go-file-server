package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Coder-Kuldeep-Singh/go-file-server/internal/models"
	"github.com/Coder-Kuldeep-Singh/go-file-server/internal/utils"
)

func HandleListFiles(w http.ResponseWriter, r *http.Request) {
	path := utils.SanitizePath(r.URL.Path[1:])
	if path == "" {
		path = models.BasePath
	}

	fullPath := filepath.Join(models.BasePath, path)

	if !strings.HasPrefix(fullPath, models.BasePath) {
		http.Error(w, utils.ACCESS_DENIED, http.StatusInternalServerError)
		return
	}

	files, err := os.ReadDir(fullPath)
	if err != nil {
		http.Error(w, utils.UNABLE_TO_LIST_FILES, http.StatusInternalServerError)
		return
	}

	var fileInfos []models.FileInfo
	for _, file := range files {
		fileInfos = append(fileInfos, models.FileInfo{
			Name:  file.Name(),
			IsDir: fmt.Sprintf("%t", file.IsDir()),
			Path:  filepath.Join(path, file.Name()),
		})
	}

	tmpl, err := template.ParseFiles(filepath.Join(models.TemplateDir, "index.html"))
	if err != nil {
		http.Error(w, utils.UNABLE_TO_LOAD_TEMPLATE, http.StatusInternalServerError)
		return
	}

	parentPath := models.BasePath
	if path != "." {
		parentPath = filepath.Dir(path)
		if parentPath == "." {
			parentPath = "" // Set to empty string for the root's parent
		}
	}

	// Calculate the previous path for the "Back" button
	previousPath := ""
	if path != "." && path != "" {
		previousPath = filepath.Dir(path)
		if previousPath == "." {
			previousPath = ""
		}
	}

	data := struct {
		Files          []models.FileInfo
		Path           string
		ParentPath     string
		PreviousPath   string
		ShowBackButton bool
	}{
		Files:          fileInfos,
		Path:           path,
		ParentPath:     parentPath,
		PreviousPath:   previousPath,
		ShowBackButton: path != "." && path != "",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, utils.UNABLE_TO_RENDER_TEMPLATE, http.StatusInternalServerError)
	}
}
