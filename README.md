# GoLang WebServer Starter Template

Project is structured in a modular way. There are following modules with their own specific purpose:

- server: contains server initialization code. Add any code here that might be tightly coupled with server.
- api: write any api specific endpoints here.
- web: write your ui (html) endpoints here.
- lib: mini libs or utilities to help develop with ease.
  - libauth: write authentication - authorization specific code here.
  - libfs: file server is here + write any fs specific code here.
  - libresponse: way to standardize your server response.
  - libtemplate: a minimal template engine, extend it with your own code.

`main.go` is the entry point of our server.

Public dir contains code served at /static route, you can chage the dir name by setting `DIR_PUBLIC` env.

Server runs at `port:3001` by default, set `PORT` env to change this.

`./run dev` command starts the server.

Use `run` executable to do various tasks.

- `./run dev` start dev server.
- `./run clean` remove existing test server binary.
- `./run tidy` to tidy and vendor your go modules.
- `./run install <pkg_name> | ./run i <pkg_name>` to install a package.
- `./run tree` to show the project tree (omit .git, vendor, etc)

> This project is tested on `ArchLinux` , so `run` binary wont run on anything but linux, preferably on `ArchLinux`.

### Project Structure

```shell
 go-websv-template  ./run tree
.
├── go.mod
├── go.sum
├── install.sh
├── README.md
├── run
└── src
    ├── api
    │   ├── router.go
    │   └── v1
    │       ├── h_readiness.go
    │       └── router.go
    ├── lib
    │   ├── libauth
    │   │   ├── auth.go
    │   │   ├── jwt.go
    │   │   ├── middleware.go
    │   │   ├── README.md
    │   │   └── session.go
    │   ├── libfirebase
    │   │   ├── oauth.go
    │   │   └── README.md
    │   ├── libfs
    │   │   ├── reader.go
    │   │   └── server.go
    │   ├── libresponse
    │   │   ├── err.go
    │   │   ├── html.go
    │   │   └── json.go
    │   ├── libtemplate
    │   │   ├── engine.go
    │   │   └── README.md
    │   └── libtemplate-idk-ai-gen-shit
    │       ├── engine.go
    │       ├── parser.go
    │       └── README.md
    ├── main.go
    ├── public
    │   ├── css
    │   │   ├── colors.scss
    │   │   ├── main.css
    │   │   ├── main.css.map
    │   │   ├── main.scss
    │   │   └── partials
    │   │       └── reset.sass
    │   ├── index.html
    │   └── js
    ├── server
    │   └── server.go
    └── web
        ├── h_readiness.go
        ├── router.go
        └── views
            ├── default-head.wc
            ├── index.layout.html
            └── test.wc
```
