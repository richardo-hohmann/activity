package propertyinstrument

import vocab "github.com/go-fed/activity/streams/vocab"

var mgr privateManager

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializeAcceptActivityStreams returns the deserialization method for
	// the "ActivityStreamsAccept" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeAcceptActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAccept, error)
	// DeserializeActivityActivityStreams returns the deserialization method
	// for the "ActivityStreamsActivity" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeActivityActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsActivity, error)
	// DeserializeAddActivityStreams returns the deserialization method for
	// the "ActivityStreamsAdd" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeAddActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAdd, error)
	// DeserializeAnnounceActivityStreams returns the deserialization method
	// for the "ActivityStreamsAnnounce" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeAnnounceActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAnnounce, error)
	// DeserializeApplicationActivityStreams returns the deserialization
	// method for the "ActivityStreamsApplication" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeApplicationActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsApplication, error)
	// DeserializeArriveActivityStreams returns the deserialization method for
	// the "ActivityStreamsArrive" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeArriveActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsArrive, error)
	// DeserializeArticleActivityStreams returns the deserialization method
	// for the "ActivityStreamsArticle" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeArticleActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsArticle, error)
	// DeserializeAudioActivityStreams returns the deserialization method for
	// the "ActivityStreamsAudio" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeAudioActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAudio, error)
	// DeserializeBlockActivityStreams returns the deserialization method for
	// the "ActivityStreamsBlock" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeBlockActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsBlock, error)
	// DeserializeCollectionActivityStreams returns the deserialization method
	// for the "ActivityStreamsCollection" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeCollectionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsCollection, error)
	// DeserializeCollectionPageActivityStreams returns the deserialization
	// method for the "ActivityStreamsCollectionPage" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeCollectionPageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsCollectionPage, error)
	// DeserializeCreateActivityStreams returns the deserialization method for
	// the "ActivityStreamsCreate" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeCreateActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsCreate, error)
	// DeserializeDeleteActivityStreams returns the deserialization method for
	// the "ActivityStreamsDelete" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeDeleteActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsDelete, error)
	// DeserializeDislikeActivityStreams returns the deserialization method
	// for the "ActivityStreamsDislike" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeDislikeActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsDislike, error)
	// DeserializeDocumentActivityStreams returns the deserialization method
	// for the "ActivityStreamsDocument" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeDocumentActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsDocument, error)
	// DeserializeEventActivityStreams returns the deserialization method for
	// the "ActivityStreamsEvent" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeEventActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsEvent, error)
	// DeserializeFlagActivityStreams returns the deserialization method for
	// the "ActivityStreamsFlag" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeFlagActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsFlag, error)
	// DeserializeFollowActivityStreams returns the deserialization method for
	// the "ActivityStreamsFollow" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeFollowActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsFollow, error)
	// DeserializeGroupActivityStreams returns the deserialization method for
	// the "ActivityStreamsGroup" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeGroupActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsGroup, error)
	// DeserializeIgnoreActivityStreams returns the deserialization method for
	// the "ActivityStreamsIgnore" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeIgnoreActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsIgnore, error)
	// DeserializeImageActivityStreams returns the deserialization method for
	// the "ActivityStreamsImage" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeImageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsImage, error)
	// DeserializeIntransitiveActivityActivityStreams returns the
	// deserialization method for the
	// "ActivityStreamsIntransitiveActivity" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeIntransitiveActivityActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsIntransitiveActivity, error)
	// DeserializeInviteActivityStreams returns the deserialization method for
	// the "ActivityStreamsInvite" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeInviteActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsInvite, error)
	// DeserializeJoinActivityStreams returns the deserialization method for
	// the "ActivityStreamsJoin" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeJoinActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsJoin, error)
	// DeserializeLeaveActivityStreams returns the deserialization method for
	// the "ActivityStreamsLeave" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeLeaveActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLeave, error)
	// DeserializeLikeActivityStreams returns the deserialization method for
	// the "ActivityStreamsLike" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeLikeActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLike, error)
	// DeserializeLinkActivityStreams returns the deserialization method for
	// the "ActivityStreamsLink" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeLinkActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLink, error)
	// DeserializeListenActivityStreams returns the deserialization method for
	// the "ActivityStreamsListen" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeListenActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsListen, error)
	// DeserializeMentionActivityStreams returns the deserialization method
	// for the "ActivityStreamsMention" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeMentionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsMention, error)
	// DeserializeMoveActivityStreams returns the deserialization method for
	// the "ActivityStreamsMove" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeMoveActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsMove, error)
	// DeserializeNoteActivityStreams returns the deserialization method for
	// the "ActivityStreamsNote" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeNoteActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsNote, error)
	// DeserializeObjectActivityStreams returns the deserialization method for
	// the "ActivityStreamsObject" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeObjectActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsObject, error)
	// DeserializeOfferActivityStreams returns the deserialization method for
	// the "ActivityStreamsOffer" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeOfferActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOffer, error)
	// DeserializeOrderedCollectionActivityStreams returns the deserialization
	// method for the "ActivityStreamsOrderedCollection" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeOrderedCollectionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOrderedCollection, error)
	// DeserializeOrderedCollectionPageActivityStreams returns the
	// deserialization method for the
	// "ActivityStreamsOrderedCollectionPage" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeOrderedCollectionPageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOrderedCollectionPage, error)
	// DeserializeOrganizationActivityStreams returns the deserialization
	// method for the "ActivityStreamsOrganization" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeOrganizationActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOrganization, error)
	// DeserializePageActivityStreams returns the deserialization method for
	// the "ActivityStreamsPage" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializePageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPage, error)
	// DeserializePersonActivityStreams returns the deserialization method for
	// the "ActivityStreamsPerson" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializePersonActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPerson, error)
	// DeserializePlaceActivityStreams returns the deserialization method for
	// the "ActivityStreamsPlace" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializePlaceActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPlace, error)
	// DeserializeProfileActivityStreams returns the deserialization method
	// for the "ActivityStreamsProfile" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeProfileActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsProfile, error)
	// DeserializePublicKeyW3IDSecurityV1 returns the deserialization method
	// for the "W3IDSecurityV1PublicKey" non-functional property in the
	// vocabulary "W3IDSecurityV1"
	DeserializePublicKeyW3IDSecurityV1() func(map[string]interface{}, map[string]string) (vocab.W3IDSecurityV1PublicKey, error)
	// DeserializeQuestionActivityStreams returns the deserialization method
	// for the "ActivityStreamsQuestion" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeQuestionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsQuestion, error)
	// DeserializeReadActivityStreams returns the deserialization method for
	// the "ActivityStreamsRead" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeReadActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsRead, error)
	// DeserializeRejectActivityStreams returns the deserialization method for
	// the "ActivityStreamsReject" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeRejectActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsReject, error)
	// DeserializeRelationshipActivityStreams returns the deserialization
	// method for the "ActivityStreamsRelationship" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeRelationshipActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsRelationship, error)
	// DeserializeRemoveActivityStreams returns the deserialization method for
	// the "ActivityStreamsRemove" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeRemoveActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsRemove, error)
	// DeserializeServiceActivityStreams returns the deserialization method
	// for the "ActivityStreamsService" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeServiceActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsService, error)
	// DeserializeTentativeAcceptActivityStreams returns the deserialization
	// method for the "ActivityStreamsTentativeAccept" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeTentativeAcceptActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTentativeAccept, error)
	// DeserializeTentativeRejectActivityStreams returns the deserialization
	// method for the "ActivityStreamsTentativeReject" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeTentativeRejectActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTentativeReject, error)
	// DeserializeTombstoneActivityStreams returns the deserialization method
	// for the "ActivityStreamsTombstone" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeTombstoneActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTombstone, error)
	// DeserializeTravelActivityStreams returns the deserialization method for
	// the "ActivityStreamsTravel" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeTravelActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTravel, error)
	// DeserializeUndoActivityStreams returns the deserialization method for
	// the "ActivityStreamsUndo" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeUndoActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsUndo, error)
	// DeserializeUpdateActivityStreams returns the deserialization method for
	// the "ActivityStreamsUpdate" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeUpdateActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsUpdate, error)
	// DeserializeVideoActivityStreams returns the deserialization method for
	// the "ActivityStreamsVideo" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeVideoActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsVideo, error)
	// DeserializeViewActivityStreams returns the deserialization method for
	// the "ActivityStreamsView" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeViewActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsView, error)
}

// SetManager sets the manager package-global variable. For internal use only, do
// not use as part of Application behavior. Must be called at golang init time.
func SetManager(m privateManager) {
	mgr = m
}
