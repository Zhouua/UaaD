# CAAD 项目：团队贡献与 AI-Native 协作指南 (CONTRIBUTING)

欢迎加入大规模信息系统开发试验（CAAD）项目。本项目采用前沿的 10人团队融合 AI 智能体 的混合开发范式。为了保证高复杂系统在自动化交付中的规范性和一致性，请全员参阅并遵守本指南。

---

## 1. 核心方法论：混合 AI 开发范式 (The Hybrid Paradigm)
为兼顾效率与系统质量，我们将团队的方法论逻辑规划为三条相互配合的智能体辅助主线：

- **架构组 (SDD 驱动)**：使用 `/skill-sdd-arch` 技能，通过明确修改 `docs/SRS.md` 来指导 AI 生成数据库模型，避免脱离设计文档的代码变更。
- **前端组 (Prompt 驱动)**：使用 `/skill-prompt-frontend` 技能，通过自然语言 `.md` 提示词文件维护意图，辅助 AI 自动生成 Vue/React 组件并完成视觉辅助验证。
- **后端组 (AI-TDD 驱动)**：使用 `/skill-aitdd-backend` 技能，由人类工程师编写高并发负载测试（例如 1000 并发压测），并利用 AI 持续重构和优化 Redis Lua 扣减逻辑，形成“测试-反馈-修复”闭环，以确保测试最终通过。

---

## 2. 代码开发与合并流水线 (The Pipeline)

### 2.1 Git 分支与命名
- **主干保护**：`main` 分支已被锁定，直接 Push 权限受限，需通过 PR 流程。
- **分支命名**：建议围绕 Issue ID 命名（例如 `feature/auth-#3`）。
- **提交规范**：每次 Commit 建议包含意图与编号，如 `feat(backend): 实现并发扣减逻辑 #14`。

### 2.2 PR 审核门槛 (PR Approval Process)
在发起 Pull Request 时，请确保满足以下常规工程标准：
1. 必须使用 `.github/PULL_REQUEST_TEMPLATE.md` 描述变更。
2. 需附带必要的端到端验证证据（如前端界面的正常渲染截图，或后端的压力测试 PASS 日志），建议统一记录于 `walkthrough.md`。
3. 必须通过 GitHub Actions 中配置的 Lint 和单元测试流水线验证。

---

## 3. 智能体工作守则 (Agent Guardrails)

如果你是在本地使用 Cursor、Copilot 或是任何大语言模型助手协助开发，根目录的 `.cursorrules` 文件将作为规则约束 AI，以确保：
- AI 能够紧贴 `SRS.md` 中的实际需求，避免凭空虚构功能。
- 在涉及 API 开发后，激励 AI 自动在 `/tmp` 输出测试脚本并以 HTTP 200/201 为基础验收标准。

---

## 4. 快速上手 (Quick Start)
如果您是刚入职的工程师，请按以下步骤准备开发环境：
1. 仔细阅读核心分析文档：`docs/SRS.md`（包含业务逻辑、API 契约和 ER 数据图）。
2. 在看板 `task_list.md` 中领取分配给您的任务工单。
3. 激活工作流，并应用 `.agents/workflows/` 下对应的团队专属技能文档，开始协同开发！
