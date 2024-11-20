package otto

import (
	"github.com/YiYuhki/otto/ast"
	"github.com/YiYuhki/otto/file"
)

type compiler struct {
	file    *file.File
	program *ast.Program
}
