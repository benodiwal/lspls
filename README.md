# Lspls

### What is LSP ?

Adding features like auto complete, go to definition, or documentation on hover for a programming language takes significant effort. Traditionally this work had to be repeated for each development tool, as each tool provides different APIs for implementing the same feature.

A Language Server is meant to provide the language-specific smarts and communicate with development tools over a protocol that enables inter-process communication.

The idea behind the Language Server Protocol (LSP) is to standardize the protocol for how such servers and development tools communicate. This way, a single Language Server can be re-used in multiple development tools, which in turn can support multiple languages with minimal effort.

LSP is a win for both language providers and tooling vendors!


### How it works ?

![Diagram](https://microsoft.github.io/language-server-protocol/overviews/lsp/img/language-server-sequence.png)

A language server runs as a separate process and development tools communicate with the server using the language protocol over JSON-RPC. Below is an example for how a tool and a language server communicate during a routine editing session:

- **The user opens a file (referred to as a document) in the tool**: The tool notifies the language server that a document is open (‘textDocument/didOpen’). From now on, the truth about the contents of the document is no longer on the file system but kept by the tool in memory. The contents now has to be synchronized between the tool and the language server.

- **The user makes edits**: The tool notifies the server about the document change (‘textDocument/didChange’) and the language representation of the document is updated by the language server. As this happens, the language server analyses this information and notifies the tool with the detected errors and warnings (‘textDocument/publishDiagnostics’).

- **The user executes “Go to Definition” on a symbol of an open document**: The tool sends a ‘textDocument/definition’ request with two parameters: (1) the document URI and (2) the text position from where the ‘Go to Definition’ request was initiated to the server. The server responds with the document URI and the position of the symbol’s definition inside the document.

- **The user closes the document (file)**: A ‘textDocument/didClose’ notification is sent from the tool informing the language server that the document is now no longer in memory. The current contents are now up to date on the file system.

### Specification:
You can refer to the officail specifiction part here:  https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification



### How to Build:

```sh
make build
```

### How to Run:
```sh
make run
```
