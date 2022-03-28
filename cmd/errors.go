package cmd

import (
	"fmt"
	"sort"
	"strings"
)

var _ error = ChainNotFoundError{}

// ChainNotFoundError is used when a requested chain does not exist.
// Its error message includes the list of known chains.
type ChainNotFoundError struct {
	Requested string
	Config    *Config
}

func (e ChainNotFoundError) Error() string {
	available := make([]string, 0, len(e.Config.Chains))
	for chainName := range e.Config.Chains {
		available = append(available, chainName)
	}
	sort.Strings(available)

	return fmt.Sprintf(
		"no chain %q found (available chains: %s)",
		e.Requested,
		strings.Join(available, ", "),
	)
}

// GRPCServiceNotFoundError is used when a requested gRPC service does not exist.
// Its error message includes the provided available services.
type GRPCServiceNotFoundError struct {
	Requested string
	Available []string
}

func (e GRPCServiceNotFoundError) Error() string {
	sort.Strings(e.Available)
	// TODO: would be nice to suggest close matches here.
	return fmt.Sprintf(
		"no service %q found (available services: %s)",
		e.Requested,
		strings.Join(e.Available, ", "),
	)
}
