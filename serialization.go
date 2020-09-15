package structbot

type SerializationType int

const (
	Json SerializationType = iota
	Yaml
	Xml
	Env
	Unknown
)

func (s SerializationType) getTagString() string {
	return [...]string{"json", "yaml", "xml", "env"}[s]
}

func (s SerializationType) String() string {
	return [...]string{"Json", "Yaml", "Xml", "Env"}[s]
}
