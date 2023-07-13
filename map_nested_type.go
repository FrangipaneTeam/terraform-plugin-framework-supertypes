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
var _ basetypes.MapTypable = MapNestedType{}

type MapNestedType struct {
	basetypes.MapType
}

// func (t MapNestedType) Equal(o attr.Type) bool {
// 	other, ok := o.(basetypes.MapType)

// 	if !ok {
// 		return false
// 	}

// 	return t.MapType.Equal(other)
// }

func (t MapNestedType) String() string {
	return "types.MapType[" + t.ElementType().String() + "]"
}

func (t MapNestedType) ValueFromMap(_ context.Context, in basetypes.MapValue) (basetypes.MapValuable, diag.Diagnostics) {
	// CustomMapValue defined in the value type section
	value := MapNestedValue{
		MapValue: in,
	}

	return value, nil
}

func (t MapNestedType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.MapType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	MapValue, ok := attrValue.(basetypes.MapValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromMap(ctx, MapValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting MapValue to MapValuable: %v", diags)
	}

	return stringValuable, nil
}

// // TerraformType returns the tftypes.Type that should be used to
// // represent this type. This constrains what user input will be
// // accepted and what kind of data can be set in state. The framework
// // will use this to translate the AttributeType to something Terraform
// // can understand.
func (t MapNestedType) TerraformType(ctx context.Context) tftypes.Type {
	return tftypes.Map{
		ElementType: t.MapType.ElementType().TerraformType(ctx),
	}
}

// GetType returns MapType of ObjectType or CustomType.
// func (a MapNestedType) GetType() attr.Type {
// 	if a.CustomType != nil {
// 		return a.CustomType
// 	}

// 	return types.MapType{
// 		ElemType: a.NestedObject.Type(),
// 	}
// }

func (t MapNestedType) ElementType() attr.Type {
	if t.MapType.ElemType == nil {
		return missingType{}
	}

	return t.MapType.ElemType
}

// func (t MapNestedType) ValueType(ctx context.Context) attr.Value {
// 	// CustomMapValue defined in the value type section
// 	return MapValue{
// 		MapValue: t.MapType.ValueType(ctx).(basetypes.MapValue),
// 	}
// }
