


# 🚀 Personal Portfolio

A modern developer portfolio built to showcase  projects, skills, and experience.

All fields — including personal information, skills, education, experience, projects, and messages — can be easily managed from the admin panel at ``` /admin ```

![Alt text](https://raw.githubusercontent.com/Amirhos-esm/portfolio/main/static/image.png)


---

## ✨ Features

* 🧑‍💻 Professional developer profile
* 📂 Project showcase section
* 🛠 Skills & technology overview
* 📱 Fully responsive design
* ⚡ Fast performance & optimized build
* 🖼 Image upload / profile assets support
* 🔐 Authentication-ready structure (extendable)

---

## 🧰 Tech Stack

### Frontend

* Next.js
* TypeScript
* Tailwind CSS
* Html/CSS with  <a href="https://github.com/a-h/templ">templ</a>

### Backend / APIs

* Golang/Graphql + REST (API routes) + templ for SSR


## 🎯 Purpose of This Project

This portfolio is designed to:

* Demonstrate real-world frontend skills
* Present my software engineering projects
* Serve as a professional personal website
* Act as a base template for future extensions

---


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





## 📬 Contact

If you'd like to connect:

* Telegram: <a href="https://t.me/amirhos_esm">@amirhos_esm</a>

---

## ⭐ Support

If you like this project:

Give it a ⭐ on GitHub 🙂

---

## 📝 License

MIT License
