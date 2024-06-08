#include "textflag.h"

// func rdtsc() (n uint64)
TEXT ·rdtsc(SB),NOSPLIT,$0
    BYTE	$0x0F // RDTSCP
	BYTE	$0x01
	BYTE	$0xF9
    SHLQ  $32, DX
    ADDQ  DX, AX
    MOVQ  AX, n+0(FP)
    RET
