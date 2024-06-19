package quantum

import "math"

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

type TGate struct {
	Gate
}

func (g TGate) WiresNeeded() int {
	return 1
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

type ToffoliGate struct {
	Gate
}

func (g ToffoliGate) WiresNeeded() int {
	return 3
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
