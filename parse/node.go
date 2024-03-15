package parse

type Ini struct {
	FileName string
	Sections []*Section
}

type Section struct {
	Name      string
	KeyValues []*KeyValue
}

type KeyValue struct {
	Key   string
	Value string
}
