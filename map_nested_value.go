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
var _ basetypes.MapValuable = MapValue{}

type MapNestedValue struct {
	basetypes.MapValue
}

func (v MapNestedValue) Type(ctx context.Context) attr.Type {
	return v.MapValue.Type(ctx)
}

func (v MapNestedValue) ToMapValue(_ context.Context) (basetypes.MapValue, diag.Diagnostics) {
	return v.MapValue, nil
}

func MapNestedNull(elementType attr.Type) MapValue {
	return MapValue{
		MapValue: basetypes.NewMapNull(elementType),
	}
}

func MapNestedUnknown(elementType attr.Type) MapValue {
	return MapValue{
		MapValue: basetypes.NewMapUnknown(elementType),
	}
}

// * CustomFunc

func (v *MapNestedValue) Get(ctx context.Context, target interface{}, allowUnhandled bool) (diag diag.Diagnostics) {
	return v.MapValue.ElementsAs(ctx, target, allowUnhandled)
}

func (v *MapNestedValue) Set(ctx context.Context, elements any) diag.Diagnostics {
	var d diag.Diagnostics
	v.MapValue, d = types.MapValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

func (v *MapNestedValue) SetNull(ctx context.Context) {
	v.MapValue = basetypes.NewMapNull(v.ElementType(ctx))
}

func (v *MapNestedValue) SetUnknown(ctx context.Context) {
	v.MapValue = basetypes.NewMapUnknown(v.ElementType(ctx))
}

func (v MapNestedValue) IsKnown() bool {
	return !v.MapValue.IsNull() && !v.MapValue.IsUnknown()
}
