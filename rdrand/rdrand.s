#include "textflag.h"

// func rdrand64() (ok bool, ret uint64)
TEXT ·rdrand64(SB),NOSPLIT,$0
    // RDRAND %RAX
    BYTE $0x48; BYTE $0x0F; BYTE $0xC7;  BYTE $0xF0
    MOVB $0, DX
    MOVB $1, BX
    // CMOVB BX, DX
    BYTE $0x0F; BYTE $0x42; BYTE $0xD3
    MOVB DX, ok+0(FP)
    MOVQ AX, ret+8(FP)
    RET

// func rdseed64() (ok bool, ret uint64)
TEXT ·rdseed64(SB),NOSPLIT,$0
    // RDSEED %RAX
    BYTE $0x48; BYTE $0x0F; BYTE $0xC7;  BYTE $0xF8
    MOVB $0, DX
    MOVB $1, BX
    // CMOVB BX, DX
    BYTE $0x0F; BYTE $0x42; BYTE $0xD3
    MOVB DX, ok+0(FP)
    MOVQ AX, ret+8(FP)
    RET

