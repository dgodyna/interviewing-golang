package main

import (
	"encoding/json"
	"fmt"
	"github.com/dmgo1014/interviewing-golang.git/pkg/generator"
	"github.com/dmgo1014/interviewing-golang.git/pkg/model"
	"github.com/google/uuid"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Generates events and saves them to the specified output file.
// The generated events will have the following types and associated probabilities:
// - Event type '1': 15%
// - Event type '2': 20%
// - Event type '3': 20%
// - Event type '5': 45%
//
// All other fields in the events will be filled with random values.
//
// Arguments:
// 1. The number of events to generate.
// 2. The path to the output file where the events will be saved.
func main() {

	// log time duration on application shutdown
	start := time.Now()
	defer func() {
		fmt.Println("================")
		fmt.Printf("Execution Time : %v\n", time.Since(start))
	}()

	// validate inputs firstly
	if len(os.Args) != 3 {
		panic(fmt.Errorf("invalid number of arguments, 2 expected, got %d", len(os.Args)-1))
	}

	// number of events is the first argument
	numEventsStr := os.Args[1]
	numEvents, err := strconv.Atoi(numEventsStr)
	if err != nil {
		panic(fmt.Errorf("unable to parse number of events : %+v", err))
	}

	// output file is the second
	outPutFile := os.Args[2]

	fmt.Printf("number event : %d\n", numEvents)
	fmt.Printf("dump output: %s\n", outPutFile)

	// act
	events := []*model.Event{}

	// generate requested number of events
	for i := 0; i < numEvents; i++ {
		events = append(events, generateEvent())
	}

	// marshall for saving
	content, err := json.Marshal(events)
	if err != nil {
		panic(fmt.Errorf("unable to marshall events : %+v", err))
	}

	// and write everything
	err = os.WriteFile(outPutFile, content, 0777)
	if err != nil {
		panic(fmt.Errorf("unable to write file : %+v", err))
	}
}

// generateEvent creates and returns a new instance of model.Event populated with random values for all its fields.
func generateEvent() *model.Event {
	return &model.Event{
		EventSource:     rand.Intn(88005553535),
		EventRef:        uuid.New().String(),
		EventType:       generateEventType(),
		EventDate:       *generator.RandomDate(),
		CallingNumber:   rand.Intn(88005553535),
		CalledNumber:    rand.Intn(88005553535),
		Location:        generator.RandomString(),
		DurationSeconds: rand.Intn(100),
		Attr1:           generator.RandomString(),
		Attr2:           generator.RandomString(),
		Attr3:           generator.RandomString(),
		Attr4:           generator.RandomString(),
		Attr5:           generator.RandomString(),
		Attr6:           generator.RandomString(),
		Attr7:           generator.RandomString(),
		Attr8:           generator.RandomString(),
	}
}

// generateEventType generates a random event type (integer) based on predefined probability distributions:
// - Type 1: 15%
// - Type 2: 20%
// - Type 3: 20%
// - Type 5: 45%
func generateEventType() int {
	r := rand.Intn(100)

	if r < 15 {
		return 1
	} else if r < 35 {
		return 2
	} else if r < 55 {
		return 3
	}
	return 5
}
