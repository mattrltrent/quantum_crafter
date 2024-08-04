package quantum

import (
	"fmt"
	"math"
)

func (g Gate) Name() string {
	return g.name
}

func (g Gate) Data() Matrix {
	return g.Matrix
}

type IdentityGate struct {
	Gate
}

func (g IdentityGate) WiresNeeded() int {
	return 1
}

func (g IdentityGate) Example() string {
	return "i0"
}

func (g IdentityGate) FullName() string {
	return "Identity"
}

func Identity(dim int) GateInterface {
	matrix := make([][]complex128, dim)
	for i := 0; i < dim; i++ {
		matrix[i] = make([]complex128, dim)
		matrix[i][i] = 1
	}
	return IdentityGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: dim,
				Cols: dim,
				Data: matrix,
			},
			name: "I",
		},
	}
}

type HadamardGate struct {
	Gate
}

func (g HadamardGate) WiresNeeded() int {
	return 1
}

func (g HadamardGate) Example() string {
	return "h0"
}

func (g HadamardGate) FullName() string {
	return "Hadamard"
}

func Hadamard() GateInterface {
	matrix := [][]complex128{
		{1 / complex(math.Sqrt2, 0), 1 / complex(math.Sqrt2, 0)},
		{1 / complex(math.Sqrt2, 0), -1 / complex(math.Sqrt2, 0)},
	}
	return HadamardGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 2,
				Cols: 2,
				Data: matrix,
			},
			name: "H",
		},
	}
}

type PauliXGate struct {
	Gate
}

func (g PauliXGate) WiresNeeded() int {
	return 1
}

func (g PauliXGate) Example() string {
	return "x0"
}

func (g PauliXGate) FullName() string {
	return "Pauli-X"
}

func PauliX() GateInterface {
	matrix := [][]complex128{
		{0, 1},
		{1, 0},
	}
	return PauliXGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 2,
				Cols: 2,
				Data: matrix,
			},
			name: "X",
		},
	}
}

type PauliYGate struct {
	Gate
}

func (g PauliYGate) WiresNeeded() int {
	return 1
}

func (g PauliYGate) Example() string {
	return "y0"
}

func (g PauliYGate) FullName() string {
	return "Pauli-Y"
}

func PauliY() GateInterface {
	matrix := [][]complex128{
		{0, -1i},
		{1i, 0},
	}
	return PauliYGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 2,
				Cols: 2,
				Data: matrix,
			},
			name: "Y",
		},
	}
}

type PauliZGate struct {
	Gate
}

func (g PauliZGate) WiresNeeded() int {
	return 1
}

func (g PauliZGate) Example() string {
	return "z0"
}

func (g PauliZGate) FullName() string {
	return "Pauli-Z"
}

func PauliZ() GateInterface {
	matrix := [][]complex128{
		{1, 0},
		{0, -1},
	}
	return PauliZGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 2,
				Cols: 2,
				Data: matrix,
			},
			name: "Z",
		},
	}
}

type CNOTGate struct {
	Gate
}

func (g CNOTGate) WiresNeeded() int {
	return 2
}

func (g CNOTGate) Example() string {
	return "cnot0,1"
}

func (g CNOTGate) FullName() string {
	return "Controlled-Not"
}

func CNOT() GateInterface {
	matrix := [][]complex128{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 1, 0},
	}
	return CNOTGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 4,
				Cols: 4,
				Data: matrix,
			},
			name: "CNOT",
		},
	}
}

type SWAPGate struct {
	Gate
}

func (g SWAPGate) WiresNeeded() int {
	return 2
}

func (g SWAPGate) Example() string {
	return "swap0,1"
}

func (g SWAPGate) FullName() string {
	return "SWAP"
}

func SWAP() GateInterface {
	matrix := [][]complex128{
		{1, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 1},
	}
	return SWAPGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 4,
				Cols: 4,
				Data: matrix,
			},
			name: "SWAP",
		},
	}
}

type CZGate struct {
	Gate
}

func (g CZGate) WiresNeeded() int {
	return 2
}

func (g CZGate) Example() string {
	return "cz0,1"
}

func (g CZGate) FullName() string {
	return "Controlled-Z"
}

func CZ() GateInterface {
	matrix := [][]complex128{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, -1},
	}
	return CZGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 4,
				Cols: 4,
				Data: matrix,
			},
			name: "CZ",
		},
	}
}

type CCXGate struct {
	Gate
}

func (g CCXGate) WiresNeeded() int {
	return 3
}

func (g CCXGate) Example() string {
	return "ccx0,1,2"
}

func (g CCXGate) FullName() string {
	return "Double-Controlled-X"
}

func CCX() GateInterface {
	matrix := [][]complex128{
		{1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 0},
	}
	return CCXGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 8,
				Cols: 8,
				Data: matrix,
			},
			name: "CCX",
		},
	}
}

type CCZGate struct {
	Gate
}

func (g CCZGate) WiresNeeded() int {
	return 3
}

func (g CCZGate) Example() string {
	return "ccz0,1,2"
}

func (g CCZGate) FullName() string {
	return "Double-Controlled-Z"
}

func CCZ() GateInterface {
	matrix := [][]complex128{
		{1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, -1},
	}
	return CCZGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 8,
				Cols: 8,
				Data: matrix,
			},
			name: "CCZ",
		},
	}
}

type TGate struct {
	Gate
}

func (g TGate) WiresNeeded() int {
	return 1
}

func (g TGate) Example() string {
	return "t0"
}

func (g TGate) FullName() string {
	return "T"
}

func T() GateInterface {
	matrix := [][]complex128{
		{1, 0},
		{0, complex(math.Cos(math.Pi/4), math.Sin(math.Pi/4))},
	}
	return TGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 2,
				Cols: 2,
				Data: matrix,
			},
			name: "T",
		},
	}
}

type SGate struct {
	Gate
}

func (g SGate) WiresNeeded() int {
	return 1
}

func (g SGate) Example() string {
	return "s0"
}

func (g SGate) FullName() string {
	return "S"
}

func S() GateInterface {
	matrix := [][]complex128{
		{1, 0},
		{0, 1i},
	}
	return SGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 2,
				Cols: 2,
				Data: matrix,
			},
			name: "S",
		},
	}
}

type PhaseGate struct {
	Gate
}

func (g PhaseGate) WiresNeeded() int {
	return 1
}

func (g PhaseGate) Example() string {
	return "p0"
}

func (g PhaseGate) FullName() string {
	return "Phase"
}

func Phase() GateInterface {
	matrix := [][]complex128{
		{1, 0},
		{0, 1i},
	}
	return PhaseGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 2,
				Cols: 2,
				Data: matrix,
			},
			name: "P",
		},
	}
}

type RxGate struct {
	Gate
	theta float64
}

func (g RxGate) WiresNeeded() int {
	return 1
}

func (g RxGate) Example() string {
	return "rx0(pi/2)"
}

func (g RxGate) FullName() string {
	return "Rotate-X"
}

func (g RxGate) Name() string {
	return fmt.Sprintf("Rx(%.2f)", g.theta)
}

func Rx(theta float64) GateInterface {
	matrix := [][]complex128{
		{complex(math.Cos(theta/2), 0), complex(0, -math.Sin(theta/2))},
		{complex(0, -math.Sin(theta/2)), complex(math.Cos(theta/2), 0)},
	}
	return RxGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 2,
				Cols: 2,
				Data: matrix,
			},
			name: "Rx",
		},
		theta: theta,
	}
}

type CRxGate struct {
	Gate
	theta float64
}

func (g CRxGate) WiresNeeded() int {
	return 2
}

func (g CRxGate) Example() string {
	return "crx0,1(pi/2)"
}

func (g CRxGate) FullName() string {
	return "Controlled-Rotate-X"
}

func (g CRxGate) Name() string {
	return fmt.Sprintf("CRx(%.2f)", g.theta)
}

func CRx(theta float64) GateInterface {
	matrix := [][]complex128{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, complex(math.Cos(theta/2), 0), complex(0, -math.Sin(theta/2))},
		{0, 0, complex(0, -math.Sin(theta/2)), complex(math.Cos(theta/2), 0)},
	}
	return CRxGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 4,
				Cols: 4,
				Data: matrix,
			},
			name: "CRx",
		},
		theta: theta,
	}
}

type RyGate struct {
	Gate
	theta float64
}

func (g RyGate) WiresNeeded() int {
	return 1
}

func (g RyGate) Example() string {
	return "ry0(pi/2)"
}

func (g RyGate) FullName() string {
	return "Rotate-Y"
}

func (g RyGate) Name() string {
	return fmt.Sprintf("Ry(%.2f)", g.theta)
}

func Ry(theta float64) GateInterface {
	matrix := [][]complex128{
		{complex(math.Cos(theta/2), 0), complex(-math.Sin(theta/2), 0)},
		{complex(math.Sin(theta/2), 0), complex(math.Cos(theta/2), 0)},
	}
	return RyGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 2,
				Cols: 2,
				Data: matrix,
			},
			name: "Ry",
		},
		theta: theta,
	}
}

type CRyGate struct {
	Gate
	theta float64
}

func (g CRyGate) WiresNeeded() int {
	return 2
}

func (g CRyGate) Example() string {
	return "cry0,1(pi/2)"
}

func (g CRyGate) FullName() string {
	return "Controlled-Rotate-Y"
}

func (g CRyGate) Name() string {
	return fmt.Sprintf("CRy(%.2f)", g.theta)
}

func CRy(theta float64) GateInterface {
	matrix := [][]complex128{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, complex(math.Cos(theta/2), 0), complex(-math.Sin(theta/2), 0)},
		{0, 0, complex(math.Sin(theta/2), 0), complex(math.Cos(theta/2), 0)},
	}
	return CRyGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 4,
				Cols: 4,
				Data: matrix,
			},
			name: "CRy",
		},
		theta: theta,
	}
}

type RzGate struct {
	Gate
	theta float64
}

func (g RzGate) WiresNeeded() int {
	return 1
}

func (g RzGate) Example() string {
	return "rz0(pi/2)"
}

func (g RzGate) FullName() string {
	return "Rotate-Z"
}

func (g RzGate) Name() string {
	return fmt.Sprintf("Rz(%.2f)", g.theta)
}

func Rz(theta float64) GateInterface {
	matrix := [][]complex128{
		{complex(math.Cos(theta/2), -math.Sin(theta/2)), 0},
		{0, complex(math.Cos(theta/2), math.Sin(theta/2))},
	}
	return RzGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 2,
				Cols: 2,
				Data: matrix,
			},
			name: "Rz",
		},
		theta: theta,
	}
}

type CRzGate struct {
	Gate
	theta float64
}

func (g CRzGate) WiresNeeded() int {
	return 2
}

func (g CRzGate) Example() string {
	return "crz0,1(pi/2)"
}

func (g CRzGate) FullName() string {
	return "Controlled-Rotate-Z"
}

func (g CRzGate) Name() string {
	return fmt.Sprintf("CRz(%.2f)", g.theta)
}

func CRz(theta float64) GateInterface {
	matrix := [][]complex128{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, complex(math.Cos(theta/2), -math.Sin(theta/2)), 0},
		{0, 0, 0, complex(math.Cos(theta/2), math.Sin(theta/2))},
	}
	return CRzGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 4,
				Cols: 4,
				Data: matrix,
			},
			name: "CRz",
		},
		theta: theta,
	}
}

type ToffoliGate struct {
	Gate
}

func (g ToffoliGate) WiresNeeded() int {
	return 3
}

func (g ToffoliGate) Example() string {
	return "toff0,1,2"
}

func (g ToffoliGate) FullName() string {
	return "Toffoli"
}

func Toffoli() GateInterface {
	matrix := [][]complex128{
		{1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 0},
	}
	return ToffoliGate{
		Gate: Gate{
			Matrix: Matrix{
				Rows: 8,
				Cols: 8,
				Data: matrix,
			},
			name: "TOFF",
		},
	}
}
