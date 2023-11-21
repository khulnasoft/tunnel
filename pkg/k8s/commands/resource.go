package commands

import (
	"context"
	"strings"

	"golang.org/x/xerrors"

	"github.com/khulnasoft-lab/tunnel-kubernetes/pkg/artifacts"
	"github.com/khulnasoft-lab/tunnel-kubernetes/pkg/k8s"
	"github.com/khulnasoft-lab/tunnel-kubernetes/pkg/tunnelk8s"
	"github.com/khulnasoft/tunnel/pkg/flag"
	"github.com/khulnasoft/tunnel/pkg/log"
)

// resourceRun runs scan on kubernetes cluster
func resourceRun(ctx context.Context, args []string, opts flag.Options, cluster k8s.Cluster) error {
	kind, name, err := extractKindAndName(args)
	if err != nil {
		return err
	}

	runner := newRunner(opts, cluster.GetCurrentContext())

	var tunnelk tunnelk8s.TunnelK8S

	tunnelk = tunnelk8s.New(cluster, log.Logger, tunnelk8s.WithExcludeOwned(opts.ExcludeOwned))

	if opts.AllNamespaces {
		tunnelk = tunnelk.AllNamespaces()
	} else {
		tunnelk = tunnelk.Namespace(getNamespace(opts, cluster.GetCurrentNamespace()))
	}

	if name == "" { // pods or configmaps etc
		if err = validateReportArguments(opts); err != nil {
			return err
		}

		targets, err := tunnelk.Resources(kind).ListArtifacts(ctx)
		if err != nil {
			return err
		}

		return runner.run(ctx, targets)
	}

	// pod/NAME or pod NAME etc
	artifact, err := tunnelk.GetArtifact(ctx, kind, name)
	if err != nil {
		return err
	}

	return runner.run(ctx, []*artifacts.Artifact{artifact})
}

func extractKindAndName(args []string) (string, string, error) {
	switch len(args) {
	case 1:
		s := strings.Split(args[0], "/")
		if len(s) != 2 {
			return args[0], "", nil
		}

		return s[0], s[1], nil
	case 2:
		return args[0], args[1], nil
	}

	return "", "", xerrors.Errorf("can't parse arguments %v. Please run `tunnel k8s` for usage.", args)
}
