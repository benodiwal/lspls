package lsp

type TextDocumentItem struct {
	URI string `json:"uri"`
	LanguageId string `json:"language_id"`
	Version int `json:"version"`
	Text string `json:"text"`
}

type TextDocumentIdentifier struct {
	URI string `json:"uri"`
}

type VersionTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}