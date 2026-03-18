# 软件需求规格说明书
# 校园活动聚合分发平台（CAAD）
## Software Requirements Specification

| 文档属性 | 内容 |
|---|---|
| **项目名称** | 校园活动聚合分发平台（Campus Activity Aggregation & Distribution Platform，简称 CAAD） |
| **文档版本** | v1.0 |
| **撰写日期** | 2026-03-12 |
| **隶属课程** | 大规模信息系统开发试验（ULSS-26） |
| **文档状态** | 正式草稿（Formal Draft） |

---

## 目录

1. 引言
2. 总体描述
3. 系统功能性需求
4. 系统非功能性需求
5. 系统架构设计
6. 数据需求
7. 接口需求
8. 迭代与交付规划
9. 附录

---

## 1. 引言（Introduction）

### 1.1 目的
本文档是《校园活动聚合分发平台（CAAD）》项目的正式需求规格说明书（SRS），面向全体开发团队成员（前端组、后端组）、测试人员及课程指导教师。文档旨在完整、准确、无歧义地描述系统的功能边界、技术约束与验证标准，是后续设计、实现与测试的基准参考文件。

### 1.2 项目背景
当前高校内各类学术讲座、体育赛事、社团活动信息散落于多个渠道（微信群、公告栏、各院系官网），普通学生难以统一获取并及时报名。尤其是热门活动开放报名瞬间，大量学生同时涌入报名接口，频繁引发系统崩溃、超额录取、数据不一致等严重工程问题。

本系统以"**数据聚合、算法驱动分发、架构保障并发**"为核心设计理念，构建一个面向全校师生的活动信息聚合、智能个性化推荐与高并发报名抢票的一体化信息平台。

### 1.3 课程对齐说明
本系统的建设完整覆盖《大规模信息系统开发试验》课程所要求的"大规模"三大维度：

| 课程定义的大规模维度 | 本系统的具体体现 |
|---|---|
| **大量的小数据** | 用户每次浏览、收藏、点击、报名成功或失败，均写入行为流水表，随用户规模持续线性增长 |
| **大规模的各种算法** | 活动热度多维加权评分算法（定时批量计算）、基于用户行为标签的协同过滤推荐算法 |
| **大量的交互** | 热门活动报名开放瞬间，预计支撑千级以上并发请求，通过 Redis 库存原子操作与消息队列削峰填谷实现 |

### 1.4 术语与缩略语

| 术语 | 解释 |
|---|---|
| CAAD | Campus Activity Aggregation & Distribution Platform，本系统简称 |
| 活动 | 泛指讲座、赛事、社团公开活动等需要线上报名的校园事件 |
| 报名 / 抢票 | 用户针对某一活动发出参与申请的行为，两词在本文中同义 |
| 名额 | 某活动允许报名的最大人数上限 |
| 热度分 | 系统后台计算的综合热度评分，用于排行与推荐 |
| 行为流水 | 用户所有操作（浏览、收藏、报名）的原始记录，不可删除，供算法使用 |
| JWT | JSON Web Token，本系统采用的无状态鉴权令牌 |
| MQ | 消息队列（Message Queue），本系统中用于抢票削峰 |

---

## 2. 总体描述（Overall Description）

### 2.1 产品愿景
一个"让每位同学都不错过任何感兴趣的校园活动"的智能聚合平台。
- **对学生**：统一入口发现活动；个性化推荐贴合兴趣；高并发报名公平不崩溃。
- **对管理员**：高效管理活动信息；活动热度数据可视；报名情况一览无余。

### 2.2 用户类型（User Classes）

| 用户类型 | 描述 | 主要操作权限 |
|---|---|---|
| **普通学生** | 持有学号的在校学生，系统的核心用户群 | 注册/登录、浏览、收藏、报名、查询个人报名记录 |
| **系统管理员** | 由学校/组织方授权，负责内容运营的人员 | 以上全部权限 + 创建/编辑/下架活动、查看所有报名数据 |

### 2.3 运行环境
- **前端**：现代浏览器（Chrome 90+, Firefox 88+, Edge 90+）
- **后端**：服务器运行于 Linux 系统的 Docker 容器内
- **部署理念**：前后端完全分离部署，前端通过 Nginx 分发静态资源并代理 API 请求到后端服务

### 2.4 约束条件
1. 后端框架约束为 MVC 架构（如 Spring Boot / Go Gin）。
2. 前端框架约束为 MVVM 架构（如 Vue 3 / React）。
3. 本系统的教学实例选题已由课程教师确认，不得与电商系统选题重叠。
4. 所有代码提交须遵循 Code Review 制度，禁止直接推送至仓库主干分支。

---

## 3. 功能性需求（Functional Requirements）

> 以下各用例 ID 为正式编号，后续任务工单须与编号关联。

### 3.1 用户身份认证模块（AUTH）

#### AUTH-01：用户注册
- **描述**：学生使用学号、用户名及密码完成账号注册。
- **前置条件**：用户尚未注册。
- **主流程**：
  1. 用户填写学号、用户名、密码、确认密码并提交。
  2. 后端校验学号格式合法性及唯一性（不可重复注册）。
  3. 密码以 Bcrypt 不可逆哈希算法加密后持久化存储。
  4. 注册成功，系统返回成功提示，引导用户前往登录页。
- **异常流程**：
  - 学号已被注册：返回错误提示"该学号已被注册"。
  - 两次密码不一致：前端实时校验拦截，后端再次校验。
  - 密码强度不达标（不足8位）：返回错误提示。
- **优先级**：高

#### AUTH-02：用户登录
- **描述**：注册用户使用学号与密码换取访问令牌（JWT Token）。
- **主流程**：
  1. 用户填写学号、密码，提交登录请求。
  2. 后端验证学号存在性，用 Bcrypt 校验密码。
  3. 验证通过后签发 JWT Token（有效期 24 小时），返回给前端。
  4. 前端将 Token 存入 localStorage，后续请求在 HTTP Header `Authorization: Bearer <token>` 中携带。
- **异常流程**：
  - 学号或密码错误：统一返回"账号或密码错误"（不区分哪个错误，以防枚举攻击）。
  - 账号已被封禁：返回"账号已被管理员停用"。
- **优先级**：高

#### AUTH-03：用户登出
- **描述**：用户主动登出，前端清除本地 Token，后端将此 Token 加入 Redis 黑名单（Token 剩余有效期内有效）。
- **优先级**：中

#### AUTH-04：管理员账号（初始化数据）
- **描述**：管理员账号不通过注册接口产生，由系统初始化时通过数据库种子脚本（Seed Script）直接写入，角色字段标记为 `ADMIN`。
- **优先级**：高

---

### 3.2 活动信息管理模块（ACTIVITY）

#### ACTIVITY-01：创建活动（管理员）
- **描述**：管理员在后台填写并发布一条新活动信息。
- **必填字段**：活动标题、活动简介、分类标签（多选，如"人工智能"、"体育"、"创业"）、主讲人/负责人、活动开始时间、报名开放时间、报名截止时间、总名额上限。
- **约束**：报名开放时间须早于活动开始时间；总名额须为正整数。
- **优先级**：高

#### ACTIVITY-02：编辑/下架活动（管理员）
- **描述**：管理员可对已发布活动进行信息修改或将其下架（从学生可见列表中隐藏）。
- **约束**：若报名已开放（已有报名记录），则不允许修改总名额上限和报名开放时间（以保障公平性）。
- **优先级**：中

#### ACTIVITY-03：浏览活动列表（学生）
- **描述**：学生可在活动列表页查看所有已发布的活动，并进行筛选与排序。
- **筛选维度**：活动分类标签（单选）、活动时间范围（近期 / 本月）、报名状态（可报名 / 已结束）。
- **排序维度**：热度排行（默认）、最新发布时间、活动即将开始时间。
- **分页**：每页展示 10 条，支持翻页。
- **优先级**：高

#### ACTIVITY-04：查看活动详情（学生）
- **描述**：学生点击某条活动后进入详情页，可查看活动全部信息。
- **展示字段**：标题、简介、主讲人、活动地点（可选）、活动时间、报名时间窗口、**当前剩余名额（实时）**、已报名人数。
- **行为记录**：用户进入详情页时，后台异步写入一条 `behavior_type=VIEW` 的行为流水，供推荐系统使用。
- **优先级**：高

#### ACTIVITY-05：收藏活动（学生）
- **描述**：学生可对感兴趣的活动进行收藏，收藏后可在"我的收藏"页统一查看。
- **行为记录**：收藏操作异步写入 `behavior_type=COLLECT` 的行为流水。
- **约束**：同一活动只能收藏一次，再次点击为取消收藏。
- **优先级**：中

---

### 3.3 高并发报名（抢票）模块（ENROLL）

#### ENROLL-01：报名状态显示
- **描述**：在活动详情页，根据当前时间与活动报名时间窗口，动态展示报名按钮状态：
  - `报名未开始`（灰色，展示距开放倒计时）
  - `立即报名`（可点击，绿色）
  - `报名已截止`（灰色）
  - `已报满`（灰色）
  - `已报名`（蓝色，不可再次点击）
- **优先级**：高

#### ENROLL-02：发起报名请求（核心）
- **描述**：学生在报名开放时间窗口内点击"立即报名"发起报名请求。
- **后端主流程（高并发防超卖三步法）**：
  1. **第一道防线（Redis 原子扣减）**：后端执行 Lua 脚本在 Redis 中原子地检查并扣减名额。若 `stock <= 0`，直接返回"报名失败，名额已满"。
  2. **第二道防线（幂等校验）**：检查 Redis 中是否已存在该用户对该活动的报名记录（防重复提交），若存在则返回"您已提交过报名请求，请勿重复操作"。
  3. **异步投递（消息队列削峰）**：通过前两道防线后，将报名请求投递至 RabbitMQ/Kafka 的报名队列，返回前端 `HTTP 202 Accepted`，并附带 `taskId`。
- **优先级**：核心，高

#### ENROLL-03：后台异步消费与落库（Worker）
- **描述**：消费者服务从报名队列中逐条消费消息，执行数据库写入操作。
- **消费流程**：
  1. 读取消息，在数据库中创建一条状态为 `PENDING` 的报名记录。
  2. 再次检查数据库中该活动已成功报名的总人数（最终一致性保障）。
  3. 若不超额，将记录状态更新为 `SUCCESS`；否则更新为 `FAIL`，并将 Redis 库存加回 1。
  4. 消费成功后向 MQ 发送 `ACK`，消息从队列移除。
- **优先级**：核心，高

#### ENROLL-04：查询报名结果
- **描述**：前端在收到 `202 Accepted` 后，利用 `taskId` 以轮询方式（每 2 秒）请求查询接口，直到得到 `SUCCESS` 或 `FAIL`。
- **优先级**：高

#### ENROLL-05：我的报名记录
- **描述**：学生在个人中心的"我的报名"页查看所有历史报名记录，包含活动名称、报名时间、最终状态。
- **优先级**：高

---

### 3.4 热度评分与排行模块（TRENDING）

#### TRENDING-01：热度评分计算（核心）
- **描述**：后台定时任务（每 10 分钟执行一次）遍历所有已发布活动，根据以下公式重新计算热度分，写入 Redis Sorted Set：

  ```
  热度分 = log₁₀(报名人数 + 1) × 5.0
           + log₁₀(浏览次数 + 1) × 2.0
           + 收藏次数 × 1.5
           + 时效系数（距活动开始越近系数越高，最大 3.0）
  ```

- **优先级**：核心，高

#### TRENDING-02：热门活动榜单
- **描述**：首页提供"本周热门活动 Top 10"榜单，直接从 Redis Sorted Set 中读取，无需查询数据库。
- **优先级**：中

---

### 3.5 个性化推荐模块（RECOMMEND）

#### RECOMMEND-01：用户标签构建
- **描述**：后台每日离线任务（每日凌晨 02:00 执行），根据近 30 天的用户行为流水，为每位用户计算兴趣标签权重向量：

  ```
  标签 T 的兴趣权重 =
    SUM( 浏览×1.0 次数 + 报名×3.0 次数 + 收藏×2.0 次数 )
    针对所有涉及标签 T 的行为记录
  ```

- **结果存储**：写入 Redis 哈希结构，有效期 25 小时（日任务刷新）。
- **优先级**：高

#### RECOMMEND-02：个性化推荐列表
- **描述**：在活动列表页的"为你推荐"分栏，展示最多 6 条个性化推荐活动。
- **计算方式**：读取用户标签权重向量，计算与各活动标签集合的余弦相似度，过滤已报名活动，按相似度降序取前 6 条。
- **冷启动策略**：新用户暂无行为数据时，推荐热度分 Top 6 的活动。
- **优先级**：高

---

## 4. 非功能性需求（Non-Functional Requirements）

### 4.1 性能需求

| 指标 | 目标值 | 测量手段 |
|---|---|---|
| 正常流量下 API 响应时间 | P95 < 200ms | JMeter 压测报告 |
| 高并发报名（1000 并发）下系统响应 | P99 < 500ms，不超卖 | JMeter 并发压测 |
| 活动列表页加载时间 | < 1 秒（含网络） | 浏览器 DevTools |

### 4.2 数据一致性需求

| 规则 | 说明 |
|---|---|
| **零超卖保证** | 任意时刻，数据库中某活动的成功报名记录总数 ≤ 该活动名额上限 |
| **幂等性保证** | 同一用户对同一活动的多次并发报名请求，最多只产生一条成功报名记录 |

### 4.3 安全性需求

| 安全要求 | 实现方式 |
|---|---|
| 用户密码安全存储 | Bcrypt 不可逆哈希，禁止明文或 MD5 |
| SQL 注入防护 | 所有数据库操作通过 ORM 参数化查询，禁止字符串拼接 SQL |
| XSS 防护 | 用户输入内容渲染前进行 HTML Entity 转义 |
| 认证鉴权 | 所有非公开 API 路由必须校验 JWT Token 有效性及角色权限 |
| 速率限制 | 报名接口设置单 IP/单用户频率限制（如每 5 秒限 1 次） |

### 4.4 可扩展性需求
- 推荐模块与抢票模块在架构上须保持解耦，可独立进行水平扩展。
- 数据库字段设计须预留扩展，避免后期 `ALTER TABLE` 影响线上服务。

---

## 5. 系统架构设计（Architecture Design）

### 5.1 整体技术架构

```
┌───────────────────────────────────────────────┐
│           用户浏览器                            │
│      前端 SPA (Vue 3 / React + Vite)           │
└───────────────────┬───────────────────────────┘
                    │ HTTP/HTTPS
┌───────────────────▼───────────────────────────┐
│     Nginx (反向代理 + 静态资源分发)              │
└───────────────────┬───────────────────────────┘
                    │
┌───────────────────▼───────────────────────────┐
│    后端应用层 (Spring Boot MVC / Gin)           │
│  用户模块 | 活动模块 | 报名模块 | 推荐模块       │
└────────┬───────────────────────┬──────────────┘
         │                       │
┌────────▼────────┐   ┌──────────▼──────────────┐
│   Redis 缓存层   │   │ 消息队列 (RabbitMQ/Kafka) │
│ - 活动名额库存   │   │ - 报名削峰队列             │
│ - 热度 SortedSet│   └──────────┬──────────────┘
│ - 用户标签向量   │              │
│ - JWT 黑名单    │   ┌──────────▼──────────────┐
└─────────────────┘   │  消费者 Worker (异步落库) │
                      └──────────┬──────────────┘
                                 │
┌────────────────────────────────▼──────────────┐
│  关系型数据库 (MySQL 8 / PostgreSQL)            │
│  users | activities | enrollments             │
│  user_behaviors | activity_scores             │
└───────────────────────────────────────────────┘
```

### 5.2 高并发报名核心时序

```
用户      前端        后端API       Redis       MQ        Worker      DB
 │         │             │            │          │            │         │
 │─报名───>│             │            │          │            │         │
 │         │──POST ─────>│            │          │            │         │
 │         │             │──Lua扣减──>│          │            │         │
 │         │             │<──OK───────│          │            │         │
 │         │             │──幂等检查─>│          │            │         │
 │         │             │<──未重复───│          │            │         │
 │         │             │────────────────publish>│           │         │
 │         │<──202 ──────│            │          │            │         │
 │         │             │            │          │─consume───>│         │
 │         │             │            │          │            │─INSERT─>│
 │─轮询───>│             │            │          │            │<─OK─────│
 │         │──GET ───────>│           │          │            │         │
 │<─成功───│             │            │          │            │         │
```

---

## 6. 数据需求（Data Requirements）

### 6.1 核心数据表

**表 1：用户表（users）**
| 字段名 | 类型 | 约束 | 说明 |
|---|---|---|---|
| id | BIGINT | PK, AUTO_INCREMENT | 主键 |
| student_id | VARCHAR(20) | UNIQUE, NOT NULL | 学号 |
| username | VARCHAR(50) | NOT NULL | 显示名称 |
| password_hash | VARCHAR(255) | NOT NULL | Bcrypt 密码 |
| role | ENUM('STUDENT','ADMIN') | DEFAULT 'STUDENT' | 角色 |
| status | ENUM('ACTIVE','BANNED') | DEFAULT 'ACTIVE' | 账号状态 |
| created_at | DATETIME | NOT NULL | 注册时间 |

**表 2：活动表（activities）**
| 字段名 | 类型 | 约束 | 说明 |
|---|---|---|---|
| id | BIGINT | PK | 主键 |
| title | VARCHAR(200) | NOT NULL | 活动标题 |
| description | TEXT | NOT NULL | 活动简介 |
| category_tags | VARCHAR(500) | NOT NULL | JSON 数组格式 |
| organizer | VARCHAR(100) | NOT NULL | 主讲人/主办方 |
| max_capacity | INT | NOT NULL | 总名额上限 |
| enroll_open_at | DATETIME | NOT NULL | 报名开放时间 |
| enroll_close_at | DATETIME | NOT NULL | 报名截止时间 |
| activity_at | DATETIME | NOT NULL | 活动举办时间 |
| status | ENUM('PUBLISHED','OFFLINE') | DEFAULT 'PUBLISHED' | 活动状态 |
| created_by | BIGINT | FK(users.id) | 创建管理员 |
| created_at | DATETIME | NOT NULL | 创建时间 |

**表 3：报名记录表（enrollments）**
| 字段名 | 类型 | 约束 | 说明 |
|---|---|---|---|
| id | BIGINT | PK | 主键 |
| user_id | BIGINT | FK(users.id), NOT NULL | 报名用户 |
| activity_id | BIGINT | FK(activities.id), NOT NULL | 目标活动 |
| status | ENUM('PENDING','SUCCESS','FAIL') | DEFAULT 'PENDING' | 报名状态 |
| created_at | DATETIME | NOT NULL | 提交时间 |
| updated_at | DATETIME | | 状态更新时间 |
| **UNIQUE KEY** | | (user_id, activity_id) | 防重复 |

**表 4：用户行为流水表（user_behaviors）**
> 此表为"大量的小数据"的核心载体，只增不删。

| 字段名 | 类型 | 约束 | 说明 |
|---|---|---|---|
| id | BIGINT | PK | 主键 |
| user_id | BIGINT | FK, NOT NULL | 行为用户 |
| activity_id | BIGINT | FK, NOT NULL | 涉及活动 |
| behavior_type | ENUM('VIEW','COLLECT','ENROLL') | NOT NULL | 行为类型 |
| created_at | DATETIME | NOT NULL | 发生时间 |
| INDEX | | (user_id, behavior_type, created_at) | 支持推荐查询 |

**表 5：活动热度快照表（activity_scores）**
| 字段名 | 类型 | 约束 | 说明 |
|---|---|---|---|
| id | BIGINT | PK | 主键 |
| activity_id | BIGINT | FK, UNIQUE | 活动外键 |
| score | DECIMAL(10,4) | NOT NULL | 综合热度分 |
| view_count | BIGINT | DEFAULT 0 | 浏览次数 |
| enroll_count | INT | DEFAULT 0 | 报名人数 |
| collect_count | INT | DEFAULT 0 | 收藏次数 |
| updated_at | DATETIME | NOT NULL | 最近计算时间 |

---

## 7. 接口需求（Interface Requirements）

### 7.1 RESTful API 规范
- **基础路径**：`/api/v1`
- **认证方式**：`Authorization: Bearer <JWT Token>`
- **数据格式**：`application/json`
- **统一响应结构**：`{ "code": 200, "message": "success", "data": {} }`

### 7.2 核心接口清单

| 方法 | 路径 | 描述 | 权限 |
|---|---|---|---|
| POST | `/api/v1/auth/register` | 用户注册 | 公开 |
| POST | `/api/v1/auth/login` | 用户登录 | 公开 |
| POST | `/api/v1/auth/logout` | 用户登出 | 已登录 |
| GET | `/api/v1/activities` | 获取活动列表（支持筛选、排序、分页） | 公开 |
| GET | `/api/v1/activities/{id}` | 获取活动详情 | 公开 |
| POST | `/api/v1/activities` | 创建活动 | 管理员 |
| PUT | `/api/v1/activities/{id}` | 编辑活动 | 管理员 |
| PATCH | `/api/v1/activities/{id}/offline` | 下架活动 | 管理员 |
| GET | `/api/v1/activities/trending` | 热门活动榜单 Top 10 | 公开 |
| POST | `/api/v1/activities/{id}/enroll` | 发起报名（抢票） | 已登录学生 |
| GET | `/api/v1/enroll/result/{taskId}` | 查询报名异步结果 | 已登录 |
| GET | `/api/v1/me/enrollments` | 我的报名记录 | 已登录 |
| POST | `/api/v1/activities/{id}/collect` | 收藏/取消收藏 | 已登录 |
| GET | `/api/v1/me/collections` | 我的收藏列表 | 已登录 |
| GET | `/api/v1/recommend` | 个性化推荐列表 | 已登录 |

---

## 8. 迭代与交付规划（Iteration & Delivery Plan）

### 第一轮迭代（Alpha）：工程基础管道

**目标**：完成用户注册/登录，贯通全部工程基础设施（CI/CD + Docker + Code Review）。

| 任务编号 | 负责组 | 任务描述 |
|---|---|---|
| ALPHA-BE-01 | 后端组 | 建立仓库，初始化项目骨架 |
| ALPHA-BE-02 | 后端组 | 建立 users 表（含数据库迁移脚本） |
| ALPHA-BE-03 | 后端组 | 实现注册接口（含 Bcrypt、参数校验、单元测试） |
| ALPHA-BE-04 | 后端组 | 实现登录接口（含 JWT 签发、单元测试） |
| ALPHA-BE-05 | 后端组 | 编写 Dockerfile，配置 CI/CD 流水线 |
| ALPHA-FE-01 | 前端组 | 建立仓库，初始化 Vue 3 / React + Vite 骨架 |
| ALPHA-FE-02 | 前端组 | 使用 Apifox 定义注册/登录接口 Mock 数据 |
| ALPHA-FE-03 | 前端组 | 实现注册页 View + ViewModel（对接 Mock） |
| ALPHA-FE-04 | 前端组 | 实现登录页 View + ViewModel（对接 Mock） |
| ALPHA-FE-05 | 前端组 | 配置 Vite 构建脚本，编写 Selenium UI 自动化测试 |

### 第二轮迭代（Beta）：核心业务功能

**目标**：完成活动管理、高并发抢票核心链路及报名查询。

关键任务：引入 Redis 和 RabbitMQ/Kafka；实现防超卖 Lua 脚本；实现消费者 Worker；前后端联调完整抢票流程；进行初次压力测试。

### 第三轮迭代（RC）：智能化功能

**目标**：完成热度算法定时任务、用户标签离线任务及个性化推荐接口。

关键任务：引入定时任务调度（Spring Scheduler / cron）；实现热度 Sorted Set 写入 Redis；实现用户标签权重计算；实现余弦相似度推荐接口。

### 最终验收标准

| 验收项 | 通过标准 |
|---|---|
| 功能验收 | Alpha/Beta/RC 所有 Issue 均已 Close，无高优先级遗留 Bug |
| 性能验收 | JMeter 1000 并发压测报告，P99 < 500ms，数据库不超额 |
| 安全验收 | SQL 注入、XSS 简要测试通过 |
| 工程流程验收 | 所有 Commit 带工单编号；所有主干合并经 PR Review；CI/CD 日志可查 |
| 文档验收 | API 文档（Apifox/Swagger）完整；数据库 ER 图完整 |

---

## 9. 附录（Appendix）

### 附录 A：课程原文要求对照

| 课程原文要求（ULSS-26-0） | 本系统的落地方式 |
|---|---|
| 大量的小数据 | `user_behaviors` 行为流水表，每次用户操作产生一条记录 |
| 大规模的各种算法 | 热度加权评分算法（TRENDING-01）；用户兴趣标签协同过滤（RECOMMEND-01/02） |
| 大量的交互 | 高并发报名模块（ENROLL-02/03）；JMeter 千并发压测验证 |

### 附录 B：技术选型汇总

| 层次 | 推荐技术 |
|---|---|
| 前端框架 | Vue 3 + Vite（或 React + Vite） |
| 后端框架 | Spring Boot 3 (Java) 或 Gin (Go) |
| 数据库 | MySQL 8 / PostgreSQL 15 |
| 缓存 | Redis 7 |
| 消息队列 | RabbitMQ 3.x 或 Kafka 3.x |
| 容器化 | Docker + Docker Compose |
| CI/CD | GitHub Actions 或 Gitee CI |
| API 文档 | Apifox 或 Swagger/OpenAPI 3 |
| 压力测试 | Apache JMeter |
| UI 自动化测试 | Selenium |


### 6.2 实体关系图 (ERD)
# 校园活动聚合分发平台 (CAAD) - 数据实体关系图 (ERD)

依据需求规格说明书 (SRS)，本系统核心数据实体及其关系如下：

```mermaid
erDiagram
    USER ||--o{ ENROLLMENT : signs_up
    USER ||--o{ USER_BEHAVIOR : performs
    USER ||--o{ ACTIVITY : creates
    ACTIVITY ||--o{ ENROLLMENT : contains
    ACTIVITY ||--o{ USER_BEHAVIOR : involves
    ACTIVITY ||--o{ ACTIVITY_SCORE : has_snapshot

    USER {
        bigint id PK
        varchar student_id UK "学号"
        varchar username "显示名称"
        varchar password_hash "加密密码"
        enum role "STUDENT/ADMIN"
        enum status "ACTIVE/BANNED"
        datetime created_at
    }

    ACTIVITY {
        bigint id PK
        varchar title "标题"
        text description "简介"
        json category_tags "标签数组"
        varchar organizer "举办方"
        int max_capacity "总名额"
        datetime enroll_open_at "开始报名时间"
        datetime enroll_close_at "截止报名时间"
        datetime activity_at "活动举行时间"
        enum status "PUBLISHED/OFFLINE"
        bigint created_by FK "管理员ID"
        datetime created_at
    }

    ENROLLMENT {
        bigint id PK
        bigint user_id FK "用户ID"
        bigint activity_id FK "活动ID"
        enum status "PENDING/SUCCESS/FAIL"
        datetime created_at
        datetime updated_at
    }

    USER_BEHAVIOR {
        bigint id PK
        bigint user_id FK "用户ID"
        bigint activity_id FK "活动ID"
        enum behavior_type "VIEW/COLLECT/ENROLL"
        datetime created_at
    }

    ACTIVITY_SCORE {
        bigint id PK
        bigint activity_id FK UK "活动ID"
        decimal score "总热度分"
        bigint view_count "累计浏览"
        int enroll_count "累计报名"
        int collect_count "累计收藏"
        datetime updated_at
    }
```

## 设计要点说明：

1.  **行为流水表 (USER_BEHAVIOR)**：作为“大量的小数据”的主要承载。所有用户操作只增不删，为后期推荐算法提供训练集。
2.  **库存一致性**：`ENROLLMENT` 表与 `ACTIVITY` 表通过应用层 (Redis) 逻辑同步，保证名额不超卖。
3.  **索引优化建议**：并在 `USER_BEHAVIOR` 的 `(user_id, created_at)` 以及 `ENROLLMENT` 的 `(activity_id, status)` 建立复合索引。

