package sequence

type SeqOpt func(*Sequence)

func WithSize(sz uint64) SeqOpt {
	return func(s *Sequence) {
		s.sz = sz
	}
}

func WithAllocate(v bool) SeqOpt {
	return func(s *Sequence) {
		if v {
			s.raw = make([]byte, s.sz)
		}
	}
}
