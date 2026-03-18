---
description: Architecture Group: Database Schema and Entity Mapping
---

# Architecture Group: Database Schema and Entity Mapping (skill-sdd-arch)

Your goal is to accurately translate ER diagrams provided in Markdown format into physical database tables and ORM models, adhering strictly to the documented design.

## 1. Strict Schema Alignment (Structure Override)
- **Baseline Alignment**: Treat the entity descriptions in `docs/SRS.md` as the definitive baseline. The generated Entity/Model code must achieve 100% field and data type mapping with the documentation.
- **Avoid Undocumented Features**: Do not introduce undocumented bidirectional relationships or complex underlying caching wrappers unless explicitly authorized in the SRS.

## 2. Automated Migration Generation (Auto-Migration)
- When the table structure evolves, do not only refactor the code files. You **must** use `run_command` to generate the corresponding migration scripts (e.g., `Flyway` V-series SQL or `golang-migrate`'s `.up.sql/.down.sql` files).
- If possible, automatically run a dry-run migration command to ensure the SQL syntax is flawless.

## 3. Verification Requirements
- // turbo
- Upon completion, please use shell scripts or automated tools to output a DDL Diff list representing the schema changes. This serves as the primary verification for this workflow.
