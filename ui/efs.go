package ui

import (
	"embed"
)

//go:embed "html" "static/css/main.css" "static/img" "static/js"
var Files embed.FS
