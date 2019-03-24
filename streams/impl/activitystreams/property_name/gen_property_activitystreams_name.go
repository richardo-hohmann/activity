package propertyname

import (
	"fmt"
	langstring "github.com/go-fed/activity/streams/values/langString"
	string1 "github.com/go-fed/activity/streams/values/string"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// ActivityStreamsNamePropertyIterator is an iterator for a property. It is
// permitted to be one of multiple value types. At most, one type of value can
// be present, or none at all. Setting a value will clear the other types of
// values so that only one of the 'Is' methods will return true. It is
// possible to clear all values, so that this property is empty.
type ActivityStreamsNamePropertyIterator struct {
	xmlschemaStringMember string
	hasStringMember       bool
	rdfLangStringMember   map[string]string
	unknown               interface{}
	iri                   *url.URL
	alias                 string
	myIdx                 int
	parent                vocab.ActivityStreamsNameProperty
}

// NewActivityStreamsNamePropertyIterator creates a new ActivityStreamsName
// property.
func NewActivityStreamsNamePropertyIterator() *ActivityStreamsNamePropertyIterator {
	return &ActivityStreamsNamePropertyIterator{alias: ""}
}

// deserializeActivityStreamsNamePropertyIterator creates an iterator from an
// element that has been unmarshalled from a text or binary format.
func deserializeActivityStreamsNamePropertyIterator(i interface{}, aliasMap map[string]string) (*ActivityStreamsNamePropertyIterator, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	if s, ok := i.(string); ok {
		u, err := url.Parse(s)
		// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
		// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
		if err == nil && len(u.Scheme) > 0 {
			this := &ActivityStreamsNamePropertyIterator{
				alias: alias,
				iri:   u,
			}
			return this, nil
		}
	}
	if v, err := string1.DeserializeString(i); err == nil {
		this := &ActivityStreamsNamePropertyIterator{
			alias:                 alias,
			hasStringMember:       true,
			xmlschemaStringMember: v,
		}
		return this, nil
	} else if v, err := langstring.DeserializeLangString(i); err == nil {
		this := &ActivityStreamsNamePropertyIterator{
			alias:               alias,
			rdfLangStringMember: v,
		}
		return this, nil
	}
	this := &ActivityStreamsNamePropertyIterator{
		alias:   alias,
		unknown: i,
	}
	return this, nil
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this ActivityStreamsNamePropertyIterator) GetIRI() *url.URL {
	return this.iri
}

// GetLanguage returns the value for the specified BCP47 language code, or an
// empty string if it is either not a language map or no value is present.
func (this ActivityStreamsNamePropertyIterator) GetLanguage(bcp47 string) string {
	if this.rdfLangStringMember == nil {
		return ""
	} else if v, ok := this.rdfLangStringMember[bcp47]; ok {
		return v
	} else {
		return ""
	}
}

// GetRDFLangString returns the value of this property. When IsRDFLangString
// returns false, GetRDFLangString will return an arbitrary value.
func (this ActivityStreamsNamePropertyIterator) GetRDFLangString() map[string]string {
	return this.rdfLangStringMember
}

// GetXMLSchemaString returns the value of this property. When IsXMLSchemaString
// returns false, GetXMLSchemaString will return an arbitrary value.
func (this ActivityStreamsNamePropertyIterator) GetXMLSchemaString() string {
	return this.xmlschemaStringMember
}

// HasAny returns true if any of the values are set, except for the natural
// language map. When true, the specific has, getter, and setter methods may
// be used to determine what kind of value there is to access and set this
// property. To determine if the property was set as a natural language map,
// use the IsRDFLangString method instead.
func (this ActivityStreamsNamePropertyIterator) HasAny() bool {
	return this.IsXMLSchemaString() ||
		this.IsRDFLangString() ||
		this.iri != nil
}

// HasLanguage returns true if the natural language map has an entry for the
// specified BCP47 language code.
func (this ActivityStreamsNamePropertyIterator) HasLanguage(bcp47 string) bool {
	if this.rdfLangStringMember == nil {
		return false
	} else {
		_, ok := this.rdfLangStringMember[bcp47]
		return ok
	}
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this ActivityStreamsNamePropertyIterator) IsIRI() bool {
	return this.iri != nil
}

// IsRDFLangString returns true if this property has a type of "langString". When
// true, use the GetRDFLangString and SetRDFLangString methods to access and
// set this property.. To determine if the property was set as a natural
// language map, use the IsRDFLangString method instead.
func (this ActivityStreamsNamePropertyIterator) IsRDFLangString() bool {
	return this.rdfLangStringMember != nil
}

// IsXMLSchemaString returns true if this property has a type of "string". When
// true, use the GetXMLSchemaString and SetXMLSchemaString methods to access
// and set this property.. To determine if the property was set as a natural
// language map, use the IsRDFLangString method instead.
func (this ActivityStreamsNamePropertyIterator) IsXMLSchemaString() bool {
	return this.hasStringMember
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsNamePropertyIterator) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
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
func (this ActivityStreamsNamePropertyIterator) KindIndex() int {
	if this.IsXMLSchemaString() {
		return 0
	}
	if this.IsRDFLangString() {
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
func (this ActivityStreamsNamePropertyIterator) LessThan(o vocab.ActivityStreamsNamePropertyIterator) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsXMLSchemaString() {
		return string1.LessString(this.GetXMLSchemaString(), o.GetXMLSchemaString())
	} else if this.IsRDFLangString() {
		return langstring.LessLangString(this.GetRDFLangString(), o.GetRDFLangString())
	} else if this.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	}
	return false
}

// Name returns the name of this property: "ActivityStreamsName".
func (this ActivityStreamsNamePropertyIterator) Name() string {
	if this.IsRDFLangString() {
		return "ActivityStreamsNameMap"
	} else {
		return "ActivityStreamsName"
	}
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this ActivityStreamsNamePropertyIterator) Next() vocab.ActivityStreamsNamePropertyIterator {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this ActivityStreamsNamePropertyIterator) Prev() vocab.ActivityStreamsNamePropertyIterator {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *ActivityStreamsNamePropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.iri = v
}

// SetLanguage sets the value for the specified BCP47 language code.
func (this *ActivityStreamsNamePropertyIterator) SetLanguage(bcp47, value string) {
	this.hasStringMember = false
	this.rdfLangStringMember = nil
	this.unknown = nil
	this.iri = nil
	if this.rdfLangStringMember == nil {
		this.rdfLangStringMember = make(map[string]string)
	}
	this.rdfLangStringMember[bcp47] = value
}

// SetRDFLangString sets the value of this property and clears the natural
// language map. Calling IsRDFLangString afterwards will return true. Calling
// IsRDFLangString afterwards returns false.
func (this *ActivityStreamsNamePropertyIterator) SetRDFLangString(v map[string]string) {
	this.clear()
	this.rdfLangStringMember = v
}

// SetXMLSchemaString sets the value of this property and clears the natural
// language map. Calling IsXMLSchemaString afterwards will return true.
// Calling IsRDFLangString afterwards returns false.
func (this *ActivityStreamsNamePropertyIterator) SetXMLSchemaString(v string) {
	this.clear()
	this.xmlschemaStringMember = v
	this.hasStringMember = true
}

// clear ensures no value and no language map for this property is set. Calling
// HasAny or any of the 'Is' methods afterwards will return false.
func (this *ActivityStreamsNamePropertyIterator) clear() {
	this.hasStringMember = false
	this.rdfLangStringMember = nil
	this.unknown = nil
	this.iri = nil
	this.rdfLangStringMember = nil
}

// serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsNamePropertyIterator) serialize() (interface{}, error) {
	if this.IsXMLSchemaString() {
		return string1.SerializeString(this.GetXMLSchemaString())
	} else if this.IsRDFLangString() {
		return langstring.SerializeLangString(this.GetRDFLangString())
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// ActivityStreamsNameProperty is the non-functional property "name". It is
// permitted to have one or more values, and of different value types.
type ActivityStreamsNameProperty struct {
	properties []*ActivityStreamsNamePropertyIterator
	alias      string
}

// DeserializeNameProperty creates a "name" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeNameProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsNameProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "name"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "name")
	}
	i, ok := m[propName]
	if !ok {
		// Attempt to find the map instead.
		i, ok = m[propName+"Map"]
	}
	if ok {
		this := &ActivityStreamsNameProperty{
			alias:      alias,
			properties: []*ActivityStreamsNamePropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeActivityStreamsNamePropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeActivityStreamsNamePropertyIterator(i, aliasMap); err != nil {
				return this, err
			} else if p != nil {
				this.properties = append(this.properties, p)
			}
		}
		// Set up the properties for iteration.
		for idx, ele := range this.properties {
			ele.parent = this
			ele.myIdx = idx
		}
		return this, nil
	}
	return nil, nil
}

// NewActivityStreamsNameProperty creates a new name property.
func NewActivityStreamsNameProperty() *ActivityStreamsNameProperty {
	return &ActivityStreamsNameProperty{alias: ""}
}

// AppendIRI appends an IRI value to the back of a list of the property "name"
func (this *ActivityStreamsNameProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &ActivityStreamsNamePropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  this.Len(),
		parent: this,
	})
}

// AppendRDFLangString appends a langString value to the back of a list of the
// property "name". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsNameProperty) AppendRDFLangString(v map[string]string) {
	this.properties = append(this.properties, &ActivityStreamsNamePropertyIterator{
		alias:               this.alias,
		myIdx:               this.Len(),
		parent:              this,
		rdfLangStringMember: v,
	})
}

// AppendXMLSchemaString appends a string value to the back of a list of the
// property "name". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsNameProperty) AppendXMLSchemaString(v string) {
	this.properties = append(this.properties, &ActivityStreamsNamePropertyIterator{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 this.Len(),
		parent:                this,
		xmlschemaStringMember: v,
	})
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this ActivityStreamsNameProperty) At(index int) vocab.ActivityStreamsNamePropertyIterator {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this ActivityStreamsNameProperty) Begin() vocab.ActivityStreamsNamePropertyIterator {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this ActivityStreamsNameProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this ActivityStreamsNameProperty) End() vocab.ActivityStreamsNamePropertyIterator {
	return nil
}

// Insert inserts an IRI value at the specified index for a property "name".
// Existing elements at that index and higher are shifted back once.
// Invalidates all iterators.
func (this *ActivityStreamsNameProperty) InsertIRI(idx int, v *url.URL) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ActivityStreamsNamePropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// InsertRDFLangString inserts a langString value at the specified index for a
// property "name". Existing elements at that index and higher are shifted
// back once. Invalidates all iterators.
func (this *ActivityStreamsNameProperty) InsertRDFLangString(idx int, v map[string]string) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ActivityStreamsNamePropertyIterator{
		alias:               this.alias,
		myIdx:               idx,
		parent:              this,
		rdfLangStringMember: v,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// InsertXMLSchemaString inserts a string value at the specified index for a
// property "name". Existing elements at that index and higher are shifted
// back once. Invalidates all iterators.
func (this *ActivityStreamsNameProperty) InsertXMLSchemaString(idx int, v string) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ActivityStreamsNamePropertyIterator{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 idx,
		parent:                this,
		xmlschemaStringMember: v,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsNameProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	for _, elem := range this.properties {
		child := elem.JSONLDContext()
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		for k, v := range child {
			m[k] = v
		}
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API method specifically needed only for alternate implementations
// for go-fed. Applications should not use this method. Panics if the index is
// out of bounds.
func (this ActivityStreamsNameProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "name" property.
func (this ActivityStreamsNameProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this ActivityStreamsNameProperty) Less(i, j int) bool {
	idx1 := this.KindIndex(i)
	idx2 := this.KindIndex(j)
	if idx1 < idx2 {
		return true
	} else if idx1 == idx2 {
		if idx1 == 0 {
			lhs := this.properties[i].GetXMLSchemaString()
			rhs := this.properties[j].GetXMLSchemaString()
			return string1.LessString(lhs, rhs)
		} else if idx1 == 1 {
			lhs := this.properties[i].GetRDFLangString()
			rhs := this.properties[j].GetRDFLangString()
			return langstring.LessLangString(lhs, rhs)
		} else if idx1 == -2 {
			lhs := this.properties[i].GetIRI()
			rhs := this.properties[j].GetIRI()
			return lhs.String() < rhs.String()
		}
	}
	return false
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this ActivityStreamsNameProperty) LessThan(o vocab.ActivityStreamsNameProperty) bool {
	l1 := this.Len()
	l2 := o.Len()
	l := l1
	if l2 < l1 {
		l = l2
	}
	for i := 0; i < l; i++ {
		if this.properties[i].LessThan(o.At(i)) {
			return true
		} else if o.At(i).LessThan(this.properties[i]) {
			return false
		}
	}
	return l1 < l2
}

// Name returns the name of this property: "name".
func (this ActivityStreamsNameProperty) Name() string {
	if this.Len() == 1 && this.At(0).IsRDFLangString() {
		return "nameMap"
	} else {
		return "name"
	}
}

// PrependIRI prepends an IRI value to the front of a list of the property "name".
func (this *ActivityStreamsNameProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*ActivityStreamsNamePropertyIterator{{
		alias:  this.alias,
		iri:    v,
		myIdx:  0,
		parent: this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependRDFLangString prepends a langString value to the front of a list of the
// property "name". Invalidates all iterators.
func (this *ActivityStreamsNameProperty) PrependRDFLangString(v map[string]string) {
	this.properties = append([]*ActivityStreamsNamePropertyIterator{{
		alias:               this.alias,
		myIdx:               0,
		parent:              this,
		rdfLangStringMember: v,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependXMLSchemaString prepends a string value to the front of a list of the
// property "name". Invalidates all iterators.
func (this *ActivityStreamsNameProperty) PrependXMLSchemaString(v string) {
	this.properties = append([]*ActivityStreamsNamePropertyIterator{{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 0,
		parent:                this,
		xmlschemaStringMember: v,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Remove deletes an element at the specified index from a list of the property
// "name", regardless of its type. Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *ActivityStreamsNameProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &ActivityStreamsNamePropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsNameProperty) Serialize() (interface{}, error) {
	s := make([]interface{}, 0, len(this.properties))
	for _, iterator := range this.properties {
		if b, err := iterator.serialize(); err != nil {
			return s, err
		} else {
			s = append(s, b)
		}
	}
	// Shortcut: if serializing one value, don't return an array -- pretty sure other Fediverse software would choke on a "type" value with array, for example.
	if len(s) == 1 {
		return s[0], nil
	}
	return s, nil
}

// SetIRI sets an IRI value to be at the specified index for the property "name".
// Panics if the index is out of bounds.
func (this *ActivityStreamsNameProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsNamePropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
}

// SetRDFLangString sets a langString value to be at the specified index for the
// property "name". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsNameProperty) SetRDFLangString(idx int, v map[string]string) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsNamePropertyIterator{
		alias:               this.alias,
		myIdx:               idx,
		parent:              this,
		rdfLangStringMember: v,
	}
}

// SetXMLSchemaString sets a string value to be at the specified index for the
// property "name". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsNameProperty) SetXMLSchemaString(idx int, v string) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsNamePropertyIterator{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 idx,
		parent:                this,
		xmlschemaStringMember: v,
	}
}

// Swap swaps the location of values at two indices for the "name" property.
func (this ActivityStreamsNameProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}
