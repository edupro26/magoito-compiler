<div align="center">

# MAGOITO Compiler

#### Compilation Techniques - Faculty of Sciences of University of Lisbon

[![Go](https://img.shields.io/badge/Built_with-Go-blue?logo=go)](https://go.dev/)
![GitHub repo size](https://img.shields.io/github/repo-size/ThreaditApp/Threadit)

</div>

## 📖 Overview

This repository contains the compiler's source code for the **MAGOITO** programming language.

MAGOITO is an imperative programming language composed of expressions
alone. A program in MAGOITO is a sequence of declarations, each announced with
keyword `const` or `fun`. Value declarations introduce an identifier and a type.

```
-- "Hello, World!" in MAGOITO

const hello : String = "Hello, World!"
fun main (_) : Unit -> Unit =
    print(hello)
```

More information about the MAGOITO language can be found in the [language specification](docs/magoito.md).

## 🔍 Implementation milestones

- [M1](docs/milestones/magoito-M1.md) - Creation of a test suite.
- [M2](docs/milestones/magoito-M2.md) - AST + Parser implementation.
- [M3](docs/milestones/magoito-M3.md) - Semantic analysis.
- [M4](magoito-M4.md) - Code generation.
