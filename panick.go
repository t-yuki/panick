package panick

//go:generate sh gen.sh $GOROOT

type Panic struct {
	p uintptr
}

func Observe() (*Panic, bool) {
	p := ptrPanic()
	if p == uintptr(0) {
		return nil, false
	}
	return &Panic{p: p}, true
}

func (p Panic) Recovered() bool {
	return ptrPanicRecovered(p.p)
}

func (p Panic) Aborted() bool {
	return ptrPanicAborted(p.p)
}

func (p Panic) Arg() interface{} {
	return ptrPanicArg(p.p)
}

func (p Panic) Link() *Panic {
	p2 := ptrPanicLink(p.p)
	if p2 != uintptr(0) {
		return &Panic{p: p2}
	}
	return nil
}

func Panicked() bool {
	p, _ := Observe()
	return p != nil && !p.Recovered() && !p.Aborted()
}

func ptrPanic() uintptr
func ptrPanicRecovered(p uintptr) bool
func ptrPanicAborted(p uintptr) bool
func ptrPanicLink(p uintptr) uintptr
func ptrPanicArg(p uintptr) interface{}
