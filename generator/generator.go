package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/ParaCAD/ParaCAD-backend/utils"
)

const tempDirName string = "/dev/shm/ParaCAD"
const openScadLibrariesIncludes string = ``

func Generate(template FilledTemplate) ([]byte, error) {
	if runtime.GOOS != "linux" {
		return nil, fmt.Errorf("runtime.GOOS != linux")
	}

	//programDir, err := os.Getwd()
	// if err != nil {
	// 	return nil, fmt.Errorf("os.Getwd(): %s", err.Error())
	// }

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

	outputFileName := template.UUID.String() + "_" + tempID + ".stl"
	outputFilePath := path.Join(tempDirName, outputFileName)

	// Generate model
	cmdArgs := []string{}
	cmdArgs = append(cmdArgs, "--export-format=binstl")
	cmdArgs = append(cmdArgs, "-o"+outputFilePath)
	cmdArgs = append(cmdArgs, scadFilePath)

	cmd := exec.Command("openscad", cmdArgs...)
	output, err := cmd.CombinedOutput()
	defer os.Remove(outputFilePath)
	if err != nil {
		return nil, fmt.Errorf("cmd.Run(): %s", string(output))
	}

	model, err := os.ReadFile(outputFilePath)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile(outputFilePath): %s", err.Error())
	}

	return model, nil
}
