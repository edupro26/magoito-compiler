## 🔍 Milestone 2

### 1. How to Build and Run the Compiler

A `Dockerfile` and a convenience script `run.sh` are provided. The script
builds a Docker image and immediately starts an interactive container:

```bash
./run.sh
```

You can also do the two steps separately:

```bash
docker build -t magoito .
docker run --rm -it magoito
```

### 2. Shift-Reduce and Reduce-Reduce Conflicts

This compiler uses a **handwritten recursive-descent parser** and does not make use of
any lexer/parser generator tool (e.g. yacc, ANTLR, etc.), the notions of shift-reduce
and reduce-reduce conflicts do not apply. Ambiguities are instead resolved directly in
code through parsing order and look-ahead logic.

### 3. How to Run the Compiler on a MAGOITO Program

Start the **CLI** via Docker.

```
Magoito CLI (early version), Go version: go1.25.9
Try "help" for more information, or run an example: mag run --example 1
>>> 
```

Inside the **CLI** prompt (`>>>`):

```
# Run a specific .mag source file
>>> mag run path/to/program.mag

# Optionally display the AST of the program
>>> mag run --ast path/to/program.mag

# Run one of the built-in examples (1, 2, or 3)
>>> mag run --example 1
```

Type `help` to see all available commands, or `quit` / `q` to exit.

> **\*Note:** any `.mag` file can be run, as long as you provide the 
> correct path, and it is located inside the project directory (or subdirectories).

### 4. How to execute the Test Suite

Inside the **CLI** prompt, run:

```
>>> mag test
```

This automatically discovers and runs every `.mag` file under the test directory:

Each test case prints one line:

```
<path/to/program.mag> - ✅
<path/to/program.mag> - ❌
```

If the result is ✅, the test passed, meaning that the program was successfully parsed.
if you see ❌, the test failed, which means that the parser produced a syntax error.