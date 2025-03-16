package libtemplate

import (
	"bytes"
	"context"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
)

// TemplateEngine is a wrapper around the templ package
type TemplateEngine struct {
	templates map[string]templ.Component
	baseDir   string
}

// NewTemplateEngine creates a new TemplateEngine
func NewTemplateEngine(baseDir string) *TemplateEngine {
	return &TemplateEngine{
		templates: make(map[string]templ.Component),
		baseDir:   baseDir,
	}
}

// LoadTemplates loads all templates from the base directory
func (te *TemplateEngine) LoadTemplates() error {
	err := filepath.Walk(te.baseDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			tmpl, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			// Assuming the template files define components
			component, err := ParseComponent(string(tmpl))
			if err != nil {
				return err
			}
			te.templates[filepath.Base(path)] = component
		}
		return nil
	})
	return err
}

// RenderTemplate renders a template with the given name and data
func (te *TemplateEngine) RenderTemplate(name string, data interface{}) (string, error) {
	tmpl, ok := te.templates[name]
	if !ok {
		return "", os.ErrNotExist
	}

	var buf bytes.Buffer
	ctx := context.Background()
	if err := tmpl.Render(ctx, &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}
