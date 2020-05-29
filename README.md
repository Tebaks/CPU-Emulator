# CPU-Emulator
Simple CPU Emulator using golang

## Operations
- START : Start of the program.
- LOAD "number" : Load to accumulator.
- STORE "memory" : Store from accumulator to memory.
- CMPM "memory" : Compare accumulator and memory if memory greater set flag to -1, if accumulator greater set flag to 1 else set flag to 0.
- CJMP "index" : If flag is 1 jump to the index.
- JUMP "index" : Jump to the index.
- LOADM "memory": Load accumulator from meory.
- DISP: print accumulator to screen.
- ADD "number" : Sum number and accumulator and store to accumulator.
- ADDM "memory": Sum accumulator and memory and store to accumulator.
- SUB "number" : Subtract number from accumulator and store to accumulator.
- SUBM "memory": Subtract memory from accumulator and store to accumulator.
- MUL "number" : Multiply number and accumulator and store to accumulator.
- MULM "memory": Multiply accumulator and memory and store to accumulator.
- HALT: End of the program.

## Test
Test app that computes the sum of the numbers between 0 to 20