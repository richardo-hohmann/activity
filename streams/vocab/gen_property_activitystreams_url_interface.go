package vocab

import "net/url"

// ActivityStreamsUrlPropertyIterator represents a single value for the "url"
// property.
type ActivityStreamsUrlPropertyIterator interface {
	// GetActivityStreamsLink returns the value of this property. When
	// IsActivityStreamsLink returns false, GetActivityStreamsLink will
	// return an arbitrary value.
	GetActivityStreamsLink() ActivityStreamsLink
	// GetActivityStreamsMention returns the value of this property. When
	// IsActivityStreamsMention returns false, GetActivityStreamsMention
	// will return an arbitrary value.
	GetActivityStreamsMention() ActivityStreamsMention
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetType returns the value in this property as a Type. Returns nil if
	// the value is not an ActivityStreams type, such as an IRI or another
	// value.
	GetType() Type
	// GetXMLSchemaAnyURI returns the value of this property. When
	// IsXMLSchemaAnyURI returns false, GetXMLSchemaAnyURI will return an
	// arbitrary value.
	GetXMLSchemaAnyURI() *url.URL
	// HasAny returns true if any of the different values is set.
	HasAny() bool
	// IsActivityStreamsLink returns true if this property has a type of
	// "Link". When true, use the GetActivityStreamsLink and
	// SetActivityStreamsLink methods to access and set this property.
	IsActivityStreamsLink() bool
	// IsActivityStreamsMention returns true if this property has a type of
	// "Mention". When true, use the GetActivityStreamsMention and
	// SetActivityStreamsMention methods to access and set this property.
	IsActivityStreamsMention() bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// IsXMLSchemaAnyURI returns true if this property has a type of "anyURI".
	// When true, use the GetXMLSchemaAnyURI and SetXMLSchemaAnyURI
	// methods to access and set this property.
	IsXMLSchemaAnyURI() bool
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
	LessThan(o ActivityStreamsUrlPropertyIterator) bool
	// Name returns the name of this property: "ActivityStreamsUrl".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() ActivityStreamsUrlPropertyIterator
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() ActivityStreamsUrlPropertyIterator
	// SetActivityStreamsLink sets the value of this property. Calling
	// IsActivityStreamsLink afterwards returns true.
	SetActivityStreamsLink(v ActivityStreamsLink)
	// SetActivityStreamsMention sets the value of this property. Calling
	// IsActivityStreamsMention afterwards returns true.
	SetActivityStreamsMention(v ActivityStreamsMention)
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetXMLSchemaAnyURI sets the value of this property. Calling
	// IsXMLSchemaAnyURI afterwards returns true.
	SetXMLSchemaAnyURI(v *url.URL)
}

// Identifies one or more links to representations of the object
//
// Example 109 (https://www.w3.org/TR/activitystreams-vocabulary/#ex124-jsonld):
//   {
//     "name": "4Q Sales Forecast",
//     "type": "Document",
//     "url": "http://example.org/4q-sales-forecast.pdf"
//   }
//
// Example 110 (https://www.w3.org/TR/activitystreams-vocabulary/#ex125-jsonld):
//   {
//     "name": "4Q Sales Forecast",
//     "type": "Document",
//     "url": {
//       "type": "owl:Class",
//       "url": "http://example.org/4q-sales-forecast.pdf"
//     }
//   }
//
// Example 111 (https://www.w3.org/TR/activitystreams-vocabulary/#ex126-jsonld):
//   {
//     "name": "4Q Sales Forecast",
//     "type": "Document",
//     "url": [
//       {
//         "mediaType": "application/pdf",
//         "type": "owl:Class",
//         "url": "http://example.org/4q-sales-forecast.pdf"
//       },
//       {
//         "mediaType": "text/html",
//         "type": "owl:Class",
//         "url": "http://example.org/4q-sales-forecast.html"
//       }
//     ]
//   }
type ActivityStreamsUrlProperty interface {
	// AppendActivityStreamsLink appends a Link value to the back of a list of
	// the property "url". Invalidates iterators that are traversing using
	// Prev.
	AppendActivityStreamsLink(v ActivityStreamsLink)
	// AppendActivityStreamsMention appends a Mention value to the back of a
	// list of the property "url". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsMention(v ActivityStreamsMention)
	// AppendIRI appends an IRI value to the back of a list of the property
	// "url"
	AppendIRI(v *url.URL)
	// AppendXMLSchemaAnyURI appends a anyURI value to the back of a list of
	// the property "url". Invalidates iterators that are traversing using
	// Prev.
	AppendXMLSchemaAnyURI(v *url.URL)
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) ActivityStreamsUrlPropertyIterator
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() ActivityStreamsUrlPropertyIterator
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() ActivityStreamsUrlPropertyIterator
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
	// Len returns the number of values that exist for the "url" property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ActivityStreamsUrlProperty) bool
	// Name returns the name of this property: "url".
	Name() string
	// PrependActivityStreamsLink prepends a Link value to the front of a list
	// of the property "url". Invalidates all iterators.
	PrependActivityStreamsLink(v ActivityStreamsLink)
	// PrependActivityStreamsMention prepends a Mention value to the front of
	// a list of the property "url". Invalidates all iterators.
	PrependActivityStreamsMention(v ActivityStreamsMention)
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "url".
	PrependIRI(v *url.URL)
	// PrependXMLSchemaAnyURI prepends a anyURI value to the front of a list
	// of the property "url". Invalidates all iterators.
	PrependXMLSchemaAnyURI(v *url.URL)
	// Remove deletes an element at the specified index from a list of the
	// property "url", regardless of its type. Panics if the index is out
	// of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetActivityStreamsLink sets a Link value to be at the specified index
	// for the property "url". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsLink(idx int, v ActivityStreamsLink)
	// SetActivityStreamsMention sets a Mention value to be at the specified
	// index for the property "url". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsMention(idx int, v ActivityStreamsMention)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "url". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// SetXMLSchemaAnyURI sets a anyURI value to be at the specified index for
	// the property "url". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetXMLSchemaAnyURI(idx int, v *url.URL)
	// Swap swaps the location of values at two indices for the "url" property.
	Swap(i, j int)
}
