package quantum

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	//! constants
	// matches gate and optional wires with delimiter "," and allows some gates to have n optional arguments in () right after wire delcariations
	gateWireRegex = regexp.MustCompile(`^([a-z]+)(\d+(?:,\d+)*)?(?:\((.*)\))?$`)
	// max wires
	maxWires = 9
	// max gates
	maxGates = 99_999

	//! errors

	ErrUnknownGate         = errors.New("unknown gate")
	ErrDuplicateWire	   = errors.New("duplicate wire")
	ErrInvalidWireFormat   = errors.New("invalid wire format")
	ErrInvalidArgument     = errors.New("invalid argument")
	ErrGateMatrixNotSquare = errors.New("gate matrix is not square")
	ErrInvalidWireCount    = errors.New("invalid wire count")
	ErrInvalidBarrier      = errors.New("invalid barrier")
	ErrTooManyGates        = errors.New(fmt.Sprintf("too many gates, max: %d", maxGates))
	ErrTooManyWires        = errors.New(fmt.Sprintf("too many wires, max: %d", maxWires))
)
