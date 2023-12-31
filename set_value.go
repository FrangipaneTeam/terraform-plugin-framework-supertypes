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
var _ basetypes.SetValuable = SetValue{}

type SetValue struct {
	basetypes.SetValue
}

func (v SetValue) Equal(o attr.Value) bool {
	other, ok := o.(SetValue)

	if !ok {
		return false
	}

	return v.SetValue.Equal(other.SetValue)
}

func (v SetValue) Type(ctx context.Context) attr.Type {
	return SetType{
		SetType: v.SetValue.Type(ctx).(basetypes.SetType),
	}
}

func (v SetValue) ToSetValue(_ context.Context) (basetypes.SetValue, diag.Diagnostics) {
	return v.SetValue, nil
}

func NewSetNull(elementType attr.Type) SetValue {
	return SetValue{
		SetValue: basetypes.NewSetNull(elementType),
	}
}

func NewSetUnknown(elementType attr.Type) SetValue {
	return SetValue{
		SetValue: basetypes.NewSetUnknown(elementType),
	}
}

func NewSetValueMust(elementType attr.Type, elements []attr.Value) SetValue {
	return SetValue{
		SetValue: basetypes.NewSetValueMust(elementType, elements),
	}
}

func NewSetValue(elementType attr.Type, elements []attr.Value) (SetValue, diag.Diagnostics) {
	x, d := basetypes.NewSetValue(elementType, elements)
	return SetValue{
		SetValue: x,
	}, d
}

func NewSetValueFrom(ctx context.Context, elementType attr.Type, elements any) (SetValue, diag.Diagnostics) {
	x, d := basetypes.NewSetValueFrom(ctx, elementType, elements)
	return SetValue{
		SetValue: x,
	}, d
}

// * CustomFunc

func (v *SetValue) Get(ctx context.Context, target interface{}, allowUnhandled bool) (diag diag.Diagnostics) {
	return v.SetValue.ElementsAs(ctx, target, allowUnhandled)
}

func (v *SetValue) Set(ctx context.Context, elements any) diag.Diagnostics {
	var d diag.Diagnostics
	v.SetValue, d = types.SetValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

func (v *SetValue) SetNull(ctx context.Context) {
	v.SetValue = basetypes.NewSetNull(v.ElementType(ctx))
}

func (v *SetValue) SetUnknown(ctx context.Context) {
	v.SetValue = basetypes.NewSetUnknown(v.ElementType(ctx))
}

func (v SetValue) IsKnown() bool {
	return !v.SetValue.IsNull() && !v.SetValue.IsUnknown()
}
