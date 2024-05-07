package lsp

type TextDocumentItem struct {
	URI string `json:"uri"`
	LanguageId string `json:"language_id"`
	Version int `json:"version"`
	Text string `json:"text"`
}