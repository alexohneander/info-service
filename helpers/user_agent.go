package helpers

import (
	"fmt"
	"strings"

	ua "github.com/mileusna/useragent"
)

func IsBrowser(userAgent string) bool {

	ua := ua.Parse(userAgent)
	fmt.Println()
	fmt.Println(ua.String)
	fmt.Println(strings.Repeat("=", len(ua.String)))
	fmt.Println("Name:", ua.Name, "v", ua.Version)
	fmt.Println("OS:", ua.OS, "v", ua.OSVersion)
	fmt.Println("Device:", ua.Device)
	if ua.Mobile {
		fmt.Println("(Mobile)")
	}
	if ua.Tablet {
		fmt.Println("(Tablet)")
	}
	if ua.Desktop {
		fmt.Println("(Desktop)")
	}
	if ua.Bot {
		fmt.Println("(Bot)")
	}
	if ua.URL != "" {
		fmt.Println(ua.URL)
	}

	return true
}
