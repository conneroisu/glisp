package domain

import (
	"fmt"
	"net/url"
)

// Text Document Request Methods
const (
	// MethodRequestTextDocumentDidOpen is the text document did open
	// request method for the language server protocol.
	//
	// Microsoft LSP Docs:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_didOpen
	MethodRequestTextDocumentDidOpen Method = "textDocument/didOpen"

	// MethodRequestTextDocumentCompletion is the text document completion
	// request method.
	//
	// Microsoft LSP Docs:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_completion
	MethodRequestTextDocumentCompletion Method = "textDocument/completion"

	// MethodRequestTextDocumentHover is the text document hover request
	// method.
	//
	// Microsoft LSP Docs:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_hover
	MethodRequestTextDocumentHover Method = "textDocument/hover"

	// MethodRequestTextDocumentSignatureHelp is the text document signature
	// help method for the LSP
	//
	// Microsoft LSP Docs:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_signatureHelp
	MethodRequestTextDocumentSignatureHelp Method = "textDocument/signatureHelp"

	// MethodRequestTextDocumentDefinition is the text document definition
	// method for the LSP
	//
	// Microsoft LSP Docs:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_definition
	MethodRequestTextDocumentDefinition Method = "textDocument/definition"

	// MethodTextDocumentReferences is the text document references method
	// for the LSP
	//
	// Microsoft LSP Docs:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_references
	MethodTextDocumentReferences Method = "textDocument/references"

	// MethodRequestTextDocumentDocumentHighlight is the text document
	// document highlight method for the LSP
	//
	// Microsoft LSP Docs:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_documentHighlight
	MethodRequestTextDocumentDocumentHighlight Method = "textDocument/documentHighlight"

	// MethodRequestTextDocumentDocumentSymbol is the text document document symbol method for the LSP
	//
	// Microsoft LSP Docs:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_documentSymbol
	MethodRequestTextDocumentDocumentSymbol Method = "textDocument/documentSymbol"

	// MethodTextDocumentFormatting is the text document formatting method for the LSP
	//
	// Microsoft LSP Docs:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_formatting
	MethodTextDocumentFormatting Method = "textDocument/formatting"

	// MethodTextDocumentDidClose is the text document did close notification method for the language server protocol.
	//
	// Microsoft LSP Docs:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_didClose
	MethodTextDocumentDidClose Method = "textDocument/didClose"

	// MethodTextDocumentRangeFormatting is the text document range formatting method for the LSP
	//
	// Microsoft LSP Docs:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_rangeFormatting
	MethodTextDocumentRangeFormatting Method = "textDocument/rangeFormatting"

	// MethodTextDocumentOnTypeFormatting is the text document on type formatting method for the LSP
	//
	// Microsoft LSP Docs:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_onTypeFormatting
	MethodTextDocumentOnTypeFormatting Method = "textDocument/onTypeFormatting"

	// MethodTextDocumentRename is the text document code action method for
	// renaming a symbol in the language server protocol.
	//
	// Microsoft LSP Docs:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_rename
	MethodTextDocumentRename Method = "textDocument/rename"

	// MethodRequestTextDocumentCodeAction is the text document code action method for the LSP
	//
	// Microsoft LSP Docs:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_codeAction
	MethodRequestTextDocumentCodeAction Method = "textDocument/codeAction"

	// MethodTextDocumentCodeLens is the text document code lens method for
	// the LSP
	//
	// Microsoft LSP Docs:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_codeLens
	MethodTextDocumentCodeLens Method = "textDocument/codeLens"

	// MethodTextDocumentDocumentLink is the text document document link method for the LSP
	//
	// Microsoft LSP Docs:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_documentLink
	MethodTextDocumentDocumentLink Method = "textDocument/documentLink"
)

// TextDocumentURI is a URI for a document.
// Over the wire, it will still be transferred as a string, but this guarantees
// that the contents of that string can be parsed as a valid URI.
type TextDocumentURI struct {
	URL url.URL `json:"url"`
}

// NotificationDidOpenTextDocument is a notification that is sent when
// the client opens a text document.
//
// Microsoft LSP Docs:
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_didOpen
type NotificationDidOpenTextDocument struct {
	// DidOpenTextDocumentNotification embeds the Notification struct
	Notification
	// Params are the parameters for the notification.
	Params DidOpenTextDocumentParams `json:"params"`
}

// DidOpenTextDocumentParams contains the text document after it has been opened.
type DidOpenTextDocumentParams struct {
	// TextDocument is the text document after it has been opened.
	TextDocument TextDocumentItem `json:"textDocument"`
}

// DidSaveTextDocumentParamsNotification is a notification for when
// the client saves a text document.
//
// Microsoft LSP Docs:
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_didSave
type DidSaveTextDocumentParamsNotification struct {
	// DidSaveTextDocumentParams embeds the Notification struct
	Notification
	// Params are the parameters for the notification.
	Params DidSaveTextDocumentParams `json:"params"`
}

// DidSaveTextDocumentParams contains the text document after it has been saved.
type DidSaveTextDocumentParams struct {
	// TextDocument is the text document after it has been saved.
	TextDocument TextDocumentURI `json:"textDocument"`
}

// CodeActionRequest is a request for a code action to the language server.
//
// Microsoft LSP Docs:
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_code
type CodeActionRequest struct {
	// CodeActionRequest embeds the Request struct
	Request
	// Params are the parameters for the code action request.
	Params TextDocumentCodeActionParams `json:"params"`
}

// TextDocumentCodeActionParams are the parameters for a code action request.
type TextDocumentCodeActionParams struct {
	// TextDocument is the text document for the code action request.
	TextDocument TextDocumentURI `json:"textDocument"`
	// Range is the range for the code action request.
	Range Range `json:"range"`
	// Context is the context for the code action request.
	Context CodeActionContext `json:"context"`
}

// TextDocumentCodeActionResponse is the response for a code action request.
type TextDocumentCodeActionResponse struct {
	// TextDocumentCodeActionResponse embeds the Response struct
	Response
	// Result is the result for the code action request.
	Result []CodeAction `json:"result"`
}

// Method returns the method for the code action response
func (r TextDocumentCodeActionResponse) Method() string {
	return "textDocument/codeAction"
}

// CodeActionContext is the context for a code action request.
type CodeActionContext struct {
	// Add fields for CodeActionContext as needed
}

// CodeAction is a code action for a given text document.
type CodeAction struct {
	// Title is the title for the code action.
	Title string `json:"title"`
	// Edit is the edit for the code action.
	Edit *WorkspaceEdit `json:"edit,omitempty"`
	// Command is the command for the code action.
	Command *Command `json:"command,omitempty"`
}

// Command is a command for a given text document.
//
// Microsoft LSP Docs:
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#command
type Command struct {
	// Title is the title for the command.
	Title string `json:"title"`
	// Command is the command for the command.
	Command string `json:"command"`
	// Arguments are the arguments for the command.
	Arguments []interface{} `json:"arguments,omitempty"`
}

// CompletionRequest is a request for a completion to the language server
//
// Microsoft LSP Docs:
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_completion
type CompletionRequest struct {
	// CompletionRequest embeds the Request struct
	Request
	// Params are the parameters for the completion request
	Params CompletionParams `json:"params"`
}

// CompletionParams is a struct for the completion params
type CompletionParams struct {
	// CompletionParams embeds the TextDocumentPositionParams struct
	TextDocumentPositionParams
}

// CompletionResponse is a response for a completion to the language server
//
// Microsoft LSP Docs:
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_completion
type CompletionResponse struct {
	// CompletionResponse embeds the Response struct
	Response
	// Result is the result of the completion request
	Result []CompletionItem `json:"result"`
}

// CompletionItem is a struct for a completion item
//
// Microsoft LSP Docs:
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_completion
type CompletionItem struct {
	// Label is the label for the completion item
	Label string `json:"label"`
	// Detail is the detail for the completion item
	Detail string `json:"detail"`
	// Documentation is the documentation for the completion item
	Documentation string `json:"documentation"`
	// Kind is the kind of the completion item
	Kind CompletionItemKind `json:"kind"`
}

// CompletionItemTag is a struct for a completion item tag
type CompletionItemTag struct {
	Deprecated bool `json:"deprecated"`
}

// TextDocumentItem is a text document.
type TextDocumentItem struct {
	// URI is the uri for the text document.
	URI string `json:"uri"`

	// LanguageID is the language id for the text document.
	LanguageID string `json:"languageId"`

	// Version is the version number of a given text document.
	Version int `json:"version"`

	// Text is the text of the text document.
	Text string `json:"text"`
}

// VersionTextDocumentIdentifier is a text document with a version number.
type VersionTextDocumentIdentifier struct {
	// VersionTextDocumentIdentifier embeds the TextDocumentIdentifier struct
	TextDocumentURI
	// Version is the version number for the text document.
	Version int `json:"version"`
}

// TextDocumentPositionParams is a text document position parameters.
type TextDocumentPositionParams struct {
	// TextDocument is the text document for the position parameters.
	TextDocument TextDocumentURI `json:"textDocument"`
	// Position is the position for the text document.
	Position Position `json:"position"`
}

// Position is a position inside a text document.
type Position struct {
	// Line is the line number for the position (zero-based).
	Line int `json:"line"`
	// Character is the character number for the position (zero-based).
	Character int `json:"character"`
}

// String returns a string representation of the position.
func (p Position) String() string {
	return fmt.Sprintf("Line: %d, Character: %d", p.Line, p.Character)
}

// Location is a location inside a resource, such as a line
// inside a text file.
type Location struct {
	// URI is the uri for the location.
	URI string `json:"uri"`
	// Range is the range for the location.
	Range Range `json:"range"`
}

// Range is a range in a text document.
type Range struct {
	// Start is the start of a given range.
	Start Position `json:"start"`
	// End is the end of a given range.
	End Position `json:"end"`
}

// WorkspaceEdit is the workspace edit object.
type WorkspaceEdit struct {
	// Changes is the changes for the workspace edit.
	Changes map[string][]TextEdit `json:"changes"`
}

// TextEdit represents an edit operation on a single text document.
type TextEdit struct {
	// Range is the range for the text edit.
	Range Range `json:"range"`
	// NewText is the new text for the text edit.
	NewText string `json:"newText"`
}

// HoverRequest is sent from the client to the server to request hover
// information.
//
// Microsoft LSP Docs:
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_hover
type HoverRequest struct {
	// HoverRequest embeeds the request struct.
	Request
	// Params are the parameters for the hover request.
	Params HoverParams `json:"params"`
}

// HoverParams is the parameters for a hover request.
type HoverParams struct {
	// TextDocumentPositionParams is the text document position parameters.
	TextDocumentPositionParams
}

// HoverResponse is the response from the server to a hover request.
type HoverResponse struct {
	// Response is the response for the hover request.
	Response
	// Result is the result for the hover request.
	Result HoverResult `json:"result"`
}

// Method returns the method for the hover response
func (r HoverResponse) Method() string {
	return "textDocument/hover"
}

// HoverResult is a result from a hover request to the client from the
// language server.
type HoverResult struct {
	// Contents is the contents for the hover result.
	Contents string `json:"contents"`
}

// TextDocumentDidChangeNotification is sent from the client to the server to signal
// that the content of a text document has changed.
//
// Microsoft LSP Docs:
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_didChange
type TextDocumentDidChangeNotification struct {
	// TextDocumentDidChangeNotification embeds the Notification struct
	Notification
	// Params are the parameters for the notification.
	Params DidChangeTextDocumentParams `json:"params"`
}

// DidChangeTextDocumentParams is sent from the client to the server to signal
// that the content of a text document has changed.
//
// Microsoft LSP Docs:
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_didChange
type DidChangeTextDocumentParams struct {
	// TextDocument is the  document that did change. The version number points to the version
	TextDocument VersionTextDocumentIdentifier `json:"textDocument"`
	// ContentChanges is the array of content changes.
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

// TextDocumentContentChangeEvent is sent from the client to the server to signal
// that the content of a text document has changed.
type TextDocumentContentChangeEvent struct {
	// The new text of the whole document.
	Text string `json:"text"`
}

// CompletionItemKind is an enum for completion item kinds.
type CompletionItemKind int

const (
	// Text is a completion item kind
	Text CompletionItemKind = iota + 1
	// CompletionItemKindMethod is a completion item kind for a method or function completion
	CompletionItemKindMethod
	// CompletionItemKindFunction is a completion item kind for a function completion
	CompletionItemKindFunction
	// CompletionItemKindConstructor is a completion item kind for a constructor completion
	CompletionItemKindConstructor
	// CompletionItemKindField is a completion item kind for a field completion
	CompletionItemKindField
	// CompletionItemKindVariable is a completion item kind for a variable completion
	CompletionItemKindVariable
	// CompletionItemKindClass is a completion item kind for a class completion
	CompletionItemKindClass
	// CompletionItemKindInterface is a completion item kind for an interface completion
	CompletionItemKindInterface
	// CompletionItemKindModule is a completion item kind for a module completion
	CompletionItemKindModule
	// CompletionItemKindProperty is a completion item kind for a property completion
	CompletionItemKindProperty
	// CompletionItemKindUnit is a completion item kind for a unit
	CompletionItemKindUnit
	// CompletionItemKindValue is a completion item kind for a value
	CompletionItemKindValue
	// CompletionItemKindEnum is a completion item kind for an enum
	CompletionItemKindEnum
	// CompletionItemKindKeyword is a completion item kind for a keyword
	CompletionItemKindKeyword
	// CompletionItemKindSnippet is a completion item kind for a snippet
	CompletionItemKindSnippet
	// CompletionItemKindColor is a completion item kind for a color
	CompletionItemKindColor
	// CompletionItemKindFile is a completion item kind for a file
	CompletionItemKindFile
	// CompletionItemKindReference is a completion item kind for a reference
	CompletionItemKindReference
	// CompletionItemKindFolder is a completion item kind for a folder
	CompletionItemKindFolder
	// CompletionItemKindEnumMember is a completion item kind for an enum member
	CompletionItemKindEnumMember
	// CompletionItemKindConstant is a completion item kind for a constant
	CompletionItemKindConstant
	// CompletionItemKindStruct is a completion item kind for a struct
	CompletionItemKindStruct
	// CompletionItemKindEvent is a completion item kind for an event
	CompletionItemKindEvent
	// CompletionItemKindOperator is a completion item kind for an operator
	CompletionItemKindOperator
	// CompletionItemKindTypeParameter is a completion item kind for a type parameter
	CompletionItemKindTypeParameter
)

// String returns the string representation of the CompletionItemKind.
func (c CompletionItemKind) String() string {
	return [...]string{
		"Text",
		"Method",
		"Function",
		"Constructor",
		"Field",
		"Variable",
		"Class",
		"Interface",
		"Module",
		"Property",
		"Unit",
		"Value",
		"Enum",
		"Keyword",
		"Snippet",
		"Color",
		"File",
		"Reference",
		"Folder",
		"EnumMember",
		"Constant",
		"Struct",
		"Event",
		"Operator",
		"TypeParameter",
	}[c-1]
}

// DidCloseTextDocumentParamsNotification is a struct for the did close text document params notification
//
// Microsoft LSP Docs:
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_didClose
type DidCloseTextDocumentParamsNotification struct {
	Notification
	Params DidCloseTextDocumentParamsNotificationParams `json:"params,required"`
}

// DidCloseTextDocumentParamsNotificationParams is a struct for the did close text document params notification params
//
// Microsoft LSP Docs:
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_didClose
type DidCloseTextDocumentParamsNotificationParams struct {
	TextDocument TextDocumentURI `json:"textDocument,required"`
}
