//go:generate sh ../gen.sh $GOROOT

package panick

// PanickedはPanic中だったらtrueを返す
func Panicked() bool {
	return ptrPanic() != uintptr(0)
}

// ptrPanicは_panicへのポインタを返す.
// 中身は別ファイルで定義されている
func ptrPanic() uintptr
