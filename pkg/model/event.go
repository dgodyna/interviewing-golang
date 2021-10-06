package model

import "time"

// Event is a single billable occurrence of product usage.
type Event struct {
	// EventSource is the source of the event, for example: telephone number.
	EventSource int `json:"event_source"`
	// EventRef is unique event identifier across all the events.
	EventRef string `json:"event_ref"`
	// EventType is type of service to which the event relates. Roaming/non-roaming.
	EventType int `json:"event_type"`
	// EventDate datetime of event.
	EventDate time.Time `json:"event_date"`
	// CallingNumber the person initiated event.
	CallingNumber int `json:"calling_number"`
	// CalledNumber the person received an event.
	CalledNumber int `json:"called_number"`
	// Location where event occurred.
	Location string `json:"location"`
	// DurationSeconds is duration of event in seconds.
	DurationSeconds int `json:"duration_seconds"`
	// Attr1 is configurable attribute number 1.
	Attr1 string `json:"attr_1"`
	// Attr2 is configurable attribute number 1.
	Attr2 string `json:"attr_2"`
	// Attr3 is configurable attribute number 1.
	Attr3 string `json:"attr_3"`
	// Attr4 is configurable attribute number 1.
	Attr4 string `json:"attr_4"`
	// Attr5 is configurable attribute number 1.
	Attr5 string `json:"attr_5"`
	// Attr6 is configurable attribute number 1.
	Attr6 string `json:"attr_6"`
	// Attr7 is configurable attribute number 1.
	Attr7 string `json:"attr_7"`
	// Attr8 is configurable attribute number 1.
	Attr8 string `json:"attr_8"`
}
