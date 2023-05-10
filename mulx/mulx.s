#include "textflag.h"

// func mulx(a, b uint64) (hi, lo uint64)
TEXT ·mulx(SB),NOSPLIT,$0-32
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), DX
    MULXQ AX, SI, DI
    MOVQ DI, hi+16(FP)
    MOVQ SI, lo+24(FP)
    RET

