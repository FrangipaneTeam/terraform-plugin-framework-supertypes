// -------------------------------------------------------------------- //
// !      DO NOT EDIT. This file is auto-generated from template      ! //
// -------------------------------------------------------------------- //
package supertypes

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.ObjectValuable = ObjectValue{}

type ObjectValue struct {
	basetypes.ObjectValue
}

func (v ObjectValue) Equal(o attr.Value) bool {
	other, ok := o.(ObjectValue)

	if !ok {
		return false
	}

	return v.ObjectValue.Equal(other.ObjectValue)
}

func (v ObjectValue) Type(_ context.Context) attr.Type {
	return ObjectType{}
}

func NewObjectNull(attributeTypes map[string]attr.Type) ObjectValue {
	return ObjectValue{
		ObjectValue: basetypes.NewObjectNull(attributeTypes),
	}
}

func NewObjectUnknown(attributeTypes map[string]attr.Type) ObjectValue {
	return ObjectValue{
		ObjectValue: basetypes.NewObjectUnknown(attributeTypes),
	}
}

// * CustomFunc

func (v *ObjectValue) Get(ctx context.Context, target interface{}, opts basetypes.ObjectAsOptions) diag.Diagnostics {
	return v.ObjectValue.As(ctx, target, opts)
}

func (v *ObjectValue) Set(ctx context.Context, structure any) diag.Diagnostics {
	var d diag.Diagnostics
	v.ObjectValue, d = basetypes.NewObjectValueFrom(ctx, v.AttributeTypes(ctx), structure)
	return d
}

func (v *ObjectValue) SetNull(ctx context.Context) {
	v.ObjectValue = basetypes.NewObjectNull(v.AttributeTypes(ctx))
}

func (v *ObjectValue) SetUnknown(ctx context.Context) {
	v.ObjectValue = basetypes.NewObjectUnknown(v.AttributeTypes(ctx))
}

func (v ObjectValue) IsKnown() bool {
	return !v.ObjectValue.IsNull() && !v.ObjectValue.IsUnknown()
}
