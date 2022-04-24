package consts

const TemplateGenServiceContent = `
// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package {PackageName}

type I{StructName} interface {
	{FuncDefinition}
}

var local{StructName} I{StructName}

func {StructName}() I{StructName} {
	return local{StructName}
}

func Register{StructName}(i I{StructName}) {
	local{StructName} = i
}
`
