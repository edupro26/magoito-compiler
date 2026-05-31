# MAGOITO M3 Report

## 1. How to build the compiler

A `Dockerfile` and a run script `run.sh` are provided. The script
builds a Docker image and immediately starts an interactive container:

```bash
./run.sh
```

You can also do the two steps separately:

```bash
docker build -t magoito .
docker run -it --rm magoito
```

## 2. How to run the whole test suite (valid and invalid programs)

Inside the Magoito CLI, run:

```bash
magoito test
```

The test command prints a summary of passed and failed tests per group.

## 3. How to find out how many tests passed and which tests failed

To see failing tests, run:

This will execute all tests and print a summary of passed and failed tests per group. To see details of failing tests, run:

```bash
magoito test --show-failed
```

You can use the flags `--all` to show all failed tests, or `--max N` to limit the number of shown failed tests to N.

## 4. How to run the compiler on a given MAGOITO program

```bash
magoito run path/to/program.mag
```

You can also run one of the provided examples:

```bash
magoito run --example 1
```

To print the AST of a program, use the `--ast` flag:

## 5. Brief description of the symbol table implementation

The symbol table is a chained structure with a `parent` pointer and a map of name to symbol. Each symbol stores a name, type, and kind (CONST, FUN, VAR). The validator creates a root table for top-level declarations, then extends it for nested scopes (function bodies and `var` scopes).

## 6. Brief description of bidirectional type checking

The implementation uses two mutually recursive functions: `inferType` synthesizes a type for expressions, while `checkType` validates an expression against an expected type.

## Notes

- A significant portion of time for this milestone was spent on migrating the previous hand written parser implementation to a ANTLR generated one. Due to this fact, the semantic analysis and type checking implementation is still in progress, despite already passing the majority of test cases therea are a lot of sections in the source code that will be revise and refactored.

- AI (GPT-4.2-Codex) was used to assist in the development; however, AI generated code was reviewed by the author to ensure correctness and adherence, with the exception of the semantics validation that as explained above still needs improvement and still contains code that as not been fully reviewed.
