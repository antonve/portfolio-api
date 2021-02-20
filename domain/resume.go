package domain

import (
	"github.com/srvc/fail"
)

// Resume is a container for a JSON resume blob
type Resume struct {
	Slug    string `json:"slug" db:"slug"`
	Body    string `json:"body" db:"body"`
	Enabled bool   `json:"enabled" db:"enabled"`
}

// ResumeLog is used to keep track of visits to a resume
type ResumeLog struct {
	ID        uint64 `json:"id" db:"id"`
	Slug      string `json:"slug" db:"slug"`
	IPAddress string `json:"ip_address" db:"ip_address"`
	UserAgent string `json:"user_agent" db:"user_agent"`
}

// Validate checks if a resume log is valid
func (r ResumeLog) Validate() (bool, error) {
	if r.ID == 0 {
		return false, fail.New("invalid `ID` supplied")
	}
	if len(r.Slug) == 0 {
		return false, fail.New("no `Slug` supplied")
	}
	if len(r.IPAddress) == 0 {
		return false, fail.New("invalid `IPAddress` supplied")
	}
	if len(r.UserAgent) == 0 {
		return false, fail.New("invalid `UserAgent` supplied")
	}

	return true, nil
}
