package main

import (
	"hve/onion-simulate/internal/onion"
	"hve/onion-simulate/internal/onion_ready"
)

func main() {
	options := onion_ready.ParseOption()
	config, err := onion_ready.ReadConfig(options.ConfigPath)
	if err != nil {
		panic("Fail to read config file\n")
	}

	// b, _ := json.MarshalIndent(config, "", "  ")
	// fmt.Println(string(b))

	// os.Exit(0)
	onion.Simulate(*config, options)
}
