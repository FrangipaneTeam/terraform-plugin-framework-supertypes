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
var _ basetypes.SetTypable = SetNestedType{}

type SetNestedType struct {
	basetypes.SetType
}

func (t SetNestedType) String() string {
	return "types.SetType[" + t.ElementType().String() + "]"
}

func (t SetNestedType) ValueFromSet(_ context.Context, in basetypes.SetValue) (basetypes.SetValuable, diag.Diagnostics) {
	// CustomSetValue defined in the value type section
	value := SetNestedValue{
		SetValue: in,
	}

	return value, nil
}

func (t SetNestedType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.SetType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	SetValue, ok := attrValue.(basetypes.SetValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromSet(ctx, SetValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting SetValue to SetValuable: %v", diags)
	}

	return stringValuable, nil
}

// TerraformType returns the tftypes.Type that should be used to
// represent this type. This constrains what user input will be
// accepted and what kind of data can be set in state. The framework
// will use this to translate the AttributeType to something Terraform
// can understand.
func (t SetNestedType) TerraformType(ctx context.Context) tftypes.Type {
	return tftypes.Set{
		ElementType: t.SetType.ElementType().TerraformType(ctx),
	}
}

func (t SetNestedType) ElementType() attr.Type {
	if t.SetType.ElemType == nil {
		return missingType{}
	}

	return t.SetType.ElemType
}
