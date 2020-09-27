// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package vultr_test

import (
	"os"
	"testing"

	builderT "github.com/hashicorp/packer/helper/builder/testing"

	. "github.com/leocp1/packer-builder-delete-vultr/pkg/vultr"
)

func TestBuilderAcc_read(t *testing.T) {
	builderT.Test(t, builderT.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Builder:  &Builder{ForceRead:true},
		Template: testBuilderAccBasic,
	})
}

func TestBuilderAcc_delete(t *testing.T) {
	builderT.Test(t, builderT.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Builder:  &Builder{ForceDelete:true},
		Template: testBuilderAccBasic,
	})
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("VULTR_API_KEY"); v == "" {
		t.Fatal("VULTR_API_KEY must be set for acceptance tests")
	}
}

const testBuilderAccBasic = `
{
	"builders": [{
		"type": "test",
		"snapshot_description": "packer-test-snapshot",
        "region_id": 4,
        "plan_id": 402,
        "os_id": 127,
        "ssh_username": "root",
		"state_timeout": "20m"
	}]
}
`
