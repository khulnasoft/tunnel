# Scanning Terraform files with Tunnel

This tutorial is focused on ways Tunnel can scan Terraform IaC configuration files.

A video tutorial on Terraform Misconfiguration scans can be found on the [KhulnaSoft Open Source YouTube account.](https://youtu.be/BWp5JLXkbBc)

**A note to tfsec users**
We have been consolidating all of our scanning-related efforts in one place, and that is Tunnel. You can read more on the decision in the [tfsec discussions.](https://github.com/khulnasoft/tfsec/discussions/1994)

## Tunnel Config Command

Terraform configuration scanning is available as part of the `tunnel config` command. This command scans all configuration files for misconfiguration issues. You can find the details within [misconfiguration scans in the Tunnel documentation.](https://khulnasoft.github.io/tunnel/latest/docs/scanner/misconfiguration/)

Command structure:

```
tunnel config <any flags you want to use> <file or directory that you would like to scan>
```

The `tunnel config` command can scan Terraform configuration, CloudFormation, Dockerfile, Kubernetes manifests, and Helm Charts for misconfiguration. Tunnel will compare the configuration found in the file with a set of best practices.

- If the configuration is following best practices, the check will pass,
- If the configuration does not define the resource of some configuration, Tunnel will assume the default configuration for the resource creation is used. In this case, the check might fail.
- If the configuration that has been defined does not follow best practices, the check will fail.

### Prerequisites

Install Tunnel on your local machines. The documentation provides several [different installation options.](https://khulnasoft.github.io/tunnel/latest/getting-started/installation/)
This tutorial will use this example [Terraform tutorial](https://github.com/Cloud-Native-Security/tunnel-demo/tree/main/bad_iac/terraform) for terraform misconfiguration scanning with Tunnel.

Git clone the tutorial and cd into the directory:

```
git clone git@github.com:Cloud-Native-Security/tunnel-demo.git
cd bad_iac/terraform
```

In this case, the folder only containes Terraform configuration files. However, you could scan a directory that contains several different configurations e.g. Kubernetes YAML manifests, Dockerfile, and Terraform. Tunnel will then detect the different configuration files and apply the right rules automatically.

## Different types of `tunnel config` scans

Below are several examples of how the tunnel config scan can be used.

General Terraform scan with tunnel:

```
tunnel config <specify the directory>
```

So if we are already in the directory that we want to scan:

```
tunnel config ./
```

### Specify the scan format

The `--format` flag changes the way that Tunnel displays the scan result:

JSON:

```
tunnel config -f json terraform-infra
```

Sarif:

```
tunnel config -f sarif terraform-infra
```

### Specifying the output location

The `--output` flag specifies the file location in which the scan result should be saved:

JSON:

```
tunnel config -f json -o example.json terraform-infra
```

Sarif:

```
tunnel config -f sarif -o example.sarif terraform-infra
```

### Filtering by severity

If you are presented with lots and lots of misconfiguration across different files, you might want to filter or the misconfiguration with the highest severity:

```
tunnel config --severity CRITICAL, MEDIUM terraform-infra
```

### Passing tf.tfvars files into `tunnel config` scans

You can pass terraform values to Tunnel to override default values found in the Terraform HCL code. More information are provided [in the documentation.](https://khulnasoft.github.io/tunnel/latest/docs/coverage/iac/terraform/#value-overrides)

```
tunnel config --tf-vars terraform.tfvars ./
```

### Custom Checks

We have lots of examples in the [documentation](https://khulnasoft.github.io/tunnel/latest/docs/scanner/misconfiguration/custom/) on how you can write and pass custom Rego checks into terraform misconfiguration scans.

## Secret and vulnerability scans

The `tunnel config` command does not perform secrete and vulnerability checks out of the box. However, you can specify as part of your `tunnel fs` scan that you would like to scan you terraform files for exposed secrets and misconfiguraction through the following flags:

```
tunnel fs --scanners secret,misconfig ./
```

The `tunnel config` command is a sub-command of the `tunnel fs` command. You can learn more about this command in the [documentation.](https://khulnasoft.github.io/tunnel/latest/docs/target/filesystem/)

## Scanning Terraform Plan files

Instead of scanning your different Terraform resources individually, you could also scan your Terraform Plan file before it is deployed for misconfiguration. This will give you insights into any misconfiguration of your resources as they would become deployed. [Here](https://khulnasoft.github.io/tunnel/latest/docs/coverage/iac/terraform/#terraform) is the link to the documentation.

Note that you need to be able to create a terraform init and plan without any errors.

## Using Tunnel in your CI/CD pipeline

Similar to tfsec, Tunnel can be used either on local developer machines or integrated into your CI/CD pipeline. There are several steps available for different pipelines, including GitHub Actions, Circle CI, GitLab, Travis and more in the tutorials section of the documentation: [https://khulnasoft.github.io/tunnel/latest/tutorials/integrations/](https://khulnasoft.github.io/tunnel/latest/tutorials/integrations/)
