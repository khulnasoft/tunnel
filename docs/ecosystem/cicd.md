# CI/CD Integrations

## Azure DevOps (Official)

[Azure Devops](https://azure.microsoft.com/en-us/products/devops/#overview) is Microsoft Azure cloud native CI/CD service.

Tunnel has a "Azure Devops Pipelines Task" for Tunnel, that lets you easily introduce security scanning into your workflow, with an integrated Azure Devops UI.

👉 Get it at: <https://github.com/aquasecurity/tunnel-azure-pipelines-task>

## GitHub Actions

[GitHub Actions](https://github.com/features/actions) is GitHub's native CI/CD and job orchestration service.

### tunnel-action (Official)

GitHub Action for integrating Tunnel into your GitHub pipeline

👉 Get it at: <https://github.com/aquasecurity/tunnel-action>

### tunnel-action (Community)

GitHub Action to scan vulnerability using Tunnel. If vulnerabilities are found by Tunnel, it creates a GitHub Issue.

👉 Get it at: <https://github.com/marketplace/actions/tunnel-action>

### tunnel-github-issues (Community)

In this action, Tunnel scans the dependency files such as package-lock.json and go.sum in your repository, then create GitHub issues according to the result.

👉 Get it at: <https://github.com/marketplace/actions/tunnel-github-issues>

## Buildkite Plugin (Community)

The tunnel buildkite plugin provides a convenient mechanism for running the open-source tunnel static analysis tool on your project.

👉 Get it at: https://github.com/equinixmetal-buildkite/tunnel-buildkite-plugin

## Dagger (Community)

[Dagger](https://dagger.io/) is CI/CD as code that runs anywhere.

The Dagger module for Tunnel provides functions for scanning container images from registries as well as Dagger Container objects from any Dagger SDK (e.g. Go, Python, Node.js, etc).

👉 Get it at: <https://daggerverse.dev/mod/github.com/jpadams/daggerverse/tunnel>

## Semaphore (Community)

[Semaphore](https://semaphoreci.com/) is a CI/CD service.

You can use Tunnel in Semaphore for scanning code, containers, infrastructure, and Kubernetes in Semaphore workflow.

👉 Get it at: <https://semaphoreci.com/blog/continuous-container-vulnerability-testing-with-tunnel>

## CircleCI (Community)

[CircleCI](https://circleci.com/) is a CI/CD service.

You can use the Tunnel Orb for Circle CI to introduce security scanning into your workflow.

👉 Get it at: <https://circleci.com/developer/orbs/orb/fifteen5/tunnel-orb>
Source: <https://github.com/15five/tunnel-orb>

## Woodpecker CI (Community)

Example Tunnel step in pipeline

```yml
pipeline:
  securitycheck:
    image: khulnasoft/tunnel:latest
    commands:
      # use any tunnel command, if exit code is 0 woodpecker marks it as passed, else it assumes it failed
      - tunnel fs --exit-code 1 --skip-dirs web/ --skip-dirs docs/ --severity MEDIUM,HIGH,CRITICAL .
```

Woodpecker does use Tunnel itself so you can [see it in use there](https://github.com/woodpecker-ci/woodpecker/pull/1163).

## Concourse CI (Community)

[Concourse CI](https://concourse-ci.org/) is a CI/CD service.

You can use Tunnel Resource in Concourse for scanning containers and introducing security scanning into your workflow.
It has capabilities to fail the pipeline, create issues, alert communication channels (using respective resources) based on Tunnel scan output.

👉 Get it at: <https://github.com/Comcast/tunnel-resource/>

## SecObserve GitHub actions and GitLab templates (Community)

[SecObserve GitHub actions and GitLab templates](https://github.com/MaibornWolff/secobserve_actions_templates) run various vulnerability scanners, providing uniform methods and parameters for launching the tools.

The Tunnel integration supports scanning Docker images and local filesystems for vulnerabilities as well as scanning IaC files for misconfigurations.

👉 Get it at: <https://github.com/MaibornWolff/secobserve_actions_templates>
