# Gin-Vue-Blog

一个基于 Gin + Vue3 的现代化博客系统，包含前后台管理界面。

[English Documentation](README_EN.md) | 中文文档

<p align="center">
  <img src="https://www.logoko.com.cn/uploadfile/icon_case/201808/5b86063b48933.png" alt="GVB Logo" width="200">
</p>

## 📚 项目介绍

Gin-Vue-Blog 是一个使用 Go 语言的 Gin 框架和 Vue3 开发的现代化博客系统。项目采用前后端分离架构，提供了完整的博客内容管理和用户管理功能。

### 🏗️ 项目结构

项目分为三个主要部分：

- [后台管理界面](gin-vue-blog_admin/README.md) - 基于 Vue3 + Ant Design Vue 的管理后台
- [后端服务](gin-vue-blog_server/README.md) - 基于 Gin 的 RESTful API 服务
- 前台展示界面 (开发中)

### ✨ 主要特性

- 🎯 基于 Gin + Vue3 的现代技术栈
- 🎨 优雅的用户界面设计
- 📦 完整的博客功能
- 🔐 权限管理系统
- 🌓 支持深色/浅色主题
- 📱 响应式设计

## 🚀 快速开始

### 环境要求

- Go >= 1.18
- Node.js >= 18.0.0
- MySQL >= 5.7
- Redis >= 6.0

### 后台管理系统启动

```bash
cd gin-vue-blog_admin
npm install
npm run dev
```

### 后端服务启动

```bash
cd gin-vue-blog_server
go mod tidy
go run main.go
```

## 📋 功能模块

### 后台管理系统

- 📊 仪表盘
- 👥 用户管理
- 📝 文章管理
- 🏷️ 标签管理
- 🖼️ 图片管理
- ⚙️ 系统设置

### 后端服务

- 🔐 JWT 认证
- 📧 邮件服务
- 💾 数据库操作
- 📤 文件上传
- 🔍 全文检索

## 🛠️ 技术栈

### 前端

- Vue 3
- Vite
- Ant Design Vue
- Pinia
- Vue Router

### 后端

- Gin
- GORM
- JWT
- Redis
- MySQL

## 📄 开源协议

本项目基于 [MIT 协议](LICENSE) 开源。

## 🤝 贡献指南

1. Fork 本仓库
2. 创建特性分支：`git checkout -b feature/AmazingFeature`
3. 提交改动：`git commit -m 'Add some AmazingFeature'`
4. 推送分支：`git push origin feature/AmazingFeature`
5. 提交 Pull Request

## 📞 联系我们

如有任何问题或建议，欢迎提交 [Issue](https://github.com/AA12-G/gin-vue-blog/issues) 或通过以下方式联系我们：

- Email: 811959019@example.com

---

<p align="center">用 ❤️ 制作</p>