// generated by compile -asmhdr from package runtime

#define g__size 400
#define g_stack 0
#define g_stackguard0 16
#define g_stackguard1 24
#define g__panic 32 // _panicフィールドの、g structの先頭からのオフセット(バイト数) // HL
#define g__defer 40
#define g_m 48

#define _panic__size 40
#define _panic_argp 0
#define _panic_arg 8 // argフィールドの、_panic structの先頭からのオフセット。interface{}なので16バイト // HL
#define _panic_link 24
#define _panic_recovered 32
#define _panic_aborted 33
