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

func (v BoolValue) Type(ctx context.Context) attr.Type {
	return BoolType{
		BoolType: v.BoolValue.Type(ctx).(basetypes.BoolType),
	}
}

func NewBoolNull() BoolValue {
	return BoolValue{
		BoolValue: basetypes.NewBoolNull(),
	}
}

func NewBoolUnknown() BoolValue {
	return BoolValue{
		BoolValue: basetypes.NewBoolUnknown(),
	}
}

func NewBoolValue(s bool) BoolValue {
	return BoolValue{
		BoolValue: basetypes.NewBoolValue(s),
	}
}

func NewBoolPointerValue(s *bool) BoolValue {
	return BoolValue{
		BoolValue: basetypes.NewBoolPointerValue(s),
	}
}

// * CustomFunc

func (v *BoolValue) Get() bool {
	return v.BoolValue.ValueBool()
}

func (v *BoolValue) GetPtr() *bool {
	return v.BoolValue.ValueBoolPointer()
}

func (v *BoolValue) Set(s bool) {
	v.BoolValue = basetypes.NewBoolValue(s)
}

func (v *BoolValue) SetNull() {
	v.BoolValue = basetypes.NewBoolNull()
}

func (v *BoolValue) SetUnknown() {
	v.BoolValue = basetypes.NewBoolUnknown()
}

func (v BoolValue) IsKnown() bool {
	return !v.BoolValue.IsNull() && !v.BoolValue.IsUnknown()
}
