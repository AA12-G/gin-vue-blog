# Gin-Vue-Blog

A modern blog system based on Gin + Vue3, including front-end and back-end management interfaces.

[中文文档](README.md) | English Documentation

<p align="center">
  <img src="https://www.logoko.com.cn/uploadfile/icon_case/201808/5b86063b48933.png" alt="GVB Logo" width="200">
</p>

## 📚 Project Introduction

Gin-Vue-Blog is a modern blog system developed using Go's Gin framework and Vue3. The project adopts a front-end and back-end separation architecture, providing complete blog content management and user management functions.

### 🏗️ Project Structure

The project consists of three main parts:

- [Admin Dashboard](gin-vue-blog_admin/README.md) - Management backend based on Vue3 + Ant Design Vue
- [Backend Service](gin-vue-blog_server/README.md) - RESTful API service based on Gin
- Frontend Display Interface (In Development)

### ✨ Key Features

- 🎯 Modern tech stack based on Gin + Vue3
- 🎨 Elegant user interface design
- 📦 Complete blog functionality
- 🔐 Permission management system
- 🌓 Support for dark/light themes
- 📱 Responsive design

## 🚀 Quick Start

### Requirements

- Go >= 1.18
- Node.js >= 18.0.0
- MySQL >= 5.7
- Redis >= 6.0

### Start Admin System

```bash
cd gin-vue-blog_admin
npm install
npm run dev
```

### Start Backend Service

```bash
cd gin-vue-blog_server
go mod tidy
go run main.go
```

## 📋 Function Modules

### Admin System

- 📊 Dashboard
- 👥 User Management
- 📝 Article Management
- 🏷️ Tag Management
- 🖼️ Image Management
- ⚙️ System Settings

### Backend Service

- 🔐 JWT Authentication
- 📧 Email Service
- 💾 Database Operations
- 📤 File Upload
- 🔍 Full-text Search

## 🛠️ Technology Stack

### Frontend

- Vue 3
- Vite
- Ant Design Vue
- Pinia
- Vue Router

### Backend

- Gin
- GORM
- JWT
- Redis
- MySQL

## 📄 License

This project is licensed under the [MIT License](LICENSE).

## 🤝 Contributing

1. Fork the repository
2. Create feature branch: `git checkout -b feature/AmazingFeature`
3. Commit changes: `git commit -m 'Add some AmazingFeature'`
4. Push branch: `git push origin feature/AmazingFeature`
5. Submit Pull Request

## 📞 Contact Us

If you have any questions or suggestions, please submit an [Issue](https://github.com/AA12-G/gin-vue-blog/issues) or contact us through:

- Email: 811959019@example.com

---

<p align="center">Made with ❤️</p> 