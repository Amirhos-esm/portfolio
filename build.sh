#!/usr/bin/env bash

set -euo pipefail

step () {
  echo
  echo "==> $1"
}

run () {
  msg="$1"
  shift
  step "$msg"
  if ! "$@"; then
    echo "ERROR: Failed at step: $msg"
    exit 1
  fi
}

# ------------------------
run "installing templ" go install github.com/a-h/templ/cmd/templ@latest

run "Generating Go files" go generate ./...

run "Preparing admin env" \
  cp admin/.env admin/v0-portfolio-admin-dashboard/.env

step "Building frontend"
cd admin/v0-portfolio-admin-dashboard || { echo "ERROR: cannot cd to admin dashboard"; exit 1; }

run "Installing frontend dependencies" pnpm install
run "Building frontend" pnpm build

run "Copying frontend output" cp -r out ../out

step "Returning to project root"
cd ../.. || { echo "ERROR: cannot return to root"; exit 1; }

mkdir -p output
run "Building Go binary" go build -o output/portfolio

run "Copying static files" cp -r static output/

echo
echo "✅ Build completed successfully"
