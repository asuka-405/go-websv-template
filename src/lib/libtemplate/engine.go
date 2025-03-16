package libtemplate

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// TemplateEngine is a simple template engine
type TemplateEngine struct {
	templates map[string]string
	baseDir   string
}

// NewTemplateEngine creates a new TemplateEngine
func NewTemplateEngine(baseDir string) *TemplateEngine {
	return &TemplateEngine{
		templates: make(map[string]string),
		baseDir:   baseDir,
	}
}

// LoadTemplates loads all HTML templates from the base directory
func (te *TemplateEngine) LoadTemplates() error {
	err := filepath.Walk(te.baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			tmpl, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			te.templates[filepath.Base(path)] = string(tmpl)
		}
		return nil
	})
	return err
}

// RenderTemplate renders a template with the given name and data
func (te *TemplateEngine) RenderTemplate(name string, data map[string]string) (string, error) {
	tmpl, ok := te.templates[name]
	if !ok {
		return "", os.ErrNotExist
	}

	for placeholder, value := range data {
		tmpl = strings.ReplaceAll(tmpl, "{{"+placeholder+"}}", value)
	}

	return tmpl, nil
}
func (te *TemplateEngine) RenderWithLogs(name string, data map[string]string) string {
	tmpl, ok := te.templates[name]
	if !ok {
		log.Fatal("Template not found: " + name)
		return ""
	}

	for placeholder, value := range data {
		tmpl = strings.ReplaceAll(tmpl, "{{"+placeholder+"}}", value)
	}

	return tmpl
}

// StitchViews stitches together a list of views into a single string
func (te *TemplateEngine) StitchViews(views []string) string {
	return strings.Join(views, "")
}
