package info

import (
	"github.com/alexohneander/info-service/model"
	"github.com/gofiber/fiber/v2"
)

// Public Functions
func GetAllClientInfos(c *fiber.Ctx) model.ClientInfo {
	clientInfo := model.ClientInfo{}

	// Get Informations
	clientInfo.IPv4 = getIPv4Address(c)
	clientInfo.UserAgent = getUserAgent(c)
	clientInfo.Language = getLanguage(c)
	clientInfo.Referer = getReferer(c)
	clientInfo.Method = getMethod(c)
	clientInfo.Encoding = getEncoding(c)
	clientInfo.MIMEType = getMIMEType(c)
	clientInfo.XForwardedFor = getXForwardedFor(c)

	return clientInfo
}

// Private Functions
func getIPv4Address(c *fiber.Ctx) string {
	// Get ip address from request header
	ipv4 := c.IP()

	if ipv4 == "" {
		return ""
	}

	return ipv4
}

func getUserAgent(c *fiber.Ctx) string {
	userAgent := c.Get("User-Agent")
	return userAgent
}

func getLanguage(c *fiber.Ctx) string {
	language := c.Get("Accept-Language")
	return language
}

func getReferer(c *fiber.Ctx) string {
	referer := c.Get("Referer")
	return referer
}

func getMethod(c *fiber.Ctx) string {
	method := c.Method()
	return method
}

func getEncoding(c *fiber.Ctx) string {
	encoding := c.Get("Accept-Encoding")
	return encoding
}

func getMIMEType(c *fiber.Ctx) string {
	mimeType := ""
	return mimeType
}

func getXForwardedFor(c *fiber.Ctx) string {
	// Get ip address from request header
	xForwardedFor := c.Get("X-Forwarded-For")
	return xForwardedFor
}
