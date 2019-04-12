package propertyrel

import (
	"fmt"
	rfc5988 "github.com/go-fed/activity/streams/values/rfc5988"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// ActivityStreamsRelPropertyIterator is an iterator for a property. It is
// permitted to be a single default-valued value type.
type ActivityStreamsRelPropertyIterator struct {
	rfcRfc5988Member string
	hasRfc5988Member bool
	unknown          interface{}
	iri              *url.URL
	alias            string
	myIdx            int
	parent           vocab.ActivityStreamsRelProperty
}

// NewActivityStreamsRelPropertyIterator creates a new ActivityStreamsRel property.
func NewActivityStreamsRelPropertyIterator() *ActivityStreamsRelPropertyIterator {
	return &ActivityStreamsRelPropertyIterator{alias: ""}
}

// deserializeActivityStreamsRelPropertyIterator creates an iterator from an
// element that has been unmarshalled from a text or binary format.
func deserializeActivityStreamsRelPropertyIterator(i interface{}, aliasMap map[string]string) (*ActivityStreamsRelPropertyIterator, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/ns/activitystreams"]; ok {
		alias = a
	}
	if s, ok := i.(string); ok {
		u, err := url.Parse(s)
		// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
		// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
		if err == nil && len(u.Scheme) > 0 {
			this := &ActivityStreamsRelPropertyIterator{
				alias: alias,
				iri:   u,
			}
			return this, nil
		}
	}
	if v, err := rfc5988.DeserializeRfc5988(i); err == nil {
		this := &ActivityStreamsRelPropertyIterator{
			alias:            alias,
			hasRfc5988Member: true,
			rfcRfc5988Member: v,
		}
		return this, nil
	}
	this := &ActivityStreamsRelPropertyIterator{
		alias:   alias,
		unknown: i,
	}
	return this, nil
}

// Get returns the value of this property. When IsRFCRfc5988 returns false, Get
// will return any arbitrary value.
func (this ActivityStreamsRelPropertyIterator) Get() string {
	return this.rfcRfc5988Member
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this ActivityStreamsRelPropertyIterator) GetIRI() *url.URL {
	return this.iri
}

// HasAny returns true if the value or IRI is set.
func (this ActivityStreamsRelPropertyIterator) HasAny() bool {
	return this.IsRFCRfc5988() || this.iri != nil
}

// IsIRI returns true if this property is an IRI.
func (this ActivityStreamsRelPropertyIterator) IsIRI() bool {
	return this.iri != nil
}

// IsRFCRfc5988 returns true if this property is set and not an IRI.
func (this ActivityStreamsRelPropertyIterator) IsRFCRfc5988() bool {
	return this.hasRfc5988Member
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsRelPropertyIterator) JSONLDContext() map[string]string {
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
func (this ActivityStreamsRelPropertyIterator) KindIndex() int {
	if this.IsRFCRfc5988() {
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
func (this ActivityStreamsRelPropertyIterator) LessThan(o vocab.ActivityStreamsRelPropertyIterator) bool {
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
	if !this.IsRFCRfc5988() && !o.IsRFCRfc5988() {
		// Both are unknowns.
		return false
	} else if this.IsRFCRfc5988() && !o.IsRFCRfc5988() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsRFCRfc5988() && o.IsRFCRfc5988() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return rfc5988.LessRfc5988(this.Get(), o.Get())
	}
}

// Name returns the name of this property: "ActivityStreamsRel".
func (this ActivityStreamsRelPropertyIterator) Name() string {
	return "ActivityStreamsRel"
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this ActivityStreamsRelPropertyIterator) Next() vocab.ActivityStreamsRelPropertyIterator {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this ActivityStreamsRelPropertyIterator) Prev() vocab.ActivityStreamsRelPropertyIterator {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// Set sets the value of this property. Calling IsRFCRfc5988 afterwards will
// return true.
func (this *ActivityStreamsRelPropertyIterator) Set(v string) {
	this.clear()
	this.rfcRfc5988Member = v
	this.hasRfc5988Member = true
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *ActivityStreamsRelPropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.iri = v
}

// clear ensures no value of this property is set. Calling IsRFCRfc5988 afterwards
// will return false.
func (this *ActivityStreamsRelPropertyIterator) clear() {
	this.unknown = nil
	this.iri = nil
	this.hasRfc5988Member = false
}

// serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsRelPropertyIterator) serialize() (interface{}, error) {
	if this.IsRFCRfc5988() {
		return rfc5988.SerializeRfc5988(this.Get())
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// ActivityStreamsRelProperty is the non-functional property "rel". It is
// permitted to have one or more values, and of different value types.
type ActivityStreamsRelProperty struct {
	properties []*ActivityStreamsRelPropertyIterator
	alias      string
}

// DeserializeRelProperty creates a "rel" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeRelProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsRelProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/ns/activitystreams"]; ok {
		alias = a
	}
	propName := "rel"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "rel")
	}
	i, ok := m[propName]

	if ok {
		this := &ActivityStreamsRelProperty{
			alias:      alias,
			properties: []*ActivityStreamsRelPropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeActivityStreamsRelPropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeActivityStreamsRelPropertyIterator(i, aliasMap); err != nil {
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

// NewActivityStreamsRelProperty creates a new rel property.
func NewActivityStreamsRelProperty() *ActivityStreamsRelProperty {
	return &ActivityStreamsRelProperty{alias: ""}
}

// AppendIRI appends an IRI value to the back of a list of the property "rel"
func (this *ActivityStreamsRelProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &ActivityStreamsRelPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  this.Len(),
		parent: this,
	})
}

// AppendRFCRfc5988 appends a rfc5988 value to the back of a list of the property
// "rel". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsRelProperty) AppendRFCRfc5988(v string) {
	this.properties = append(this.properties, &ActivityStreamsRelPropertyIterator{
		alias:            this.alias,
		hasRfc5988Member: true,
		myIdx:            this.Len(),
		parent:           this,
		rfcRfc5988Member: v,
	})
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this ActivityStreamsRelProperty) At(index int) vocab.ActivityStreamsRelPropertyIterator {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this ActivityStreamsRelProperty) Begin() vocab.ActivityStreamsRelPropertyIterator {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this ActivityStreamsRelProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this ActivityStreamsRelProperty) End() vocab.ActivityStreamsRelPropertyIterator {
	return nil
}

// Insert inserts an IRI value at the specified index for a property "rel".
// Existing elements at that index and higher are shifted back once.
// Invalidates all iterators.
func (this *ActivityStreamsRelProperty) InsertIRI(idx int, v *url.URL) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ActivityStreamsRelPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// InsertRFCRfc5988 inserts a rfc5988 value at the specified index for a property
// "rel". Existing elements at that index and higher are shifted back once.
// Invalidates all iterators.
func (this *ActivityStreamsRelProperty) InsertRFCRfc5988(idx int, v string) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ActivityStreamsRelPropertyIterator{
		alias:            this.alias,
		hasRfc5988Member: true,
		myIdx:            idx,
		parent:           this,
		rfcRfc5988Member: v,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsRelProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/ns/activitystreams": this.alias}
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
func (this ActivityStreamsRelProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "rel" property.
func (this ActivityStreamsRelProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this ActivityStreamsRelProperty) Less(i, j int) bool {
	idx1 := this.KindIndex(i)
	idx2 := this.KindIndex(j)
	if idx1 < idx2 {
		return true
	} else if idx1 == idx2 {
		if idx1 == 0 {
			lhs := this.properties[i].Get()
			rhs := this.properties[j].Get()
			return rfc5988.LessRfc5988(lhs, rhs)
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
func (this ActivityStreamsRelProperty) LessThan(o vocab.ActivityStreamsRelProperty) bool {
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

// Name returns the name of this property: "rel".
func (this ActivityStreamsRelProperty) Name() string {
	return "rel"
}

// PrependIRI prepends an IRI value to the front of a list of the property "rel".
func (this *ActivityStreamsRelProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*ActivityStreamsRelPropertyIterator{{
		alias:  this.alias,
		iri:    v,
		myIdx:  0,
		parent: this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependRFCRfc5988 prepends a rfc5988 value to the front of a list of the
// property "rel". Invalidates all iterators.
func (this *ActivityStreamsRelProperty) PrependRFCRfc5988(v string) {
	this.properties = append([]*ActivityStreamsRelPropertyIterator{{
		alias:            this.alias,
		hasRfc5988Member: true,
		myIdx:            0,
		parent:           this,
		rfcRfc5988Member: v,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Remove deletes an element at the specified index from a list of the property
// "rel", regardless of its type. Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *ActivityStreamsRelProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &ActivityStreamsRelPropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsRelProperty) Serialize() (interface{}, error) {
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

// Set sets a rfc5988 value to be at the specified index for the property "rel".
// Panics if the index is out of bounds. Invalidates all iterators.
func (this *ActivityStreamsRelProperty) Set(idx int, v string) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsRelPropertyIterator{
		alias:            this.alias,
		hasRfc5988Member: true,
		myIdx:            idx,
		parent:           this,
		rfcRfc5988Member: v,
	}
}

// SetIRI sets an IRI value to be at the specified index for the property "rel".
// Panics if the index is out of bounds.
func (this *ActivityStreamsRelProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsRelPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
}

// Swap swaps the location of values at two indices for the "rel" property.
func (this ActivityStreamsRelProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}
