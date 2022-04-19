package swagger

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/swagger/example/docs"
)

func Test_Swagger(t *testing.T) {
	app := *fiber.New()

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
			req, _ := http.NewRequest("GET", tt.url, nil)
			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf(`%s: %s`, t.Name(), err)
			}

			if resp.StatusCode != tt.statusCode {
				t.Fatalf(`%s: StatusCode: got %v - expected %v`, t.Name(), resp.StatusCode, tt.statusCode)
			}

			if tt.contentType != "" {
				ct := resp.Header.Get("Content-Type")
				if ct != tt.contentType {
					t.Fatalf(`%s: Content-Type: got %s - expected %s`, t.Name(), ct, tt.contentType)
				}
			}
		})
	}
}
