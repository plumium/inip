package parse

type Ini struct {
	FileName string     `json:"fileName"`
	Sections []*Section `json:"sections"`
}

type Section struct {
	Name      string      `json:"name"`
	KeyValues []*KeyValue `json:"keyValues"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
