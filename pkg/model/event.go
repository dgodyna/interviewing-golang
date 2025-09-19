// Package model defines the core data structures for network event processing.
// This package contains the Event type which represents a single network event
// with all required fields for downstream rating and billing systems.
package model

import "time"

// Event represents a network communication event (call, SMS, data session) that will be
// processed by downstream rating systems. Each event contains billing-relevant information
// and must maintain data integrity for accurate charging and reporting.
//
// The event structure matches the database schema requirements and supports
// high-throughput generation and processing scenarios.
type Event struct {
	// EventSource identifies the client/system that generated this event.
	// This field ensures event traceability and supports multi-tenant scenarios.
	EventSource int `json:"event_source"`

	// EventRef is a globally unique identifier for this specific event.
	// Must be unique across all events to prevent duplicate processing.
	EventRef string `json:"event_ref"`

	// EventType categorizes the service type for rating purposes.
	// Valid values: 1 (15% - Standard calls), 2 (20% - Premium services),
	// 3 (20% - International calls), 5 (45% - Complex routing).
	// Distribution must be maintained precisely for accurate cost modeling.
	EventType int `json:"event_type"`

	// EventDate records the exact timestamp when the event occurred.
	// Used for time-based rating, peak/off-peak calculations, and reporting.
	EventDate time.Time `json:"event_date"`

	// CallingNumber identifies the originating party (caller's phone number).
	// Required for customer billing and fraud detection.
	CallingNumber int `json:"calling_number"`

	// CalledNumber identifies the destination party (recipient's phone number).
	// Used for routing decisions and interconnect billing.
	CalledNumber int `json:"called_number"`

	// Location specifies the geographic area where the event originated.
	// Critical for location-based rating and regulatory compliance.
	Location string `json:"location"`

	// DurationSeconds contains the total duration of the event in seconds.
	// Primary metric for billing calculations in voice services.
	DurationSeconds int `json:"duration_seconds"`

	// Attr1 through Attr8 provide extensible storage for additional event metadata.
	// These fields support custom attributes without requiring schema changes,
	// enabling flexible business rule implementation and feature extensions.

	// Attr1 stores custom attribute data (e.g., service quality metrics).
	Attr1 string `json:"attr_1"`
	// Attr2 stores custom attribute data (e.g., routing information).
	Attr2 string `json:"attr_2"`
	// Attr3 stores custom attribute data (e.g., device information).
	Attr3 string `json:"attr_3"`
	// Attr4 stores custom attribute data (e.g., promotion codes).
	Attr4 string `json:"attr_4"`
	// Attr5 stores custom attribute data (e.g., network quality indicators).
	Attr5 string `json:"attr_5"`
	// Attr6 stores custom attribute data (e.g., customer segment).
	Attr6 string `json:"attr_6"`
	// Attr7 stores custom attribute data (e.g., partner identifiers).
	Attr7 string `json:"attr_7"`
	// Attr8 stores custom attribute data (e.g., compliance flags).
	Attr8 string `json:"attr_8"`
}
