package structbot

type ConfigureFileType int

const (
	Json ConfigureFileType = iota
	Yaml
	Xml
	Env
)

func (c ConfigureFileType) getTagString() string {
	return [...]string{"json", "yaml", "xml", "env"}[c]
}
