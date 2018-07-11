package thesaurus

type Thesuarus interface {
	Synonyms(term string) ([]string, error)
}