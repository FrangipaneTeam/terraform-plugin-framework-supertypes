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
var _ basetypes.Float64Valuable = Float64Value{}

type Float64Value struct {
	basetypes.Float64Value
}

func (v Float64Value) Equal(o attr.Value) bool {
	other, ok := o.(Float64Value)

	if !ok {
		return false
	}

	return v.Float64Value.Equal(other.Float64Value)
}

func (v Float64Value) Type(ctx context.Context) attr.Type {
	return Float64Type{
		Float64Type: v.Float64Value.Type(ctx).(basetypes.Float64Type),
	}
}

func NewFloat64Null() Float64Value {
	return Float64Value{
		Float64Value: basetypes.NewFloat64Null(),
	}
}

func NewFloat64Unknown() Float64Value {
	return Float64Value{
		Float64Value: basetypes.NewFloat64Unknown(),
	}
}

func NewFloat64Value(s float64) Float64Value {
	return Float64Value{
		Float64Value: basetypes.NewFloat64Value(s),
	}
}

func NewFloat64PointerValue(s *float64) Float64Value {
	return Float64Value{
		Float64Value: basetypes.NewFloat64PointerValue(s),
	}
}

// * CustomFunc

func (v *Float64Value) Get() float64 {
	return v.Float64Value.ValueFloat64()
}

func (v *Float64Value) GetPtr() *float64 {
	return v.Float64Value.ValueFloat64Pointer()
}

func (v *Float64Value) Set(s float64) {
	v.Float64Value = basetypes.NewFloat64Value(s)
}

func (v *Float64Value) SetNull() {
	v.Float64Value = basetypes.NewFloat64Null()
}

func (v *Float64Value) SetUnknown() {
	v.Float64Value = basetypes.NewFloat64Unknown()
}

func (v Float64Value) IsKnown() bool {
	return !v.Float64Value.IsNull() && !v.Float64Value.IsUnknown()
}
