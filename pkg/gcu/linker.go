package gcu

import (
	"fmt"
	"strings"
)

// Description:
//
//	Represents all linker specific information used by the go compiler.
type Linker struct {

	// The linker flags used by the go compiler.
	Flags []LinkerFlag
}

// Description:
//
//	Represents a single linker flag.
type LinkerFlag struct {

	// The flag name.
	Flag string

	// The value of the flag.
	Value string
}

// Description:
//
//	Creates a new linker instance.
//
// Returns:
//
//	A reference to the created instance.
func NewLinker() *Linker {
	return &Linker{}
}

// Description:
//
//	Adds a flag to the linker.
//
// Parameters:
//
//	flag The flag to add.
func (linker *Linker) AddFlag(flag *LinkerFlag) {
	linker.Flags = append(linker.Flags, *flag)
}

// Description:
//
//	Formats all linker flags.
//	Outputs a string which can be directly passed to the go compiler.
//
// Returns:
//
//	A command line string containing all registered linker flags.
func (linker *Linker) Compile() string {
	result := ""

	for index, flag := range linker.Flags {
		compiledFlag := flag.Compile()

		if index == 0 {
			result = compiledFlag
			continue
		}

		result = strings.Join([]string{result, compiledFlag}, " ")
	}

	return result
}

// Description:
//
//	Creates a new linker flag instance.
//
// Returns:
//
//	A reference to the created instance.
func NewLinkerFlag(flag string, value string) *LinkerFlag {
	return &LinkerFlag{
		Flag:  flag,
		Value: value,
	}
}

// Description:
//
//	Formats the linker flag.
//	Outputs a string which can be directly passed as a linker flag to the go compiler.
//
// Returns:
//
//	The linker flag as expected by the go compiler.
func (flag *LinkerFlag) Compile() string {
	if len(flag.Value) == 0 {
		return flag.Flag
	}

	return fmt.Sprintf(`%s "%s"`, flag.Flag, flag.Value)
}
