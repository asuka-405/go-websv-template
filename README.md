# GoLang WebServer Starter Template

Dependencies:

- go-chi
- go-chi/cors
- godotenv

## After Clone Setup

- Do migrate mod name to your github repo.
- Also update all the local imports for same.

### Project Structure

.
├── go.mod
├── go.sum
├── README.md
├── run
└── src
├── api
│   ├── api.go
│   ├── server.go
│   └── v1
│   ├── h_readiness.go
│   └── router.go
├── lib
│   ├── fs
│   │   └── fs.go
│   └── libresponse
│   ├── err.go
│   └── json.go
├── main.go
└── public
└── index.html

8 directories, 13 files
