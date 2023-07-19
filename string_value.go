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
var _ basetypes.StringValuable = StringValue{}

type StringValue struct {
	basetypes.StringValue
}

func (v StringValue) Equal(o attr.Value) bool {
	other, ok := o.(StringValue)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v StringValue) Type(ctx context.Context) attr.Type {
	return StringType{
		StringType: v.StringValue.Type(ctx).(basetypes.StringType),
	}
}

func NewStringNull() StringValue {
	return StringValue{
		StringValue: basetypes.NewStringNull(),
	}
}

func NewStringUnknown() StringValue {
	return StringValue{
		StringValue: basetypes.NewStringUnknown(),
	}
}

func NewStringValue(s string) StringValue {
	return StringValue{
		StringValue: basetypes.NewStringValue(s),
	}
}

func NewStringPointerValue(s *string) StringValue {
	return StringValue{
		StringValue: basetypes.NewStringPointerValue(s),
	}
}

// * CustomFunc

func (v *StringValue) Get() string {
	return v.StringValue.ValueString()
}

func (v *StringValue) GetPtr() *string {
	return v.StringValue.ValueStringPointer()
}

func (v *StringValue) Set(s string) {
	v.StringValue = basetypes.NewStringValue(s)
}

func (v *StringValue) SetNull() {
	v.StringValue = basetypes.NewStringNull()
}

func (v *StringValue) SetUnknown() {
	v.StringValue = basetypes.NewStringUnknown()
}

func (v StringValue) IsKnown() bool {
	return !v.StringValue.IsNull() && !v.StringValue.IsUnknown()
}
