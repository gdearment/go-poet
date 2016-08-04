package gopoet

type CodeBlock interface {
	String() string
	Packages() []ImportSpec
}

type MethodSpec struct {
	// CodeBlock

	FuncSpec
	ReceiverName string
	Receiver     *StructSpec
}

type Identifier struct {
	Name string
	Type ImportSpec
}

type IdentifierParameter struct {
	Identifier
}

type IdentifierField struct {
	Identifier
	Tag string
}

type ImportSpec interface {
	GetPackage() string
	GetPackageAlias() string
	NeedsQualifier() bool
	GetName() string
}

type InterfaceSpec struct {
	// CodeBlock

	Name               string
	Comment            string
	EmbeddedInterfaces []ImportSpec
	Methods            []FuncSpec
}

type StructSpec struct {
	// CodeBlock

	Name    string
	Comment string
	Fields  []IdentifierField
}

type Statement struct {
	// CodeBlock

	Format    string
	Arguments []interface{}
	Indent    int
}