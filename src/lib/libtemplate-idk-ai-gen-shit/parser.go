package libtemplate

import (
	"errors"
	"os"
	"os/exec"
	"plugin"
	"strings"

	"github.com/a-h/templ"
)

// parseComponent loads and parses the component from the template file
func ParseComponent(tmpl string) (templ.Component, error) {
	// Create a temporary file to store the template code
	tmpFile, err := os.CreateTemp("", "*.go")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	// Write the template code to the temporary file
	if _, err := tmpFile.WriteString(tmpl); err != nil {
		return nil, err
	}
	if err := tmpFile.Close(); err != nil {
		return nil, err
	}

	// Compile the temporary file into a plugin
	pluginPath := strings.TrimSuffix(tmpFile.Name(), ".go") + ".so"
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", pluginPath, tmpFile.Name())
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	defer os.Remove(pluginPath)

	// Load the compiled plugin
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return nil, err
	}

	// Lookup the exported component symbol
	symbol, err := p.Lookup("Component")
	if err != nil {
		return nil, err
	}

	// Assert the symbol to be of type templ.Component
	component, ok := symbol.(templ.Component)
	if !ok {
		return nil, errors.New("invalid component type")
	}

	return component, nil
}
