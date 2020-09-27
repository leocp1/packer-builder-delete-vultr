// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Read or delete vultr snapshots
package vultr

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer/packer"
	"github.com/vultr/govultr"
	"github.com/vultr/packer-builder-vultr/vultr"
)

const (
	// BuilderID is the unique ID for the builder
	BuilderID = "packer.vultr"

	// Name of the executable if being used for deletion
	DeleteExeName = "packer-builder-delete-vultr"
)

// Builder ...
type Builder struct {
	// Change whether the builder will read or delete if os.Args[0] cannot be
	// relied on. ForceRead takes priority over ForceDelete
	ForceDelete bool
	ForceRead bool

	config vultr.Config
}

func (b *Builder) ConfigSpec() hcldec.ObjectSpec {
	return b.config.FlatMapstructure().HCL2Spec()
}

func (b *Builder) Prepare(raws ...interface{}) ([]string, []string, error) {
	warnings, errs := b.config.Prepare(raws...)
	if errs != nil {
		return nil, warnings, errs
	}
	return nil, nil, nil
}

// Determine if the executable is named
// packer-builder-read-vultr
// or
// packer-builder-delete-vultr
// to figure out what operation we should run
func ReadOnly() bool {
	return !strings.HasPrefix(
		filepath.Base(os.Args[0]),
		DeleteExeName,
	)
}

func (b *Builder) Run(
	ctx context.Context,
	ui packer.Ui,
	hook packer.Hook,
) (ret packer.Artifact, err error) {

	description := b.config.Description

	ui.Say(fmt.Sprintf(
		"Reading Vultr snapshots with Description=%s",
		description,
	))

	a := &Artifact{
		ID:           "",
		ResourceName: description,
		Count:        0,
	}

	cli := govultr.NewClient(nil, b.config.APIKey)

	ssh := &govultr.SnapshotServiceHandler{Client: cli}
	all_snapshots, err := ssh.List(ctx)
	if err != nil {
		return a, nil
	}

	snapshots := make([]govultr.Snapshot, 0)
	for _, s := range all_snapshots {
		if s.Description == description {
			snapshots = append(snapshots, s)
		}
	}

	a.Count = len(snapshots)

	// Sort snapshots by date
	sort.Slice(snapshots, func(lhs, rhs int) bool {
		tfmt := "2006-01-02 15:04:05"
		lhst, _ := time.Parse(tfmt, snapshots[lhs].DateCreated)
		rhst, _ := time.Parse(tfmt, snapshots[rhs].DateCreated)
		return lhst.Before(rhst)
	})
	if len(snapshots) > 0 {
		a.ID = snapshots[len(snapshots)-1].SnapshotID
	}

	if !b.ForceRead && (b.ForceDelete || !ReadOnly()) {
		ui.Say(fmt.Sprintf(
			"Deleting Vultr snapshots with Description=%s",
			description,
		))
		for _, s := range snapshots {
			ui.Say(fmt.Sprintf(
				"Deleting snapshot %s",
				s.SnapshotID,
			))
			err = ssh.Delete(ctx, s.SnapshotID)
			if err != nil {
				return a, err
			}
		}
	}

	return a, nil
}
