---
description: Backend AI-TDD: Concurrency Testing and Automated Iteration
---

# Backend AI-TDD: Concurrency Testing and Automated Iteration (skill-aitdd-backend)

In this module, functional correctness under high concurrency is essential. The strategy is to utilize load testing to identify bottlenecks and iteratively improve the system.

## 1. Test-First Strategy for High Concurrency (Hardcore TDD Sandbox)
- **Test Generation**: If no load testing scripts exist for the target module, please begin by writing a concurrency test (e.g., spawning 1000 `Goroutines` or using `JMeter/HTTP` clients to stress the target endpoint).
- **Diagnostics**: Run the test to collect any resulting stack traces, such as database deadlocks or oversold inventory errors. This log will serve as the baseline for your refactoring.

## 2. Performance and Optimization (Performance Under Extreme Constraint)
- **Optimized Locking**: Avoid relying solely on heavy pessimistic database locks (`SELECT ... FOR UPDATE`) in high-traffic scenarios, as this can exhaust the connection pool.
- **Architectural Upgrades**: Implement more efficient atomicity mechanisms, such as **Redis Lua scripts** for transaction isolation. Ensure that the logic processes quickly in the distributed cache layer and relies on asynchronous message queues for persistence.

## 3. Automated Execution Loop (The Self-Healing Death Loop)
- // turbo-all
- After completing the Redis/DB layer refactoring, follow this iterative process: `Run Load Test -> Analyze Error Logs -> Refactor Code -> Re-run Load Test`.
- Continue this autonomous debugging cycle until all concurrency-related exceptions are resolved.
- Output requirement: Include the terminal output showing a fully passed test execution (e.g., 1000 requests with 0 errors) in your final delivery notes.
