package impl

import "github.com/t-yuki/panick/internal"

func init() {
	iface.GetPanic["go1.6.4"] = GetPanic
}

func GetPanic() iface.Panic {
	if p := getPanic(); p != uintptr(0) {
		return &Panic{p: p}
	}
	return nil
}

type Panic struct {
	p uintptr
}

func (p Panic) Recovered() bool {
	return panicRecovered(p.p)
}

func (p Panic) Aborted() bool {
	return panicAborted(p.p)
}

func (p Panic) Arg() interface{} {
	return panicArg(p.p)
}

func (p Panic) Link() iface.Panic {
	p2 := panicLink(p.p)
	if p2 != uintptr(0) {
		return &Panic{p: p2}
	}
	return nil
}

func getPanic() uintptr
func panicRecovered(p uintptr) bool
func panicAborted(p uintptr) bool
func panicLink(p uintptr) uintptr
func panicArg(p uintptr) interface{}
