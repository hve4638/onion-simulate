package onion

// import (
// 	"fmt"
// 	"hve/onion-simulate/internal/types"
// 	"math/rand"
// 	"slices"
// 	"time"
// )

// func (op *UserOP) Simulate() {
// 	for _, routine := range op.routines {
// 		go op.simulateRoutine(&routine)
// 	}

// 	for {
// 		packet := <-op.packetCh

// 		ch, ok := op.circuitChs[packet.circId]
// 		if ok {
// 			*ch <- packet
// 		} else {
// 			fmt.Println("Circuit channel not found")
// 			continue
// 		}
// 	}
// }

// func (op *UserOP) simulateRoutine(rt *Routine) {
// 	for {
// 		now := op.network.timer.Now() / 1000
// 		if !inPeriod(rt.Period, now) {
// 			time.Sleep(timeUntilNextPeriod(rt.Period, now))
// 			continue
// 		}
// 		// fmt.Println("Routine", rt.Name, "is not running")

// 		repeat := int64(0)
// 		for repeat < rt.RepeatCount {
// 			// fmt.Println("Routine", rt.Name, "is running")
// 			now := op.network.timer.Now()
// 			if !inPeriod(rt.Period, now) {
// 				break
// 			}

// 			op.runRoutine(rt)

// 			interval := randInRange(rt.RepeatInterval)
// 			time.Sleep(time.Duration(interval) * time.Millisecond)

// 			repeat++
// 		}
// 	}
// }

// func inPeriod(period Period, now int64) bool {
// 	return period.time.Min <= now && now <= period.time.Max
// }

// func timeUntilNextPeriod(p Period, now int64) time.Duration {
// 	return 1 * time.Second
// }

// func randInRange(r types.IntRange) int {
// 	if r.Max <= r.Min {
// 		return r.Min
// 	}
// 	return rand.Intn(r.Max-r.Min+1) + r.Min
// }

// func (op *UserOP) runRoutine(routine *Routine) {
// 	op.makeCircuit(3)
// }

// func (op *UserOP) makeCircuit(depth int) {
// 	if depth <= 0 || depth > len(op.network.Routers) {
// 		return
// 	}
// 	selectedRouters := op.getRouters(depth)

// 	circId := op.nextCircuitId
// 	op.nextCircuitId++

// 	ch := make(chan Packet)
// 	op.circuitChs[circId] = &ch
// 	for i, node := range selectedRouters {
// 		var packet Packet
// 		if i == 0 {
// 			packet = op.MakeCreatePacket(circId, node)
// 		} else {
// 			packet = op.MakeExtendPacket(circId, node)
// 		}
// 		packet.Send()
// 		<-ch
// 	}
// }

// func (op *UserOP) getRouters(depth int) []*OnionRouter {
// 	selectedRouters := make([]*OnionRouter, 0, depth)

// 	for len(selectedRouters) < depth {
// 		nodeIndex := int(op.rand.Int31n(int32(len(op.network.Routers))))
// 		router := op.network.Routers[nodeIndex]

// 		duplicate := slices.Contains(selectedRouters, router)
// 		if !duplicate {
// 			selectedRouters = append(selectedRouters, router)
// 		}
// 	}
// 	return selectedRouters
// }
