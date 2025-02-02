package advanced

import (
	_ "embed"
)

//go:embed files/websocket/imports/standard_library.tmpl
var stdLibWebsocketImports []byte

//go:embed files/websocket/imports/fiber.tmpl
var fiberWebsocketTemplImports []byte

func StdLibWebsocketTemplImportsTemplate() []byte {
	return stdLibWebsocketImports
}

func FiberWebsocketTemplImportsTemplate() []byte {
	return fiberWebsocketTemplImports
}
