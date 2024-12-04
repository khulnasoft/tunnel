# CircleCI

```
$ cat .circleci/config.yml
jobs:
  build:
    docker:
      - image: docker:stable-git
    steps:
      - checkout
      - setup_remote_docker
      - run:
          name: Build image
          command: docker build -t tunnel-ci-test:${CIRCLE_SHA1} .
      - run:
          name: Install tunnel
          command: |
            apk add --update-cache --upgrade curl
            curl -sfL https://raw.githubusercontent.com/khulnasoft/tunnel/main/contrib/install.sh | sh -s -- -b /usr/local/bin
      - run:
          name: Scan the local image with tunnel
          command: tunnel image --exit-code 0 --no-progress tunnel-ci-test:${CIRCLE_SHA1}
workflows:
  version: 2
  release:
    jobs:
      - build
```

[Example][example]
[Repository][repository]

[example]: https://circleci.com/gh/aquasecurity/tunnel-ci-test
[repository]: https://github.com/aquasecurity/tunnel-ci-test
