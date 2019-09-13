package streams

import (
	"context"
	"errors"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// TypeResolver resolves ActivityStreams values based on their type name.
type TypeResolver struct {
	callbacks []interface{}
}

// NewTypeResolver creates a new Resolver that examines the type of an
// ActivityStream value to determine what callback function to pass the
// concretely typed value. The callback is guaranteed to receive a value whose
// underlying ActivityStreams type matches the concrete interface name in its
// signature. The callback functions must be of the form:
//
//   func(context.Context, <TypeInterface>) error
//
// where TypeInterface is the code-generated interface for an ActivityStream
// type. An error is returned if a callback function does not match this
// signature.
func NewTypeResolver(callbacks ...interface{}) (*TypeResolver, error) {
	for _, cb := range callbacks {
		// Each callback function must satisfy one known function signature, or else we will generate a runtime error instead of silently fail.
		switch cb.(type) {
		case func(context.Context, vocab.ActivityStreamsAccept) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsActivity) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsAdd) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsAnnounce) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsApplication) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsArrive) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsArticle) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsAudio) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsBlock) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsCollection) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsCollectionPage) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsCreate) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsDelete) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsDislike) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsDocument) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsEvent) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsFlag) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsFollow) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsGroup) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsIgnore) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsImage) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsIntransitiveActivity) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsInvite) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsJoin) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsLeave) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsLike) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsLink) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsListen) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsMention) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsMove) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsNote) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsObject) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsOffer) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsOrderedCollection) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsOrderedCollectionPage) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsOrganization) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsPage) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsPerson) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsPlace) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsProfile) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsPublicKey) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsQuestion) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsRead) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsReject) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsRelationship) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsRemove) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsService) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsTentativeAccept) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsTentativeReject) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsTombstone) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsTravel) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsUndo) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsUpdate) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsVideo) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsView) error:
			// Do nothing, this callback has a correct signature.
		default:
			return nil, errors.New("a callback function is of the wrong signature and would never be called")
		}
	}
	return &TypeResolver{callbacks: callbacks}, nil
}

// Resolve applies the first callback function whose signature accepts the
// ActivityStreams value's type. This strictly assures that the callback
// function will only be passed ActivityStream objects whose type matches its
// interface. Returns an error if the ActivityStreams type does not match
// callbackers, is not a type handled by the generated code, or the value
// passed in is not go-fed compatible.
func (this TypeResolver) Resolve(ctx context.Context, o ActivityStreamsInterface) error {
	for _, i := range this.callbacks {
		if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Accept" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsAccept) error); ok {
				if v, ok := o.(vocab.ActivityStreamsAccept); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Activity" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsActivity) error); ok {
				if v, ok := o.(vocab.ActivityStreamsActivity); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Add" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsAdd) error); ok {
				if v, ok := o.(vocab.ActivityStreamsAdd); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Announce" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsAnnounce) error); ok {
				if v, ok := o.(vocab.ActivityStreamsAnnounce); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Application" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsApplication) error); ok {
				if v, ok := o.(vocab.ActivityStreamsApplication); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Arrive" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsArrive) error); ok {
				if v, ok := o.(vocab.ActivityStreamsArrive); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Article" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsArticle) error); ok {
				if v, ok := o.(vocab.ActivityStreamsArticle); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Audio" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsAudio) error); ok {
				if v, ok := o.(vocab.ActivityStreamsAudio); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Block" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsBlock) error); ok {
				if v, ok := o.(vocab.ActivityStreamsBlock); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Collection" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsCollection) error); ok {
				if v, ok := o.(vocab.ActivityStreamsCollection); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "CollectionPage" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsCollectionPage) error); ok {
				if v, ok := o.(vocab.ActivityStreamsCollectionPage); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Create" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsCreate) error); ok {
				if v, ok := o.(vocab.ActivityStreamsCreate); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Delete" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsDelete) error); ok {
				if v, ok := o.(vocab.ActivityStreamsDelete); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Dislike" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsDislike) error); ok {
				if v, ok := o.(vocab.ActivityStreamsDislike); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Document" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsDocument) error); ok {
				if v, ok := o.(vocab.ActivityStreamsDocument); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Event" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsEvent) error); ok {
				if v, ok := o.(vocab.ActivityStreamsEvent); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Flag" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsFlag) error); ok {
				if v, ok := o.(vocab.ActivityStreamsFlag); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Follow" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsFollow) error); ok {
				if v, ok := o.(vocab.ActivityStreamsFollow); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Group" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsGroup) error); ok {
				if v, ok := o.(vocab.ActivityStreamsGroup); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Ignore" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsIgnore) error); ok {
				if v, ok := o.(vocab.ActivityStreamsIgnore); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Image" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsImage) error); ok {
				if v, ok := o.(vocab.ActivityStreamsImage); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "IntransitiveActivity" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsIntransitiveActivity) error); ok {
				if v, ok := o.(vocab.ActivityStreamsIntransitiveActivity); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Invite" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsInvite) error); ok {
				if v, ok := o.(vocab.ActivityStreamsInvite); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Join" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsJoin) error); ok {
				if v, ok := o.(vocab.ActivityStreamsJoin); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Leave" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsLeave) error); ok {
				if v, ok := o.(vocab.ActivityStreamsLeave); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Like" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsLike) error); ok {
				if v, ok := o.(vocab.ActivityStreamsLike); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Link" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsLink) error); ok {
				if v, ok := o.(vocab.ActivityStreamsLink); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Listen" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsListen) error); ok {
				if v, ok := o.(vocab.ActivityStreamsListen); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Mention" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsMention) error); ok {
				if v, ok := o.(vocab.ActivityStreamsMention); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Move" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsMove) error); ok {
				if v, ok := o.(vocab.ActivityStreamsMove); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Note" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsNote) error); ok {
				if v, ok := o.(vocab.ActivityStreamsNote); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Object" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsObject) error); ok {
				if v, ok := o.(vocab.ActivityStreamsObject); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Offer" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsOffer) error); ok {
				if v, ok := o.(vocab.ActivityStreamsOffer); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "OrderedCollection" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsOrderedCollection) error); ok {
				if v, ok := o.(vocab.ActivityStreamsOrderedCollection); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "OrderedCollectionPage" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsOrderedCollectionPage) error); ok {
				if v, ok := o.(vocab.ActivityStreamsOrderedCollectionPage); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Organization" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsOrganization) error); ok {
				if v, ok := o.(vocab.ActivityStreamsOrganization); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Page" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsPage) error); ok {
				if v, ok := o.(vocab.ActivityStreamsPage); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Person" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsPerson) error); ok {
				if v, ok := o.(vocab.ActivityStreamsPerson); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Place" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsPlace) error); ok {
				if v, ok := o.(vocab.ActivityStreamsPlace); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Profile" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsProfile) error); ok {
				if v, ok := o.(vocab.ActivityStreamsProfile); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "PublicKey" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsPublicKey) error); ok {
				if v, ok := o.(vocab.ActivityStreamsPublicKey); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Question" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsQuestion) error); ok {
				if v, ok := o.(vocab.ActivityStreamsQuestion); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Read" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsRead) error); ok {
				if v, ok := o.(vocab.ActivityStreamsRead); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Reject" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsReject) error); ok {
				if v, ok := o.(vocab.ActivityStreamsReject); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Relationship" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsRelationship) error); ok {
				if v, ok := o.(vocab.ActivityStreamsRelationship); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Remove" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsRemove) error); ok {
				if v, ok := o.(vocab.ActivityStreamsRemove); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Service" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsService) error); ok {
				if v, ok := o.(vocab.ActivityStreamsService); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "TentativeAccept" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsTentativeAccept) error); ok {
				if v, ok := o.(vocab.ActivityStreamsTentativeAccept); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "TentativeReject" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsTentativeReject) error); ok {
				if v, ok := o.(vocab.ActivityStreamsTentativeReject); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Tombstone" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsTombstone) error); ok {
				if v, ok := o.(vocab.ActivityStreamsTombstone); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Travel" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsTravel) error); ok {
				if v, ok := o.(vocab.ActivityStreamsTravel); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Undo" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsUndo) error); ok {
				if v, ok := o.(vocab.ActivityStreamsUndo); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Update" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsUpdate) error); ok {
				if v, ok := o.(vocab.ActivityStreamsUpdate); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Video" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsVideo) error); ok {
				if v, ok := o.(vocab.ActivityStreamsVideo); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "View" {
			if fn, ok := i.(func(context.Context, vocab.ActivityStreamsView) error); ok {
				if v, ok := o.(vocab.ActivityStreamsView); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		} else {
			return ErrUnhandledType
		}
	}
	return ErrNoCallbackMatch
}
