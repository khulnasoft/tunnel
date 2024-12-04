# Tunnel Project Principles

This document outlines the guiding principles and governance framework for the Tunnel project.

## Core Principles

Tunnel is a security scanner focused on static analysis and designed with simplicity and security at its core.
All new proposals to the project must adhere to the following principles.

### Static Analysis (No Runtime Required)

Tunnel operates without requiring container or VM image startups, eliminating the need for Docker or similar runtimes, except for scanning images stored within a container runtime.
This approach enhances security and efficiency by minimizing dependencies.

### External Dependency Free (Single Binary)

Operating as a single binary, Tunnel is independent of external environments and avoids executing external OS commands or processes.
If specific functionality, like Maven's, is needed, Tunnel opts for internal reimplementations or processing outputs of the tool without direct execution of external tools.

This approach obviously requires more effort but significantly reduces security risks associated with executing OS commands and dependency errors due to external environment versions.
Simplifying the scanner's use by making it operational immediately upon binary download facilitates easier initiation of scans.

### No Setup Required

Tunnel must be ready to use immediately after installation.
It's unacceptable for Tunnel not to function without setting up a database or writing configuration files by default.
Such setups should only be necessary for users requiring specific customizations.

Security often isn't a top priority for many organizations and can be easily deferred.
Tunnel aims to lower the barrier to entry by simplifying the setup process, making it easier for users to start securing their projects.

### Security Focus

Tunnel prioritizes the identification of security issues, excluding features unrelated to security, such as performance metrics or content listings of container images.
It can, however, produce and output intermediate representations like SBOMs for comprehensive security assessments.

Tunnel serves as a tool with opinions on security, used to warn users about potential issues.

### Detecting Unintended States

Tunnel is designed to detect unintended vulnerable states in projects, such as the use of vulnerable versions of dependencies or misconfigurations in Infrastructure as Code (IaC) that may unintentionally expose servers to the internet.
The focus is on identifying developer mistakes or undesirable states, not on detecting intentional attacks, such as malicious images and malware.

## Out of Scope Features

KhulnaSoft Security offers a premium version with several features not available in the open-source Tunnel project.
While detailed information can be found [here][tunnel-khulnasoft], it's beneficial to highlight specific functionalities frequently inquired about:

### Runtime Security

As mentioned in [the Core Principles](#static-analysis-no-runtime-required), Tunnel is a static analysis security scanner, making runtime security outside its scope.
Runtime security needs are addressed by [Tracee][tracee] or [the commercial version of KhulnaSoft Security]().

### Intentional Attacks

As mentioned in [the Core Principles](#detecting-unintended-states), detection of intentional attacks, such as malware or malicious container images, is not covered by Tunnel and is supported in [the commercial version][khulnasoft].

### User Interface

Tunnel primarily operates via CLI for displaying results, with a richer UI available in [the commercial version][khulnasoft].

[tunnel-khulnasoft]: ../commercial/compare.md
[tracee]: https://github.com/khulnasoft/tracee
[khulnasoft]: https://www.khulnasoft.com/
