package footsites

import "bytes"

type SendDobRequest struct {
	Csrfmiddlewaretoken string `json:"csrftoken"`
	Dob                 string `json:"dob"`
	country             string `json:"country"`
}

type BrowserInfo struct {
	ScreenWidth    int    `json:"screenWidth"`
	ScreenHeight   int    `json:"screenHeight"`
	ColorDepth     int    `json:"colorDepth"`
	UserAgent      string `json:"userAgent"`
	TimeZoneOffset int    `json:"timeZoneOffset"`
	Language       string `json:"language"`
	JavaEnabled    bool   `json:"javaEnabled"`
}

type Header struct {
	cookie            []string
	additionalHeaders []additionalHeaders
	content           *bytes.Reader
	contentType       string
}
type additionalHeaders struct {
	key   string
	value string
}
