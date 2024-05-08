package analysis

import (
	"fmt"
	"strings"

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

	line := position.Line
	if line <= 0 {
		line += 1
	}

	return lsp.DefinitionResponse {
				Response: lsp.Response{
					RPC: "2.0",
					ID: &id,
				},
				Result: lsp.Location{
					URI: uri,
					Range: lsp.Range {
						Start: lsp.Position{
							Line: line - 1,
							Character: 0,
						},
						End: lsp.Position{
							Line: line - 1,
							Character: 0,
						},
					},
				},
			}
}

func (s *State) TextDocumentCodeAction(id int, uri string) lsp.TextDocumentCodeActionResponse {
	text := s.Documents[uri]

	actions := []lsp.CodeAction{}

	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")

		// if "VS Code" is present
		if idx >= 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit {
				{
					Range: LineRange(row, idx, idx+len("VS Code")),
					NewText: "Neovim",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Replace VS Code with a superior editor",
				Edit: &lsp.WorkspaceEdit{Changes: replaceChange},
			})

			censorChange := map[string][]lsp.TextEdit{}
			censorChange[uri] = []lsp.TextEdit {
				{
					Range: LineRange(row, idx, idx+len("VS Code")),
					NewText: "VS C*de",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Censor to VS C*de",
				Edit: &lsp.WorkspaceEdit{Changes: censorChange},
			})

			humbleChange := map[string][]lsp.TextEdit{}
			humbleChange[uri] = []lsp.TextEdit{
				{
					Range: LineRange(row, idx, idx+len("VS Code")),
					NewText: "VS Code is also good :))",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "I am humble !!",
				Edit: &lsp.WorkspaceEdit{Changes: humbleChange},
			})

		}
	}

	return lsp.TextDocumentCodeActionResponse {
		Response: lsp.Response{
			RPC: "2.0",
			ID: &id,
		},
		Result: actions,
	}

}

func (s *State) TextDocumentCompletion(id int) lsp.CompletionResponse {

	items := []lsp.CompletionItem {
		{
			Label: "Neovim (BTW)",
			Detail: "Cool Editor !!",
			Documentation: "Hey there !!",
		},
	}

	return lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID: &id,
		},
		Result: items,
	}

}

func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line: line,
			Character: start,
		},
		End: lsp.Position{
			Line: line,
			Character: end,
		},
	}
}