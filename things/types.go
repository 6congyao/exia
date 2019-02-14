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

// ThingRepository specifies a thing persistence API.
type ThingRepository interface {
	// Save persists the thing. Successful operation is indicated by non-nil
	// error response.
	Save(Thing) (string, error)

	// Update performs an update to the existing thing. A non-nil error is
	// returned to indicate operation failure.
	Update(Thing) error

	// RetrieveByID retrieves the thing having the provided identifier, that is owned
	// by the specified user.
	RetrieveByID(string, string) (Thing, error)

	// RetrieveByKey returns thing ID for given thing key.
	RetrieveByKey(string) (string, error)

	// RetrieveAll retrieves the subset of things owned by the specified user.
	RetrieveAll(string, uint64, uint64) ThingsPage

	// RetrieveByChannel retrieves the subset of things owned by the specified
	// user and connected to specified channel.
	RetrieveByChannel(string, string, uint64, uint64) ThingsPage

	// Remove removes the thing having the provided identifier, that is owned
	// by the specified user.
	Remove(string, string) error
}

// ThingCache contains thing caching interface.
type ThingCache interface {
	// Save stores pair thing key, thing id.
	Save(string, string) error

	// ID returns thing ID for given key.
	ID(string) (string, error)

	// Removes thing from cache.
	Remove(string) error
}

// Service specifies an API that must be fullfiled by the domain service
// implementation, and all of its decorators (e.g. logging & metrics).
type Service interface {
	// AddThing adds new thing to the user identified by the provided key.
	AddThing(string, Thing) (Thing, error)
}
