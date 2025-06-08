package openscadgenerator

const tempDirName string = "/dev/shm/ParaCAD"
const openScadLibrariesIncludes string = ``

type OpenSCADGenerator struct{}

func NewOpenSCADGenerator() OpenSCADGenerator {
	return OpenSCADGenerator{}
}
