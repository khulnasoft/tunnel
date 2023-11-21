package commands

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/khulnasoft-lab/tunnel-kubernetes/pkg/k8s"
	"github.com/khulnasoft-lab/tunnel-kubernetes/pkg/tunnelk8s"
	"github.com/khulnasoft/tunnel/pkg/flag"
	"github.com/khulnasoft/tunnel/pkg/log"
)

// namespaceRun runs scan on kubernetes cluster
func namespaceRun(ctx context.Context, opts flag.Options, cluster k8s.Cluster) error {
	if err := validateReportArguments(opts); err != nil {
		return err
	}
	var tunnelk tunnelk8s.TunnelK8S
	if opts.AllNamespaces {
		tunnelk = tunnelk8s.New(cluster, log.Logger).AllNamespaces()
	} else {
		tunnelk = tunnelk8s.New(cluster, log.Logger).Namespace(getNamespace(opts, cluster.GetCurrentNamespace()))
	}

	artifacts, err := tunnelk.ListArtifacts(ctx)
	if err != nil {
		return xerrors.Errorf("get k8s artifacts error: %w", err)
	}

	runner := newRunner(opts, cluster.GetCurrentContext())
	return runner.run(ctx, artifacts)
}

func getNamespace(opts flag.Options, currentNamespace string) string {
	if len(opts.K8sOptions.Namespace) > 0 {
		return opts.K8sOptions.Namespace
	}

	return currentNamespace
}
