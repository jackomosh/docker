# Architectural Decision Record: Multi-Stage Container Isolation

## Status
Approved

## Context
The system requires complete environment consistency. By mapping raw banner files directly to the root workspace layout, deployment configurations must be kept clean to minimize final container image sizes.

## Decision
We utilize a statically compiled 2-stage scratch build configuration:
1. The compiler stage optimizes caching by separating module parsing passes.
2. The runtime configuration strips structural debugging layers, drops system access privileges down to a standard non-root shell profile (`runtimeuser`), and mirrors the local tracking directory layouts perfectly.

## Consequences
* **Pros:** Runtime footprints drop significantly (< 20MB total).
* **Cons:** Internal interactive shell inspection binaries are removed from execution space.