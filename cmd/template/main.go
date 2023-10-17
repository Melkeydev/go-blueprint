package template

func MainTemplate() []byte {
	return []byte(`
package main

import "valentine/cmd"

func main() {
	cmd.Execute()
}
`)
}
