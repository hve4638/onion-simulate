package onion_ready

import (
	"flag"
	"hve/onion-simulate/internal/types"
)

func ParseOption() types.Options {
	var options types.Options

	seedOpt := flag.Int64("seed", 0, "")
	debugOpt := flag.Bool("debug", false, "")
	configPathOpt := flag.String("config", "", "")
	logOutputPathOpt := flag.String("output", "", "")

	flag.Parse()
	options.Seed = *seedOpt
	options.Debug = *debugOpt
	options.ConfigPath = *configPathOpt
	options.LogOutputPath = *logOutputPathOpt

	return options
}
