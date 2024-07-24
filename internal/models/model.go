package models

const (
	BasePath = "files" // Base directory for file operations
	// use / for giving access of root directory of system
	TemplateDir = "./internal/templates"
)

type FileInfo struct {
	Name  string
	IsDir string
	Path  string
}
