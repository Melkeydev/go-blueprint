package web

import (
	"bytes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HelloWebHandler(c *fiber.Ctx) error {
	// Parse form data
	if err := c.BodyParser(c); err != nil {
		innerErr := c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		if innerErr != nil {
			log.Fatalf("Could not send error in HelloWebHandler: %e", innerErr)
		}
	}

	// Get the name from the form data
	name := c.FormValue("name")

	// Render the component
	component := HelloPost(name)
	buf := new(bytes.Buffer)
	err := component.Render(c.Context(), buf)
	if err != nil {
		errorString := fmt.Sprintf("Error rendering in HelloWebHandler: %e", err)
		innerErr := c.Status(fiber.StatusBadRequest).SendString(errorString)
		if innerErr != nil {
			log.Fatalf("Could not send error in HelloWebHandler: %e", innerErr)
		}
		log.Fatalf(errorString)
	}

	// Send the response
	err = c.Status(fiber.StatusOK).SendString(buf.String())
	if err != nil {
		log.Fatalf("Could not send OK in HelloWebHandler: %e", err)
	}
	return nil
}
