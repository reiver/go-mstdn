package local

import (
	"github.com/reiver/go-mstdn/ent"
)

// The "/api/v1/streaming/public/local" API end-points returns a series of events.
// Event is used to represent one of those events.
//
// And event has a 'name'.
// Currently the possible names are:
//
// • "delete"
// • "update"
//
// And event also has a status.
//
// If the event is an "update", then all of (or most of) the 'status' should be filled in.
//
// If the event is a "delete", then only the 'ID' will be filled in.
type Event struct {
	Name string
	Status ent.Status
}
