package model

type ClientInfo struct {
	IPv4          string `json:"ipv4`
	IPv6          string `json:"ipv6`
	UserAgent     string `json:"userAgent`
	Language      string `json:"language`
	Referer       string `json:referer`
	Method        string `json:"method`
	Encoding      string `json:"encoding`
	MIMEType      string `json:"mimeType`
	XForwardedFor string `json:"x-forwarded-for`
}
