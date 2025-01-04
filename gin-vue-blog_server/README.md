# GVB Server - Gin 后端服务

<p align="center">
  <img src="https://gin-gonic.com/zh-cn/logo.jpg" alt="Gin Logo" width="180">
</p>

<p align="center">
  <a href="https://golang.org/"><img src="https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat&logo=go" alt="go version"></a>
  <a href="https://github.com/gin-gonic/gin"><img src="https://img.shields.io/badge/Gin-1.9.0-brightgreen?style=flat&logo=gin" alt="gin version"></a>
  <a href="https://gorm.io/"><img src="https://img.shields.io/badge/GORM-1.24.0-blue?style=flat" alt="gorm version"></a>
  <a href="https://redis.io/"><img src="https://img.shields.io/badge/Redis-6.0+-DC382D?style=flat&logo=redis" alt="redis version"></a>
</p>

## 📚 项目介绍

GVB Server 是基于 Gin 框架开发的博客系统后端服务，提供了完整的 RESTful API 接口。采用 JWT 认证、GORM 数据库操作、Redis 缓存等技术，实现了高性能、安全可靠的后端服务。

### ✨ 特性

- 🔐 JWT 身份认证
- 📝 文章 CRUD 操作
- 👤 用户管理系统
- 🏷️ 标签管理
- 📧 邮件服务
- 📤 文件上传
- 🔍 全文检索
- 📦 Redis 缓存
- 🛡️ 中间件机制
- 📊 数据统计

## 🚀 快速开始

### 环境要求

- Go >= 1.18
- MySQL >= 5.7
- Redis >= 6.0

### 配置文件

在 `config/config.yaml` 中配置数据库等信息：

```yaml
mysql:
  host: localhost
  port: 3306
  db: gvb_db
  user: root
  password: your_password
  log_level: dev

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0
```

### 启动服务

```bash
# 下载依赖
go mod tidy

# 运行服务
go run main.go
```

## 📁 项目结构

```
gin-vue-blog_server
├── api/                # API 接口层
│   ├── advert_api/    # 广告相关接口
│   ├── images_api/    # 图片处理接口
│   ├── menu_api/      # 菜单管理接口
│   ├── user_api/      # 用户相关接口
│   └── ...
├── config/            # 配置文件
├── global/            # 全局变量
├── middleware/        # 中间件
├── models/            # 数据模型
├── routers/          # 路由配置
├── service/          # 业务逻辑层
├── utils/            # 工具函数
└── main.go           # 程序入口
```

## 🔗 API 文档

### 用户相关

```
POST   /api/users/login        # 用户登录
POST   /api/users/register     # 用户注册
GET    /api/users/info         # 获取用户信息
PUT    /api/users/info         # 更新用户信息
```

### 文章相关

```
POST   /api/articles           # 创建文章
GET    /api/articles          # 文章列表
GET    /api/articles/:id      # 文章详情
PUT    /api/articles/:id      # 更新文章
DELETE /api/articles/:id      # 删除文章
```

### 标签相关

```
POST   /api/tags              # 创建标签
GET    /api/tags             # 标签列表
PUT    /api/tags/:id         # 更新标签
DELETE /api/tags/:id         # 删除标签
```

更多 API 详情请参考 Swagger 文档。

## 🛠️ 核心依赖

- [Gin](https://github.com/gin-gonic/gin) - Web 框架
- [GORM](https://gorm.io/) - ORM 框架
- [JWT-Go](https://github.com/golang-jwt/jwt) - JWT 认证
- [Go-Redis](https://github.com/go-redis/redis) - Redis 客户端
- [Viper](https://github.com/spf13/viper) - 配置管理
- [Zap](https://github.com/uber-go/zap) - 日志管理

## 📦 中间件

- JWT 认证中间件
- CORS 跨域中间件
- 日志记录中间件
- 错误处理中间件
- 权限验证中间件

## 🔐 安全特性

- 密码加密存储
- JWT 令牌认证
- 请求频率限制
- SQL 注入防护
- XSS 防护

## 💡 开发规范

- 遵循 RESTful API 设计规范
- 使用统一的错误处理机制
- 规范的日志记录
- 统一的响应格式

## 🤝 贡献指南

1. Fork 本仓库
2. 创建特性分支
3. 提交代码
4. 创建 Pull Request

## 📄 开源协议

本项目基于 [MIT 协议](../LICENSE) 开源。

## 📞 联系我们

如有任何问题或建议，欢迎提交 [Issue](https://github.com/AA12-G/gin-vue-blog/issues)。

---

<p align="center">用 ❤️ 制作</p> 