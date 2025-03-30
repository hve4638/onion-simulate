package onion

type OnionCMD int

const (
	Relay OnionCMD = iota
	Data
	Begin
	End
	Teardown
	Connected
	Create
	Created
	Extend
	Extended
	Truncate
	Truncated
	Sendme
	Drop
)
