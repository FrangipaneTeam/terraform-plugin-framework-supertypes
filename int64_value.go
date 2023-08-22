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

func (v *Int64Value) SetPtr(s *int64) {
	if s == nil {
		v.Int64Value = basetypes.NewInt64Null()
		return
	}

	v.Int64Value = basetypes.NewInt64PointerValue(s)
}

func (v Int64Value) SetInt(s int) {
	v.Int64Value = basetypes.NewInt64Value(int64(s))
}

func (v Int64Value) SetInt8(s int8) {
	v.Int64Value = basetypes.NewInt64Value(int64(s))
}

func (v Int64Value) SetInt16(s int16) {
	v.Int64Value = basetypes.NewInt64Value(int64(s))
}

func (v Int64Value) SetInt32(s int32) {
	v.Int64Value = basetypes.NewInt64Value(int64(s))
}

func (v Int64Value) SetInt64(s int64) {
	v.Set(s)
}

func (v Int64Value) SetIntPtr(s *int) {
	v.Int64Value = basetypes.NewInt64PointerValue(int64(s))
}

func (v Int64Value) SetInt8Ptr(s *int8) {
	v.Int64Value = basetypes.NewInt64PointerValue(int64(s))
}

func (v Int64Value) SetInt16Ptr(s *int16) {
	v.Int64Value = basetypes.NewInt64PointerValue(int64(s))
}

func (v Int64Value) SetInt32Ptr(s *int32) {
	v.Int64Value = basetypes.NewInt64PointerValue(int64(s))
}

func (v Int64Value) SetInt64Ptr(s *int64) {
	v.SetPtr(s)
}

func (v Int64Value) GetInt() int {
	return int(v.Get())
}

func (v Int64Value) GetInt8() int8 {
	return int8(v.Get())
}

func (v Int64Value) GetInt16() int16 {
	return int16(v.Get())
}

func (v Int64Value) GetInt32() int32 {
	return int32(v.Get())
}

func (v Int64Value) GetInt64() int64 {
	return v.Get()
}

func (v Int64Value) GetIntPtr() *int {
	if v.IsKnown() {
		i := int(v.Get())
		return &i
	}

	return nil
}

func (v Int64Value) GetInt8Ptr() *int8 {
	if v.IsKnown() {
		i := int8(v.Get())
		return &i
	}

	return nil
}

func (v Int64Value) GetInt16Ptr() *int16 {
	if v.IsKnown() {
		i := int16(v.Get())
		return &i
	}

	return nil
}

func (v Int64Value) GetInt32Ptr() *int32 {
	if v.IsKnown() {
		i := int32(v.Get())
		return &i
	}

	return nil
}

func (v Int64Value) GetInt64Ptr() *int64 {
	if v.IsKnown() {
		i := v.Get()
		return &i
	}

	return nil
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
