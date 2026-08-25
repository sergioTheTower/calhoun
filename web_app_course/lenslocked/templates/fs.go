// Package templates embeds the .gohtml files and any other files needed for the application.
package templates

import "embed"

//go:embed *
var FS embed.FS
