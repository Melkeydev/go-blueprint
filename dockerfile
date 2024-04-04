FROM ghcr.io/catthehacker/ubuntu:full-latest
USER 0

# Fixing issue https://github.com/nektos/act/issues/935 for local test

# [Linting Generated Blueprints/install_dependencies]   💬  ::debug::Caching tool go 1.22.2 x64
# [Linting Generated Blueprints/install_dependencies]   💬  ::debug::source dir: /tmp/e84efa31-5b27-4bac-95b2-51cc641ed4a4/go
# [Linting Generated Blueprints/install_dependencies]   💬  ::debug::destination /opt/hostedtoolcache/go/1.22.2/x64
# [Linting Generated Blueprints/install_dependencies]   ❗  ::error::Failed to download version 1.22.2: Error: EACCES: permission denied, mkdir '/opt/hostedtoolcache/go/1.22.2'
# [Linting Generated Blueprints/install_dependencies]   ❌  Failure - Main Setup Go
# [Linting Generated Blueprints/install_dependencies] exitcode '1': failure
# [Linting Generated Blueprints/install_dependencies] 🏁  Job failed

