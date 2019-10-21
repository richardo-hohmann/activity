package typepublickey

import vocab "github.com/go-fed/activity/streams/vocab"

// A public key represents a public cryptographical key for a user
type W3IDSecurityV1PublicKey struct {
	JSONLDId                   vocab.JSONLDIdProperty
	W3IDSecurityV1Owner        vocab.W3IDSecurityV1OwnerProperty
	W3IDSecurityV1PublicKeyPem vocab.W3IDSecurityV1PublicKeyPemProperty
	alias                      string
	unknown                    map[string]interface{}
}

// DeserializePublicKey creates a PublicKey from a map representation that has
// been unmarshalled from a text or binary format.
func DeserializePublicKey(m map[string]interface{}, aliasMap map[string]string) (*W3IDSecurityV1PublicKey, error) {
	alias := ""
	if a, ok := aliasMap["https://w3id.org/security/v1"]; ok {
		alias = a
	}
	this := &W3IDSecurityV1PublicKey{
		alias:   alias,
		unknown: make(map[string]interface{}),
	}

	// Begin: Known property deserialization
	if p, err := mgr.DeserializeIdPropertyJSONLD()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.JSONLDId = p
	}
	if p, err := mgr.DeserializeOwnerPropertyW3IDSecurityV1()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.W3IDSecurityV1Owner = p
	}
	if p, err := mgr.DeserializePublicKeyPemPropertyW3IDSecurityV1()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.W3IDSecurityV1PublicKeyPem = p
	}
	// End: Known property deserialization

	// Begin: Unknown deserialization
	for k, v := range m {
		// Begin: Code that ensures a property name is unknown
		if k == "id" {
			continue
		} else if k == "owner" {
			continue
		} else if k == "publicKeyPem" {
			continue
		} // End: Code that ensures a property name is unknown

		this.unknown[k] = v
	}
	// End: Unknown deserialization

	return this, nil
}

// IsOrExtendsPublicKey returns true if the other provided type is the PublicKey
// type or extends from the PublicKey type.
func IsOrExtendsPublicKey(other vocab.Type) bool {
	if other.GetTypeName() == "PublicKey" {
		return true
	}
	return PublicKeyIsExtendedBy(other)
}

// NewW3IDSecurityV1PublicKey creates a new PublicKey type
func NewW3IDSecurityV1PublicKey() *W3IDSecurityV1PublicKey {
	return &W3IDSecurityV1PublicKey{
		alias:   "",
		unknown: make(map[string]interface{}, 0),
	}
}

// PublicKeyIsDisjointWith returns true if the other provided type is disjoint
// with the PublicKey type.
func PublicKeyIsDisjointWith(other vocab.Type) bool {
	// Shortcut implementation: is not disjoint with anything.
	return false
}

// PublicKeyIsExtendedBy returns true if the other provided type extends from the
// PublicKey type. Note that it returns false if the types are the same; see
// the "IsOrExtendsPublicKey" variant instead.
func PublicKeyIsExtendedBy(other vocab.Type) bool {
	// Shortcut implementation: is not extended by anything.
	return false
}

// W3IDSecurityV1PublicKeyExtends returns true if the PublicKey type extends from
// the other type.
func W3IDSecurityV1PublicKeyExtends(other vocab.Type) bool {
	// Shortcut implementation: this does not extend anything.
	return false
}

// GetJSONLDId returns the "id" property if it exists, and nil otherwise.
func (this W3IDSecurityV1PublicKey) GetJSONLDId() vocab.JSONLDIdProperty {
	return this.JSONLDId
}

// GetTypeName returns the name of this type.
func (this W3IDSecurityV1PublicKey) GetTypeName() string {
	return "PublicKey"
}

// GetUnknownProperties returns the unknown properties for the PublicKey type.
// Note that this should not be used by app developers. It is only used to
// help determine which implementation is LessThan the other. Developers who
// are creating a different implementation of this type's interface can use
// this method in their LessThan implementation, but routine ActivityPub
// applications should not use this to bypass the code generation tool.
func (this W3IDSecurityV1PublicKey) GetUnknownProperties() map[string]interface{} {
	return this.unknown
}

// GetW3IDSecurityV1Owner returns the "owner" property if it exists, and nil
// otherwise.
func (this W3IDSecurityV1PublicKey) GetW3IDSecurityV1Owner() vocab.W3IDSecurityV1OwnerProperty {
	return this.W3IDSecurityV1Owner
}

// GetW3IDSecurityV1PublicKeyPem returns the "publicKeyPem" property if it exists,
// and nil otherwise.
func (this W3IDSecurityV1PublicKey) GetW3IDSecurityV1PublicKeyPem() vocab.W3IDSecurityV1PublicKeyPemProperty {
	return this.W3IDSecurityV1PublicKeyPem
}

// IsExtending returns true if the PublicKey type extends from the other type.
func (this W3IDSecurityV1PublicKey) IsExtending(other vocab.Type) bool {
	return W3IDSecurityV1PublicKeyExtends(other)
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// type and the specific properties that are set. The value in the map is the
// alias used to import the type and its properties.
func (this W3IDSecurityV1PublicKey) JSONLDContext() map[string]string {
	m := map[string]string{"https://w3id.org/security/v1": this.alias}
	m = this.helperJSONLDContext(this.JSONLDId, m)
	m = this.helperJSONLDContext(this.W3IDSecurityV1Owner, m)
	m = this.helperJSONLDContext(this.W3IDSecurityV1PublicKeyPem, m)

	return m
}

// LessThan computes if this PublicKey is lesser, with an arbitrary but stable
// determination.
func (this W3IDSecurityV1PublicKey) LessThan(o vocab.W3IDSecurityV1PublicKey) bool {
	// Begin: Compare known properties
	// Compare property "id"
	if lhs, rhs := this.JSONLDId, o.GetJSONLDId(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "owner"
	if lhs, rhs := this.W3IDSecurityV1Owner, o.GetW3IDSecurityV1Owner(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "publicKeyPem"
	if lhs, rhs := this.W3IDSecurityV1PublicKeyPem, o.GetW3IDSecurityV1PublicKeyPem(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// End: Compare known properties

	// Begin: Compare unknown properties (only by number of them)
	if len(this.unknown) < len(o.GetUnknownProperties()) {
		return true
	} else if len(o.GetUnknownProperties()) < len(this.unknown) {
		return false
	} // End: Compare unknown properties (only by number of them)

	// All properties are the same.
	return false
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format.
func (this W3IDSecurityV1PublicKey) Serialize() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	// Begin: Serialize known properties
	// Maybe serialize property "id"
	if this.JSONLDId != nil {
		if i, err := this.JSONLDId.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.JSONLDId.Name()] = i
		}
	}
	// Maybe serialize property "owner"
	if this.W3IDSecurityV1Owner != nil {
		if i, err := this.W3IDSecurityV1Owner.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.W3IDSecurityV1Owner.Name()] = i
		}
	}
	// Maybe serialize property "publicKeyPem"
	if this.W3IDSecurityV1PublicKeyPem != nil {
		if i, err := this.W3IDSecurityV1PublicKeyPem.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.W3IDSecurityV1PublicKeyPem.Name()] = i
		}
	}
	// End: Serialize known properties

	// Begin: Serialize unknown properties
	for k, v := range this.unknown {
		// To be safe, ensure we aren't overwriting a known property
		if _, has := m[k]; !has {
			m[k] = v
		}
	}
	// End: Serialize unknown properties

	return m, nil
}

// SetJSONLDId sets the "id" property.
func (this *W3IDSecurityV1PublicKey) SetJSONLDId(i vocab.JSONLDIdProperty) {
	this.JSONLDId = i
}

// SetW3IDSecurityV1Owner sets the "owner" property.
func (this *W3IDSecurityV1PublicKey) SetW3IDSecurityV1Owner(i vocab.W3IDSecurityV1OwnerProperty) {
	this.W3IDSecurityV1Owner = i
}

// SetW3IDSecurityV1PublicKeyPem sets the "publicKeyPem" property.
func (this *W3IDSecurityV1PublicKey) SetW3IDSecurityV1PublicKeyPem(i vocab.W3IDSecurityV1PublicKeyPemProperty) {
	this.W3IDSecurityV1PublicKeyPem = i
}

// VocabularyURI returns the vocabulary's URI as a string.
func (this W3IDSecurityV1PublicKey) VocabularyURI() string {
	return "https://w3id.org/security/v1"
}

// helperJSONLDContext obtains the context uris and their aliases from a property,
// if it is not nil.
func (this W3IDSecurityV1PublicKey) helperJSONLDContext(i jsonldContexter, toMerge map[string]string) map[string]string {
	if i == nil {
		return toMerge
	}
	for k, v := range i.JSONLDContext() {
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		toMerge[k] = v
	}
	return toMerge
}
