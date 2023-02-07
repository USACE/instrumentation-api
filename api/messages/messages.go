package messages

// Message is a response message; typically used for 400/500 level API responses
// Used in cases where error cannot be sent back to client;
// when it potentially contains information about resouces for which a user is unauthorized
type Message struct {
	Message string `json:"message"`
}

// Unauthorized is the a default unauthorized (403) message
var Unauthorized = Message{Message: "User is not authorized to use this resource"}

// NotFound is the a default not found (404) message
var NotFound = Message{Message: "Not Found"}

// BadRequest is the a default not found (404) message
var BadRequest = Message{Message: "Bad Request"}

// InternalServerError is the a default not found (404) message
var InternalServerError = Message{Message: "Internal Server Error"}

var MalformedID = Message{Message: "Malformed ID"}

func MissingQueryParameter(param string) *Message {
	return &Message{Message: "Missing query parameter " + param}
}

func MatchRouteParam(param string) *Message {
	return &Message{Message: "Object " + param + " does not match Route Param"}
}
