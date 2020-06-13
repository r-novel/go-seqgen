package sequence

import (
	"fmt"
	"math/rand"
	"time"
)

type Sequence struct {
	sz  uint64
	raw []byte
	src rand.Source
}

func NewSequnce(opts ...SeqOpt) string {
	const (
		defaultSize = 64
		defaultSrc  = rand.NewSource(time.Now().UnixNano())
	)

	s := &Sequence{
		sz:  defaultSize,
		src: defaultSrc,
	}

	for _, opt := range opts {
		opt(h)
	}

	return s
}

func (s *Sequence) Generate() error {
	if s.src == nil || s.sz == 0 {
		return fmt.Errorf("error with allocate data;")
	}

	for i, cache, remain := s.sz-1,
		s.src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = s.src.Int63(), letterIdxMax
		}

		if idx := uint(cache & letterIdxMask); idx < len(letterBytes) {
			s.raw[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return nil
}
