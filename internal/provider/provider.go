package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure DemoProvider implements the provider.Provider interface.
var _ provider.Provider = &DemoProvider{}

// DemoProvider implements a basic provider.
type DemoProvider struct {
	version string
}

// New: returns a function that returns a provider.Provider
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &DemoProvider{version: version}
	}
}

// Metadata sets the provider type name and version.
func (p *DemoProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "demo"
	resp.Version = p.version
}

// Schema defines the provider schema (empty in this example).
func (p *DemoProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Demo Terraform Provider",
	}
}

// Configure would normally prepare a client for API calls. For now, it's empty.
func (p *DemoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// For our simple example, we don't need to configure any clients.
}

// DataSources returns nil as this provider doesn't define any data sources.
func (p *DemoProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}

// Resources returns the list of resources supported by the provider.
func (p *DemoProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewExampleResource,
	}
}
