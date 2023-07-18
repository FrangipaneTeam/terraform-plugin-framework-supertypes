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
	return MapType{
		MapType: v.MapValue.Type(ctx).(basetypes.MapType),
	}
}

func (v MapValue) ToMapValue(_ context.Context) (basetypes.MapValue, diag.Diagnostics) {
	return v.MapValue, nil
}

func NewMapNull(elementType attr.Type) MapValue {
	return MapValue{
		MapValue: basetypes.NewMapNull(elementType),
	}
}

func NewMapUnknown(elementType attr.Type) MapValue {
	return MapValue{
		MapValue: basetypes.NewMapUnknown(elementType),
	}
}

func NewMapValueMust(elementType attr.Type, elements map[string]attr.Value) MapValue {
	return MapValue{
		MapValue: basetypes.NewMapValueMust(elementType, elements),
	}
}

func NewMapValue(elementType attr.Type, elements map[string]attr.Value) (MapValue, diag.Diagnostics) {
	x, d := basetypes.NewMapValue(elementType, elements)
	return MapValue{
		MapValue: x,
	}, d
}

func NewMapValueFrom(ctx context.Context, elementType attr.Type, elements any) (MapValue, diag.Diagnostics) {
	x, d := basetypes.NewMapValueFrom(ctx, elementType, elements)
	return MapValue{
		MapValue: x,
	}, d
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
