package quantum

type Matrix struct {
	Rows int
	Cols int
	Data [][]complex128
}

type Gate struct {
	Matrix Matrix
	name   string
}

type GateInterface interface {
	WiresNeeded() int
	Name() string
	Data() Matrix
}

type CircuitGate struct {
	Gate  GateInterface
	Wires []int
}

type Circuit struct {
	Gates []CircuitGate
}

type Result struct {
	StateVector         map[string]complex128
	StateVectorSymbolic map[string]string
	Probabilities       map[string]float64
}
