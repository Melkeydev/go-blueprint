package template

// defines the files to be created

func MainTemplate() []byte {
	return []byte(`
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
`)
}
