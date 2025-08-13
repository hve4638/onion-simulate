package types

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
	RepeatCount      int64
	RepeatInterval   IntRange
	Period           Period
	Endpoint         OnionNode
	CommunicateCount IntRange
}

type Period struct {
	week Week
	time Int64Range
}
