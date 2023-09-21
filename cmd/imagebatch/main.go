package main

import (
	"fmt"

	"github.com/mhborthwick/image-batch-processor/pkg/processor"
)

func main() {
	inputDir := "~/input"
	outputDir := "~/output"
	err := processor.ProcessImgsInBatch(inputDir, outputDir)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success, your new images are in ", outputDir)
	}
}
