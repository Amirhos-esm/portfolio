## 🛠️ Build Instructions

You can build the project using either the manual steps or the Linux helper script.

---

## ✅ Method 1 — Manual Build

### 1. Generate Go sources

```
go generate ./...
```

### 2. Copy environment file

```
cp admin/.env admin/v0-portfolio-admin-dashboard/.env
```

### 3. Enter admin dashboard directory

```
cd admin/v0-portfolio-admin-dashboard
```

### 4. Install dependencies

```
pnpm install
```

### 5. Build frontend

```
pnpm build
```

### 6. Copy static output

```
cp -r out ../out
```

### 7. Return to project root and build Go binary

```
cd ../..
go build -o output
```

---

## ✅ Method 2 — Linux Build Script (Recommended)

On Linux, you can run the included build script:

```
./build.sh
```

If needed, make it executable first:

```
chmod +x build.sh
```

---


