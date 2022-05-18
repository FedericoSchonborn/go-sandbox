package color

type RGBA32 struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Alpha uint8
}

func NewRGBA32(red, green, blue, alpha uint8) RGBA32 {
	return RGBA32{
		Red:   red,
		Green: green,
		Blue:  blue,
		Alpha: alpha,
	}
}

func DecodeRGBA32(value uint32) RGBA32 {
	return RGBA32{
		Red:   uint8((value >> 24) & 0xFF),
		Green: uint8((value >> 16) & 0xFF),
		Blue:  uint8((value >> 8) & 0xFF),
		Alpha: uint8(value & 0xFF),
	}
}

func (c RGBA32) RGB32() RGB32 {
	return RGB32{
		Red:   c.Red,
		Green: c.Green,
		Blue:  c.Blue,
	}
}

func (c RGBA32) EqualRGB32(other RGB32) bool {
	return other.EqualRGBA32(c)
}

func (c RGBA32) EqualRGBA32(other RGBA32) bool {
	return c.Red == other.Red && c.Green == other.Green && c.Blue == other.Blue && c.Alpha == other.Alpha
}

func (c RGBA32) Encode() uint32 {
	return uint32(c.Red)<<24 | uint32(c.Green)<<16 | uint32(c.Blue)<<8 | uint32(c.Alpha)
}
