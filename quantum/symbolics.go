package quantum

import (
	"fmt"
	"math"
	"math/cmplx"
)

// ways to convert "constant" similar-ish values to their symbolic form
var constants = map[string]complex128{
	"0":                    0,
	"1":                    1,
	"i":                    1i,
	"1/sqrt(2)":            complex(1/math.Sqrt(2), 0),
	"-1/sqrt(2)":           complex(-1/math.Sqrt(2), 0),
	"1/sqrt(2)+i/2":        complex(1/math.Sqrt(2), 0.5),
	"1/sqrt(2)-i/2":        complex(1/math.Sqrt(2), -0.5),
	"-1/sqrt(2)+i/2":       complex(-1/math.Sqrt(2), 0.5),
	"-1/sqrt(2)-i/2":       complex(-1/math.Sqrt(2), -0.5),
	"1/2":                  complex(0.5, 0),
	"-1/2":                 complex(-0.5, 0),
	"i/2":                  complex(0, 0.5),
	"-i/2":                 complex(0, -0.5),
	"1/sqrt(2)+i/sqrt(2)":  complex(1/math.Sqrt(2), 1/math.Sqrt(2)),
	"1/sqrt(2)-i/sqrt(2)":  complex(1/math.Sqrt(2), -1/math.Sqrt(2)),
	"-1/sqrt(2)+i/sqrt(2)": complex(-1/math.Sqrt(2), 1/math.Sqrt(2)),
	"-1/sqrt(2)-i/sqrt(2)": complex(-1/math.Sqrt(2), -1/math.Sqrt(2)),
	"pi":                   math.Pi,
}

// map of complex128s to map of strings
func SymbofyMap(s map[string]complex128) map[string]string {
	result := make(map[string]string)
	for k, v := range s {
		result[k] = Symbofy(v)
	}
	return result
}

// symbofy a single complex num
func Symbofy(s complex128) string {
	const epsilon = 1e-9

	for name, val := range constants {
		if cmplx.Abs(s-val) < epsilon {
			return name
		}
	}

	// truncate again
	truncated := complex128(complex(float64(real(s)), float64(imag(s))))
	return fmt.Sprintf("%.5f", truncated)
}
