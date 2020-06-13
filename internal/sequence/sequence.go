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

func NewSequnce(opts ...SeqOpt) *Sequence {
	const defaultSize = 64
	var defaultSrc = rand.NewSource(time.Now().UnixNano())

	s := &Sequence{
		sz:  defaultSize,
		src: defaultSrc,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Sequence) String() string {
	return string(s.raw)
}

func (s *Sequence) Generate() error {
	if s.src == nil || s.sz == 0 {
		return fmt.Errorf("error with allocate data;")
	}

	for i, cache, remain := s.sz-1,
		s.src.Int63(), letterIdxMax; i > 0; {
		if remain == 0 {
			cache, remain = s.src.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(letters) {
			s.raw[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return nil
}
