package swagger

import (
	"fmt"
	"html/template"
	"path"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/utils/v2"
	swaggerFiles "github.com/swaggo/files/v2"
	"github.com/swaggo/swag"
)

const (
	defaultDocURL = "doc.json"
	defaultIndex  = "index.html"
)

var HandlerDefault = New()

// New returns custom handler
func New(config ...Config) fiber.Handler {
	cfg := configDefault(config...)

	index, err := template.New("swagger_index.html").Parse(indexTmpl)
	if err != nil {
		panic(fmt.Errorf("fiber: swagger middleware error -> %w", err))
	}

	var (
		prefix string
		once   sync.Once
	)

	return func(c fiber.Ctx) error {
		// Set prefix
		once.Do(
			func() {
				prefix = strings.ReplaceAll(c.Route().Path, "*", "")

				forwardedPrefix := getForwardedPrefix(c)
				if forwardedPrefix != "" {
					prefix = forwardedPrefix + prefix
				}

				// Set doc url
				if len(cfg.URL) == 0 {
					cfg.URL = path.Join(prefix, defaultDocURL)
				}
			},
		)

		p := c.Path(utils.CopyString(c.Params("*")))

		switch p {
		case defaultIndex:
			c.Type("html")
			return index.Execute(c, cfg)
		case defaultDocURL:
			var doc string
			if doc, err = swag.ReadDoc(cfg.InstanceName); err != nil {
				return err
			}
			return c.Type("json").SendString(doc)
		case "", "/":
			return c.Redirect().Status(fiber.StatusMovedPermanently).To(path.Join(prefix, defaultIndex))
		default:
			return static.New("/", static.Config{
				FS:     swaggerFiles.FS,
				Browse: true,
			})(c)
		}
	}
}

func getForwardedPrefix(c fiber.Ctx) string {
	header := c.GetReqHeaders()["X-Forwarded-Prefix"]

	if len(header) == 0 {
		return ""
	}

	prefix := ""

	for _, rawPrefix := range header {
		endIndex := len(rawPrefix)
		for endIndex > 1 && rawPrefix[endIndex-1] == '/' {
			endIndex--
		}

		if endIndex != len(rawPrefix) {
			prefix += rawPrefix[:endIndex]
		} else {
			prefix += rawPrefix
		}
	}

	return prefix
}
