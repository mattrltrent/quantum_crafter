package cli

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"github.com/mattrltrent/quantum_crafter/quantum"
	"golang.org/x/term"
)

func init() {
	color.Output = colorable.NewColorableStdout()
}

// color vars
var (
	white        = color.New(color.FgWhite, color.Bold).SprintFunc()
	whiteFprintf = color.New(color.FgWhite, color.Bold).FprintfFunc()
	whitePrintf  = color.New(color.FgWhite, color.Bold).PrintfFunc()
	whitePrintln = color.New(color.FgWhite, color.Bold).PrintlnFunc()
	redPrintln   = color.New(color.FgRed, color.Bold).PrintlnFunc()
)

// help menu
func PrintHelp() {
	redPrintln("Usage:")
	whitePrintln("  help                  - shows this help message")
	whitePrintln("  version               - shows tool version")
	whitePrintln("  repo                  - opens the github repository")
	whitePrintln("  gates                 - lists available gates")
	whitePrintln("  run \"<gates here>\"    - executes the quantum circuit with the provided gates")
	redPrintln("Circuit examples:")
	whitePrintln("  run \"z2 x1 cnot0,1 cz2,3 rz0(-pi/2*(-3^2)) toff1,2,3\"      - random gates")
	whitePrintln("  run \"h0 cnot0,1\"                                           - creates bell pair")
	redPrintln("Notes:")
	whitePrintln("  indexing starts at 0")
	whitePrintln("  arithmetic operations must be explicit (yes: 2*pi, no: 2pi)")
	redPrintln("How to apply gates by type:")
	whitePrintln("  h0          - hadamard gate on wire 0")
	whitePrintln("  cnot0,1     - controlled not gate with control wire 0 and target wire 1")
	whitePrintln("  rx0(pi/2)   - rotate x gate on wire 0 by pi/2 radians")
	whitePrintln("  crz0,1(pi)  - controlled z rotation with control on wire 0 and target 1 for rotation pi radians")
	whitePrintln("  toff0,1,2   - toffoli gate with control wires 0, 1 and target wire 2")
	redPrintln("Arithmetic operations for rotational gates:")
	whitePrintln("  pi          - defined constant")
	whitePrintln("  ()          - order of operators")
	whitePrintln("  +           - addition")
	whitePrintln("  -           - subtraction")
	whitePrintln("  *           - multiplication")
	whitePrintln("  /           - floating point division")
	whitePrintln("  ^           - exponentiation")
}

// print available gates
func PrintGates() {
	gates := []quantum.GateInterface{
		quantum.Identity(2),
		quantum.Hadamard(),
		quantum.PauliX(),
		quantum.PauliY(),
		quantum.PauliZ(),
		quantum.CNOT(),
		quantum.CZ(),
		quantum.SWAP(),
		quantum.Toffoli(),
		quantum.T(),
		quantum.S(),
		quantum.Phase(),
		quantum.Rx(1),
		quantum.Ry(1),
		quantum.Rz(1),
		quantum.CRx(1),
		quantum.CRy(1),
		quantum.CRz(1),
		quantum.CCX(),
		quantum.CCZ(),
	}

	redPrintln("Gates & example usage:")
	for _, gate := range gates {
		// show gate.Name() and gate.Example()
		whitePrintf("%s: %s\n", gate.FullName(), gate.Example())
	}
}

// version
func PrintVersion() {
	whitePrintf("%s\n", Version)
}

// open repo in browser
func OpenRepo() {
	exec.Command("open", RepoURL).Run()
}

// execute interactively
func ExecuteCircuit(gates []string) {
	// decode args
	for i, gate := range gates {
		decoded, err := url.QueryUnescape(gate)
		if err != nil {
			whitePrintf("Error decoding argument: %v\n", err)
			return
		}
		gates[i] = decoded
	}

	circuit, err := quantum.NewCircuit(gates)
	if err != nil {
		whitePrintf("Error creating circuit: %v\n", err)
		return
	}

	RunInteractiveCLI(&circuit)
}

func getSingleKey() (rune, error) {
	var buf [1]byte
	if _, err := syscall.Read(syscall.Stdin, buf[:]); err != nil {
		return 0, err
	}
	return rune(buf[0]), nil
}

func enableRawMode() (*term.State, error) {
	oldState, err := term.MakeRaw(int(syscall.Stdin))
	if err != nil {
		return nil, err
	}
	return oldState, nil
}

func disableRawMode(state *term.State) {
	term.Restore(int(syscall.Stdin), state)
}

func RunInteractiveCLI(circuit *quantum.Circuit) {
	state, err := enableRawMode()
	if err != nil {
		fmt.Println("Failed to enable raw mode:", err)
		return
	}
	defer disableRawMode(state)

	atBarrier := len(circuit.Gates)
	clearScreen()
	circuit.Draw(atBarrier)

	for {
		key, err := getSingleKey()
		if err != nil {
			fmt.Println("Failed to read key:", err)
			break
		}

		switch key {
		case 'q':
			return
		// also quit on ctrl+c
		case 3:
			return
		case 'j':
			if atBarrier > 1 {
				atBarrier--
				clearScreen()
				circuit.Draw(atBarrier)
			}
		case 'k':
			if atBarrier < len(circuit.Gates) {
				atBarrier++
				clearScreen()
				circuit.Draw(atBarrier)
			}
		}
	}
}

// clear terminal screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Run() {
	if len(os.Args) < 2 {
		PrintHelp()
		return
	}

	switch os.Args[1] {
	case "gates":
		PrintGates()
	case "repo":
		OpenRepo()
	case "help", "-h", "--help":
		PrintHelp()
	case "version", "-v", "--version":
		PrintVersion()
	case "run":
		if len(os.Args) < 3 {
			PrintHelp()
			return
		}
		gates := strings.Split(os.Args[2], " ")
		ExecuteCircuit(gates)
	default:
		PrintHelp()
	}
}
