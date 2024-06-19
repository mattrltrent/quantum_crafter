package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"github.com/mattrltrent/quantum_crafter/quantum"
	"github.com/rodaine/table"
)

func init() {
	color.Output = colorable.NewColorableStdout()
}

// color vars
var (
	white        = color.New(color.FgWhite).SprintFunc()
	whiteFprintf = color.New(color.FgWhite).FprintfFunc()
	whitePrintf  = color.New(color.FgWhite).PrintfFunc()
	whitePrintln = color.New(color.FgWhite).PrintlnFunc()
)

// help menu
func PrintHelp() {
	whitePrintln("Usage:")
	whitePrintln("----------------")
	whitePrintln("  help       - shows this help message")
	whitePrintln("  version    - shows tool version")
	whitePrintln("  repo       - opens the GitHub repository")
	whitePrintln("  gates      - lists available gates")
	whitePrintln("  <circuit>  - executes the quantum circuit with the provided gates")
	whitePrintln("----------------")
	whitePrintln("Circuit examples:")
	whitePrintln("----------------")
	whitePrintln("  z2 x1 cnot0,1 cz2,3 toff1,2,3      - random gates")
	whitePrintln("  h0 cnot0,1                         - creates bell pair")
	whitePrintln("----------------")
}

// print available gates
func PrintGates() {
	whitePrintln("Available gates:")
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
	}

	for _, gate := range gates {
		whitePrintf("  - %s\n", gate.Name())
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

// execute the circuit
func ExecuteCircuit(gates []string) {
	circuit, err := quantum.NewCircuit(gates)
	if err != nil {
		whitePrintf("Error creating circuit: %v\n", err)
		return
	}

	err = circuit.Draw()
	if err != nil {
		whitePrintf("Error drawing circuit: %v\n", err)
		return
	}

	// input # to get state at what barrier
	whitePrintf("\nExecute to barrier # (default all): ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		whitePrintln("Error reading input")
		return
	}
	whitePrintln()

	input = strings.TrimSpace(input)
	barrier, err := strconv.Atoi(input)
	if err != nil {
		barrier = len(circuit.Gates)
	}

	result, err := circuit.ExecuteToBarrier(barrier)
	if err != nil {
		whitePrintf("Error executing circuit: %v\n", err)
		return
	}

	// create slice to store the probs
	probabilities := make([]struct {
		Key   string
		Value float64
		State string
	}, 0)

	for key, value := range result.Probabilities {
		if value > 0 {
			probabilities = append(probabilities, struct {
				Key   string
				Value float64
				State string
			}{
				Key:   key,
				Value: value,
				State: result.StateVectorSymbolic[key],
			})
		}
	}

	// sort/order by prob
	sort.Slice(probabilities, func(i, j int) bool {
		return probabilities[i].Value > probabilities[j].Value
	})

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Non-zero basis vectors", "Amplitudes", "Probabilities")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithPadding(5)

	for _, p := range probabilities {
		// trunc to 5 decimals
		truncatedValue := fmt.Sprintf("%.2f%%", float64(int(p.Value*1e5))/1e5*100)
		tbl.AddRow(p.Key, p.State, truncatedValue)
	}

	tbl.Print()
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
	default:
		gates := os.Args[1:]
		ExecuteCircuit(gates)
	}
}
