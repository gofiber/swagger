package swagger

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/swaggo/swag"
)

type mockedSwag struct{}

func (s *mockedSwag) ReadDoc() string {
	return `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server.",
        "title": "Swagger Example API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "petstore.swagger.io",
    "basePath": "/v2",
    "paths": {}
}`
}

func Test_Swagger(t *testing.T) {
	app := fiber.New()

	swag.Register(swag.Name, &mockedSwag{})

	app.Get("/swag/*", HandlerDefault)

	tests := []struct {
		name        string
		url         string
		statusCode  int
		contentType string
	}{
		{
			name:        "Should be returns status 200 with 'text/html' content-type",
			url:         "/swag/index.html",
			statusCode:  200,
			contentType: "text/html",
		},
		{
			name:        "Should be returns status 200 with 'application/json' content-type",
			url:         "/swag/doc.json",
			statusCode:  200,
			contentType: "application/json",
		},
		{
			name:        "Should be returns status 200 with 'image/png' content-type",
			url:         "/swag/favicon-16x16.png",
			statusCode:  200,
			contentType: "image/png",
		},
		{
			name:       "Should return status 301",
			url:        "/swag/",
			statusCode: 301,
		},
		{
			name:       "Should return status 404",
			url:        "/swag/notfound",
			statusCode: 404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := app.Test(req)
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != tt.statusCode {
				t.Fatalf(`StatusCode: got %v - expected %v`, resp.StatusCode, tt.statusCode)
			}

			if tt.contentType != "" {
				ct := resp.Header.Get("Content-Type")
				if ct != tt.contentType {
					t.Fatalf(`Content-Type: got %s - expected %s`, ct, tt.contentType)
				}
			}
		})
	}
}
