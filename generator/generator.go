package generator

type Generator interface {
	GenerateModel(template FilledTemplate) ([]byte, error)
}
