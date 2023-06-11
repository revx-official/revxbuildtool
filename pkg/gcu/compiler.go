package gcu

import "strings"

// Description:
//
// 	Represents all compiler specific information used by the go compiler.
type Compiler struct {

	// A list of compiler tags.
	Tags []string
}

// Description:
//
//	Creates a new compiler instance.
//
// Returns:
//
//	A reference to the created instance.
func NewCompiler() *Compiler {
	return &Compiler{}
}

// Description:
//
//	Adds a tag to the compiler.
//
// Parameters:
//
//	tag The tag to add.
func (compiler *Compiler) AddTag(tag string) {
	compiler.Tags = append(compiler.Tags, tag)
}

// Description:
//
//	Formats all compiler tags.
//	Outputs a string which can be directly passed to the go compiler.
//
// Returns:
//
//	A command line string containing all registered compiler tags.
func (compiler *Compiler) CompileTags() string {
	result := ""

	for index, tag := range compiler.Tags {
		if index == 0 {
			result = tag
			continue
		}

		result = strings.Join([]string{result, tag}, ",")
	}

	return result
}
