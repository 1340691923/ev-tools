package frontend

import (
	"embed"
)

//go:embed dist/index.html
//go:embed dist/assets/**
var StatisFs embed.FS
