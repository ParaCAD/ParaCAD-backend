package openscadgenerator

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/ParaCAD/ParaCAD-backend/generator"
	"github.com/ParaCAD/ParaCAD-backend/utils"
)

func (g OpenSCADGenerator) GeneratePreview(template generator.FilledTemplate) ([]byte, error) {
	if runtime.GOOS != "linux" {
		return nil, fmt.Errorf("runtime.GOOS != linux")
	}

	// Ensure temp dir exists
	err := os.MkdirAll(tempDirName, 0700)
	if err != nil {
		return nil, fmt.Errorf("os.MkdirAll(tempDirName, 0700): %s", err.Error())
	}

	// Create temporary .scad file
	tempID := utils.CreateRandomString(6)
	scadFileName := template.UUID.String() + "_" + tempID + ".scad"
	scadFilePath := path.Join(tempDirName, scadFileName)
	scadFile, err := os.Create(scadFilePath)
	if err != nil {
		return nil, fmt.Errorf("os.Create(scadFilePath): %s", err.Error())
	}
	defer os.Remove(scadFile.Name())

	// Include all libraries
	scadFile.WriteString(openScadLibrariesIncludes)

	// Add parameters
	// If issues arise, could switch to using -D flag
	for _, param := range template.Params {
		scadParam := param.Key + "=" + param.Value + ";\n"
		_, err := scadFile.WriteString(scadParam)
		if err != nil {
			return nil, fmt.Errorf("scadFile.WriteString(scadParam) %s: %s", param.Key, err.Error())
		}
	}

	// Write template code
	scadFile.Write(template.Template)

	// Close file
	err = scadFile.Close()
	if err != nil {
		return nil, fmt.Errorf("scadFile.Close(): %s", err.Error())
	}

	outputFileName := template.UUID.String() + "_" + tempID + ".png"
	outputFilePath := path.Join(tempDirName, outputFileName)

	// Generate model
	cmdArgs := buildArgs(FormatPNG, outputFilePath, scadFilePath)

	cmd := exec.Command("openscad", cmdArgs...)
	output, err := cmd.CombinedOutput()
	defer os.Remove(outputFilePath)
	if err != nil {
		return nil, fmt.Errorf("cmd.Run(): %s", string(output))
	}

	preview, err := os.ReadFile(outputFilePath)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile(outputFilePath): %s", err.Error())
	}

	return preview, nil
}
