package transform

import (
	"github.com/prisma/prisma-client-go/generator/ast/dmmf"
)

type AST struct {
	dmmf *dmmf.Document

	// Scalars describe a list of scalar types, such as Int, String, DateTime, etc.
	Scalars []string `json:"scalars"`

	// Enums contains all enums
	Enums []Enum `json:"enums"`

	// ReadFilters describe a list of scalar types and the respective read operations
	ReadFilters []Filter `json:"readFilters"`

	// WriteFilters describe a list of scalar types and the respective read operations
	WriteFilters []Filter `json:"writeFilters"`
}

func New(document *dmmf.Document) *AST {
	ast := &AST{
		dmmf: document,
	}

	// first, fetch types
	ast.Scalars = ast.scalars()
	ast.Enums = ast.enums()

	// fetch data which is needed for the query api, which require ast types
	ast.ReadFilters = ast.readFilters()
	ast.WriteFilters = ast.writeFilters()

	return ast
}

func (r *AST) pick(name string) *dmmf.CoreType {
	for _, i := range r.dmmf.Schema.InputObjectTypes.Prisma {
		if string(i.Name) == name {
			return &i
		}
	}
	return nil
}