package types

type Field struct {
	Name      string
	IsArray   bool
	IsPointer bool
	Type      Type
	Tags      []Tag
}
