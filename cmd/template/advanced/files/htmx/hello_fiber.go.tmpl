package web

import (
	"bytes"

	"github.com/gofiber/fiber/v2"
)

func HelloWebHandler(c *fiber.Ctx) error {
	// Parse form data
	if err := c.BodyParser(c); err != nil {
		c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	// Get the name from the form data
	name := c.FormValue("name")

	// Render the component
	component := HelloPost(name)
	buf := new(bytes.Buffer)
	component.Render(c.Context(), buf)

	// Send the response
	c.Status(fiber.StatusOK).SendString(buf.String())
	return nil
}
