package propertyinbox

import vocab "github.com/go-fed/activity/streams/vocab"

var mgr privateManager

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializeOrderedCollectionActivityStreams returns the deserialization
	// method for the "ActivityStreamsOrderedCollection" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeOrderedCollectionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOrderedCollection, error)
	// DeserializeOrderedCollectionPageActivityStreams returns the
	// deserialization method for the
	// "ActivityStreamsOrderedCollectionPage" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeOrderedCollectionPageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOrderedCollectionPage, error)
}

// SetManager sets the manager package-global variable. For internal use only, do
// not use as part of Application behavior. Must be called at golang init time.
func SetManager(m privateManager) {
	mgr = m
}
