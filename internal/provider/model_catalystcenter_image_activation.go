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
	"context"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ImageActivation struct {
	Id                        types.String `tfsdk:"id"`
	DeviceUuid                types.String `tfsdk:"device_uuid"`
	ImageUuidList             types.Set    `tfsdk:"image_uuid_list"`
	ActivateLowerImageVersion types.Bool   `tfsdk:"activate_lower_image_version"`
	DeviceUpgradeMode         types.String `tfsdk:"device_upgrade_mode"`
	DistributeIfNeeded        types.Bool   `tfsdk:"distribute_if_needed"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ImageActivation) getPath() string {
	return "/dna/intent/api/v1/image/activation/device"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ImageActivation) toBody(ctx context.Context, state ImageActivation) string {
	body := ""
	if !data.DeviceUuid.IsNull() {
		body, _ = sjson.Set(body, "0.deviceUuid", data.DeviceUuid.ValueString())
	}
	if !data.ImageUuidList.IsNull() {
		var values []string
		data.ImageUuidList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "0.imageUuidList", values)
	}
	if !data.ActivateLowerImageVersion.IsNull() {
		body, _ = sjson.Set(body, "0.activateLowerImageVersion", data.ActivateLowerImageVersion.ValueBool())
	}
	if !data.DeviceUpgradeMode.IsNull() {
		body, _ = sjson.Set(body, "0.deviceUpgradeMode", data.DeviceUpgradeMode.ValueString())
	}
	if !data.DistributeIfNeeded.IsNull() {
		body, _ = sjson.Set(body, "0.distributeIfNeeded", data.DistributeIfNeeded.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ImageActivation) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("0.deviceUuid"); value.Exists() {
		data.DeviceUuid = types.StringValue(value.String())
	} else {
		data.DeviceUuid = types.StringNull()
	}
	if value := res.Get("0.imageUuidList"); value.Exists() && len(value.Array()) > 0 {
		data.ImageUuidList = helpers.GetStringSet(value.Array())
	} else {
		data.ImageUuidList = types.SetNull(types.StringType)
	}
	if value := res.Get("0.activateLowerImageVersion"); value.Exists() {
		data.ActivateLowerImageVersion = types.BoolValue(value.Bool())
	} else {
		data.ActivateLowerImageVersion = types.BoolNull()
	}
	if value := res.Get("0.deviceUpgradeMode"); value.Exists() {
		data.DeviceUpgradeMode = types.StringValue(value.String())
	} else {
		data.DeviceUpgradeMode = types.StringNull()
	}
	if value := res.Get("0.distributeIfNeeded"); value.Exists() {
		data.DistributeIfNeeded = types.BoolValue(value.Bool())
	} else {
		data.DistributeIfNeeded = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ImageActivation) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("0.deviceUuid"); value.Exists() && !data.DeviceUuid.IsNull() {
		data.DeviceUuid = types.StringValue(value.String())
	} else {
		data.DeviceUuid = types.StringNull()
	}
	if value := res.Get("0.imageUuidList"); value.Exists() && !data.ImageUuidList.IsNull() {
		data.ImageUuidList = helpers.GetStringSet(value.Array())
	} else {
		data.ImageUuidList = types.SetNull(types.StringType)
	}
	if value := res.Get("0.activateLowerImageVersion"); value.Exists() && !data.ActivateLowerImageVersion.IsNull() {
		data.ActivateLowerImageVersion = types.BoolValue(value.Bool())
	} else {
		data.ActivateLowerImageVersion = types.BoolNull()
	}
	if value := res.Get("0.deviceUpgradeMode"); value.Exists() && !data.DeviceUpgradeMode.IsNull() {
		data.DeviceUpgradeMode = types.StringValue(value.String())
	} else {
		data.DeviceUpgradeMode = types.StringNull()
	}
	if value := res.Get("0.distributeIfNeeded"); value.Exists() && !data.DistributeIfNeeded.IsNull() {
		data.DistributeIfNeeded = types.BoolValue(value.Bool())
	} else {
		data.DistributeIfNeeded = types.BoolNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *ImageActivation) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DeviceUuid.IsNull() {
		return false
	}
	if !data.ImageUuidList.IsNull() {
		return false
	}
	if !data.ActivateLowerImageVersion.IsNull() {
		return false
	}
	if !data.DeviceUpgradeMode.IsNull() {
		return false
	}
	if !data.DistributeIfNeeded.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
