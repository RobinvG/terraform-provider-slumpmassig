// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"terraform-provider-slumpmassig/internal/products"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &ProductsResource{}
var _ resource.ResourceWithImportState = &ProductsResource{}

func NewProductsResource() resource.Resource {
	return &ProductsResource{}
}

// ExampleResource defines the resource implementation.
type ProductsResource struct {
	ID         types.String `tfsdk:"id"`
	Length     types.Int64  `tfsdk:"length"`
	Spongecase types.String   `tfsdk:"spongecase"`
	L33t       types.String   `tfsdk:"l33t"`
	Diacritics types.String   `tfsdk:"diacritics"`
	Result     types.String `tfsdk:"result"`
}

// ExampleResourceModel describes the resource data model.
type ProductsResourceModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	Id                    types.String `tfsdk:"id"`
}

func (r *ProductsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_products"
}

func (r *ProductsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Get random products of a known swedish warehouse.",
		Attributes: map[string]schema.Attribute{

			"length": schema.Int64Attribute{
				Description: "The maximum length of the generated product",
				Optional:    true,
				Computed:   true,
				Default:     int64default.StaticInt64(0),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Validators: []validator.Int64{
					int64validator.AtLeast(5),
				},
			},

			"spongecase": schema.StringAttribute{
				Description: "The generated random product in sponge case. Default value is `false`.",
				Sensitive:   true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},

			"l33t": schema.StringAttribute{
				Description: "The generated random product in l33t  . Default value is `false`.",
				Sensitive:   true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"diacritics": schema.StringAttribute{
				Description: "The generated random product with converted diacritics . Default value is `false`.",
				Sensitive:   true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"result": schema.StringAttribute{
				Description: "The generated random product.",
				Computed:    true,
				Sensitive:   true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},

			"id": schema.StringAttribute{
				Description: "A static value used internally by Terraform, this should not be referenced in configurations.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *ProductsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
}

func (r *ProductsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ProductsResource

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	params := products.ProductsParams{
		Length:     plan.Length.ValueInt64(),
	}

	result, err := params.ReturnProduct()
	if err != nil {
		resp.Diagnostics.AddError("Error generating product", err.Error())
		return
	}
	plan.Result = types.StringValue(result.Result)
	plan.Spongecase = types.StringValue(result.Spongecase)
	plan.L33t = types.StringValue(result.L33t)
	plan.Diacritics = types.StringValue(result.Diacritics)
	plan.ID = types.StringValue("none")
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)


}

func (r *ProductsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *ProductsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var model ProductsResource
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *ProductsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *ProductsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
}
