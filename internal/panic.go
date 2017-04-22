package iface

var GetPanic = map[string]func() Panic{}

type Panic interface {
	Recovered() bool
	Aborted() bool
	Arg() interface{}
	Link() Panic
}
