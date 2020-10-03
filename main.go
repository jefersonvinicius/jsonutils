package main

import (
	"io/ioutil"
	"jsonutils/common"
	"log"
	"os"
)

func executeCommand(command common.Command, output common.Output) {
	result := command.Execute()
	output.Show(result)
}

// Args type
type Args struct {
	Command    string
	InputFile  string
	InTerminal bool
	OutputFile string
}

func normalizeArgs(args []string) Args {
	argsNormalized := Args{}
	argsNormalized.Command = args[1]
	argsNormalized.InputFile = args[2]
	for _, arg := range args {
		if arg == "--inTerminal" {
			argsNormalized.InTerminal = true
		}
	}

	if !argsNormalized.InTerminal {
		if len(args) < 4 {
			log.Fatalln("Invalid output file")
		}
		argsNormalized.OutputFile = args[3]
	}

	return argsNormalized
}

func main() {
	args := normalizeArgs(os.Args)

	var selectedOutput string
	if args.InTerminal {
		selectedOutput = "terminal"
	} else {
		selectedOutput = "file"
	}

	content, err := ioutil.ReadFile(args.InputFile)
	if err != nil {
		log.Fatal(err)
	}

	availableCommands := map[string]common.Command{
		"minify": common.MinifyCommand{InputData: content},
		"format": common.FormatCommand{InputData: content},
	}

	availableOutputs := map[string]common.Output{
		"file":     common.FileOutput{Path: args.OutputFile},
		"terminal": common.TerminalOutput{},
	}

	command, commandExists := availableCommands[args.Command]
	if !commandExists {
		log.Fatalf("'%s' is invalid comand option", args.Command)
	}

	output := availableOutputs[selectedOutput]

	executeCommand(command, output)
}
