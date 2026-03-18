# CAAD 任务工单清单 (Alpha 阶段)

请将以下列表录入到 GitHub/Gitee 的 Issues 中，并按 ID 分配给对应组员。

## 后端开发组 (Backend)

| ID | 标题 | 详细描述 | 优先级 |
|---|---|---|---|
| **#1** | [ALPHA-BE-01] 初始化项目骨架 | 使用 Maven/Go Mod 初始化项目，确认目录结构符合 MVC 规范。 | 高 |
| **#2** | [ALPHA-BE-02] 数据库 DDL 与驱动配置 | 编写 SQL 脚本创建 `users` 表，配置数据库连接池。 | 高 |
| **#3** | [ALPHA-BE-03] 注册接口设计与实现 | 实现 `POST /api/v1/auth/register`，包含 Bcrypt 加密及数据校验，由于是基建期，需配套 JUnit/Go Test 单元测试。 | 高 |
| **#4** | [ALPHA-BE-04] 登录接口与 JWT 签发 | 实现 `POST /api/v1/auth/login`，成功后签发 JWT。 | 高 |
| **#5** | [ALPHA-BE-05] 容器化与 CI 脚本 | 编写 Dockerfile 及 `.github/workflows/ci.yml`。 | 中 |

## 前端开发组 (Frontend)

| ID | 标题 | 详细描述 | 优先级 |
|---|---|---|---|
| **#6** | [ALPHA-FE-01] 初始化前端项目 | 使用 Vite 初始化 Vue/React 项目，配置 Axios 与 baseURL。 | 高 |
| **#7** | [ALPHA-FE-02] Apifox 接口 Mock 定义 | 依据接口规范，在 Apifox 录入注册、登录接口 Mock 数据。 | 高 |
| **:construction: #8** | [ALPHA-FE-03] 用户注册页面开发 | 实现注册视图 (View)，与 ViewModel 绑定，对接 Mock 接口。 | 高 |
| **:construction: #9** | [ALPHA-FE-04] 用户登录页面开发 | 实现登录视图 (View) 与 Token 本地存储逻辑。 | 高 |
| **#10** | [ALPHA-FE-05] UI 自动化测试初探 | 配置 Selenium 并编写一个基础的登录成功断言测试用例。 | 中 |
