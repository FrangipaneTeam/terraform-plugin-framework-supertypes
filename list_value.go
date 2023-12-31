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

type ListValue struct {
	basetypes.ListValue
}

func (v ListValue) Equal(o attr.Value) bool {
	other, ok := o.(ListValue)

	if !ok {
		return false
	}

	return v.ListValue.Equal(other.ListValue)
}

func (v ListValue) Type(ctx context.Context) attr.Type {
	return ListType{
		ListType: v.ListValue.Type(ctx).(basetypes.ListType),
	}
}

func (v ListValue) ToListValue(_ context.Context) (basetypes.ListValue, diag.Diagnostics) {
	return v.ListValue, nil
}

func NewListNull(elementType attr.Type) ListValue {
	return ListValue{
		ListValue: basetypes.NewListNull(elementType),
	}
}

func NewListUnknown(elementType attr.Type) ListValue {
	return ListValue{
		ListValue: basetypes.NewListUnknown(elementType),
	}
}

func NewListValueMust(elementType attr.Type, elements []attr.Value) ListValue {
	return ListValue{
		ListValue: basetypes.NewListValueMust(elementType, elements),
	}
}

func NewListValue(elementType attr.Type, elements []attr.Value) (ListValue, diag.Diagnostics) {
	x, d := basetypes.NewListValue(elementType, elements)
	return ListValue{
		ListValue: x,
	}, d
}

func NewListValueFrom(ctx context.Context, elementType attr.Type, elements any) (ListValue, diag.Diagnostics) {
	x, d := basetypes.NewListValueFrom(ctx, elementType, elements)
	return ListValue{
		ListValue: x,
	}, d
}

// * CustomFunc

func (v *ListValue) Get(ctx context.Context, target interface{}, allowUnhandled bool) (diag diag.Diagnostics) {
	return v.ListValue.ElementsAs(ctx, target, allowUnhandled)
}

func (v *ListValue) Set(ctx context.Context, elements any) diag.Diagnostics {
	var d diag.Diagnostics
	v.ListValue, d = types.ListValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

func (v *ListValue) SetNull(ctx context.Context) {
	v.ListValue = basetypes.NewListNull(v.ElementType(ctx))
}

func (v *ListValue) SetUnknown(ctx context.Context) {
	v.ListValue = basetypes.NewListUnknown(v.ElementType(ctx))
}

func (v ListValue) IsKnown() bool {
	return !v.ListValue.IsNull() && !v.ListValue.IsUnknown()
}
