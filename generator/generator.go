package generator

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/ParaCAD/ParaCAD-backend/utils"
)

const tempDirName string = "temp"
const openScadLibrariesIncludes string = ``

func Generate(template FilledTemplate, format ModelFileFormat) ([]byte, error) {
	if runtime.GOOS != "linux" {
		return nil, fmt.Errorf("runtime.GOOS != linux")
	}

	programDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd(): %s", err.Error())
	}

	// Ensure temp dir exists
	err = os.MkdirAll(tempDirName, 0700)
	if err != nil {
		return nil, fmt.Errorf("os.MkdirAll(tempDirName, 0700): %s", err.Error())
	}

	// Create temporary .scad file
	tempID := utils.CreateRandomString(6)
	scadFileName := template.UUID.String() + "_" + tempID + ".scad"
	scadFilePath := path.Join(programDir, tempDirName, scadFileName)
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

	// Generate model
	cmdArgs := []string{}
	cmdArgs = append(cmdArgs, "--export-format="+string(format))
	cmdArgs = append(cmdArgs, "-o-")
	cmdArgs = append(cmdArgs, scadFilePath)

	cmd := exec.Command("openscad", cmdArgs...)
	fmt.Println(cmd.String())
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("cmd.Run(): %s", stderr.String())
	}

	model := stdout.Bytes()

	return model, nil
}
