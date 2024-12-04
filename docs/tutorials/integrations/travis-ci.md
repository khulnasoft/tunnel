# Travis CI

```
$ cat .travis.yml
services:
  - docker

env:
  global:
    - COMMIT=${TRAVIS_COMMIT::8}

before_install:
  - docker build -t tunnel-ci-test:${COMMIT} .
  - export VERSION=$(curl --silent "https://api.github.com/repos/khulnasoft/tunnel/releases/latest" | grep '"tag_name":' | sed -E 's/.*"v([^"]+)".*/\1/')
  - wget https://github.com/khulnasoft/tunnel/releases/download/v${VERSION}/tunnel_${VERSION}_Linux-64bit.tar.gz
  - tar zxvf tunnel_${VERSION}_Linux-64bit.tar.gz
script:
  - ./tunnel image --exit-code 0 --severity HIGH --no-progress tunnel-ci-test:${COMMIT}
  - ./tunnel image --exit-code 1 --severity CRITICAL --no-progress tunnel-ci-test:${COMMIT}
cache:
  directories:
    - $HOME/.cache/tunnel
```

[Example][example]
[Repository][repository]

[example]: https://travis-ci.org/aquasecurity/tunnel-ci-test
[repository]: https://github.com/aquasecurity/tunnel-ci-test
