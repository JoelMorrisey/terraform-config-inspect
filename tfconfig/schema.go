// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfconfig

import (
	"github.com/hashicorp/hcl/v2"
)

var rootSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "terraform",
			LabelNames: nil,
		},
		{
			Type:       "variable",
			LabelNames: []string{"name"},
		},
		{
			Type:       "output",
			LabelNames: []string{"name"},
		},
		{
			Type:       "provider",
			LabelNames: []string{"name"},
		},
		{
			Type:       "resource",
			LabelNames: []string{"type", "name"},
		},
		{
			Type:       "data",
			LabelNames: []string{"type", "name"},
		},
		{
			Type:       "module",
			LabelNames: []string{"name"},
		},
	},
}

var terraformBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "required_version",
		},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type: "required_providers",
		},
		{
			Type:       "backend",
			LabelNames: []string{"name"},
		},
	},
}

var providerConfigSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "version",
		},
		{
			Name: "alias",
		},
	},
}

var backendConfigSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "bucket",
		},
		{
			Name: "key",
		},
		{
			Name: "region",
		},
		{
			Name: "dynamodb_table",
		},
		{
			Name: "path",
		},
	},
}

var variableSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "type",
		},
		{
			Name: "description",
		},
		{
			Name: "default",
		},
		{
			Name: "sensitive",
		},
	},
}

var outputSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "description",
		},
		{
			Name: "sensitive",
		},
	},
}

var moduleCallSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "source",
		},
		{
			Name: "version",
		},
		{
			Name: "providers",
		},
	},
}

var resourceSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "provider",
		},
	},
}
