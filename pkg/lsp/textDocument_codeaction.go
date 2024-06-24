package lsp

type CodeActionRequest struct {
	Request
	Params TextDocumentCodeActionsParams `json:"params"`
}

type TextDocumentCodeActionsParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Range Range `json:"range"`
	Context CodeActionContext `json:"context"`
}

type TextDocumentCodeActionResponse struct {
	Response
	Result []CodeAction `json:"result"`
}

type CodeActionContext struct {}

type CodeAction struct {
	Title string `json:"title"`
	Edit *WorkspaceEdit `json:"edit,omitempty"`
	Command *Command `json:"command,omitempty"` 
}

type Command struct {
	Title string `json:"title"`
	Command string `json:"command"`
	Arguments []interface{} `json:"arguments,omitempty"`
}
