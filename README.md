![logo](./public/logo.png)

<div style="text-align: center;">
  <h1>
    Introducing the Ultimate Golang Blueprint Library
  </h1>
</div>

Go Blueprint is a CLI tool that allows users to spin up a Go project with the corresponding structure seamlessly. It also
gives the option to integrate with one of the more popular Go frameworks (and the list is growing with new features)!

### Why Would I use this?

- Easy to set up and install
- Have the entire Go structure already established
- Setting up a Go HTTP server (or Fasthttp with Fiber)
- Integrate with a popular frameworks
- Focus on the actual code of your application

## Table of Contents

- [Install](#install)
- [Frameworks Supported](#frameworks-supported)
- [Database Support](#database-support)
- [Usage Example](#usage-example)
- [GitHub Stats](#github-stats)
- [License](#license)

<a id="install"></a>

<h2>
  <picture>
    <img src="./public/install.gif?raw=true" width="60px" style="margin-right: 1px;">
  </picture>
  Install
</h2>

```sh
go install github.com/melkeydev/go-blueprint@latest
```

This installs a go binary that will automatically bind to your $GOPATH

Then in a new terminal run:

```
go-blueprint create
```

You can also use the provided flags to set up a project without interacting with the UI.

```
go-blueprint create --name my-project --framework gin --driver postgres
```

See `go-blueprint create -h` for all the options and shorthands.

<a id="frameworks-supported"></a>

<h2>
  <picture>
    <img src="./public/frameworks.gif?raw=true" width="60px" style="margin-right: 1px;">
  </picture>
  Frameworks Supported
</h2>

- [Chi](https://github.com/go-chi/chi)
- [Gin](https://github.com/gin-gonic/gin)
- [Fiber](https://github.com/gofiber/fiber)
- [HttpRouter](https://github.com/julienschmidt/httprouter)
- [Gorilla/mux](https://github.com/gorilla/mux)
- [Echo](https://github.com/labstack/echo)

<a id="database-support"></a>

<h2>
  <picture>
    <img src="./public/database.gif?raw=true" width="45px" style="margin-right: 15px;">
  </picture>
  Database Support
</h2>

Go Blueprint now offers enhanced database support, allowing you to choose your preferred database driver during project setup. Use the `--driver` or `-d` flag to specify the database driver you want to integrate into your project.

### Supported Database Drivers

Choose from a variety of supported database drivers:

- [Mysql](https://github.com/go-sql-driver/mysql)
- [Postgres](https://github.com/lib/pq)
- [Sqlite](https://github.com/mattn/go-sqlite3)
- [Mongo](go.mongodb.org/mongo-driver)

<a id="usage-example"></a>

<h2>
  <picture>
    <img src="./public/example.gif?raw=true" width="60px" style="margin-right: 1px;">
  </picture>
  Usage Example
</h2>

Here's an example of setting up a project with a specific database driver:

```bash
go-blueprint create --name my-project --framework gin --driver postgres
```

<p align="center">
  <img src="./public/blueprint_1.png" alt="Starter Image" width="800"/>
</p>

<a id="github-stats"></a>

<h2>
  <picture>
    <img src="./public/stats.gif?raw=true" width="45px" style="margin-right: 10px;">
  </picture>
  Github Stats
</h2>

<p align="center">
  <img alt="Alt" src="https://repobeats.axiom.co/api/embed/7c4be18864d441f961be61186ce49b5471a9e7bf.svg" title="Repobeats analytics image"/>
</p>

<a id="license"></a>

<h2>
  <picture>
    <img src="./public/license.gif?raw=true" width="50px" style="margin-right: 1px;">
  </picture>
  Licence
</h2>

Licensed under [MIT License](./LICENSE)
