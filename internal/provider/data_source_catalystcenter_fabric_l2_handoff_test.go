// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceCcFabricL2Handoff(t *testing.T) {
	if os.Getenv("SDA") == "" {
		t.Skip("skipping test, set environment variable SDA")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l2_handoff.test", "network_device_id", "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l2_handoff.test", "interface_name", "GigabitEthernet1/0/4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l2_handoff.test", "internal_vlan_id", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l2_handoff.test", "external_vlan_id", "400"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcFabricL2HandoffPrerequisitesConfig + testAccDataSourceCcFabricL2HandoffConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcFabricL2HandoffPrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}
resource "catalystcenter_fabric_site" "test" {
  site_id                     = catalystcenter_area.test.id
  pub_sub_enabled             = false
  authentication_profile_name = "No Authentication"
  depends_on = [catalystcenter_area.test]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcFabricL2HandoffConfig() string {
	config := `resource "catalystcenter_fabric_l2_handoff" "test" {` + "\n"
	config += `	network_device_id = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"` + "\n"
	config += `	fabric_id = catalystcenter_fabric_site.test.id` + "\n"
	config += `	interface_name = "GigabitEthernet1/0/4"` + "\n"
	config += `	internal_vlan_id = 300` + "\n"
	config += `	external_vlan_id = 400` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_fabric_l2_handoff" "test" {
			id = catalystcenter_fabric_l2_handoff.test.id
			network_device_id = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
			fabric_id = catalystcenter_fabric_site.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
