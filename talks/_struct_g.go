// +build ignore

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// START OMIT
package runtime

type g struct {
	stack       stack   // offset known to runtime/cgo
	stackguard0 uintptr // offset known to liblink
	stackguard1 uintptr // offset known to liblink

	_panic *_panic // innermost panic - offset known to liblink // HL
	_defer *_defer // innermost defer
	m      *m      // current m; offset known to arm liblink
	// ...
	goid int64 // HL
	// END OMIT
	waitsince      int64  // approx time when the g become blocked
	waitreason     string // if status==Gwaiting
	schedlink      guintptr
	preempt        bool     // preemption signal, duplicates stackguard0 = stackpreempt
	paniconfault   bool     // panic (instead of crash) on unexpected fault address
	preemptscan    bool     // preempted g does scan for gc
	gcscandone     bool     // g has scanned stack; protected by _Gscan bit in status
	gcscanvalid    bool     // false at start of gc cycle, true if G has not run since last scan; transition from true to false by calling queueRescan and false to true by calling dequeueRescan
	throwsplit     bool     // must not split stack
	raceignore     int8     // ignore race detection events
	sysblocktraced bool     // StartTrace has emitted EvGoInSyscall about this goroutine
	sysexitticks   int64    // cputicks when syscall has returned (for tracing)
	traceseq       uint64   // trace event sequencer
	tracelastp     puintptr // last P emitted an event for this goroutine
	lockedm        *m
	sig            uint32
	writebuf       []byte
	sigcode0       uintptr
	sigcode1       uintptr
	sigpc          uintptr
	gopc           uintptr // pc of go statement that created this goroutine
	startpc        uintptr // pc of goroutine function
	racectx        uintptr
	waiting        *sudog    // sudog structures this g is waiting on (that have a valid elem ptr); in lock order
	cgoCtxt        []uintptr // cgo traceback context
}
