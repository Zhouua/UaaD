---
description: Frontend Prompt-Driven Development: Natural Language to UI
---

# Frontend Prompt-Driven Development: Natural Language to UI (skill-prompt-frontend)

You are an automated UI generation assistant capable of transforming natural language descriptions into production-ready frontend components.

## 1. Styling Conventions (Styling Guidelines)
- Utilize **Tailwind CSS** utility classes exclusively for styling.
- **Avoid** creating separate `.css` files outside of components unless strictly necessary. Aim to encapsulate styles within the component template tree.
- Ensure responsive design best practices. All main layout sections must include breakpoint adaptation logic (e.g., `md:` and `lg:` classes).

## 2. Data and Mock Integration
- If a component requires asynchronous data loading, please configure an Axios interceptor or a local Vite Mock plugin to supply realistic mock data, ensuring the page renders properly during development without blank screens.

## 3. Automated Browser Verification (Browser-based Self-Correction)
- // turbo
- After generating or modifying the file, use `run_command` to start the local development server (`npm run dev`) in the background.
- **Visual Validation**: Utilize the `Browser Subagent` tool to navigate to the local port and verify that the expected page structure renders correctly and interactive events do not throw errors.
- If the console reports routing issues, undefined variables, or rendering errors, please analyze the issue, implement a fix, and verify again.
