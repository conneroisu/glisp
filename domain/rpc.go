package domain

// Message is a message in the language server protocol.
//
// Microsoft LSP Docs:
// https://microsoft.github.io/language-server-protocol/specification#message
type Message[
	T any,
] struct {
	// RPC is the rpc method for the message.
	RPC string `json:"jsonrpc"`
	// ID is the id of the message.
	ID *int `json:"id,omitempty"`
	// Params is the params of the message.
	Params T `json:"params,omitempty"`
	// Method is the method of the message.
	Method *Method `json:"method,omitempty"`
}
