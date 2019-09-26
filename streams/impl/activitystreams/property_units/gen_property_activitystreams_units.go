package propertyunits

import (
	"fmt"
	anyuri "github.com/go-fed/activity/streams/values/anyURI"
	string1 "github.com/go-fed/activity/streams/values/string"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// ActivityStreamsUnitsProperty is the functional property "units". It is
// permitted to be one of multiple value types. At most, one type of value can
// be present, or none at all. Setting a value will clear the other types of
// values so that only one of the 'Is' methods will return true. It is
// possible to clear all values, so that this property is empty.
type ActivityStreamsUnitsProperty struct {
	xmlschemaStringMember string
	hasStringMember       bool
	xmlschemaAnyURIMember *url.URL
	unknown               interface{}
	alias                 string
}

// DeserializeUnitsProperty creates a "units" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeUnitsProperty(m map[string]interface{}, aliasMap map[string]string) (*ActivityStreamsUnitsProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/ns/activitystreams"]; ok {
		alias = a
	}
	propName := "units"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "units")
	}
	i, ok := m[propName]

	if ok {
		if v, err := string1.DeserializeString(i); err == nil {
			this := &ActivityStreamsUnitsProperty{
				alias:                 alias,
				hasStringMember:       true,
				xmlschemaStringMember: v,
			}
			return this, nil
		} else if v, err := anyuri.DeserializeAnyURI(i); err == nil {
			this := &ActivityStreamsUnitsProperty{
				alias:                 alias,
				xmlschemaAnyURIMember: v,
			}
			return this, nil
		}
		this := &ActivityStreamsUnitsProperty{
			alias:   alias,
			unknown: i,
		}
		return this, nil
	}
	return nil, nil
}

// NewActivityStreamsUnitsProperty creates a new units property.
func NewActivityStreamsUnitsProperty() *ActivityStreamsUnitsProperty {
	return &ActivityStreamsUnitsProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling HasAny or any of the
// 'Is' methods afterwards will return false.
func (this *ActivityStreamsUnitsProperty) Clear() {
	this.hasStringMember = false
	this.xmlschemaAnyURIMember = nil
	this.unknown = nil
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this ActivityStreamsUnitsProperty) GetIRI() *url.URL {
	return this.xmlschemaAnyURIMember
}

// GetXMLSchemaAnyURI returns the value of this property. When IsXMLSchemaAnyURI
// returns false, GetXMLSchemaAnyURI will return an arbitrary value.
func (this ActivityStreamsUnitsProperty) GetXMLSchemaAnyURI() *url.URL {
	return this.xmlschemaAnyURIMember
}

// GetXMLSchemaString returns the value of this property. When IsXMLSchemaString
// returns false, GetXMLSchemaString will return an arbitrary value.
func (this ActivityStreamsUnitsProperty) GetXMLSchemaString() string {
	return this.xmlschemaStringMember
}

// HasAny returns true if any of the different values is set.
func (this ActivityStreamsUnitsProperty) HasAny() bool {
	return this.IsXMLSchemaString() ||
		this.IsXMLSchemaAnyURI()
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this ActivityStreamsUnitsProperty) IsIRI() bool {
	return this.xmlschemaAnyURIMember != nil
}

// IsXMLSchemaAnyURI returns true if this property has a type of "anyURI". When
// true, use the GetXMLSchemaAnyURI and SetXMLSchemaAnyURI methods to access
// and set this property.
func (this ActivityStreamsUnitsProperty) IsXMLSchemaAnyURI() bool {
	return this.xmlschemaAnyURIMember != nil
}

// IsXMLSchemaString returns true if this property has a type of "string". When
// true, use the GetXMLSchemaString and SetXMLSchemaString methods to access
// and set this property.
func (this ActivityStreamsUnitsProperty) IsXMLSchemaString() bool {
	return this.hasStringMember
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsUnitsProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/ns/activitystreams": this.alias}
	var child map[string]string

	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this ActivityStreamsUnitsProperty) KindIndex() int {
	if this.IsXMLSchemaString() {
		return 0
	}
	if this.IsXMLSchemaAnyURI() {
		return 1
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this ActivityStreamsUnitsProperty) LessThan(o vocab.ActivityStreamsUnitsProperty) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsXMLSchemaString() {
		return string1.LessString(this.GetXMLSchemaString(), o.GetXMLSchemaString())
	} else if this.IsXMLSchemaAnyURI() {
		return anyuri.LessAnyURI(this.GetXMLSchemaAnyURI(), o.GetXMLSchemaAnyURI())
	}
	return false
}

// Name returns the name of this property: "units".
func (this ActivityStreamsUnitsProperty) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "units"
	} else {
		return "units"
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsUnitsProperty) Serialize() (interface{}, error) {
	if this.IsXMLSchemaString() {
		return string1.SerializeString(this.GetXMLSchemaString())
	} else if this.IsXMLSchemaAnyURI() {
		return anyuri.SerializeAnyURI(this.GetXMLSchemaAnyURI())
	}
	return this.unknown, nil
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *ActivityStreamsUnitsProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.SetXMLSchemaAnyURI(v)
}

// SetXMLSchemaAnyURI sets the value of this property. Calling IsXMLSchemaAnyURI
// afterwards returns true.
func (this *ActivityStreamsUnitsProperty) SetXMLSchemaAnyURI(v *url.URL) {
	this.Clear()
	this.xmlschemaAnyURIMember = v
}

// SetXMLSchemaString sets the value of this property. Calling IsXMLSchemaString
// afterwards returns true.
func (this *ActivityStreamsUnitsProperty) SetXMLSchemaString(v string) {
	this.Clear()
	this.xmlschemaStringMember = v
	this.hasStringMember = true
}
