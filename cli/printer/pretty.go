package printer

import (
	"fmt"
	"magoito-compiler/internal/ast"
	"strings"
)

// Pretty prints a source representation of a Magoito Program.
//
//	fun f (n) : Int -> Int =
//	    if ((n % 2) == 0) then
//	        (n / 2)
//	    else
//	        ((3 * n) + 1)
func Pretty(p *ast.Program) {
	for i, decl := range p.Declarations {
		if i > 0 {
			fmt.Println()
		}
		fmt.Println(prettyDecl(decl))
	}
}

func prettyDecl(d ast.Declaration) string {
	switch n := d.(type) {
	case *ast.ConstDecl:
		return fmt.Sprintf("const %s : %s = %s", n.Name, prettyType(n.Type), prettyExpr(n.Value, ""))
	case *ast.FunDecl:
		params := strings.Join(n.Params, ", ")
		body := prettyExpr(n.Body, "    ")
		return fmt.Sprintf("fun %s (%s) : %s =\n    %s", n.Name, params, prettyType(n.Type), body)
	}
	return ""
}

func prettyType(t ast.Type) string {
	switch n := t.(type) {
	case *ast.BasicType:
		return n.Name
	case *ast.RecordType:
		if len(n.Fields) == 0 {
			return "{}"
		}
		parts := make([]string, len(n.Fields))
		for i, f := range n.Fields {
			parts[i] = f.Label + ": " + prettyType(f.Type)
		}
		return "{" + strings.Join(parts, ", ") + "}"
	case *ast.FunctionType:
		params := make([]string, len(n.Params))
		for i, pt := range n.Params {
			params[i] = prettyType(pt)
		}
		paramStr := strings.Join(params, ", ")
		if len(n.Params) > 1 {
			paramStr = "(" + paramStr + ")"
		}
		return paramStr + " -> " + prettyType(n.Return)
	}
	return "?"
}

func prettyExpr(e ast.Expr, indent string) string {
	switch n := e.(type) {
	case *ast.IntLiteral:
		return fmt.Sprintf("%d", n.Value)
	case *ast.StringLiteral:
		return fmt.Sprintf("%q", n.Value)
	case *ast.BoolLiteral:
		return fmt.Sprintf("%t", n.Value)
	case *ast.UnitLiteral:
		return "unit"
	case *ast.Identifier:
		return n.Name
	case *ast.BinaryExpr:
		return fmt.Sprintf("(%s %s %s)", prettyExpr(n.Left, indent), n.Op, prettyExpr(n.Right, indent))
	case *ast.UnaryExpr:
		return fmt.Sprintf("%s%s", n.Op, prettyExpr(n.Operand, indent))
	case *ast.SequenceExpr:
		parts := flattenSeq(e, indent)
		return strings.Join(parts, ";\n"+indent)
	case *ast.IfExpr:
		cond := prettyExpr(n.Condition, indent)
		inner := indent + "    "
		then := prettyExpr(n.Then, inner)
		if n.Else != nil {
			els := prettyExpr(n.Else, inner)
			return fmt.Sprintf("if %s then\n%s%s\n%selse\n%s%s",
				cond, inner, then, indent, inner, els)
		}
		return fmt.Sprintf("if %s then\n%s%s", cond, inner, then)
	case *ast.WhileExpr:
		cond := prettyExpr(n.Condition, indent)
		if _, ok := n.Body.(*ast.SequenceExpr); ok {
			inner := indent + "    "
			body := prettyExpr(n.Body, inner)
			return fmt.Sprintf("while %s do (\n%s%s\n%s)", cond, inner, body, indent)
		}
		return fmt.Sprintf("while %s do %s", cond, prettyExpr(n.Body, indent))
	case *ast.VarDecl:
		return fmt.Sprintf("var %s : %s = %s", n.Name, prettyType(n.Type), prettyExpr(n.Value, indent))
	case *ast.Assignment:
		return fmt.Sprintf("%s = %s", n.Name, prettyExpr(n.Value, indent))
	case *ast.CallExpr:
		args := make([]string, len(n.Args))
		for i, arg := range n.Args {
			args[i] = prettyExpr(arg, indent)
		}
		return fmt.Sprintf("%s(%s)", n.Callee, strings.Join(args, ", "))
	case *ast.RecordExpr:
		fields := make([]string, len(n.Fields))
		for i, f := range n.Fields {
			fields[i] = fmt.Sprintf("%s = %s", f.Label, prettyExpr(f.Value, indent))
		}
		return "{" + strings.Join(fields, ", ") + "}"
	case *ast.ProjectionExpr:
		return fmt.Sprintf("%s.%s", prettyExpr(n.Record, indent), n.Field)
	}
	return "<unknown expr>"
}

// flattenSeq recursively collects all elements of a SequenceExpr into a flat slice.
func flattenSeq(e ast.Expr, indent string) []string {
	if seq, ok := e.(*ast.SequenceExpr); ok {
		return append(flattenSeq(seq.Left, indent), flattenSeq(seq.Right, indent)...)
	}
	return []string{prettyExpr(e, indent)}
}
