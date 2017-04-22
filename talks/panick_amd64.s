// see also: https://golang.org/doc/asm

// START OMIT
#include "go_asm.h" // マクロ定義
#include "go_tls.h" // get_tlsの定義
#include "textflag.h" // NOSPLIT等の定義

// func ptrPanic() (ret int64)
TEXT ·ptrPanic(SB),NOSPLIT,$0-8 // 関数宣言。int64=8 byteを返り値用にスタックにとる。中黒(·)に注意！ // HL
                                // Thread Local Storageにレジスタ(≒変数) CXに格納
	get_tls(CX)                 // CX := get_tls()  // HL
                                // マクロ `g` を使ってCXからstruct gを辿り、レジスタBXに格納
	MOVQ	g(CX), BX           // BX := CX.g // HL
                                // マクロ  `g__panic` を使ってフィールドpanicをレジスタCXに格納
	MOVQ	g__panic(BX), CX    // CX := BX.panic // HL
                                // CXを返り値変数retに格納
	MOVQ	CX, ret+0(FP)       // ret = CX // HL
	RET // return
// END OMIT