package propertycontent

import (
	"fmt"
	langstring "github.com/go-fed/activity/streams/values/langString"
	string1 "github.com/go-fed/activity/streams/values/string"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// ActivityStreamsContentPropertyIterator is an iterator for a property. It is
// permitted to be one of multiple value types. At most, one type of value can
// be present, or none at all. Setting a value will clear the other types of
// values so that only one of the 'Is' methods will return true. It is
// possible to clear all values, so that this property is empty.
type ActivityStreamsContentPropertyIterator struct {
	xmlschemaStringMember string
	hasStringMember       bool
	rdfLangStringMember   map[string]string
	unknown               interface{}
	iri                   *url.URL
	alias                 string
	myIdx                 int
	parent                vocab.ActivityStreamsContentProperty
}

// NewActivityStreamsContentPropertyIterator creates a new ActivityStreamsContent
// property.
func NewActivityStreamsContentPropertyIterator() *ActivityStreamsContentPropertyIterator {
	return &ActivityStreamsContentPropertyIterator{alias: ""}
}

// deserializeActivityStreamsContentPropertyIterator creates an iterator from an
// element that has been unmarshalled from a text or binary format.
func deserializeActivityStreamsContentPropertyIterator(i interface{}, aliasMap map[string]string) (*ActivityStreamsContentPropertyIterator, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	if s, ok := i.(string); ok {
		u, err := url.Parse(s)
		// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
		// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
		if err == nil && len(u.Scheme) > 0 {
			this := &ActivityStreamsContentPropertyIterator{
				alias: alias,
				iri:   u,
			}
			return this, nil
		}
	}
	if v, err := string1.DeserializeString(i); err == nil {
		this := &ActivityStreamsContentPropertyIterator{
			alias:                 alias,
			hasStringMember:       true,
			xmlschemaStringMember: v,
		}
		return this, nil
	} else if v, err := langstring.DeserializeLangString(i); err == nil {
		this := &ActivityStreamsContentPropertyIterator{
			alias:               alias,
			rdfLangStringMember: v,
		}
		return this, nil
	}
	this := &ActivityStreamsContentPropertyIterator{
		alias:   alias,
		unknown: i,
	}
	return this, nil
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this ActivityStreamsContentPropertyIterator) GetIRI() *url.URL {
	return this.iri
}

// GetLanguage returns the value for the specified BCP47 language code, or an
// empty string if it is either not a language map or no value is present.
func (this ActivityStreamsContentPropertyIterator) GetLanguage(bcp47 string) string {
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
func (this ActivityStreamsContentPropertyIterator) GetRDFLangString() map[string]string {
	return this.rdfLangStringMember
}

// GetXMLSchemaString returns the value of this property. When IsXMLSchemaString
// returns false, GetXMLSchemaString will return an arbitrary value.
func (this ActivityStreamsContentPropertyIterator) GetXMLSchemaString() string {
	return this.xmlschemaStringMember
}

// HasAny returns true if any of the values are set, except for the natural
// language map. When true, the specific has, getter, and setter methods may
// be used to determine what kind of value there is to access and set this
// property. To determine if the property was set as a natural language map,
// use the IsRDFLangString method instead.
func (this ActivityStreamsContentPropertyIterator) HasAny() bool {
	return this.IsXMLSchemaString() ||
		this.IsRDFLangString() ||
		this.iri != nil
}

// HasLanguage returns true if the natural language map has an entry for the
// specified BCP47 language code.
func (this ActivityStreamsContentPropertyIterator) HasLanguage(bcp47 string) bool {
	if this.rdfLangStringMember == nil {
		return false
	} else {
		_, ok := this.rdfLangStringMember[bcp47]
		return ok
	}
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this ActivityStreamsContentPropertyIterator) IsIRI() bool {
	return this.iri != nil
}

// IsRDFLangString returns true if this property has a type of "langString". When
// true, use the GetRDFLangString and SetRDFLangString methods to access and
// set this property.. To determine if the property was set as a natural
// language map, use the IsRDFLangString method instead.
func (this ActivityStreamsContentPropertyIterator) IsRDFLangString() bool {
	return this.rdfLangStringMember != nil
}

// IsXMLSchemaString returns true if this property has a type of "string". When
// true, use the GetXMLSchemaString and SetXMLSchemaString methods to access
// and set this property.. To determine if the property was set as a natural
// language map, use the IsRDFLangString method instead.
func (this ActivityStreamsContentPropertyIterator) IsXMLSchemaString() bool {
	return this.hasStringMember
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsContentPropertyIterator) JSONLDContext() map[string]string {
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
func (this ActivityStreamsContentPropertyIterator) KindIndex() int {
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
func (this ActivityStreamsContentPropertyIterator) LessThan(o vocab.ActivityStreamsContentPropertyIterator) bool {
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

// Name returns the name of this property: "ActivityStreamsContent".
func (this ActivityStreamsContentPropertyIterator) Name() string {
	if this.IsRDFLangString() {
		return "ActivityStreamsContentMap"
	} else {
		return "ActivityStreamsContent"
	}
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this ActivityStreamsContentPropertyIterator) Next() vocab.ActivityStreamsContentPropertyIterator {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this ActivityStreamsContentPropertyIterator) Prev() vocab.ActivityStreamsContentPropertyIterator {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *ActivityStreamsContentPropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.iri = v
}

// SetLanguage sets the value for the specified BCP47 language code.
func (this *ActivityStreamsContentPropertyIterator) SetLanguage(bcp47, value string) {
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
func (this *ActivityStreamsContentPropertyIterator) SetRDFLangString(v map[string]string) {
	this.clear()
	this.rdfLangStringMember = v
}

// SetXMLSchemaString sets the value of this property and clears the natural
// language map. Calling IsXMLSchemaString afterwards will return true.
// Calling IsRDFLangString afterwards returns false.
func (this *ActivityStreamsContentPropertyIterator) SetXMLSchemaString(v string) {
	this.clear()
	this.xmlschemaStringMember = v
	this.hasStringMember = true
}

// clear ensures no value and no language map for this property is set. Calling
// HasAny or any of the 'Is' methods afterwards will return false.
func (this *ActivityStreamsContentPropertyIterator) clear() {
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
func (this ActivityStreamsContentPropertyIterator) serialize() (interface{}, error) {
	if this.IsXMLSchemaString() {
		return string1.SerializeString(this.GetXMLSchemaString())
	} else if this.IsRDFLangString() {
		return langstring.SerializeLangString(this.GetRDFLangString())
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// ActivityStreamsContentProperty is the non-functional property "content". It is
// permitted to have one or more values, and of different value types.
type ActivityStreamsContentProperty struct {
	properties []*ActivityStreamsContentPropertyIterator
	alias      string
}

// DeserializeContentProperty creates a "content" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeContentProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsContentProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "content"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "content")
	}
	i, ok := m[propName]
	if !ok {
		// Attempt to find the map instead.
		i, ok = m[propName+"Map"]
	}
	if ok {
		this := &ActivityStreamsContentProperty{
			alias:      alias,
			properties: []*ActivityStreamsContentPropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeActivityStreamsContentPropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeActivityStreamsContentPropertyIterator(i, aliasMap); err != nil {
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

// NewActivityStreamsContentProperty creates a new content property.
func NewActivityStreamsContentProperty() *ActivityStreamsContentProperty {
	return &ActivityStreamsContentProperty{alias: ""}
}

// AppendIRI appends an IRI value to the back of a list of the property "content"
func (this *ActivityStreamsContentProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &ActivityStreamsContentPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  this.Len(),
		parent: this,
	})
}

// AppendRDFLangString appends a langString value to the back of a list of the
// property "content". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsContentProperty) AppendRDFLangString(v map[string]string) {
	this.properties = append(this.properties, &ActivityStreamsContentPropertyIterator{
		alias:               this.alias,
		myIdx:               this.Len(),
		parent:              this,
		rdfLangStringMember: v,
	})
}

// AppendXMLSchemaString appends a string value to the back of a list of the
// property "content". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsContentProperty) AppendXMLSchemaString(v string) {
	this.properties = append(this.properties, &ActivityStreamsContentPropertyIterator{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 this.Len(),
		parent:                this,
		xmlschemaStringMember: v,
	})
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this ActivityStreamsContentProperty) At(index int) vocab.ActivityStreamsContentPropertyIterator {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this ActivityStreamsContentProperty) Begin() vocab.ActivityStreamsContentPropertyIterator {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this ActivityStreamsContentProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this ActivityStreamsContentProperty) End() vocab.ActivityStreamsContentPropertyIterator {
	return nil
}

// Insert inserts an IRI value at the specified index for a property "content".
// Existing elements at that index and higher are shifted back once.
// Invalidates all iterators.
func (this *ActivityStreamsContentProperty) InsertIRI(idx int, v *url.URL) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ActivityStreamsContentPropertyIterator{
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
// property "content". Existing elements at that index and higher are shifted
// back once. Invalidates all iterators.
func (this *ActivityStreamsContentProperty) InsertRDFLangString(idx int, v map[string]string) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ActivityStreamsContentPropertyIterator{
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
// property "content". Existing elements at that index and higher are shifted
// back once. Invalidates all iterators.
func (this *ActivityStreamsContentProperty) InsertXMLSchemaString(idx int, v string) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ActivityStreamsContentPropertyIterator{
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
func (this ActivityStreamsContentProperty) JSONLDContext() map[string]string {
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
func (this ActivityStreamsContentProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "content" property.
func (this ActivityStreamsContentProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this ActivityStreamsContentProperty) Less(i, j int) bool {
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
func (this ActivityStreamsContentProperty) LessThan(o vocab.ActivityStreamsContentProperty) bool {
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

// Name returns the name of this property: "content".
func (this ActivityStreamsContentProperty) Name() string {
	if this.Len() == 1 && this.At(0).IsRDFLangString() {
		return "contentMap"
	} else {
		return "content"
	}
}

// PrependIRI prepends an IRI value to the front of a list of the property
// "content".
func (this *ActivityStreamsContentProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*ActivityStreamsContentPropertyIterator{{
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
// property "content". Invalidates all iterators.
func (this *ActivityStreamsContentProperty) PrependRDFLangString(v map[string]string) {
	this.properties = append([]*ActivityStreamsContentPropertyIterator{{
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
// property "content". Invalidates all iterators.
func (this *ActivityStreamsContentProperty) PrependXMLSchemaString(v string) {
	this.properties = append([]*ActivityStreamsContentPropertyIterator{{
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
// "content", regardless of its type. Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *ActivityStreamsContentProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &ActivityStreamsContentPropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsContentProperty) Serialize() (interface{}, error) {
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

// SetIRI sets an IRI value to be at the specified index for the property
// "content". Panics if the index is out of bounds.
func (this *ActivityStreamsContentProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsContentPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
}

// SetRDFLangString sets a langString value to be at the specified index for the
// property "content". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsContentProperty) SetRDFLangString(idx int, v map[string]string) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsContentPropertyIterator{
		alias:               this.alias,
		myIdx:               idx,
		parent:              this,
		rdfLangStringMember: v,
	}
}

// SetXMLSchemaString sets a string value to be at the specified index for the
// property "content". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsContentProperty) SetXMLSchemaString(idx int, v string) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsContentPropertyIterator{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 idx,
		parent:                this,
		xmlschemaStringMember: v,
	}
}

// Swap swaps the location of values at two indices for the "content" property.
func (this ActivityStreamsContentProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}
