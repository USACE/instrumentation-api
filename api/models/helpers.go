package models

import "bytes"

// JSONType is a helper to infer the JSON type from []byte; object or array
// https://stackoverflow.com/questions/55014001/check-if-json-is-object-or-array
func JSONType(b []byte) string {
	leadChars := bytes.TrimLeft(b, " \t\r\n")
	if len(leadChars) > 0 && leadChars[0] == '[' {
		return "ARRAY"
	} else if len(leadChars) > 0 && leadChars[0] == '{' {
		return "OBJECT"
	}
	return ""
}
