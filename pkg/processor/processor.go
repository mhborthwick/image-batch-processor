package processor

import "fmt"

func processImg(
	inputDir string,
	outputDir string,
) error {
	return nil
}

func ProcessImgsInBatch(
	inputDir string,
	outputDir string,
) error {
	// TODO: Loop
	err := processImg(inputDir, outputDir)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success!")
	}
	return nil
}
