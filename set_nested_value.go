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

type SetNestedValue struct {
	basetypes.SetValue
}

func (v SetNestedValue) Type(ctx context.Context) attr.Type {
	return v.SetValue.Type(ctx)
}

func (v SetNestedValue) ToSetValue(_ context.Context) (basetypes.SetValue, diag.Diagnostics) {
	return v.SetValue, nil
}

func SetNestedNull(elementType attr.Type) SetValue {
	return SetValue{
		SetValue: basetypes.NewSetNull(elementType),
	}
}

func SetNestedUnknown(elementType attr.Type) SetValue {
	return SetValue{
		SetValue: basetypes.NewSetUnknown(elementType),
	}
}

// * CustomFunc

func (v *SetNestedValue) Get(ctx context.Context, target interface{}, allowUnhandled bool) (diag diag.Diagnostics) {
	return v.SetValue.ElementsAs(ctx, target, allowUnhandled)
}

func (v *SetNestedValue) Set(ctx context.Context, elements any) diag.Diagnostics {
	var d diag.Diagnostics
	v.SetValue, d = types.SetValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

func (v *SetNestedValue) SetNull(ctx context.Context) {
	v.SetValue = basetypes.NewSetNull(v.ElementType(ctx))
}

func (v *SetNestedValue) SetUnknown(ctx context.Context) {
	v.SetValue = basetypes.NewSetUnknown(v.ElementType(ctx))
}

func (v SetNestedValue) IsKnown() bool {
	return !v.SetValue.IsNull() && !v.SetValue.IsUnknown()
}
