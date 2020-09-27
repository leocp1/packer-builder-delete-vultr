// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package vultr

import (
	"strconv"
)

type Artifact struct {
	// The ID of the snapshot
	ID string

	// The Description of the snapshot
	ResourceName string

	// The number of snapshots
	Count int
}

func (a *Artifact) BuilderId() string {
	return BuilderID
}

func (a *Artifact) Files() []string {
	return nil
}

func (a *Artifact) Id() string {
	return a.ID
}

func (a *Artifact) String() string {
	return strconv.Itoa(a.Count)
}

func (a *Artifact) State(name string) interface{} {
	return nil
}

func (a *Artifact) Destroy() error {
	return nil
}
