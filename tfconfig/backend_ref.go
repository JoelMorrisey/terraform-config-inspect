package tfconfig

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"strings"
)

// S3BackendConfig represents a block of info for S3 that would be contained in a backend block in a terraform block configuration
type S3BackendConfig struct {
	Bucket        string `json:"bucket"`
	Key           string `json:"key"`
	Region        string `json:"region"`
	DynamodbTable string `json:"dynamodbTable,omitempty"`
}

// LocalBackendConfig represents a block of info for a local that would be contained in a backend block in a terraform block configuration
type LocalBackendConfig struct {
	Path string `json:"path"`
}

// BackendConfig represents a backend block in the terraform block configuration
type BackendConfig struct {
	Type  string              `json:"type"`
	S3    *S3BackendConfig    `json:"s3,omitempty"`
	Local *LocalBackendConfig `json:"local,omitempty"`
}

func DefaultBackendBlock(dir string) *BackendConfig {
	reqs := &BackendConfig{}
	reqs.Local = &LocalBackendConfig{}
	reqs.Type = "local"
	reqs.Local.Path = dir + "/terraform.tfstate"
	return reqs
}

func DecodeBackendBlock(block *hcl.Block) (*BackendConfig, hcl.Diagnostics) {
	reqs := &BackendConfig{}
	diags := hcl.Diagnostics{}
	reqs.Type = block.Labels[0]
	switch strings.ToLower(block.Labels[0]) {
	case "s3":
		diags = s3BackendDecode(block, reqs)
	case "local":
		diags = localBackendDecode(block, reqs)
	}
	return reqs, diags
}

func s3BackendDecode(block *hcl.Block, reqs *BackendConfig) hcl.Diagnostics {
	reqs.S3 = &S3BackendConfig{}
	content, _, diags := block.Body.PartialContent(backendConfigSchema)

	valDiags := gohcl.DecodeExpression(content.Attributes["bucket"].Expr, nil, &reqs.S3.Bucket)
	diags = append(diags, valDiags...)

	valDiags = gohcl.DecodeExpression(content.Attributes["key"].Expr, nil, &reqs.S3.Key)
	diags = append(diags, valDiags...)

	valDiags = gohcl.DecodeExpression(content.Attributes["region"].Expr, nil, &reqs.S3.Region)
	diags = append(diags, valDiags...)

	if content.Attributes["dynamodb_table"] != nil {
		valDiags = gohcl.DecodeExpression(content.Attributes["dynamodb_table"].Expr, nil, &reqs.S3.DynamodbTable)
		diags = append(diags, valDiags...)
	}

	return diags
}

func localBackendDecode(block *hcl.Block, reqs *BackendConfig) hcl.Diagnostics {
	reqs.Local = &LocalBackendConfig{}
	content, _, diags := block.Body.PartialContent(backendConfigSchema)

	valDiags := gohcl.DecodeExpression(content.Attributes["path"].Expr, nil, &reqs.Local.Path)
	diags = append(diags, valDiags...)
	//
	//valDiags = gohcl.DecodeExpression(content.Attributes["key"].Expr, nil, &reqs.S3.Key)
	//diags = append(diags, valDiags...)
	//
	//valDiags = gohcl.DecodeExpression(content.Attributes["region"].Expr, nil, &reqs.S3.Region)
	//diags = append(diags, valDiags...)

	return diags
}
