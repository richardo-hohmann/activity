package exp

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

// TODO: Deserialize & generated structs preserve "unknown" elements.
// TODO: Make Serialize/Deserialize/KindIndex funcs that return jen.Code
// TODO: Document and comment this code.
// TODO: Split out methods as structs.

const (
	getMethod                 = "Get"
	setMethod                 = "Set"
	hasMethod                 = "Has"
	clearMethod               = "Clear"
	iteratorClearMethod       = "clear"
	isMethod                  = "Is"
	appendMethod              = "Append"
	prependMethod             = "Prepend"
	removeMethod              = "Remove"
	lenMethod                 = "Len"
	swapMethod                = "Swap"
	lessMethod                = "Less"
	kindIndexMethod           = "kindIndex"
	serializeMethod           = "Serialize"
	deserializeMethod         = "Deserialize"
	nameMethod                = "Name"
	serializeIteratorMethod   = "serialize"
	deserializeIteratorMethod = "deserialize"
)

func join(s []*jen.Statement) *jen.Statement {
	return joinLines(s, true)
}

func joinSingleLine(s []*jen.Statement) *jen.Statement {
	return joinLines(s, false)
}

func joinLines(s []*jen.Statement, double bool) *jen.Statement {
	r := &jen.Statement{}
	for i, stmt := range s {
		if i > 0 {
			r.Line()
			if double {
				r.Line()
			}
		}
		r.Add(stmt.Clone())
	}
	return r
}

type Identifier struct {
	LowerName string
	CamelName string
}

type Kind struct {
	Name                  Identifier
	ConcreteKind          string
	Nilable               bool
	HasNaturalLanguageMap bool
	SerializeFnName       string
	DeserializeFnName     string
	LessFnName            string
}

type PropertyGenerator struct {
	Name       Identifier
	Kinds      []Kind
	asIterator bool
}

func (p *PropertyGenerator) structName() string {
	if p.asIterator {
		return p.Name.CamelName
	}
	return fmt.Sprintf("%sProperty", p.Name.CamelName)
}

func (p *PropertyGenerator) propertyName() string {
	return p.Name.LowerName
}

func (p *PropertyGenerator) deserializeFnName() string {
	if p.asIterator {
		return fmt.Sprintf("%s%s", deserializeIteratorMethod, p.Name.CamelName)
	}
	return fmt.Sprintf("%s%sProperty", deserializeMethod, p.Name.CamelName)
}

func (p *PropertyGenerator) getFnName(i int) string {
	if len(p.Kinds) == 1 {
		return getMethod
	}
	return fmt.Sprintf("%s%s", getMethod, p.kindCamelName(i))
}

func (p *PropertyGenerator) serializeFnName() string {
	if p.asIterator {
		return serializeIteratorMethod
	}
	return serializeMethod
}

func (p *PropertyGenerator) kindCamelName(i int) string {
	return p.Kinds[i].Name.CamelName
}

func (p *PropertyGenerator) memberName(i int) string {
	return fmt.Sprintf("%sMember", p.Kinds[i].Name.LowerName)
}

func (p *PropertyGenerator) hasMemberName(i int) string {
	if len(p.Kinds) == 1 && p.Kinds[0].Nilable {
		panic("PropertyGenerator.hasMemberName called for nilable single value")
	}
	return fmt.Sprintf("has%sMember", p.Kinds[i].Name.CamelName)
}

func (p *PropertyGenerator) clearMethodName() string {
	if p.asIterator {
		return iteratorClearMethod
	}
	return clearMethod
}

func (p *PropertyGenerator) commonFuncs() jen.Code {
	memberFunc := jen.Func().Params(
		jen.Id("t").Id(p.structName()),
	)
	return jen.Commentf(
		"%s returns the name of this property: %q.", nameMethod, p.propertyName(),
	).Line().Add(memberFunc.Clone().Id(nameMethod).Params().Params(
		jen.String(),
	).Block(
		jen.Return(
			jen.Lit(p.propertyName()),
		),
	))
}

func (p *PropertyGenerator) isMethodName(i int) string {
	return fmt.Sprintf("%s%s", isMethod, p.kindCamelName(i))
}
