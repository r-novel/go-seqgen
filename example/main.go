package main

import (
	"fmt"

	"github.com/r-novel/go-seqgen/internal/sequence"
)

func main() {
	sq := sequence.NewSequnce(sequence.WithSize(128), sequence.WithAllocate(true))
	if sq != nil {
		if e := sq.Generate(); e == nil {
			fmt.Printf("sequence: %s\n", sq)
		}
	}
}
