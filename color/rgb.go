package color

type RGB32 struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func NewRGB32(red, green, blue uint8) RGB32 {
	return RGB32{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

func DecodeRGB32(value uint32) RGB32 {
	return RGB32{
		Red:   uint8((value >> 16) & 0xFF),
		Green: uint8((value >> 8) & 0xFF),
		Blue:  uint8(value & 0xFF),
	}
}

func (c RGB32) RGBA32(alpha uint8) RGBA32 {
	return RGBA32{
		Red:   c.Red,
		Green: c.Green,
		Blue:  c.Blue,
		Alpha: alpha,
	}
}

func (c RGB32) EqualRGB32(other RGB32) bool {
	return c.Red == other.Red && c.Green == other.Green && c.Blue == other.Blue
}

func (c RGB32) EqualRGBA32(other RGBA32) bool {
	return other.EqualRGB32(c)
}

func (c RGB32) Encode() uint32 {
	return uint32(c.Red)<<16 | uint32(c.Green)<<8 | uint32(c.Blue)
}
