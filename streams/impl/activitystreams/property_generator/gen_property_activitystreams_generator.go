package propertygenerator

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// ActivityStreamsGeneratorPropertyIterator is an iterator for a property. It is
// permitted to be one of multiple value types. At most, one type of value can
// be present, or none at all. Setting a value will clear the other types of
// values so that only one of the 'Is' methods will return true. It is
// possible to clear all values, so that this property is empty.
type ActivityStreamsGeneratorPropertyIterator struct {
	activitystreamsObjectMember                vocab.ActivityStreamsObject
	activitystreamsLinkMember                  vocab.ActivityStreamsLink
	activitystreamsAcceptMember                vocab.ActivityStreamsAccept
	activitystreamsActivityMember              vocab.ActivityStreamsActivity
	activitystreamsAddMember                   vocab.ActivityStreamsAdd
	activitystreamsAnnounceMember              vocab.ActivityStreamsAnnounce
	activitystreamsApplicationMember           vocab.ActivityStreamsApplication
	activitystreamsArriveMember                vocab.ActivityStreamsArrive
	activitystreamsArticleMember               vocab.ActivityStreamsArticle
	activitystreamsAudioMember                 vocab.ActivityStreamsAudio
	activitystreamsBlockMember                 vocab.ActivityStreamsBlock
	activitystreamsCollectionMember            vocab.ActivityStreamsCollection
	activitystreamsCollectionPageMember        vocab.ActivityStreamsCollectionPage
	activitystreamsCreateMember                vocab.ActivityStreamsCreate
	activitystreamsDeleteMember                vocab.ActivityStreamsDelete
	activitystreamsDislikeMember               vocab.ActivityStreamsDislike
	activitystreamsDocumentMember              vocab.ActivityStreamsDocument
	activitystreamsEventMember                 vocab.ActivityStreamsEvent
	activitystreamsFlagMember                  vocab.ActivityStreamsFlag
	activitystreamsFollowMember                vocab.ActivityStreamsFollow
	activitystreamsGroupMember                 vocab.ActivityStreamsGroup
	activitystreamsIgnoreMember                vocab.ActivityStreamsIgnore
	activitystreamsImageMember                 vocab.ActivityStreamsImage
	activitystreamsIntransitiveActivityMember  vocab.ActivityStreamsIntransitiveActivity
	activitystreamsInviteMember                vocab.ActivityStreamsInvite
	activitystreamsJoinMember                  vocab.ActivityStreamsJoin
	activitystreamsLeaveMember                 vocab.ActivityStreamsLeave
	activitystreamsLikeMember                  vocab.ActivityStreamsLike
	activitystreamsListenMember                vocab.ActivityStreamsListen
	activitystreamsMentionMember               vocab.ActivityStreamsMention
	activitystreamsMoveMember                  vocab.ActivityStreamsMove
	activitystreamsNoteMember                  vocab.ActivityStreamsNote
	activitystreamsOfferMember                 vocab.ActivityStreamsOffer
	activitystreamsOrderedCollectionMember     vocab.ActivityStreamsOrderedCollection
	activitystreamsOrderedCollectionPageMember vocab.ActivityStreamsOrderedCollectionPage
	activitystreamsOrganizationMember          vocab.ActivityStreamsOrganization
	activitystreamsPageMember                  vocab.ActivityStreamsPage
	activitystreamsPersonMember                vocab.ActivityStreamsPerson
	activitystreamsPlaceMember                 vocab.ActivityStreamsPlace
	activitystreamsProfileMember               vocab.ActivityStreamsProfile
	activitystreamsQuestionMember              vocab.ActivityStreamsQuestion
	activitystreamsReadMember                  vocab.ActivityStreamsRead
	activitystreamsRejectMember                vocab.ActivityStreamsReject
	activitystreamsRelationshipMember          vocab.ActivityStreamsRelationship
	activitystreamsRemoveMember                vocab.ActivityStreamsRemove
	activitystreamsServiceMember               vocab.ActivityStreamsService
	activitystreamsTentativeAcceptMember       vocab.ActivityStreamsTentativeAccept
	activitystreamsTentativeRejectMember       vocab.ActivityStreamsTentativeReject
	activitystreamsTombstoneMember             vocab.ActivityStreamsTombstone
	activitystreamsTravelMember                vocab.ActivityStreamsTravel
	activitystreamsUndoMember                  vocab.ActivityStreamsUndo
	activitystreamsUpdateMember                vocab.ActivityStreamsUpdate
	activitystreamsVideoMember                 vocab.ActivityStreamsVideo
	activitystreamsViewMember                  vocab.ActivityStreamsView
	unknown                                    interface{}
	iri                                        *url.URL
	alias                                      string
	myIdx                                      int
	parent                                     vocab.ActivityStreamsGeneratorProperty
}

// NewActivityStreamsGeneratorPropertyIterator creates a new
// ActivityStreamsGenerator property.
func NewActivityStreamsGeneratorPropertyIterator() *ActivityStreamsGeneratorPropertyIterator {
	return &ActivityStreamsGeneratorPropertyIterator{alias: ""}
}

// deserializeActivityStreamsGeneratorPropertyIterator creates an iterator from an
// element that has been unmarshalled from a text or binary format.
func deserializeActivityStreamsGeneratorPropertyIterator(i interface{}, aliasMap map[string]string) (*ActivityStreamsGeneratorPropertyIterator, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	if s, ok := i.(string); ok {
		u, err := url.Parse(s)
		// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
		// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
		if err == nil && len(u.Scheme) > 0 {
			this := &ActivityStreamsGeneratorPropertyIterator{
				alias: alias,
				iri:   u,
			}
			return this, nil
		}
	}
	if m, ok := i.(map[string]interface{}); ok {
		if v, err := mgr.DeserializeObjectActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsObjectMember: v,
				alias:                       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeLinkActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsLinkMember: v,
				alias:                     alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeAcceptActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsAcceptMember: v,
				alias:                       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeActivityActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsActivityMember: v,
				alias:                         alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeAddActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsAddMember: v,
				alias:                    alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeAnnounceActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsAnnounceMember: v,
				alias:                         alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeApplicationActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsApplicationMember: v,
				alias:                            alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeArriveActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsArriveMember: v,
				alias:                       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeArticleActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsArticleMember: v,
				alias:                        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeAudioActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsAudioMember: v,
				alias:                      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeBlockActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsBlockMember: v,
				alias:                      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeCollectionActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsCollectionMember: v,
				alias:                           alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeCollectionPageActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsCollectionPageMember: v,
				alias:                               alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeCreateActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsCreateMember: v,
				alias:                       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeDeleteActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsDeleteMember: v,
				alias:                       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeDislikeActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsDislikeMember: v,
				alias:                        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeDocumentActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsDocumentMember: v,
				alias:                         alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeEventActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsEventMember: v,
				alias:                      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeFlagActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsFlagMember: v,
				alias:                     alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeFollowActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsFollowMember: v,
				alias:                       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeGroupActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsGroupMember: v,
				alias:                      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeIgnoreActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsIgnoreMember: v,
				alias:                       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeImageActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsImageMember: v,
				alias:                      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeIntransitiveActivityActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsIntransitiveActivityMember: v,
				alias: alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeInviteActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsInviteMember: v,
				alias:                       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeJoinActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsJoinMember: v,
				alias:                     alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeLeaveActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsLeaveMember: v,
				alias:                      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeLikeActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsLikeMember: v,
				alias:                     alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeListenActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsListenMember: v,
				alias:                       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeMentionActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsMentionMember: v,
				alias:                        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeMoveActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsMoveMember: v,
				alias:                     alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeNoteActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsNoteMember: v,
				alias:                     alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeOfferActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsOfferMember: v,
				alias:                      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeOrderedCollectionActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsOrderedCollectionMember: v,
				alias:                                  alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeOrderedCollectionPageActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsOrderedCollectionPageMember: v,
				alias: alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeOrganizationActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsOrganizationMember: v,
				alias:                             alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializePageActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsPageMember: v,
				alias:                     alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializePersonActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsPersonMember: v,
				alias:                       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializePlaceActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsPlaceMember: v,
				alias:                      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeProfileActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsProfileMember: v,
				alias:                        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeQuestionActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsQuestionMember: v,
				alias:                         alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeReadActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsReadMember: v,
				alias:                     alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeRejectActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsRejectMember: v,
				alias:                       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeRelationshipActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsRelationshipMember: v,
				alias:                             alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeRemoveActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsRemoveMember: v,
				alias:                       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeServiceActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsServiceMember: v,
				alias:                        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeTentativeAcceptActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsTentativeAcceptMember: v,
				alias:                                alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeTentativeRejectActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsTentativeRejectMember: v,
				alias:                                alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeTombstoneActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsTombstoneMember: v,
				alias:                          alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeTravelActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsTravelMember: v,
				alias:                       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeUndoActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsUndoMember: v,
				alias:                     alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeUpdateActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsUpdateMember: v,
				alias:                       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeVideoActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsVideoMember: v,
				alias:                      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeViewActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsGeneratorPropertyIterator{
				activitystreamsViewMember: v,
				alias:                     alias,
			}
			return this, nil
		}
	}
	this := &ActivityStreamsGeneratorPropertyIterator{
		alias:   alias,
		unknown: i,
	}
	return this, nil
}

// GetActivityStreamsAccept returns the value of this property. When
// IsActivityStreamsAccept returns false, GetActivityStreamsAccept will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsAccept() vocab.ActivityStreamsAccept {
	return this.activitystreamsAcceptMember
}

// GetActivityStreamsActivity returns the value of this property. When
// IsActivityStreamsActivity returns false, GetActivityStreamsActivity will
// return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsActivity() vocab.ActivityStreamsActivity {
	return this.activitystreamsActivityMember
}

// GetActivityStreamsAdd returns the value of this property. When
// IsActivityStreamsAdd returns false, GetActivityStreamsAdd will return an
// arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsAdd() vocab.ActivityStreamsAdd {
	return this.activitystreamsAddMember
}

// GetActivityStreamsAnnounce returns the value of this property. When
// IsActivityStreamsAnnounce returns false, GetActivityStreamsAnnounce will
// return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsAnnounce() vocab.ActivityStreamsAnnounce {
	return this.activitystreamsAnnounceMember
}

// GetActivityStreamsApplication returns the value of this property. When
// IsActivityStreamsApplication returns false, GetActivityStreamsApplication
// will return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsApplication() vocab.ActivityStreamsApplication {
	return this.activitystreamsApplicationMember
}

// GetActivityStreamsArrive returns the value of this property. When
// IsActivityStreamsArrive returns false, GetActivityStreamsArrive will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsArrive() vocab.ActivityStreamsArrive {
	return this.activitystreamsArriveMember
}

// GetActivityStreamsArticle returns the value of this property. When
// IsActivityStreamsArticle returns false, GetActivityStreamsArticle will
// return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsArticle() vocab.ActivityStreamsArticle {
	return this.activitystreamsArticleMember
}

// GetActivityStreamsAudio returns the value of this property. When
// IsActivityStreamsAudio returns false, GetActivityStreamsAudio will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsAudio() vocab.ActivityStreamsAudio {
	return this.activitystreamsAudioMember
}

// GetActivityStreamsBlock returns the value of this property. When
// IsActivityStreamsBlock returns false, GetActivityStreamsBlock will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsBlock() vocab.ActivityStreamsBlock {
	return this.activitystreamsBlockMember
}

// GetActivityStreamsCollection returns the value of this property. When
// IsActivityStreamsCollection returns false, GetActivityStreamsCollection
// will return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsCollection() vocab.ActivityStreamsCollection {
	return this.activitystreamsCollectionMember
}

// GetActivityStreamsCollectionPage returns the value of this property. When
// IsActivityStreamsCollectionPage returns false,
// GetActivityStreamsCollectionPage will return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsCollectionPage() vocab.ActivityStreamsCollectionPage {
	return this.activitystreamsCollectionPageMember
}

// GetActivityStreamsCreate returns the value of this property. When
// IsActivityStreamsCreate returns false, GetActivityStreamsCreate will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsCreate() vocab.ActivityStreamsCreate {
	return this.activitystreamsCreateMember
}

// GetActivityStreamsDelete returns the value of this property. When
// IsActivityStreamsDelete returns false, GetActivityStreamsDelete will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsDelete() vocab.ActivityStreamsDelete {
	return this.activitystreamsDeleteMember
}

// GetActivityStreamsDislike returns the value of this property. When
// IsActivityStreamsDislike returns false, GetActivityStreamsDislike will
// return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsDislike() vocab.ActivityStreamsDislike {
	return this.activitystreamsDislikeMember
}

// GetActivityStreamsDocument returns the value of this property. When
// IsActivityStreamsDocument returns false, GetActivityStreamsDocument will
// return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsDocument() vocab.ActivityStreamsDocument {
	return this.activitystreamsDocumentMember
}

// GetActivityStreamsEvent returns the value of this property. When
// IsActivityStreamsEvent returns false, GetActivityStreamsEvent will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsEvent() vocab.ActivityStreamsEvent {
	return this.activitystreamsEventMember
}

// GetActivityStreamsFlag returns the value of this property. When
// IsActivityStreamsFlag returns false, GetActivityStreamsFlag will return an
// arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsFlag() vocab.ActivityStreamsFlag {
	return this.activitystreamsFlagMember
}

// GetActivityStreamsFollow returns the value of this property. When
// IsActivityStreamsFollow returns false, GetActivityStreamsFollow will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsFollow() vocab.ActivityStreamsFollow {
	return this.activitystreamsFollowMember
}

// GetActivityStreamsGroup returns the value of this property. When
// IsActivityStreamsGroup returns false, GetActivityStreamsGroup will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsGroup() vocab.ActivityStreamsGroup {
	return this.activitystreamsGroupMember
}

// GetActivityStreamsIgnore returns the value of this property. When
// IsActivityStreamsIgnore returns false, GetActivityStreamsIgnore will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsIgnore() vocab.ActivityStreamsIgnore {
	return this.activitystreamsIgnoreMember
}

// GetActivityStreamsImage returns the value of this property. When
// IsActivityStreamsImage returns false, GetActivityStreamsImage will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsImage() vocab.ActivityStreamsImage {
	return this.activitystreamsImageMember
}

// GetActivityStreamsIntransitiveActivity returns the value of this property. When
// IsActivityStreamsIntransitiveActivity returns false,
// GetActivityStreamsIntransitiveActivity will return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsIntransitiveActivity() vocab.ActivityStreamsIntransitiveActivity {
	return this.activitystreamsIntransitiveActivityMember
}

// GetActivityStreamsInvite returns the value of this property. When
// IsActivityStreamsInvite returns false, GetActivityStreamsInvite will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsInvite() vocab.ActivityStreamsInvite {
	return this.activitystreamsInviteMember
}

// GetActivityStreamsJoin returns the value of this property. When
// IsActivityStreamsJoin returns false, GetActivityStreamsJoin will return an
// arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsJoin() vocab.ActivityStreamsJoin {
	return this.activitystreamsJoinMember
}

// GetActivityStreamsLeave returns the value of this property. When
// IsActivityStreamsLeave returns false, GetActivityStreamsLeave will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsLeave() vocab.ActivityStreamsLeave {
	return this.activitystreamsLeaveMember
}

// GetActivityStreamsLike returns the value of this property. When
// IsActivityStreamsLike returns false, GetActivityStreamsLike will return an
// arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsLike() vocab.ActivityStreamsLike {
	return this.activitystreamsLikeMember
}

// GetActivityStreamsLink returns the value of this property. When
// IsActivityStreamsLink returns false, GetActivityStreamsLink will return an
// arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsLink() vocab.ActivityStreamsLink {
	return this.activitystreamsLinkMember
}

// GetActivityStreamsListen returns the value of this property. When
// IsActivityStreamsListen returns false, GetActivityStreamsListen will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsListen() vocab.ActivityStreamsListen {
	return this.activitystreamsListenMember
}

// GetActivityStreamsMention returns the value of this property. When
// IsActivityStreamsMention returns false, GetActivityStreamsMention will
// return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsMention() vocab.ActivityStreamsMention {
	return this.activitystreamsMentionMember
}

// GetActivityStreamsMove returns the value of this property. When
// IsActivityStreamsMove returns false, GetActivityStreamsMove will return an
// arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsMove() vocab.ActivityStreamsMove {
	return this.activitystreamsMoveMember
}

// GetActivityStreamsNote returns the value of this property. When
// IsActivityStreamsNote returns false, GetActivityStreamsNote will return an
// arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsNote() vocab.ActivityStreamsNote {
	return this.activitystreamsNoteMember
}

// GetActivityStreamsObject returns the value of this property. When
// IsActivityStreamsObject returns false, GetActivityStreamsObject will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsObject() vocab.ActivityStreamsObject {
	return this.activitystreamsObjectMember
}

// GetActivityStreamsOffer returns the value of this property. When
// IsActivityStreamsOffer returns false, GetActivityStreamsOffer will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsOffer() vocab.ActivityStreamsOffer {
	return this.activitystreamsOfferMember
}

// GetActivityStreamsOrderedCollection returns the value of this property. When
// IsActivityStreamsOrderedCollection returns false,
// GetActivityStreamsOrderedCollection will return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsOrderedCollection() vocab.ActivityStreamsOrderedCollection {
	return this.activitystreamsOrderedCollectionMember
}

// GetActivityStreamsOrderedCollectionPage returns the value of this property.
// When IsActivityStreamsOrderedCollectionPage returns false,
// GetActivityStreamsOrderedCollectionPage will return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsOrderedCollectionPage() vocab.ActivityStreamsOrderedCollectionPage {
	return this.activitystreamsOrderedCollectionPageMember
}

// GetActivityStreamsOrganization returns the value of this property. When
// IsActivityStreamsOrganization returns false, GetActivityStreamsOrganization
// will return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsOrganization() vocab.ActivityStreamsOrganization {
	return this.activitystreamsOrganizationMember
}

// GetActivityStreamsPage returns the value of this property. When
// IsActivityStreamsPage returns false, GetActivityStreamsPage will return an
// arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsPage() vocab.ActivityStreamsPage {
	return this.activitystreamsPageMember
}

// GetActivityStreamsPerson returns the value of this property. When
// IsActivityStreamsPerson returns false, GetActivityStreamsPerson will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsPerson() vocab.ActivityStreamsPerson {
	return this.activitystreamsPersonMember
}

// GetActivityStreamsPlace returns the value of this property. When
// IsActivityStreamsPlace returns false, GetActivityStreamsPlace will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsPlace() vocab.ActivityStreamsPlace {
	return this.activitystreamsPlaceMember
}

// GetActivityStreamsProfile returns the value of this property. When
// IsActivityStreamsProfile returns false, GetActivityStreamsProfile will
// return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsProfile() vocab.ActivityStreamsProfile {
	return this.activitystreamsProfileMember
}

// GetActivityStreamsQuestion returns the value of this property. When
// IsActivityStreamsQuestion returns false, GetActivityStreamsQuestion will
// return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsQuestion() vocab.ActivityStreamsQuestion {
	return this.activitystreamsQuestionMember
}

// GetActivityStreamsRead returns the value of this property. When
// IsActivityStreamsRead returns false, GetActivityStreamsRead will return an
// arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsRead() vocab.ActivityStreamsRead {
	return this.activitystreamsReadMember
}

// GetActivityStreamsReject returns the value of this property. When
// IsActivityStreamsReject returns false, GetActivityStreamsReject will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsReject() vocab.ActivityStreamsReject {
	return this.activitystreamsRejectMember
}

// GetActivityStreamsRelationship returns the value of this property. When
// IsActivityStreamsRelationship returns false, GetActivityStreamsRelationship
// will return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsRelationship() vocab.ActivityStreamsRelationship {
	return this.activitystreamsRelationshipMember
}

// GetActivityStreamsRemove returns the value of this property. When
// IsActivityStreamsRemove returns false, GetActivityStreamsRemove will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsRemove() vocab.ActivityStreamsRemove {
	return this.activitystreamsRemoveMember
}

// GetActivityStreamsService returns the value of this property. When
// IsActivityStreamsService returns false, GetActivityStreamsService will
// return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsService() vocab.ActivityStreamsService {
	return this.activitystreamsServiceMember
}

// GetActivityStreamsTentativeAccept returns the value of this property. When
// IsActivityStreamsTentativeAccept returns false,
// GetActivityStreamsTentativeAccept will return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsTentativeAccept() vocab.ActivityStreamsTentativeAccept {
	return this.activitystreamsTentativeAcceptMember
}

// GetActivityStreamsTentativeReject returns the value of this property. When
// IsActivityStreamsTentativeReject returns false,
// GetActivityStreamsTentativeReject will return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsTentativeReject() vocab.ActivityStreamsTentativeReject {
	return this.activitystreamsTentativeRejectMember
}

// GetActivityStreamsTombstone returns the value of this property. When
// IsActivityStreamsTombstone returns false, GetActivityStreamsTombstone will
// return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsTombstone() vocab.ActivityStreamsTombstone {
	return this.activitystreamsTombstoneMember
}

// GetActivityStreamsTravel returns the value of this property. When
// IsActivityStreamsTravel returns false, GetActivityStreamsTravel will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsTravel() vocab.ActivityStreamsTravel {
	return this.activitystreamsTravelMember
}

// GetActivityStreamsUndo returns the value of this property. When
// IsActivityStreamsUndo returns false, GetActivityStreamsUndo will return an
// arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsUndo() vocab.ActivityStreamsUndo {
	return this.activitystreamsUndoMember
}

// GetActivityStreamsUpdate returns the value of this property. When
// IsActivityStreamsUpdate returns false, GetActivityStreamsUpdate will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsUpdate() vocab.ActivityStreamsUpdate {
	return this.activitystreamsUpdateMember
}

// GetActivityStreamsVideo returns the value of this property. When
// IsActivityStreamsVideo returns false, GetActivityStreamsVideo will return
// an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsVideo() vocab.ActivityStreamsVideo {
	return this.activitystreamsVideoMember
}

// GetActivityStreamsView returns the value of this property. When
// IsActivityStreamsView returns false, GetActivityStreamsView will return an
// arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetActivityStreamsView() vocab.ActivityStreamsView {
	return this.activitystreamsViewMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this ActivityStreamsGeneratorPropertyIterator) GetIRI() *url.URL {
	return this.iri
}

// GetType returns the value in this property as a Type. Returns nil if the value
// is not an ActivityStreams type, such as an IRI or another value.
func (this ActivityStreamsGeneratorPropertyIterator) GetType() vocab.Type {
	if this.IsActivityStreamsObject() {
		return this.GetActivityStreamsObject()
	}
	if this.IsActivityStreamsLink() {
		return this.GetActivityStreamsLink()
	}
	if this.IsActivityStreamsAccept() {
		return this.GetActivityStreamsAccept()
	}
	if this.IsActivityStreamsActivity() {
		return this.GetActivityStreamsActivity()
	}
	if this.IsActivityStreamsAdd() {
		return this.GetActivityStreamsAdd()
	}
	if this.IsActivityStreamsAnnounce() {
		return this.GetActivityStreamsAnnounce()
	}
	if this.IsActivityStreamsApplication() {
		return this.GetActivityStreamsApplication()
	}
	if this.IsActivityStreamsArrive() {
		return this.GetActivityStreamsArrive()
	}
	if this.IsActivityStreamsArticle() {
		return this.GetActivityStreamsArticle()
	}
	if this.IsActivityStreamsAudio() {
		return this.GetActivityStreamsAudio()
	}
	if this.IsActivityStreamsBlock() {
		return this.GetActivityStreamsBlock()
	}
	if this.IsActivityStreamsCollection() {
		return this.GetActivityStreamsCollection()
	}
	if this.IsActivityStreamsCollectionPage() {
		return this.GetActivityStreamsCollectionPage()
	}
	if this.IsActivityStreamsCreate() {
		return this.GetActivityStreamsCreate()
	}
	if this.IsActivityStreamsDelete() {
		return this.GetActivityStreamsDelete()
	}
	if this.IsActivityStreamsDislike() {
		return this.GetActivityStreamsDislike()
	}
	if this.IsActivityStreamsDocument() {
		return this.GetActivityStreamsDocument()
	}
	if this.IsActivityStreamsEvent() {
		return this.GetActivityStreamsEvent()
	}
	if this.IsActivityStreamsFlag() {
		return this.GetActivityStreamsFlag()
	}
	if this.IsActivityStreamsFollow() {
		return this.GetActivityStreamsFollow()
	}
	if this.IsActivityStreamsGroup() {
		return this.GetActivityStreamsGroup()
	}
	if this.IsActivityStreamsIgnore() {
		return this.GetActivityStreamsIgnore()
	}
	if this.IsActivityStreamsImage() {
		return this.GetActivityStreamsImage()
	}
	if this.IsActivityStreamsIntransitiveActivity() {
		return this.GetActivityStreamsIntransitiveActivity()
	}
	if this.IsActivityStreamsInvite() {
		return this.GetActivityStreamsInvite()
	}
	if this.IsActivityStreamsJoin() {
		return this.GetActivityStreamsJoin()
	}
	if this.IsActivityStreamsLeave() {
		return this.GetActivityStreamsLeave()
	}
	if this.IsActivityStreamsLike() {
		return this.GetActivityStreamsLike()
	}
	if this.IsActivityStreamsListen() {
		return this.GetActivityStreamsListen()
	}
	if this.IsActivityStreamsMention() {
		return this.GetActivityStreamsMention()
	}
	if this.IsActivityStreamsMove() {
		return this.GetActivityStreamsMove()
	}
	if this.IsActivityStreamsNote() {
		return this.GetActivityStreamsNote()
	}
	if this.IsActivityStreamsOffer() {
		return this.GetActivityStreamsOffer()
	}
	if this.IsActivityStreamsOrderedCollection() {
		return this.GetActivityStreamsOrderedCollection()
	}
	if this.IsActivityStreamsOrderedCollectionPage() {
		return this.GetActivityStreamsOrderedCollectionPage()
	}
	if this.IsActivityStreamsOrganization() {
		return this.GetActivityStreamsOrganization()
	}
	if this.IsActivityStreamsPage() {
		return this.GetActivityStreamsPage()
	}
	if this.IsActivityStreamsPerson() {
		return this.GetActivityStreamsPerson()
	}
	if this.IsActivityStreamsPlace() {
		return this.GetActivityStreamsPlace()
	}
	if this.IsActivityStreamsProfile() {
		return this.GetActivityStreamsProfile()
	}
	if this.IsActivityStreamsQuestion() {
		return this.GetActivityStreamsQuestion()
	}
	if this.IsActivityStreamsRead() {
		return this.GetActivityStreamsRead()
	}
	if this.IsActivityStreamsReject() {
		return this.GetActivityStreamsReject()
	}
	if this.IsActivityStreamsRelationship() {
		return this.GetActivityStreamsRelationship()
	}
	if this.IsActivityStreamsRemove() {
		return this.GetActivityStreamsRemove()
	}
	if this.IsActivityStreamsService() {
		return this.GetActivityStreamsService()
	}
	if this.IsActivityStreamsTentativeAccept() {
		return this.GetActivityStreamsTentativeAccept()
	}
	if this.IsActivityStreamsTentativeReject() {
		return this.GetActivityStreamsTentativeReject()
	}
	if this.IsActivityStreamsTombstone() {
		return this.GetActivityStreamsTombstone()
	}
	if this.IsActivityStreamsTravel() {
		return this.GetActivityStreamsTravel()
	}
	if this.IsActivityStreamsUndo() {
		return this.GetActivityStreamsUndo()
	}
	if this.IsActivityStreamsUpdate() {
		return this.GetActivityStreamsUpdate()
	}
	if this.IsActivityStreamsVideo() {
		return this.GetActivityStreamsVideo()
	}
	if this.IsActivityStreamsView() {
		return this.GetActivityStreamsView()
	}

	return nil
}

// HasAny returns true if any of the different values is set.
func (this ActivityStreamsGeneratorPropertyIterator) HasAny() bool {
	return this.IsActivityStreamsObject() ||
		this.IsActivityStreamsLink() ||
		this.IsActivityStreamsAccept() ||
		this.IsActivityStreamsActivity() ||
		this.IsActivityStreamsAdd() ||
		this.IsActivityStreamsAnnounce() ||
		this.IsActivityStreamsApplication() ||
		this.IsActivityStreamsArrive() ||
		this.IsActivityStreamsArticle() ||
		this.IsActivityStreamsAudio() ||
		this.IsActivityStreamsBlock() ||
		this.IsActivityStreamsCollection() ||
		this.IsActivityStreamsCollectionPage() ||
		this.IsActivityStreamsCreate() ||
		this.IsActivityStreamsDelete() ||
		this.IsActivityStreamsDislike() ||
		this.IsActivityStreamsDocument() ||
		this.IsActivityStreamsEvent() ||
		this.IsActivityStreamsFlag() ||
		this.IsActivityStreamsFollow() ||
		this.IsActivityStreamsGroup() ||
		this.IsActivityStreamsIgnore() ||
		this.IsActivityStreamsImage() ||
		this.IsActivityStreamsIntransitiveActivity() ||
		this.IsActivityStreamsInvite() ||
		this.IsActivityStreamsJoin() ||
		this.IsActivityStreamsLeave() ||
		this.IsActivityStreamsLike() ||
		this.IsActivityStreamsListen() ||
		this.IsActivityStreamsMention() ||
		this.IsActivityStreamsMove() ||
		this.IsActivityStreamsNote() ||
		this.IsActivityStreamsOffer() ||
		this.IsActivityStreamsOrderedCollection() ||
		this.IsActivityStreamsOrderedCollectionPage() ||
		this.IsActivityStreamsOrganization() ||
		this.IsActivityStreamsPage() ||
		this.IsActivityStreamsPerson() ||
		this.IsActivityStreamsPlace() ||
		this.IsActivityStreamsProfile() ||
		this.IsActivityStreamsQuestion() ||
		this.IsActivityStreamsRead() ||
		this.IsActivityStreamsReject() ||
		this.IsActivityStreamsRelationship() ||
		this.IsActivityStreamsRemove() ||
		this.IsActivityStreamsService() ||
		this.IsActivityStreamsTentativeAccept() ||
		this.IsActivityStreamsTentativeReject() ||
		this.IsActivityStreamsTombstone() ||
		this.IsActivityStreamsTravel() ||
		this.IsActivityStreamsUndo() ||
		this.IsActivityStreamsUpdate() ||
		this.IsActivityStreamsVideo() ||
		this.IsActivityStreamsView() ||
		this.iri != nil
}

// IsActivityStreamsAccept returns true if this property has a type of "Accept".
// When true, use the GetActivityStreamsAccept and SetActivityStreamsAccept
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsAccept() bool {
	return this.activitystreamsAcceptMember != nil
}

// IsActivityStreamsActivity returns true if this property has a type of
// "Activity". When true, use the GetActivityStreamsActivity and
// SetActivityStreamsActivity methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsActivity() bool {
	return this.activitystreamsActivityMember != nil
}

// IsActivityStreamsAdd returns true if this property has a type of "Add". When
// true, use the GetActivityStreamsAdd and SetActivityStreamsAdd methods to
// access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsAdd() bool {
	return this.activitystreamsAddMember != nil
}

// IsActivityStreamsAnnounce returns true if this property has a type of
// "Announce". When true, use the GetActivityStreamsAnnounce and
// SetActivityStreamsAnnounce methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsAnnounce() bool {
	return this.activitystreamsAnnounceMember != nil
}

// IsActivityStreamsApplication returns true if this property has a type of
// "Application". When true, use the GetActivityStreamsApplication and
// SetActivityStreamsApplication methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsApplication() bool {
	return this.activitystreamsApplicationMember != nil
}

// IsActivityStreamsArrive returns true if this property has a type of "Arrive".
// When true, use the GetActivityStreamsArrive and SetActivityStreamsArrive
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsArrive() bool {
	return this.activitystreamsArriveMember != nil
}

// IsActivityStreamsArticle returns true if this property has a type of "Article".
// When true, use the GetActivityStreamsArticle and SetActivityStreamsArticle
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsArticle() bool {
	return this.activitystreamsArticleMember != nil
}

// IsActivityStreamsAudio returns true if this property has a type of "Audio".
// When true, use the GetActivityStreamsAudio and SetActivityStreamsAudio
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsAudio() bool {
	return this.activitystreamsAudioMember != nil
}

// IsActivityStreamsBlock returns true if this property has a type of "Block".
// When true, use the GetActivityStreamsBlock and SetActivityStreamsBlock
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsBlock() bool {
	return this.activitystreamsBlockMember != nil
}

// IsActivityStreamsCollection returns true if this property has a type of
// "Collection". When true, use the GetActivityStreamsCollection and
// SetActivityStreamsCollection methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsCollection() bool {
	return this.activitystreamsCollectionMember != nil
}

// IsActivityStreamsCollectionPage returns true if this property has a type of
// "CollectionPage". When true, use the GetActivityStreamsCollectionPage and
// SetActivityStreamsCollectionPage methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsCollectionPage() bool {
	return this.activitystreamsCollectionPageMember != nil
}

// IsActivityStreamsCreate returns true if this property has a type of "Create".
// When true, use the GetActivityStreamsCreate and SetActivityStreamsCreate
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsCreate() bool {
	return this.activitystreamsCreateMember != nil
}

// IsActivityStreamsDelete returns true if this property has a type of "Delete".
// When true, use the GetActivityStreamsDelete and SetActivityStreamsDelete
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsDelete() bool {
	return this.activitystreamsDeleteMember != nil
}

// IsActivityStreamsDislike returns true if this property has a type of "Dislike".
// When true, use the GetActivityStreamsDislike and SetActivityStreamsDislike
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsDislike() bool {
	return this.activitystreamsDislikeMember != nil
}

// IsActivityStreamsDocument returns true if this property has a type of
// "Document". When true, use the GetActivityStreamsDocument and
// SetActivityStreamsDocument methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsDocument() bool {
	return this.activitystreamsDocumentMember != nil
}

// IsActivityStreamsEvent returns true if this property has a type of "Event".
// When true, use the GetActivityStreamsEvent and SetActivityStreamsEvent
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsEvent() bool {
	return this.activitystreamsEventMember != nil
}

// IsActivityStreamsFlag returns true if this property has a type of "Flag". When
// true, use the GetActivityStreamsFlag and SetActivityStreamsFlag methods to
// access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsFlag() bool {
	return this.activitystreamsFlagMember != nil
}

// IsActivityStreamsFollow returns true if this property has a type of "Follow".
// When true, use the GetActivityStreamsFollow and SetActivityStreamsFollow
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsFollow() bool {
	return this.activitystreamsFollowMember != nil
}

// IsActivityStreamsGroup returns true if this property has a type of "Group".
// When true, use the GetActivityStreamsGroup and SetActivityStreamsGroup
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsGroup() bool {
	return this.activitystreamsGroupMember != nil
}

// IsActivityStreamsIgnore returns true if this property has a type of "Ignore".
// When true, use the GetActivityStreamsIgnore and SetActivityStreamsIgnore
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsIgnore() bool {
	return this.activitystreamsIgnoreMember != nil
}

// IsActivityStreamsImage returns true if this property has a type of "Image".
// When true, use the GetActivityStreamsImage and SetActivityStreamsImage
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsImage() bool {
	return this.activitystreamsImageMember != nil
}

// IsActivityStreamsIntransitiveActivity returns true if this property has a type
// of "IntransitiveActivity". When true, use the
// GetActivityStreamsIntransitiveActivity and
// SetActivityStreamsIntransitiveActivity methods to access and set this
// property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsIntransitiveActivity() bool {
	return this.activitystreamsIntransitiveActivityMember != nil
}

// IsActivityStreamsInvite returns true if this property has a type of "Invite".
// When true, use the GetActivityStreamsInvite and SetActivityStreamsInvite
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsInvite() bool {
	return this.activitystreamsInviteMember != nil
}

// IsActivityStreamsJoin returns true if this property has a type of "Join". When
// true, use the GetActivityStreamsJoin and SetActivityStreamsJoin methods to
// access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsJoin() bool {
	return this.activitystreamsJoinMember != nil
}

// IsActivityStreamsLeave returns true if this property has a type of "Leave".
// When true, use the GetActivityStreamsLeave and SetActivityStreamsLeave
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsLeave() bool {
	return this.activitystreamsLeaveMember != nil
}

// IsActivityStreamsLike returns true if this property has a type of "Like". When
// true, use the GetActivityStreamsLike and SetActivityStreamsLike methods to
// access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsLike() bool {
	return this.activitystreamsLikeMember != nil
}

// IsActivityStreamsLink returns true if this property has a type of "Link". When
// true, use the GetActivityStreamsLink and SetActivityStreamsLink methods to
// access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsLink() bool {
	return this.activitystreamsLinkMember != nil
}

// IsActivityStreamsListen returns true if this property has a type of "Listen".
// When true, use the GetActivityStreamsListen and SetActivityStreamsListen
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsListen() bool {
	return this.activitystreamsListenMember != nil
}

// IsActivityStreamsMention returns true if this property has a type of "Mention".
// When true, use the GetActivityStreamsMention and SetActivityStreamsMention
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsMention() bool {
	return this.activitystreamsMentionMember != nil
}

// IsActivityStreamsMove returns true if this property has a type of "Move". When
// true, use the GetActivityStreamsMove and SetActivityStreamsMove methods to
// access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsMove() bool {
	return this.activitystreamsMoveMember != nil
}

// IsActivityStreamsNote returns true if this property has a type of "Note". When
// true, use the GetActivityStreamsNote and SetActivityStreamsNote methods to
// access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsNote() bool {
	return this.activitystreamsNoteMember != nil
}

// IsActivityStreamsObject returns true if this property has a type of "Object".
// When true, use the GetActivityStreamsObject and SetActivityStreamsObject
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsObject() bool {
	return this.activitystreamsObjectMember != nil
}

// IsActivityStreamsOffer returns true if this property has a type of "Offer".
// When true, use the GetActivityStreamsOffer and SetActivityStreamsOffer
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsOffer() bool {
	return this.activitystreamsOfferMember != nil
}

// IsActivityStreamsOrderedCollection returns true if this property has a type of
// "OrderedCollection". When true, use the GetActivityStreamsOrderedCollection
// and SetActivityStreamsOrderedCollection methods to access and set this
// property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsOrderedCollection() bool {
	return this.activitystreamsOrderedCollectionMember != nil
}

// IsActivityStreamsOrderedCollectionPage returns true if this property has a type
// of "OrderedCollectionPage". When true, use the
// GetActivityStreamsOrderedCollectionPage and
// SetActivityStreamsOrderedCollectionPage methods to access and set this
// property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsOrderedCollectionPage() bool {
	return this.activitystreamsOrderedCollectionPageMember != nil
}

// IsActivityStreamsOrganization returns true if this property has a type of
// "Organization". When true, use the GetActivityStreamsOrganization and
// SetActivityStreamsOrganization methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsOrganization() bool {
	return this.activitystreamsOrganizationMember != nil
}

// IsActivityStreamsPage returns true if this property has a type of "Page". When
// true, use the GetActivityStreamsPage and SetActivityStreamsPage methods to
// access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsPage() bool {
	return this.activitystreamsPageMember != nil
}

// IsActivityStreamsPerson returns true if this property has a type of "Person".
// When true, use the GetActivityStreamsPerson and SetActivityStreamsPerson
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsPerson() bool {
	return this.activitystreamsPersonMember != nil
}

// IsActivityStreamsPlace returns true if this property has a type of "Place".
// When true, use the GetActivityStreamsPlace and SetActivityStreamsPlace
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsPlace() bool {
	return this.activitystreamsPlaceMember != nil
}

// IsActivityStreamsProfile returns true if this property has a type of "Profile".
// When true, use the GetActivityStreamsProfile and SetActivityStreamsProfile
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsProfile() bool {
	return this.activitystreamsProfileMember != nil
}

// IsActivityStreamsQuestion returns true if this property has a type of
// "Question". When true, use the GetActivityStreamsQuestion and
// SetActivityStreamsQuestion methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsQuestion() bool {
	return this.activitystreamsQuestionMember != nil
}

// IsActivityStreamsRead returns true if this property has a type of "Read". When
// true, use the GetActivityStreamsRead and SetActivityStreamsRead methods to
// access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsRead() bool {
	return this.activitystreamsReadMember != nil
}

// IsActivityStreamsReject returns true if this property has a type of "Reject".
// When true, use the GetActivityStreamsReject and SetActivityStreamsReject
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsReject() bool {
	return this.activitystreamsRejectMember != nil
}

// IsActivityStreamsRelationship returns true if this property has a type of
// "Relationship". When true, use the GetActivityStreamsRelationship and
// SetActivityStreamsRelationship methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsRelationship() bool {
	return this.activitystreamsRelationshipMember != nil
}

// IsActivityStreamsRemove returns true if this property has a type of "Remove".
// When true, use the GetActivityStreamsRemove and SetActivityStreamsRemove
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsRemove() bool {
	return this.activitystreamsRemoveMember != nil
}

// IsActivityStreamsService returns true if this property has a type of "Service".
// When true, use the GetActivityStreamsService and SetActivityStreamsService
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsService() bool {
	return this.activitystreamsServiceMember != nil
}

// IsActivityStreamsTentativeAccept returns true if this property has a type of
// "TentativeAccept". When true, use the GetActivityStreamsTentativeAccept and
// SetActivityStreamsTentativeAccept methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsTentativeAccept() bool {
	return this.activitystreamsTentativeAcceptMember != nil
}

// IsActivityStreamsTentativeReject returns true if this property has a type of
// "TentativeReject". When true, use the GetActivityStreamsTentativeReject and
// SetActivityStreamsTentativeReject methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsTentativeReject() bool {
	return this.activitystreamsTentativeRejectMember != nil
}

// IsActivityStreamsTombstone returns true if this property has a type of
// "Tombstone". When true, use the GetActivityStreamsTombstone and
// SetActivityStreamsTombstone methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsTombstone() bool {
	return this.activitystreamsTombstoneMember != nil
}

// IsActivityStreamsTravel returns true if this property has a type of "Travel".
// When true, use the GetActivityStreamsTravel and SetActivityStreamsTravel
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsTravel() bool {
	return this.activitystreamsTravelMember != nil
}

// IsActivityStreamsUndo returns true if this property has a type of "Undo". When
// true, use the GetActivityStreamsUndo and SetActivityStreamsUndo methods to
// access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsUndo() bool {
	return this.activitystreamsUndoMember != nil
}

// IsActivityStreamsUpdate returns true if this property has a type of "Update".
// When true, use the GetActivityStreamsUpdate and SetActivityStreamsUpdate
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsUpdate() bool {
	return this.activitystreamsUpdateMember != nil
}

// IsActivityStreamsVideo returns true if this property has a type of "Video".
// When true, use the GetActivityStreamsVideo and SetActivityStreamsVideo
// methods to access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsVideo() bool {
	return this.activitystreamsVideoMember != nil
}

// IsActivityStreamsView returns true if this property has a type of "View". When
// true, use the GetActivityStreamsView and SetActivityStreamsView methods to
// access and set this property.
func (this ActivityStreamsGeneratorPropertyIterator) IsActivityStreamsView() bool {
	return this.activitystreamsViewMember != nil
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this ActivityStreamsGeneratorPropertyIterator) IsIRI() bool {
	return this.iri != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsGeneratorPropertyIterator) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	var child map[string]string
	if this.IsActivityStreamsObject() {
		child = this.GetActivityStreamsObject().JSONLDContext()
	} else if this.IsActivityStreamsLink() {
		child = this.GetActivityStreamsLink().JSONLDContext()
	} else if this.IsActivityStreamsAccept() {
		child = this.GetActivityStreamsAccept().JSONLDContext()
	} else if this.IsActivityStreamsActivity() {
		child = this.GetActivityStreamsActivity().JSONLDContext()
	} else if this.IsActivityStreamsAdd() {
		child = this.GetActivityStreamsAdd().JSONLDContext()
	} else if this.IsActivityStreamsAnnounce() {
		child = this.GetActivityStreamsAnnounce().JSONLDContext()
	} else if this.IsActivityStreamsApplication() {
		child = this.GetActivityStreamsApplication().JSONLDContext()
	} else if this.IsActivityStreamsArrive() {
		child = this.GetActivityStreamsArrive().JSONLDContext()
	} else if this.IsActivityStreamsArticle() {
		child = this.GetActivityStreamsArticle().JSONLDContext()
	} else if this.IsActivityStreamsAudio() {
		child = this.GetActivityStreamsAudio().JSONLDContext()
	} else if this.IsActivityStreamsBlock() {
		child = this.GetActivityStreamsBlock().JSONLDContext()
	} else if this.IsActivityStreamsCollection() {
		child = this.GetActivityStreamsCollection().JSONLDContext()
	} else if this.IsActivityStreamsCollectionPage() {
		child = this.GetActivityStreamsCollectionPage().JSONLDContext()
	} else if this.IsActivityStreamsCreate() {
		child = this.GetActivityStreamsCreate().JSONLDContext()
	} else if this.IsActivityStreamsDelete() {
		child = this.GetActivityStreamsDelete().JSONLDContext()
	} else if this.IsActivityStreamsDislike() {
		child = this.GetActivityStreamsDislike().JSONLDContext()
	} else if this.IsActivityStreamsDocument() {
		child = this.GetActivityStreamsDocument().JSONLDContext()
	} else if this.IsActivityStreamsEvent() {
		child = this.GetActivityStreamsEvent().JSONLDContext()
	} else if this.IsActivityStreamsFlag() {
		child = this.GetActivityStreamsFlag().JSONLDContext()
	} else if this.IsActivityStreamsFollow() {
		child = this.GetActivityStreamsFollow().JSONLDContext()
	} else if this.IsActivityStreamsGroup() {
		child = this.GetActivityStreamsGroup().JSONLDContext()
	} else if this.IsActivityStreamsIgnore() {
		child = this.GetActivityStreamsIgnore().JSONLDContext()
	} else if this.IsActivityStreamsImage() {
		child = this.GetActivityStreamsImage().JSONLDContext()
	} else if this.IsActivityStreamsIntransitiveActivity() {
		child = this.GetActivityStreamsIntransitiveActivity().JSONLDContext()
	} else if this.IsActivityStreamsInvite() {
		child = this.GetActivityStreamsInvite().JSONLDContext()
	} else if this.IsActivityStreamsJoin() {
		child = this.GetActivityStreamsJoin().JSONLDContext()
	} else if this.IsActivityStreamsLeave() {
		child = this.GetActivityStreamsLeave().JSONLDContext()
	} else if this.IsActivityStreamsLike() {
		child = this.GetActivityStreamsLike().JSONLDContext()
	} else if this.IsActivityStreamsListen() {
		child = this.GetActivityStreamsListen().JSONLDContext()
	} else if this.IsActivityStreamsMention() {
		child = this.GetActivityStreamsMention().JSONLDContext()
	} else if this.IsActivityStreamsMove() {
		child = this.GetActivityStreamsMove().JSONLDContext()
	} else if this.IsActivityStreamsNote() {
		child = this.GetActivityStreamsNote().JSONLDContext()
	} else if this.IsActivityStreamsOffer() {
		child = this.GetActivityStreamsOffer().JSONLDContext()
	} else if this.IsActivityStreamsOrderedCollection() {
		child = this.GetActivityStreamsOrderedCollection().JSONLDContext()
	} else if this.IsActivityStreamsOrderedCollectionPage() {
		child = this.GetActivityStreamsOrderedCollectionPage().JSONLDContext()
	} else if this.IsActivityStreamsOrganization() {
		child = this.GetActivityStreamsOrganization().JSONLDContext()
	} else if this.IsActivityStreamsPage() {
		child = this.GetActivityStreamsPage().JSONLDContext()
	} else if this.IsActivityStreamsPerson() {
		child = this.GetActivityStreamsPerson().JSONLDContext()
	} else if this.IsActivityStreamsPlace() {
		child = this.GetActivityStreamsPlace().JSONLDContext()
	} else if this.IsActivityStreamsProfile() {
		child = this.GetActivityStreamsProfile().JSONLDContext()
	} else if this.IsActivityStreamsQuestion() {
		child = this.GetActivityStreamsQuestion().JSONLDContext()
	} else if this.IsActivityStreamsRead() {
		child = this.GetActivityStreamsRead().JSONLDContext()
	} else if this.IsActivityStreamsReject() {
		child = this.GetActivityStreamsReject().JSONLDContext()
	} else if this.IsActivityStreamsRelationship() {
		child = this.GetActivityStreamsRelationship().JSONLDContext()
	} else if this.IsActivityStreamsRemove() {
		child = this.GetActivityStreamsRemove().JSONLDContext()
	} else if this.IsActivityStreamsService() {
		child = this.GetActivityStreamsService().JSONLDContext()
	} else if this.IsActivityStreamsTentativeAccept() {
		child = this.GetActivityStreamsTentativeAccept().JSONLDContext()
	} else if this.IsActivityStreamsTentativeReject() {
		child = this.GetActivityStreamsTentativeReject().JSONLDContext()
	} else if this.IsActivityStreamsTombstone() {
		child = this.GetActivityStreamsTombstone().JSONLDContext()
	} else if this.IsActivityStreamsTravel() {
		child = this.GetActivityStreamsTravel().JSONLDContext()
	} else if this.IsActivityStreamsUndo() {
		child = this.GetActivityStreamsUndo().JSONLDContext()
	} else if this.IsActivityStreamsUpdate() {
		child = this.GetActivityStreamsUpdate().JSONLDContext()
	} else if this.IsActivityStreamsVideo() {
		child = this.GetActivityStreamsVideo().JSONLDContext()
	} else if this.IsActivityStreamsView() {
		child = this.GetActivityStreamsView().JSONLDContext()
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
func (this ActivityStreamsGeneratorPropertyIterator) KindIndex() int {
	if this.IsActivityStreamsObject() {
		return 0
	}
	if this.IsActivityStreamsLink() {
		return 1
	}
	if this.IsActivityStreamsAccept() {
		return 2
	}
	if this.IsActivityStreamsActivity() {
		return 3
	}
	if this.IsActivityStreamsAdd() {
		return 4
	}
	if this.IsActivityStreamsAnnounce() {
		return 5
	}
	if this.IsActivityStreamsApplication() {
		return 6
	}
	if this.IsActivityStreamsArrive() {
		return 7
	}
	if this.IsActivityStreamsArticle() {
		return 8
	}
	if this.IsActivityStreamsAudio() {
		return 9
	}
	if this.IsActivityStreamsBlock() {
		return 10
	}
	if this.IsActivityStreamsCollection() {
		return 11
	}
	if this.IsActivityStreamsCollectionPage() {
		return 12
	}
	if this.IsActivityStreamsCreate() {
		return 13
	}
	if this.IsActivityStreamsDelete() {
		return 14
	}
	if this.IsActivityStreamsDislike() {
		return 15
	}
	if this.IsActivityStreamsDocument() {
		return 16
	}
	if this.IsActivityStreamsEvent() {
		return 17
	}
	if this.IsActivityStreamsFlag() {
		return 18
	}
	if this.IsActivityStreamsFollow() {
		return 19
	}
	if this.IsActivityStreamsGroup() {
		return 20
	}
	if this.IsActivityStreamsIgnore() {
		return 21
	}
	if this.IsActivityStreamsImage() {
		return 22
	}
	if this.IsActivityStreamsIntransitiveActivity() {
		return 23
	}
	if this.IsActivityStreamsInvite() {
		return 24
	}
	if this.IsActivityStreamsJoin() {
		return 25
	}
	if this.IsActivityStreamsLeave() {
		return 26
	}
	if this.IsActivityStreamsLike() {
		return 27
	}
	if this.IsActivityStreamsListen() {
		return 28
	}
	if this.IsActivityStreamsMention() {
		return 29
	}
	if this.IsActivityStreamsMove() {
		return 30
	}
	if this.IsActivityStreamsNote() {
		return 31
	}
	if this.IsActivityStreamsOffer() {
		return 32
	}
	if this.IsActivityStreamsOrderedCollection() {
		return 33
	}
	if this.IsActivityStreamsOrderedCollectionPage() {
		return 34
	}
	if this.IsActivityStreamsOrganization() {
		return 35
	}
	if this.IsActivityStreamsPage() {
		return 36
	}
	if this.IsActivityStreamsPerson() {
		return 37
	}
	if this.IsActivityStreamsPlace() {
		return 38
	}
	if this.IsActivityStreamsProfile() {
		return 39
	}
	if this.IsActivityStreamsQuestion() {
		return 40
	}
	if this.IsActivityStreamsRead() {
		return 41
	}
	if this.IsActivityStreamsReject() {
		return 42
	}
	if this.IsActivityStreamsRelationship() {
		return 43
	}
	if this.IsActivityStreamsRemove() {
		return 44
	}
	if this.IsActivityStreamsService() {
		return 45
	}
	if this.IsActivityStreamsTentativeAccept() {
		return 46
	}
	if this.IsActivityStreamsTentativeReject() {
		return 47
	}
	if this.IsActivityStreamsTombstone() {
		return 48
	}
	if this.IsActivityStreamsTravel() {
		return 49
	}
	if this.IsActivityStreamsUndo() {
		return 50
	}
	if this.IsActivityStreamsUpdate() {
		return 51
	}
	if this.IsActivityStreamsVideo() {
		return 52
	}
	if this.IsActivityStreamsView() {
		return 53
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
func (this ActivityStreamsGeneratorPropertyIterator) LessThan(o vocab.ActivityStreamsGeneratorPropertyIterator) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsActivityStreamsObject() {
		return this.GetActivityStreamsObject().LessThan(o.GetActivityStreamsObject())
	} else if this.IsActivityStreamsLink() {
		return this.GetActivityStreamsLink().LessThan(o.GetActivityStreamsLink())
	} else if this.IsActivityStreamsAccept() {
		return this.GetActivityStreamsAccept().LessThan(o.GetActivityStreamsAccept())
	} else if this.IsActivityStreamsActivity() {
		return this.GetActivityStreamsActivity().LessThan(o.GetActivityStreamsActivity())
	} else if this.IsActivityStreamsAdd() {
		return this.GetActivityStreamsAdd().LessThan(o.GetActivityStreamsAdd())
	} else if this.IsActivityStreamsAnnounce() {
		return this.GetActivityStreamsAnnounce().LessThan(o.GetActivityStreamsAnnounce())
	} else if this.IsActivityStreamsApplication() {
		return this.GetActivityStreamsApplication().LessThan(o.GetActivityStreamsApplication())
	} else if this.IsActivityStreamsArrive() {
		return this.GetActivityStreamsArrive().LessThan(o.GetActivityStreamsArrive())
	} else if this.IsActivityStreamsArticle() {
		return this.GetActivityStreamsArticle().LessThan(o.GetActivityStreamsArticle())
	} else if this.IsActivityStreamsAudio() {
		return this.GetActivityStreamsAudio().LessThan(o.GetActivityStreamsAudio())
	} else if this.IsActivityStreamsBlock() {
		return this.GetActivityStreamsBlock().LessThan(o.GetActivityStreamsBlock())
	} else if this.IsActivityStreamsCollection() {
		return this.GetActivityStreamsCollection().LessThan(o.GetActivityStreamsCollection())
	} else if this.IsActivityStreamsCollectionPage() {
		return this.GetActivityStreamsCollectionPage().LessThan(o.GetActivityStreamsCollectionPage())
	} else if this.IsActivityStreamsCreate() {
		return this.GetActivityStreamsCreate().LessThan(o.GetActivityStreamsCreate())
	} else if this.IsActivityStreamsDelete() {
		return this.GetActivityStreamsDelete().LessThan(o.GetActivityStreamsDelete())
	} else if this.IsActivityStreamsDislike() {
		return this.GetActivityStreamsDislike().LessThan(o.GetActivityStreamsDislike())
	} else if this.IsActivityStreamsDocument() {
		return this.GetActivityStreamsDocument().LessThan(o.GetActivityStreamsDocument())
	} else if this.IsActivityStreamsEvent() {
		return this.GetActivityStreamsEvent().LessThan(o.GetActivityStreamsEvent())
	} else if this.IsActivityStreamsFlag() {
		return this.GetActivityStreamsFlag().LessThan(o.GetActivityStreamsFlag())
	} else if this.IsActivityStreamsFollow() {
		return this.GetActivityStreamsFollow().LessThan(o.GetActivityStreamsFollow())
	} else if this.IsActivityStreamsGroup() {
		return this.GetActivityStreamsGroup().LessThan(o.GetActivityStreamsGroup())
	} else if this.IsActivityStreamsIgnore() {
		return this.GetActivityStreamsIgnore().LessThan(o.GetActivityStreamsIgnore())
	} else if this.IsActivityStreamsImage() {
		return this.GetActivityStreamsImage().LessThan(o.GetActivityStreamsImage())
	} else if this.IsActivityStreamsIntransitiveActivity() {
		return this.GetActivityStreamsIntransitiveActivity().LessThan(o.GetActivityStreamsIntransitiveActivity())
	} else if this.IsActivityStreamsInvite() {
		return this.GetActivityStreamsInvite().LessThan(o.GetActivityStreamsInvite())
	} else if this.IsActivityStreamsJoin() {
		return this.GetActivityStreamsJoin().LessThan(o.GetActivityStreamsJoin())
	} else if this.IsActivityStreamsLeave() {
		return this.GetActivityStreamsLeave().LessThan(o.GetActivityStreamsLeave())
	} else if this.IsActivityStreamsLike() {
		return this.GetActivityStreamsLike().LessThan(o.GetActivityStreamsLike())
	} else if this.IsActivityStreamsListen() {
		return this.GetActivityStreamsListen().LessThan(o.GetActivityStreamsListen())
	} else if this.IsActivityStreamsMention() {
		return this.GetActivityStreamsMention().LessThan(o.GetActivityStreamsMention())
	} else if this.IsActivityStreamsMove() {
		return this.GetActivityStreamsMove().LessThan(o.GetActivityStreamsMove())
	} else if this.IsActivityStreamsNote() {
		return this.GetActivityStreamsNote().LessThan(o.GetActivityStreamsNote())
	} else if this.IsActivityStreamsOffer() {
		return this.GetActivityStreamsOffer().LessThan(o.GetActivityStreamsOffer())
	} else if this.IsActivityStreamsOrderedCollection() {
		return this.GetActivityStreamsOrderedCollection().LessThan(o.GetActivityStreamsOrderedCollection())
	} else if this.IsActivityStreamsOrderedCollectionPage() {
		return this.GetActivityStreamsOrderedCollectionPage().LessThan(o.GetActivityStreamsOrderedCollectionPage())
	} else if this.IsActivityStreamsOrganization() {
		return this.GetActivityStreamsOrganization().LessThan(o.GetActivityStreamsOrganization())
	} else if this.IsActivityStreamsPage() {
		return this.GetActivityStreamsPage().LessThan(o.GetActivityStreamsPage())
	} else if this.IsActivityStreamsPerson() {
		return this.GetActivityStreamsPerson().LessThan(o.GetActivityStreamsPerson())
	} else if this.IsActivityStreamsPlace() {
		return this.GetActivityStreamsPlace().LessThan(o.GetActivityStreamsPlace())
	} else if this.IsActivityStreamsProfile() {
		return this.GetActivityStreamsProfile().LessThan(o.GetActivityStreamsProfile())
	} else if this.IsActivityStreamsQuestion() {
		return this.GetActivityStreamsQuestion().LessThan(o.GetActivityStreamsQuestion())
	} else if this.IsActivityStreamsRead() {
		return this.GetActivityStreamsRead().LessThan(o.GetActivityStreamsRead())
	} else if this.IsActivityStreamsReject() {
		return this.GetActivityStreamsReject().LessThan(o.GetActivityStreamsReject())
	} else if this.IsActivityStreamsRelationship() {
		return this.GetActivityStreamsRelationship().LessThan(o.GetActivityStreamsRelationship())
	} else if this.IsActivityStreamsRemove() {
		return this.GetActivityStreamsRemove().LessThan(o.GetActivityStreamsRemove())
	} else if this.IsActivityStreamsService() {
		return this.GetActivityStreamsService().LessThan(o.GetActivityStreamsService())
	} else if this.IsActivityStreamsTentativeAccept() {
		return this.GetActivityStreamsTentativeAccept().LessThan(o.GetActivityStreamsTentativeAccept())
	} else if this.IsActivityStreamsTentativeReject() {
		return this.GetActivityStreamsTentativeReject().LessThan(o.GetActivityStreamsTentativeReject())
	} else if this.IsActivityStreamsTombstone() {
		return this.GetActivityStreamsTombstone().LessThan(o.GetActivityStreamsTombstone())
	} else if this.IsActivityStreamsTravel() {
		return this.GetActivityStreamsTravel().LessThan(o.GetActivityStreamsTravel())
	} else if this.IsActivityStreamsUndo() {
		return this.GetActivityStreamsUndo().LessThan(o.GetActivityStreamsUndo())
	} else if this.IsActivityStreamsUpdate() {
		return this.GetActivityStreamsUpdate().LessThan(o.GetActivityStreamsUpdate())
	} else if this.IsActivityStreamsVideo() {
		return this.GetActivityStreamsVideo().LessThan(o.GetActivityStreamsVideo())
	} else if this.IsActivityStreamsView() {
		return this.GetActivityStreamsView().LessThan(o.GetActivityStreamsView())
	} else if this.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	}
	return false
}

// Name returns the name of this property: "ActivityStreamsGenerator".
func (this ActivityStreamsGeneratorPropertyIterator) Name() string {
	return "ActivityStreamsGenerator"
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this ActivityStreamsGeneratorPropertyIterator) Next() vocab.ActivityStreamsGeneratorPropertyIterator {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this ActivityStreamsGeneratorPropertyIterator) Prev() vocab.ActivityStreamsGeneratorPropertyIterator {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// SetActivityStreamsAccept sets the value of this property. Calling
// IsActivityStreamsAccept afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsAccept(v vocab.ActivityStreamsAccept) {
	this.clear()
	this.activitystreamsAcceptMember = v
}

// SetActivityStreamsActivity sets the value of this property. Calling
// IsActivityStreamsActivity afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsActivity(v vocab.ActivityStreamsActivity) {
	this.clear()
	this.activitystreamsActivityMember = v
}

// SetActivityStreamsAdd sets the value of this property. Calling
// IsActivityStreamsAdd afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsAdd(v vocab.ActivityStreamsAdd) {
	this.clear()
	this.activitystreamsAddMember = v
}

// SetActivityStreamsAnnounce sets the value of this property. Calling
// IsActivityStreamsAnnounce afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsAnnounce(v vocab.ActivityStreamsAnnounce) {
	this.clear()
	this.activitystreamsAnnounceMember = v
}

// SetActivityStreamsApplication sets the value of this property. Calling
// IsActivityStreamsApplication afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsApplication(v vocab.ActivityStreamsApplication) {
	this.clear()
	this.activitystreamsApplicationMember = v
}

// SetActivityStreamsArrive sets the value of this property. Calling
// IsActivityStreamsArrive afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsArrive(v vocab.ActivityStreamsArrive) {
	this.clear()
	this.activitystreamsArriveMember = v
}

// SetActivityStreamsArticle sets the value of this property. Calling
// IsActivityStreamsArticle afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsArticle(v vocab.ActivityStreamsArticle) {
	this.clear()
	this.activitystreamsArticleMember = v
}

// SetActivityStreamsAudio sets the value of this property. Calling
// IsActivityStreamsAudio afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsAudio(v vocab.ActivityStreamsAudio) {
	this.clear()
	this.activitystreamsAudioMember = v
}

// SetActivityStreamsBlock sets the value of this property. Calling
// IsActivityStreamsBlock afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsBlock(v vocab.ActivityStreamsBlock) {
	this.clear()
	this.activitystreamsBlockMember = v
}

// SetActivityStreamsCollection sets the value of this property. Calling
// IsActivityStreamsCollection afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsCollection(v vocab.ActivityStreamsCollection) {
	this.clear()
	this.activitystreamsCollectionMember = v
}

// SetActivityStreamsCollectionPage sets the value of this property. Calling
// IsActivityStreamsCollectionPage afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsCollectionPage(v vocab.ActivityStreamsCollectionPage) {
	this.clear()
	this.activitystreamsCollectionPageMember = v
}

// SetActivityStreamsCreate sets the value of this property. Calling
// IsActivityStreamsCreate afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsCreate(v vocab.ActivityStreamsCreate) {
	this.clear()
	this.activitystreamsCreateMember = v
}

// SetActivityStreamsDelete sets the value of this property. Calling
// IsActivityStreamsDelete afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsDelete(v vocab.ActivityStreamsDelete) {
	this.clear()
	this.activitystreamsDeleteMember = v
}

// SetActivityStreamsDislike sets the value of this property. Calling
// IsActivityStreamsDislike afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsDislike(v vocab.ActivityStreamsDislike) {
	this.clear()
	this.activitystreamsDislikeMember = v
}

// SetActivityStreamsDocument sets the value of this property. Calling
// IsActivityStreamsDocument afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsDocument(v vocab.ActivityStreamsDocument) {
	this.clear()
	this.activitystreamsDocumentMember = v
}

// SetActivityStreamsEvent sets the value of this property. Calling
// IsActivityStreamsEvent afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsEvent(v vocab.ActivityStreamsEvent) {
	this.clear()
	this.activitystreamsEventMember = v
}

// SetActivityStreamsFlag sets the value of this property. Calling
// IsActivityStreamsFlag afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsFlag(v vocab.ActivityStreamsFlag) {
	this.clear()
	this.activitystreamsFlagMember = v
}

// SetActivityStreamsFollow sets the value of this property. Calling
// IsActivityStreamsFollow afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsFollow(v vocab.ActivityStreamsFollow) {
	this.clear()
	this.activitystreamsFollowMember = v
}

// SetActivityStreamsGroup sets the value of this property. Calling
// IsActivityStreamsGroup afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsGroup(v vocab.ActivityStreamsGroup) {
	this.clear()
	this.activitystreamsGroupMember = v
}

// SetActivityStreamsIgnore sets the value of this property. Calling
// IsActivityStreamsIgnore afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsIgnore(v vocab.ActivityStreamsIgnore) {
	this.clear()
	this.activitystreamsIgnoreMember = v
}

// SetActivityStreamsImage sets the value of this property. Calling
// IsActivityStreamsImage afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsImage(v vocab.ActivityStreamsImage) {
	this.clear()
	this.activitystreamsImageMember = v
}

// SetActivityStreamsIntransitiveActivity sets the value of this property. Calling
// IsActivityStreamsIntransitiveActivity afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsIntransitiveActivity(v vocab.ActivityStreamsIntransitiveActivity) {
	this.clear()
	this.activitystreamsIntransitiveActivityMember = v
}

// SetActivityStreamsInvite sets the value of this property. Calling
// IsActivityStreamsInvite afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsInvite(v vocab.ActivityStreamsInvite) {
	this.clear()
	this.activitystreamsInviteMember = v
}

// SetActivityStreamsJoin sets the value of this property. Calling
// IsActivityStreamsJoin afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsJoin(v vocab.ActivityStreamsJoin) {
	this.clear()
	this.activitystreamsJoinMember = v
}

// SetActivityStreamsLeave sets the value of this property. Calling
// IsActivityStreamsLeave afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsLeave(v vocab.ActivityStreamsLeave) {
	this.clear()
	this.activitystreamsLeaveMember = v
}

// SetActivityStreamsLike sets the value of this property. Calling
// IsActivityStreamsLike afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsLike(v vocab.ActivityStreamsLike) {
	this.clear()
	this.activitystreamsLikeMember = v
}

// SetActivityStreamsLink sets the value of this property. Calling
// IsActivityStreamsLink afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsLink(v vocab.ActivityStreamsLink) {
	this.clear()
	this.activitystreamsLinkMember = v
}

// SetActivityStreamsListen sets the value of this property. Calling
// IsActivityStreamsListen afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsListen(v vocab.ActivityStreamsListen) {
	this.clear()
	this.activitystreamsListenMember = v
}

// SetActivityStreamsMention sets the value of this property. Calling
// IsActivityStreamsMention afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsMention(v vocab.ActivityStreamsMention) {
	this.clear()
	this.activitystreamsMentionMember = v
}

// SetActivityStreamsMove sets the value of this property. Calling
// IsActivityStreamsMove afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsMove(v vocab.ActivityStreamsMove) {
	this.clear()
	this.activitystreamsMoveMember = v
}

// SetActivityStreamsNote sets the value of this property. Calling
// IsActivityStreamsNote afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsNote(v vocab.ActivityStreamsNote) {
	this.clear()
	this.activitystreamsNoteMember = v
}

// SetActivityStreamsObject sets the value of this property. Calling
// IsActivityStreamsObject afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsObject(v vocab.ActivityStreamsObject) {
	this.clear()
	this.activitystreamsObjectMember = v
}

// SetActivityStreamsOffer sets the value of this property. Calling
// IsActivityStreamsOffer afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsOffer(v vocab.ActivityStreamsOffer) {
	this.clear()
	this.activitystreamsOfferMember = v
}

// SetActivityStreamsOrderedCollection sets the value of this property. Calling
// IsActivityStreamsOrderedCollection afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsOrderedCollection(v vocab.ActivityStreamsOrderedCollection) {
	this.clear()
	this.activitystreamsOrderedCollectionMember = v
}

// SetActivityStreamsOrderedCollectionPage sets the value of this property.
// Calling IsActivityStreamsOrderedCollectionPage afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsOrderedCollectionPage(v vocab.ActivityStreamsOrderedCollectionPage) {
	this.clear()
	this.activitystreamsOrderedCollectionPageMember = v
}

// SetActivityStreamsOrganization sets the value of this property. Calling
// IsActivityStreamsOrganization afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsOrganization(v vocab.ActivityStreamsOrganization) {
	this.clear()
	this.activitystreamsOrganizationMember = v
}

// SetActivityStreamsPage sets the value of this property. Calling
// IsActivityStreamsPage afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsPage(v vocab.ActivityStreamsPage) {
	this.clear()
	this.activitystreamsPageMember = v
}

// SetActivityStreamsPerson sets the value of this property. Calling
// IsActivityStreamsPerson afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsPerson(v vocab.ActivityStreamsPerson) {
	this.clear()
	this.activitystreamsPersonMember = v
}

// SetActivityStreamsPlace sets the value of this property. Calling
// IsActivityStreamsPlace afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsPlace(v vocab.ActivityStreamsPlace) {
	this.clear()
	this.activitystreamsPlaceMember = v
}

// SetActivityStreamsProfile sets the value of this property. Calling
// IsActivityStreamsProfile afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsProfile(v vocab.ActivityStreamsProfile) {
	this.clear()
	this.activitystreamsProfileMember = v
}

// SetActivityStreamsQuestion sets the value of this property. Calling
// IsActivityStreamsQuestion afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsQuestion(v vocab.ActivityStreamsQuestion) {
	this.clear()
	this.activitystreamsQuestionMember = v
}

// SetActivityStreamsRead sets the value of this property. Calling
// IsActivityStreamsRead afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsRead(v vocab.ActivityStreamsRead) {
	this.clear()
	this.activitystreamsReadMember = v
}

// SetActivityStreamsReject sets the value of this property. Calling
// IsActivityStreamsReject afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsReject(v vocab.ActivityStreamsReject) {
	this.clear()
	this.activitystreamsRejectMember = v
}

// SetActivityStreamsRelationship sets the value of this property. Calling
// IsActivityStreamsRelationship afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsRelationship(v vocab.ActivityStreamsRelationship) {
	this.clear()
	this.activitystreamsRelationshipMember = v
}

// SetActivityStreamsRemove sets the value of this property. Calling
// IsActivityStreamsRemove afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsRemove(v vocab.ActivityStreamsRemove) {
	this.clear()
	this.activitystreamsRemoveMember = v
}

// SetActivityStreamsService sets the value of this property. Calling
// IsActivityStreamsService afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsService(v vocab.ActivityStreamsService) {
	this.clear()
	this.activitystreamsServiceMember = v
}

// SetActivityStreamsTentativeAccept sets the value of this property. Calling
// IsActivityStreamsTentativeAccept afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsTentativeAccept(v vocab.ActivityStreamsTentativeAccept) {
	this.clear()
	this.activitystreamsTentativeAcceptMember = v
}

// SetActivityStreamsTentativeReject sets the value of this property. Calling
// IsActivityStreamsTentativeReject afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsTentativeReject(v vocab.ActivityStreamsTentativeReject) {
	this.clear()
	this.activitystreamsTentativeRejectMember = v
}

// SetActivityStreamsTombstone sets the value of this property. Calling
// IsActivityStreamsTombstone afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsTombstone(v vocab.ActivityStreamsTombstone) {
	this.clear()
	this.activitystreamsTombstoneMember = v
}

// SetActivityStreamsTravel sets the value of this property. Calling
// IsActivityStreamsTravel afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsTravel(v vocab.ActivityStreamsTravel) {
	this.clear()
	this.activitystreamsTravelMember = v
}

// SetActivityStreamsUndo sets the value of this property. Calling
// IsActivityStreamsUndo afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsUndo(v vocab.ActivityStreamsUndo) {
	this.clear()
	this.activitystreamsUndoMember = v
}

// SetActivityStreamsUpdate sets the value of this property. Calling
// IsActivityStreamsUpdate afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsUpdate(v vocab.ActivityStreamsUpdate) {
	this.clear()
	this.activitystreamsUpdateMember = v
}

// SetActivityStreamsVideo sets the value of this property. Calling
// IsActivityStreamsVideo afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsVideo(v vocab.ActivityStreamsVideo) {
	this.clear()
	this.activitystreamsVideoMember = v
}

// SetActivityStreamsView sets the value of this property. Calling
// IsActivityStreamsView afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetActivityStreamsView(v vocab.ActivityStreamsView) {
	this.clear()
	this.activitystreamsViewMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *ActivityStreamsGeneratorPropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.iri = v
}

// clear ensures no value of this property is set. Calling HasAny or any of the
// 'Is' methods afterwards will return false.
func (this *ActivityStreamsGeneratorPropertyIterator) clear() {
	this.activitystreamsObjectMember = nil
	this.activitystreamsLinkMember = nil
	this.activitystreamsAcceptMember = nil
	this.activitystreamsActivityMember = nil
	this.activitystreamsAddMember = nil
	this.activitystreamsAnnounceMember = nil
	this.activitystreamsApplicationMember = nil
	this.activitystreamsArriveMember = nil
	this.activitystreamsArticleMember = nil
	this.activitystreamsAudioMember = nil
	this.activitystreamsBlockMember = nil
	this.activitystreamsCollectionMember = nil
	this.activitystreamsCollectionPageMember = nil
	this.activitystreamsCreateMember = nil
	this.activitystreamsDeleteMember = nil
	this.activitystreamsDislikeMember = nil
	this.activitystreamsDocumentMember = nil
	this.activitystreamsEventMember = nil
	this.activitystreamsFlagMember = nil
	this.activitystreamsFollowMember = nil
	this.activitystreamsGroupMember = nil
	this.activitystreamsIgnoreMember = nil
	this.activitystreamsImageMember = nil
	this.activitystreamsIntransitiveActivityMember = nil
	this.activitystreamsInviteMember = nil
	this.activitystreamsJoinMember = nil
	this.activitystreamsLeaveMember = nil
	this.activitystreamsLikeMember = nil
	this.activitystreamsListenMember = nil
	this.activitystreamsMentionMember = nil
	this.activitystreamsMoveMember = nil
	this.activitystreamsNoteMember = nil
	this.activitystreamsOfferMember = nil
	this.activitystreamsOrderedCollectionMember = nil
	this.activitystreamsOrderedCollectionPageMember = nil
	this.activitystreamsOrganizationMember = nil
	this.activitystreamsPageMember = nil
	this.activitystreamsPersonMember = nil
	this.activitystreamsPlaceMember = nil
	this.activitystreamsProfileMember = nil
	this.activitystreamsQuestionMember = nil
	this.activitystreamsReadMember = nil
	this.activitystreamsRejectMember = nil
	this.activitystreamsRelationshipMember = nil
	this.activitystreamsRemoveMember = nil
	this.activitystreamsServiceMember = nil
	this.activitystreamsTentativeAcceptMember = nil
	this.activitystreamsTentativeRejectMember = nil
	this.activitystreamsTombstoneMember = nil
	this.activitystreamsTravelMember = nil
	this.activitystreamsUndoMember = nil
	this.activitystreamsUpdateMember = nil
	this.activitystreamsVideoMember = nil
	this.activitystreamsViewMember = nil
	this.unknown = nil
	this.iri = nil
}

// serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsGeneratorPropertyIterator) serialize() (interface{}, error) {
	if this.IsActivityStreamsObject() {
		return this.GetActivityStreamsObject().Serialize()
	} else if this.IsActivityStreamsLink() {
		return this.GetActivityStreamsLink().Serialize()
	} else if this.IsActivityStreamsAccept() {
		return this.GetActivityStreamsAccept().Serialize()
	} else if this.IsActivityStreamsActivity() {
		return this.GetActivityStreamsActivity().Serialize()
	} else if this.IsActivityStreamsAdd() {
		return this.GetActivityStreamsAdd().Serialize()
	} else if this.IsActivityStreamsAnnounce() {
		return this.GetActivityStreamsAnnounce().Serialize()
	} else if this.IsActivityStreamsApplication() {
		return this.GetActivityStreamsApplication().Serialize()
	} else if this.IsActivityStreamsArrive() {
		return this.GetActivityStreamsArrive().Serialize()
	} else if this.IsActivityStreamsArticle() {
		return this.GetActivityStreamsArticle().Serialize()
	} else if this.IsActivityStreamsAudio() {
		return this.GetActivityStreamsAudio().Serialize()
	} else if this.IsActivityStreamsBlock() {
		return this.GetActivityStreamsBlock().Serialize()
	} else if this.IsActivityStreamsCollection() {
		return this.GetActivityStreamsCollection().Serialize()
	} else if this.IsActivityStreamsCollectionPage() {
		return this.GetActivityStreamsCollectionPage().Serialize()
	} else if this.IsActivityStreamsCreate() {
		return this.GetActivityStreamsCreate().Serialize()
	} else if this.IsActivityStreamsDelete() {
		return this.GetActivityStreamsDelete().Serialize()
	} else if this.IsActivityStreamsDislike() {
		return this.GetActivityStreamsDislike().Serialize()
	} else if this.IsActivityStreamsDocument() {
		return this.GetActivityStreamsDocument().Serialize()
	} else if this.IsActivityStreamsEvent() {
		return this.GetActivityStreamsEvent().Serialize()
	} else if this.IsActivityStreamsFlag() {
		return this.GetActivityStreamsFlag().Serialize()
	} else if this.IsActivityStreamsFollow() {
		return this.GetActivityStreamsFollow().Serialize()
	} else if this.IsActivityStreamsGroup() {
		return this.GetActivityStreamsGroup().Serialize()
	} else if this.IsActivityStreamsIgnore() {
		return this.GetActivityStreamsIgnore().Serialize()
	} else if this.IsActivityStreamsImage() {
		return this.GetActivityStreamsImage().Serialize()
	} else if this.IsActivityStreamsIntransitiveActivity() {
		return this.GetActivityStreamsIntransitiveActivity().Serialize()
	} else if this.IsActivityStreamsInvite() {
		return this.GetActivityStreamsInvite().Serialize()
	} else if this.IsActivityStreamsJoin() {
		return this.GetActivityStreamsJoin().Serialize()
	} else if this.IsActivityStreamsLeave() {
		return this.GetActivityStreamsLeave().Serialize()
	} else if this.IsActivityStreamsLike() {
		return this.GetActivityStreamsLike().Serialize()
	} else if this.IsActivityStreamsListen() {
		return this.GetActivityStreamsListen().Serialize()
	} else if this.IsActivityStreamsMention() {
		return this.GetActivityStreamsMention().Serialize()
	} else if this.IsActivityStreamsMove() {
		return this.GetActivityStreamsMove().Serialize()
	} else if this.IsActivityStreamsNote() {
		return this.GetActivityStreamsNote().Serialize()
	} else if this.IsActivityStreamsOffer() {
		return this.GetActivityStreamsOffer().Serialize()
	} else if this.IsActivityStreamsOrderedCollection() {
		return this.GetActivityStreamsOrderedCollection().Serialize()
	} else if this.IsActivityStreamsOrderedCollectionPage() {
		return this.GetActivityStreamsOrderedCollectionPage().Serialize()
	} else if this.IsActivityStreamsOrganization() {
		return this.GetActivityStreamsOrganization().Serialize()
	} else if this.IsActivityStreamsPage() {
		return this.GetActivityStreamsPage().Serialize()
	} else if this.IsActivityStreamsPerson() {
		return this.GetActivityStreamsPerson().Serialize()
	} else if this.IsActivityStreamsPlace() {
		return this.GetActivityStreamsPlace().Serialize()
	} else if this.IsActivityStreamsProfile() {
		return this.GetActivityStreamsProfile().Serialize()
	} else if this.IsActivityStreamsQuestion() {
		return this.GetActivityStreamsQuestion().Serialize()
	} else if this.IsActivityStreamsRead() {
		return this.GetActivityStreamsRead().Serialize()
	} else if this.IsActivityStreamsReject() {
		return this.GetActivityStreamsReject().Serialize()
	} else if this.IsActivityStreamsRelationship() {
		return this.GetActivityStreamsRelationship().Serialize()
	} else if this.IsActivityStreamsRemove() {
		return this.GetActivityStreamsRemove().Serialize()
	} else if this.IsActivityStreamsService() {
		return this.GetActivityStreamsService().Serialize()
	} else if this.IsActivityStreamsTentativeAccept() {
		return this.GetActivityStreamsTentativeAccept().Serialize()
	} else if this.IsActivityStreamsTentativeReject() {
		return this.GetActivityStreamsTentativeReject().Serialize()
	} else if this.IsActivityStreamsTombstone() {
		return this.GetActivityStreamsTombstone().Serialize()
	} else if this.IsActivityStreamsTravel() {
		return this.GetActivityStreamsTravel().Serialize()
	} else if this.IsActivityStreamsUndo() {
		return this.GetActivityStreamsUndo().Serialize()
	} else if this.IsActivityStreamsUpdate() {
		return this.GetActivityStreamsUpdate().Serialize()
	} else if this.IsActivityStreamsVideo() {
		return this.GetActivityStreamsVideo().Serialize()
	} else if this.IsActivityStreamsView() {
		return this.GetActivityStreamsView().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// ActivityStreamsGeneratorProperty is the non-functional property "generator". It
// is permitted to have one or more values, and of different value types.
type ActivityStreamsGeneratorProperty struct {
	properties []*ActivityStreamsGeneratorPropertyIterator
	alias      string
}

// DeserializeGeneratorProperty creates a "generator" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeGeneratorProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsGeneratorProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "generator"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "generator")
	}
	i, ok := m[propName]

	if ok {
		this := &ActivityStreamsGeneratorProperty{
			alias:      alias,
			properties: []*ActivityStreamsGeneratorPropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeActivityStreamsGeneratorPropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeActivityStreamsGeneratorPropertyIterator(i, aliasMap); err != nil {
				return this, err
			} else if p != nil {
				this.properties = append(this.properties, p)
			}
		}
		// Set up the properties for iteration.
		for idx, ele := range this.properties {
			ele.parent = this
			ele.myIdx = idx
		}
		return this, nil
	}
	return nil, nil
}

// NewActivityStreamsGeneratorProperty creates a new generator property.
func NewActivityStreamsGeneratorProperty() *ActivityStreamsGeneratorProperty {
	return &ActivityStreamsGeneratorProperty{alias: ""}
}

// AppendActivityStreamsAccept appends a Accept value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsAccept(v vocab.ActivityStreamsAccept) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsAcceptMember: v,
		alias:                       this.alias,
		myIdx:                       this.Len(),
		parent:                      this,
	})
}

// AppendActivityStreamsActivity appends a Activity value to the back of a list of
// the property "generator". Invalidates iterators that are traversing using
// Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsActivity(v vocab.ActivityStreamsActivity) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsActivityMember: v,
		alias:                         this.alias,
		myIdx:                         this.Len(),
		parent:                        this,
	})
}

// AppendActivityStreamsAdd appends a Add value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsAdd(v vocab.ActivityStreamsAdd) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsAddMember: v,
		alias:                    this.alias,
		myIdx:                    this.Len(),
		parent:                   this,
	})
}

// AppendActivityStreamsAnnounce appends a Announce value to the back of a list of
// the property "generator". Invalidates iterators that are traversing using
// Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsAnnounce(v vocab.ActivityStreamsAnnounce) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsAnnounceMember: v,
		alias:                         this.alias,
		myIdx:                         this.Len(),
		parent:                        this,
	})
}

// AppendActivityStreamsApplication appends a Application value to the back of a
// list of the property "generator". Invalidates iterators that are traversing
// using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsApplication(v vocab.ActivityStreamsApplication) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsApplicationMember: v,
		alias:                            this.alias,
		myIdx:                            this.Len(),
		parent:                           this,
	})
}

// AppendActivityStreamsArrive appends a Arrive value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsArrive(v vocab.ActivityStreamsArrive) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsArriveMember: v,
		alias:                       this.alias,
		myIdx:                       this.Len(),
		parent:                      this,
	})
}

// AppendActivityStreamsArticle appends a Article value to the back of a list of
// the property "generator". Invalidates iterators that are traversing using
// Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsArticle(v vocab.ActivityStreamsArticle) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsArticleMember: v,
		alias:                        this.alias,
		myIdx:                        this.Len(),
		parent:                       this,
	})
}

// AppendActivityStreamsAudio appends a Audio value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsAudio(v vocab.ActivityStreamsAudio) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsAudioMember: v,
		alias:                      this.alias,
		myIdx:                      this.Len(),
		parent:                     this,
	})
}

// AppendActivityStreamsBlock appends a Block value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsBlock(v vocab.ActivityStreamsBlock) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsBlockMember: v,
		alias:                      this.alias,
		myIdx:                      this.Len(),
		parent:                     this,
	})
}

// AppendActivityStreamsCollection appends a Collection value to the back of a
// list of the property "generator". Invalidates iterators that are traversing
// using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsCollection(v vocab.ActivityStreamsCollection) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsCollectionMember: v,
		alias:                           this.alias,
		myIdx:                           this.Len(),
		parent:                          this,
	})
}

// AppendActivityStreamsCollectionPage appends a CollectionPage value to the back
// of a list of the property "generator". Invalidates iterators that are
// traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsCollectionPage(v vocab.ActivityStreamsCollectionPage) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsCollectionPageMember: v,
		alias:                               this.alias,
		myIdx:                               this.Len(),
		parent:                              this,
	})
}

// AppendActivityStreamsCreate appends a Create value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsCreate(v vocab.ActivityStreamsCreate) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsCreateMember: v,
		alias:                       this.alias,
		myIdx:                       this.Len(),
		parent:                      this,
	})
}

// AppendActivityStreamsDelete appends a Delete value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsDelete(v vocab.ActivityStreamsDelete) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsDeleteMember: v,
		alias:                       this.alias,
		myIdx:                       this.Len(),
		parent:                      this,
	})
}

// AppendActivityStreamsDislike appends a Dislike value to the back of a list of
// the property "generator". Invalidates iterators that are traversing using
// Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsDislike(v vocab.ActivityStreamsDislike) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsDislikeMember: v,
		alias:                        this.alias,
		myIdx:                        this.Len(),
		parent:                       this,
	})
}

// AppendActivityStreamsDocument appends a Document value to the back of a list of
// the property "generator". Invalidates iterators that are traversing using
// Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsDocument(v vocab.ActivityStreamsDocument) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsDocumentMember: v,
		alias:                         this.alias,
		myIdx:                         this.Len(),
		parent:                        this,
	})
}

// AppendActivityStreamsEvent appends a Event value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsEvent(v vocab.ActivityStreamsEvent) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsEventMember: v,
		alias:                      this.alias,
		myIdx:                      this.Len(),
		parent:                     this,
	})
}

// AppendActivityStreamsFlag appends a Flag value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsFlag(v vocab.ActivityStreamsFlag) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsFlagMember: v,
		alias:                     this.alias,
		myIdx:                     this.Len(),
		parent:                    this,
	})
}

// AppendActivityStreamsFollow appends a Follow value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsFollow(v vocab.ActivityStreamsFollow) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsFollowMember: v,
		alias:                       this.alias,
		myIdx:                       this.Len(),
		parent:                      this,
	})
}

// AppendActivityStreamsGroup appends a Group value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsGroup(v vocab.ActivityStreamsGroup) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsGroupMember: v,
		alias:                      this.alias,
		myIdx:                      this.Len(),
		parent:                     this,
	})
}

// AppendActivityStreamsIgnore appends a Ignore value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsIgnore(v vocab.ActivityStreamsIgnore) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsIgnoreMember: v,
		alias:                       this.alias,
		myIdx:                       this.Len(),
		parent:                      this,
	})
}

// AppendActivityStreamsImage appends a Image value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsImage(v vocab.ActivityStreamsImage) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsImageMember: v,
		alias:                      this.alias,
		myIdx:                      this.Len(),
		parent:                     this,
	})
}

// AppendActivityStreamsIntransitiveActivity appends a IntransitiveActivity value
// to the back of a list of the property "generator". Invalidates iterators
// that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsIntransitiveActivity(v vocab.ActivityStreamsIntransitiveActivity) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsIntransitiveActivityMember: v,
		alias:  this.alias,
		myIdx:  this.Len(),
		parent: this,
	})
}

// AppendActivityStreamsInvite appends a Invite value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsInvite(v vocab.ActivityStreamsInvite) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsInviteMember: v,
		alias:                       this.alias,
		myIdx:                       this.Len(),
		parent:                      this,
	})
}

// AppendActivityStreamsJoin appends a Join value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsJoin(v vocab.ActivityStreamsJoin) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsJoinMember: v,
		alias:                     this.alias,
		myIdx:                     this.Len(),
		parent:                    this,
	})
}

// AppendActivityStreamsLeave appends a Leave value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsLeave(v vocab.ActivityStreamsLeave) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsLeaveMember: v,
		alias:                      this.alias,
		myIdx:                      this.Len(),
		parent:                     this,
	})
}

// AppendActivityStreamsLike appends a Like value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsLike(v vocab.ActivityStreamsLike) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsLikeMember: v,
		alias:                     this.alias,
		myIdx:                     this.Len(),
		parent:                    this,
	})
}

// AppendActivityStreamsLink appends a Link value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsLink(v vocab.ActivityStreamsLink) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsLinkMember: v,
		alias:                     this.alias,
		myIdx:                     this.Len(),
		parent:                    this,
	})
}

// AppendActivityStreamsListen appends a Listen value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsListen(v vocab.ActivityStreamsListen) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsListenMember: v,
		alias:                       this.alias,
		myIdx:                       this.Len(),
		parent:                      this,
	})
}

// AppendActivityStreamsMention appends a Mention value to the back of a list of
// the property "generator". Invalidates iterators that are traversing using
// Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsMention(v vocab.ActivityStreamsMention) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsMentionMember: v,
		alias:                        this.alias,
		myIdx:                        this.Len(),
		parent:                       this,
	})
}

// AppendActivityStreamsMove appends a Move value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsMove(v vocab.ActivityStreamsMove) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsMoveMember: v,
		alias:                     this.alias,
		myIdx:                     this.Len(),
		parent:                    this,
	})
}

// AppendActivityStreamsNote appends a Note value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsNote(v vocab.ActivityStreamsNote) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsNoteMember: v,
		alias:                     this.alias,
		myIdx:                     this.Len(),
		parent:                    this,
	})
}

// AppendActivityStreamsObject appends a Object value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsObject(v vocab.ActivityStreamsObject) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsObjectMember: v,
		alias:                       this.alias,
		myIdx:                       this.Len(),
		parent:                      this,
	})
}

// AppendActivityStreamsOffer appends a Offer value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsOffer(v vocab.ActivityStreamsOffer) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsOfferMember: v,
		alias:                      this.alias,
		myIdx:                      this.Len(),
		parent:                     this,
	})
}

// AppendActivityStreamsOrderedCollection appends a OrderedCollection value to the
// back of a list of the property "generator". Invalidates iterators that are
// traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsOrderedCollection(v vocab.ActivityStreamsOrderedCollection) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsOrderedCollectionMember: v,
		alias:                                  this.alias,
		myIdx:                                  this.Len(),
		parent:                                 this,
	})
}

// AppendActivityStreamsOrderedCollectionPage appends a OrderedCollectionPage
// value to the back of a list of the property "generator". Invalidates
// iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsOrderedCollectionPage(v vocab.ActivityStreamsOrderedCollectionPage) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsOrderedCollectionPageMember: v,
		alias:  this.alias,
		myIdx:  this.Len(),
		parent: this,
	})
}

// AppendActivityStreamsOrganization appends a Organization value to the back of a
// list of the property "generator". Invalidates iterators that are traversing
// using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsOrganization(v vocab.ActivityStreamsOrganization) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsOrganizationMember: v,
		alias:                             this.alias,
		myIdx:                             this.Len(),
		parent:                            this,
	})
}

// AppendActivityStreamsPage appends a Page value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsPage(v vocab.ActivityStreamsPage) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsPageMember: v,
		alias:                     this.alias,
		myIdx:                     this.Len(),
		parent:                    this,
	})
}

// AppendActivityStreamsPerson appends a Person value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsPerson(v vocab.ActivityStreamsPerson) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsPersonMember: v,
		alias:                       this.alias,
		myIdx:                       this.Len(),
		parent:                      this,
	})
}

// AppendActivityStreamsPlace appends a Place value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsPlace(v vocab.ActivityStreamsPlace) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsPlaceMember: v,
		alias:                      this.alias,
		myIdx:                      this.Len(),
		parent:                     this,
	})
}

// AppendActivityStreamsProfile appends a Profile value to the back of a list of
// the property "generator". Invalidates iterators that are traversing using
// Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsProfile(v vocab.ActivityStreamsProfile) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsProfileMember: v,
		alias:                        this.alias,
		myIdx:                        this.Len(),
		parent:                       this,
	})
}

// AppendActivityStreamsQuestion appends a Question value to the back of a list of
// the property "generator". Invalidates iterators that are traversing using
// Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsQuestion(v vocab.ActivityStreamsQuestion) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsQuestionMember: v,
		alias:                         this.alias,
		myIdx:                         this.Len(),
		parent:                        this,
	})
}

// AppendActivityStreamsRead appends a Read value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsRead(v vocab.ActivityStreamsRead) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsReadMember: v,
		alias:                     this.alias,
		myIdx:                     this.Len(),
		parent:                    this,
	})
}

// AppendActivityStreamsReject appends a Reject value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsReject(v vocab.ActivityStreamsReject) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsRejectMember: v,
		alias:                       this.alias,
		myIdx:                       this.Len(),
		parent:                      this,
	})
}

// AppendActivityStreamsRelationship appends a Relationship value to the back of a
// list of the property "generator". Invalidates iterators that are traversing
// using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsRelationship(v vocab.ActivityStreamsRelationship) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsRelationshipMember: v,
		alias:                             this.alias,
		myIdx:                             this.Len(),
		parent:                            this,
	})
}

// AppendActivityStreamsRemove appends a Remove value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsRemove(v vocab.ActivityStreamsRemove) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsRemoveMember: v,
		alias:                       this.alias,
		myIdx:                       this.Len(),
		parent:                      this,
	})
}

// AppendActivityStreamsService appends a Service value to the back of a list of
// the property "generator". Invalidates iterators that are traversing using
// Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsService(v vocab.ActivityStreamsService) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsServiceMember: v,
		alias:                        this.alias,
		myIdx:                        this.Len(),
		parent:                       this,
	})
}

// AppendActivityStreamsTentativeAccept appends a TentativeAccept value to the
// back of a list of the property "generator". Invalidates iterators that are
// traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsTentativeAccept(v vocab.ActivityStreamsTentativeAccept) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsTentativeAcceptMember: v,
		alias:                                this.alias,
		myIdx:                                this.Len(),
		parent:                               this,
	})
}

// AppendActivityStreamsTentativeReject appends a TentativeReject value to the
// back of a list of the property "generator". Invalidates iterators that are
// traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsTentativeReject(v vocab.ActivityStreamsTentativeReject) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsTentativeRejectMember: v,
		alias:                                this.alias,
		myIdx:                                this.Len(),
		parent:                               this,
	})
}

// AppendActivityStreamsTombstone appends a Tombstone value to the back of a list
// of the property "generator". Invalidates iterators that are traversing
// using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsTombstone(v vocab.ActivityStreamsTombstone) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsTombstoneMember: v,
		alias:                          this.alias,
		myIdx:                          this.Len(),
		parent:                         this,
	})
}

// AppendActivityStreamsTravel appends a Travel value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsTravel(v vocab.ActivityStreamsTravel) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsTravelMember: v,
		alias:                       this.alias,
		myIdx:                       this.Len(),
		parent:                      this,
	})
}

// AppendActivityStreamsUndo appends a Undo value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsUndo(v vocab.ActivityStreamsUndo) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsUndoMember: v,
		alias:                     this.alias,
		myIdx:                     this.Len(),
		parent:                    this,
	})
}

// AppendActivityStreamsUpdate appends a Update value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsUpdate(v vocab.ActivityStreamsUpdate) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsUpdateMember: v,
		alias:                       this.alias,
		myIdx:                       this.Len(),
		parent:                      this,
	})
}

// AppendActivityStreamsVideo appends a Video value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsVideo(v vocab.ActivityStreamsVideo) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsVideoMember: v,
		alias:                      this.alias,
		myIdx:                      this.Len(),
		parent:                     this,
	})
}

// AppendActivityStreamsView appends a View value to the back of a list of the
// property "generator". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsGeneratorProperty) AppendActivityStreamsView(v vocab.ActivityStreamsView) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsViewMember: v,
		alias:                     this.alias,
		myIdx:                     this.Len(),
		parent:                    this,
	})
}

// AppendIRI appends an IRI value to the back of a list of the property "generator"
func (this *ActivityStreamsGeneratorProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &ActivityStreamsGeneratorPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  this.Len(),
		parent: this,
	})
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this ActivityStreamsGeneratorProperty) At(index int) vocab.ActivityStreamsGeneratorPropertyIterator {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this ActivityStreamsGeneratorProperty) Begin() vocab.ActivityStreamsGeneratorPropertyIterator {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this ActivityStreamsGeneratorProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this ActivityStreamsGeneratorProperty) End() vocab.ActivityStreamsGeneratorPropertyIterator {
	return nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsGeneratorProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	for _, elem := range this.properties {
		child := elem.JSONLDContext()
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		for k, v := range child {
			m[k] = v
		}
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API method specifically needed only for alternate implementations
// for go-fed. Applications should not use this method. Panics if the index is
// out of bounds.
func (this ActivityStreamsGeneratorProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "generator" property.
func (this ActivityStreamsGeneratorProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this ActivityStreamsGeneratorProperty) Less(i, j int) bool {
	idx1 := this.KindIndex(i)
	idx2 := this.KindIndex(j)
	if idx1 < idx2 {
		return true
	} else if idx1 == idx2 {
		if idx1 == 0 {
			lhs := this.properties[i].GetActivityStreamsObject()
			rhs := this.properties[j].GetActivityStreamsObject()
			return lhs.LessThan(rhs)
		} else if idx1 == 1 {
			lhs := this.properties[i].GetActivityStreamsLink()
			rhs := this.properties[j].GetActivityStreamsLink()
			return lhs.LessThan(rhs)
		} else if idx1 == 2 {
			lhs := this.properties[i].GetActivityStreamsAccept()
			rhs := this.properties[j].GetActivityStreamsAccept()
			return lhs.LessThan(rhs)
		} else if idx1 == 3 {
			lhs := this.properties[i].GetActivityStreamsActivity()
			rhs := this.properties[j].GetActivityStreamsActivity()
			return lhs.LessThan(rhs)
		} else if idx1 == 4 {
			lhs := this.properties[i].GetActivityStreamsAdd()
			rhs := this.properties[j].GetActivityStreamsAdd()
			return lhs.LessThan(rhs)
		} else if idx1 == 5 {
			lhs := this.properties[i].GetActivityStreamsAnnounce()
			rhs := this.properties[j].GetActivityStreamsAnnounce()
			return lhs.LessThan(rhs)
		} else if idx1 == 6 {
			lhs := this.properties[i].GetActivityStreamsApplication()
			rhs := this.properties[j].GetActivityStreamsApplication()
			return lhs.LessThan(rhs)
		} else if idx1 == 7 {
			lhs := this.properties[i].GetActivityStreamsArrive()
			rhs := this.properties[j].GetActivityStreamsArrive()
			return lhs.LessThan(rhs)
		} else if idx1 == 8 {
			lhs := this.properties[i].GetActivityStreamsArticle()
			rhs := this.properties[j].GetActivityStreamsArticle()
			return lhs.LessThan(rhs)
		} else if idx1 == 9 {
			lhs := this.properties[i].GetActivityStreamsAudio()
			rhs := this.properties[j].GetActivityStreamsAudio()
			return lhs.LessThan(rhs)
		} else if idx1 == 10 {
			lhs := this.properties[i].GetActivityStreamsBlock()
			rhs := this.properties[j].GetActivityStreamsBlock()
			return lhs.LessThan(rhs)
		} else if idx1 == 11 {
			lhs := this.properties[i].GetActivityStreamsCollection()
			rhs := this.properties[j].GetActivityStreamsCollection()
			return lhs.LessThan(rhs)
		} else if idx1 == 12 {
			lhs := this.properties[i].GetActivityStreamsCollectionPage()
			rhs := this.properties[j].GetActivityStreamsCollectionPage()
			return lhs.LessThan(rhs)
		} else if idx1 == 13 {
			lhs := this.properties[i].GetActivityStreamsCreate()
			rhs := this.properties[j].GetActivityStreamsCreate()
			return lhs.LessThan(rhs)
		} else if idx1 == 14 {
			lhs := this.properties[i].GetActivityStreamsDelete()
			rhs := this.properties[j].GetActivityStreamsDelete()
			return lhs.LessThan(rhs)
		} else if idx1 == 15 {
			lhs := this.properties[i].GetActivityStreamsDislike()
			rhs := this.properties[j].GetActivityStreamsDislike()
			return lhs.LessThan(rhs)
		} else if idx1 == 16 {
			lhs := this.properties[i].GetActivityStreamsDocument()
			rhs := this.properties[j].GetActivityStreamsDocument()
			return lhs.LessThan(rhs)
		} else if idx1 == 17 {
			lhs := this.properties[i].GetActivityStreamsEvent()
			rhs := this.properties[j].GetActivityStreamsEvent()
			return lhs.LessThan(rhs)
		} else if idx1 == 18 {
			lhs := this.properties[i].GetActivityStreamsFlag()
			rhs := this.properties[j].GetActivityStreamsFlag()
			return lhs.LessThan(rhs)
		} else if idx1 == 19 {
			lhs := this.properties[i].GetActivityStreamsFollow()
			rhs := this.properties[j].GetActivityStreamsFollow()
			return lhs.LessThan(rhs)
		} else if idx1 == 20 {
			lhs := this.properties[i].GetActivityStreamsGroup()
			rhs := this.properties[j].GetActivityStreamsGroup()
			return lhs.LessThan(rhs)
		} else if idx1 == 21 {
			lhs := this.properties[i].GetActivityStreamsIgnore()
			rhs := this.properties[j].GetActivityStreamsIgnore()
			return lhs.LessThan(rhs)
		} else if idx1 == 22 {
			lhs := this.properties[i].GetActivityStreamsImage()
			rhs := this.properties[j].GetActivityStreamsImage()
			return lhs.LessThan(rhs)
		} else if idx1 == 23 {
			lhs := this.properties[i].GetActivityStreamsIntransitiveActivity()
			rhs := this.properties[j].GetActivityStreamsIntransitiveActivity()
			return lhs.LessThan(rhs)
		} else if idx1 == 24 {
			lhs := this.properties[i].GetActivityStreamsInvite()
			rhs := this.properties[j].GetActivityStreamsInvite()
			return lhs.LessThan(rhs)
		} else if idx1 == 25 {
			lhs := this.properties[i].GetActivityStreamsJoin()
			rhs := this.properties[j].GetActivityStreamsJoin()
			return lhs.LessThan(rhs)
		} else if idx1 == 26 {
			lhs := this.properties[i].GetActivityStreamsLeave()
			rhs := this.properties[j].GetActivityStreamsLeave()
			return lhs.LessThan(rhs)
		} else if idx1 == 27 {
			lhs := this.properties[i].GetActivityStreamsLike()
			rhs := this.properties[j].GetActivityStreamsLike()
			return lhs.LessThan(rhs)
		} else if idx1 == 28 {
			lhs := this.properties[i].GetActivityStreamsListen()
			rhs := this.properties[j].GetActivityStreamsListen()
			return lhs.LessThan(rhs)
		} else if idx1 == 29 {
			lhs := this.properties[i].GetActivityStreamsMention()
			rhs := this.properties[j].GetActivityStreamsMention()
			return lhs.LessThan(rhs)
		} else if idx1 == 30 {
			lhs := this.properties[i].GetActivityStreamsMove()
			rhs := this.properties[j].GetActivityStreamsMove()
			return lhs.LessThan(rhs)
		} else if idx1 == 31 {
			lhs := this.properties[i].GetActivityStreamsNote()
			rhs := this.properties[j].GetActivityStreamsNote()
			return lhs.LessThan(rhs)
		} else if idx1 == 32 {
			lhs := this.properties[i].GetActivityStreamsOffer()
			rhs := this.properties[j].GetActivityStreamsOffer()
			return lhs.LessThan(rhs)
		} else if idx1 == 33 {
			lhs := this.properties[i].GetActivityStreamsOrderedCollection()
			rhs := this.properties[j].GetActivityStreamsOrderedCollection()
			return lhs.LessThan(rhs)
		} else if idx1 == 34 {
			lhs := this.properties[i].GetActivityStreamsOrderedCollectionPage()
			rhs := this.properties[j].GetActivityStreamsOrderedCollectionPage()
			return lhs.LessThan(rhs)
		} else if idx1 == 35 {
			lhs := this.properties[i].GetActivityStreamsOrganization()
			rhs := this.properties[j].GetActivityStreamsOrganization()
			return lhs.LessThan(rhs)
		} else if idx1 == 36 {
			lhs := this.properties[i].GetActivityStreamsPage()
			rhs := this.properties[j].GetActivityStreamsPage()
			return lhs.LessThan(rhs)
		} else if idx1 == 37 {
			lhs := this.properties[i].GetActivityStreamsPerson()
			rhs := this.properties[j].GetActivityStreamsPerson()
			return lhs.LessThan(rhs)
		} else if idx1 == 38 {
			lhs := this.properties[i].GetActivityStreamsPlace()
			rhs := this.properties[j].GetActivityStreamsPlace()
			return lhs.LessThan(rhs)
		} else if idx1 == 39 {
			lhs := this.properties[i].GetActivityStreamsProfile()
			rhs := this.properties[j].GetActivityStreamsProfile()
			return lhs.LessThan(rhs)
		} else if idx1 == 40 {
			lhs := this.properties[i].GetActivityStreamsQuestion()
			rhs := this.properties[j].GetActivityStreamsQuestion()
			return lhs.LessThan(rhs)
		} else if idx1 == 41 {
			lhs := this.properties[i].GetActivityStreamsRead()
			rhs := this.properties[j].GetActivityStreamsRead()
			return lhs.LessThan(rhs)
		} else if idx1 == 42 {
			lhs := this.properties[i].GetActivityStreamsReject()
			rhs := this.properties[j].GetActivityStreamsReject()
			return lhs.LessThan(rhs)
		} else if idx1 == 43 {
			lhs := this.properties[i].GetActivityStreamsRelationship()
			rhs := this.properties[j].GetActivityStreamsRelationship()
			return lhs.LessThan(rhs)
		} else if idx1 == 44 {
			lhs := this.properties[i].GetActivityStreamsRemove()
			rhs := this.properties[j].GetActivityStreamsRemove()
			return lhs.LessThan(rhs)
		} else if idx1 == 45 {
			lhs := this.properties[i].GetActivityStreamsService()
			rhs := this.properties[j].GetActivityStreamsService()
			return lhs.LessThan(rhs)
		} else if idx1 == 46 {
			lhs := this.properties[i].GetActivityStreamsTentativeAccept()
			rhs := this.properties[j].GetActivityStreamsTentativeAccept()
			return lhs.LessThan(rhs)
		} else if idx1 == 47 {
			lhs := this.properties[i].GetActivityStreamsTentativeReject()
			rhs := this.properties[j].GetActivityStreamsTentativeReject()
			return lhs.LessThan(rhs)
		} else if idx1 == 48 {
			lhs := this.properties[i].GetActivityStreamsTombstone()
			rhs := this.properties[j].GetActivityStreamsTombstone()
			return lhs.LessThan(rhs)
		} else if idx1 == 49 {
			lhs := this.properties[i].GetActivityStreamsTravel()
			rhs := this.properties[j].GetActivityStreamsTravel()
			return lhs.LessThan(rhs)
		} else if idx1 == 50 {
			lhs := this.properties[i].GetActivityStreamsUndo()
			rhs := this.properties[j].GetActivityStreamsUndo()
			return lhs.LessThan(rhs)
		} else if idx1 == 51 {
			lhs := this.properties[i].GetActivityStreamsUpdate()
			rhs := this.properties[j].GetActivityStreamsUpdate()
			return lhs.LessThan(rhs)
		} else if idx1 == 52 {
			lhs := this.properties[i].GetActivityStreamsVideo()
			rhs := this.properties[j].GetActivityStreamsVideo()
			return lhs.LessThan(rhs)
		} else if idx1 == 53 {
			lhs := this.properties[i].GetActivityStreamsView()
			rhs := this.properties[j].GetActivityStreamsView()
			return lhs.LessThan(rhs)
		} else if idx1 == -2 {
			lhs := this.properties[i].GetIRI()
			rhs := this.properties[j].GetIRI()
			return lhs.String() < rhs.String()
		}
	}
	return false
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this ActivityStreamsGeneratorProperty) LessThan(o vocab.ActivityStreamsGeneratorProperty) bool {
	l1 := this.Len()
	l2 := o.Len()
	l := l1
	if l2 < l1 {
		l = l2
	}
	for i := 0; i < l; i++ {
		if this.properties[i].LessThan(o.At(i)) {
			return true
		} else if o.At(i).LessThan(this.properties[i]) {
			return false
		}
	}
	return l1 < l2
}

// Name returns the name of this property: "generator".
func (this ActivityStreamsGeneratorProperty) Name() string {
	return "generator"
}

// PrependActivityStreamsAccept prepends a Accept value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsAccept(v vocab.ActivityStreamsAccept) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsAcceptMember: v,
		alias:                       this.alias,
		myIdx:                       0,
		parent:                      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsActivity prepends a Activity value to the front of a list
// of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsActivity(v vocab.ActivityStreamsActivity) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsActivityMember: v,
		alias:                         this.alias,
		myIdx:                         0,
		parent:                        this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsAdd prepends a Add value to the front of a list of the
// property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsAdd(v vocab.ActivityStreamsAdd) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsAddMember: v,
		alias:                    this.alias,
		myIdx:                    0,
		parent:                   this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsAnnounce prepends a Announce value to the front of a list
// of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsAnnounce(v vocab.ActivityStreamsAnnounce) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsAnnounceMember: v,
		alias:                         this.alias,
		myIdx:                         0,
		parent:                        this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsApplication prepends a Application value to the front of
// a list of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsApplication(v vocab.ActivityStreamsApplication) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsApplicationMember: v,
		alias:                            this.alias,
		myIdx:                            0,
		parent:                           this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsArrive prepends a Arrive value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsArrive(v vocab.ActivityStreamsArrive) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsArriveMember: v,
		alias:                       this.alias,
		myIdx:                       0,
		parent:                      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsArticle prepends a Article value to the front of a list
// of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsArticle(v vocab.ActivityStreamsArticle) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsArticleMember: v,
		alias:                        this.alias,
		myIdx:                        0,
		parent:                       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsAudio prepends a Audio value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsAudio(v vocab.ActivityStreamsAudio) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsAudioMember: v,
		alias:                      this.alias,
		myIdx:                      0,
		parent:                     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsBlock prepends a Block value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsBlock(v vocab.ActivityStreamsBlock) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsBlockMember: v,
		alias:                      this.alias,
		myIdx:                      0,
		parent:                     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsCollection prepends a Collection value to the front of a
// list of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsCollection(v vocab.ActivityStreamsCollection) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsCollectionMember: v,
		alias:                           this.alias,
		myIdx:                           0,
		parent:                          this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsCollectionPage prepends a CollectionPage value to the
// front of a list of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsCollectionPage(v vocab.ActivityStreamsCollectionPage) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsCollectionPageMember: v,
		alias:                               this.alias,
		myIdx:                               0,
		parent:                              this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsCreate prepends a Create value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsCreate(v vocab.ActivityStreamsCreate) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsCreateMember: v,
		alias:                       this.alias,
		myIdx:                       0,
		parent:                      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsDelete prepends a Delete value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsDelete(v vocab.ActivityStreamsDelete) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsDeleteMember: v,
		alias:                       this.alias,
		myIdx:                       0,
		parent:                      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsDislike prepends a Dislike value to the front of a list
// of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsDislike(v vocab.ActivityStreamsDislike) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsDislikeMember: v,
		alias:                        this.alias,
		myIdx:                        0,
		parent:                       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsDocument prepends a Document value to the front of a list
// of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsDocument(v vocab.ActivityStreamsDocument) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsDocumentMember: v,
		alias:                         this.alias,
		myIdx:                         0,
		parent:                        this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsEvent prepends a Event value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsEvent(v vocab.ActivityStreamsEvent) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsEventMember: v,
		alias:                      this.alias,
		myIdx:                      0,
		parent:                     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsFlag prepends a Flag value to the front of a list of the
// property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsFlag(v vocab.ActivityStreamsFlag) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsFlagMember: v,
		alias:                     this.alias,
		myIdx:                     0,
		parent:                    this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsFollow prepends a Follow value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsFollow(v vocab.ActivityStreamsFollow) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsFollowMember: v,
		alias:                       this.alias,
		myIdx:                       0,
		parent:                      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsGroup prepends a Group value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsGroup(v vocab.ActivityStreamsGroup) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsGroupMember: v,
		alias:                      this.alias,
		myIdx:                      0,
		parent:                     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsIgnore prepends a Ignore value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsIgnore(v vocab.ActivityStreamsIgnore) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsIgnoreMember: v,
		alias:                       this.alias,
		myIdx:                       0,
		parent:                      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsImage prepends a Image value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsImage(v vocab.ActivityStreamsImage) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsImageMember: v,
		alias:                      this.alias,
		myIdx:                      0,
		parent:                     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsIntransitiveActivity prepends a IntransitiveActivity
// value to the front of a list of the property "generator". Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsIntransitiveActivity(v vocab.ActivityStreamsIntransitiveActivity) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsIntransitiveActivityMember: v,
		alias:  this.alias,
		myIdx:  0,
		parent: this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsInvite prepends a Invite value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsInvite(v vocab.ActivityStreamsInvite) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsInviteMember: v,
		alias:                       this.alias,
		myIdx:                       0,
		parent:                      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsJoin prepends a Join value to the front of a list of the
// property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsJoin(v vocab.ActivityStreamsJoin) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsJoinMember: v,
		alias:                     this.alias,
		myIdx:                     0,
		parent:                    this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsLeave prepends a Leave value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsLeave(v vocab.ActivityStreamsLeave) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsLeaveMember: v,
		alias:                      this.alias,
		myIdx:                      0,
		parent:                     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsLike prepends a Like value to the front of a list of the
// property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsLike(v vocab.ActivityStreamsLike) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsLikeMember: v,
		alias:                     this.alias,
		myIdx:                     0,
		parent:                    this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsLink prepends a Link value to the front of a list of the
// property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsLink(v vocab.ActivityStreamsLink) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsLinkMember: v,
		alias:                     this.alias,
		myIdx:                     0,
		parent:                    this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsListen prepends a Listen value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsListen(v vocab.ActivityStreamsListen) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsListenMember: v,
		alias:                       this.alias,
		myIdx:                       0,
		parent:                      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsMention prepends a Mention value to the front of a list
// of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsMention(v vocab.ActivityStreamsMention) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsMentionMember: v,
		alias:                        this.alias,
		myIdx:                        0,
		parent:                       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsMove prepends a Move value to the front of a list of the
// property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsMove(v vocab.ActivityStreamsMove) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsMoveMember: v,
		alias:                     this.alias,
		myIdx:                     0,
		parent:                    this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsNote prepends a Note value to the front of a list of the
// property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsNote(v vocab.ActivityStreamsNote) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsNoteMember: v,
		alias:                     this.alias,
		myIdx:                     0,
		parent:                    this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsObject prepends a Object value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsObject(v vocab.ActivityStreamsObject) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsObjectMember: v,
		alias:                       this.alias,
		myIdx:                       0,
		parent:                      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsOffer prepends a Offer value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsOffer(v vocab.ActivityStreamsOffer) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsOfferMember: v,
		alias:                      this.alias,
		myIdx:                      0,
		parent:                     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsOrderedCollection prepends a OrderedCollection value to
// the front of a list of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsOrderedCollection(v vocab.ActivityStreamsOrderedCollection) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsOrderedCollectionMember: v,
		alias:                                  this.alias,
		myIdx:                                  0,
		parent:                                 this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsOrderedCollectionPage prepends a OrderedCollectionPage
// value to the front of a list of the property "generator". Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsOrderedCollectionPage(v vocab.ActivityStreamsOrderedCollectionPage) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsOrderedCollectionPageMember: v,
		alias:  this.alias,
		myIdx:  0,
		parent: this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsOrganization prepends a Organization value to the front
// of a list of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsOrganization(v vocab.ActivityStreamsOrganization) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsOrganizationMember: v,
		alias:                             this.alias,
		myIdx:                             0,
		parent:                            this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsPage prepends a Page value to the front of a list of the
// property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsPage(v vocab.ActivityStreamsPage) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsPageMember: v,
		alias:                     this.alias,
		myIdx:                     0,
		parent:                    this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsPerson prepends a Person value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsPerson(v vocab.ActivityStreamsPerson) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsPersonMember: v,
		alias:                       this.alias,
		myIdx:                       0,
		parent:                      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsPlace prepends a Place value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsPlace(v vocab.ActivityStreamsPlace) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsPlaceMember: v,
		alias:                      this.alias,
		myIdx:                      0,
		parent:                     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsProfile prepends a Profile value to the front of a list
// of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsProfile(v vocab.ActivityStreamsProfile) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsProfileMember: v,
		alias:                        this.alias,
		myIdx:                        0,
		parent:                       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsQuestion prepends a Question value to the front of a list
// of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsQuestion(v vocab.ActivityStreamsQuestion) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsQuestionMember: v,
		alias:                         this.alias,
		myIdx:                         0,
		parent:                        this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsRead prepends a Read value to the front of a list of the
// property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsRead(v vocab.ActivityStreamsRead) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsReadMember: v,
		alias:                     this.alias,
		myIdx:                     0,
		parent:                    this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsReject prepends a Reject value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsReject(v vocab.ActivityStreamsReject) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsRejectMember: v,
		alias:                       this.alias,
		myIdx:                       0,
		parent:                      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsRelationship prepends a Relationship value to the front
// of a list of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsRelationship(v vocab.ActivityStreamsRelationship) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsRelationshipMember: v,
		alias:                             this.alias,
		myIdx:                             0,
		parent:                            this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsRemove prepends a Remove value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsRemove(v vocab.ActivityStreamsRemove) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsRemoveMember: v,
		alias:                       this.alias,
		myIdx:                       0,
		parent:                      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsService prepends a Service value to the front of a list
// of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsService(v vocab.ActivityStreamsService) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsServiceMember: v,
		alias:                        this.alias,
		myIdx:                        0,
		parent:                       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsTentativeAccept prepends a TentativeAccept value to the
// front of a list of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsTentativeAccept(v vocab.ActivityStreamsTentativeAccept) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsTentativeAcceptMember: v,
		alias:                                this.alias,
		myIdx:                                0,
		parent:                               this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsTentativeReject prepends a TentativeReject value to the
// front of a list of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsTentativeReject(v vocab.ActivityStreamsTentativeReject) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsTentativeRejectMember: v,
		alias:                                this.alias,
		myIdx:                                0,
		parent:                               this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsTombstone prepends a Tombstone value to the front of a
// list of the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsTombstone(v vocab.ActivityStreamsTombstone) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsTombstoneMember: v,
		alias:                          this.alias,
		myIdx:                          0,
		parent:                         this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsTravel prepends a Travel value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsTravel(v vocab.ActivityStreamsTravel) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsTravelMember: v,
		alias:                       this.alias,
		myIdx:                       0,
		parent:                      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsUndo prepends a Undo value to the front of a list of the
// property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsUndo(v vocab.ActivityStreamsUndo) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsUndoMember: v,
		alias:                     this.alias,
		myIdx:                     0,
		parent:                    this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsUpdate prepends a Update value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsUpdate(v vocab.ActivityStreamsUpdate) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsUpdateMember: v,
		alias:                       this.alias,
		myIdx:                       0,
		parent:                      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsVideo prepends a Video value to the front of a list of
// the property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsVideo(v vocab.ActivityStreamsVideo) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsVideoMember: v,
		alias:                      this.alias,
		myIdx:                      0,
		parent:                     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsView prepends a View value to the front of a list of the
// property "generator". Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) PrependActivityStreamsView(v vocab.ActivityStreamsView) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		activitystreamsViewMember: v,
		alias:                     this.alias,
		myIdx:                     0,
		parent:                    this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependIRI prepends an IRI value to the front of a list of the property
// "generator".
func (this *ActivityStreamsGeneratorProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*ActivityStreamsGeneratorPropertyIterator{{
		alias:  this.alias,
		iri:    v,
		myIdx:  0,
		parent: this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Remove deletes an element at the specified index from a list of the property
// "generator", regardless of its type. Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &ActivityStreamsGeneratorPropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsGeneratorProperty) Serialize() (interface{}, error) {
	s := make([]interface{}, 0, len(this.properties))
	for _, iterator := range this.properties {
		if b, err := iterator.serialize(); err != nil {
			return s, err
		} else {
			s = append(s, b)
		}
	}
	// Shortcut: if serializing one value, don't return an array -- pretty sure other Fediverse software would choke on a "type" value with array, for example.
	if len(s) == 1 {
		return s[0], nil
	}
	return s, nil
}

// SetActivityStreamsAccept sets a Accept value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsAccept(idx int, v vocab.ActivityStreamsAccept) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsAcceptMember: v,
		alias:                       this.alias,
		myIdx:                       idx,
		parent:                      this,
	}
}

// SetActivityStreamsActivity sets a Activity value to be at the specified index
// for the property "generator". Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsActivity(idx int, v vocab.ActivityStreamsActivity) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsActivityMember: v,
		alias:                         this.alias,
		myIdx:                         idx,
		parent:                        this,
	}
}

// SetActivityStreamsAdd sets a Add value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsAdd(idx int, v vocab.ActivityStreamsAdd) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsAddMember: v,
		alias:                    this.alias,
		myIdx:                    idx,
		parent:                   this,
	}
}

// SetActivityStreamsAnnounce sets a Announce value to be at the specified index
// for the property "generator". Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsAnnounce(idx int, v vocab.ActivityStreamsAnnounce) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsAnnounceMember: v,
		alias:                         this.alias,
		myIdx:                         idx,
		parent:                        this,
	}
}

// SetActivityStreamsApplication sets a Application value to be at the specified
// index for the property "generator". Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsApplication(idx int, v vocab.ActivityStreamsApplication) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsApplicationMember: v,
		alias:                            this.alias,
		myIdx:                            idx,
		parent:                           this,
	}
}

// SetActivityStreamsArrive sets a Arrive value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsArrive(idx int, v vocab.ActivityStreamsArrive) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsArriveMember: v,
		alias:                       this.alias,
		myIdx:                       idx,
		parent:                      this,
	}
}

// SetActivityStreamsArticle sets a Article value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsArticle(idx int, v vocab.ActivityStreamsArticle) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsArticleMember: v,
		alias:                        this.alias,
		myIdx:                        idx,
		parent:                       this,
	}
}

// SetActivityStreamsAudio sets a Audio value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsAudio(idx int, v vocab.ActivityStreamsAudio) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsAudioMember: v,
		alias:                      this.alias,
		myIdx:                      idx,
		parent:                     this,
	}
}

// SetActivityStreamsBlock sets a Block value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsBlock(idx int, v vocab.ActivityStreamsBlock) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsBlockMember: v,
		alias:                      this.alias,
		myIdx:                      idx,
		parent:                     this,
	}
}

// SetActivityStreamsCollection sets a Collection value to be at the specified
// index for the property "generator". Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsCollection(idx int, v vocab.ActivityStreamsCollection) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsCollectionMember: v,
		alias:                           this.alias,
		myIdx:                           idx,
		parent:                          this,
	}
}

// SetActivityStreamsCollectionPage sets a CollectionPage value to be at the
// specified index for the property "generator". Panics if the index is out of
// bounds. Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsCollectionPage(idx int, v vocab.ActivityStreamsCollectionPage) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsCollectionPageMember: v,
		alias:                               this.alias,
		myIdx:                               idx,
		parent:                              this,
	}
}

// SetActivityStreamsCreate sets a Create value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsCreate(idx int, v vocab.ActivityStreamsCreate) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsCreateMember: v,
		alias:                       this.alias,
		myIdx:                       idx,
		parent:                      this,
	}
}

// SetActivityStreamsDelete sets a Delete value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsDelete(idx int, v vocab.ActivityStreamsDelete) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsDeleteMember: v,
		alias:                       this.alias,
		myIdx:                       idx,
		parent:                      this,
	}
}

// SetActivityStreamsDislike sets a Dislike value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsDislike(idx int, v vocab.ActivityStreamsDislike) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsDislikeMember: v,
		alias:                        this.alias,
		myIdx:                        idx,
		parent:                       this,
	}
}

// SetActivityStreamsDocument sets a Document value to be at the specified index
// for the property "generator". Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsDocument(idx int, v vocab.ActivityStreamsDocument) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsDocumentMember: v,
		alias:                         this.alias,
		myIdx:                         idx,
		parent:                        this,
	}
}

// SetActivityStreamsEvent sets a Event value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsEvent(idx int, v vocab.ActivityStreamsEvent) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsEventMember: v,
		alias:                      this.alias,
		myIdx:                      idx,
		parent:                     this,
	}
}

// SetActivityStreamsFlag sets a Flag value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsFlag(idx int, v vocab.ActivityStreamsFlag) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsFlagMember: v,
		alias:                     this.alias,
		myIdx:                     idx,
		parent:                    this,
	}
}

// SetActivityStreamsFollow sets a Follow value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsFollow(idx int, v vocab.ActivityStreamsFollow) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsFollowMember: v,
		alias:                       this.alias,
		myIdx:                       idx,
		parent:                      this,
	}
}

// SetActivityStreamsGroup sets a Group value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsGroup(idx int, v vocab.ActivityStreamsGroup) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsGroupMember: v,
		alias:                      this.alias,
		myIdx:                      idx,
		parent:                     this,
	}
}

// SetActivityStreamsIgnore sets a Ignore value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsIgnore(idx int, v vocab.ActivityStreamsIgnore) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsIgnoreMember: v,
		alias:                       this.alias,
		myIdx:                       idx,
		parent:                      this,
	}
}

// SetActivityStreamsImage sets a Image value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsImage(idx int, v vocab.ActivityStreamsImage) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsImageMember: v,
		alias:                      this.alias,
		myIdx:                      idx,
		parent:                     this,
	}
}

// SetActivityStreamsIntransitiveActivity sets a IntransitiveActivity value to be
// at the specified index for the property "generator". Panics if the index is
// out of bounds. Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsIntransitiveActivity(idx int, v vocab.ActivityStreamsIntransitiveActivity) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsIntransitiveActivityMember: v,
		alias:  this.alias,
		myIdx:  idx,
		parent: this,
	}
}

// SetActivityStreamsInvite sets a Invite value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsInvite(idx int, v vocab.ActivityStreamsInvite) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsInviteMember: v,
		alias:                       this.alias,
		myIdx:                       idx,
		parent:                      this,
	}
}

// SetActivityStreamsJoin sets a Join value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsJoin(idx int, v vocab.ActivityStreamsJoin) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsJoinMember: v,
		alias:                     this.alias,
		myIdx:                     idx,
		parent:                    this,
	}
}

// SetActivityStreamsLeave sets a Leave value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsLeave(idx int, v vocab.ActivityStreamsLeave) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsLeaveMember: v,
		alias:                      this.alias,
		myIdx:                      idx,
		parent:                     this,
	}
}

// SetActivityStreamsLike sets a Like value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsLike(idx int, v vocab.ActivityStreamsLike) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsLikeMember: v,
		alias:                     this.alias,
		myIdx:                     idx,
		parent:                    this,
	}
}

// SetActivityStreamsLink sets a Link value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsLink(idx int, v vocab.ActivityStreamsLink) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsLinkMember: v,
		alias:                     this.alias,
		myIdx:                     idx,
		parent:                    this,
	}
}

// SetActivityStreamsListen sets a Listen value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsListen(idx int, v vocab.ActivityStreamsListen) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsListenMember: v,
		alias:                       this.alias,
		myIdx:                       idx,
		parent:                      this,
	}
}

// SetActivityStreamsMention sets a Mention value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsMention(idx int, v vocab.ActivityStreamsMention) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsMentionMember: v,
		alias:                        this.alias,
		myIdx:                        idx,
		parent:                       this,
	}
}

// SetActivityStreamsMove sets a Move value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsMove(idx int, v vocab.ActivityStreamsMove) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsMoveMember: v,
		alias:                     this.alias,
		myIdx:                     idx,
		parent:                    this,
	}
}

// SetActivityStreamsNote sets a Note value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsNote(idx int, v vocab.ActivityStreamsNote) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsNoteMember: v,
		alias:                     this.alias,
		myIdx:                     idx,
		parent:                    this,
	}
}

// SetActivityStreamsObject sets a Object value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsObject(idx int, v vocab.ActivityStreamsObject) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsObjectMember: v,
		alias:                       this.alias,
		myIdx:                       idx,
		parent:                      this,
	}
}

// SetActivityStreamsOffer sets a Offer value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsOffer(idx int, v vocab.ActivityStreamsOffer) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsOfferMember: v,
		alias:                      this.alias,
		myIdx:                      idx,
		parent:                     this,
	}
}

// SetActivityStreamsOrderedCollection sets a OrderedCollection value to be at the
// specified index for the property "generator". Panics if the index is out of
// bounds. Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsOrderedCollection(idx int, v vocab.ActivityStreamsOrderedCollection) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsOrderedCollectionMember: v,
		alias:                                  this.alias,
		myIdx:                                  idx,
		parent:                                 this,
	}
}

// SetActivityStreamsOrderedCollectionPage sets a OrderedCollectionPage value to
// be at the specified index for the property "generator". Panics if the index
// is out of bounds. Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsOrderedCollectionPage(idx int, v vocab.ActivityStreamsOrderedCollectionPage) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsOrderedCollectionPageMember: v,
		alias:  this.alias,
		myIdx:  idx,
		parent: this,
	}
}

// SetActivityStreamsOrganization sets a Organization value to be at the specified
// index for the property "generator". Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsOrganization(idx int, v vocab.ActivityStreamsOrganization) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsOrganizationMember: v,
		alias:                             this.alias,
		myIdx:                             idx,
		parent:                            this,
	}
}

// SetActivityStreamsPage sets a Page value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsPage(idx int, v vocab.ActivityStreamsPage) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsPageMember: v,
		alias:                     this.alias,
		myIdx:                     idx,
		parent:                    this,
	}
}

// SetActivityStreamsPerson sets a Person value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsPerson(idx int, v vocab.ActivityStreamsPerson) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsPersonMember: v,
		alias:                       this.alias,
		myIdx:                       idx,
		parent:                      this,
	}
}

// SetActivityStreamsPlace sets a Place value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsPlace(idx int, v vocab.ActivityStreamsPlace) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsPlaceMember: v,
		alias:                      this.alias,
		myIdx:                      idx,
		parent:                     this,
	}
}

// SetActivityStreamsProfile sets a Profile value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsProfile(idx int, v vocab.ActivityStreamsProfile) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsProfileMember: v,
		alias:                        this.alias,
		myIdx:                        idx,
		parent:                       this,
	}
}

// SetActivityStreamsQuestion sets a Question value to be at the specified index
// for the property "generator". Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsQuestion(idx int, v vocab.ActivityStreamsQuestion) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsQuestionMember: v,
		alias:                         this.alias,
		myIdx:                         idx,
		parent:                        this,
	}
}

// SetActivityStreamsRead sets a Read value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsRead(idx int, v vocab.ActivityStreamsRead) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsReadMember: v,
		alias:                     this.alias,
		myIdx:                     idx,
		parent:                    this,
	}
}

// SetActivityStreamsReject sets a Reject value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsReject(idx int, v vocab.ActivityStreamsReject) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsRejectMember: v,
		alias:                       this.alias,
		myIdx:                       idx,
		parent:                      this,
	}
}

// SetActivityStreamsRelationship sets a Relationship value to be at the specified
// index for the property "generator". Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsRelationship(idx int, v vocab.ActivityStreamsRelationship) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsRelationshipMember: v,
		alias:                             this.alias,
		myIdx:                             idx,
		parent:                            this,
	}
}

// SetActivityStreamsRemove sets a Remove value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsRemove(idx int, v vocab.ActivityStreamsRemove) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsRemoveMember: v,
		alias:                       this.alias,
		myIdx:                       idx,
		parent:                      this,
	}
}

// SetActivityStreamsService sets a Service value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsService(idx int, v vocab.ActivityStreamsService) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsServiceMember: v,
		alias:                        this.alias,
		myIdx:                        idx,
		parent:                       this,
	}
}

// SetActivityStreamsTentativeAccept sets a TentativeAccept value to be at the
// specified index for the property "generator". Panics if the index is out of
// bounds. Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsTentativeAccept(idx int, v vocab.ActivityStreamsTentativeAccept) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsTentativeAcceptMember: v,
		alias:                                this.alias,
		myIdx:                                idx,
		parent:                               this,
	}
}

// SetActivityStreamsTentativeReject sets a TentativeReject value to be at the
// specified index for the property "generator". Panics if the index is out of
// bounds. Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsTentativeReject(idx int, v vocab.ActivityStreamsTentativeReject) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsTentativeRejectMember: v,
		alias:                                this.alias,
		myIdx:                                idx,
		parent:                               this,
	}
}

// SetActivityStreamsTombstone sets a Tombstone value to be at the specified index
// for the property "generator". Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsTombstone(idx int, v vocab.ActivityStreamsTombstone) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsTombstoneMember: v,
		alias:                          this.alias,
		myIdx:                          idx,
		parent:                         this,
	}
}

// SetActivityStreamsTravel sets a Travel value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsTravel(idx int, v vocab.ActivityStreamsTravel) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsTravelMember: v,
		alias:                       this.alias,
		myIdx:                       idx,
		parent:                      this,
	}
}

// SetActivityStreamsUndo sets a Undo value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsUndo(idx int, v vocab.ActivityStreamsUndo) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsUndoMember: v,
		alias:                     this.alias,
		myIdx:                     idx,
		parent:                    this,
	}
}

// SetActivityStreamsUpdate sets a Update value to be at the specified index for
// the property "generator". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsUpdate(idx int, v vocab.ActivityStreamsUpdate) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsUpdateMember: v,
		alias:                       this.alias,
		myIdx:                       idx,
		parent:                      this,
	}
}

// SetActivityStreamsVideo sets a Video value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsVideo(idx int, v vocab.ActivityStreamsVideo) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsVideoMember: v,
		alias:                      this.alias,
		myIdx:                      idx,
		parent:                     this,
	}
}

// SetActivityStreamsView sets a View value to be at the specified index for the
// property "generator". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsGeneratorProperty) SetActivityStreamsView(idx int, v vocab.ActivityStreamsView) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		activitystreamsViewMember: v,
		alias:                     this.alias,
		myIdx:                     idx,
		parent:                    this,
	}
}

// SetIRI sets an IRI value to be at the specified index for the property
// "generator". Panics if the index is out of bounds.
func (this *ActivityStreamsGeneratorProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsGeneratorPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
}

// Swap swaps the location of values at two indices for the "generator" property.
func (this ActivityStreamsGeneratorProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}
