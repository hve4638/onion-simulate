package onion

import "hve/onion-simulate/internal/types"

type Week uint

const (
	Monday Week = 1 << iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
	Everyday Week = Monday | Tuesday | Wednesday | Thursday | Friday | Saturday | Sunday
)

type Routines []Routine

type Routine struct {
	Name             string
	RepeatCount      types.IntRange
	RepeatInterval   types.IntRange
	Period           Period
	Endpoint         OnionNode
	CommunicateCount types.IntRange
}

type Period struct {
	week Week
	time types.Int64Range
}
