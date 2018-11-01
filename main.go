package main

import (
	"flag"
	"fmt"
	"os"

	"encoding/csv"

	"go.uber.org/zap"
)

func main() {
	var (
		attendeesFile = flag.String("attendees", "attendees.csv", "Input file containing lunch attendees")
		groupFile     = flag.String("groups", "groups.csv", "Output file containing recommended lunch groups")
		iterations    = flag.Int("iterations", 10000, "Number of iterations to simulate")
		historyFile   = flag.String("history", "history.db", "Sqlite Database containing group history")
	)

	flag.Parse()

	logger := zap.NewExample()
	defer logger.Sync()

	logger.Info("Reading input from file", zap.String("input", *attendeesFile))
	logger.Info("Will output groups to file", zap.String("output", *groupFile))
	logger.Info("Reading history from file", zap.String("history", *historyFile))
	logger.Info(fmt.Sprintf("Will run %v iterations", iterations))

	f, err := os.Open(*attendeesFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		panic(err)
	}

	attendees := []string{}
	for _, line := range lines {
		attendees = append(attendees, line[0])
	}

	logger.Info(fmt.Sprintf("Found %d lunch attendees", len(attendees)))

}
