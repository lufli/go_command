package thesaurus

type BigHuge struct {
	APIKey string
}

type words struct {
	Syn []string `json: "syn"`
}
type synonyms struct {
	Noun *words `json: "noun"`
	Verb *words `json: "verb"`
}

func (b *BigHuge) Synonuyms(term string) ([]string, error) {
	var syns []string
}