package parse

func Parse(name string, input string) *Ini {
	lex := NewLexer(name, input)
	ini := &Ini{FileName: name, Sections: make([]*Section, 0)}
	go lex.Run()
	var (
		sec *Section
		kv  *KeyValue
	)
	for t := range lex.Token {
		switch t.Type {
		case tokenSection:
			sec = &Section{Name: t.Value, KeyValues: make([]*KeyValue, 0)}
			ini.Sections = append(ini.Sections, sec)
		case tokenKey:
			kv = &KeyValue{Key: t.Value, Value: ""}
		case tokenValue:
			kv.Value = t.Value
			sec.KeyValues = append(sec.KeyValues, kv)
		}
	}
	return ini
}
