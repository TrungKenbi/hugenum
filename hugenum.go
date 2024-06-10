package hugenum

import (
	"fmt"
	"math"
	"strconv"
)

const (
	maxMagnitude = 12 // max power magnitude diff for operands.
	tenCubed     = 1e3
)

// powTenToName is a map containing the names for powers of 10 (only multiples of 3 for Engineering Notation).
var powTenToName = map[int]string{
	0:   "",
	3:   "thousand",
	6:   "million",
	9:   "billion",
	12:  "trillion",
	15:  "quadrillion",
	18:  "quintillion",
	21:  "sextillion",
	24:  "septillion",
	27:  "octillion",
	30:  "nonillion",
	33:  "decillion",
	36:  "undecillion",
	39:  "duodecillion",
	42:  "tredecillion",
	45:  "quattuordecillion",
	48:  "quindecillion",
	51:  "sedecillion",
	54:  "septendecillion",
	57:  "octodecillion",
	60:  "novendecillion",
	63:  "vigintillion",
	66:  "unvigintillion",
	69:  "duovigintillion",
	72:  "tresvigintillion",
	75:  "quattuorvigintillion",
	78:  "quinvigintillion",
	81:  "sesvigintillion",
	84:  "septemvigintillion",
	87:  "octovigintillion",
	90:  "novemvigintillion",
	93:  "trigintillion",
	96:  "untrigintillion",
	99:  "duotrigintillion",
	102: "trestrigintillion",
	105: "quattuortrigintillion",
	108: "quintrigintillion",
	111: "sestrigintillion",
	114: "septentrigintillion",
	117: "octotrigintillion",
	120: "noventrigintillion",
	123: "quadragintillion",
	126: "unquadragintillion",
	129: "duoquadragintillion",
	132: "tresquadragintillion",
	135: "quattuorquadragintillion",
	138: "quinquadragintillion",
	141: "sesquadragintillion",
	144: "septenquadragintillion",
	147: "octoquadragintillion",
	150: "novenquadragintillion",
	153: "quinquagintillion",
	156: "unquinquagintillion",
	159: "duoquinquagintillion",
	162: "tresquinquagintillion",
	165: "quattuorquinquagintillion",
	168: "quinquinquagintillion",
	171: "sesquinquagintillion",
	174: "septenquinquagintillion",
	177: "octoquinquagintillion",
	180: "novenquinquagintillion",
	183: "sexagintillion",
	186: "unsexagintillion",
	189: "duosexagintillion",
	192: "tresexagintillion",
	195: "quattuorsexagintillion",
	198: "quinsexagintillion",
	201: "sesexagintillion",
	204: "septensexagintillion",
	207: "octosexagintillion",
	210: "novensexagintillion",
	213: "septuagintillion",
	216: "unseptuagintillion",
	219: "duoseptuagintillion",
	222: "treseptuagintillion",
	225: "quattuorseptuagintillion",
	228: "quinseptuagintillion",
	231: "seseptuagintillion",
	234: "septenseptuagintillion",
	237: "octoseptuagintillion",
	240: "novenseptuagintillion",
	243: "octogintillion",
	246: "unoctogintillion",
	249: "duooctogintillion",
	252: "tresoctogintillion",
	255: "quattuoroctogintillion",
	258: "quinoctogintillion",
	261: "sexoctogintillion",
	264: "septemoctogintillion",
	267: "octooctogintillion",
	270: "novemoctogintillion",
	273: "nonagintillion",
	276: "unnonagintillion",
	279: "duononagintillion",
	282: "trenonagintillion",
	285: "quattuornonagintillion",
	288: "quinnonagintillion",
	291: "senonagintillion",
	294: "septenonagintillion",
	297: "octononagintillion",
	300: "novenonagintillion",
	303: "centillion",
	306: "uncentillion",
	309: "duocentillion",
	312: "trescentillion",
	315: "quattuorcentillion",
	318: "quincentillion",
	321: "sexcentillion",
	324: "septencentillion",
	327: "octocentillion",
	330: "novencentillion",
	333: "decicentillion",
}

// BigNum represents a big number.
type BigNum struct {
	value float64
	exp   int
}

// normalize normalizes a number to engineering notation.
func (b *BigNum) normalize() {
	if b.value < 1 && b.exp != 0 {
		// e.g., 0.1E6 is converted to 100E3 ([0.1, 6] = [100, 3])
		b.value *= tenCubed
		b.exp -= 3
	} else if b.value >= tenCubed {
		// e.g., 10000E3 is converted to 10E6 ([10000, 3] = [10, 6])
		for b.value >= tenCubed {
			b.value /= tenCubed
			b.exp += 3
		}
	} else if b.value <= 0 {
		b.exp = 0
		b.value = 0
	}
}

// align computes the equivalent number at 1.Eexp (note: assumes exp is greater than b.exp).
func (b *BigNum) align(exp int) {
	d := exp - b.exp
	if d > 0 {
		if d <= maxMagnitude {
			b.value /= math.Pow(10, float64(d))
		} else {
			b.value = 0
		}
		b.exp = exp
	}
}

// Add adds a number to b.
func (b *BigNum) Add(other *BigNum) {
	if other.exp < b.exp {
		other.align(b.exp)
	} else {
		b.align(other.exp)
	}
	b.value += other.value
	b.normalize()
}

// Subtract subtracts a number from b.
func (b *BigNum) Subtract(other *BigNum) {
	if other.exp < b.exp {
		other.align(b.exp)
	} else {
		b.align(other.exp)
	}
	b.value -= other.value
	b.normalize()
}

// Multiply multiplies b by a factor.
func (b *BigNum) Multiply(factor float64) {
	// We do not support negative numbers.
	if factor >= 0 {
		b.value *= factor
		b.normalize()
	}
}

// Divide divides b by a divisor.
func (b *BigNum) Divide(divisor float64) {
	if divisor > 0 {
		b.value /= divisor
		b.normalize()
	}
}

// GetValue returns the number value as a string with the specified precision.
func (b *BigNum) GetValue(precision int) string {
	if precision <= 0 {
		precision = 3
	}
	return strconv.FormatFloat(b.value, 'f', precision, 64)
}

// GetExpName returns the exponent name as a string.
func (b *BigNum) GetExpName() string {
	return powTenToName[b.exp]
}

// GetExp returns the exponent as a string.
func (b *BigNum) GetExp() int {
	return b.exp
}

// String returns a string representation of the BigNum.
func (b *BigNum) String() string {
	expName := powTenToName[b.exp]
	if expName == "" {
		return strconv.FormatFloat(b.value, 'f', 3, 64)
	}
	return fmt.Sprintf("%g %s", b.value, expName)
}
