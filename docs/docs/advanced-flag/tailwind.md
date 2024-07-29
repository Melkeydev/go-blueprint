Tailwind is closely coupled with the advanced HTMX flag, and HTMX will be automatically used if you select Tailwind in your project.

We do not introduce outside dependencies automatically, and you need compile output.css (file is empty by default) with the Tailwind CLI tool.

The project tree would look like this:

```bash
/ (Root)
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── web/
│       ├── assets/
│       │   ├── css/
│       │   │   ├── input.css
│       │   │   └── output.css
│       │   └── js/
│       │       └── htmx.min.js
│       ├── base.templ
│       ├── base_templ.go
│       ├── efs.go
│       ├── hello.go
│       ├── hello.templ
│       └── hello_templ.go
├── internal/
│   └── server/
│       ├── routes.go
│       ├── routes_test.go
│       └── server.go
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── tailwind.config.js
```

## Standalone Tailwind CLI

The idea is not to use Node.js and npm to build `output.css`. To achieve this, visit the [official repository](https://github.com/tailwindlabs/tailwindcss/releases/latest) and download the latest release version for your OS and Arch:

```bash
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.4/tailwindcss-linux-x64
```

Give execution permission:

```bash
chmod +x tailwindcss-linux-x64
```

## Compile and minify your CSS

```bash
./tailwindcss-linux-x64 -i cmd/web/assets/css/input.css -o cmd/web/assets/css/output.css
```

## Use Tailwind CSS in your project

By default, CSS examples are not included in the codebase.
Update base.templ and hello.templ, then rerun templ generate to see the changes at the localhost:PORT/web endpoint.

## Standalone Rustywind

[Rustywind](https://github.com/avencera/rustywind) is a CLI tool for oraganizing and formatting Tailwind classes by [avencera](https://github.com/avencera)

```bash
curl -LSfs https://avencera.github.io/rustywind/install.sh | sh -s -- --git avencera/rustywind
```

## Rustywind formatting

Rustywind will check and format all `.templ` files inside cmd/web folder

```bash
rustywind cmd/web/assets/css/input.css --write
rustywind --write $(shell find cmd/web -name '*.templ')
```
