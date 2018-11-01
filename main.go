package main

import (
	"flag"

	"go.uber.org/zap"
)

func main() {
	var (
		attendeesFile = flag.String("attendees", "attendees.csv", "Input file containing lunch attendees")
		groupFile     = flag.String("groups", "groups.csv", "Output file containing recommended lunch groups")
		iterations    = flag.Int("iterations", 10000, "Number of iterations to simulate")
		historyFile   = flag.String("history", "history.", "Sqlite Database containing group history")
	)

	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	logger.Info("Reading input from file", zap.String("input", *attendeesFile))
	logger.Info("Will output groups to file", zap.String("output", *groupFile))
	logger.Info("Reading history from file", zap.String("history", *historyFile))
	logger.Infof("Will run %v iterations", iterations)
}
