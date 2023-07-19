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
var _ basetypes.Int64Valuable = Int64Value{}

type Int64Value struct {
	basetypes.Int64Value
}

func (v Int64Value) Equal(o attr.Value) bool {
	other, ok := o.(Int64Value)

	if !ok {
		return false
	}

	return v.Int64Value.Equal(other.Int64Value)
}

func (v Int64Value) Type(ctx context.Context) attr.Type {
	return Int64Type{
		Int64Type: v.Int64Value.Type(ctx).(basetypes.Int64Type),
	}
}

func NewInt64Null() Int64Value {
	return Int64Value{
		Int64Value: basetypes.NewInt64Null(),
	}
}

func NewInt64Unknown() Int64Value {
	return Int64Value{
		Int64Value: basetypes.NewInt64Unknown(),
	}
}

func NewInt64Value(s int64) Int64Value {
	return Int64Value{
		Int64Value: basetypes.NewInt64Value(s),
	}
}

func NewInt64PointerValue(s *int64) Int64Value {
	return Int64Value{
		Int64Value: basetypes.NewInt64PointerValue(s),
	}
}

// * CustomFunc

func (v *Int64Value) Get() int64 {
	return v.Int64Value.ValueInt64()
}

func (v *Int64Value) GetPtr() *int64 {
	return v.Int64Value.ValueInt64Pointer()
}

func (v *Int64Value) Set(s int64) {
	v.Int64Value = basetypes.NewInt64Value(s)
}

func (v *Int64Value) SetNull() {
	v.Int64Value = basetypes.NewInt64Null()
}

func (v *Int64Value) SetUnknown() {
	v.Int64Value = basetypes.NewInt64Unknown()
}

func (v Int64Value) IsKnown() bool {
	return !v.Int64Value.IsNull() && !v.Int64Value.IsUnknown()
}
