package footsites

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
)

//Default headers with functionality to set the host, content type and add 1-off hard-coded cookies.
func AddHeaders(header Header, host string) http.Header {
	var x = http.Header{
		"origin": {"https://club.pokemon.com"},
	}

	if header.content != nil {
		x.Set("Content-Length", fmt.Sprint(header.content.Size()))
	}

	if len(header.additionalHeaders) > 0 {
		for _, element := range header.additionalHeaders {
			x.Set(element.key, element.value)
		}
	}

	if header.contentType == "json" {
		x.Set("content-type", "application/json")
		x.Set("accept", "application/json")
	} else if header.contentType == "multipart" {
		x.Set("content-type", "multipart/form-data; boundary=----WebKitFormBoundary45pI4iftSbnzXGQ1")

	} else {
		x.Set("content-type", "application/x-www-form-urlencoded")
	}

	if len(header.cookie) > 0 {
		buildString := ""
		for i := 0; i < len(header.cookie); i++ {
			buildString += header.cookie[i] + "; "
		}
		x.Set("Cookie", buildString+strings.Join(x.Values("Cookie"), "; "))
	}

	return x
}

//Used in multiple methods to extract key values
func ExtractValue(body, elementType, targetType, typeValue string, optionalAttribute ...string) string {
	var val = ""
	var value = "value"
	if len(optionalAttribute) > 0 {
		value = optionalAttribute[0]
	}
	_body := soup.HTMLParse(body)
	el := _body.Find(elementType, targetType, typeValue)
	if el.Error == nil {
		element := el.Pointer.Attr
		for _, v := range element {
			if v.Key == value {
				//Locate the authKey attribute value within this node
				val = v.Val
			}
		}
	}

	return val
}

//Task helpers, to loop functions and log failures to console
//In future change this so FunctionToRun and Name are in a 'function' object WIll probably build on this in future.
type TaskTemplate struct {
	functionToRun func() bool
	name          string
	complete      bool
	concurrent    bool
	delay         time.Duration
}

type SessionInfo struct {
	OutJSSRC        string
	JSTOKEN         string
	AgentEncoded    string
	LanguageEncoded string
	SVRTIME         string
	IGGY            string

	RequestAgent string
	AexOffset    string
	Browser      string
	Version      string
	OsName       string
	Appname      string
	AppPlatform  string
	Height       int
	Width        int
	AllPlugins   string
	Referer      string
	IntLoc       string
	GetOffset    int
}

func loopFunctionB(task TaskTemplate) {
	for {
		taskResult := task.functionToRun()
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(5000)
		time.Sleep((time.Duration(r) + task.delay) * time.Millisecond)

		if taskResult {
			break
		} else {
			fmt.Printf("%s task failed - retrying\n", task.name)
		}
	}
}
