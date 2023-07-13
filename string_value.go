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

func (v StringValue) Type(_ context.Context) attr.Type {
	// CustomStringType defined in the schema type section
	return StringType{}
}

func StringNull() StringValue {
	return StringValue{
		StringValue: basetypes.NewStringNull(),
	}
}

func StringUnknown() StringValue {
	return StringValue{
		StringValue: basetypes.NewStringUnknown(),
	}
}

// * CustomFunc

func (v *StringValue) Get() string {
	return v.StringValue.ValueString()
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
