#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"

// see also: https://golang.org/doc/asm

// func ptrPanic() int64
TEXT ·ptrPanic(SB),NOSPLIT,$0-8
	get_tls(CX)
	MOVQ	g(CX), BX
	MOVQ	g__panic(BX), CX
	MOVQ	CX, ret+0(FP)
	RET

// func ptrPanicRecovered(uintptr) bool
TEXT ·ptrPanicRecovered(SB),NOSPLIT,$0-9
	MOVQ 	p+0(FP), CX
	MOVQ	_panic_recovered(CX), BX
	MOVB	BX, ret+8(FP)
	RET

// func ptrPanicAborted(uintptr) bool
TEXT ·ptrPanicAborted(SB),NOSPLIT,$0-9
	MOVQ 	p+0(FP), CX
	MOVQ	_panic_aborted(CX), BX
	MOVB	BX, ret+8(FP)
	RET

// func ptrPanicLink(uintptr) uintptr
TEXT ·ptrPanicLink(SB),NOSPLIT,$0-16
	MOVQ 	p+0(FP), CX
	MOVQ	_panic_link(CX), BX
	MOVQ	BX, ret+8(FP)
	RET
