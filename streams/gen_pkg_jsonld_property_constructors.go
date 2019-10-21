package streams

import (
	propertyid "github.com/go-fed/activity/streams/impl/jsonld/property_id"
	propertytype "github.com/go-fed/activity/streams/impl/jsonld/property_type"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// NewJSONLDJSONLDTypeProperty creates a new JSONLDTypeProperty
func NewJSONLDTypeProperty() vocab.JSONLDTypeProperty {
	return propertytype.NewJSONLDTypeProperty()
}

// NewJSONLDJSONLDIdProperty creates a new JSONLDIdProperty
func NewJSONLDIdProperty() vocab.JSONLDIdProperty {
	return propertyid.NewJSONLDIdProperty()
}
