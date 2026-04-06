package printer

import (
	"fmt"
	"magoito-compiler/internal/ast"
	"strings"
)

// Print a indented tree representation
// of a AST Magoito Program to stdout.
//
//	Program
//	├── FunDecl: main
//	│   ├── Params: [_]
//	│   ├── Type: Unit -> Unit
//	│   └── Body
//	│       └── Sequence
//	│           ├── VarDecl: i : Int
//	│           │   └── IntLiteral: 0
//	│           └── ...
func Print(p *ast.Program) {
	fmt.Println("Program")
	for i, decl := range p.Declarations {
		last := i == len(p.Declarations)-1
		printDecl(decl, "", last)
	}
}

const (
	tee    = "├── "
	corner = "└── "
	pipe   = "│   "
	blank  = "    "
)

func branch(last bool) string {
	if last {
		return corner
	}
	return tee
}

func childPrefix(last bool) string {
	if last {
		return blank
	}
	return pipe
}

func printDecl(d ast.Declaration, prefix string, last bool) {
	switch n := d.(type) {
	case *ast.ConstDecl:
		fmt.Printf("%s%sConstDecl: %s : %s\n", prefix, branch(last), n.Name, printType(n.Type))
		printExpr(n.Value, prefix+childPrefix(last), true)
	case *ast.FunDecl:
		fmt.Printf("%s%sFunDecl: %s\n", prefix, branch(last), n.Name)
		p := prefix + childPrefix(last)
		fmt.Printf("%s%sParams: (%s)\n", p, tee, strings.Join(n.Params, ", "))
		fmt.Printf("%s%sType: %s\n", p, tee, printType(n.Type))
		fmt.Printf("%s%sBody\n", p, corner)
		printExpr(n.Body, p+blank, true)
	}
}

func printType(t ast.Type) string {
	switch n := t.(type) {
	case *ast.BasicType:
		return n.Name
	case *ast.RecordType:
		if len(n.Fields) == 0 {
			return "{}"
		}
		parts := make([]string, len(n.Fields))
		for i, f := range n.Fields {
			parts[i] = f.Label + ": " + printType(f.Type)
		}
		return "{" + strings.Join(parts, ", ") + "}"
	case *ast.FunctionType:
		params := make([]string, len(n.Params))
		for i, pt := range n.Params {
			params[i] = printType(pt)
		}
		paramStr := strings.Join(params, ", ")
		if len(n.Params) > 1 {
			paramStr = "(" + paramStr + ")"
		}
		return paramStr + " -> " + printType(n.Return)
	}
	return "?"
}

func printExpr(e ast.Expr, prefix string, last bool) {
	pfx := prefix + branch(last)
	childPfx := prefix + childPrefix(last)

	switch n := e.(type) {
	case *ast.IntLiteral:
		fmt.Printf("%sIntLiteral: %d\n", pfx, n.Value)
	case *ast.StringLiteral:
		fmt.Printf("%sStringLiteral: %q\n", pfx, n.Value)
	case *ast.BoolLiteral:
		fmt.Printf("%sBoolLiteral: %t\n", pfx, n.Value)
	case *ast.UnitLiteral:
		fmt.Printf("%sUnitLiteral\n", pfx)
	case *ast.Identifier:
		fmt.Printf("%sIdentifier: %s\n", pfx, n.Name)
	case *ast.BinaryExpr:
		fmt.Printf("%sBinaryExpr: %s\n", pfx, n.Op)
		printExpr(n.Left, childPfx, false)
		printExpr(n.Right, childPfx, true)
	case *ast.UnaryExpr:
		fmt.Printf("%sUnaryExpr: %s\n", pfx, n.Op)
		printExpr(n.Operand, childPfx, true)
	case *ast.SequenceExpr:
		fmt.Printf("%sSequence\n", pfx)
		printExpr(n.Left, childPfx, false)
		printExpr(n.Right, childPfx, true)
	case *ast.IfExpr:
		fmt.Printf("%sIfExpr\n", pfx)
		fmt.Printf("%s%sCondition\n", childPfx, tee)
		printExpr(n.Condition, childPfx+pipe, true)
		if n.Else != nil {
			fmt.Printf("%s%sThen\n", childPfx, tee)
			printExpr(n.Then, childPfx+pipe, true)
			fmt.Printf("%s%sElse\n", childPfx, corner)
			printExpr(n.Else, childPfx+blank, true)
		} else {
			fmt.Printf("%s%sThen\n", childPfx, corner)
			printExpr(n.Then, childPfx+blank, true)
		}
	case *ast.WhileExpr:
		fmt.Printf("%sWhileExpr\n", pfx)
		fmt.Printf("%s%sCondition\n", childPfx, tee)
		printExpr(n.Condition, childPfx+pipe, true)
		fmt.Printf("%s%sBody\n", childPfx, corner)
		printExpr(n.Body, childPfx+blank, true)
	case *ast.VarDecl:
		fmt.Printf("%sVarDecl: %s : %s\n", pfx, n.Name, printType(n.Type))
		printExpr(n.Value, childPfx, true)
	case *ast.Assignment:
		fmt.Printf("%sAssignment: %s\n", pfx, n.Name)
		printExpr(n.Value, childPfx, true)
	case *ast.CallExpr:
		fmt.Printf("%sCallExpr: %s\n", pfx, n.Callee)
		for i, arg := range n.Args {
			printExpr(arg, childPfx, i == len(n.Args)-1)
		}
	case *ast.RecordExpr:
		fmt.Printf("%sRecordExpr\n", pfx)
		for i, f := range n.Fields {
			fieldLast := i == len(n.Fields)-1
			fmt.Printf("%s%sField: %s\n", childPfx, branch(fieldLast), f.Label)
			printExpr(f.Value, childPfx+childPrefix(fieldLast), true)
		}
	case *ast.ProjectionExpr:
		fmt.Printf("%sProjection: .%s\n", pfx, n.Field)
		printExpr(n.Record, childPfx, true)
	default:
		fmt.Printf("%s<unknown expr>\n", pfx)
	}
}
