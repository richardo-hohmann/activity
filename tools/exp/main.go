package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/cjslep/activity/tools/exp/convert"
	"github.com/cjslep/activity/tools/exp/gen"
	"github.com/cjslep/activity/tools/exp/rdf"
	"github.com/cjslep/activity/tools/exp/rdf/owl"
	"github.com/cjslep/activity/tools/exp/rdf/rdfs"
	"github.com/cjslep/activity/tools/exp/rdf/rfc"
	"github.com/cjslep/activity/tools/exp/rdf/schema"
	"github.com/cjslep/activity/tools/exp/rdf/xsd"
	"io/ioutil"
	"os"
	"strings"
)

var registry *rdf.RDFRegistry

func mustAddOntology(o rdf.Ontology) {
	if registry == nil {
		registry = rdf.NewRDFRegistry()
	}
	if err := registry.AddOntology(o); err != nil {
		panic(err)
	}
}

func init() {
	mustAddOntology(&xsd.XMLOntology{Package: "xml"})
	mustAddOntology(&owl.OWLOntology{})
	mustAddOntology(&rdf.RDFOntology{Package: "rdf"})
	mustAddOntology(&rdfs.RDFSchemaOntology{})
	mustAddOntology(&schema.SchemaOntology{})
	mustAddOntology(&rfc.RFCOntology{Package: "rfc"})
}

var (
	input = flag.String("input", "spec.json", "Input JSON-LD specification used to generate Go code.")
	// TODO: Be more rigorous when applying this. Also, clear the default value I am using for convenience.
	prefix     = flag.String("prefix", "github.com/cjslep/activity/tools/exp/tmp", "Package prefix to use for all generated package paths. This should be the prefix in the GOPATH directory if generating in a subdirectory.")
	individual = flag.Bool("individual", false, "Whether to generate types and properties in individual packages.")
)

type list []string

func (l *list) String() string {
	return strings.Join(*l, ",")
}

func (l *list) Set(v string) error {
	vals := strings.Split(v, ",")
	*l = append(*l, vals...)
	return nil
}

func main() {
	flag.Parse()
	// TODO: Flag validation

	b, err := ioutil.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	var inputJSON map[string]interface{}
	err = json.Unmarshal(b, &inputJSON)
	if err != nil {
		panic(err)
	}
	p, err := rdf.ParseVocabulary(registry, inputJSON)
	if err != nil {
		panic(err)
	}
	policy := convert.FlatUnderRoot
	if *individual {
		policy = convert.IndividualUnderRoot
	}
	fmt.Printf("Vocab Name: %q\n", p.Vocab.Name)
	c := &convert.Converter{
		Registry:       registry,
		GenRoot:        gen.NewPackageManager(*prefix, "gen"),
		VocabularyName: p.Vocab.Name,
		PackagePolicy:  policy,
	}
	f, err := c.Convert(p)
	if err != nil {
		panic(err)
	}
	for _, file := range f {
		if e := os.MkdirAll("./"+file.Directory, 0777); e != nil {
			panic(e)
		}
		if e := file.F.Save("./" + file.Directory + "/" + file.FileName); e != nil {
			panic(e)
		}
	}
	fmt.Printf("done\n")
}
