package cicd

type GithubActionTemplate struct{}

func (a GithubActionTemplate) Dockerfile() []byte {
	return MakeDockerfile()
}

func (a GithubActionTemplate) Pipline() []byte {
	return []byte(`name: Go-test
on: [push]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4
      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21.x'
      - name: Build
        run: go build -v ./...
      - name: Test with the Go CLI
        run: go test ./... 
	`)
}