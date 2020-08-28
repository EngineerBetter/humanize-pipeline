package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func fatalIf(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

var debug = false

func main() {
	pipelinePath := parseArguments()

	originalYaml, err := ioutil.ReadFile(pipelinePath)
	fatalIf(err)

	reorderedYaml := reorderKnownSchema(originalYaml)

	orderedCompleteYaml := addNodesThatAreUnknownToSchema(originalYaml, reorderedYaml)

	fmt.Println("---")
	fmt.Print(string(orderedCompleteYaml))
}

func parseArguments() string {
	if len(os.Args) == 0 {
		os.Exit(2)
	}
	pipelineArgIdx := 1
	if os.Args[1] == "-d" {
		fmt.Fprintf(os.Stderr, "DEBUG: activating debug mode\n")
		debug = true
		if len(os.Args) < 3 {
			os.Exit(2)
		}
		pipelineArgIdx = 2
	}
	pipelinePath := os.Args[pipelineArgIdx]
	if pipelinePath == "" {
		os.Exit(2)
	}
	return pipelinePath
}

func reorderKnownSchema(originalYaml []byte) []byte {
	pipeline := &Pipeline{}
	err := yaml.Unmarshal(originalYaml, pipeline)
	fatalIf(err)

	yamlOutput, err := yaml.Marshal(*pipeline)
	fatalIf(err)
	return yamlOutput
}

func addNodesThatAreUnknownToSchema(originalYaml []byte, reorderedYaml []byte) []byte {
	completeTree := &yaml.MapSlice{}
	err := yaml.Unmarshal(originalYaml, &completeTree)
	fatalIf(err)

	orderedTree := &yaml.MapSlice{}
	err = yaml.Unmarshal(reorderedYaml, &orderedTree)
	fatalIf(err)

	storeDest := func(modified yaml.MapSlice) {
		if debug {
			fmt.Fprintf(os.Stderr, "DEBUG: storing modified MapSlice at path '/' (root MapSlice)\n")
		}
		orderedTree = &modified
	}
	err = appendMissingNodes(*completeTree, *orderedTree, storeDest, "/")
	fatalIf(err)

	orderedCompleteYaml, err := yaml.Marshal(orderedTree)
	fatalIf(err)

	return orderedCompleteYaml
}
