package tftestutil

import (
	"testing"

	"github.com/khulnasoft/tunnel/internal/testutil"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/terraform/parser"
	"github.com/khulnasoft/tunnel/pkg/iac/terraform"
)

func CreateModulesFromSource(t *testing.T, source, ext string) terraform.Modules {
	fs := testutil.CreateFS(t, map[string]string{
		"source" + ext: source,
	})
	p := parser.New(fs, "", parser.OptionStopOnHCLError(true))
	if err := p.ParseFS(t.Context(), "."); err != nil {
		t.Fatal(err)
	}
	modules, _, err := p.EvaluateAll(t.Context())
	if err != nil {
		t.Fatalf("parse error: %s", err)
	}
	return modules
}
