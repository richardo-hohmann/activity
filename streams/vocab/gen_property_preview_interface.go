package vocab

import "net/url"

// PreviewPropertyIteratorInterface represents a single value for the "preview"
// property.
type PreviewPropertyIteratorInterface interface {
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetLink returns the value of this property. When IsLink returns false,
	// GetLink will return an arbitrary value.
	GetLink() LinkInterface
	// GetObject returns the value of this property. When IsObject returns
	// false, GetObject will return an arbitrary value.
	GetObject() ObjectInterface
	// HasAny returns true if any of the different values is set.
	HasAny() bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// IsLink returns true if this property has a type of "Link". When true,
	// use the GetLink and SetLink methods to access and set this property.
	IsLink() bool
	// IsObject returns true if this property has a type of "Object". When
	// true, use the GetObject and SetObject methods to access and set
	// this property.
	IsObject() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o PreviewPropertyIteratorInterface) bool
	// Name returns the name of this property: "preview".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() PreviewPropertyIteratorInterface
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() PreviewPropertyIteratorInterface
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetLink sets the value of this property. Calling IsLink afterwards
	// returns true.
	SetLink(v LinkInterface)
	// SetObject sets the value of this property. Calling IsObject afterwards
	// returns true.
	SetObject(v ObjectInterface)
}

// Identifies an entity that provides a preview of this object.
//
// Example 102 (https://www.w3.org/TR/activitystreams-vocabulary/#ex106-jsonld):
//   {
//     "duration": "PT2H30M",
//     "name": "Cool New Movie",
//     "preview": {
//       "duration": "PT1M",
//       "name": "Trailer",
//       "type": "Video",
//       "url": {
//         "mediaType": "video/mkv",
//         "url": "http://example.org/trailer.mkv"
//       }
//     },
//     "type": "Video"
//   }
type PreviewPropertyInterface interface {
	// AppendIRI appends an IRI value to the back of a list of the property
	// "preview"
	AppendIRI(v *url.URL)
	// AppendLink appends a Link value to the back of a list of the property
	// "preview". Invalidates iterators that are traversing using Prev.
	AppendLink(v LinkInterface)
	// AppendObject appends a Object value to the back of a list of the
	// property "preview". Invalidates iterators that are traversing using
	// Prev.
	AppendObject(v ObjectInterface)
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) PreviewPropertyIteratorInterface
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() PreviewPropertyIteratorInterface
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() PreviewPropertyIteratorInterface
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API method specifically needed only for alternate
	// implementations for go-fed. Applications should not use this
	// method. Panics if the index is out of bounds.
	KindIndex(idx int) int
	// Len returns the number of values that exist for the "preview" property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o PreviewPropertyInterface) bool
	// Name returns the name of this property: "preview".
	Name() string
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "preview".
	PrependIRI(v *url.URL)
	// PrependLink prepends a Link value to the front of a list of the
	// property "preview". Invalidates all iterators.
	PrependLink(v LinkInterface)
	// PrependObject prepends a Object value to the front of a list of the
	// property "preview". Invalidates all iterators.
	PrependObject(v ObjectInterface)
	// Remove deletes an element at the specified index from a list of the
	// property "preview", regardless of its type. Panics if the index is
	// out of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "preview". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// SetLink sets a Link value to be at the specified index for the property
	// "preview". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetLink(idx int, v LinkInterface)
	// SetObject sets a Object value to be at the specified index for the
	// property "preview". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetObject(idx int, v ObjectInterface)
	// Swap swaps the location of values at two indices for the "preview"
	// property.
	Swap(i, j int)
}
