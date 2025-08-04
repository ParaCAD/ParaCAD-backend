package openscadgenerator

import (
	"fmt"
)

const (
	FormatSTL = "binstl"
	FormatPNG = "png"

	ColorSchemeBeforeDawn = "BeforeDawn"
)

func buildArgs(format, outputFile, scadFilePath string) []string {
	args := []string{}
	args = append(args, fmt.Sprintf("--export-format=%s", format))

	if format == FormatPNG {
		args = append(args, fmt.Sprintf("--colorscheme=%s", ColorSchemeBeforeDawn))
	}

	args = append(args, fmt.Sprintf("-o%s", outputFile))
	args = append(args, scadFilePath)

	return args
}
