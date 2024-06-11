package handler

import (
	"strings"

	"github.com/alexohneander/info-service/service/info"
	"github.com/gofiber/fiber/v2"
)

func Default(c *fiber.Ctx) error {
	clientInfo := info.GetAllClientInfos(c)

	if strings.Contains(clientInfo.UserAgent, "curl") {
		return c.SendString(clientInfo.IPv4)
	}

	return c.Render("index", fiber.Map{
		"clientInfo": clientInfo,
	})
}

func IP(c *fiber.Ctx) error {
	return c.SendString(info.GetIPv4Address(c))
}

func UA(c *fiber.Ctx) error {
	return c.SendString(info.GetUserAgent(c))
}

func Lang(c *fiber.Ctx) error {
	return c.SendString(info.GetLanguage(c))
}
