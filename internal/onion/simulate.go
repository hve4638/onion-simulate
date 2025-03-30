package onion

import (
	"fmt"
	onion_log "hve/onion-simulate/internal/onion_log"
	"hve/onion-simulate/internal/types"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Simulate(config types.Config, option types.Options) {
	onionNetwork := NewOnionNetwork(uint32(config.NodeCount.Amount), option.Seed)
	fmt.Println("Simulation started with options:", option)

	onionNetwork.Init(&config)

	onionNetwork.ShowOption()
	onionNetwork.Simulate()

	go writeLog("log.txt", &onionNetwork.logs)
	fmt.Println("Simulation started")

	go waitForExit()

	// go func() {
	// 	for {
	// 		time.Sleep(1 * time.Second)
	// 		fmt.Println(onionNetwork.timer.Now())
	// 	}
	// }()
	// fmt.Scanln()
	// onionNetwork.Users[0].makeCircuit(4)
	// fmt.Scanln()

	// go test(&onionNetwork)

	select {}
}

func writeLog(filename string, logs *onion_log.ConcurrentLog) {
	for {
		time.Sleep(1 * time.Second)
		entries := logs.MergeAndClear()

		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		defer file.Close()

		for _, entry := range entries {
			file.WriteString(fmt.Sprintf("%s\n", entry.String()))
		}
	}
}

func waitForExit() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	fmt.Println("\nReceived interrupt signal. Shutting down...")
	fmt.Println("Simulation finished")

	os.Exit(0)
}
