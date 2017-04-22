package panick_test

import (
	"runtime"
	"sync"
	"testing"

	"github.com/t-yuki/panick"
)

func TestPanic_Link(t *testing.T) {
	defer func() {
		p, _ := panick.Observe()
		p2 := p.Link()
		if !p2.Recovered() {
			t.Fatal("test1 panic should be recovered")
		}
		a, ok := p2.Arg().(string)
		if !ok || a != "test1" {
			t.Fatal("Arg should catch the arg")
		}
		recover()
	}()
	defer func() {
		recover()
		panic("test2")
	}()
	panic("test1")
}

func TestPanicked_nil(t *testing.T) {
	defer func() {
		if !panick.Panicked() {
			t.Fatal("Panicked should detect nil panic")
		}
		if e := recover(); e != nil {
			t.Fatal("panic should be nil")
		}
	}()
	panic(nil)
}

func TestPaniced(t *testing.T) {
	ok := false
	if panick.Panicked() {
		t.Fatal("it shouldn't panicked")
	}
	defer func() {
		if !ok {
			t.Fatal("Panicked should detect panic")
		}

		if e := recover(); e == nil {
			t.Fatal("Panicked shouldn't recover")
		}
		if panick.Panicked() {
			t.Fatal("Panicked shouldn't detect recovered panic")
		}
	}()
	func() {
		defer func() {
			if panick.Panicked() {
				ok = true
			}
		}()
		panic("panic")
	}()
	t.Fatal("Panicked shouldn't recover")
}

func TestPanic_Aborted(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if p, _ := panick.Observe(); !p.Aborted() {
				t.Fatal("Goexit should aborts panic")
			}
			if e := recover(); e != nil {
				t.Fatal("Goexit should hide panic")
			}
		}()
		defer runtime.Goexit()
		defer func() {
			p, _ := panick.Observe()
			if p.Aborted() {
				t.Fatal("panic shouldn't be aborted")
			}
		}()
		panic("test")
	}()
	wg.Wait()
}
