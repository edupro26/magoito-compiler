package semantics

import "magoito-compiler/internal/ast"

func containsFunctionType(t ast.Type) bool {
	switch typ := t.(type) {
	case *ast.FunctionType:
		return true
	case *ast.RecordType:
		for _, field := range typ.Fields {
			if containsFunctionType(field.Type) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func recordFieldMap(rt *ast.RecordType) map[string]ast.Type {
	m := make(map[string]ast.Type, len(rt.Fields))
	for _, f := range rt.Fields {
		m[f.Label] = f.Type
	}
	return m
}

func formatType(t ast.Type) string {
	if t == nil {
		return "?"
	}
	switch tt := t.(type) {
	case *ast.BasicType:
		return tt.Name
	case *ast.RecordType:
		if len(tt.Fields) == 0 {
			return "{}"
		}
		parts := make([]string, len(tt.Fields))
		for i, f := range tt.Fields {
			parts[i] = f.Label + ": " + formatType(f.Type)
		}
		return "{" + joinStrings(parts) + "}"
	case *ast.FunctionType:
		params := make([]string, len(tt.Params))
		for i, p := range tt.Params {
			params[i] = formatType(p)
		}
		paramStr := joinStrings(params)
		if len(tt.Params) > 1 {
			paramStr = "(" + paramStr + ")"
		}
		return paramStr + " -> " + formatType(tt.Return)
	default:
		return "?"
	}
}

func joinStrings(parts []string) string {
	if len(parts) == 0 {
		return ""
	}
	res := parts[0]
	for i := 1; i < len(parts); i++ {
		res += ", " + parts[i]
	}
	return res
}
