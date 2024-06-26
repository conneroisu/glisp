package domain

import "strings"

// Request is a request to a language server.
//
// Microsoft LSP Docs:
// https://microsoft.github.io/language-server-protocol/specification#requestMessage
type Request struct {
	// RPC is the rpc method for the request
	RPC string `json:"jsonrpc"`
	// ID is the id of the request
	ID int `json:"id,omitempty"`
	// Method is the method for the request
	Method string `json:"method"`
}

// Response is the response in an lanaguage server request.
type Response struct {
	// RPC is the rpc method for the response
	RPC string `json:"jsonrpc"`
	// ID is the id of the response
	ID int `json:"id,omitempty"`
	// Result is the result of the response
	Result interface{} `json:"result,omitempty"`
	// Error is the error of the response
	Error *Error `json:"error,omitempty"`
}

// Error is the error of a language server.
type Error struct {
	// Code is the code of the error
	Code ErrorCode `json:"code"`
	// Message is the message of the error
	Message string `json:"message"`
	// Data is the data of the error
	Data interface{} `json:"data,omitempty"`
}

// Notification is a notification from a LSP
type Notification struct {
	// RPC is the rpc method for the notification.
	RPC string `json:"jsonrpc"`
	// Method is the method for the notification.
	Method string `json:"method"`
}

// Method represents an lanaguage server method.
type Method string

// String returns the string representation of the method
func (m Method) String() string {
	return string(m)
}

// Prefix returns true if the method starts with the input
func (m Method) Prefix(input string) bool {
	return strings.HasPrefix(string(m), input)
}
