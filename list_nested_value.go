// -------------------------------------------------------------------- //
// !      DO NOT EDIT. This file is auto-generated from template      ! //
// -------------------------------------------------------------------- //
package supertypes

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.ListValuable = ListValue{}

type ListNestedValue struct {
	basetypes.ListValue
}

func (v ListNestedValue) Type(ctx context.Context) attr.Type {
	return v.ListValue.Type(ctx)
}

func (v ListNestedValue) ToListValue(_ context.Context) (basetypes.ListValue, diag.Diagnostics) {
	return v.ListValue, nil
}

func NewListNestedNull(elementType attr.Type) ListValue {
	return ListValue{
		ListValue: basetypes.NewListNull(elementType),
	}
}

func NewListNestedUnknown(elementType attr.Type) ListValue {
	return ListValue{
		ListValue: basetypes.NewListUnknown(elementType),
	}
}

// * CustomFunc

func (v *ListNestedValue) Get(ctx context.Context, target interface{}, allowUnhandled bool) (diag diag.Diagnostics) {
	return v.ListValue.ElementsAs(ctx, target, allowUnhandled)
}

func (v *ListNestedValue) Set(ctx context.Context, elements any) diag.Diagnostics {
	var d diag.Diagnostics
	v.ListValue, d = types.ListValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

func (v *ListNestedValue) SetNull(ctx context.Context) {
	v.ListValue = basetypes.NewListNull(v.ElementType(ctx))
}

func (v *ListNestedValue) SetUnknown(ctx context.Context) {
	v.ListValue = basetypes.NewListUnknown(v.ElementType(ctx))
}

func (v ListNestedValue) IsKnown() bool {
	return !v.ListValue.IsNull() && !v.ListValue.IsUnknown()
}
