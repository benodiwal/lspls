package analysis

type State struct {
	Documents map[string]string // map of file names to content of the files
}

func NewState() State {
	return State {
		Documents: map[string]string{},
	}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}