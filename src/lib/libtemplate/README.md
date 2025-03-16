## Usage

### 1. Define Your Views Directory

Create a directory to store your HTML view files. Use placeholders in the format `{{placeholder}}` for dynamic content.

Example view file (`views/example.html`):

```html
<!DOCTYPE html>

<html>
	<head>
		<title>{{title}}</title>
	</head>

	<body>
		<h1>{{header}}</h1>
		<p>{{content}}</p>
	</body>
</html>
```

### 2. Initialize the Template Engine

Initialize the template engine with the base directory where your view files are located.

```go
package main

import (
    "log"
    "path/to/libtemplate"
)

func main() {
    // Define the directory where your views are stored
    viewsDir := "/path/to/views"

    te := libtemplate.NewTemplateEngine(viewsDir)
    err := te.LoadTemplates()
    if err != nil {
        log.Fatal(err)
    }

    data := map[string]string{
        "title":   "My Page",
        "header":  "Welcome to My Page",
        "content": "This is the content of my page.",
    }

    rendered, err := te.RenderTemplate("example.html", data)
    if err != nil {
        log.Fatal(err)
    }

    log.Println(rendered)
}
```

### 3. Stitching Views

You can stitch together a list of views into a single string.

Example usage:

```go
package main

import (
    "log"
    "path/to/libtemplate"
)

func main() {
    // Define the directory where your views are stored
    viewsDir := "/path/to/views"

    te := libtemplate.NewTemplateEngine(viewsDir)
    err := te.LoadTemplates()
    if err != nil {
        log.Fatal(err)
    }

    view1 := "<div>View 1</div>"
    view2 := "<div>View 2</div>"

    stitched := te.StitchViews([]string{view1, view2})

    log.Println(stitched)
}
```

## Internal Working

### Template Engine

The `TemplateEngine` struct is a simple wrapper that loads HTML templates, replaces placeholders with provided data, and supports stitching together lists of views.

```go
type TemplateEngine struct {
    templates map[string]string
    baseDir   string
}
```

#### NewTemplateEngine

The `NewTemplateEngine` function initializes a new `TemplateEngine` with the given base directory.

```go
func NewTemplateEngine(baseDir string) *TemplateEngine {
    return &TemplateEngine{
        templates: make(map[string]string),
        baseDir:   baseDir,
    }
}
```

#### LoadTemplates

The `LoadTemplates` method loads all `.html` files from the base directory into the `templates` map.

**func** **(**te** \*\*\***TemplateEngine**)** **LoadTemplates**(**)** **error** \*\*{

```go
func (te *TemplateEngine) LoadTemplates() error {
    err := filepath.Walk(te.baseDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && filepath.Ext(path) == ".html" {
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
```

#### RenderTemplate

The `RenderTemplate` method renders a template with the given name and data, returning the rendered string.

```go
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
```

#### StitchViews

The `StitchViews` method stitches together a list of views into a single string.

```go
func (te *TemplateEngine) StitchViews(views []string) string {
    return strings.Join(views, "")
}
```

## Conclusion

This template engine allows you to define your templates as HTML files with placeholders and dynamically load and render them. It also supports stitching together multiple views into a single string. Define your views directory, load the templates, and render them with the provided data. ```
