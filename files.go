package swagger

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var SwaggerFileSystem fs.FS
func init() {
    var err error
    SwaggerFileSystem, err = fs.Sub(dist, "dist")
    if err != nil {
        log.Fatalf("Failed to initialize SwaggerFileSystem: %v", err)
    }
}
var dist embed.FS

// SwaggerFileSystem holds embedded swagger ui files
var SwaggerFileSystem, _ = fs.Sub(dist, "dist")
