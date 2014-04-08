// Package core defines the Siegfried struct and Identifier/Identification interfaces.
package core

import (
	"io"
	"sync"

	"github.com/richardlehane/siegfried/pkg/core/siegreader"
)

type Siegfried struct {
	identifiers []Identifier // at present only one identifier (the PRONOM identifier) is used, but can add other identifiers e.g. for FILE sigs
	buffer      *siegreader.Buffer
	name        string
}

type Identifier interface {
	Identify(*siegreader.Buffer, string, chan Identification, *sync.WaitGroup)
}

type Identification interface {
	String() string
	Confidence() float64 // how certain is this identification?
	Basis() string       // on what grounds was this identification made?
}

func NewSiegfried() *Siegfried {
	s := new(Siegfried)
	s.identifiers = make([]Identifier, 0, 1)
	s.buffer = siegreader.New()
	return s
}

func (s *Siegfried) AddIdentifier(i Identifier) {
	s.identifiers = append(s.identifiers, i)
}

// Identify applies the set of identifiers to a reader and file name. If the file name is not known, use an empty string instead.
func (s *Siegfried) Identify(r io.Reader, n string) (chan Identification, error) {
	err := s.buffer.SetSource(r)
	if err != nil && err != io.EOF {
		return nil, err
	}
	s.name = n
	ret := make(chan Identification)
	go s.identify(ret)
	return ret, nil
}

func (s *Siegfried) identify(ret chan Identification) {
	var wg sync.WaitGroup
	for _, v := range s.identifiers {
		wg.Add(1)
		go v.Identify(s.buffer, s.name, ret, &wg)
	}
	wg.Wait()
	close(ret)
}
