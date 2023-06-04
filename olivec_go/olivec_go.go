package olivec_go

import "unsafe"

const OLIVEC_AA_RES = 2

func OlivecSwap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func condToInt(condition bool) int {
	if condition {
		return 1
	} else {
		return 0
	}
}

func OlivecSign(x int) int {
	return condToInt(x > 0) - condToInt(x < 0)
}

func OlivecAbs(x int) int {
	return OlivecSign(x) * x
}

const OlivecDefaultFontHeight = 6
const OlivecDefaultFontWidth = 6

type OlivecFont struct {
	width, height uint
	glyphs        *[128][OlivecDefaultFontHeight][OlivecDefaultFontWidth]byte
}

var OlivecDefaultGlyphs = [128][OlivecDefaultFontHeight][OlivecDefaultFontWidth]byte{
	'a': {
		{0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
	},
	'b': {
		{1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0},

		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 0, 0},
	},
	'c': {
		{0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'd': {

		{0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
	},

	'e': {
		{0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},

		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
	},
	'f': {
		{0, 0, 1, 1, 0},
		{0, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	},
	'g': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'h': {

		{1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
	},

	'i': {
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},

		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
	},
	'j': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'k': {
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 0, 0},
		{0, 1, 1, 0, 0},

		{0, 1, 0, 1, 0},
	},
	'l': {

		{0, 1, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 1, 1, 0},
	},

	'm': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'n': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'o': {
		{0, 0, 0, 0, 0},

		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'p': {
		{1, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 0, 0},

		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
	},
	'q': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'r': {
		{0, 0, 0, 0, 0},
		{1, 0, 1, 1, 0},
		{1, 1, 0, 0, 1},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
	},
	's': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	't': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'u': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'v': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'w': {
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},

		{0, 1, 1, 1, 1},
	},
	'x': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'y': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'z': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},

	'A': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'B': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'C': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'D': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'E': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'F': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'G': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'H': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'I': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'J': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'K': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'L': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'M': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'N': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'O': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'P': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'Q': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'R': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'S': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'T': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'U': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'V': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'W': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'X': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'Y': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},
	'Z': {
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
	},

	'0': {
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},

		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'1': {
		{0, 0, 1, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 1, 1, 0},
	},
	'2': {
		{0, 1, 1, 0, 0},

		{1, 0, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0},
	},
	'3': {
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},

		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'4': {
		{0, 0, 1, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 1, 1},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0},
	},
	'5': {
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},

		{1, 1, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'6': {
		{0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},

		{0, 1, 1, 0, 0},
	},
	'7': {

		{1, 1, 1, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	},

	'8': {
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},

		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'9': {
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},

	',': {
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},

		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0},
	},

	'.': {
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
	},
	'-': {
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},

		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	},
}

var OlivecDefaultFont = OlivecFont{
	glyphs: &OlivecDefaultGlyphs,
	width:  OlivecDefaultFontWidth,
	height: OlivecDefaultFontHeight,
}

type OlivecCanvas struct {
	pixels *[]uint32
	width  uint
	height uint
	stride uint
}

func OlivecCanvasNull() *OlivecCanvas {
	return &OlivecCanvas{
		pixels: nil,
	}
}

func OlivecPixel(oc *OlivecCanvas, x, y uint) *uint32 {
	return &((*oc.pixels)[y*oc.stride+x])
}

type OlivecNormalizedRect struct {
	x1, x2 int
	y1, y2 int

	ox1, ox2 int
	oy1, oy2 int
}

func OlivecCanvasNew(pixels *[]uint32, width, height, stride uint) *OlivecCanvas {
	return &OlivecCanvas{
		pixels: pixels,
		width:  width,
		height: height,
		stride: stride,
	}
}

func OlivecNormalizeRect(x, y, w, h int, canvas_width, canvas_height uint, nr *OlivecNormalizedRect) bool {
	if w == 0 {
		return false
	}
	if h == 0 {
		return false
	}

	nr.ox1 = x
	nr.oy1 = y

	nr.ox2 = nr.ox1 + OlivecSign(w)*(OlivecAbs(w)-1)
	if nr.ox1 > nr.ox2 {
		OlivecSwap(&nr.ox1, &nr.ox2)
	}
	nr.oy2 = nr.oy1 + OlivecSign(h)*(OlivecAbs(h)-1)
	if nr.oy1 > nr.oy2 {
		OlivecSwap(&nr.oy1, &nr.oy2)
	}

	if nr.ox1 >= int(canvas_width) {
		return false
	}
	if nr.ox2 < 0 {
		return false
	}
	if nr.oy1 >= int(canvas_height) {
		return false
	}
	if nr.oy2 < 0 {
		return false
	}

	nr.x1 = nr.ox1
	nr.y1 = nr.oy1
	nr.x2 = nr.ox2

	nr.y2 = nr.oy2

	if nr.x1 < 0 {
		nr.x1 = 0
	}
	if nr.x2 >= int(canvas_width) {
		nr.x2 = int(canvas_width) - 1
	}

	if nr.y1 < 0 {
		nr.y1 = 0
	}
	if nr.y2 >= int(canvas_height) {
		nr.y2 = int(canvas_height) - 1
	}

	return true
}

func OlivecSubcanvas(oc *OlivecCanvas, x, y, w, h int) *OlivecCanvas {
	nr := OlivecNormalizedRect{}
	if !OlivecNormalizeRect(x, y, w, h, oc.width, oc.height, &nr) {
		return OlivecCanvasNull()
	}

	oc.pixels = &[]uint32{
		*OlivecPixel(oc, uint(nr.x1), uint(nr.y1)),
	}
	oc.width = uint(nr.x2 - nr.x1 + 1)
	oc.height = uint(nr.y2 - nr.y1 + 1)
	return oc
}

func OlivecRed(color uint32) uint32 {
	return (color & 0x000000FF) >> (8 * 0)
}

func OlivecGreen(color uint32) uint32 {
	return (color & 0x0000FF00) >> (8 * 1)
}

func OlivecBlue(color uint32) uint32 {
	return (color & 0x00FF0000) >> (8 * 2)
}

func OlivecAlpha(color uint32) uint32 {
	return (color & 0xFF000000) >> (8 * 3)
}

func OlivecRgba(r, g, b, a uint32) uint32 {
	return ((r & 0xFF) << (8 * 0)) | ((g & 0xFF) << (8 * 1)) | ((b & 0xFF) << (8 * 2)) | ((a & 0xFF) << (8 * 3))
}

func OlivecBlendColor(c1 *uint32, c2 uint32) {
	r1 := OlivecRed(*c1)
	g1 := OlivecGreen(*c1)
	b1 := OlivecBlue(*c1)
	a1 := OlivecAlpha(*c1)

	r2 := OlivecRed(c2)
	g2 := OlivecGreen(c2)
	b2 := OlivecBlue(c2)
	a2 := OlivecAlpha(c2)

	r1 = (r1*(255-a2) + r2*a2) / 255
	if r1 > 255 {
		r1 = 255
	}
	g1 = (g1*(255-a2) + g2*a2) / 255
	if g1 > 255 {
		g1 = 255
	}
	b1 = (b1*(255-a2) + b2*a2) / 255
	if b1 > 255 {
		b1 = 255
	}

	*c1 = OlivecRgba(r1, g1, b1, a1)
}

func OlivecFill(oc *OlivecCanvas, color uint32) {
	for y := 0; y < int(oc.height); y++ {
		for x := 0; x < int(oc.width); x++ {
			*OlivecPixel(oc, uint(x), uint(y)) = color
		}
	}
}

func OlivecRect(oc *OlivecCanvas, x, y, w, h int, color uint32) {
	nr := &OlivecNormalizedRect{}
	if !OlivecNormalizeRect(x, y, w, h, oc.width, oc.height, nr) {
		return
	}
	for x := nr.x1; x <= nr.x2; x++ {
		for y := nr.y1; y <= nr.y2; y++ {
			OlivecBlendColor(OlivecPixel(oc, uint(x), uint(y)), color)
		}
	}
}

func OlivecFrame(oc *OlivecCanvas, x, y, w, h int, t uint, color uint32) {
	if t == 0 {
		return
	}

	x1 := x
	y1 := y
	x2 := x1 + OlivecSign(w)*(OlivecAbs(w)-1)
	if x1 > x2 {
		OlivecSwap(&x1, &x2)
	}
	y2 := y1 + OlivecSign(h)*(OlivecAbs(h)-1)
	if y1 > y2 {
		OlivecSwap(&y1, &y2)
	}

	intT := int(t)
	OlivecRect(oc, x1-intT/2, y1-intT/2, (x2-x1+1)+intT/2*2, intT, color)
	OlivecRect(oc, x1-intT/2, y1-intT/2, intT, (y2-y1+1)+intT/2*2, color)
	OlivecRect(oc, x1-intT/2, y2+intT/2, (x2-x1+1)+intT/2*2, -intT, color)
	OlivecRect(oc, x2+intT/2, y1-intT/2, -intT, (y2-y1+1)+intT/2*2, color)
}

func OlivecEclipse(oc *OlivecCanvas, cx, cy, rx, ry int, color uint32) {
	nr := &OlivecNormalizedRect{}
	rx1 := rx + OlivecSign(rx)
	ry1 := ry + OlivecSign(ry)
	if !OlivecNormalizeRect(cx-rx1, cy-ry1, 2*rx1, 2*ry1, oc.width, oc.height, nr) {
		return
	}

	for y := nr.y1; y <= nr.y2; y++ {
		for x := nr.x1; x <= nr.x2; x++ {
			nx := (float32(x) + 0.5 - float32(nr.x1)) / (2.0 * float32(rx1))
			ny := (float32(y) + 0.5 - float32(nr.y1)) / (2.0 * float32(ry1))
			dx := nx - 0.5
			dy := ny - 0.5
			if dx*dx+dy*dy <= 0.5*0.5 {
				*OlivecPixel(oc, uint(x), uint(y)) = color
			}
		}
	}
}

func OlivecCircle(oc *OlivecCanvas, cx, cy, r int, color uint32) {
	nr := &OlivecNormalizedRect{}
	r1 := r + OlivecSign(r)
	if !OlivecNormalizeRect(cx-r1, cy-r1, 2*r1, 2*r1, oc.width, oc.height, nr) {
		return
	}

	for y := nr.y1; y <= nr.y2; y++ {
		for x := nr.x1; x <= nr.x2; x++ {
			count := 0
			for sox := 0; sox < OLIVEC_AA_RES; sox++ {
				for soy := 0; soy < OLIVEC_AA_RES; soy++ {
					res1 := OLIVEC_AA_RES + 1
					dx := (x*res1*2 + 2 + sox*2 - res1*cx*2 - res1)
					dy := (y*res1*2 + 2 + soy*2 - res1*cy*2 - res1)
					if dx*dx+dy*dy <= res1*res1*r*r*2*2 {
						count++
					}
				}
			}
			alpha := ((color & 0xFF000000) >> (3 * 8)) * uint32(count) / OLIVEC_AA_RES / OLIVEC_AA_RES
			updated_color := (color & 0x00FFFFFF) | (alpha << (3 * 8))
			OlivecBlendColor(OlivecPixel(oc, uint(x), uint(y)), updated_color)
		}
	}
}

func OlivecInBounds(oc *OlivecCanvas, x, y int) bool {
	return 0 <= x && x < int(oc.width) && 0 <= y && y <= int(oc.height)
}

func OlivecLine(oc *OlivecCanvas, x1, y1, x2, y2 int, color uint32) {
	dx := x2 - x1
	dy := y2 - y1

	if dx == 0 && dy == 0 {
		if OlivecInBounds(oc, x1, y2) {
			OlivecBlendColor(OlivecPixel(oc, uint(x1), uint(y1)), color)
		}
		return
	}

	if OlivecAbs(dx) > OlivecAbs(dy) {
		if x1 > x2 {
			OlivecSwap(&x1, &x2)
			OlivecSwap(&y1, &y2)
		}

		for x := x1; x <= x2; x++ {
			y := dy*(x-x1)/dx + y1
			if OlivecInBounds(oc, x, y) {
				OlivecBlendColor(OlivecPixel(oc, uint(x), uint(y)), color)
			}
		}
	} else {
		if y1 > y2 {
			OlivecSwap(&x1, &x2)
			OlivecSwap(&y1, &y2)
		}

		for y := y1; y <= y2; y++ {
			x := dx*(y-y1)/dy + x1
			if OlivecInBounds(oc, x, y) {
				OlivecBlendColor(OlivecPixel(oc, uint(x), uint(y)), color)
			}
		}
	}
}

func MixColors2(c1, c2 uint32, u1, det int) uint32 {
	r1 := OlivecRed(c1)
	g1 := OlivecGreen(c1)
	b1 := OlivecBlue(c1)
	a1 := OlivecAlpha(c1)

	r2 := OlivecRed(c2)
	g2 := OlivecGreen(c2)
	b2 := OlivecBlue(c2)
	a2 := OlivecAlpha(c2)

	if det != 0 {
		u2 := det - u1
		r4 := (int(r1)*u2 + int(r2)*u1) / det
		g4 := (int(g1)*u2 + int(g2)*u1) / det
		b4 := (int(b1)*u2 + int(b2)*u1) / det
		a4 := (int(a1)*u2 + int(a2)*u1) / det

		return OlivecRgba(uint32(r4), uint32(g4), uint32(b4), uint32(a4))
	}

	return 0
}

func MixColors3(c1, c2, c3 uint32, u1, u2, det int) uint32 {
	r1 := OlivecRed(c1)
	g1 := OlivecGreen(c1)
	b1 := OlivecBlue(c1)
	a1 := OlivecAlpha(c1)

	r2 := OlivecRed(c2)
	g2 := OlivecGreen(c2)
	b2 := OlivecBlue(c2)
	a2 := OlivecAlpha(c2)

	r3 := OlivecRed(c3)
	g3 := OlivecGreen(c3)
	b3 := OlivecBlue(c3)
	a3 := OlivecAlpha(c3)

	if det != 0 {
		u3 := det - u1 - u2
		r4 := (int(r1)*u2 + int(r2)*u1 + int(r3)*u3) / det
		g4 := (int(g1)*u2 + int(g2)*u1 + int(g3)*u3) / det
		b4 := (int(b1)*u2 + int(b2)*u1 + int(b3)*u3) / det
		a4 := (int(a1)*u2 + int(a2)*u1 + int(a3)*u3) / det

		return OlivecRgba(uint32(r4), uint32(g4), uint32(b4), uint32(a4))
	}

	return 0
}

func OlivecBarycentric(x1, y1, x2, y2, x3, y3, xp, yp int, u1, u2, det *int) bool {
	*det = ((x1-x3)*(y2-y3) - (x2-x3)*(y1-y3))
	*u1 = ((y2-y3)*(xp-x3) + (x3-x2)*(yp-y3))
	*u2 = ((y3-y1)*(xp-x3) + (x1-x3)*(yp-y3))
	u3 := *det - *u1 - *u2

	return (OlivecSign(*u1) == OlivecSign(*det) || *u1 == 0) && (OlivecSign(*u2) == OlivecSign(*det) || *u2 == 0) && (OlivecSign(u3) == OlivecSign(*det) || u3 == 0)
}

func OlivecNormalizeTriangle(width, height uint, x1, y1, x2, y2, x3, y3 int, lx, hx, ly, hy *int) bool {
	*lx = x1
	*hx = x1
	if *lx > x2 {
		*lx = x2
	}
	if *lx > x3 {
		*lx = x3
	}
	if *hx < x2 {
		*hx = x2
	}
	if *hx < x3 {
		*hx = x3
	}
	if *lx < 0 {
		*lx = 0
	}
	if uint(*lx) >= width {
		return false
	}
	if *hx < 0 {
		return false
	}
	if uint(*hx) >= width {
		*hx = int(width) - 1
	}

	*ly = y1
	*hy = y1
	if *ly > y2 {
		*ly = y2
	}
	if *ly > y3 {
		*ly = y3
	}
	if *hy < y2 {
		*hy = y2
	}

	if *hy < y3 {
		*hy = y3
	}
	if *ly < 0 {
		*ly = 0
	}
	if uint(*ly) >= height {
		return false
	}
	if *hy < 0 {
		return false
	}
	if uint(*hy) >= height {
		*hy = int(height) - 1
	}

	return true
}

func OlivecTriangle3c(oc *OlivecCanvas, x1, y1, x2, y2, x3, y3 int, c1, c2, c3 uint32) {
	var lx, hx, ly, hy int
	if OlivecNormalizeTriangle(oc.width, oc.height, x1, y1, x2, y2, x3, y3, &lx, &hx, &ly, &hy) {
		for y := ly; y <= hy; y++ {
			for x := lx; x <= hx; x++ {
				var u1, u2, det int
				if OlivecBarycentric(x1, y1, x2, y2, x3, y3, x, y, &u1, &u2, &det) {
					OlivecBlendColor(OlivecPixel(oc, uint(x), uint(y)), MixColors3(c1, c2, c3, u1, u2, det))
				}
			}
		}
	}
}

func OlivecTriangle3z(oc *OlivecCanvas, x1, y1, x2, y2, x3, y3 int, z1, z2, z3 float32) {
	var lx, hx, ly, hy int
	if OlivecNormalizeTriangle(oc.width, oc.height, x1, y1, x2, y2, x3, y3, &lx, &hx, &ly, &hy) {
		for y := ly; y <= hy; y++ {
			for x := lx; x <= hy; x++ {
				var u1, u2, det int
				if OlivecBarycentric(x1, y1, x2, y2, x3, y3, x, y, &u1, &u2, &det) {
					z := z1*float32(u1)/float32(det) + z2*float32(u2)/float32(det) + z3*(float32(det)-float32(u1)-float32(u2))/float32(det)
					*OlivecPixel(oc, uint(x), uint(y)) = uint32(z)
				}
			}
		}
	}
}

func OlivecTriangle3uv(oc *OlivecCanvas, x1, y1, x2, y2, x3, y3 int, tx1, ty1, tx2, ty2, tx3, ty3, z1, z2, z3 float32, texture *OlivecCanvas) {
	var lx, hx, ly, hy int
	if OlivecNormalizeTriangle(oc.width, oc.height, x1, y1, x2, y2, x3, y3, &lx, &hx, &ly, &hy) {
		for y := ly; y <= hy; y++ {
			for x := lx; x <= hx; x++ {
				var u1, u2, det int
				if OlivecBarycentric(x1, y1, x2, y2, x3, y3, x, y, &u1, &u2, &det) {
					u3 := det - u1 - u2
					z := z1*float32(u1)/float32(det) + z2*float32(u2)/float32(det) + z3*(float32(det-u1-u2))/float32(det)
					tx := tx1*float32(u1)/float32(det) + tx2*float32(u2)/float32(det) + tx3*float32(u3)/float32(det)
					ty := ty1*float32(u1)/float32(det) + ty2*float32(u2)/float32(det) + ty3*float32(u3)/float32(det)

					texture_x := int(tx / z * float32(texture.width))
					if texture_x < 0 {
						texture_x = 0
					}
					if uint(texture_x) >= texture.width {
						texture_x = int(texture.width) - 1
					}

					texture_y := int(ty / z * float32(texture.height))
					if texture_y < 0 {
						texture_y = 0
					}
					if uint(texture_y) >= texture.height {
						texture_y = int(texture.height) - 1
					}
					*OlivecPixel(oc, uint(x), uint(y)) = *OlivecPixel(texture, uint(texture_x), uint(texture_y))
				}
			}
		}
	}
}

func OlivecTriangle3uvBilinear(oc *OlivecCanvas, x1, y1, x2, y2, x3, y3 int, tx1, ty1, tx2, ty2, tx3, ty3, z1, z2, z3 float32, texture *OlivecCanvas) {
	var lx, hx, ly, hy int
	if OlivecNormalizeTriangle(oc.width, oc.height, x1, y1, x2, y2, x3, y3, &lx, &hx, &ly, &hy) {
		for y := ly; y <= hy; y++ {
			for x := lx; x <= hx; x++ {
				var u1, u2, det int
				if OlivecBarycentric(x1, y1, x2, y2, x3, y3, x, y, &u1, &u2, &det) {
					u3 := det - u1 - u2
					z := z1*float32(u1)/float32(det) + z2*float32(u2)/float32(det) + z3*(float32(det-u1-u2))/float32(det)
					tx := tx1*float32(u1)/float32(det) + tx2*float32(u2)/float32(det) + tx3*float32(u3)/float32(det)
					ty := ty1*float32(u1)/float32(det) + ty2*float32(u2)/float32(det) + ty3*float32(u3)/float32(det)

					texture_x := tx / z * float32(texture.width)
					if texture_x < 0 {
						texture_x = 0
					}
					if texture_x >= float32(texture.width) {
						texture_x = float32(texture.width) - 1
					}

					texture_y := ty / z * float32(texture.height)
					if texture_y < 0 {
						texture_y = 0
					}
					if texture_y >= float32(texture.height) {
						texture_y = float32(texture.height) - 1
					}
					precision := 100
					*OlivecPixel(oc, uint(x), uint(y)) = OlivecPixelBilinear(texture, int(texture_x)*precision, int(texture_y)*precision, precision, precision)
				}
			}
		}
	}
}

func OlivecTriangle(oc *OlivecCanvas, x1, y1, x2, y2, x3, y3 int, color uint32) {
	var lx, hx, ly, hy int
	for y := ly; y <= hy; y++ {
		for x := lx; x <= hx; x++ {
			var u1, u2, det int
			if OlivecBarycentric(x1, y1, x2, y2, x3, y3, x, y, &u1, &u2, &det) {
				OlivecBlendColor(OlivecPixel(oc, uint(x), uint(y)), color)
			}
		}
	}
}

func OlivecText(oc *OlivecCanvas, text string, tx, ty int, font *OlivecFont, glyphSize uint, color uint32) {
	for i := 0; text[i] != 0; i++ {
		gx := tx + i*int(font.width)*int(glyphSize)
		gy := ty

		firstCharAscii := int(text[0])
		sizeofChar := unsafe.Sizeof(byte(0))
		glyph := &font.glyphs[firstCharAscii*int(sizeofChar)*int(font.width)*int(font.height)]
		for dy := 0; uint(dy) < font.height; dy++ {
			for dx := 0; uint(dx) < font.width; dx++ {
				px := gx + dx*int(glyphSize)
				py := gy + dy*int(glyphSize)
				if 0 <= px && px < int(oc.width) && 0 <= py && py < int(oc.height) {
					// FIXME: remove [0] from glyph array access
					if glyph[dy*int(font.width)+dx][0] != 0 {
						OlivecRect(oc, px, py, int(glyphSize), int(glyphSize), color)
					}
				}
			}
		}
	}
}

func OlivecSpriteBlend(oc *OlivecCanvas, x, y, w, h int, sprite *OlivecCanvas) {
	if sprite.width == 0 {
		return
	}
	if sprite.height == 0 {
		return
	}

	nr := &OlivecNormalizedRect{}
	if !OlivecNormalizeRect(x, y, w, h, oc.width, oc.height, nr) {
		return
	}

	xa := nr.ox1
	if w < 0 {
		xa = nr.ox2
	}
	ya := nr.oy1
	if h < 0 {
		ya = nr.oy2
	}
	for y := nr.y1; y <= nr.y2; y++ {
		for x := nr.x1; x <= nr.x2; x++ {
			nx := uint((x - xa) * (int(sprite.width)) / w)
			ny := uint((y - ya) * (int(sprite.height)) / h)
			OlivecBlendColor(OlivecPixel(oc, uint(x), uint(y)), *OlivecPixel(sprite, nx, ny))
		}
	}
}

func OlivecSpriteCopy(oc *OlivecCanvas, x, y, w, h int, sprite *OlivecCanvas) {
	if sprite.width == 0 {
		return
	}
	if sprite.height == 0 {
		return
	}

	nr := &OlivecNormalizedRect{}
	if !OlivecNormalizeRect(x, y, w, h, oc.width, oc.height, nr) {
		return
	}

	xa := nr.ox1
	if w < 0 {
		xa = nr.ox2
	}
	ya := nr.oy1
	if h < 0 {
		ya = nr.oy2
	}
	for y := nr.y1; y <= nr.y2; y++ {
		for x := nr.x1; x <= nr.x2; x++ {
			nx := uint((x - xa) * (int(sprite.width)) / w)
			ny := uint((y - ya) * (int(sprite.height)) / h)
			*OlivecPixel(oc, uint(x), uint(y)) = *OlivecPixel(sprite, nx, ny)
		}
	}
}

func OlivecPixelBilinear(sprite *OlivecCanvas, nx, ny, w, h int) uint32 {
	px := nx % w
	py := ny % h

	x1 := nx / w
	x2 := nx / w
	y1 := ny / h
	y2 := ny / h
	if px < w/2 {
		px += w / 2
		x1 -= 1
		if x1 < 0 {
			x1 = 0
		}
	} else {
		px -= w / 2
		x2 += 1
		if uint(x2) >= sprite.width {
			x2 = int(sprite.width) - 1
		}
	}

	if py < h/2 {
		py += h / 2
		y1 -= 1
		if y1 < 0 {
			y1 = 0
		}

	} else {
		py -= h / 2
		y2 += 1
		if uint(y2) >= sprite.height {
			y2 = int(sprite.height) - 1
		}

	}

	return MixColors2(MixColors2(*OlivecPixel(sprite, uint(x1), uint(y1)),
		*OlivecPixel(sprite, uint(x2), uint(y1)),
		px, w),

		MixColors2(*OlivecPixel(sprite, uint(x1), uint(y2)),
			*OlivecPixel(sprite, uint(x2), uint(y2)),
			px, w),
		py, h)
}

func OlivecSpriteCopyBilinear(oc *OlivecCanvas, x, y, w, h int, sprite *OlivecCanvas) {
	if w <= 0 {
		return
	}
	if h <= 0 {
		return
	}

	nr := &OlivecNormalizedRect{}
	if !OlivecNormalizeRect(x, y, w, h, oc.width, oc.height, nr) {
		return
	}

	for y := nr.y1; y <= nr.y2; y++ {
		for x := nr.x1; x <= nr.x2; x++ {
			nx := uint(x-nr.ox1) * sprite.width
			ny := uint(y-nr.oy1) * sprite.height
			*OlivecPixel(oc, uint(x), uint(y)) = OlivecPixelBilinear(sprite, int(nx), int(ny), w, h)
		}
	}
}
