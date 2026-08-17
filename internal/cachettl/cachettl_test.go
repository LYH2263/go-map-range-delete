package cachettl

import (
	"testing"
	"time"
)

func TestPurgeExpired(t *testing.T) {
	now := time.Unix(100, 0)
	s := &Store{expire: map[string]time.Time{
		"old": now.Add(-time.Second),
		"new": now.Add(time.Hour),
	}}
	s.PurgeExpired(now)
	if _, ok := s.expire["old"]; ok {
		t.Fatal("old should be purged")
	}
	if _, ok := s.expire["new"]; !ok {
		t.Fatal("new should remain")
	}
}
