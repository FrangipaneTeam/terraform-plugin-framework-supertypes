// -------------------------------------------------------------------- //
// !      DO NOT EDIT. This file is auto-generated from template      ! //
// -------------------------------------------------------------------- //
package supertypes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.SetTypable = SetType{}

type SetType struct {
	basetypes.SetType
}

func (t SetType) Equal(o attr.Type) bool {
	other, ok := o.(SetType)

	if !ok {
		return false
	}

	return t.SetType.Equal(other.SetType)
}

func (t SetType) String() string {
	return "types.SetType[" + t.ElementType().String() + "]"
}

func (t SetType) ValueFromSet(_ context.Context, in basetypes.SetValue) (basetypes.SetValuable, diag.Diagnostics) {
	value := SetValue{
		SetValue: in,
	}

	return value, nil
}

func (t SetType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.SetType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	SetValue, ok := attrValue.(basetypes.SetValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	SetValuable, diags := t.ValueFromSet(ctx, SetValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting SetValue to SetValuable: %v", diags)
	}

	return SetValuable, nil
}

func (t SetType) TerraformType(ctx context.Context) tftypes.Type {
	return tftypes.Set{
		ElementType: t.ElementType().TerraformType(ctx),
	}
}

func (t SetType) ElementType() attr.Type {
	if t.SetType.ElemType == nil {
		return missingType{}
	}

	return t.SetType.ElemType
}

func (t SetType) ValueType(ctx context.Context) attr.Value {
	return SetValue{
		SetValue: t.SetType.ValueType(ctx).(basetypes.SetValue),
	}
}
