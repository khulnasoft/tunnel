package server

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/khulnasoft/tunnel/pkg/cache"
	"github.com/khulnasoft/tunnel/pkg/commands/operation"
	"github.com/khulnasoft/tunnel/pkg/db"
	"github.com/khulnasoft/tunnel/pkg/flag"
	"github.com/khulnasoft/tunnel/pkg/log"
	"github.com/khulnasoft/tunnel/pkg/module"
	rpcServer "github.com/khulnasoft/tunnel/pkg/rpc/server"
)

// Run runs the scan
func Run(ctx context.Context, opts flag.Options) (err error) {
	log.InitLogger(opts.Debug, opts.Quiet)

	// configure cache dir
	cacheClient, cleanup, err := cache.New(opts.CacheOpts())
	if err != nil {
		return xerrors.Errorf("server cache error: %w", err)
	}
	defer cleanup()

	// download the database file
	if err = operation.DownloadDB(ctx, opts.AppVersion, opts.CacheDir, opts.DBRepositories,
		true, opts.SkipDBUpdate, opts.RegistryOpts()); err != nil {
		return err
	}

	if opts.DownloadDBOnly {
		return nil
	}

	if err = db.Init(db.Dir(opts.CacheDir)); err != nil {
		return xerrors.Errorf("error in vulnerability DB initialize: %w", err)
	}

	// Initialize WASM modules
	m, err := module.NewManager(ctx, module.Options{
		Dir:            opts.ModuleDir,
		EnabledModules: opts.EnabledModules,
	})
	if err != nil {
		return xerrors.Errorf("WASM module error: %w", err)
	}
	m.Register()

	server := rpcServer.NewServer(opts.AppVersion, opts.Listen, opts.CacheDir, opts.Token, opts.TokenHeader,
		opts.PathPrefix, opts.DBRepositories, opts.RegistryOpts())
	return server.ListenAndServe(ctx, cacheClient, opts.SkipDBUpdate)
}
