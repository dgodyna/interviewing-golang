package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dmgo1014/interviewing-golang.git/pkg/model"
	"github.com/xo/dburl"
	"io/ioutil"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// "postgresql://nrm:nrm@pg:5432/nrm?sslmode=disable"

// Loader will read generated dump and load it in provided DB.
//
// arg 1 is DB URL for database to load data
// atg 2 is path to file to load
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

	inputFile := os.Args[2]

	fmt.Printf("input file: %s\n", inputFile)

	dbUrl := os.Args[1]
	url, err := dburl.Parse(dbUrl)
	if err != nil {
		panic(fmt.Errorf("unable to parse database URL '%s' : %+v", url, err))
	}

	eventRaw, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(fmt.Errorf("unable to read input file : %+v", err))
	}

	var events []*model.Event

	err = json.Unmarshal(eventRaw, &events)
	if err != nil {
		panic(fmt.Errorf("unable to unmarshall event file content : %+v", err))
	}

	fmt.Printf("Total events to load : %d\n", len(events))

	db, err := sql.Open("postgres", url.DSN)
	if err != nil {
		panic(fmt.Errorf("unable to connecto to database : %+v", err))
	}

	tx, err := db.Begin()
	if err != nil {
		panic(fmt.Errorf("unable to start transaction : %+v", err))
	}
	defer db.Close()

	for _, e := range events {
		err = load(tx, e)
		if err != nil {
			tx.Rollback()
			panic(fmt.Errorf("unable to load event : %+v", err))
		}
	}

	fmt.Printf("sucessfully loaded %d events\n", len(events))

	tx.Commit()

}

// load will save event to database.
func load(tx *sql.Tx, event *model.Event) error {

	q := `
insert into event(event_source, event_ref, event_type, event_date, calling_number, called_number, location,
                  duration_seconds, attr_1, attr_2, attr_3, attr_4, attr_5, attr_6, attr_7, attr_8)
values ($1, $2, $3, %s, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
`

	// we have to format query to use function for converting time
	q = fmt.Sprintf(q, timeToTimestampNoTz(&event.EventDate))

	_, err := tx.Exec(q,
		event.EventSource,
		event.EventRef,
		event.EventType,
		event.CallingNumber,
		event.CalledNumber,
		event.Location,
		event.DurationSeconds,
		event.Attr1,
		event.Attr2,
		event.Attr3,
		event.Attr4,
		event.Attr5,
		event.Attr6,
		event.Attr7,
		event.Attr8,
	)

	return err
}

// timeToTimestampNoTz will format go time to timestamp - thus will allow us to use epoch time
// and don't rely on client and server timezones.
func timeToTimestampNoTz(t *time.Time) string {
	return fmt.Sprintf("to_timestamp(cast(%d as bigint))::date", t.Unix())
}
