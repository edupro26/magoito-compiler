# MAGOITO M4 Report

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

**Important**: Do not forget to specify `-it` in the `docker run` command, otherwise you won't be able to interact with the CLI, since the container is meant to be used in interactive mode.

## 2. How to run the whole test suite (valid and invalid programs)

Inside the Magoito CLI, run:

```bash
magoito test
```

The test command prints a summary of passed and failed tests per group.

## 3. How to find out how many tests passed and which tests failed

To see failing tests, run:

```bash
magoito test --show-failed
```

You can use the flags `--all` to show all failed tests, or `--max N` to limit the number of shown failed tests to N.

## 4. How to run the compiler on a given MAGOITO program

To run the compiler on a MAGOITO program, first compile it to LLVM IR.

```bash
magoito [file.mag]
```

Then run the generated LLVM IR with `lli`:

```bash
lli [file.ll]
```

Alternatively, you can run the whole process in one step with:

```bash
magoito run [file.mag]
```

## 5. A brief description of how you implemented short-circuit boolean expressions

Short-circuit is lowered into explicit control-flow. For `&&` we branch to the RHS only when LHS is true; for `||` we branch to the RHS only when LHS is false. The result is a `phi` in a merge block that selects either the RHS value or the short-circuit constant (`false` for `&&`, `true` for `||`).

## Notes

- The code generator is built with `github.com/llir/llvm`.

- GPT-5.2-Codex was used to assist in the development
