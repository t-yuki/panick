package panick

import (
	"runtime"

	"github.com/t-yuki/panick/internal"
	_ "github.com/t-yuki/panick/internal/go1.7.5"
	_ "github.com/t-yuki/panick/internal/go1.8"
	_ "github.com/t-yuki/panick/internal/go1.8.1"
)

//go:generate sh gen.sh $GOROOT

type Panic struct {
	t iface.Panic
}

func Observe() (*Panic, bool) {
	get, ok := iface.GetPanic[runtime.Version()]
	if !ok {
		get, ok = iface.GetPanic["HEAD"]
		if !ok {
			panic("unsupported runtime:" + runtime.Version())
		}
	}

	t := get()
	if t != nil {
		return &Panic{t: t}, true
	}
	return nil, false
}

func (p Panic) Recovered() bool {
	return p.t.Recovered()
}

func (p Panic) Aborted() bool {
	return p.t.Aborted()
}

func (p Panic) Arg() interface{} {
	return p.t.Arg()
}

func (p Panic) Link() *Panic {
	t := p.t.Link()
	if t != nil {
		return &Panic{t: t}
	}
	return nil
}

func Panicked() bool {
	p, _ := Observe()
	return p != nil && !p.Recovered() && !p.Aborted()
}
