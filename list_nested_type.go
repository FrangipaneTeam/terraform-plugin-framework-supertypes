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
var _ basetypes.ListTypable = ListNestedType{}

type ListNestedType struct {
	basetypes.ListType
}

// func (t ListNestedType) Equal(o attr.Type) bool {
// 	other, ok := o.(basetypes.ListType)

// 	if !ok {
// 		return false
// 	}

// 	return t.ListType.Equal(other)
// }

func (t ListNestedType) String() string {
	return "types.ListType[" + t.ElementType().String() + "]"
}

func (t ListNestedType) ValueFromList(_ context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	// CustomListValue defined in the value type section
	value := ListNestedValue{
		ListValue: in,
	}

	return value, nil
}

func (t ListNestedType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ListType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	ListValue, ok := attrValue.(basetypes.ListValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromList(ctx, ListValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ListValue to ListValuable: %v", diags)
	}

	return stringValuable, nil
}

// // TerraformType returns the tftypes.Type that should be used to
// // represent this type. This constrains what user input will be
// // accepted and what kind of data can be set in state. The framework
// // will use this to translate the AttributeType to something Terraform
// // can understand.
func (t ListNestedType) TerraformType(ctx context.Context) tftypes.Type {
	return tftypes.List{
		ElementType: t.ListType.ElementType().TerraformType(ctx),
	}
}

// GetType returns ListType of ObjectType or CustomType.
// func (a ListNestedType) GetType() attr.Type {
// 	if a.CustomType != nil {
// 		return a.CustomType
// 	}

// 	return types.ListType{
// 		ElemType: a.NestedObject.Type(),
// 	}
// }

func (t ListNestedType) ElementType() attr.Type {
	if t.ListType.ElemType == nil {
		return missingType{}
	}

	return t.ListType.ElemType
}

// func (t ListNestedType) ValueType(ctx context.Context) attr.Value {
// 	// CustomListValue defined in the value type section
// 	return ListValue{
// 		ListValue: t.ListType.ValueType(ctx).(basetypes.ListValue),
// 	}
// }
