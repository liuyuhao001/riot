// Copyright 2016 ego authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package storage

import (
	"os"
	"testing"

	"github.com/vcaesar/tt"
)

func TestOpenOrCreateBolt(t *testing.T) {
	db, err := OpenBolt("bolt_test")
	tt.Expect(t, "<nil>", err)
	db.Close()

	db, err = OpenBolt("bolt_test")
	tt.Expect(t, "<nil>", err)
	err = db.Set([]byte("key1"), []byte("value1"))
	tt.Expect(t, "<nil>", err)

	has, err := db.Has([]byte("key1"))
	tt.Equal(t, nil, err)
	if err == nil {
		tt.Equal(t, true, has)
	}

	buffer := make([]byte, 100)
	buffer, err = db.Get([]byte("key1"))
	tt.Expect(t, "<nil>", err)
	tt.Expect(t, "value1", string(buffer))

	walFile := db.WALName()
	db.Close()
	os.Remove(walFile)
	os.Remove("bolt_test")
}
