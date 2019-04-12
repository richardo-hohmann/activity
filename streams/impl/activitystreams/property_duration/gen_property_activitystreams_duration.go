package propertyduration

import (
	"fmt"
	duration "github.com/go-fed/activity/streams/values/duration"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
	"time"
)

// ActivityStreamsDurationProperty is the functional property "duration". It is
// permitted to be a single default-valued value type.
type ActivityStreamsDurationProperty struct {
	xmlschemaDurationMember time.Duration
	hasDurationMember       bool
	unknown                 interface{}
	iri                     *url.URL
	alias                   string
}

// DeserializeDurationProperty creates a "duration" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeDurationProperty(m map[string]interface{}, aliasMap map[string]string) (*ActivityStreamsDurationProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/ns/activitystreams"]; ok {
		alias = a
	}
	propName := "duration"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "duration")
	}
	i, ok := m[propName]

	if ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
			if err == nil && len(u.Scheme) > 0 {
				this := &ActivityStreamsDurationProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if v, err := duration.DeserializeDuration(i); err == nil {
			this := &ActivityStreamsDurationProperty{
				alias:                   alias,
				hasDurationMember:       true,
				xmlschemaDurationMember: v,
			}
			return this, nil
		}
		this := &ActivityStreamsDurationProperty{
			alias:   alias,
			unknown: i,
		}
		return this, nil
	}
	return nil, nil
}

// NewActivityStreamsDurationProperty creates a new duration property.
func NewActivityStreamsDurationProperty() *ActivityStreamsDurationProperty {
	return &ActivityStreamsDurationProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling IsXMLSchemaDuration
// afterwards will return false.
func (this *ActivityStreamsDurationProperty) Clear() {
	this.unknown = nil
	this.iri = nil
	this.hasDurationMember = false
}

// Get returns the value of this property. When IsXMLSchemaDuration returns false,
// Get will return any arbitrary value.
func (this ActivityStreamsDurationProperty) Get() time.Duration {
	return this.xmlschemaDurationMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this ActivityStreamsDurationProperty) GetIRI() *url.URL {
	return this.iri
}

// HasAny returns true if the value or IRI is set.
func (this ActivityStreamsDurationProperty) HasAny() bool {
	return this.IsXMLSchemaDuration() || this.iri != nil
}

// IsIRI returns true if this property is an IRI.
func (this ActivityStreamsDurationProperty) IsIRI() bool {
	return this.iri != nil
}

// IsXMLSchemaDuration returns true if this property is set and not an IRI.
func (this ActivityStreamsDurationProperty) IsXMLSchemaDuration() bool {
	return this.hasDurationMember
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsDurationProperty) JSONLDContext() map[string]string {
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
func (this ActivityStreamsDurationProperty) KindIndex() int {
	if this.IsXMLSchemaDuration() {
		return 0
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
func (this ActivityStreamsDurationProperty) LessThan(o vocab.ActivityStreamsDurationProperty) bool {
	// LessThan comparison for if either or both are IRIs.
	if this.IsIRI() && o.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	} else if this.IsIRI() {
		// IRIs are always less than other values, none, or unknowns
		return true
	} else if o.IsIRI() {
		// This other, none, or unknown value is always greater than IRIs
		return false
	}
	// LessThan comparison for the single value or unknown value.
	if !this.IsXMLSchemaDuration() && !o.IsXMLSchemaDuration() {
		// Both are unknowns.
		return false
	} else if this.IsXMLSchemaDuration() && !o.IsXMLSchemaDuration() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsXMLSchemaDuration() && o.IsXMLSchemaDuration() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return duration.LessDuration(this.Get(), o.Get())
	}
}

// Name returns the name of this property: "duration".
func (this ActivityStreamsDurationProperty) Name() string {
	return "duration"
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsDurationProperty) Serialize() (interface{}, error) {
	if this.IsXMLSchemaDuration() {
		return duration.SerializeDuration(this.Get())
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// Set sets the value of this property. Calling IsXMLSchemaDuration afterwards
// will return true.
func (this *ActivityStreamsDurationProperty) Set(v time.Duration) {
	this.Clear()
	this.xmlschemaDurationMember = v
	this.hasDurationMember = true
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *ActivityStreamsDurationProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}
