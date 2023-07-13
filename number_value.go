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

func (v NumberValue) Type(_ context.Context) attr.Type {
	return NumberType{}
}

func NumberNull() NumberValue {
	return NumberValue{
		NumberValue: basetypes.NewNumberNull(),
	}
}

func NumberUnknown() NumberValue {
	return NumberValue{
		NumberValue: basetypes.NewNumberUnknown(),
	}
}

// * CustomFunc

func (v *NumberValue) Get() *big.Float {
	return v.NumberValue.ValueBigFloat()
}

func (v NumberValue) Set(s *big.Float) {
	v.NumberValue = basetypes.NewNumberValue(s)
}

func (v NumberValue) SetNull() {
	v.NumberValue = basetypes.NewNumberNull()
}

func (v NumberValue) SetUnknown() {
	v.NumberValue = basetypes.NewNumberUnknown()
}

func (v NumberValue) IsKnown() bool {
	return !v.NumberValue.IsNull() && !v.NumberValue.IsUnknown()
}
