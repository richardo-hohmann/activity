package propertyfollowing

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// FollowingProperty is the functional property "following". It is permitted to be
// one of multiple value types. At most, one type of value can be present, or
// none at all. Setting a value will clear the other types of values so that
// only one of the 'Is' methods will return true. It is possible to clear all
// values, so that this property is empty.
type FollowingProperty struct {
	OrderedCollectionMember     vocab.OrderedCollectionInterface
	CollectionMember            vocab.CollectionInterface
	CollectionPageMember        vocab.CollectionPageInterface
	OrderedCollectionPageMember vocab.OrderedCollectionPageInterface
	unknown                     interface{}
	iri                         *url.URL
	alias                       string
}

// DeserializeFollowingProperty creates a "following" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeFollowingProperty(m map[string]interface{}, aliasMap map[string]string) (*FollowingProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "following"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "following")
	}
	i, ok := m[propName]

	if ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
			if err == nil && len(u.Scheme) > 0 {
				this := &FollowingProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if m, ok := i.(map[string]interface{}); ok {
			if v, err := mgr.DeserializeOrderedCollectionActivityStreams()(m, aliasMap); err == nil {
				this := &FollowingProperty{
					OrderedCollectionMember: v,
					alias:                   alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeCollectionActivityStreams()(m, aliasMap); err == nil {
				this := &FollowingProperty{
					CollectionMember: v,
					alias:            alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeCollectionPageActivityStreams()(m, aliasMap); err == nil {
				this := &FollowingProperty{
					CollectionPageMember: v,
					alias:                alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeOrderedCollectionPageActivityStreams()(m, aliasMap); err == nil {
				this := &FollowingProperty{
					OrderedCollectionPageMember: v,
					alias:                       alias,
				}
				return this, nil
			}
		}
		this := &FollowingProperty{
			alias:   alias,
			unknown: i,
		}
		return this, nil
	}
	return nil, nil
}

// NewFollowingProperty creates a new following property.
func NewFollowingProperty() *FollowingProperty {
	return &FollowingProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling HasAny or any of the
// 'Is' methods afterwards will return false.
func (this *FollowingProperty) Clear() {
	this.OrderedCollectionMember = nil
	this.CollectionMember = nil
	this.CollectionPageMember = nil
	this.OrderedCollectionPageMember = nil
	this.unknown = nil
	this.iri = nil
}

// GetCollection returns the value of this property. When IsCollection returns
// false, GetCollection will return an arbitrary value.
func (this FollowingProperty) GetCollection() vocab.CollectionInterface {
	return this.CollectionMember
}

// GetCollectionPage returns the value of this property. When IsCollectionPage
// returns false, GetCollectionPage will return an arbitrary value.
func (this FollowingProperty) GetCollectionPage() vocab.CollectionPageInterface {
	return this.CollectionPageMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this FollowingProperty) GetIRI() *url.URL {
	return this.iri
}

// GetOrderedCollection returns the value of this property. When
// IsOrderedCollection returns false, GetOrderedCollection will return an
// arbitrary value.
func (this FollowingProperty) GetOrderedCollection() vocab.OrderedCollectionInterface {
	return this.OrderedCollectionMember
}

// GetOrderedCollectionPage returns the value of this property. When
// IsOrderedCollectionPage returns false, GetOrderedCollectionPage will return
// an arbitrary value.
func (this FollowingProperty) GetOrderedCollectionPage() vocab.OrderedCollectionPageInterface {
	return this.OrderedCollectionPageMember
}

// HasAny returns true if any of the different values is set.
func (this FollowingProperty) HasAny() bool {
	return this.IsOrderedCollection() ||
		this.IsCollection() ||
		this.IsCollectionPage() ||
		this.IsOrderedCollectionPage() ||
		this.iri != nil
}

// IsCollection returns true if this property has a type of "Collection". When
// true, use the GetCollection and SetCollection methods to access and set
// this property.
func (this FollowingProperty) IsCollection() bool {
	return this.CollectionMember != nil
}

// IsCollectionPage returns true if this property has a type of "CollectionPage".
// When true, use the GetCollectionPage and SetCollectionPage methods to
// access and set this property.
func (this FollowingProperty) IsCollectionPage() bool {
	return this.CollectionPageMember != nil
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this FollowingProperty) IsIRI() bool {
	return this.iri != nil
}

// IsOrderedCollection returns true if this property has a type of
// "OrderedCollection". When true, use the GetOrderedCollection and
// SetOrderedCollection methods to access and set this property.
func (this FollowingProperty) IsOrderedCollection() bool {
	return this.OrderedCollectionMember != nil
}

// IsOrderedCollectionPage returns true if this property has a type of
// "OrderedCollectionPage". When true, use the GetOrderedCollectionPage and
// SetOrderedCollectionPage methods to access and set this property.
func (this FollowingProperty) IsOrderedCollectionPage() bool {
	return this.OrderedCollectionPageMember != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this FollowingProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	var child map[string]string
	if this.IsOrderedCollection() {
		child = this.GetOrderedCollection().JSONLDContext()
	} else if this.IsCollection() {
		child = this.GetCollection().JSONLDContext()
	} else if this.IsCollectionPage() {
		child = this.GetCollectionPage().JSONLDContext()
	} else if this.IsOrderedCollectionPage() {
		child = this.GetOrderedCollectionPage().JSONLDContext()
	}
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
func (this FollowingProperty) KindIndex() int {
	if this.IsOrderedCollection() {
		return 0
	}
	if this.IsCollection() {
		return 1
	}
	if this.IsCollectionPage() {
		return 2
	}
	if this.IsOrderedCollectionPage() {
		return 3
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
func (this FollowingProperty) LessThan(o vocab.FollowingPropertyInterface) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsOrderedCollection() {
		return this.GetOrderedCollection().LessThan(o.GetOrderedCollection())
	} else if this.IsCollection() {
		return this.GetCollection().LessThan(o.GetCollection())
	} else if this.IsCollectionPage() {
		return this.GetCollectionPage().LessThan(o.GetCollectionPage())
	} else if this.IsOrderedCollectionPage() {
		return this.GetOrderedCollectionPage().LessThan(o.GetOrderedCollectionPage())
	} else if this.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	}
	return false
}

// Name returns the name of this property: "following".
func (this FollowingProperty) Name() string {
	return "following"
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this FollowingProperty) Serialize() (interface{}, error) {
	if this.IsOrderedCollection() {
		return this.GetOrderedCollection().Serialize()
	} else if this.IsCollection() {
		return this.GetCollection().Serialize()
	} else if this.IsCollectionPage() {
		return this.GetCollectionPage().Serialize()
	} else if this.IsOrderedCollectionPage() {
		return this.GetOrderedCollectionPage().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// SetCollection sets the value of this property. Calling IsCollection afterwards
// returns true.
func (this *FollowingProperty) SetCollection(v vocab.CollectionInterface) {
	this.Clear()
	this.CollectionMember = v
}

// SetCollectionPage sets the value of this property. Calling IsCollectionPage
// afterwards returns true.
func (this *FollowingProperty) SetCollectionPage(v vocab.CollectionPageInterface) {
	this.Clear()
	this.CollectionPageMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *FollowingProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}

// SetOrderedCollection sets the value of this property. Calling
// IsOrderedCollection afterwards returns true.
func (this *FollowingProperty) SetOrderedCollection(v vocab.OrderedCollectionInterface) {
	this.Clear()
	this.OrderedCollectionMember = v
}

// SetOrderedCollectionPage sets the value of this property. Calling
// IsOrderedCollectionPage afterwards returns true.
func (this *FollowingProperty) SetOrderedCollectionPage(v vocab.OrderedCollectionPageInterface) {
	this.Clear()
	this.OrderedCollectionPageMember = v
}
