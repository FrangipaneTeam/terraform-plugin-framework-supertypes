// -------------------------------------------------------------------- //
// !      DO NOT EDIT. This file is auto-generated from template      ! //
// -------------------------------------------------------------------- //
package supertypes

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.NumberValuable = NumberValue{}

type NumberValue struct {
	basetypes.NumberValue
}

func (v NumberValue) Equal(o attr.Value) bool {
	other, ok := o.(NumberValue)

	if !ok {
		return false
	}

	return v.NumberValue.Equal(other.NumberValue)
}

func (v NumberValue) Type(ctx context.Context) attr.Type {
	return NumberType{
		NumberType: v.NumberValue.Type(ctx).(basetypes.NumberType),
	}
}

func NewNumberNull() NumberValue {
	return NumberValue{
		NumberValue: basetypes.NewNumberNull(),
	}
}

func NewNumberUnknown() NumberValue {
	return NumberValue{
		NumberValue: basetypes.NewNumberUnknown(),
	}
}

func NewNumberValue(s *big.Float) NumberValue {
	return NumberValue{
		NumberValue: basetypes.NewNumberValue(s),
	}
}

// * CustomFunc

func (v *NumberValue) Get() *big.Float {
	return v.NumberValue.ValueBigFloat()
}

func (v *NumberValue) Set(s *big.Float) {

	v.NumberValue = basetypes.NewNumberValue(s)
}

func (v *NumberValue) SetPtr(s **big.Float) {
	if s == nil {
		v.NumberValue = basetypes.NewNumberNull()
		return
	}

	v.NumberValue = basetypes.NewNumberPointerValue(s)
}

func (v *NumberValue) SetNull() {
	v.NumberValue = basetypes.NewNumberNull()
}

func (v *NumberValue) SetUnknown() {
	v.NumberValue = basetypes.NewNumberUnknown()
}

func (v NumberValue) IsKnown() bool {
	return !v.NumberValue.IsNull() && !v.NumberValue.IsUnknown()
}
