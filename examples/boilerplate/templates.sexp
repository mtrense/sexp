(boilerplate 
    (title "Simple Golang project template")
    (maintainer "Max Trense" "dev@trense.info")
    (url "https://github.com/mtrense/tmpl8-golang")

    (string "name" "Project name"
        (default $directory.base))
    )
    (boolean "gomod" "Run `go mod init`"
        (default true)
    )
    (string "projecturl" "Go project URL"
        (default (format "github.com/%s/%s" $env.github_user $var.name))
        (when $var.gomod)
    )
    (boolean "cli" "Command Line Application"
        (default true)
    )
    (boolean "webservice" "Create a web service listener"
        (default false)
    )
    (boolean "makefile" "Create a Makefile"
        (default true)
    )
    (boolean "docker-compose" "Create a docker-compose file"
        (default false)
    )
    (boolean "envrc" "Create a local .envrc file"
        (default true)
    )
    (boolean "database" "Configure a postgresql database"
        (default false)
        (when $var.docker-compose))
    )
    (boolean "logging" "Configure logging"
        (default true)
    )
    (multiple-choice "libraries" "Libraries to include"
        (element "github.com/rs/zerolog/log"    "Zero Allocation JSON Logger")
        (element "github.com/osteele/liquid"    "A complete Liquid template engine in Go")
        (element "github.com/antonmedv/expr"    "Expression evaluation engine for Go: fast, non-Turing complete, dynamic typing, static typing")
        (element "github.com/kataras/iris/v12"  "The fastest community-driven web framework for Go")
        (element "github.com/rjeczalik/notify"  "File system event notification library on steroids.")
        (element "github.com/gonum/gonum"       "Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more")
        (element "github.com/gizak/termui"      "Golang terminal dashboard")
        (element "github.com/jmoiron/sqlx"      "general purpose extensions to golang's database/sql")
        (element "github.com/rivo/tview"        "Rich interactive widgets for terminal-based UIs written in Go")
        (element "github.com/jroimartin/gocui"  "Minimalist Go package aimed at creating Console User Interfaces")
        (element "github.com/jedib0t/go-pretty" "Pretty print Tables and more in golang")
        (element "github.com/davecgh/go-spew"   "Implements a deep pretty printer for Go data structures to aid in debugging")
    )
    (template "Makefile"
        (file "makefile.tmpl")
        (when $var.makefile)
    )
    (template ".envrc"
        (file "envrc.tmpl")
        (when $var.envrc)
    )
)