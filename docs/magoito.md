## MAGOITO by example

**MAGOITO** is an imperative programming language composed of expressions alone.

A program in **MAGOITO** is a sequence of declarations, each announced with keyword `const` or `fun`. 
Value declarations introduce an identifier and a type. For example:

``` const maxSpeed : Int = 120 ```

``` const fever : Int = 37 + 5 ```

Declarations may also introduce functions, as in:

``` fun succ (n) : Int -> Int = n + 1 ```

Functions can be recursive:

```
fun add (n, m) : (Int, Int) -> Int = 
    if n == 0 then m else succ(add (n - 1, m))
```

or even mutually recursive:

```
fun even (x) : Int -> Bool = 
    x == 0 || odd(x - 1)
```

```
fun odd (x) : Int -> Bool = 
    x != 0 && even(x - 1)
```

Imperative means state changing. How do we change state? By means of (imperative) variables and assignment. Here’s addition in imperative style:

```
fun addI (n, m) : (Int, Int) -> Int =
    var sum : Int = m;
    while n > 0 do (
        sum = sum + 1;
        n = n - 1
    );
    sum 
```

A few points worth of notice:

- The body of function `varI` is composed of one expression only
- This expression is composed of two expressions, *separated* by a semicolon (the second of such expression is again a semicolon expression);
- The value of the function is given by the last expression in the semicolon separated sequence, namely, `sum`;
- The body of the `while` expression is a sequence of two expressions; we enclose them in parentheses.

Each expression — no exception — has a value. This includes `while` loops, variable declarations and assignments. The value of `while` is `unit`, 
the only value of type `Unit`. This means that we can “store” a while loop in a variable, as in

```var x : Unit = while i > 0 do i = i - 1```

or “return” a while loop from a function, as in

```fun f (n) : Int -> Unit = while n > 0 do n = n - 1```

A variable declaration or an assignment stores a value in a variable and returns the value. Given the below function declaration, a call to `f(unit)`
yields value 3.

```fun f (_) : Unit -> Int = var x : Int = var y : Int = 3```

Equipped with the `Unit` value, an if-then expression `if` exp1 `then` exp is an abbreviation for an if-then-else `if` exp1 `then` exp2 `else unit`.

Let us now look at the support for records. A record is a (possibly empty) sequence of fields; each field is composed of a label (that is, an 
identifier) and an expression. For example, `{user = "tcomp", id = 0 + 0 * 0}` is a record with two fields, the first named `user` with value 
`"tcomp"`, second named `id` with expression `0 + 0 * 0`.

The syntax of records types is similar; this time we use a colon rather than an equals sign: `{user: String, id: Int}`. In any case, record 
expressions or record types, the order of the fields is irrelevant, so that `{user = "tcomp", id: 0 + 0 * 0}` and `{id: 0, user: "tcomp"}` denote
the same value, and `{user: String, id: Int}` and `{id: Int, user: String}` denote the same type.

The elements in a record can be obtained via projection. For example `{user = "tcomp", id: 0 + 0 * 0}`. `id` evaluates to `0`. Here’s a program that
manipulates 2D points and prints number 3.

`const origin : {x: Int, y: Int} = {x = 0, y = 0}`

```
fun move (p, deltax, deltay) : ({x: Int, y: Int}, Int, Int) -> {y: Int, x: Int} =
    {x = p.x + deltax, y = p.y + deltay}
```

```
fun main (_) : Unit -> Unit =
    print (move(origin, 3, 5).x)
```

Records can be nested, as in 
```
{point: {x: Int, y: Int}, colour: String}
```

## The syntax of **MAGOITO**, informally

### Lexing

- Identifiers (variable and function names, record labels) start with a letter and are followed by zero or more letters, digits, underscore 
symbols **(_)** and single quotes **(')**
- Integer values are either 0 or start with an optional sign, followed by a non-zero digit and then followed by zero or more digits
- Strings are sequences of characters (not including new line), enclosed in quotes **(")**
- Comments, line comments only, start with **--**

### Wildcards

- Wildcards may appear only in binding positions:
    ```
    var _ : Unit = 17; unit
    fun` f (_) : Int -> Int = 5
    const _ : Bool = true
    ```
- Hence, wildcards may not appear in expressions: `2 * _`
- Wildcards may appear more than once, even if with different types:
    ```
    fun f (_, _) : (Unit, Bool) -> Int = 5
    ```

### Expressions

- Variables are identifiers
- Literals: `-2147483648, ..., -1, 0, 1, ..., 2147483647`, `true`, `false`, `Unit`, in addition to quote-enclosed string literals
- Binary operators: `;` `+` `-` `*` `/` `%` `^` `==` `!=` `<` `<=` `>` `>=` `!` `||` `&&`
- Unary operators: `-` `!`
- Function call: `id(exp1, ... ,expn)` with n ≥ 1
- Assignment: `id = exp`
- Variable declaration: `var id : type = exp`
- Conditional: `if exp1 then exp2 else exp3 and if exp1 then exp`
- While loop: `while exp1 do exp`
- Record construction: `{id1 = exp1, ..., idn = expn}` with n ≥ 0
- Record projection: `exp.id`
- Sequential composition: `exp1 ; exp2`
- Parenthetical expression: `(exp)`

### Operator precedence and associativity 

- Binary operators `+` `-` `*` `/` `%` `==` `!=` `<` `<=` `>` `%` `>=` `!` `||` `&&` `.` take the precedence and associativity of the 
Java programming language
- The unary minus binds tighter than any other arithmetic operator
- The power operator, `^`, is right associative and binds tighter than any other arithmetic operator (different from unary minus). For example,
`2 ^ - 3 ^ 4 * 5` is to be understood as `(2 ^ ((-3) ^ 4)) * 5`
- The sequencing operator, `;`, associates to the right and binds loser than all other operators. For example, `1 ; 2 || 3 ; 4` is to be understood
as `1 ; ((2 || 3) ; 4)` , and `var x = 4 ; 5 ` means `(var x = 4) ; 5`
- The record projection operator associates to the left and binds tighter than any other binary operator
- The precedence of `while` loops and conditionals (both `if-then-else` and `if-then`) seats between that of `;` and `||`. For example, 
`while b do false || true` should be understood as `while b do (false || true)`, but `while b do false ; true` should be understood as 
`(while b do false) ; true`. Similarly, for conditionals
- Keywords `then` and `else` associate to the right, so that `if a then if b then c else d` is to be understood as `if a then (if b then c else d)`
- The arrow type operator, `->`, associates to the right, so that `Bool -> Int -> String` means `Bool -> (Int -> String)`

### Top-level declarations

- Constants: `const : type = exp`
- Functions: `fun id (id1,...,idn) : type = exp` with n ≥ 1

Functions are *not* first class values: they cannot be stored in variables or in records, neither can they be passed as parameters.

### Programs 

A non-empty sequence of declarations.

### Types

- Basic: `Int` , `Bool` , `Unit` , `String`
- Record: `{id1: type1,...,idn: typen}` with n ≥ 0
- Function: `type -> type` or `(type1,..., typen) -> type` with n ≥ 2. There are no zero-ary functions. If needed, use `f (_) : Unit -> type = exp`
and call `f(unit)`

# Validation

### Top-level declarations

There is a “small” restriction on `const` declaration:

- Expression `exp` in a declaration `const id : type = exp` must be made solely of literals, operators, record expressions and projections, in such
a way that it can be evaluated at compile time. 

The validation rule for programs is as follows. A program is valid if:

- The identifiers of all its (top level) declarations are pairwise distinct and
- All its (top level) declarations are valid and
- There is a function with signature `fun main (x) : Unit -> Unit`

### Typing 

The types for expressions are as follows.

- The type of a literal is the corresponding type. For example, the type of `5` is `Int`
- The type of call `print(exp1,...,expn)` is `Unit` if `n = 1` and `exp` has a type (any type)
- The type of call `id(exp1,...,expn)` (with `id` different from `print`) is `type` if `id` is of type `(type1,...,typem) ->type` and `n=m ` and 
each `expi` has type `typei`
- The type of variable declaration `var id : type = exp` is `type` if `exp` is of type `type`. Furthermore, if the declaration appears at the left 
of a semicolon `var id : type = exp1 ; exp2`, then the type of `id` is used to validate occurrences of `id` in `exp`
- The type of an identifier is that more recently introduced by a `var` expression or a `const` declaration
- The type of conditional `if exp1 then exp2 else exp3` is `type` if `exp1` is of type `Bool` and both `exp2` and `exp3` are of type `type`. In the
case of `if exp1 then exp2`, expression `exp2` must be of type `Unit`
- The type of while loop `while exp1 do exp2` is `Unit` if `exp1` is of type `Bool` and `exp2` has a type (any type)
- The type of assignment `id = e` is type if `id` and `e` share the same type `type`
- The type of a record `{id1 = exp1, ...,idn = expn}` is `{id1: type1, ...,idn: typen}` if each `expi` is of type `typei`
- The type of record projection `exp.idk` is `typek` if `exp` is of type `{id1: type1, ...,idn: typen}`
- The type of sequential composition `exp1; exp2` is the type of `exp2` if `exp1` has a type (any type)
- The type of the remaining operators (binary and unary) are as customary in programming languages

The rules for top-level declarations are as follows.

- Constant declaration `const id: type = e` is valid if `e` is of type `type`
- Function declaration `fun id (id1,...,idn) : type = exp` is valid if `type` is of the form `(type1,...,typem)->type'` and `n=m` and `exp` is of 
type `type'` in a context augmented with `id:type,id1:type1,...,idn:typen`
- No variable or function may be named `print`
- Function declarations may be (mutually) recursive
- Functions and constants cannot be declared twice, even if with different signatures

# The behaviour of programs

How are programs evaluated?

A **MAGOITO** program executes starting from evaluating all `const` declarations. Then a call to function `main(unit)` is performed. Here’s a few pointers.

- Function arguments are evaluated left-to-right until they all become values. The corresponding function is then called (**MAGOITO** is call-by-value)
- Most binary expressions `exp1` `op` `exp2` evaluate `exp1` and `exp2` (in this order) and then apply the corresponding operator
- The exception are boolean expressions. These are evaluated in a *short-circuit* manner: they are evaluated left-to-right until the truth value of the
expression can be determined. In particular, in expression `exp1 && exp2`, if `exp1` turns out to be false, then `exp2` is *not* evaluated. Dually for disjunction
- Variable declaration and function parameters introduce storage to hold values of the corresponding type
- For non-function types, the value in the store may be read via the identifier expression `id` and updated via expression `id = exp`
- The value of `var id = exp` and that of `id = exp` is the value of `exp`
- A conditional expression `if exp1 then exp2 else exp3` evaluates to `exp2` or to `exp3` according to the truth value of expression `exp1`. The value
of the conditional is the value of `exp1` or `exp2`
- A `while exp1 do exp2` loop evaluates `exp2` while `exp1` remains true. The value of the loop is `unit`
- The expressions in a record `{id1 = exp1, ...,idn = expn}` are evaluated left to right
- Record projection `exp.idk` first evaluates `exp` to obtain a record `{id1 = value1, ...,idn = valuen}`; then the value of the projection is `valuek`