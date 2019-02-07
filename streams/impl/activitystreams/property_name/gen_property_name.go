package propertyname

import (
	"fmt"
	langstring "github.com/go-fed/activity/streams/values/langString"
	string1 "github.com/go-fed/activity/streams/values/string"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// NamePropertyIterator is an iterator for a property. It is permitted to be one
// of multiple value types. At most, one type of value can be present, or none
// at all. Setting a value will clear the other types of values so that only
// one of the 'Is' methods will return true. It is possible to clear all
// values, so that this property is empty.
type NamePropertyIterator struct {
	stringMember     string
	hasStringMember  bool
	langStringMember map[string]string
	unknown          interface{}
	iri              *url.URL
	alias            string
	langMap          map[string]string
	myIdx            int
	parent           vocab.NamePropertyInterface
}

// NewNamePropertyIterator creates a new name property.
func NewNamePropertyIterator() *NamePropertyIterator {
	return &NamePropertyIterator{alias: ""}
}

// deserializeNamePropertyIterator creates an iterator from an element that has
// been unmarshalled from a text or binary format.
func deserializeNamePropertyIterator(i interface{}, aliasMap map[string]string) (*NamePropertyIterator, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	if s, ok := i.(string); ok {
		u, err := url.Parse(s)
		// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
		// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
		if err == nil && len(u.Scheme) > 0 {
			this := &NamePropertyIterator{
				alias: alias,
				iri:   u,
			}
			return this, nil
		}
	}
	if v, err := string1.DeserializeString(i); err == nil {
		this := &NamePropertyIterator{
			alias:           alias,
			hasStringMember: true,
			stringMember:    v,
		}
		return this, nil
	} else if v, err := langstring.DeserializeLangString(i); err == nil {
		this := &NamePropertyIterator{
			alias:            alias,
			langStringMember: v,
		}
		return this, nil
	}
	this := &NamePropertyIterator{
		alias:   alias,
		unknown: i,
	}
	return this, nil
	return nil, fmt.Errorf("could not deserialize %q property", "name")
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this NamePropertyIterator) GetIRI() *url.URL {
	return this.iri
}

// GetLangString returns the value of this property. When IsLangString returns
// false, GetLangString will return an arbitrary value.
func (this NamePropertyIterator) GetLangString() map[string]string {
	return this.langStringMember
}

// GetLanguage returns the value for the specified BCP47 language code, or an
// empty string if it is either not a language map or no value is present.
func (this NamePropertyIterator) GetLanguage(bcp47 string) string {
	if this.langMap == nil {
		return ""
	} else if v, ok := this.langMap[bcp47]; ok {
		return v
	} else {
		return ""
	}
}

// GetString returns the value of this property. When IsString returns false,
// GetString will return an arbitrary value.
func (this NamePropertyIterator) GetString() string {
	return this.stringMember
}

// HasAny returns true if any of the values are set, except for the natural
// language map. When true, the specific has, getter, and setter methods may
// be used to determine what kind of value there is to access and set this
// property. To determine if the property was set as a natural language map,
// use the IsLanguageMap method instead.
func (this NamePropertyIterator) HasAny() bool {
	return this.IsString() ||
		this.IsLangString() ||
		this.iri != nil
}

// HasLanguage returns true if the natural language map has an entry for the
// specified BCP47 language code.
func (this NamePropertyIterator) HasLanguage(bcp47 string) bool {
	if this.langMap == nil {
		return false
	} else {
		_, ok := this.langMap[bcp47]
		return ok
	}
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this NamePropertyIterator) IsIRI() bool {
	return this.iri != nil
}

// IsLangString returns true if this property has a type of "langString". When
// true, use the GetLangString and SetLangString methods to access and set
// this property.. To determine if the property was set as a natural language
// map, use the IsLanguageMap method instead.
func (this NamePropertyIterator) IsLangString() bool {
	return this.langStringMember != nil
}

// IsLanguageMap determines if this property is represented by a natural language
// map. When true, use HasLanguage, GetLanguage, and SetLanguage methods to
// access and mutate the natural language map. The clear method can be used to
// clear the natural language map. Note that this method is only used for
// natural language representations, and does not determine the presence nor
// absence of other values for this property.
func (this NamePropertyIterator) IsLanguageMap() bool {
	return this.langMap != nil
}

// IsString returns true if this property has a type of "string". When true, use
// the GetString and SetString methods to access and set this property.. To
// determine if the property was set as a natural language map, use the
// IsLanguageMap method instead.
func (this NamePropertyIterator) IsString() bool {
	return this.hasStringMember
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this NamePropertyIterator) JSONLDContext() map[string]string {
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
func (this NamePropertyIterator) KindIndex() int {
	if this.IsString() {
		return 0
	}
	if this.IsLangString() {
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
func (this NamePropertyIterator) LessThan(o vocab.NamePropertyIteratorInterface) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsString() {
		return string1.LessString(this.GetString(), o.GetString())
	} else if this.IsLangString() {
		return langstring.LessLangString(this.GetLangString(), o.GetLangString())
	} else if this.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	}
	return false
}

// Name returns the name of this property: "name".
func (this NamePropertyIterator) Name() string {
	return "name"
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this NamePropertyIterator) Next() vocab.NamePropertyIteratorInterface {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this NamePropertyIterator) Prev() vocab.NamePropertyIteratorInterface {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *NamePropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.iri = v
}

// SetLangString sets the value of this property and clears the natural language
// map. Calling IsLangString afterwards will return true. Calling
// IsLanguageMap afterwards returns false.
func (this *NamePropertyIterator) SetLangString(v map[string]string) {
	this.clear()
	this.langStringMember = v
}

// SetLanguage sets the value for the specified BCP47 language code.
func (this *NamePropertyIterator) SetLanguage(bcp47, value string) {
	this.hasStringMember = false
	this.langStringMember = nil
	this.unknown = nil
	this.iri = nil
	if this.langMap == nil {
		this.langMap = make(map[string]string)
	}
	this.langMap[bcp47] = value
}

// SetString sets the value of this property and clears the natural language map.
// Calling IsString afterwards will return true. Calling IsLanguageMap
// afterwards returns false.
func (this *NamePropertyIterator) SetString(v string) {
	this.clear()
	this.stringMember = v
	this.hasStringMember = true
}

// clear ensures no value and no language map for this property is set. Calling
// HasAny or any of the 'Is' methods afterwards will return false.
func (this *NamePropertyIterator) clear() {
	this.hasStringMember = false
	this.langStringMember = nil
	this.unknown = nil
	this.iri = nil
	this.langMap = nil
}

// serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this NamePropertyIterator) serialize() (interface{}, error) {
	if this.IsString() {
		return string1.SerializeString(this.GetString())
	} else if this.IsLangString() {
		return langstring.SerializeLangString(this.GetLangString())
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// NameProperty is the non-functional property "name". It is permitted to have one
// or more values, and of different value types.
type NameProperty struct {
	properties []*NamePropertyIterator
	alias      string
}

// DeserializeNameProperty creates a "name" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeNameProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.NamePropertyInterface, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "name"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "name")
	}
	if i, ok := m[propName]; ok {
		this := &NameProperty{
			alias:      alias,
			properties: []*NamePropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeNamePropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeNamePropertyIterator(i, aliasMap); err != nil {
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

// NewNameProperty creates a new name property.
func NewNameProperty() *NameProperty {
	return &NameProperty{alias: ""}
}

// AppendIRI appends an IRI value to the back of a list of the property "name"
func (this *NameProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &NamePropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  this.Len(),
		parent: this,
	})
}

// AppendLangString appends a langString value to the back of a list of the
// property "name". Invalidates iterators that are traversing using Prev.
func (this *NameProperty) AppendLangString(v map[string]string) {
	this.properties = append(this.properties, &NamePropertyIterator{
		alias:            this.alias,
		langStringMember: v,
		myIdx:            this.Len(),
		parent:           this,
	})
}

// AppendString appends a string value to the back of a list of the property
// "name". Invalidates iterators that are traversing using Prev.
func (this *NameProperty) AppendString(v string) {
	this.properties = append(this.properties, &NamePropertyIterator{
		alias:           this.alias,
		hasStringMember: true,
		myIdx:           this.Len(),
		parent:          this,
		stringMember:    v,
	})
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this NameProperty) At(index int) vocab.NamePropertyIteratorInterface {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this NameProperty) Begin() vocab.NamePropertyIteratorInterface {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this NameProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this NameProperty) End() vocab.NamePropertyIteratorInterface {
	return nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this NameProperty) JSONLDContext() map[string]string {
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
func (this NameProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "name" property.
func (this NameProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this NameProperty) Less(i, j int) bool {
	idx1 := this.KindIndex(i)
	idx2 := this.KindIndex(j)
	if idx1 < idx2 {
		return true
	} else if idx1 == idx2 {
		if idx1 == 0 {
			lhs := this.properties[i].GetString()
			rhs := this.properties[j].GetString()
			return string1.LessString(lhs, rhs)
		} else if idx1 == 1 {
			lhs := this.properties[i].GetLangString()
			rhs := this.properties[j].GetLangString()
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
func (this NameProperty) LessThan(o vocab.NamePropertyInterface) bool {
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
func (this NameProperty) Name() string {
	return "name"
}

// PrependIRI prepends an IRI value to the front of a list of the property "name".
func (this *NameProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*NamePropertyIterator{{
		alias:  this.alias,
		iri:    v,
		myIdx:  0,
		parent: this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependLangString prepends a langString value to the front of a list of the
// property "name". Invalidates all iterators.
func (this *NameProperty) PrependLangString(v map[string]string) {
	this.properties = append([]*NamePropertyIterator{{
		alias:            this.alias,
		langStringMember: v,
		myIdx:            0,
		parent:           this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependString prepends a string value to the front of a list of the property
// "name". Invalidates all iterators.
func (this *NameProperty) PrependString(v string) {
	this.properties = append([]*NamePropertyIterator{{
		alias:           this.alias,
		hasStringMember: true,
		myIdx:           0,
		parent:          this,
		stringMember:    v,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Remove deletes an element at the specified index from a list of the property
// "name", regardless of its type. Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *NameProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &NamePropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this NameProperty) Serialize() (interface{}, error) {
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
func (this *NameProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &NamePropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
}

// SetLangString sets a langString value to be at the specified index for the
// property "name". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *NameProperty) SetLangString(idx int, v map[string]string) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &NamePropertyIterator{
		alias:            this.alias,
		langStringMember: v,
		myIdx:            idx,
		parent:           this,
	}
}

// SetString sets a string value to be at the specified index for the property
// "name". Panics if the index is out of bounds. Invalidates all iterators.
func (this *NameProperty) SetString(idx int, v string) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &NamePropertyIterator{
		alias:           this.alias,
		hasStringMember: true,
		myIdx:           idx,
		parent:          this,
		stringMember:    v,
	}
}

// Swap swaps the location of values at two indices for the "name" property.
func (this NameProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}