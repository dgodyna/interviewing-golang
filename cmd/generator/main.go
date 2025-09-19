// Package main implements a high-volume network event data generator for performance testing.
//
// This command-line application generates realistic network events (calls, SMS, data sessions)
// with proper probability distributions and randomized data fields. The generator is designed
// to simulate real-world telco mediation system outputs for downstream rating system testing.
//
// Performance Characteristics:
// - Current implementation: ~5s for 1M events (baseline)
// - Target optimization: Sub-second generation for 1M events
// - Memory usage: Scales linearly with event count (optimization opportunity)
//
// Usage:
//
//	generator <num_events> <output_file>
//
// Example:
//
//	generator 1000000 events.json
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/dmgo1014/interviewing-golang.git/pkg/generator"
	"github.com/dmgo1014/interviewing-golang.git/pkg/model"
	"github.com/google/uuid"
)

// main orchestrates the event generation process with timing measurements.
//
// The application follows a simple pipeline:
// 1. Parse and validate command-line arguments
// 2. Generate the specified number of events
// 3. Serialize all events to JSON format
// 4. Write the complete dataset to the output file
//
// Event Type Distribution (must be maintained precisely):
// - Type 1: 15% (Standard calls)
// - Type 2: 20% (Premium services)
// - Type 3: 20% (International calls)
// - Type 5: 45% (Complex routing)
//
// Command-line Arguments:
// 1. num_events: Integer specifying the number of events to generate
// 2. output_file: Path where the JSON event data will be saved
func main() {
	// Performance monitoring: Track total execution time for benchmarking
	start := time.Now()
	defer func() {
		fmt.Println("================")
		fmt.Printf("Execution Time : %v\n", time.Since(start))
	}()

	// Input validation: Ensure exactly 2 arguments are provided
	if len(os.Args) != 3 {
		panic(fmt.Errorf("invalid number of arguments, 2 expected, got %d", len(os.Args)-1))
	}

	// Parse event count from the first command-line argument
	numEventsStr := os.Args[1]
	numEvents, err := strconv.Atoi(numEventsStr)
	if err != nil {
		panic(fmt.Errorf("unable to parse number of events : %+v", err))
	}

	// Extract output file path from second command-line argument
	outPutFile := os.Args[2]

	fmt.Printf("number event : %d\n", numEvents)
	fmt.Printf("dump output: %s\n", outPutFile)

	// Event generation phase: Create all events in memory
	events := []*model.Event{}

	// Sequential event generation loop
	for i := 0; i < numEvents; i++ {
		events = append(events, generateEvent())
	}

	// Convert events to JSON format
	content, err := json.Marshal(events)
	if err != nil {
		panic(fmt.Errorf("unable to marshall events : %+v", err))
	}

	// File output phase: Write complete JSON to disk
	err = os.WriteFile(outPutFile, content, 0777)
	if err != nil {
		panic(fmt.Errorf("unable to write file : %+v", err))
	}
}

// generateEvent creates a single network event with realistic random data.
//
// This function populates all required fields of a model.Event with appropriate
// random values that simulate real-world network event characteristics:
//
// Field Generation Strategy:
// - EventSource/Numbers: Large integers simulating phone numbers (up to 88005553535)
// - EventRef: UUID for guaranteed uniqueness across all events
// - EventType: Probability-based selection maintaining required distribution
// - EventDate: Historical timestamp between 2010-2020
// - Location: Random alphanumeric string representing geographic codes
// - DurationSeconds: Random duration 0-99 seconds for call length simulation
// - Attr1-8: Random strings for extensible metadata storage
func generateEvent() *model.Event {
	return &model.Event{
		EventSource:     rand.Intn(88005553535),   // Simulated client identifier
		EventRef:        uuid.New().String(),      // Guaranteed unique event ID
		EventType:       generateEventType(),      // Probability-distributed type
		EventDate:       *generator.RandomDate(),  // Historical timestamp
		CallingNumber:   rand.Intn(88005553535),   // Originating phone number
		CalledNumber:    rand.Intn(88005553535),   // Destination phone number
		Location:        generator.RandomString(), // Geographic location code
		DurationSeconds: rand.Intn(100),           // Call duration 0-99 seconds
		Attr1:           generator.RandomString(), // Custom attribute 1
		Attr2:           generator.RandomString(), // Custom attribute 2
		Attr3:           generator.RandomString(), // Custom attribute 3
		Attr4:           generator.RandomString(), // Custom attribute 4
		Attr5:           generator.RandomString(), // Custom attribute 5
		Attr6:           generator.RandomString(), // Custom attribute 6
		Attr7:           generator.RandomString(), // Custom attribute 7
		Attr8:           generator.RandomString(), // Custom attribute 8
	}
}

// generateEventType produces event types following the required probability distribution.
//
// The distribution reflects real-world telco traffic patterns where complex routing
// events (type 5) are most common, while standard calls (type 1) are least frequent.
// This distribution directly impacts downstream rating system resource planning.
//
// Probability Mapping:
// - Random 0-14   (15 values) → Type 1 (15%): Standard calls
// - Random 15-34  (20 values) → Type 2 (20%): Premium services
// - Random 35-54  (20 values) → Type 3 (20%): International calls
// - Random 55-99  (45 values) → Type 5 (45%): Complex routing
func generateEventType() int {
	r := rand.Intn(100)

	if r < 15 {
		return 1 // 15% - Standard calls
	} else if r < 35 {
		return 2 // 20% - Premium services
	} else if r < 55 {
		return 3 // 20% - International calls
	}
	return 5 // 45% - Complex routing
}
