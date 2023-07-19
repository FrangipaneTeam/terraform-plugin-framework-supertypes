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
var _ basetypes.ObjectValuable = SingleNestedValue{}

type SingleNestedValue struct {
	basetypes.ObjectValue
}

func (v SingleNestedValue) Type(ctx context.Context) attr.Type {
	return v.ObjectValue.Type(ctx)
}

func (v SingleNestedValue) ToObjectValue(_ context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return v.ObjectValue, nil
}

func NewSingleNestedNull(attributeTypes map[string]attr.Type) SingleNestedValue {
	return SingleNestedValue{
		ObjectValue: basetypes.NewObjectNull(attributeTypes),
	}
}

func NewSingleNestedUnknown(attributeTypes map[string]attr.Type) SingleNestedValue {
	return SingleNestedValue{
		ObjectValue: basetypes.NewObjectUnknown(attributeTypes),
	}
}

// * CustomFunc

func (v *SingleNestedValue) Get(ctx context.Context, target interface{}, opts basetypes.ObjectAsOptions) (diag diag.Diagnostics) {
	return v.ObjectValue.As(ctx, target, opts)
}

func (v *SingleNestedValue) Set(ctx context.Context, structure any) diag.Diagnostics {
	var d diag.Diagnostics
	v.ObjectValue, d = types.ObjectValueFrom(ctx, v.AttributeTypes(ctx), structure)
	return d
}

func (v *SingleNestedValue) SetNull(ctx context.Context) {
	v.ObjectValue = basetypes.NewObjectNull(v.AttributeTypes(ctx))
}

func (v *SingleNestedValue) SetUnknown(ctx context.Context) {
	v.ObjectValue = basetypes.NewObjectUnknown(v.AttributeTypes(ctx))
}

func (v SingleNestedValue) IsKnown() bool {
	return !v.ObjectValue.IsNull() && !v.ObjectValue.IsUnknown()
}
