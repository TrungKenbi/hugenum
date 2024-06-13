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
var powTenToName = map[int32]string{
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

// HugeNum represents a big number.
type HugeNum struct {
	Value float64
	Exp   int32
}

func New(value float64, exp int32) *HugeNum {
	return &HugeNum{
		Value: value,
		Exp:   exp,
	}
}

func NewExp0(value float64) *HugeNum {
	return New(value, 0)
}

// normalize normalizes a number to engineering notation.
func (h *HugeNum) normalize() {
	if h.Value == 0 {
		h.Exp = 0
		return
	}

	// Handle negative values
	sign := 1.0
	if h.Value < 0 {
		sign = -1.0
		h.Value = -h.Value
	}

	// Adjust value and exponent to get value in range [1, 1000)
	for h.Value >= tenCubed {
		h.Value /= tenCubed
		h.Exp += 3
	}
	for h.Value < 1 {
		h.Value *= tenCubed
		h.Exp -= 3
	}

	// Adjust exponent to be a multiple of 3
	remainder := h.Exp % 3
	h.Value *= math.Pow10(int(remainder))
	h.Exp -= remainder

	// Restore sign
	h.Value *= sign
}

// align computes the equivalent number at 1.Eexp (note: assumes Exp is greater than b.exp).
func (h *HugeNum) align(exp int32) {
	d := exp - h.Exp
	if d > 0 {
		if d <= maxMagnitude {
			h.Value /= math.Pow(10, float64(d))
		} else {
			h.Value = 0
		}
		h.Exp = exp
	}
}

// Add adds a number to b.
func (h *HugeNum) Add(other *HugeNum) {
	if other.Exp < h.Exp {
		other.align(h.Exp)
	} else {
		h.align(other.Exp)
	}
	h.Value += other.Value
	h.normalize()
}

// Subtract subtracts a number from b.
func (h *HugeNum) Subtract(other *HugeNum) {
	if other.Exp < h.Exp {
		other.align(h.Exp)
	} else {
		h.align(other.Exp)
	}
	h.Value -= other.Value
	h.normalize()
}

// Multiply multiplies two HugeNum instances.
func (h *HugeNum) Multiply(other *HugeNum) {
	h.Value *= other.Value
	h.Exp += other.Exp
	h.normalize()
}

// MultiplyFactor multiplies b by a factor.
func (h *HugeNum) MultiplyFactor(factor float64) {
	// We do not support negative numbers.
	if factor >= 0 {
		h.Value *= factor
		h.normalize()
	}
}

// Divide divides b by a divisor.
func (h *HugeNum) Divide(divisor float64) {
	if divisor > 0 {
		h.Value /= divisor
		h.normalize()
	}
}

// PowTen raises b to the power of 10^n.
func (h *HugeNum) PowTen(n int32) {
	h.Exp += n
	h.normalize()
}

// ExpName returns the exponent name as a string.
func (h *HugeNum) ExpName() string {
	return powTenToName[h.Exp]
}

// String returns a string representation of the HugeNum.
func (h *HugeNum) String() string {
	expName := h.ExpName()
	if expName == "" {
		return strconv.FormatFloat(h.Value, 'f', 3, 64)
	}
	return fmt.Sprintf("%g %s", h.Value, expName)
}
