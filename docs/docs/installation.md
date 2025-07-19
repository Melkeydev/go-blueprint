---
hide:
  - toc
---

Go-Blueprint provides a convenient CLI tool to effortlessly set up your Go projects. Follow the steps below to install the tool on your system.

## Binary Installation

To install the Go-Blueprint CLI tool as a binary, run the following command:

```sh
go install github.com/melkeydev/go-blueprint@latest
```

This command installs the Go-Blueprint binary, automatically binding it to your `$GOPATH`.

> If you’re using Zsh, you’ll need to add it manually to `~/.zshrc`.

> After running the installation command, you need to update your `PATH` environment variable. To do this, you need to find out the correct `GOPATH` for your system. You can do this by running the following command:
> Check your `GOPATH`
>
> ```
> go env GOPATH
> ```
>
> Then, add the following line to your `~/.zshrc` file:
>
> ```
> GOPATH=$HOME/go  PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
> ```
>
> Save the changes to your `~/.zshrc` file by running the following command:
>
> ```
> source ~/.zshrc
> ```

## NPM Install

If you prefer using Node.js package manager, you can install Go-Blueprint via NPM. This method is convenient for developers who are already working in JavaScript/Node.js environments and want to integrate Go-Blueprint into their existing workflow.

```bash
npm install -g @Melkeydev/go-blueprint
```

The `-g` flag installs Go-Blueprint globally, making it accessible from any directory on your system.

## Homebrew Install

For macOS and Linux users, Homebrew provides a simple way to install Go-Blueprint. Homebrew automatically handles dependencies and keeps the tool updated through its package management system.

```bash
brew install go-blueprint
```

After installation via Homebrew, Go-Blueprint will be automatically added to your PATH, making it immediately available in your terminal.

## Building and Installing from Source

If you prefer to build and install Go-Blueprint directly from the source code, you can follow these steps:

Clone the Go-Blueprint repository from GitHub:

```sh
git clone https://github.com/melkeydev/go-blueprint
```

Build the Go-Blueprint binary:

```sh
go build
```

Install in your `$PATH` to make it accessible system-wide:

```sh
go install
```

Verify the installation by running:

```sh
go-blueprint version
```

This should display the version information of the installed Go-Blueprint.

Now you have successfully built and installed Go-Blueprint from the source code.
