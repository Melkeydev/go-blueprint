#!/bin/bash

set -euo pipefail

VERSION="$1"
PACKAGE_NAME="go-blueprint"
MAIN_PACKAGE_DIR="npm-package"
PLATFORM_PACKAGES_DIR="platform-packages"

rm -rf "$MAIN_PACKAGE_DIR" "$PLATFORM_PACKAGES_DIR"

mkdir -p "$MAIN_PACKAGE_DIR/bin" "$PLATFORM_PACKAGES_DIR"

declare -A PLATFORM_MAP=(
    ["go-blueprint_${VERSION}_Darwin_all"]="darwin-x64,darwin-arm64"
    ["go-blueprint_${VERSION}_Linux_x86_64"]="linux-x64"
    ["go-blueprint_${VERSION}_Linux_arm64"]="linux-arm64"
    ["go-blueprint_${VERSION}_Windows_x86_64"]="win32-x64"
    ["go-blueprint_${VERSION}_Windows_arm64"]="win32-arm64"
)

declare -A OS_MAP=(
    ["darwin-x64"]="darwin"
    ["darwin-arm64"]="darwin"
    ["linux-x64"]="linux"
    ["linux-arm64"]="linux"
    ["win32-x64"]="win32"
    ["win32-arm64"]="win32"
)

declare -A CPU_MAP=(
    ["darwin-x64"]="x64"
    ["darwin-arm64"]="arm64"
    ["linux-x64"]="x64"
    ["linux-arm64"]="arm64"
    ["win32-x64"]="x64"
    ["win32-arm64"]="arm64"
)

OPTIONAL_DEPS=""
for archive in dist/*.tar.gz dist/*.zip; do
    if [ -f "$archive" ]; then
        archive_name=$(basename "$archive")
        archive_name="${archive_name%.tar.gz}"
        archive_name="${archive_name%.zip}"
        
        platform_keys="${PLATFORM_MAP[$archive_name]:-}"
        
        if [ -n "$platform_keys" ]; then
            echo "Processing $archive for platforms: $platform_keys"
            
            IFS=',' read -ra PLATFORM_ARRAY <<< "$platform_keys"
            for platform_key in "${PLATFORM_ARRAY[@]}"; do
                platform_key=$(echo "$platform_key" | xargs)
                
                echo "  Creating package for platform: $platform_key"
                
                platform_package_dir="$PLATFORM_PACKAGES_DIR/$PACKAGE_NAME-$platform_key"
                mkdir -p "$platform_package_dir/bin"
                
                if [[ "$archive" == *.tar.gz ]]; then
                    tar -xzf "$archive" -C "$platform_package_dir/bin"
                else
                    unzip -j "$archive" -d "$platform_package_dir/bin"
                fi
                
                for doc_file in README.md README README.txt LICENSE LICENSE.md LICENSE.txt; do
                    if [ -f "$platform_package_dir/bin/$doc_file" ]; then
                        mv "$platform_package_dir/bin/$doc_file" "$platform_package_dir/"
                    fi
                done
                
                ls -l "$platform_package_dir/bin"
                chmod +x "$platform_package_dir/bin/"*
                
                os_value="${OS_MAP[$platform_key]}"
                cpu_value="${CPU_MAP[$platform_key]}"
                
                files_array='["bin/"]'
                for doc_file in README.md README README.txt LICENSE LICENSE.md LICENSE.txt; do
                    if [ -f "$platform_package_dir/$doc_file" ]; then
                        files_array="${files_array%]}, \"$doc_file\"]"
                    fi
                done
                
                binary_name="go-blueprint"
                if [[ "$os_value" == "win32" ]]; then
                    binary_name="go-blueprint.exe"
                fi
                
                cat > "$platform_package_dir/package.json" << EOF
{
  "name": "$PACKAGE_NAME-$platform_key",
  "version": "$VERSION",
  "description": "Platform-specific binary for $PACKAGE_NAME ($platform_key)",
  "os": ["$os_value"],
  "cpu": ["$cpu_value"],
  "bin": {
    "go-blueprint": "bin/$binary_name"
  },
  "files": $files_array,
  "repository": {
    "type": "git",
    "url": "https://github.com/Melkeydev/go-blueprint.git"
  },
  "author": "Melkeydev",
  "license": "MIT"
}
EOF
                
                if [ -n "$OPTIONAL_DEPS" ]; then
                    OPTIONAL_DEPS="$OPTIONAL_DEPS,"
                fi
                OPTIONAL_DEPS="$OPTIONAL_DEPS\"$PACKAGE_NAME-$platform_key\": \"$VERSION\""
            done
        fi
    fi
done

cat > "$MAIN_PACKAGE_DIR/bin/go-blueprint" << 'EOF'
#!/usr/bin/env node

const { execFileSync } = require('child_process')

const packageName = 'go-blueprint'

const platformPackages = {
  'darwin-x64': `${packageName}-darwin-x64`,
  'darwin-arm64': `${packageName}-darwin-arm64`,
  'linux-x64': `${packageName}-linux-x64`,
  'linux-arm64': `${packageName}-linux-arm64`,
  'win32-x64': `${packageName}-win32-x64`,
  'win32-arm64': `${packageName}-win32-arm64`
}

function getBinaryPath() {
  const platformKey = `${process.platform}-${process.arch}`
  const platformPackageName = platformPackages[platformKey]

  if (!platformPackageName) {
    console.error(`Platform ${platformKey} is not supported!`)
    process.exit(1)
  }

  try {
    const binaryName = process.platform === 'win32' ? 'go-blueprint.exe' : 'go-blueprint'
    return require.resolve(`${platformPackageName}/bin/${binaryName}`)
  } catch (e) {
    process.exit(1)
  }
}

try {
  const binaryPath = getBinaryPath()
  execFileSync(binaryPath, process.argv.slice(2), { stdio: 'inherit' })
} catch (error) {
  console.error('Failed to execute go-blueprint:', error.message)
  process.exit(1)
}
EOF

chmod +x "$MAIN_PACKAGE_DIR/bin/go-blueprint"

cat > "$MAIN_PACKAGE_DIR/package.json" << EOF
{
  "name": "$PACKAGE_NAME",
  "version": "$VERSION",
  "description": "A CLI for scaffolding Go projects with modern tooling",
  "main": "index.js",
  "bin": {
    "go-blueprint": "bin/go-blueprint"
  },
  "optionalDependencies": {
    $OPTIONAL_DEPS
  },
  "keywords": ["go", "golang", "cli"],
  "author": "Melkeydev",
  "license": "MIT",
  "repository": {
    "type": "git",
    "url": "https://github.com/Melkeydev/go-blueprint.git"
  },
  "homepage": "https://github.com/Melkeydev/go-blueprint",
  "engines": {
    "node": ">=14.0.0"
  },
  "files": [
    "bin/",
    "index.js",
    "README.md"
  ]
}
EOF

cat > "$MAIN_PACKAGE_DIR/index.js" << 'EOF'
const { execFileSync } = require('child_process')
const path = require('path')

const binaryName = process.platform === 'win32' ? 'go-blueprint.exe' : 'go-blueprint'

const packageName = 'go-blueprint'

const platformPackages = {
  'darwin-x64': `${packageName}-darwin-x64`,
  'darwin-arm64': `${packageName}-darwin-arm64`,
  'linux-x64': `${packageName}-linux-x64`,
  'linux-arm64': `${packageName}-linux-arm64`,
  'win32-x64': `${packageName}-win32-x64`,
  'win32-arm64': `${packageName}-win32-arm64`
}

function getBinaryPath() {
  const platformKey = `${process.platform}-${process.arch}`
  const platformPackageName = platformPackages[platformKey]

  if (!platformPackageName) {
    throw new Error(`Platform ${platformKey} is not supported!`)
  }

  try {
    return require.resolve(`${platformPackageName}/bin/${binaryName}`)
  } catch (e) {
    throw new Error(`Platform-specific package ${platformPackageName} not found.`)
  }
}

module.exports = {
  getBinaryPath,
  run: function(...args) {
    const binaryPath = getBinaryPath()
    return execFileSync(binaryPath, args, { stdio: 'inherit' })
  }
}
EOF

first_platform_dir=$(ls -1d "$PLATFORM_PACKAGES_DIR"/* | head -1 2>/dev/null || echo "")
if [ -n "$first_platform_dir" ] && [ -f "$first_platform_dir/README.md" ]; then
    cp "$first_platform_dir/README.md" "$MAIN_PACKAGE_DIR/"
fi