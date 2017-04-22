package panick_test

import (
	"testing"

	"github.com/t-yuki/panick"
)

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

func TestPanicked(t *testing.T) {
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
