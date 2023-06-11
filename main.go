package main

import (
	"bytes"
	"flag"
	"os/exec"
	"strings"

	"github.com/revx-official/revxbuildtool/pkg/gcu"
	"github.com/revx-official/revxbuildtool/pkg/git"
	"github.com/revx-official/revxbuildtool/pkg/revx"

	"github.com/revx-official/log/pkg/console"
)

const (
	RevxPackagePath = "github.com/revx/pkg/revx"
)

func main() {
	Infof := console.NewConsoleColor().PrintFunc()
	Errorf := console.NewConsoleColor(console.FgColorRed).PrintFunc()

	var flagTrace bool
	var flagLocal bool
	var flagRelease bool
	var flagRevxInfoFile string

	flag.BoolVar(&flagTrace, "trace", false, "Enable trace logging.")
	flag.BoolVar(&flagLocal, "local", false, "Build revx for local development execution.")
	flag.BoolVar(&flagRelease, "release", false, "Build revx for release.")
	flag.StringVar(&flagRevxInfoFile, "revx-info", "version.yaml", "Specify the revx version info file.")

	flag.Parse()

	Infof("revx build tool\n")
	Infof("copyright © 2023 revx\n")
	Infof("all rights reserved\n")
	Infof("\n")

	revxInfo, err := revx.UnmarshalFromFile(flagRevxInfoFile)

	if err != nil {
		Errorf("error: %s", err.Error())
		return
	}

	Infof("revx info (%s):\n", flagRevxInfoFile)
	Infof("  app: %s\n", revxInfo.App)
	Infof("  model: %s\n", revxInfo.Model)
	Infof("  version: %s\n", revxInfo.Version)
	Infof("\n")

	commitHash, err := git.GetCurrentCommitHash()

	if err != nil {
		Errorf("error: %s\n", err.Error())
	}

	Infof("commit hash:\n")
	Infof("  %s\n", commitHash)
	Infof("\n")

	linker := gcu.NewLinker()

	linkerFlagRevxApp := gcu.NewLinkerFlag("-X", RevxPackagePath+".RevxApp="+revxInfo.App)
	linkerFlagRevxModel := gcu.NewLinkerFlag("-X", RevxPackagePath+".RevxModel="+revxInfo.Model)
	linkerFlagRevxVersion := gcu.NewLinkerFlag("-X", RevxPackagePath+".RevxVersion="+revxInfo.Version)
	linkerFlagRevxCommit := gcu.NewLinkerFlag("-X", RevxPackagePath+".RevxCommit="+commitHash)

	linker.AddFlag(linkerFlagRevxApp)
	linker.AddFlag(linkerFlagRevxModel)
	linker.AddFlag(linkerFlagRevxVersion)
	linker.AddFlag(linkerFlagRevxCommit)

	compiler := gcu.NewCompiler()
	if flagLocal {
		compiler.AddTag("local")
	}

	Infof("compiler tags:\n")
	for _, tag := range compiler.Tags {
		Infof("  %s\n", tag)
	}

	Infof("\n")

	Infof("linker flags:\n")
	Infof("  %s\n", linkerFlagRevxApp.Compile())
	Infof("  %s\n", linkerFlagRevxModel.Compile())
	Infof("  %s\n", linkerFlagRevxVersion.Compile())
	Infof("  %s\n", linkerFlagRevxCommit.Compile())

	linkerFlags := linker.Compile()
	compilerTags := compiler.CompileTags()

	Infof("\n")
	Infof("compiling ...\n")

	cmd := exec.Command("go", "build", "-tags", compilerTags, "-ldflags", linkerFlags, "-o", "bin/revx-daemon", "cmd/daemon.go")

	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer

	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	err = cmd.Run()
	if err != nil {
		Errorf("go: %s\n", err.Error())
	}

	output := outBuffer.String()
	errput := errBuffer.String()

	if len(output) > 0 {
		for _, line := range strings.Split(output, "\n") {
			trimmed := strings.TrimSpace(line)

			if len(trimmed) > 0 {
				Infof("go: %s\n", line)
			}
		}
	}

	if len(errput) > 0 {
		for _, line := range strings.Split(errput, "\n") {
			trimmed := strings.TrimSpace(line)

			if len(trimmed) > 0 {
				Errorf("go: %s\n", line)
			}
		}
	}
}
