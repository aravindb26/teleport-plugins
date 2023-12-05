// Code generated by _gen/main.go DO NOT EDIT
/*
Copyright 2015-2022 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"context"

	apitypes "github.com/gravitational/teleport/api/types"
    
	"github.com/gravitational/trace"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/gravitational/teleport-plugins/terraform/tfschema"
)

// dataSourceTeleportSAMLConnectorType is the data source metadata type
type dataSourceTeleportSAMLConnectorType struct{}

// dataSourceTeleportSAMLConnector is the resource
type dataSourceTeleportSAMLConnector struct {
	p Provider
}

// GetSchema returns the data source schema
func (r dataSourceTeleportSAMLConnectorType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfschema.GenSchemaSAMLConnectorV2(ctx)
}

// NewDataSource creates the empty data source
func (r dataSourceTeleportSAMLConnectorType) NewDataSource(_ context.Context, p tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	return dataSourceTeleportSAMLConnector{
		p: *(p.(*Provider)),
	}, nil
}

// Read reads teleport SAMLConnector
func (r dataSourceTeleportSAMLConnector) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var id types.String
	diags := req.Config.GetAttribute(ctx, path.Root("metadata").AtName("name"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	samlConnectorI, err := r.p.Client.GetSAMLConnector(ctx, id.Value, true)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading SAMLConnector", trace.Wrap(err), "saml"))
		return
	}

    var state types.Object
	
	samlConnector := samlConnectorI.(*apitypes.SAMLConnectorV2)
	diags = tfschema.CopySAMLConnectorV2ToTerraform(ctx, samlConnector, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
