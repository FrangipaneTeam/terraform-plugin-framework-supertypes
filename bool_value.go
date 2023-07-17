// -------------------------------------------------------------------- //
// !      DO NOT EDIT. This file is auto-generated from template      ! //
// -------------------------------------------------------------------- //
package supertypes

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.BoolValuable = BoolValue{}

type BoolValue struct {
	basetypes.BoolValue
}

func (v BoolValue) Equal(o attr.Value) bool {
	other, ok := o.(BoolValue)

	if !ok {
		return false
	}

	return v.BoolValue.Equal(other.BoolValue)
}

func (v BoolValue) Type(_ context.Context) attr.Type {
	return BoolType{}
}

func BoolNull() BoolValue {
	return BoolValue{
		BoolValue: basetypes.NewBoolNull(),
	}
}

func BoolUnknown() BoolValue {
	return BoolValue{
		BoolValue: basetypes.NewBoolUnknown(),
	}
}

// * CustomFunc

func (v *BoolValue) Get() bool {
	return v.BoolValue.ValueBool()
}

func (v BoolValue) Set(s bool) {
	v.BoolValue = basetypes.NewBoolValue(s)
}

func (v BoolValue) SetNull() {
	v.BoolValue = basetypes.NewBoolNull()
}

func (v BoolValue) SetUnknown() {
	v.BoolValue = basetypes.NewBoolUnknown()
}

func (v BoolValue) IsKnown() bool {
	return !v.BoolValue.IsNull() && !v.BoolValue.IsUnknown()
}
