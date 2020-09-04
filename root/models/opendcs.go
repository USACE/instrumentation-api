package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Site is an instrument, represented as an OpenDCS Site
type Site struct {
	Elevation      string   `xml:"Elevation"`
	ElevationUnits string   `xml:"ElevationUnits"`
	Description    string   `xml:"Description"`
	SiteName       SiteName `xml:"SiteName"`
}

// SiteName is SiteName
type SiteName struct {
	ID       uuid.UUID `xml:",chardata"`
	NameType string    `xml:",attr"`
}

// AsSite returns an instrument represented as an OpenDCS Site
func (n *Instrument) AsSite() Site {
	return Site{
		Elevation:      "",
		ElevationUnits: "",
		Description:    n.Name,
		SiteName: SiteName{
			ID:       n.ID,
			NameType: "uuid",
		},
	}
}

// ListOpendcsSites returns an array of instruments from the database
// And formats them as OpenDCS Sites
func ListOpendcsSites(db *sqlx.DB) ([]Site, error) {

	nn, err := ListInstruments(db)
	if err != nil {
		return make([]Site, 0), err
	}
	ss := make([]Site, len(nn))
	for idx := range nn {
		ss[idx] = nn[idx].AsSite()
	}
	return ss, nil
}
