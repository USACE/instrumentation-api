package messages

// Message is a response message; typically used for 400/500 level API responses
// Used in cases where error cannot be sent back to client;
// when it potentially contains information about resouces for which a user is unauthorized

// Unauthorized is the a default unauthorized (403) message
var Unauthorized = "User is not authorized to use this resource"

// NotFound is the a default not found (404) message
var NotFound = "Not Found"

// BadRequest is the a default not found (404) message
var BadRequest = "Bad Request"

// InternalServerError is the a default not found (404) message
var InternalServerError = "Internal Server Error"

var MalformedID = "Malformed ID"

func MissingQueryParameter(param string) string {
	return "Missing query parameter " + param
}

func MatchRouteParam(param string) string {
	return "Object " + param + " does not match Route Param"
}
