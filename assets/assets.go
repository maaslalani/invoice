package assets

import (
	_ "embed"
)

//go:embed "fonts/Inter Variable/Inter.ttf"
var InterFont []byte

//go:embed "fonts/Inter Hinted for Windows/Desktop/Inter-Bold.ttf"
var InterBoldFont []byte
