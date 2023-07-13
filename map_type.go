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
var _ basetypes.MapTypable = MapType{}

type MapType struct {
	basetypes.MapType
}

func (t MapType) Equal(o attr.Type) bool {
	other, ok := o.(basetypes.MapType)

	if !ok {
		return false
	}

	return t.MapType.Equal(other)
}

func (t MapType) String() string {
	return "types.MapType[" + t.ElementType().String() + "]"
}

func (t MapType) ValueFromMap(_ context.Context, in basetypes.MapValue) (basetypes.MapValuable, diag.Diagnostics) {
	value := MapValue{
		MapValue: in,
	}

	return value, nil
}

func (t MapType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.MapType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	MapValue, ok := attrValue.(basetypes.MapValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	MapValuable, diags := t.ValueFromMap(ctx, MapValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting MapValue to MapValuable: %v", diags)
	}

	return MapValuable, nil
}

func (t MapType) TerraformType(ctx context.Context) tftypes.Type {
	return tftypes.Map{
		ElementType: t.ElementType().TerraformType(ctx),
	}
}

func (t MapType) ElementType() attr.Type {
	if t.MapType.ElemType == nil {
		return missingType{}
	}

	return t.MapType.ElemType
}

func (t MapType) ValueType(ctx context.Context) attr.Value {
	return MapValue{
		MapValue: t.MapType.ValueType(ctx).(basetypes.MapValue),
	}
}
