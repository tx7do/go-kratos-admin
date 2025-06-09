package assets

import _ "embed"

//go:embed openapi.yaml
var OpenApiData []byte

//go:embed rbac.rego
var OpaRbacRego []byte
