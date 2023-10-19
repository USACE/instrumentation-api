package message

// Message is a response message; typically used for 400/500 level API responses
// Used in cases where error cannot be sent back to client;
// when it potentially contains information about resouces for which a user is unauthorized

const Unauthorized = "User is not authorized to use this resource"

const NotFound = "Not Found"

const BadRequest = "Bad Request"

const InternalServerError = "Internal Server Error"

const MalformedID = "Malformed ID"

func MissingQueryParameter(param string) string {
	return "Missing query parameter " + param
}

func MatchRouteParam(param string) string {
	return "Object " + param + " does not match Route Param"
}
