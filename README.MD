---

# 🚀 Go 语言基础入门课程

[![Go Version](https://img.shields.io/badge/Go-1.18%2B-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-Welcome-brightgreen.svg)](CONTRIBUTING.md)

欢迎来到 Go 语言基础入门课程！本项目是一个通过 **丰富的代码示例** 和 **详尽的中文讲解** 相结合的方式来系统学习 Google Go 编程语言基础的教程仓库。无论你是编程新手，还是有其他语言经验的开发者，都能从这里扎实地迈出 Go 语言编程的第一步。

## 📖 课程简介

本课程旨在提供一种高效、直观的学习方式。我们坚信“代码是最好的老师”，因此每个核心概念都配有可运行的示例代码 (`*.go` 文件)。你可以直接阅读代码旁的注释，也可以查看对应的讲解文档，通过理论与实践相结合的方式，深刻理解 Go 语言的精髓。

通过本课程，你将系统地掌握 Go 语言的核心语法和编程思想，为后续进行 Web 开发、微服务、云计算、命令行工具等领域的开发打下坚实的基础。

## 🧭 课程内容概览

本课程将涵盖以下 Go 语言基础知识模块：

| 模块 | 内容描述 | 代码示例 |
| :--- | :--- | :--- |
| **1. Go 语言概览** | 介绍 Go 的诞生背景、核心特点（如静态强类型、并发模型、垃圾回收、编译速度快）以及开发环境搭建（安装、配置 `GOROOT`/`GOPATH`、Hello World）。 | [`01-hello-world`](./01-hello-world/) |
| **2. 基本数据类型** | 学习 Go 的内建基本类型：`int`, `float64`, `string`, `bool` 等，以及变量的声明与初始化（`var` vs `:=`）。 | [`02-basic-types`](./02-basic-types/) |
| **3. 程序控制结构** | 掌握控制程序流程的语句：条件分支 `if`/`else`、`switch`，循环 `for`（Go 只有一种循环语句），以及 `break`, `continue`, `goto` 的使用。 | [`03-control-flow`](./03-control-flow/) |
| **4. 引用类型详解** | 深入理解 Go 中三种重要的引用类型：**切片（Slice）**（动态数组）、**映射（Map）**（键值对集合）、**通道（Channel）**（goroutine 间通信）。 | [`04-reference-types`](./04-reference-types/) |
| **5. 函数（Function）** | 学习如何定义和调用函数、多返回值、命名返回值、可变参数函数、函数作为值和类型。 | [`05-functions`](./05-functions/) |
| **6. 方法（Method）与面向对象** | 探索 Go 独特的“面向对象”编程方式：如何为类型定义方法、值接收者 vs 指针接收者、以及如何通过组合（Embedding）实现类似继承的效果。 | [`06-methods-oop`](./06-methods-oop/) |
| **7. 错误处理** | 了解 Go 的错误处理哲学，如何使用 `error` 接口和 `panic`/`recover` 机制。 | [`07-error-handling`](./07-error-handling/) |
| **(附加) 并发初探** | 简要介绍 Go 最强大的特性：Goroutine 和 Channel 的基本用法，体验“用 Go 的方式”编写并发程序。 | [`08-concurrency`](./08-concurrency/) |

## 🛠️ 开始学习

### 先决条件

- **Go 环境**：请确保你的电脑上已安装 **Go 1.18 或更高版本**。可以从 [官方下载页面](https://golang.org/dl/) 获取。
- **代码编辑器**：推荐使用 VS Code（配合 Go 插件）或 GoLand 等带有 Go 语言支持的编辑器。

### 如何使用本仓库

1.  **克隆仓库到本地**：
    ```bash
    git clone https://github.com/Zqwater/golang_primer.git
    cd golang_primer
    ```

2.  **按顺序学习**：建议从 `01-hello-world` 目录开始，依次进入各个模块的目录。
3.  **阅读与运行**：
    - 阅读每个目录下的 `README.md` 文件获取概念讲解。
    - 查看并运行目录中的 `.go` 文件，观察输出结果。
    - 尝试修改代码中的参数或逻辑，加深理解。

4.  **动手实践**：每个模块后都鼓励你创建自己的 `playground.go` 文件，编写代码来练习刚学到的知识。

## 🤝 贡献

我们非常欢迎和鼓励任何形式的贡献！

- **发现问题？** 欢迎提交 [Issue](https://github.com/your-username/your-golang-basic-course/issues)。
- **有改进建议？** 或者想补充更多例子？欢迎 Fork 本仓库并提交 [Pull Request](https://github.com/your-username/your-golang-basic-course/pulls)。
- 请参阅 [贡献指南](./CONTRIBUTING.md) 了解详情。

## 📜 许可证

本项目采用 [MIT 许可证](./LICENSE)。你可以自由地使用本项目的代码进行学习和创作。

## 🙏 致谢

感谢 Go 语言开源社区提供的优秀文档和资源。

---

**Happy Coding! Gopher! 🦡**

---