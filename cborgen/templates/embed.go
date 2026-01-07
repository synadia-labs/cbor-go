package templates

import "embed"

// FS exposes the codegen templates used by cborgen
// for per-struct encode/decode generation.
//
//go:embed *.go.tpl
var FS embed.FS

