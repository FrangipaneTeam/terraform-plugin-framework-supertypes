package supertypes_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	supertypes "github.com/FrangipaneTeam/terraform-plugin-framework-supertypes"
)

func TestAttributeTypes(t *testing.T) {
	t.Parallel()

	type struct1 struct{}
	type struct2 struct {
		Name            types.String `tfsdk:"name"`
		ID              types.Int64  `tfsdk:"id"`
		IncludeProperty types.Bool   `tfsdk:"include_property"`
	}

	ctx := context.Background()
	got := supertypes.AttributeTypesMust[struct1](ctx)
	wanted := map[string]attr.Type{}

	if diff := cmp.Diff(got, wanted); diff != "" {
		t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}

	_, err := supertypes.AttributeTypes[int](ctx)

	if err == nil {
		t.Fatalf("expected error")
	}

	got, err = supertypes.AttributeTypes[struct2](ctx)

	if err != nil {
		t.Fatalf("unexpected error")
	}

	wanted = map[string]attr.Type{
		"name":             types.StringType,
		"id":               types.Int64Type,
		"include_property": types.BoolType,
	}

	if diff := cmp.Diff(got, wanted); diff != "" {
		t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}
}
