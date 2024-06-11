package info

import (
	"testing"

	"github.com/alexohneander/info-service/model"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func TestGetAllClientInfos(t *testing.T) {
	// Mock Fiber Context
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})

	// Headers
	c.Request().Header.Add("X-Forwarded-For", "42.42.42.42")
	c.Request().Header.Add("User-Agent", "fasthttp")
	c.Request().Header.Add("Accept-Language", "de")

	got := GetAllClientInfos(c)

	want := model.ClientInfo{
		IPv4:          "0.0.0.0",
		UserAgent:     "fasthttp",
		Language:      "de",
		Method:        "GET",
		XForwardedFor: "42.42.42.42",
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
