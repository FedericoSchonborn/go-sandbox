package rgb

import "github.com/FedericoSchonborn/go-sandbox/color"

var (
	Red     = color.NewRGB32(255, 0, 0)
	Green   = color.NewRGB32(0, 255, 0)
	Blue    = color.NewRGB32(0, 0, 255)
	Cyan    = color.NewRGB32(0, 255, 255)
	Magenta = color.NewRGB32(255, 0, 255)
	Yellow  = color.NewRGB32(255, 255, 0)
	White   = color.NewRGB32(255, 255, 255)
	Black   = color.NewRGB32(0, 0, 0)
	None    = color.NewRGBA32(0, 0, 0, 0)
)
