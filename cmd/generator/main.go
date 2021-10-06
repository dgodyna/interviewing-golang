package main

import (
	"encoding/json"
	"fmt"
	"github.com/dmgo1014/interviewing-golang.git/pkg/generator"
	"github.com/dmgo1014/interviewing-golang.git/pkg/model"
	"github.com/google/uuid"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// run generation of costed eventsand save them to provided file.
// costed event will have following types and probability:
// * type 1 - 15%
// * type 2 - 20%
// * type 3 - 20&
// * type 5 - 45%
//
// rest of fields will be filled randomly.
//
// arg 1 - number of events to generate
// arg 2 - output file.
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

	numEventsStr := os.Args[1]
	numEvents, err := strconv.Atoi(numEventsStr)
	if err != nil {
		panic(fmt.Errorf("unable to parse number of events : %w", err))
	}

	outPutFile := os.Args[2]

	fmt.Printf("number event : %d\n", numEvents)
	fmt.Printf("dump output: %s\n", outPutFile)

	// act
	var events []*model.Event

	// generate requested number of events
	for i := 0; i < numEvents; i++ {
		events = append(events, generateEvent())
	}

	// marshall for saving
	content, err := json.Marshal(events)
	if err != nil {
		panic(fmt.Errorf("unable to marshall events : %w", err))
	}

	// and write everything
	err = ioutil.WriteFile(outPutFile, content, 0777)
	if err != nil {
		panic(fmt.Errorf("unable to write file : %w", err))
	}
}

// generateEvent will create a new instance of event with some random values.
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

// generateEventType will generate event type with following probability;
// * type 1 - 15%
// * type 2 - 20%
// * type 3 - 20&
// * type 5 - 45%
func generateEventType() int {
	r := rand.Intn(100)
	switch {
	case r < 15:
		return 1
	case r < 35:
		return 2
	case r < 55:
		return 3
	default:
		return 5
	}
}
