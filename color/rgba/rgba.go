package rgba

import "github.com/FedericoSchonborn/go-sandbox/color"

var (
	Red     = color.NewRGBA32(255, 0, 0, 255)
	Green   = color.NewRGBA32(0, 255, 0, 255)
	Blue    = color.NewRGBA32(0, 0, 255, 255)
	Cyan    = color.NewRGBA32(0, 255, 255, 255)
	Magenta = color.NewRGBA32(255, 0, 255, 255)
	Yellow  = color.NewRGBA32(255, 255, 0, 255)
	White   = color.NewRGBA32(255, 255, 255, 255)
	Black   = color.NewRGBA32(0, 0, 0, 255)
	None    = color.NewRGBA32(0, 0, 0, 0)
)
