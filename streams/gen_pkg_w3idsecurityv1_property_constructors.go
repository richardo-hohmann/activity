package streams

import (
	propertyowner "github.com/go-fed/activity/streams/impl/w3idsecurityv1/property_owner"
	propertypublickey "github.com/go-fed/activity/streams/impl/w3idsecurityv1/property_publickey"
	propertypublickeypem "github.com/go-fed/activity/streams/impl/w3idsecurityv1/property_publickeypem"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// NewW3IDSecurityV1W3IDSecurityV1OwnerProperty creates a new
// W3IDSecurityV1OwnerProperty
func NewW3IDSecurityV1OwnerProperty() vocab.W3IDSecurityV1OwnerProperty {
	return propertyowner.NewW3IDSecurityV1OwnerProperty()
}

// NewW3IDSecurityV1W3IDSecurityV1PublicKeyProperty creates a new
// W3IDSecurityV1PublicKeyProperty
func NewW3IDSecurityV1PublicKeyProperty() vocab.W3IDSecurityV1PublicKeyProperty {
	return propertypublickey.NewW3IDSecurityV1PublicKeyProperty()
}

// NewW3IDSecurityV1W3IDSecurityV1PublicKeyPemProperty creates a new
// W3IDSecurityV1PublicKeyPemProperty
func NewW3IDSecurityV1PublicKeyPemProperty() vocab.W3IDSecurityV1PublicKeyPemProperty {
	return propertypublickeypem.NewW3IDSecurityV1PublicKeyPemProperty()
}
