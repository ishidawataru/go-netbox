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

type RackGroupIdentifier struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type RackGroup struct {
	ID   int            `json:"id"`
	Name string         `json:"name"`
	Slug string         `json:"slug"`
	Site SiteIdentifier `json:"site"`
}

type Rack struct {
	ID          int                  `json:"id"`
	Name        string               `json:"name"`
	DisplayName string               `json:"display_name"`
	FacilityID  string               `json:"facility_id"`
	Site        *SiteIdentifier      `json:"site"`
	Group       *RackGroupIdentifier `json:"group"`
	UHeight     int                  `json:"u_height"`
	Comments    string               `json:"comments"`
}

type RackIdentifier struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	FacilityID  string `json:facility_id"`
}

// GetRack retrieves a Rack object from NetBox by its ID.
func (s *DCIMService) GetRack(id int) (*Rack, error) {
	r := new(Rack)
	if err := s.get("racks", id, r); err != nil {
		return nil, err
	}
	return r, nil
}

// ListRacks retrives a list of Rack objects from NetBox.
func (s *DCIMService) ListRacks() ([]*Rack, error) {
	var rs []*Rack
	if err := s.list("racks", &rs); err != nil {
		return nil, err
	}
	return rs, nil
}
