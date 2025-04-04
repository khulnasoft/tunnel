package sum

import ftypes "github.com/khulnasoft/tunnel/pkg/fanal/types"

var (
	// docker run --name gomod --rm -it golang:1.15 bash
	// export USER=gomod
	// mkdir repo
	// cd repo
	// go mod init github.com/org/repo
	// go get golang.org/x/xerrors
	// go list -m all | awk 'NR>1 {sub(/^v/, "", $2); printf("{\""$1"\", \""$2"\", },\n")}'
	GoModNormal = []ftypes.Package{
		{Name: "golang.org/x/xerrors", Version: "v0.0.0-20200804184101-5ec99f83aff1"},
	}

	// https://github.com/uudashr/gopkgs/blob/616744904701ef01d868da4b66aad0e6856c361d/v2/go.sum
	GoModEmptyLine = []ftypes.Package{
		{Name: "github.com/karrick/godirwalk", Version: "v1.12.0"},
		{Name: "github.com/pkg/errors", Version: "v0.8.1"},
	}

	// docker run --name gomod --rm -it golang:1.15 bash
	// export USER=gomod
	// mkdir repo
	// cd repo
	// go mod init github.com/org/repo
	// go get golang.org/x/xerrors
	// go get github.com/urfave/cli
	// go get github.com/stretchr/testify
	// go get github.com/BurntSushi/toml
	// go list -m all | awk 'NR>1 {sub(/^v/, "", $2); printf("{\""$1"\", \""$2"\", },\n")}'
	GoModMany = []ftypes.Package{
		{Name: "github.com/BurntSushi/toml", Version: "v0.3.1"},
		{Name: "github.com/cpuguy83/go-md2man/v2", Version: "v2.0.0-20190314233015-f79a8a8ca69d"},
		{Name: "github.com/davecgh/go-spew", Version: "v1.1.0"},
		{Name: "github.com/pmezard/go-difflib", Version: "v1.0.0"},
		{Name: "github.com/russross/blackfriday/v2", Version: "v2.0.1"},
		{Name: "github.com/shurcooL/sanitized_anchor_name", Version: "v1.0.0"},
		{Name: "github.com/stretchr/objx", Version: "v0.1.0"},
		{Name: "github.com/stretchr/testify", Version: "v1.7.0"},
		{Name: "github.com/urfave/cli", Version: "v1.22.5"},
		{Name: "golang.org/x/xerrors", Version: "v0.0.0-20200804184101-5ec99f83aff1"},
		{Name: "gopkg.in/check.v1", Version: "v0.0.0-20161208181325-20d25e280405"},
		{Name: "gopkg.in/yaml.v2", Version: "v2.2.2"},
		{Name: "gopkg.in/yaml.v3", Version: "v3.0.0-20200313102051-9f266ea9e77c"},
	}

	// docker run --name gomod --rm -it golang:1.15 bash
	// export USER=gomod
	// mkdir repo
	// cd repo
	// go mod init github.com/org/repo
	// go get github.com/khulnasoft/tunnel
	// go list -m all | awk 'NR>1 {sub(/^v/, "", $2); printf("{\""$1"\", \""$2"\", },\n")}'
	GoModTunnel = []ftypes.Package{
		{Name: "cloud.google.com/go", Version: "v0.65.0"},
		{Name: "cloud.google.com/go/bigquery", Version: "v1.8.0"},
		{Name: "cloud.google.com/go/datastore", Version: "v1.1.0"},
		{Name: "cloud.google.com/go/pubsub", Version: "v1.3.1"},
		{Name: "cloud.google.com/go/storage", Version: "v1.10.0"},
		{Name: "dmitri.shuralyov.com/gpu/mtl", Version: "v0.0.0-20190408044501-666a987793e9"},
		{Name: "github.com/Azure/azure-sdk-for-go", Version: "v38.0.0+incompatible"},
		{Name: "github.com/Azure/go-ansiterm", Version: "v0.0.0-20170929234023-d6e3b3328b78"},
		{Name: "github.com/Azure/go-autorest/autorest", Version: "v0.9.3"},
		{Name: "github.com/Azure/go-autorest/autorest/adal", Version: "v0.8.1"},
		{Name: "github.com/Azure/go-autorest/autorest/date", Version: "v0.2.0"},
		{Name: "github.com/Azure/go-autorest/autorest/mocks", Version: "v0.3.0"},
		{Name: "github.com/Azure/go-autorest/autorest/to", Version: "v0.3.0"},
		{Name: "github.com/Azure/go-autorest/autorest/validation", Version: "v0.1.0"},
		{Name: "github.com/Azure/go-autorest/logger", Version: "v0.1.0"},
		{Name: "github.com/Azure/go-autorest/tracing", Version: "v0.5.0"},
		{Name: "github.com/BurntSushi/toml", Version: "v0.3.1"},
		{Name: "github.com/BurntSushi/xgb", Version: "v0.0.0-20160522181843-27f122750802"},
		{Name: "github.com/GoogleCloudPlatform/docker-credential-gcr", Version: "v1.5.0"},
		{Name: "github.com/GoogleCloudPlatform/k8s-cloud-provider", Version: "v0.0.0-20190822182118-27a4ced34534"},
		{Name: "github.com/Microsoft/go-winio", Version: "v0.4.15-0.20190919025122-fc70bd9a86b5"},
		{Name: "github.com/Microsoft/hcsshim", Version: "v0.8.6"},
		{Name: "github.com/NYTimes/gziphandler", Version: "v0.0.0-20170623195520-56545f4a5d46"},
		{Name: "github.com/OneOfOne/xxhash", Version: "v1.2.7"},
		{Name: "github.com/PuerkitoBio/purell", Version: "v1.1.1"},
		{Name: "github.com/PuerkitoBio/urlesc", Version: "v0.0.0-20170810143723-de5bf2ad4578"},
		{Name: "github.com/VividCortex/ewma", Version: "v1.1.1"},
		{Name: "github.com/alcortesm/tgz", Version: "v0.0.0-20161220082320-9c5fe88206d7"},
		{Name: "github.com/alecthomas/template", Version: "v0.0.0-20160405071501-a0175ee3bccc"},
		{Name: "github.com/alecthomas/units", Version: "v0.0.0-20151022065526-2efee857e7cf"},
		{Name: "github.com/alicebob/gopher-json", Version: "v0.0.0-20200520072559-a9ecdc9d1d3a"},
		{Name: "github.com/alicebob/miniredis/v2", Version: "v2.14.1"},
		{Name: "github.com/anmitsu/go-shlex", Version: "v0.0.0-20161002113705-648efa622239"},
		{Name: "github.com/khulnasoft-lab/bolt-fixtures", Version: "v0.0.0-20200903104109-d34e7f983986"},
		{Name: "github.com/khulnasoft/fanal", Version: "v0.0.0-20210119051230-28c249da7cfd"},
		{Name: "github.com/khulnasoft/go-dep-parser", Version: "v0.0.0-20201028043324-889d4a92b8e0"},
		{Name: "github.com/khulnasoft-lab/go-gem-version", Version: "v0.0.0-20201115065557-8eed6fe000ce"},
		{Name: "github.com/khulnasoft-lab/go-npm-version", Version: "v0.0.0-20201110091526-0b796d180798"},
		{Name: "github.com/khulnasoft-lab/go-pep440-version", Version: "v0.0.0-20210121094942-22b2f8951d46"},
		{Name: "github.com/khulnasoft/goversify", Version: "v0.0.0-20210121072130-637058cfe492"},
		{Name: "github.com/khulnasoft-lab/testdocker", Version: "v0.0.0-20210106133225-0b17fe083674"},
		{Name: "github.com/khulnasoft/tunnel", Version: "v0.16.0"},
		{Name: "github.com/khulnasoft-lab/tunnel-db", Version: "v0.0.0-20210105160501-c5bf4e153277"},
		{Name: "github.com/khulnasoft/vuln-list-update", Version: "v0.0.0-20191016075347-3d158c2bf9a2"},
		{Name: "github.com/araddon/dateparse", Version: "v0.0.0-20190426192744-0d74ffceef83"},
		{Name: "github.com/armon/consul-api", Version: "v0.0.0-20180202201655-eb2c6b5be1b6"},
		{Name: "github.com/armon/go-socks5", Version: "v0.0.0-20160902184237-e75332964ef5"},
		{Name: "github.com/aws/aws-sdk-go", Version: "v1.27.1"},
		{Name: "github.com/beorn7/perks", Version: "v1.0.0"},
		{Name: "github.com/bgentry/speakeasy", Version: "v0.1.0"},
		{Name: "github.com/blang/semver", Version: "v3.5.0+incompatible"},
		{Name: "github.com/briandowns/spinner", Version: "v1.12.0"},
		{Name: "github.com/caarlos0/env/v6", Version: "v6.0.0"},
		{Name: "github.com/cenkalti/backoff", Version: "v2.2.1+incompatible"},
		{Name: "github.com/census-instrumentation/opencensus-proto", Version: "v0.2.1"},
		{Name: "github.com/cespare/xxhash/v2", Version: "v2.1.1"},
		{Name: "github.com/cheggaaa/pb/v3", Version: "v3.0.3"},
		{Name: "github.com/chzyer/logex", Version: "v1.1.10"},
		{Name: "github.com/chzyer/readline", Version: "v0.0.0-20180603132655-2972be24d48e"},
		{Name: "github.com/chzyer/test", Version: "v0.0.0-20180213035817-a1ea475d72b1"},
		{Name: "github.com/client9/misspell", Version: "v0.3.4"},
		{Name: "github.com/cncf/udpa/go", Version: "v0.0.0-20191209042840-269d4d468f6f"},
		{Name: "github.com/cockroachdb/datadriven", Version: "v0.0.0-20190809214429-80d97fb3cbaa"},
		{Name: "github.com/containerd/containerd", Version: "v1.3.3"},
		{Name: "github.com/containerd/continuity", Version: "v0.0.0-20190426062206-aaeac12a7ffc"},
		{Name: "github.com/coreos/etcd", Version: "v3.3.10+incompatible"},
		{Name: "github.com/coreos/go-etcd", Version: "v2.0.0+incompatible"},
		{Name: "github.com/coreos/go-oidc", Version: "v2.1.0+incompatible"},
		{Name: "github.com/coreos/go-semver", Version: "v0.3.0"},
		{Name: "github.com/coreos/go-systemd", Version: "v0.0.0-20190321100706-95778dfbb74e"},
		{Name: "github.com/coreos/pkg", Version: "v0.0.0-20180108230652-97fdf19511ea"},
		{Name: "github.com/cpuguy83/go-md2man", Version: "v1.0.10"},
		{Name: "github.com/cpuguy83/go-md2man/v2", Version: "v2.0.0"},
		{Name: "github.com/creack/pty", Version: "v1.1.9"},
		{Name: "github.com/davecgh/go-spew", Version: "v1.1.1"},
		{Name: "github.com/deckarep/golang-set", Version: "v1.7.1"},
		{Name: "github.com/dgrijalva/jwt-go", Version: "v3.2.0+incompatible"},
		{Name: "github.com/dgryski/go-rendezvous", Version: "v0.0.0-20200823014737-9f7001d12a5f"},
		{Name: "github.com/dnaeon/go-vcr", Version: "v1.0.1"},
		{Name: "github.com/docker/cli", Version: "v0.0.0-20191017083524-a8ff7f821017"},
		{Name: "github.com/docker/distribution", Version: "v2.7.1+incompatible"},
		{Name: "github.com/docker/docker", Version: "v1.4.2-0.20190924003213-a8608b5b67c7"},
		{Name: "github.com/docker/docker-credential-helpers", Version: "v0.6.3"},
		{Name: "github.com/docker/go-connections", Version: "v0.4.0"},
		{Name: "github.com/docker/go-units", Version: "v0.4.0"},
		{Name: "github.com/docker/spdystream", Version: "v0.0.0-20160310174837-449fdfce4d96"},
		{Name: "github.com/dustin/go-humanize", Version: "v1.0.0"},
		{Name: "github.com/elazarl/goproxy", Version: "v0.0.0-20200809112317-0581fc3aee2d"},
		{Name: "github.com/elazarl/goproxy/ext", Version: "v0.0.0-20200809112317-0581fc3aee2d"},
		{Name: "github.com/emicklei/go-restful", Version: "v2.9.5+incompatible"},
		{Name: "github.com/emirpasic/gods", Version: "v1.12.0"},
		{Name: "github.com/envoyproxy/go-control-plane", Version: "v0.9.4"},
		{Name: "github.com/envoyproxy/protoc-gen-validate", Version: "v0.1.0"},
		{Name: "github.com/evanphx/json-patch", Version: "v4.2.0+incompatible"},
		{Name: "github.com/fatih/color", Version: "v1.10.0"},
		{Name: "github.com/flynn/go-shlex", Version: "v0.0.0-20150515145356-3f9db97f8568"},
		{Name: "github.com/fsnotify/fsnotify", Version: "v1.4.9"},
		{Name: "github.com/ghodss/yaml", Version: "v1.0.0"},
		{Name: "github.com/gin-contrib/sse", Version: "v0.1.0"},
		{Name: "github.com/gin-gonic/gin", Version: "v1.5.0"},
		{Name: "github.com/gliderlabs/ssh", Version: "v0.2.2"},
		{Name: "github.com/go-git/gcfg", Version: "v1.5.0"},
		{Name: "github.com/go-git/go-billy/v5", Version: "v5.0.0"},
		{Name: "github.com/go-git/go-git-fixtures/v4", Version: "v4.0.1"},
		{Name: "github.com/go-git/go-git/v5", Version: "v5.0.0"},
		{Name: "github.com/go-gl/glfw", Version: "v0.0.0-20190409004039-e6da0acd62b1"},
		{Name: "github.com/go-gl/glfw/v3.3/glfw", Version: "v0.0.0-20200222043503-6f7a984d4dc4"},
		{Name: "github.com/go-kit/kit", Version: "v0.8.0"},
		{Name: "github.com/go-logfmt/logfmt", Version: "v0.3.0"},
		{Name: "github.com/go-logr/logr", Version: "v0.1.0"},
		{Name: "github.com/go-openapi/jsonpointer", Version: "v0.19.3"},
		{Name: "github.com/go-openapi/jsonreference", Version: "v0.19.3"},
		{Name: "github.com/go-openapi/spec", Version: "v0.19.3"},
		{Name: "github.com/go-openapi/swag", Version: "v0.19.5"},
		{Name: "github.com/go-playground/locales", Version: "v0.13.0"},
		{Name: "github.com/go-playground/universal-translator", Version: "v0.17.0"},
		{Name: "github.com/go-redis/redis", Version: "v6.15.7+incompatible"},
		{Name: "github.com/go-redis/redis/v8", Version: "v8.4.0"},
		{Name: "github.com/go-restruct/restruct", Version: "v0.0.0-20191227155143-5734170a48a1"},
		{Name: "github.com/go-sql-driver/mysql", Version: "v1.5.0"},
		{Name: "github.com/go-stack/stack", Version: "v1.8.0"},
		{Name: "github.com/gobwas/glob", Version: "v0.2.3"},
		{Name: "github.com/goccy/go-yaml", Version: "v1.8.2"},
		{Name: "github.com/gogo/protobuf", Version: "v1.3.1"},
		{Name: "github.com/golang/glog", Version: "v0.0.0-20160126235308-23def4e6c14b"},
		{Name: "github.com/golang/groupcache", Version: "v0.0.0-20200121045136-8c9f03a8e57e"},
		{Name: "github.com/golang/mock", Version: "v1.4.4"},
		{Name: "github.com/golang/protobuf", Version: "v1.4.2"},
		{Name: "github.com/google/btree", Version: "v1.0.0"},
		{Name: "github.com/google/go-cmp", Version: "v0.5.3"},
		{Name: "github.com/google/go-containerregistry", Version: "v0.0.0-20200331213917-3d03ed9b1ca2"},
		{Name: "github.com/google/go-github/v28", Version: "v28.1.1"},
		{Name: "github.com/google/go-querystring", Version: "v1.0.0"},
		{Name: "github.com/google/gofuzz", Version: "v1.0.0"},
		{Name: "github.com/google/martian", Version: "v2.1.0+incompatible"},
		{Name: "github.com/google/martian/v3", Version: "v3.0.0"},
		{Name: "github.com/google/pprof", Version: "v0.0.0-20200708004538-1a94d8640e99"},
		{Name: "github.com/google/renameio", Version: "v0.1.0"},
		{Name: "github.com/google/subcommands", Version: "v1.0.1"},
		{Name: "github.com/google/uuid", Version: "v1.1.1"},
		{Name: "github.com/google/wire", Version: "v0.3.0"},
		{Name: "github.com/googleapis/gax-go/v2", Version: "v2.0.5"},
		{Name: "github.com/googleapis/gnostic", Version: "v0.2.2"},
		{Name: "github.com/gophercloud/gophercloud", Version: "v0.1.0"},
		{Name: "github.com/gopherjs/gopherjs", Version: "v0.0.0-20200217142428-fce0ec30dd00"},
		{Name: "github.com/gorilla/context", Version: "v1.1.1"},
		{Name: "github.com/gorilla/mux", Version: "v1.7.4"},
		{Name: "github.com/gorilla/websocket", Version: "v1.4.0"},
		{Name: "github.com/gregjones/httpcache", Version: "v0.0.0-20180305231024-9cad4c3443a7"},
		{Name: "github.com/grpc-ecosystem/go-grpc-middleware", Version: "v1.0.1-0.20190118093823-f849b5445de4"},
		{Name: "github.com/grpc-ecosystem/go-grpc-prometheus", Version: "v1.2.0"},
		{Name: "github.com/grpc-ecosystem/grpc-gateway", Version: "v1.9.5"},
		{Name: "github.com/hashicorp/errwrap", Version: "v1.0.0"},
		{Name: "github.com/hashicorp/go-multierror", Version: "v1.1.0"},
		{Name: "github.com/hashicorp/go-version", Version: "v1.2.1"},
		{Name: "github.com/hashicorp/golang-lru", Version: "v0.5.3"},
		{Name: "github.com/hashicorp/hcl", Version: "v1.0.0"},
		{Name: "github.com/hpcloud/tail", Version: "v1.0.0"},
		{Name: "github.com/ianlancetaylor/demangle", Version: "v0.0.0-20181102032728-5e5cf60278f6"},
		{Name: "github.com/imdario/mergo", Version: "v0.3.5"},
		{Name: "github.com/inconshreveable/mousetrap", Version: "v1.0.0"},
		{Name: "github.com/jbenet/go-context", Version: "v0.0.0-20150711004518-d14ea06fba99"},
		{Name: "github.com/jessevdk/go-flags", Version: "v1.4.0"},
		{Name: "github.com/jmespath/go-jmespath", Version: "v0.0.0-20180206201540-c2b33e8439af"},
		{Name: "github.com/joefitzgerald/rainbow-reporter", Version: "v0.1.0"},
		{Name: "github.com/jonboulle/clockwork", Version: "v0.1.0"},
		{Name: "github.com/json-iterator/go", Version: "v1.1.8"},
		{Name: "github.com/jstemmer/go-junit-report", Version: "v0.9.1"},
		{Name: "github.com/jtolds/gls", Version: "v4.20.0+incompatible"},
		{Name: "github.com/julienschmidt/httprouter", Version: "v1.2.0"},
		{Name: "github.com/kevinburke/ssh_config", Version: "v0.0.0-20190725054713-01f96b0aa0cd"},
		{Name: "github.com/kisielk/errcheck", Version: "v1.2.0"},
		{Name: "github.com/kisielk/gotool", Version: "v1.0.0"},
		{Name: "github.com/knqyf263/go-apk-version", Version: "v0.0.0-20200609155635-041fdbb8563f"},
		{Name: "github.com/knqyf263/go-deb-version", Version: "v0.0.0-20190517075300-09fca494f03d"},
		{Name: "github.com/knqyf263/go-rpm-version", Version: "v0.0.0-20170716094938-74609b86c936"},
		{Name: "github.com/knqyf263/go-rpmdb", Version: "v0.0.0-20201215100354-a9e3110d8ee1"},
		{Name: "github.com/knqyf263/nested", Version: "v0.0.1"},
		{Name: "github.com/konsorten/go-windows-terminal-sequences", Version: "v1.0.2"},
		{Name: "github.com/kr/logfmt", Version: "v0.0.0-20140226030751-b84e30acd515"},
		{Name: "github.com/kr/pretty", Version: "v0.1.0"},
		{Name: "github.com/kr/pty", Version: "v1.1.5"},
		{Name: "github.com/kr/text", Version: "v0.2.0"},
		{Name: "github.com/kylelemons/godebug", Version: "v1.1.0"},
		{Name: "github.com/leodido/go-urn", Version: "v1.2.0"},
		{Name: "github.com/magiconair/properties", Version: "v1.8.0"},
		{Name: "github.com/mailru/easyjson", Version: "v0.7.0"},
		{Name: "github.com/mattn/go-colorable", Version: "v0.1.8"},
		{Name: "github.com/mattn/go-isatty", Version: "v0.0.12"},
		{Name: "github.com/mattn/go-jsonpointer", Version: "v0.0.0-20180225143300-37667080efed"},
		{Name: "github.com/mattn/go-runewidth", Version: "v0.0.9"},
		{Name: "github.com/matttproud/golang_protobuf_extensions", Version: "v1.0.1"},
		{Name: "github.com/maxbrunsfeld/counterfeiter/v6", Version: "v6.2.2"},
		{Name: "github.com/mitchellh/go-homedir", Version: "v1.1.0"},
		{Name: "github.com/mitchellh/mapstructure", Version: "v1.1.2"},
		{Name: "github.com/modern-go/concurrent", Version: "v0.0.0-20180306012644-bacd9c7ef1dd"},
		{Name: "github.com/modern-go/reflect2", Version: "v1.0.1"},
		{Name: "github.com/morikuni/aec", Version: "v1.0.0"},
		{Name: "github.com/munnerz/goautoneg", Version: "v0.0.0-20191010083416-a7dc8b61c822"},
		{Name: "github.com/mwitkow/go-conntrack", Version: "v0.0.0-20161129095857-cc309e4a2223"},
		{Name: "github.com/mxk/go-flowrate", Version: "v0.0.0-20140419014527-cca7078d478f"},
		{Name: "github.com/niemeyer/pretty", Version: "v0.0.0-20200227124842-a10e7caefd8e"},
		{Name: "github.com/nxadm/tail", Version: "v1.4.4"},
		{Name: "github.com/olekukonko/tablewriter", Version: "v0.0.2-0.20190607075207-195002e6e56a"},
		{Name: "github.com/onsi/ginkgo", Version: "v1.14.2"},
		{Name: "github.com/onsi/gomega", Version: "v1.10.3"},
		{Name: "github.com/open-policy-agent/opa", Version: "v0.21.1"},
		{Name: "github.com/opencontainers/go-digest", Version: "v1.0.0-rc1"},
		{Name: "github.com/opencontainers/image-spec", Version: "v1.0.2-0.20190823105129-775207bd45b6"},
		{Name: "github.com/opencontainers/runc", Version: "v0.1.1"},
		{Name: "github.com/parnurzeal/gorequest", Version: "v0.2.16"},
		{Name: "github.com/pelletier/go-toml", Version: "v1.2.0"},
		{Name: "github.com/peterbourgon/diskv", Version: "v2.0.1+incompatible"},
		{Name: "github.com/peterh/liner", Version: "v0.0.0-20170211195444-bf27d3ba8e1d"},
		{Name: "github.com/pkg/errors", Version: "v0.9.1"},
		{Name: "github.com/pmezard/go-difflib", Version: "v1.0.0"},
		{Name: "github.com/pquerna/cachecontrol", Version: "v0.0.0-20171018203845-0dec1b30a021"},
		{Name: "github.com/prometheus/client_golang", Version: "v1.0.0"},
		{Name: "github.com/prometheus/client_model", Version: "v0.0.0-20190812154241-14fe0d1b01d4"},
		{Name: "github.com/prometheus/common", Version: "v0.4.1"},
		{Name: "github.com/prometheus/procfs", Version: "v0.0.2"},
		{Name: "github.com/rcrowley/go-metrics", Version: "v0.0.0-20181016184325-3113b8401b8a"},
		{Name: "github.com/remyoudompheng/bigfft", Version: "v0.0.0-20170806203942-52369c62f446"},
		{Name: "github.com/rogpeppe/fastuuid", Version: "v0.0.0-20150106093220-6724a57986af"},
		{Name: "github.com/rogpeppe/go-charset", Version: "v0.0.0-20180617210344-2471d30d28b4"},
		{Name: "github.com/rogpeppe/go-internal", Version: "v1.3.0"},
		{Name: "github.com/rubiojr/go-vhd", Version: "v0.0.0-20160810183302-0bfd3b39853c"},
		{Name: "github.com/russross/blackfriday", Version: "v1.5.2"},
		{Name: "github.com/russross/blackfriday/v2", Version: "v2.0.1"},
		{Name: "github.com/saracen/walker", Version: "v0.0.0-20191201085201-324a081bae7e"},
		{Name: "github.com/satori/go.uuid", Version: "v1.2.0"},
		{Name: "github.com/sclevine/spec", Version: "v1.2.0"},
		{Name: "github.com/sergi/go-diff", Version: "v1.1.0"},
		{Name: "github.com/shurcooL/sanitized_anchor_name", Version: "v1.0.0"},
		{Name: "github.com/simplereach/timeutils", Version: "v1.2.0"},
		{Name: "github.com/sirupsen/logrus", Version: "v1.5.0"},
		{Name: "github.com/smartystreets/assertions", Version: "v1.2.0"},
		{Name: "github.com/smartystreets/goconvey", Version: "v1.6.4"},
		{Name: "github.com/soheilhy/cmux", Version: "v0.1.4"},
		{Name: "github.com/sosedoff/gitkit", Version: "v0.2.0"},
		{Name: "github.com/spf13/afero", Version: "v1.2.2"},
		{Name: "github.com/spf13/cast", Version: "v1.3.0"},
		{Name: "github.com/spf13/cobra", Version: "v0.0.5"},
		{Name: "github.com/spf13/jwalterweatherman", Version: "v1.0.0"},
		{Name: "github.com/spf13/pflag", Version: "v1.0.5"},
		{Name: "github.com/spf13/viper", Version: "v1.3.2"},
		{Name: "github.com/stretchr/objx", Version: "v0.3.0"},
		{Name: "github.com/stretchr/testify", Version: "v1.6.1"},
		{Name: "github.com/testcontainers/testcontainers-go", Version: "v0.3.1"},
		{Name: "github.com/tmc/grpc-websocket-proxy", Version: "v0.0.0-20170815181823-89b8d40f7ca8"},
		{Name: "github.com/twitchtv/twirp", Version: "v5.10.1+incompatible"},
		{Name: "github.com/ugorji/go", Version: "v1.1.7"},
		{Name: "github.com/ugorji/go/codec", Version: "v1.1.7"},
		{Name: "github.com/urfave/cli", Version: "v1.22.5"},
		{Name: "github.com/urfave/cli/v2", Version: "v2.3.0"},
		{Name: "github.com/vdemeester/k8s-pkg-credentialprovider", Version: "v1.17.4"},
		{Name: "github.com/vmware/govmomi", Version: "v0.20.3"},
		{Name: "github.com/xanzy/ssh-agent", Version: "v0.2.1"},
		{Name: "github.com/xiang90/probing", Version: "v0.0.0-20190116061207-43a291ad63a2"},
		{Name: "github.com/xordataexchange/crypt", Version: "v0.0.3-0.20170626215501-b2862e3d0a77"},
		{Name: "github.com/yashtewari/glob-intersection", Version: "v0.0.0-20180916065949-5c77d914dd0b"},
		{Name: "github.com/yuin/goldmark", Version: "v1.1.32"},
		{Name: "github.com/yuin/gopher-lua", Version: "v0.0.0-20191220021717-ab39c6098bdb"},
		{Name: "go.etcd.io/bbolt", Version: "v1.3.5"},
		{Name: "go.etcd.io/etcd", Version: "v0.0.0-20191023171146-3cf2f69b5738"},
		{Name: "go.opencensus.io", Version: "v0.22.4"},
		{Name: "go.opentelemetry.io/otel", Version: "v0.14.0"},
		{Name: "go.uber.org/atomic", Version: "v1.5.1"},
		{Name: "go.uber.org/multierr", Version: "v1.4.0"},
		{Name: "go.uber.org/tools", Version: "v0.0.0-20190618225709-2cfd321de3ee"},
		{Name: "go.uber.org/zap", Version: "v1.13.0"},
		{Name: "golang.org/x/crypto", Version: "v0.0.0-20201002170205-7f63de1d35b0"},
		{Name: "golang.org/x/exp", Version: "v0.0.0-20200224162631-6cc2880d07d6"},
		{Name: "golang.org/x/image", Version: "v0.0.0-20190802002840-cff245a6509b"},
		{Name: "golang.org/x/lint", Version: "v0.0.0-20200302205851-738671d3881b"},
		{Name: "golang.org/x/mobile", Version: "v0.0.0-20190719004257-d2bd2a29d028"},
		{Name: "golang.org/x/mod", Version: "v0.3.0"},
		{Name: "golang.org/x/net", Version: "v0.0.0-20201006153459-a7d1128ccaa0"},
		{Name: "golang.org/x/oauth2", Version: "v0.0.0-20201208152858-08078c50e5b5"},
		{Name: "golang.org/x/sync", Version: "v0.0.0-20200625203802-6e8e738ad208"},
		{Name: "golang.org/x/sys", Version: "v0.0.0-20201006155630-ac719f4daadf"},
		{Name: "golang.org/x/text", Version: "v0.3.3"},
		{Name: "golang.org/x/time", Version: "v0.0.0-20191024005414-555d28b269f0"},
		{Name: "golang.org/x/tools", Version: "v0.0.0-20200825202427-b303f430e36d"},
		{Name: "golang.org/x/xerrors", Version: "v0.0.0-20200804184101-5ec99f83aff1"},
		{Name: "gonum.org/v1/gonum", Version: "v0.0.0-20190331200053-3d26580ed485"},
		{Name: "gonum.org/v1/netlib", Version: "v0.0.0-20190331212654-76723241ea4e"},
		{Name: "google.golang.org/api", Version: "v0.30.0"},
		{Name: "google.golang.org/appengine", Version: "v1.6.6"},
		{Name: "google.golang.org/genproto", Version: "v0.0.0-20200825200019-8632dd797987"},
		{Name: "google.golang.org/grpc", Version: "v1.31.0"},
		{Name: "google.golang.org/protobuf", Version: "v1.25.0"},
		{Name: "gopkg.in/alecthomas/kingpin.v2", Version: "v2.2.6"},
		{Name: "gopkg.in/check.v1", Version: "v1.0.0-20200902074654-038fdea0a05b"},
		{Name: "gopkg.in/cheggaaa/pb.v1", Version: "v1.0.28"},
		{Name: "gopkg.in/errgo.v2", Version: "v2.1.0"},
		{Name: "gopkg.in/fsnotify.v1", Version: "v1.4.7"},
		{Name: "gopkg.in/gcfg.v1", Version: "v1.2.0"},
		{Name: "gopkg.in/go-playground/assert.v1", Version: "v1.2.1"},
		{Name: "gopkg.in/go-playground/validator.v9", Version: "v9.31.0"},
		{Name: "gopkg.in/inf.v0", Version: "v0.9.1"},
		{Name: "gopkg.in/mgo.v2", Version: "v2.0.0-20180705113604-9856a29383ce"},
		{Name: "gopkg.in/natefinch/lumberjack.v2", Version: "v2.0.0"},
		{Name: "gopkg.in/resty.v1", Version: "v1.12.0"},
		{Name: "gopkg.in/square/go-jose.v2", Version: "v2.2.2"},
		{Name: "gopkg.in/tomb.v1", Version: "v1.0.0-20141024135613-dd632973f1e7"},
		{Name: "gopkg.in/warnings.v0", Version: "v0.1.2"},
		{Name: "gopkg.in/yaml.v2", Version: "v2.4.0"},
		{Name: "gopkg.in/yaml.v3", Version: "v3.0.0-20200615113413-eeeca48fe776"},
		{Name: "gotest.tools", Version: "v2.2.0+incompatible"},
		{Name: "honnef.co/go/tools", Version: "v0.0.1-2020.1.4"},
		{Name: "k8s.io/api", Version: "v0.17.4"},
		{Name: "k8s.io/apimachinery", Version: "v0.17.4"},
		{Name: "k8s.io/apiserver", Version: "v0.17.4"},
		{Name: "k8s.io/client-go", Version: "v0.17.4"},
		{Name: "k8s.io/cloud-provider", Version: "v0.17.4"},
		{Name: "k8s.io/code-generator", Version: "v0.17.2"},
		{Name: "k8s.io/component-base", Version: "v0.17.4"},
		{Name: "k8s.io/csi-translation-lib", Version: "v0.17.4"},
		{Name: "k8s.io/gengo", Version: "v0.0.0-20190822140433-26a664648505"},
		{Name: "k8s.io/klog", Version: "v1.0.0"},
		{Name: "k8s.io/klog/v2", Version: "v2.0.0"},
		{Name: "k8s.io/kube-openapi", Version: "v0.0.0-20191107075043-30be4d16710a"},
		{Name: "k8s.io/legacy-cloud-providers", Version: "v0.17.4"},
		{Name: "k8s.io/utils", Version: "v0.0.0-20201110183641-67b214c5f920"},
		{Name: "modernc.org/cc", Version: "v1.0.0"},
		{Name: "modernc.org/golex", Version: "v1.0.0"},
		{Name: "modernc.org/mathutil", Version: "v1.0.0"},
		{Name: "modernc.org/strutil", Version: "v1.0.0"},
		{Name: "modernc.org/xc", Version: "v1.0.0"},
		{Name: "moul.io/http2curl", Version: "v1.0.0"},
		{Name: "rsc.io/binaryregexp", Version: "v0.2.0"},
		{Name: "rsc.io/quote/v3", Version: "v3.1.0"},
		{Name: "rsc.io/sampler", Version: "v1.3.0"},
		{Name: "sigs.k8s.io/structured-merge-diff", Version: "v1.0.1-0.20191108220359-b1b620dd3f06"},
		{Name: "sigs.k8s.io/yaml", Version: "v1.1.0"},
	}
)
