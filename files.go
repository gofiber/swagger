package swagger

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var dist embed.FS

// SwaggerFileSystem holds embedded swagger ui files
var SwaggerFileSystem, _ = fs.Sub(dist, "dist")
