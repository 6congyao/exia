/*
 * Copyright (c) 2019. LuCongyao <6congyao@gmail.com> .
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this work except in compliance with the License.
 * You may obtain a copy of the License in the LICENSE file, or at:
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package things

import "strings"

// Thing represents a thing. Each thing is owned by one user, and
// it is assigned with the unique identifier and (temporary) access key.
type Thing struct {
	ID       string
	Owner    string
	Type     string
	Name     string
	Key      string
	Metadata string
}

// ThingsPage contains page related metadata as well as list of things that
// belong to this page.
type ThingsPage struct {
	PageMetadata
	Things []Thing
}

var thingTypes = map[string]bool{
	"app":    true,
	"device": true,
}

// Validate returns an error if thing representation is invalid.
func (c *Thing) Validate() error {
	if c.Type = strings.ToLower(c.Type); !thingTypes[c.Type] {
		return ErrMalformedEntity
	}

	return nil
}
