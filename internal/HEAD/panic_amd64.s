#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"

// see also: https://golang.org/doc/asm

// func getPanic() uintptr
TEXT ·getPanic(SB),NOSPLIT,$0-8
	get_tls(CX)
	MOVQ	g(CX), BX
	MOVQ	g__panic(BX), CX
	MOVQ	CX, ret+0(FP)
	RET

// func panicRecovered(uintptr) bool
TEXT ·panicRecovered(SB),NOSPLIT,$0-9
	MOVQ 	p+0(FP), CX
	MOVQ	_panic_recovered(CX), BX
	MOVB	BX, ret+8(FP)
	RET

// func panicAborted(uintptr) bool
TEXT ·panicAborted(SB),NOSPLIT,$0-9
	MOVQ 	p+0(FP), CX
	MOVQ	_panic_aborted(CX), BX
	MOVB	BX, ret+8(FP)
	RET

// func panicLink(uintptr) uintptr
TEXT ·panicLink(SB),NOSPLIT,$0-16
	MOVQ 	p+0(FP), CX
	MOVQ	_panic_link(CX), BX
	MOVQ	BX, ret+8(FP)
	RET

// func panicArg(uintptr) interface{}
TEXT ·panicArg(SB),NOSPLIT,$0-24
	MOVQ 	p+0(FP), CX
	MOVQ	_panic_arg(CX), BX
	MOVQ	BX, ret_type+8(FP)
	ADDQ	$iface_data, CX
	MOVQ	_panic_arg(CX), BX
	MOVQ	BX, ret_data+16(FP)
	RET
