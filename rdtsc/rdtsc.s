#include "textflag.h"

// func rdtsc() (n uint64)
TEXT ·rdtsc(SB),NOSPLIT,$0
    RDTSCP
    SHLQ  $32, DX
    ADDQ  DX, AX
    MOVQ  AX, n+0(FP)
    RET
