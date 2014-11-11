// Copyright 2014 Richard Lehane. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Core siegfried defaults
package config

import (
	"net/http"
	"path/filepath"
	"time"
)

var siegfried = struct {
	version          [3]int // Siegfried version (i.e. of the sf tool)
	home             string // Home directory used by both sf and r2d2 tools
	signature        string // Name of signature file
	signatureVersion int    // Version of the signature file (this is used for the update service)
	// Defaults for processing bytematcher signatures. These control the segmentation.
	distance  int // The acceptable distance between two frames before they will be segmented (default is 8192)
	rng       int // The acceptable range between two frames before they will be segmented (default is 0-2049)
	choices   int // The acceptable number of plain sequences generated from a single segment
	varLength int // The acceptable length of a variable byte sequence (longer the better to reduce false matches)
	// Config for using the update service.
	updateURL       string // URL for the update service (a JSON file that indicates whether update necessary and where can be found)
	updateTimeout   time.Duration
	updateTransport *http.Transport
}{
	version:          [3]int{0, 6, 0},
	signature:        "pronom.gob",
	signatureVersion: 5,
	distance:         8192,
	rng:              2049,
	choices:          64,
	varLength:        1,
	updateURL:        "http://www.itforarchivists.com/siegfried/update",
	updateTimeout:    30 * time.Second,
	updateTransport:  &http.Transport{Proxy: http.ProxyFromEnvironment},
}

// GETTERS

func Version() [3]int {
	return siegfried.version
}

func Home() string {
	return siegfried.home
}

func Signature() string {
	if filepath.Dir(siegfried.signature) == "." {
		return filepath.Join(siegfried.home, siegfried.signature)
	}
	return siegfried.signature
}

func SignatureBase() string {
	return siegfried.signature
}

func SignatureVersion() int {
	return siegfried.signatureVersion
}

func Distance() int {
	return siegfried.distance
}

func Range() int {
	return siegfried.rng
}

func Choices() int {
	return siegfried.choices
}

func VarLength() int {
	return siegfried.varLength
}

func BMOptions() (int, int, int, int) {
	return siegfried.distance, siegfried.rng, siegfried.choices, siegfried.varLength
}

func UpdateOptions() (string, time.Duration, *http.Transport) {
	return siegfried.updateURL, siegfried.updateTimeout, siegfried.updateTransport
}

// SETTERS

func SetHome(h string) func() private {
	return func() private {
		siegfried.home = h
		return private{}
	}
}

func SetSignature(s string) func() private {
	return func() private {
		siegfried.signature = s
		return private{}
	}
}

func SetDistance(i int) func() private {
	return func() private {
		siegfried.distance = i
		return private{}
	}
}

func SetRange(i int) func() private {
	return func() private {
		siegfried.rng = i
		return private{}
	}
}

func SetChoices(i int) func() private {
	return func() private {
		siegfried.choices = i
		return private{}
	}
}

func SetVarLength(i int) func() private {
	return func() private {
		siegfried.varLength = i
		return private{}
	}
}