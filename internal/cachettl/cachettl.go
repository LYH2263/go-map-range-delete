package cachettl

import "time"

type Store struct {
	expire map[string]time.Time
}

func (s *Store) PurgeExpired(now time.Time) {
	for k, exp := range s.expire {
		if !exp.After(now) {
			delete(s.expire, k)
		}
	}
}
