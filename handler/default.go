package handler

import (
	"strings"

	"github.com/alexohneander/info-service/service/info"
	"github.com/gofiber/fiber/v2"
)

func Default(c *fiber.Ctx) error {
	clientInfo := info.GetAllClientInfos(c)

	// Debug
	// helpers.IsBrowser(clientInfo.UserAgent)

	if strings.Contains(clientInfo.UserAgent, "curl") {
		return c.SendString(clientInfo.IPv4)
	}

	return c.Render("index", fiber.Map{
		"clientInfo": clientInfo,
	})
}
