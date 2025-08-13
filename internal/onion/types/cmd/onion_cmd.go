package cmd

import (
	"fmt"
)

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

func (cmd OnionCMD) String() string {
	switch cmd {
	case Relay:
		return "Relay"
	case Data:
		return "Data"
	case Begin:
		return "Begin"
	case End:
		return "End"
	case Teardown:
		return "Teardown"
	case Connected:
		return "Connected"
	case Create:
		return "Create"
	case Created:
		return "Created"
	case Extend:
		return "Extend"
	case Extended:
		return "Extended"
	case Truncate:
		return "Truncate"
	case Truncated:
		return "Truncated"
	case Sendme:
		return "Sendme"
	case Drop:
		return "Drop"
	default:
		return fmt.Sprintf("UnknownOnionCMD(%d)", int(cmd))
	}
}
