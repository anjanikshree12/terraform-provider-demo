package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure ExampleResource satisfies required interfaces.
var _ resource.Resource = &ExampleResource{}
var _ resource.ResourceWithImportState = &ExampleResource{}

// NewExampleResource returns a new instance of ExampleResource.
func NewExampleResource() resource.Resource {
	return &ExampleResource{}
}

// ExampleResource is the resource implementation.
type ExampleResource struct {
	// In a real provider, you might set an API client here.
	client *http.Client
}

// ExampleResourceModel describes the resource data model.
type ExampleResourceModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	Id                    types.String `tfsdk:"id"`
}

// Metadata sets the resource type name.
func (r *ExampleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_example"
}

// Schema defines the schema for the resource.
func (r *ExampleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "An example resource that demonstrates basic CRUD operations.",
		Attributes: map[string]schema.Attribute{
			"configurable_attribute": schema.StringAttribute{
				MarkdownDescription: "An example attribute that can be set by the user.",
				Optional:            true,
			},
			"defaulted": schema.StringAttribute{
				MarkdownDescription: "An attribute with a default value if not provided.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("example value when not configured"),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The identifier for this resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

// Configure attaches any provider-level data to the resource. Here we simulate using an HTTP client.
func (r *ExampleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*http.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.client = client
}

// Create simulates resource creation by setting a static ID.
func (r *ExampleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ExampleResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Simulate creation by setting an ID.
	data.Id = types.StringValue("example-id")
	tflog.Trace(ctx, "Created example resource")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read simply copies the state over.
func (r *ExampleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ExampleResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update copies the plan to state.
func (r *ExampleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ExampleResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op for this example.
func (r *ExampleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Nothing to do for deletion in this simple example.
}

// ImportState allows the resource to be imported.
func (r *ExampleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
