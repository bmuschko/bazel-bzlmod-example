# Bazel Bzlmod Multi-Language Example

This repository demonstrates [Bzlmod](https://bazel.build/external/overview), Bazel's module-based external dependency management system, across three languages: Go, Java, and Python. Each subproject is a fully independent Bazel module with its own `MODULE.bazel`, dependency graph, and toolchain configuration.

## Repository Structure

```
├── go/                     # Go module
│   ├── MODULE.bazel        # rules_go + gazelle
│   ├── BUILD.bazel
│   ├── go.mod / go.sum
│   ├── main.go
│   └── main_test.go
├── java/                   # Java module
│   ├── MODULE.bazel        # rules_java + rules_jvm_external
│   ├── BUILD.bazel
│   ├── .bazelrc
│   └── src/
│       ├── main/java/com/example/App.java
│       └── test/java/com/example/AppTest.java
└── python/                 # Python module
    ├── MODULE.bazel        # rules_python + pip
    ├── BUILD.bazel
    ├── requirements_lock.txt
    ├── main.py
    └── main_test.py
```

Each subproject has its own `.bazelversion` (Bazel 9.0.2) and `MODULE.bazel.lock` for reproducible builds. There is no root-level Bazel workspace — the subprojects are fully decoupled.

## How Bzlmod Works in This Repository

### MODULE.bazel

Each subproject declares a [`MODULE.bazel`](https://bazel.build/external/module) at its root. This file replaces the legacy `WORKSPACE` file and serves two purposes: it marks the project root and declares all external dependencies.

Dependencies are added with [`bazel_dep()`](https://bazel.build/rules/lib/globals/module#bazel_dep), which fetches rule sets from the [Bazel Central Registry (BCR)](https://registry.bazel.build):

```starlark
bazel_dep(name = "rules_go", version = "0.60.0")
```

### Module Extensions

Language-specific dependency managers (Go modules, Maven, pip) are integrated through [module extensions](https://bazel.build/external/extension). Each extension bridges a language ecosystem's native dependency format into Bazel's dependency graph:

| Language | Extension | Dependency Source | Rule Set |
|----------|-----------|-------------------|----------|
| Go | `go_deps` | `go.mod` | [rules_go](https://github.com/bazel-contrib/rules_go) + [gazelle](https://github.com/bazelbuild/bazel-gazelle) |
| Java | `maven` | Maven coordinates | [rules_java](https://github.com/bazelbuild/rules_java) + [rules_jvm_external](https://github.com/bazel-contrib/rules_jvm_external) |
| Python | `pip` | `requirements_lock.txt` | [rules_python](https://github.com/bazel-contrib/rules_python) |

Extensions are invoked with [`use_extension()`](https://bazel.build/rules/lib/globals/module#use_extension) and their resolved repositories are made available to BUILD files via [`use_repo()`](https://bazel.build/rules/lib/globals/module#use_repo).

### Lock File

Each subproject has a [`MODULE.bazel.lock`](https://bazel.build/external/lockfile) that pins the entire resolved dependency graph. This ensures builds are reproducible across machines and CI without re-resolving dependencies.

## Building and Testing

Each subproject is built independently from its own directory:

```bash
cd go && bazelisk build //...
cd java && bazelisk build //...
cd python && bazelisk build //...
```

To run tests:

```bash
cd go && bazelisk test //...
cd java && bazelisk test //...
cd python && bazelisk test //...
```

## CI

The GitHub Actions [workflow](.github/workflows/build.yml) uses a matrix strategy to build and test all three subprojects in parallel.

## Bzlmod Documentation

- [External Dependencies Overview](https://bazel.build/external/overview) — introduction to bzlmod
- [MODULE.bazel Reference](https://bazel.build/rules/lib/globals/module) — API reference for `module()`, `bazel_dep()`, `use_extension()`, etc.
- [Module Extensions](https://bazel.build/external/extension) — how extensions bridge language ecosystems into Bazel
- [Lock Files](https://bazel.build/external/lockfile) — reproducible dependency resolution
- [Migration Guide](https://bazel.build/external/migration) — migrating from WORKSPACE to bzlmod
- [Bazel Central Registry](https://registry.bazel.build) — browsable index of available modules
