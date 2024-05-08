package analysis

import (
	"fmt"

	"github.com/benodiwal/lspls/lsp"
)

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

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	document := s.Documents[uri]

	return lsp.HoverResponse {
				Response: lsp.Response{
					RPC: "2.0",
					ID: &id,
				},
				Result: lsp.HoverResult {
					Contents: fmt.Sprintf("File: %s, Characters: %d", uri, len(document)),
				},
			}
}

func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {

	return lsp.DefinitionResponse {
				Response: lsp.Response{
					RPC: "2.0",
					ID: &id,
				},
				Result: lsp.Location{
					URI: uri,
					Range: lsp.Range {
						Start: lsp.Position{
							Line: position.Line - 1,
							Character: 0,
						},
						End: lsp.Position{
							Line: position.Line - 1,
							Character: 0,
						},
					},
				},
			}
}