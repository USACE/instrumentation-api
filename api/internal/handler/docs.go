package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const localServer = "http://localhost:8080"

func (h *ApiHandler) CreateDocHtmlHandler(apidoc []byte, serverBaseUrl string, authJwtMocked bool) (echo.HandlerFunc, error) {
	var apidocJson map[string]interface{}
	if err := json.Unmarshal(apidoc, &apidocJson); err != nil {
		return nil, err
	}
	url := serverBaseUrl
	if url == "" {
		url = localServer
	}
	server := map[string]string{"url": url}
	apidocJson["servers"] = []map[string]string{server}

	newApiDoc, err := json.Marshal(apidocJson)
	if err != nil {
		return nil, err
	}
	// if running locally, prefill mock bearer token
	var authOptions string
	if authJwtMocked {
		authOptions = mockAuthOptions
	}

	htmlContent := fmt.Sprintf(htmlTmpl, newApiDoc, authOptions)
	return func(c echo.Context) error {
		return c.HTML(http.StatusOK, htmlContent)
	}, nil
}

const htmlTmpl = `<!doctype html>
<html>
  <head>
    <title>MIDAS API Reference</title>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
  </head>
  <body>
    <script
      id="api-reference"
      type="application/json"
    >%s</script>
    <script>
      var configuration = {%s theme: 'default', layout: 'classic'}
      document.getElementById('api-reference').dataset.configuration =
        JSON.stringify(configuration)
    </script>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
  </body>
</html>`

const mockAuthOptions = `authentication: {
    preferredSecurityScheme: 'Bearer',
    apiKey: {
        token: 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwibmFtZSI6IlVzZXIuQXBwbGljYXRpb25BZG1pbiIsImdpdmVuX25hbWUiOiJNb2NrIEFkbWluIiwicHJlZmVycmVkX25hbWUiOiJNb2NrIEFkbWluIiwiY2FjVUlEIjoiMiIsIng1MDlfcHJlc2VudGVkIjp0cnVlLCJpYXQiOjE1MTYyMzkwMjIsImV4cCI6MjAwMDAwMDAwMCwicm9sZXMiOltdfQ.m3XTmSQmy_T6ywoLChBXqhXEC_dMg9K-zjRM7_WbGl8',
    },
},`
