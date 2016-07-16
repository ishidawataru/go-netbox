// Copyright 2016 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netbox

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClientDCIMGetRack(t *testing.T) {
	wantRack := testRack(1)

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/dcim/racks/1", wantRack))
	defer done()

	gotRack, err := c.DCIM.GetRack(wantRack.ID)
	if err != nil {
		t.Fatalf("unexpected error from Client.DCIM.GetRack: %v", err)
	}

	if want, got := *wantRack, *gotRack; !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpeted Rack:\n- want: %v\n- got: %v", want, got)
	}
}
