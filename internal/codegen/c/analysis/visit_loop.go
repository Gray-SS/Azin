package analysis

import (
	"github.com/azin-lang/Azin/pkg/ast"
)

func (a *Analyzer) visitLoop(fn string, stmt *ast.WhileStmt) {
	for _, child := range stmt.Body {
		a.visitStmt(fn, child)
	}
}
