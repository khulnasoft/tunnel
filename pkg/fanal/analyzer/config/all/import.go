package all

import (
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/config/azurearm"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/config/cloudformation"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/config/dockerfile"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/config/helm"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/config/json"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/config/k8s"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/config/terraform"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/config/terraformplan/json"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/config/terraformplan/snapshot"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/config/yaml"
)
