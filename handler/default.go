package handler

import (
	"github.com/alexohneander/info-service/service/info"
	"github.com/gofiber/fiber/v2"
)

func Default(c *fiber.Ctx) error {

	clientInfo := info.GetAllClientInfos(c)
	// clientInfoString := fmt.Sprintf("%v", clientInfo)

	return c.JSON(clientInfo)
}
