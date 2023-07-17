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

type MapValue struct {
	basetypes.MapValue
}

func (v MapValue) Equal(o attr.Value) bool {
	other, ok := o.(MapValue)

	if !ok {
		return false
	}

	return v.MapValue.Equal(other.MapValue)
}

func (v MapValue) Type(ctx context.Context) attr.Type {
	// CustomMapType defined in the schema type section
	return v.MapValue.Type(ctx)
}

func (v MapValue) ToMapValue(_ context.Context) (basetypes.MapValue, diag.Diagnostics) {
	return v.MapValue, nil
}

func MapNull(elementType attr.Type) MapValue {
	return MapValue{
		MapValue: basetypes.NewMapNull(elementType),
	}
}

func MapUnknown(elementType attr.Type) MapValue {
	return MapValue{
		MapValue: basetypes.NewMapUnknown(elementType),
	}
}

// * CustomFunc

func (v *MapValue) Get(ctx context.Context, target interface{}, allowUnhandled bool) (diag diag.Diagnostics) {
	return v.MapValue.ElementsAs(ctx, target, allowUnhandled)
}

func (v *MapValue) Set(ctx context.Context, elements any) diag.Diagnostics {
	var d diag.Diagnostics
	v.MapValue, d = types.MapValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

func (v *MapValue) SetNull(ctx context.Context) {
	v.MapValue = basetypes.NewMapNull(v.ElementType(ctx))
}

func (v *MapValue) SetUnknown(ctx context.Context) {
	v.MapValue = basetypes.NewMapUnknown(v.ElementType(ctx))
}

func (v MapValue) IsKnown() bool {
	return !v.MapValue.IsNull() && !v.MapValue.IsUnknown()
}
