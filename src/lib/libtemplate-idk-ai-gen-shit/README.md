# Template Engine

This project provides a template engine using the `github.com/a-h/templ` package. The engine allows you to load and render templates defined as Go components.

## Usage

### 1. Define Your Components

Create Go files that define your components. Each component should implement the `templ.Component` interface and be exported as `Component`.

Example component file (`example.go`):

```go
package main

import (
    "context"
    "io"

    "github.com/a-h/templ"
)

type ExampleComponent struct{}

func (c ExampleComponent) Render(ctx context.Context, w io.Writer) error {
    _, err := w.Write([]byte("<h1>Hello, World!</h1>"))
    return err
}

var Component templ.Component = ExampleComponent{}
```

### 2. Initialize the Template Engine

Initialize the template engine with the base directory where your component files are located.

```go
package main

import (
    "log"

    "path/to/libtemplate"
)

func main() {
    te := libtemplate.NewTemplateEngine("/path/to/templates")
    err := te.LoadTemplates()
    if err != nil {
        log.Fatal(err)
    }

    data := struct {
        Title string
    }{
        Title: "My Page",
    }

    rendered, err := te.RenderTemplate("example.go", data)
    if err != nil {
        log.Fatal(err)
    }

    log.Println(rendered)
}
```

### 3. Nesting Components

You can nest components within other components. Define your components in Go files and use them within other components.

Example nested component file (`post.go`):

```go
package main

import (
    "context"
    "io"

    "github.com/a-h/templ"
)

type PostComponent struct {
    Title   string
    Content string
}

func (c PostComponent) Render(ctx context.Context, w io.Writer) error {
    _, err := w.Write([]byte("<div><h2>" + c.Title + "</h2><p>" + c.Content + "</p></div>"))
    return err
}

var Component templ.Component = PostComponent{}
```

Example parent component file (`blog.go`):

```go
package main

import (
    "context"
    "io"

    "github.com/a-h/templ"
)

type BlogComponent struct {
    Posts []templ.Component
}

func (c BlogComponent) Render(ctx context.Context, w io.Writer) error {
    _, err := w.Write([]byte("<div>"))
    if err != nil {
        return err
    }
    for _, post := range c.Posts {
        if err := post.Render(ctx, w); err != nil {
            return err
        }
    }
    _, err = w.Write([]byte("</div>"))
    return err
}

var Component templ.Component = BlogComponent{}
```

### 4. Rendering Nested Components

To render nested components, pass the nested components as data to the parent component.

Example usage:

```go
package main

import (
    "log"

    "path/to/libtemplate"
)

func main() {
    te := libtemplate.NewTemplateEngine("/path/to/templates")
    err := te.LoadTemplates()
    if err != nil {
        log.Fatal(err)
    }

    post1 := libtemplate.PostComponent{Title: "Post 1", Content: "Content of post 1"}
    post2 := libtemplate.PostComponent{Title: "Post 2", Content: "Content of post 2"}

    data := struct {
        Posts []templ.Component
    }{
        Posts: []templ.Component{post1, post2},
    }

    rendered, err := te.RenderTemplate("blog.go", data)
    if err != nil {
        log.Fatal(err)
    }

    log.Println(rendered)
}
```

## Internal Working

### Template Engine

The `TemplateEngine` struct is a wrapper around the `templ` package. It holds the templates and the base directory where the templates are stored.

```go
type TemplateEngine struct {
    templates map[string]templ.Component
    baseDir   string
}
```

#### NewTemplateEngine

The [NewTemplateEngine](vscode-file://vscode-app/opt/visual-studio-code/resources/app/out/vs/code/electron-sandbox/workbench/workbench.html) function initializes a new `TemplateEngine` with the given base directory.

```go
func NewTemplateEngine(baseDir string) *TemplateEngine {
    return &TemplateEngine{
        templates: make(map[string]templ.Component),
        baseDir:   baseDir,
    }
}
```

#### LoadTemplates

The `LoadTemplates` method loads all `.go` files from the base directory into the [templates](vscode-file://vscode-app/opt/visual-studio-code/resources/app/out/vs/code/electron-sandbox/workbench/workbench.html) map. It uses the [ParseComponent](vscode-file://vscode-app/opt/visual-studio-code/resources/app/out/vs/code/electron-sandbox/workbench/workbench.html) function to parse the component from the template file.

```go
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
```

#### RenderTemplate

The `RenderTemplate` method renders a template with the given name and data, returning the rendered string.

```go
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
```

#### Parser

The [ParseComponent](vscode-file://vscode-app/opt/visual-studio-code/resources/app/out/vs/code/electron-sandbox/workbench/workbench.html) function dynamically loads and parses the Go files as components using the `plugin` package.

```go
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
```

This function:

1. Creates a temporary file to store the template code.
2. Writes the template code to the temporary file.
3. Compiles the temporary file into a plugin.
4. Loads the compiled plugin.
5. Looks up the exported component symbol and asserts it to be of type `templ.Component`.
