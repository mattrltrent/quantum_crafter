package quantum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func NewCircuit(gates []string) (Circuit, error) {
	circuit := Circuit{}
	for _, name := range gates {
		gate, err := nameToCircuitGate(name)
		if err != nil {
			return Circuit{}, err
		}
		circuit.Gates = append(circuit.Gates, gate)
	}

	numOfGates := len(circuit.Gates)
	if numOfGates > maxGates {
		return Circuit{}, ErrTooManyGates
	}

	return circuit, nil
}

func nameToCircuitGate(name string) (CircuitGate, error) {

	// make lowercase
	name = strings.ToLower(name)

	match := gateWireRegex.FindStringSubmatch(name)
	if match == nil {
		return CircuitGate{}, ErrUnknownGate
	}

	gateName := match[1]
	wireStr := match[2]

	var gate GateInterface
	wires := []int{}
	if wireStr != "" {
		wireParts := strings.Split(wireStr, ",")
		for _, ws := range wireParts {
			wire, err := strconv.Atoi(ws)
			if err != nil || wire < 0 {
				return CircuitGate{}, ErrInvalidWireFormat
			}
			if wire > maxWires {
				return CircuitGate{}, ErrTooManyWires
			}
			wires = append(wires, wire)
		}
	}

	switch gateName {
	case "i":
		gate = Identity(2)
	case "h":
		gate = Hadamard()
	case "x":
		gate = PauliX()
	case "y":
		gate = PauliY()
	case "z":
		gate = PauliZ()
	case "cnot":
		gate = CNOT()
	case "swap":
		gate = SWAP()
	case "cz":
		gate = CZ()
	case "t":
		gate = T()
	case "s":
		gate = S()
	case "p":
		gate = Phase()
	case "toff":
		gate = Toffoli()
	default:
		return CircuitGate{}, ErrUnknownGate
	}

	if len(wires) != gate.WiresNeeded() {
		return CircuitGate{}, fmt.Errorf("%s gate requires %d wire(s)", gateName, gate.WiresNeeded())
	}

	return CircuitGate{Gate: gate, Wires: wires}, nil
}

func (c *Circuit) Draw() error {
	// attempt executing first to ensure the circuit is valid
	_, err := c.Execute()
	if err != nil {
		return err
	}

	// color directives
	qubitColor := color.New(color.FgCyan).SprintfFunc()
	wireColor := color.New(color.FgWhite).SprintfFunc()
	barrierColor := color.New(color.BgRed).SprintfFunc()
	gateColor := color.New(color.FgYellow).SprintfFunc()

	longestNameSize := 0
	for _, gate := range c.Gates {
		if len(gate.Gate.Name()) > longestNameSize {
			longestNameSize = len(gate.Gate.Name())
		}
	}

	var sb strings.Builder
	numQubits := 0
	for _, gate := range c.Gates {
		for _, wire := range gate.Wires {
			if wire >= numQubits {
				numQubits = wire + 1
			}
		}
	}

	qubitLines := make([]string, numQubits)
	for i := 0; i < numQubits; i++ {
		qubitLines[i] = qubitColor("q[%d]|0>", i)
	}

	barrierPositions := []int{}

	for _, gate := range c.Gates {
		gateStr := gate.Gate.Name()
		segmentSize := len(gateStr) + 4
		padding := segmentSize - len(gateStr)
		leftPad := (padding + 1) / 2
		rightPad := padding / 2

		if len(gate.Wires) == 1 {
			wire := gate.Wires[0]
			qubitLines[wire] += fmt.Sprintf("%s%s%s", wireColor(strings.Repeat("-", leftPad)), gateColor(strings.ToUpper(gateStr)), wireColor(strings.Repeat("-", rightPad)))
			for i := 0; i < numQubits; i++ {
				if i != wire {
					qubitLines[i] += wireColor(strings.Repeat("-", segmentSize))
				}
			}
		} else if len(gate.Wires) == 2 {
			control := gate.Wires[0]
			target := gate.Wires[1]
			minWire := control
			maxWire := target
			if control > target {
				minWire = target
				maxWire = control
			}

			controlStr := gateColor(strings.Repeat("•", len(gateStr)))
			controlStrPadded := fmt.Sprintf("%s%s%s", wireColor(strings.Repeat("-", leftPad)), controlStr, wireColor(strings.Repeat("-", rightPad)))

			for i := 0; i < numQubits; i++ {
				if i == control {
					qubitLines[i] += controlStrPadded
				} else if i == target {
					qubitLines[i] += fmt.Sprintf("%s%s%s", wireColor(strings.Repeat("-", leftPad)), gateColor(strings.ToUpper(gateStr)), wireColor(strings.Repeat("-", rightPad)))
				} else if i > minWire && i < maxWire {
					qubitLines[i] += wireColor(strings.Repeat("-", segmentSize))
				} else {
					qubitLines[i] += wireColor(strings.Repeat("-", segmentSize))
				}
			}
		} else if len(gate.Wires) == 3 {
			control1 := gate.Wires[0]
			control2 := gate.Wires[1]
			target := gate.Wires[2]
			controlStr := gateColor(strings.Repeat("•", len(gateStr)))
			controlStrPadded := fmt.Sprintf("%s%s%s", wireColor(strings.Repeat("-", leftPad)), controlStr, wireColor(strings.Repeat("-", rightPad)))

			for i := 0; i < numQubits; i++ {
				if i == control1 || i == control2 {
					qubitLines[i] += controlStrPadded
				} else if i == target {
					qubitLines[i] += fmt.Sprintf("%s%s%s", wireColor(strings.Repeat("-", leftPad)), gateColor(strings.ToUpper(gateStr)), wireColor(strings.Repeat("-", rightPad)))
				} else {
					qubitLines[i] += wireColor(strings.Repeat("-", segmentSize))
				}
			}
		}

		for i := 0; i < numQubits; i++ {
			qubitLines[i] += barrierColor("|")
		}
		barrierPositions = append(barrierPositions, len(qubitLines[0])-1)
	}

	for i := range qubitLines {
		sb.WriteString(qubitLines[i])
		if i < len(qubitLines)-1 {
			sb.WriteString("\n")
		}
	}

	// add column (barrier) #s
	sb.WriteString("\n")
	sb.WriteString(strings.Repeat(" ", 7))
	for i := range barrierPositions {
		offset := 0
		// for each length of string(i) > 1, add 1 to offset
		if i > 1 {
			offset = len(strconv.Itoa(i)) - 1
		}
		sb.WriteString(fmt.Sprintf("%s%d", strings.Repeat(" ", len(string(c.Gates[i].Gate.Name()))+4-offset), i+1))
	}

	fmt.Println(sb.String())
	return nil
}

func (c *Circuit) Execute() (Result, error) {
	return c.ExecuteToBarrier(len(c.Gates))
}

func (c *Circuit) ExecuteToBarrier(atBarrier int) (Result, error) {
	if atBarrier < 1 || atBarrier > len(c.Gates) {
		return Result{}, ErrInvalidBarrier
	}

	numQubits := 0
	for _, gate := range c.Gates {
		for _, wire := range gate.Wires {
			if wire >= numQubits {
				numQubits = wire + 1
			}
		}
	}

	stateVector := NewMatrix(1<<numQubits, 1)
	stateVector.Data[0][0] = 1

	// apply gates up to the barrier
	gatesToApply := c.Gates
	if atBarrier > 0 && atBarrier < len(c.Gates) {
		gatesToApply = c.Gates[:atBarrier]
	}

	newStateVector, err := applyGates(stateVector, gatesToApply, numQubits)
	if err != nil {
		return Result{}, err
	}

	labeledStateVector := make(map[string]complex128)
	for i := 0; i < (1 << numQubits); i++ {
		key := strings.Join(intToBitString(i, numQubits), "")
		labeledStateVector[key] = newStateVector.Data[i][0]
	}

	result := Result{
		StateVector:         labeledStateVector,
		StateVectorSymbolic: SymbofyMap(labeledStateVector),
		Probabilities:       Probabilities(labeledStateVector),
	}

	return result, nil
}

func Probabilities(stateVector map[string]complex128) map[string]float64 {
	probabilities := make(map[string]float64)
	for key, value := range stateVector {
		probabilities[key] = real(value * value)
	}
	return probabilities
}

func intToBitString(value, length int) []string {
	bitString := make([]string, length)
	for i := 0; i < length; i++ {
		if value&(1<<i) != 0 {
			bitString[length-i-1] = "1"
		} else {
			bitString[length-i-1] = "0"
		}
	}
	return bitString
}

func applyGates(stateVector Matrix, gates []CircuitGate, numQubits int) (Matrix, error) {
	for _, gate := range gates {
		if gate.Gate.Data().Rows != gate.Gate.Data().Cols {
			return Matrix{}, ErrGateMatrixNotSquare
		}

		fullGate := createFullGateMatrix(gate, numQubits)
		if !fullGate.CanMultiply(&stateVector) {
			return Matrix{}, errors.New("cannot multiply matrices")
		}
		stateVector = fullGate.MustMultiply(&stateVector)
	}
	return stateVector, nil
}

func createFullGateMatrix(gate CircuitGate, numQubits int) Matrix {
	fullGate := Identity(1 << numQubits).Data()

	if len(gate.Wires) == 1 {
		for i := 0; i < numQubits; i++ {
			if i == gate.Wires[0] {
				fullGate = applyGateToQubit(gate.Gate.Data(), i, numQubits)
			}
		}
	} else if len(gate.Wires) == 2 {
		fullGate = applyTwoQubitGate(gate.Gate.Data(), gate.Wires, numQubits)
	} else if len(gate.Wires) == 3 {
		// triple gates
		fullGate = applyThreeQubitGate(gate.Gate.Data(), gate.Wires, numQubits)
	}

	return fullGate
}

func applyGateToQubit(gate Matrix, qubit, numQubits int) Matrix {
	if qubit == 0 {
		identity := Identity(1 << (numQubits - 1)).Data()
		return tensorGateMatrix(&gate, &identity)
	} else if qubit == numQubits-1 {
		identity := Identity(1 << (numQubits - 1)).Data()
		return tensorGateMatrix(&identity, &gate)
	} else {
		left := Identity(1 << qubit).Data()
		right := Identity(1 << (numQubits - qubit - 1)).Data()
		temp := tensorGateMatrix(&left, &gate)
		return tensorGateMatrix(&temp, &right)
	}
}

func applyTwoQubitGate(gate Matrix, wires []int, numQubits int) Matrix {
	if len(wires) != 2 {
		panic(ErrInvalidWireCount)
	}

	control, target := wires[0], wires[1]
	if control > target {
		control, target = target, control
	}

	left := Identity(1 << control).Data()
	middle := gate
	right := Identity(1 << (numQubits - target - 1)).Data()

	if target-control > 1 {
		middleLeft := Identity(1 << (target - control - 1)).Data()
		middle = tensorGateMatrix(&middleLeft, &middle)
	}

	temp := tensorGateMatrix(&left, &middle)
	return tensorGateMatrix(&temp, &right)
}

func applyThreeQubitGate(gate Matrix, wires []int, numQubits int) Matrix {
	if len(wires) != 3 {
		panic(ErrInvalidWireCount)
	}

	control1, control2, target := wires[0], wires[1], wires[2]
	if control1 > control2 {
		control1, control2 = control2, control1
	}
	if control1 > target {
		control1, target = target, control1
	}
	if control2 > target {
		control2, target = target, control2
	}

	left := Identity(1 << control1).Data()
	middle := gate
	right := Identity(1 << (numQubits - target - 1)).Data()

	if target-control2 > 1 {
		middleLeft := Identity(1 << (target - control2 - 1)).Data()
		middle = tensorGateMatrix(&middleLeft, &middle)
	}
	if control2-control1 > 1 {
		middleLeft := Identity(1 << (control2 - control1 - 1)).Data()
		middle = tensorGateMatrix(&middleLeft, &middle)
	}

	temp := tensorGateMatrix(&left, &middle)
	return tensorGateMatrix(&temp, &right)
}
