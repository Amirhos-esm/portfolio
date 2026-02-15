bash
#!/usr/bin/env bash
set -e

echo "==> Generating Go files"
go generate ./...

echo "==> Preparing admin env"
cp admin/.env admin/v0-portfolio-admin-dashboard/.env

echo "==> Building frontend"
cd admin/v0-portfolio-admin-dashboard
pnpm install
pnpm build

echo "==> Copying frontend output"
cp -r out ../out

echo "==> Building Go binary"
cd ../..
go build -o output

echo "==> Build completed successfully"