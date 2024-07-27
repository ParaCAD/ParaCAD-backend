package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/ParaCAD/ParaCAD-backend/utils"
)

const tempDirName string = "temp"
const openScadLibrariesIncludes string = `

`

func Generate(template FilledTemplate) ([]byte, error) {
	if runtime.GOOS != "linux" {
		return nil, fmt.Errorf("this program only runs on real operating systems (install linux and try again)")
	}

	// Ensure temp dir exists
	err := os.MkdirAll(tempDirName, 0700)
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory")
	}

	// Create temporary .scad file
	scadFileName := template.UUID.String() + "-" + utils.CreateRandomString(6) + ".scad"
	scadFile, err := os.Create(path.Join(tempDirName, scadFileName))
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file for %s", template.UUID)
	}
	// TODO: uncomment
	//defer os.Remove(scadFile.Name())

	// Include all libraries
	scadFile.WriteString(openScadLibrariesIncludes)

	// Add parameters
	for _, param := range template.Params {
		scadParam := param.Key + "=" + param.Value + "\n"
		_, err := scadFile.WriteString(scadParam)
		if err != nil {
			return nil, fmt.Errorf("failed to write param %s=%s for %s", param.Key, param.Value, template.UUID)
		}
	}

	// Write template code
	scadFile.Write(template.Template)

	// Generate model
	cmd := exec.Command("openscad")

	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run openSCAD for %s", template.UUID)
	}

	return nil, nil
}
