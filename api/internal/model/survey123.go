package model

import "github.com/google/uuid"

type Survey123Payload struct {
	Survey123ID uuid.UUID `json:"survey123_id"`
}

type Survey123EquivalencyTable struct {
	Survey123ID uuid.UUID                        `json:"survey123_id" db:"survey123_id"`
	Rows        dbJSONSlice[EquivalencyTableRow] `json:"rows" db:"fields"`
}
