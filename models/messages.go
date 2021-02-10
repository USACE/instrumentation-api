package models

// Message is a response message; typically used for 400/500 level API responses
type Message struct {
	Message string `json:"message"`
}

// DefaultMessageUnauthorized is the a default unauthorized (403) message
var DefaultMessageUnauthorized = Message{Message: "User is not authorized to use this resource"}

// DefaultMessageNotFound is the a default not found (404) message
var DefaultMessageNotFound = Message{Message: "Not Found"}

// DefaultMessageBadRequest is the a default not found (404) message
var DefaultMessageBadRequest = Message{Message: "Bad Request"}

// DefaultMessageInternalServerError is the a default not found (404) message
var DefaultMessageInternalServerError = Message{Message: "Internal Server Error"}
