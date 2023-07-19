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
var _ basetypes.ListTypable = ListType{}

type ListType struct {
	basetypes.ListType
}

func (t ListType) Equal(o attr.Type) bool {
	other, ok := o.(ListType)

	if !ok {
		return false
	}

	return t.ListType.Equal(other.ListType)
}

func (t ListType) String() string {
	return "supertypes.ListType[" + t.ElementType().String() + "]"
}

func (t ListType) ValueFromList(_ context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	value := ListValue{
		ListValue: in,
	}

	return value, nil
}

func (t ListType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ListType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	ListValue, ok := attrValue.(basetypes.ListValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	ListValuable, diags := t.ValueFromList(ctx, ListValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ListValue to ListValuable: %v", diags)
	}

	return ListValuable, nil
}

func (t ListType) TerraformType(ctx context.Context) tftypes.Type {
	return tftypes.List{
		ElementType: t.ElementType().TerraformType(ctx),
	}
}

func (t ListType) ElementType() attr.Type {
	if t.ListType.ElemType == nil {
		return missingType{}
	}

	return t.ListType.ElemType
}

func (t ListType) ValueType(ctx context.Context) attr.Value {
	return ListValue{
		ListValue: t.ListType.ValueType(ctx).(basetypes.ListValue),
	}
}
