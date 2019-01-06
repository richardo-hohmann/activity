// Package RFC contains ontology values that are defined in RFCs, BCPs, and
// other miscellaneous standards.
package rfc

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/codegen"
	"github.com/cjslep/activity/tools/exp/rdf"
	"github.com/dave/jennifer/jen"
	"net/url"
	"strings"
)

const (
	rfcSpec   = "https://tools.ietf.org/html/"
	bcp47Spec = "bcp47"
)

type RFCOntology struct {
	Package string
}

func (o *RFCOntology) SpecURI() string {
	return rfcSpec
}

func (o *RFCOntology) Load() ([]rdf.RDFNode, error) {
	return o.LoadAsAlias("")
}

func (o *RFCOntology) LoadAsAlias(s string) ([]rdf.RDFNode, error) {
	return []rdf.RDFNode{
		&rdf.AliasedDelegate{
			Spec:     rfcSpec,
			Alias:    s,
			Name:     bcp47Spec,
			Delegate: &bcp47{pkg: o.Package},
		},
	}, nil
}

func (o *RFCOntology) LoadSpecificAsAlias(alias, name string) ([]rdf.RDFNode, error) {
	switch name {
	case bcp47Spec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &bcp47{pkg: o.Package},
			},
		}, nil
	}
	return nil, fmt.Errorf("rfc ontology cannot find %q to alias to %q", name, alias)
}

func (o *RFCOntology) LoadElement(name string, payload map[string]interface{}) ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *RFCOntology) GetByName(name string) (rdf.RDFNode, error) {
	name = strings.TrimPrefix(name, o.SpecURI())
	switch name {
	case bcp47Spec:
		return &bcp47{pkg: o.Package}, nil
	}
	return nil, fmt.Errorf("rfc ontology could not find node for name %s", name)
}

var _ rdf.RDFNode = &bcp47{}

type bcp47 struct {
	pkg string
}

func (b *bcp47) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("bcp47 langaugetag cannot be entered")
}

func (b *bcp47) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("bcp47 languagetag cannot be exited")
}

func (b *bcp47) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	v := ctx.Result.GetReference(rfcSpec)
	if len(v.Values[bcp47Spec].Name) == 0 {
		u, err := url.Parse(rfcSpec + bcp47Spec)
		if err != nil {
			return true, err
		}
		val := &rdf.VocabularyValue{
			Name:           bcp47Spec,
			URI:            u,
			DefinitionType: jen.String(),
			Zero:           "\"\"",
			IsNilable:      false,
			SerializeFn: rdf.SerializeValueFunction(
				b.pkg,
				bcp47Spec,
				jen.String(),
				[]jen.Code{
					jen.Return(
						jen.Id(codegen.This()),
						jen.Nil(),
					),
				}),
			DeserializeFn: rdf.DeserializeValueFunction(
				b.pkg,
				bcp47Spec,
				jen.String(),
				[]jen.Code{
					jen.If(
						jen.List(
							jen.Id("s"),
							jen.Id("ok"),
						).Op(":=").Id(codegen.This()).Assert(jen.String()),
						jen.Id("ok"),
					).Block(
						jen.Return(
							jen.Id("s"),
							jen.Nil(),
						),
					).Else().Block(
						jen.Return(
							jen.Lit(""),
							jen.Qual("fmt", "Errorf").Call(
								jen.Lit("%v cannot be interpreted as a string for bcp47 languagetag"),
								jen.Id(codegen.This()),
							),
						),
					),
				}),
			LessFn: rdf.LessFunction(
				b.pkg,
				bcp47Spec,
				jen.String(),
				[]jen.Code{
					jen.Return(
						jen.Id("lhs").Op("<").Id("rhs"),
					),
				}),
		}
		if err = v.SetValue(bcp47Spec, val); err != nil {
			return true, err
		}
	}
	return true, nil
}
