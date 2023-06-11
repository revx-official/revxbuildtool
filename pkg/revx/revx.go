package revx

import (
	"gopkg.in/yaml.v3"

	"io/ioutil"
)

// Description:
//
//	The representation of the revx version info file.
type RevxInfo struct {

	// The app name.
	App string `yaml:"app" json:"app"`

	// The revx model.
	Model string `yaml:"model" json:"model"`

	// The revx version.
	Version string `yaml:"version" json:"version"`

	// The current revx commit.
	Commit string `yaml:"commit" json:"commit"`
}

// Description:
//
//	Unmarshals revx version information from the given file.
//
// Parameters:
//
//	The path to the file containing the version information.
//
// Returns:
//
//	The revx version info or an error, if unmarshaling fails.
func UnmarshalFromFile(path string) (*RevxInfo, error) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	revxInfo := RevxInfo{}
	err = yaml.Unmarshal(content, &revxInfo)

	if err != nil {
		return nil, err
	}

	return &revxInfo, nil
}
