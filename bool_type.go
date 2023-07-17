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
var _ basetypes.BoolTypable = BoolType{}

type BoolType struct {
	basetypes.BoolType
}

func (t BoolType) Equal(o attr.Type) bool {
	other, ok := o.(BoolType)

	if !ok {
		return false
	}

	return t.BoolType.Equal(other.BoolType)
}

func (t BoolType) String() string {
	return "SuperTypesStringType"
}

func (t BoolType) ValueFromBool(_ context.Context, in basetypes.BoolValue) (basetypes.BoolValuable, diag.Diagnostics) {
	value := BoolValue{
		BoolValue: in,
	}

	return value, nil
}

func (t BoolType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.BoolType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	BoolValue, ok := attrValue.(basetypes.BoolValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	BoolValuable, diags := t.ValueFromBool(ctx, BoolValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting BoolValue to BoolValuable: %v", diags)
	}

	return BoolValuable, nil
}

func (t BoolType) ValueType(_ context.Context) attr.Value {
	return BoolValue{}
}
