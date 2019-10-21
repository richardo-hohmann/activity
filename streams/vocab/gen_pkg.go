package vocab

// Type represents an ActivityStreams type.
type Type interface {
	// GetJSONLDId returns the "id" property if it exists, and nil otherwise.
	GetJSONLDId() JSONLDIdProperty
	// GetTypeName returns the ActivityStreams type name.
	GetTypeName() string
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format.
	Serialize() (map[string]interface{}, error)
	// SetJSONLDId sets the "id" property.
	SetJSONLDId(JSONLDIdProperty)
	// VocabularyURI returns the vocabulary's URI as a string.
	VocabularyURI() string
}
