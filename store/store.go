package store

import (
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/swissinfo-ch/logd/ring"
)

type Store struct {
	rings     map[string]*ring.Ring
	fallback  *ring.Ring
	numWrites atomic.Uint64
}

type Cfg struct {
	RingSizes    map[string]uint32
	FallbackSize uint32
}

func NewStore(cfg *Cfg) *Store {
	s := &Store{
		rings:    make(map[string]*ring.Ring, len(cfg.RingSizes)),
		fallback: ring.NewRing(cfg.FallbackSize),
	}
	for key, size := range cfg.RingSizes {
		s.rings[key] = ring.NewRing(size)
	}
	return s
}

// Write writes to the ring of key, or fallback ring
func (s *Store) Write(key string, data []byte) {
	s.numWrites.Add(uint64(1))
	part := s.rings[key]
	if part == nil {
		s.fallback.Write(data)
		return
	}
	part.Write(data)
}

func (s *Store) Heads() map[string]uint32 {
	heads := make(map[string]uint32, len(s.rings)+1)
	for key, ring := range s.rings {
		heads[key] = ring.Head()
	}
	heads["_fallback"] = s.fallback.Head()
	return heads
}

func (s *Store) Sizes() map[string]uint32 {
	sizes := make(map[string]uint32, len(s.rings)+1)
	for key, ring := range s.rings {
		sizes[key] = ring.Size()
	}
	sizes["_fallback"] = s.fallback.Size()
	return sizes
}

// Read reads up to limit items, from offset,
// all rings with the given key prefix
func (s *Store) Read(keyPrefix string, offset, limit uint32) <-chan []byte {
	fmt.Println("store.Read()", keyPrefix, offset, limit)
	out := make(chan []byte)
	go func() {
		defer close(out)
		// try to read from exact ring
		exactRing := s.rings[keyPrefix]
		if exactRing != nil {
			fmt.Println("exact match", keyPrefix)
			for d := range exactRing.Read(offset, limit) {
				out <- d
			}
			return
		}
		var count uint32
		// ranging through rings for prefix
		for key, r := range s.rings {
			if strings.HasPrefix(key, keyPrefix) {
				fmt.Println(keyPrefix, "matched", key)
				for d := range r.Read(offset, limit-count) {
					out <- d
					count++
					if count >= limit {
						return
					}
				}
			}
		}
		// fallback
		for d := range s.fallback.Read(offset, limit-count) {
			out <- d
			count++
			if count >= limit {
				return
			}
		}
	}()
	return out
}

func (s *Store) NumWrites() uint64 {
	return s.numWrites.Load()
}