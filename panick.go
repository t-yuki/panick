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
	recovered bool
	aborted   bool
	arg       interface{}
	link      *Panic
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
		return walk(t), true
	}
	return nil, false
}

func walk(t iface.Panic) *Panic {
	if t == nil {
		return nil
	}
	p := &Panic{
		aborted:   t.Aborted(),
		recovered: t.Recovered(),
		arg:       t.Arg(),
		link:      walk(t.Link()),
	}
	return p
}

func (p Panic) Recovered() bool {
	return p.recovered
}

func (p Panic) Aborted() bool {
	return p.aborted
}

func (p Panic) Arg() interface{} {
	return p.arg
}

func (p Panic) Link() *Panic {
	return p.link
}

func Panicked() bool {
	p, _ := Observe()
	return p != nil && !p.Recovered() && !p.Aborted()
}
