package xsd

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/codegen"
	"github.com/cjslep/activity/tools/exp/rdf"
	"github.com/dave/jennifer/jen"
	"net/url"
	"strings"
)

const (
	xmlSpec                = "http://www.w3.org/2001/XMLSchema#"
	anyURISpec             = "anyURI"
	dateTimeSpec           = "dateTime"
	floatSpec              = "float"
	stringSpec             = "string"
	booleanSpec            = "boolean"
	nonNegativeIntegerSpec = "nonNegativeInteger"
	durationSpec           = "duration"
)

type XMLOntology struct {
	Package string
}

func (o *XMLOntology) SpecURI() string {
	return xmlSpec
}

func (o *XMLOntology) Load() ([]rdf.RDFNode, error) {
	return o.LoadAsAlias("")
}

func (o *XMLOntology) LoadAsAlias(s string) ([]rdf.RDFNode, error) {
	return []rdf.RDFNode{
		&rdf.AliasedDelegate{
			Spec:     xmlSpec,
			Alias:    s,
			Name:     anyURISpec,
			Delegate: &anyURI{pkg: o.Package},
		},
		&rdf.AliasedDelegate{
			Spec:     xmlSpec,
			Alias:    s,
			Name:     dateTimeSpec,
			Delegate: &dateTime{pkg: o.Package},
		},
		&rdf.AliasedDelegate{
			Spec:     xmlSpec,
			Alias:    s,
			Name:     floatSpec,
			Delegate: &float{pkg: o.Package},
		},
		&rdf.AliasedDelegate{
			Spec:     xmlSpec,
			Alias:    s,
			Name:     stringSpec,
			Delegate: &xmlString{pkg: o.Package},
		},
		&rdf.AliasedDelegate{
			Spec:     xmlSpec,
			Alias:    s,
			Name:     booleanSpec,
			Delegate: &boolean{pkg: o.Package},
		},
		&rdf.AliasedDelegate{
			Spec:     xmlSpec,
			Alias:    s,
			Name:     nonNegativeIntegerSpec,
			Delegate: &nonNegativeInteger{pkg: o.Package},
		},
		&rdf.AliasedDelegate{
			Spec:     xmlSpec,
			Alias:    s,
			Name:     durationSpec,
			Delegate: &duration{pkg: o.Package},
		},
	}, nil
}

func (o *XMLOntology) LoadSpecificAsAlias(alias, name string) ([]rdf.RDFNode, error) {
	switch name {
	case anyURISpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &anyURI{pkg: o.Package},
			},
		}, nil
	case dateTimeSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &dateTime{pkg: o.Package},
			},
		}, nil
	case floatSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &float{pkg: o.Package},
			},
		}, nil
	case stringSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &xmlString{pkg: o.Package},
			},
		}, nil
	case booleanSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &boolean{pkg: o.Package},
			},
		}, nil
	case nonNegativeIntegerSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &nonNegativeInteger{pkg: o.Package},
			},
		}, nil
	case durationSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &duration{pkg: o.Package},
			},
		}, nil
	}
	return nil, fmt.Errorf("xml ontology cannot find %q to alias to %q", name, alias)
}

func (o *XMLOntology) LoadElement(name string, payload map[string]interface{}) ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *XMLOntology) GetByName(name string) (rdf.RDFNode, error) {
	name = strings.TrimPrefix(name, o.SpecURI())
	switch name {
	case anyURISpec:
		return &anyURI{pkg: o.Package}, nil
	case dateTimeSpec:
		return &dateTime{pkg: o.Package}, nil
	case floatSpec:
		return &float{pkg: o.Package}, nil
	case stringSpec:
		return &xmlString{pkg: o.Package}, nil
	case booleanSpec:
		return &boolean{pkg: o.Package}, nil
	case nonNegativeIntegerSpec:
		return &nonNegativeInteger{pkg: o.Package}, nil
	case durationSpec:
		return &duration{pkg: o.Package}, nil
	}
	return nil, fmt.Errorf("xsd ontology could not find node for name %s", name)
}

var _ rdf.RDFNode = &anyURI{}

type anyURI struct {
	pkg string
}

func (a *anyURI) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd anyURI cannot be entered")
}

func (a *anyURI) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd anyURI cannot be exited")
}

func (a *anyURI) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	v := ctx.Result.GetReference(xmlSpec)
	if len(v.Values[anyURISpec].Name) == 0 {
		u, err := url.Parse(xmlSpec + anyURISpec)
		if err != nil {
			return true, err
		}
		val := &rdf.VocabularyValue{
			Name:           anyURISpec,
			URI:            u,
			DefinitionType: "*url.URL",
			Zero:           "&url.URL{}",
			SerializeFn: rdf.SerializeValueFunction(
				a.pkg,
				anyURISpec,
				jen.Op("*").Qual("net/url", "URL"),
				[]jen.Code{
					jen.Return(
						jen.Id(codegen.This()).Dot("String").Call(),
						jen.Nil(),
					),
				}),
			DeserializeFn: rdf.DeserializeValueFunction(
				a.pkg,
				anyURISpec,
				jen.Op("*").Qual("net/url", "URL"),
				[]jen.Code{
					jen.Var().Id("u").Op("*").Qual("net/url", "URL"),
					jen.Var().Err().Error(),
					jen.If(
						jen.List(
							jen.Id("s"),
							jen.Id("ok"),
						).Op(":=").Id(codegen.This()).Assert(jen.String()),
						jen.Id("ok"),
					).Block(
						jen.List(
							jen.Id("u"),
							jen.Err(),
						).Op("=").Qual("net/url", "Parse").Call(jen.Id("s")),
						jen.If(
							jen.Err().Op("!=").Nil(),
						).Block(
							jen.Err().Op("=").Qual("fmt", "Errorf").Call(
								jen.Lit("%v cannot be interpreted as a xsd:anyURI: %s"),
								jen.Id(codegen.This()),
								jen.Err(),
							),
						),
					).Else().Block(
						jen.Err().Op("=").Qual("fmt", "Errorf").Call(
							jen.Lit("%v cannot be interpreted as a string for xsd:anyURI"),
							jen.Id(codegen.This()),
						),
					),
					jen.Return(jen.List(
						jen.Id("u"),
						jen.Err(),
					)),
				}),
			LessFn: rdf.LessFunction(
				a.pkg,
				anyURISpec,
				jen.Op("*").Qual("net/url", "URL"),
				[]jen.Code{
					jen.Return(
						jen.Id("lhs").Dot("String").Call().Op("<").Id("rhs").Dot("String").Call(),
					),
				}),
		}
		if err = v.SetValue(anyURISpec, val); err != nil {
			return true, err
		}
	}
	return true, nil
}

var _ rdf.RDFNode = &dateTime{}

type dateTime struct {
	pkg string
}

func (d *dateTime) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd dateTime cannot be entered")
}

func (d *dateTime) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd dateTime cannot be exited")
}

func (d *dateTime) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	v := ctx.Result.GetReference(xmlSpec)
	if len(v.Values[dateTimeSpec].Name) == 0 {
		u, err := url.Parse(xmlSpec + dateTimeSpec)
		if err != nil {
			return true, err
		}
		val := &rdf.VocabularyValue{
			Name:           dateTimeSpec,
			URI:            u,
			DefinitionType: "time.Time",
			Zero:           "&time.Time{}",
			SerializeFn: rdf.SerializeValueFunction(
				d.pkg,
				dateTimeSpec,
				jen.Qual("time", "Time"),
				[]jen.Code{
					jen.Return(
						jen.Id(codegen.This()).Dot("Format").Call(jen.Qual("time", "RFC3339")),
						jen.Nil(),
					),
				}),
			DeserializeFn: rdf.DeserializeValueFunction(
				d.pkg,
				dateTimeSpec,
				jen.Qual("time", "Time"),
				[]jen.Code{
					jen.Var().Id("tmp").Qual("time", "Time"),
					jen.Var().Err().Error(),
					jen.If(
						jen.List(
							jen.Id("s"),
							jen.Id("ok"),
						).Op(":=").Id(codegen.This()).Assert(jen.String()),
						jen.Id("ok"),
					).Block(
						jen.List(
							jen.Id("tmp"),
							jen.Err(),
						).Op("=").Qual("time", "Parse").Call(
							jen.Qual("time", "RFC3339"),
							jen.Id("s"),
						),
						jen.If(
							jen.Err().Op("!=").Nil(),
						).Block(
							jen.List(
								jen.Id("tmp"),
								jen.Err(),
							).Op("=").Qual("time", "Parse").Call(
								jen.Lit("2006-01-02T15:04Z07:00"),
								jen.Id("s"),
							),
							jen.If(
								jen.Err().Op("!=").Nil(),
							).Block(
								jen.Err().Op("=").Qual("fmt", "Errorf").Call(
									jen.Lit("%v cannot be interpreted as xsd:datetime"),
									jen.Id(codegen.This()),
								),
							),
						),
					).Else().Block(
						jen.Err().Op("=").Qual("fmt", "Errorf").Call(
							jen.Lit("%v cannot be interpreted as a string for xsd:datetime"),
							jen.Id(codegen.This()),
						),
					),
					jen.Return(jen.List(
						jen.Id("tmp"),
						jen.Err(),
					)),
				}),
			LessFn: rdf.LessFunction(
				d.pkg,
				dateTimeSpec,
				jen.Qual("time", "Time"),
				[]jen.Code{
					jen.Return(jen.Id("lhs").Dot("Before").Call(jen.Id("rhs"))),
				}),
		}
		if err = v.SetValue(dateTimeSpec, val); err != nil {
			return true, err
		}
	}
	return true, nil
}

var _ rdf.RDFNode = &float{}

type float struct {
	pkg string
}

func (f *float) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd float cannot be entered")
}

func (f *float) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd float cannot be exited")
}

func (f *float) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	v := ctx.Result.GetReference(xmlSpec)
	if len(v.Values[floatSpec].Name) == 0 {
		u, err := url.Parse(xmlSpec + floatSpec)
		if err != nil {
			return true, err
		}
		val := &rdf.VocabularyValue{
			Name:           floatSpec,
			URI:            u,
			DefinitionType: "float32",
			Zero:           "0.0",
			SerializeFn: rdf.SerializeValueFunction(
				f.pkg,
				floatSpec,
				jen.Id("float32"),
				[]jen.Code{
					jen.Return(
						jen.Id(codegen.This()),
						jen.Nil(),
					),
				}),
			DeserializeFn: rdf.DeserializeValueFunction(
				f.pkg,
				floatSpec,
				jen.Id("float32"),
				[]jen.Code{
					jen.If(
						jen.List(
							jen.Id("f"),
							jen.Id("ok"),
						).Op(":=").Id(codegen.This()).Assert(jen.Float32()),
						jen.Id("ok"),
					).Block(
						jen.Return(
							jen.Id("f"),
							jen.Nil(),
						),
					).Else().Block(
						jen.Return(
							jen.Lit(0),
							jen.Qual("fmt", "Errorf").Call(
								jen.Lit("%v cannot be interpreted as a float32 for xsd:float"),
								jen.Id(codegen.This()),
							),
						),
					),
				}),
			LessFn: rdf.LessFunction(
				f.pkg,
				floatSpec,
				jen.Id("float32"),
				[]jen.Code{
					jen.Return(
						jen.Id("lhs").Op("<").Id("rhs"),
					),
				}),
		}
		if err = v.SetValue(floatSpec, val); err != nil {
			return true, err
		}
	}
	return true, nil
}

var _ rdf.RDFNode = &xmlString{}

type xmlString struct {
	pkg string
}

func (*xmlString) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd string cannot be entered")
}

func (*xmlString) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd string cannot be exited")
}

func (s *xmlString) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	v := ctx.Result.GetReference(xmlSpec)
	if len(v.Values[stringSpec].Name) == 0 {
		u, err := url.Parse(xmlSpec + stringSpec)
		if err != nil {
			return true, err
		}
		val := &rdf.VocabularyValue{
			Name:           stringSpec,
			URI:            u,
			DefinitionType: "string",
			Zero:           "\"\"",
			SerializeFn: rdf.SerializeValueFunction(
				s.pkg,
				stringSpec,
				jen.Id("string"),
				[]jen.Code{
					jen.Return(
						jen.Id(codegen.This()),
						jen.Nil(),
					),
				}),
			DeserializeFn: rdf.DeserializeValueFunction(
				s.pkg,
				stringSpec,
				jen.Id("string"),
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
								jen.Lit("%v cannot be interpreted as a string for xsd:string"),
								jen.Id(codegen.This()),
							),
						),
					),
				}),
			LessFn: rdf.LessFunction(
				s.pkg,
				stringSpec,
				jen.Id("string"),
				[]jen.Code{
					jen.Return(
						jen.Id("lhs").Op("<").Id("rhs"),
					),
				}),
		}
		if err = v.SetValue(stringSpec, val); err != nil {
			return true, err
		}
	}
	return true, nil
}

var _ rdf.RDFNode = &boolean{}

type boolean struct {
	pkg string
}

func (*boolean) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd boolean cannot be entered")
}

func (*boolean) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd boolean cannot be exited")
}

func (b *boolean) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	v := ctx.Result.GetReference(xmlSpec)
	if len(v.Values[booleanSpec].Name) == 0 {
		u, err := url.Parse(xmlSpec + booleanSpec)
		if err != nil {
			return true, err
		}
		val := &rdf.VocabularyValue{
			Name:           booleanSpec,
			URI:            u,
			DefinitionType: "bool",
			Zero:           "false",
			SerializeFn: rdf.SerializeValueFunction(
				b.pkg,
				booleanSpec,
				jen.Id("bool"),
				[]jen.Code{
					jen.Return(
						jen.Id(codegen.This()),
						jen.Nil(),
					),
				}),
			DeserializeFn: rdf.DeserializeValueFunction(
				b.pkg,
				booleanSpec,
				jen.Id("bool"),
				[]jen.Code{
					jen.If(
						jen.List(
							jen.Id("b"),
							jen.Id("ok"),
						).Op(":=").Id(codegen.This()).Assert(jen.Bool()),
						jen.Id("ok"),
					).Block(
						jen.Return(
							jen.Id("b"),
							jen.Nil(),
						),
					).Else().If(
						jen.List(
							jen.Id("f"),
							jen.Id("ok"),
						).Op(":=").Id(codegen.This()).Assert(jen.Float64()),
						jen.Id("ok"),
					).Block(
						jen.If(
							jen.Id("f").Op("==").Lit(0),
						).Block(
							jen.Return(
								jen.False(),
								jen.Nil(),
							),
						).Else().If(
							jen.Id("f").Op("==").Lit(1),
						).Block(
							jen.Return(
								jen.True(),
								jen.Nil(),
							),
						).Else().Block(
							jen.Return(
								jen.False(),
								jen.Qual("fmt", "Errorf").Call(
									jen.Lit("%v cannot be interpreted as a bool float64 for xsd:boolean"),
									jen.Id(codegen.This()),
								),
							),
						),
					).Else().Block(
						jen.Return(
							jen.False(),
							jen.Qual("fmt", "Errorf").Call(
								jen.Lit("%v cannot be interpreted as a bool for xsd:boolean"),
								jen.Id(codegen.This()),
							),
						),
					),
				}),
			LessFn: rdf.LessFunction(
				b.pkg,
				booleanSpec,
				jen.Id("bool"),
				[]jen.Code{
					jen.Commentf("Booleans don't have a natural ordering, so we pick that truth is greater than falsehood."),
					jen.Return(
						jen.Op("!").Id("lhs").Op("&&").Id("rhs"),
					),
				}),
		}
		if err = v.SetValue(booleanSpec, val); err != nil {
			return true, err
		}
	}
	return true, nil
}

var _ rdf.RDFNode = &nonNegativeInteger{}

type nonNegativeInteger struct {
	pkg string
}

func (*nonNegativeInteger) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd nonNegativeInteger cannot be entered")
}

func (*nonNegativeInteger) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd nonNegativeInteger cannot be exited")
}

func (n *nonNegativeInteger) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	v := ctx.Result.GetReference(xmlSpec)
	if len(v.Values[nonNegativeIntegerSpec].Name) == 0 {
		u, err := url.Parse(xmlSpec + nonNegativeIntegerSpec)
		if err != nil {
			return true, err
		}
		val := &rdf.VocabularyValue{
			Name:           nonNegativeIntegerSpec,
			URI:            u,
			DefinitionType: "int",
			Zero:           "0",
			SerializeFn: rdf.SerializeValueFunction(
				n.pkg,
				nonNegativeIntegerSpec,
				jen.Id("int"),
				[]jen.Code{
					jen.Return(
						jen.Id(codegen.This()),
						jen.Nil(),
					),
				}),
			DeserializeFn: rdf.DeserializeValueFunction(
				n.pkg,
				nonNegativeIntegerSpec,
				jen.Id("int"),
				[]jen.Code{
					jen.If(
						jen.List(
							jen.Id("i"),
							jen.Id("ok"),
						).Op(":=").Id(codegen.This()).Assert(jen.Float64()),
						jen.Id("ok"),
					).Block(
						jen.Id("n").Op(":=").Int().Call(jen.Id("i")),
						jen.If(
							jen.Id("n").Op(">=").Lit(0),
						).Block(
							jen.Return(
								jen.Id("n"),
								jen.Nil(),
							),
						).Else().Block(
							jen.Return(
								jen.Lit(0),
								jen.Qual("fmt", "Errorf").Call(
									jen.Lit("%v is a negative integer for xsd:nonNegativeInteger"),
									jen.Id(codegen.This()),
								),
							),
						),
					).Else().Block(
						jen.Return(
							jen.Lit(0),
							jen.Qual("fmt", "Errorf").Call(
								jen.Lit("%v cannot be interpreted as a float for xsd:nonNegativeInteger"),
								jen.Id(codegen.This()),
							),
						),
					),
				}),
			LessFn: rdf.LessFunction(
				n.pkg,
				nonNegativeIntegerSpec,
				jen.Id("int"),
				[]jen.Code{
					jen.Return(
						jen.Id("lhs").Op("<").Id("rhs"),
					),
				}),
		}
		if err = v.SetValue(nonNegativeIntegerSpec, val); err != nil {
			return true, err
		}
	}
	return true, nil
}

var _ rdf.RDFNode = &duration{}

type duration struct {
	pkg string
}

func (*duration) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd duration cannot be entered")
}

func (*duration) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd duration cannot be exited")
}

func (d *duration) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	v := ctx.Result.GetReference(xmlSpec)
	if len(v.Values[durationSpec].Name) == 0 {
		u, err := url.Parse(xmlSpec + durationSpec)
		if err != nil {
			return true, err
		}
		val := &rdf.VocabularyValue{
			Name:           durationSpec,
			URI:            u,
			DefinitionType: "time.Duration",
			Zero:           "time.Duration(0)",
			SerializeFn: rdf.SerializeValueFunction(
				d.pkg,
				durationSpec,
				jen.Qual("time", "Duration"),
				[]jen.Code{
					jen.Commentf("Seriously questioning my life choices."),
					jen.Id("s").Op(":=").Lit("P"),
					jen.If(
						jen.Id(codegen.This()).Op("<").Lit(0),
					).Block(
						jen.Id("s").Op("=").Lit("-P"),
						jen.Id(codegen.This()).Op("=").Lit(-1).Op("*").Id(codegen.This()),
					),
					jen.Var().Id("tally").Qual("time", "Duration"),
					// Years
					jen.Commentf("Assume 8760 Hours per 365 days, cannot account for leap years in xsd:duration. :("),
					jen.If(
						jen.Id("years").Op(":=").Id(codegen.This()).Dot("Hours").Call().Op("/").Lit(8760.0),
						jen.Id("years").Op(">=").Lit(1),
					).Block(
						jen.Id("nYears").Op(":=").Int64().Call(
							jen.Qual("math", "Floor").Call(
								jen.Id("years"),
							),
						),
						jen.Id("tally").Op("+=").Qual("time", "Duration").Call(
							jen.Id("nYears"),
						).Op("*").Lit(8760).Op("*").Qual("time", "Hour"),
						jen.Id("s").Op("=").Qual("fmt", "Sprintf").Call(
							jen.Lit("%s%dY"),
							jen.Id("s"),
							jen.Id("nYears"),
						),
					),
					// Months
					jen.Commentf("Assume 30 days per month, cannot account for months lasting 31, 30, 29, or 28 days in xsd:duration. :("),
					jen.If(
						jen.Id("months").Op(":=").Parens(
							jen.Id(codegen.This()).Dot("Hours").Call().Op("-").Id("tally").Dot("Hours").Call(),
						).Op("/").Lit(720.0),
						jen.Id("months").Op(">=").Lit(1),
					).Block(
						jen.Id("nMonths").Op(":=").Int64().Call(
							jen.Qual("math", "Floor").Call(
								jen.Id("months"),
							),
						),
						jen.Id("tally").Op("+=").Qual("time", "Duration").Call(
							jen.Id("nMonths"),
						).Op("*").Lit(720).Op("*").Qual("time", "Hour"),
						jen.Id("s").Op("=").Qual("fmt", "Sprintf").Call(
							jen.Lit("%s%dM"),
							jen.Id("s"),
							jen.Id("nMonths"),
						),
					),
					// Days
					jen.If(
						jen.Id("days").Op(":=").Parens(
							jen.Id(codegen.This()).Dot("Hours").Call().Op("-").Id("tally").Dot("Hours").Call(),
						).Op("/").Lit(24.0),
						jen.Id("days").Op(">=").Lit(1),
					).Block(
						jen.Id("nDays").Op(":=").Int64().Call(
							jen.Qual("math", "Floor").Call(
								jen.Id("days"),
							),
						),
						jen.Id("tally").Op("+=").Qual("time", "Duration").Call(
							jen.Id("nDays"),
						).Op("*").Lit(24).Op("*").Qual("time", "Hour"),
						jen.Id("s").Op("=").Qual("fmt", "Sprintf").Call(
							jen.Lit("%s%dD"),
							jen.Id("s"),
							jen.Id("nDays"),
						),
					),
					jen.If(
						jen.Id("tally").Op("<").Id(codegen.This()),
					).Block(
						jen.Id("s").Op("=").Qual("fmt", "Sprintf").Call(
							jen.Lit("%sT"),
							jen.Id("s"),
						),
						// Hours
						jen.If(
							jen.Id("hours").Op(":=").Id(codegen.This()).Dot("Hours").Call().Op("-").Id("tally").Dot("Hours").Call(),
							jen.Id("hours").Op(">=").Lit(1),
						).Block(
							jen.Id("nHours").Op(":=").Int64().Call(
								jen.Qual("math", "Floor").Call(
									jen.Id("hours"),
								),
							),
							jen.Id("tally").Op("+=").Qual("time", "Duration").Call(
								jen.Id("nHours"),
							).Op("*").Qual("time", "Hour"),
							jen.Id("s").Op("=").Qual("fmt", "Sprintf").Call(
								jen.Lit("%s%dH"),
								jen.Id("s"),
								jen.Id("nHours"),
							),
						),
						// Minutes
						jen.If(
							jen.Id("minutes").Op(":=").Id(codegen.This()).Dot("Minutes").Call().Op("-").Id("tally").Dot("Minutes").Call(),
							jen.Id("minutes").Op(">=").Lit(1),
						).Block(
							jen.Id("nMinutes").Op(":=").Int64().Call(
								jen.Qual("math", "Floor").Call(
									jen.Id("minutes"),
								),
							),
							jen.Id("tally").Op("+=").Qual("time", "Duration").Call(
								jen.Id("nMinutes"),
							).Op("*").Qual("time", "Minute"),
							jen.Id("s").Op("=").Qual("fmt", "Sprintf").Call(
								jen.Lit("%s%dM"),
								jen.Id("s"),
								jen.Id("nMinutes"),
							),
						),
						// Seconds
						jen.If(
							jen.Id("seconds").Op(":=").Id(codegen.This()).Dot("Seconds").Call().Op("-").Id("tally").Dot("Seconds").Call(),
							jen.Id("seconds").Op(">=").Lit(1),
						).Block(
							jen.Id("nSeconds").Op(":=").Int64().Call(
								jen.Qual("math", "Floor").Call(
									jen.Id("seconds"),
								),
							),
							jen.Id("tally").Op("+=").Qual("time", "Duration").Call(
								jen.Id("nSeconds"),
							).Op("*").Qual("time", "Second"),
							jen.Id("s").Op("=").Qual("fmt", "Sprintf").Call(
								jen.Lit("%s%dS"),
								jen.Id("s"),
								jen.Id("nSeconds"),
							),
						),
					),
					jen.Return(
						jen.Id("s"),
						jen.Nil(),
					),
				}),
			DeserializeFn: rdf.DeserializeValueFunction(
				d.pkg,
				durationSpec,
				jen.Qual("time", "Duration"),
				[]jen.Code{
					jen.Commentf("Maybe this time it will be easier."),
					jen.If(
						jen.List(
							jen.Id("s"),
							jen.Id("ok"),
						).Op(":=").Id(codegen.This()).Assert(jen.String()),
						jen.Id("ok"),
					).Block(
						jen.Id("isNeg").Op(":=").False(),
						jen.If(
							jen.Id("s").Index(jen.Lit(0)).Op("==").LitRune('-'),
						).Block(
							jen.Id("isNeg").Op("=").True(),
							jen.Id("s").Op("=").Id("s").Index(jen.Lit(1), jen.Empty()),
						),
						jen.If(
							jen.Id("s").Index(jen.Lit(0)).Op("!=").LitRune('P'),
						).Block(
							jen.Return(
								jen.Lit(0),
								jen.Qual("fmt", "Errorf").Call(
									jen.Lit("%s malformed: missing 'P' for xsd:duration"),
									jen.Id("s"),
								),
							),
						),
						jen.Id("re").Op(":=").Qual("regexp", "MustCompile").Call(
							jen.Lit("P(\\d*Y)?(\\d*M)?(\\d*D)?(T(\\d*H)?(\\d*M)?(\\d*S)?)?"),
						),
						jen.Id("res").Op(":=").Id("re").Dot("FindStringSubmatch").Call(jen.Id("s")),
						jen.Var().Id("dur").Qual("time", "Duration"),
						// Years
						jen.Id("nYear").Op(":=").Id("res").Index(jen.Lit(1)),
						jen.If(
							jen.Len(
								jen.Id("nYear"),
							).Op(">").Lit(0),
						).Block(
							jen.Id("nYear").Op("=").Id("nYear").Index(
								jen.Empty(),
								jen.Len(
									jen.Id("nYear"),
								).Op("-").Lit(1),
							),
							jen.List(
								jen.Id("vYear"),
								jen.Err(),
							).Op(":=").Qual("strconv", "ParseInt").Call(
								jen.Id("nYear"),
								jen.Lit(10),
								jen.Lit(64),
							),
							jen.If(
								jen.Err().Op("!=").Nil(),
							).Block(
								jen.Return(
									jen.Lit(0),
									jen.Err(),
								),
							),
							jen.Commentf("Assume 8760 Hours per 365 days, cannot account for leap years in xsd:duration. :("),
							jen.Id("dur").Op("+=").Qual("time", "Duration").Call(
								jen.Id("vYear"),
							).Op("*").Qual("time", "Hour").Op("*").Lit(8760),
						),
						// Months
						jen.Id("nMonth").Op(":=").Id("res").Index(jen.Lit(2)),
						jen.If(
							jen.Len(
								jen.Id("nMonth"),
							).Op(">").Lit(0),
						).Block(
							jen.Id("nMonth").Op("=").Id("nMonth").Index(
								jen.Empty(),
								jen.Len(
									jen.Id("nMonth"),
								).Op("-").Lit(1),
							),
							jen.List(
								jen.Id("vMonth"),
								jen.Err(),
							).Op(":=").Qual("strconv", "ParseInt").Call(
								jen.Id("nMonth"),
								jen.Lit(10),
								jen.Lit(64),
							),
							jen.If(
								jen.Err().Op("!=").Nil(),
							).Block(
								jen.Return(
									jen.Lit(0),
									jen.Err(),
								),
							),
							jen.Commentf("Assume 30 days per month, cannot account for months lasting 31, 30, 29, or 28 days in xsd:duration. :("),
							jen.Id("dur").Op("+=").Qual("time", "Duration").Call(
								jen.Id("vMonth"),
							).Op("*").Qual("time", "Hour").Op("*").Lit(720),
						),
						// Days
						jen.Id("nDay").Op(":=").Id("res").Index(jen.Lit(3)),
						jen.If(
							jen.Len(
								jen.Id("nDay"),
							).Op(">").Lit(0),
						).Block(
							jen.Id("nDay").Op("=").Id("nDay").Index(
								jen.Empty(),
								jen.Len(
									jen.Id("nDay"),
								).Op("-").Lit(1),
							),
							jen.List(
								jen.Id("vDay"),
								jen.Err(),
							).Op(":=").Qual("strconv", "ParseInt").Call(
								jen.Id("nDay"),
								jen.Lit(10),
								jen.Lit(64),
							),
							jen.If(
								jen.Err().Op("!=").Nil(),
							).Block(
								jen.Return(
									jen.Lit(0),
									jen.Err(),
								),
							),
							jen.Id("dur").Op("+=").Qual("time", "Duration").Call(
								jen.Id("vDay"),
							).Op("*").Qual("time", "Hour").Op("*").Lit(24),
						),
						// Hours
						jen.Id("nHour").Op(":=").Id("res").Index(jen.Lit(5)),
						jen.If(
							jen.Len(
								jen.Id("nHour"),
							).Op(">").Lit(0),
						).Block(
							jen.Id("nHour").Op("=").Id("nHour").Index(
								jen.Empty(),
								jen.Len(
									jen.Id("nHour"),
								).Op("-").Lit(1),
							),
							jen.List(
								jen.Id("vHour"),
								jen.Err(),
							).Op(":=").Qual("strconv", "ParseInt").Call(
								jen.Id("nHour"),
								jen.Lit(10),
								jen.Lit(64),
							),
							jen.If(
								jen.Err().Op("!=").Nil(),
							).Block(
								jen.Return(
									jen.Lit(0),
									jen.Err(),
								),
							),
							jen.Id("dur").Op("+=").Qual("time", "Duration").Call(
								jen.Id("vHour"),
							).Op("*").Qual("time", "Hour"),
						),
						// Minutes
						jen.Id("nMinute").Op(":=").Id("res").Index(jen.Lit(6)),
						jen.If(
							jen.Len(
								jen.Id("nMinute"),
							).Op(">").Lit(0),
						).Block(
							jen.Id("nMinute").Op("=").Id("nMinute").Index(
								jen.Empty(),
								jen.Len(
									jen.Id("nMinute"),
								).Op("-").Lit(1),
							),
							jen.List(
								jen.Id("vMinute"),
								jen.Err(),
							).Op(":=").Qual("strconv", "ParseInt").Call(
								jen.Id("nMinute"),
								jen.Lit(10),
								jen.Lit(64),
							),
							jen.If(
								jen.Err().Op("!=").Nil(),
							).Block(
								jen.Return(
									jen.Lit(0),
									jen.Err(),
								),
							),
							jen.Id("dur").Op("+=").Qual("time", "Duration").Call(
								jen.Id("vMinute"),
							).Op("*").Qual("time", "Minute"),
						),
						// Seconds
						jen.Id("nSecond").Op(":=").Id("res").Index(jen.Lit(7)),
						jen.If(
							jen.Len(
								jen.Id("nSecond"),
							).Op(">").Lit(0),
						).Block(
							jen.Id("nSecond").Op("=").Id("nSecond").Index(
								jen.Empty(),
								jen.Len(
									jen.Id("nSecond"),
								).Op("-").Lit(1),
							),
							jen.List(
								jen.Id("vSecond"),
								jen.Err(),
							).Op(":=").Qual("strconv", "ParseInt").Call(
								jen.Id("nSecond"),
								jen.Lit(10),
								jen.Lit(64),
							),
							jen.If(
								jen.Err().Op("!=").Nil(),
							).Block(
								jen.Return(
									jen.Lit(0),
									jen.Err(),
								),
							),
							jen.Id("dur").Op("+=").Qual("time", "Duration").Call(
								jen.Id("vSecond"),
							).Op("*").Qual("time", "Second"),
						),
						jen.If(
							jen.Id("isNeg"),
						).Block(
							jen.Id("dur").Op("*=").Lit(-1),
						),
						jen.Return(
							jen.Id("dur"),
							jen.Nil(),
						),
					).Else().Block(
						jen.Return(
							jen.Lit(0),
							jen.Qual("fmt", "Errorf").Call(
								jen.Lit("%v cannot be interpreted as a string for xsd:duration"),
								jen.Id(codegen.This()),
							),
						),
					),
				}),
			LessFn: rdf.LessFunction(
				d.pkg,
				durationSpec,
				jen.Qual("time", "Duration"),
				[]jen.Code{
					jen.Return(
						jen.Id("lhs").Op("<").Id("rhs"),
					),
				}),
		}
		if err = v.SetValue(durationSpec, val); err != nil {
			return true, err
		}
	}
	return true, nil
}
