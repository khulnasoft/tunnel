# Terminology

This page explains the terminology system used in Tunnel, helping users understand the specific terms and concepts unique to the Tunnel ecosystem.

**Inclusion Criteria**

1. Core Components of Tunnel

   - Primary features such as Scanner, Target
   - Essential components such as Scan Assets (tunnel-db, tunnel-java-db)
   - Components that users directly interact with

2. Tunnel-specific Terms
   - Terms unique to Tunnel (e.g., VEX Hub)
   - Terms that have special meaning in Tunnel's context (e.g., Plugin, Module)

**Exclusion Criteria**

1. General Terms

   - Common security/technical terms (e.g., CVE, CVSS, Container, Registry)
   - Standard industry terminology

2. Implementation Details
   - Internal workings of components
   - Usage instructions (these belong in feature documentation)

## Core Concepts

### Target

Types of artifacts that Tunnel can scan, like container images and filesystem.

### Scanner

Tunnel's built-in security scanning engines. Tunnel has four main scanners:

- [Vulnerability Scanner](../scanner/vulnerability.md)
- [Misconfiguration Scanner](../scanner/misconfiguration/index.md)
- [Secret Scanner](../scanner/secret.md)
- [License Scanner](../scanner/license.md)

!!! note
SBOM is not a scanner but an output format option.

### Scan Assets

External data that Tunnel downloads (if needed for scanner) and uses during scanning:

- [Vulnerability Database (Tunnel DB, tunnel-db)](#vulnerability-database-tunnel-db-tunnel-db): Database containing vulnerability information
- [Java Index Database (Tunnel Java DB, tunnel-java-db)](#java-index-database-tunnel-java-db-tunnel-java-db): Database for Java artifact identification
- [Checks Bundle (tunnel-checks)](#checks-bundle): Archive containing misconfiguration detection rules
- [VEX Repository](#vex-repository): Repository containing VEX documents

## Vulnerability Scanning

### Vulnerability Database (Tunnel DB, tunnel-db)

The core vulnerability database required for vulnerability detection.
Contains comprehensive vulnerability information for multiple ecosystems.
Distributed via OCI registry.

Managed at https://go.khulnasoft.com/tunnel-db.

The vulnerability database is built from a GitHub repository that collects and stores vulnerability information from various data sources.
This repository serves as the foundation for building the Tunnel DB.

Managed at:

- https://github.com/khulnasoft/vuln-list
- https://github.com/khulnasoft/vuln-list-nvd
- https://github.com/khulnasoft/vuln-list-redhat
- https://github.com/khulnasoft/vuln-list-debian
- etc.

### Java Index Database (Tunnel Java DB, tunnel-java-db)

Specialized database used for identifying Java libraries and their components during JAR/WAR/PAR/EAR scanning.
Distributed via OCI registry.

Managed at https://github.com/aquasecurity/trivy-java-db.

## Misconfiguration Scanning

When the context does not clearly indicate these terms are related to misconfiguration scanning, they may be prefixed with "Misconfiguration" for clarity.
For example, "Check" may be referred to as "Misconfiguration Check", and "Checks Bundle" as "Misconfiguration Checks Bundle".

### Check

A Rego file that defines rules for detecting misconfigurations in various types of IaC files.

### Built-in Checks

Default set of checks distributed through [the tunnel-checks repository](https://github.com/aquasecurity/trivy-checks), providing standard security and configuration best practices.

### Checks Bundle

A tar.gz archive containing [the built-in checks](#built-in-checks), distributed via OCI registry.

## Secret Scanning

### Rule

Pattern matching rules used to detect hardcoded secrets and sensitive information.
Each rule consists of:

- Metadata (ID, Category, Title, etc.)
- Regular expressions for matching sensitive patterns
- Additional context for detection accuracy

## Kubernetes Integration

### KBOM (Kubernetes Bill of Materials)

A specialized SBOM format for Kubernetes clusters that includes detailed information about the cluster's components.

## VEX (Vulnerability Exploitability eXchange)

### VEX Repository

A repository system that stores VEX documents following [the VEX Repository Specification](https://github.com/khulnasoft/vex-repo-spec).
VEX repositories help users manage and share information about vulnerability applicability and exploitability.

For detailed information about VEX repositories, see [the document](../supply-chain/vex/repo.md).

### VEX Hub

The default VEX repository managed by KhulnaSoft Security at https://github.com/khulnasoft/vexhub.
It primarily aggregates VEX documents published by package maintainers in their source repositories.
VEX Hub serves as a central point for collecting and distributing vulnerability applicability information for OSS projects.

## Cache System

### Cache Types

The cache directory contains several distinct types of data:

- [Vulnerability Database](#vulnerability-database-tunnel-db-tunnel-db)
- [Java Index Database](#java-index-database-tunnel-java-db-tunnel-java-db)
- [Misconfiguration Checks](#misconfiguration-scanning)
- [VEX Repositories](#vex-repository)
- [Scan Cache](#scan-cache)

### Asset Cache

Downloaded assets like vulnerability databases and Java index databases.

### Scan Cache

A caching mechanism that stores analysis results from previous scans to speed up subsequent scans.
For container image scanning, the scan cache stores analysis results including package names and versions per layer.

For detailed information about caching, see [the document](../configuration/cache.md).

## Plugin System

### Plugin

An add-on tool that integrates with Tunnel to extend its core functionality.
Plugins can be written in any programming language and integrate seamlessly with Tunnel CLI, appearing in Tunnel help and subcommands.
They can be installed and removed independently without affecting the core Tunnel installation.

For detailed information about plugins, see [the document](../plugin/index.md).

### Plugin Index (tunnel-plugin-index)

A centralized registry that lists available Tunnel plugins, managed at https://github.com/khulnasoft/tunnel-plugin-index.
The index maintains a curated list of official and community plugins, providing metadata such as plugin names, descriptions, and maintainers.
It enables plugin discovery through the `tunnel plugin search` command and facilitates automatic plugin installation and updates.

For detailed information about the plugin index, see [the document](../plugin/user-guide.md).

## Module System

### Module

A WebAssembly-based extension mechanism that allows custom scanning logic without modifying the Tunnel binary.
Modules can modify scan results by analyzing files or post-processing results.

For detailed information about modules, see [the document](../advanced/modules.md).
