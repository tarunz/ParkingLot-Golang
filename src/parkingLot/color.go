package parkinglot

// Color enumeration
type Color int

const (
	White  Color = 0
	Black  Color = 1
	Blue   Color = 2
	Green  Color = 3
	Red    Color = 4
	Yellow Color = 5
	Pink   Color = 6
	Violet Color = 7
	Orange Color = 8
)

var s2c = map[string]Color{
	"White":  0,
	"Black":  1,
	"Blue":   2,
	"Green":  3,
	"Red":    4,
	"Yellow": 5,
	"Pink":   6,
	"Violet": 7,
	"Orange": 8,
}

var c2s = [...]string{
	"White",
	"Black",
	"Blue",
	"Green",
	"Red",
	"Yellow",
	"Pink",
	"Violet",
	"Orange",
}

func StringToColor(color string) Color {
	return s2c[color]
}

func (c Color) ColorToString() string {
	return c2s[c]
}
