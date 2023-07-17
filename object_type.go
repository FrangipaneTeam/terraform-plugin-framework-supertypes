// -------------------------------------------------------------------- //
// !      DO NOT EDIT. This file is auto-generated from template      ! //
// -------------------------------------------------------------------- //
package supertypes

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.ObjectTypable = ObjectType{}

type ObjectType struct {
	basetypes.ObjectType
}

func (t ObjectType) Equal(o attr.Type) bool {
	other, ok := o.(ObjectType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t ObjectType) String() string {
	var res strings.Builder
	res.WriteString("types.ObjectType[")
	keys := make([]string, 0, len(t.AttrTypes))
	for k := range t.AttrTypes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for pos, key := range keys {
		if pos != 0 {
			res.WriteString(", ")
		}
		res.WriteString(`"` + key + `":`)
		res.WriteString(t.AttrTypes[key].String())
	}
	res.WriteString("]")
	return res.String()
}

func (t ObjectType) ValueFromObject(_ context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	value := ObjectValue{
		ObjectValue: in,
	}

	return value, nil
}

func (t ObjectType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ObjectType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	ObjectValue, ok := attrValue.(basetypes.ObjectValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	ObjectValuable, diags := t.ValueFromObject(ctx, ObjectValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ObjectValue to ObjectValuable: %v", diags)
	}

	return ObjectValuable, nil
}

func (t ObjectType) ValueType(_ context.Context) attr.Value {
	return ObjectValue{}
}
