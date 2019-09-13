package streams

import (
	typeaccept "github.com/go-fed/activity/streams/impl/activitystreams/type_accept"
	typeactivity "github.com/go-fed/activity/streams/impl/activitystreams/type_activity"
	typeadd "github.com/go-fed/activity/streams/impl/activitystreams/type_add"
	typeannounce "github.com/go-fed/activity/streams/impl/activitystreams/type_announce"
	typeapplication "github.com/go-fed/activity/streams/impl/activitystreams/type_application"
	typearrive "github.com/go-fed/activity/streams/impl/activitystreams/type_arrive"
	typearticle "github.com/go-fed/activity/streams/impl/activitystreams/type_article"
	typeaudio "github.com/go-fed/activity/streams/impl/activitystreams/type_audio"
	typeblock "github.com/go-fed/activity/streams/impl/activitystreams/type_block"
	typecollection "github.com/go-fed/activity/streams/impl/activitystreams/type_collection"
	typecollectionpage "github.com/go-fed/activity/streams/impl/activitystreams/type_collectionpage"
	typecreate "github.com/go-fed/activity/streams/impl/activitystreams/type_create"
	typedelete "github.com/go-fed/activity/streams/impl/activitystreams/type_delete"
	typedislike "github.com/go-fed/activity/streams/impl/activitystreams/type_dislike"
	typedocument "github.com/go-fed/activity/streams/impl/activitystreams/type_document"
	typeevent "github.com/go-fed/activity/streams/impl/activitystreams/type_event"
	typeflag "github.com/go-fed/activity/streams/impl/activitystreams/type_flag"
	typefollow "github.com/go-fed/activity/streams/impl/activitystreams/type_follow"
	typegroup "github.com/go-fed/activity/streams/impl/activitystreams/type_group"
	typeignore "github.com/go-fed/activity/streams/impl/activitystreams/type_ignore"
	typeimage "github.com/go-fed/activity/streams/impl/activitystreams/type_image"
	typeintransitiveactivity "github.com/go-fed/activity/streams/impl/activitystreams/type_intransitiveactivity"
	typeinvite "github.com/go-fed/activity/streams/impl/activitystreams/type_invite"
	typejoin "github.com/go-fed/activity/streams/impl/activitystreams/type_join"
	typeleave "github.com/go-fed/activity/streams/impl/activitystreams/type_leave"
	typelike "github.com/go-fed/activity/streams/impl/activitystreams/type_like"
	typelink "github.com/go-fed/activity/streams/impl/activitystreams/type_link"
	typelisten "github.com/go-fed/activity/streams/impl/activitystreams/type_listen"
	typemention "github.com/go-fed/activity/streams/impl/activitystreams/type_mention"
	typemove "github.com/go-fed/activity/streams/impl/activitystreams/type_move"
	typenote "github.com/go-fed/activity/streams/impl/activitystreams/type_note"
	typeobject "github.com/go-fed/activity/streams/impl/activitystreams/type_object"
	typeoffer "github.com/go-fed/activity/streams/impl/activitystreams/type_offer"
	typeorderedcollection "github.com/go-fed/activity/streams/impl/activitystreams/type_orderedcollection"
	typeorderedcollectionpage "github.com/go-fed/activity/streams/impl/activitystreams/type_orderedcollectionpage"
	typeorganization "github.com/go-fed/activity/streams/impl/activitystreams/type_organization"
	typepage "github.com/go-fed/activity/streams/impl/activitystreams/type_page"
	typeperson "github.com/go-fed/activity/streams/impl/activitystreams/type_person"
	typeplace "github.com/go-fed/activity/streams/impl/activitystreams/type_place"
	typeprofile "github.com/go-fed/activity/streams/impl/activitystreams/type_profile"
	typepublickey "github.com/go-fed/activity/streams/impl/activitystreams/type_publickey"
	typequestion "github.com/go-fed/activity/streams/impl/activitystreams/type_question"
	typeread "github.com/go-fed/activity/streams/impl/activitystreams/type_read"
	typereject "github.com/go-fed/activity/streams/impl/activitystreams/type_reject"
	typerelationship "github.com/go-fed/activity/streams/impl/activitystreams/type_relationship"
	typeremove "github.com/go-fed/activity/streams/impl/activitystreams/type_remove"
	typeservice "github.com/go-fed/activity/streams/impl/activitystreams/type_service"
	typetentativeaccept "github.com/go-fed/activity/streams/impl/activitystreams/type_tentativeaccept"
	typetentativereject "github.com/go-fed/activity/streams/impl/activitystreams/type_tentativereject"
	typetombstone "github.com/go-fed/activity/streams/impl/activitystreams/type_tombstone"
	typetravel "github.com/go-fed/activity/streams/impl/activitystreams/type_travel"
	typeundo "github.com/go-fed/activity/streams/impl/activitystreams/type_undo"
	typeupdate "github.com/go-fed/activity/streams/impl/activitystreams/type_update"
	typevideo "github.com/go-fed/activity/streams/impl/activitystreams/type_video"
	typeview "github.com/go-fed/activity/streams/impl/activitystreams/type_view"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// ActivityStreamsActivityStreamsAcceptExtends returns true if Accept extends from
// the other's type.
func ActivityStreamsActivityStreamsAcceptExtends(other vocab.Type) bool {
	return typeaccept.ActivityStreamsAcceptExtends(other)
}

// ActivityStreamsActivityStreamsActivityExtends returns true if Activity extends
// from the other's type.
func ActivityStreamsActivityStreamsActivityExtends(other vocab.Type) bool {
	return typeactivity.ActivityStreamsActivityExtends(other)
}

// ActivityStreamsActivityStreamsAddExtends returns true if Add extends from the
// other's type.
func ActivityStreamsActivityStreamsAddExtends(other vocab.Type) bool {
	return typeadd.ActivityStreamsAddExtends(other)
}

// ActivityStreamsActivityStreamsAnnounceExtends returns true if Announce extends
// from the other's type.
func ActivityStreamsActivityStreamsAnnounceExtends(other vocab.Type) bool {
	return typeannounce.ActivityStreamsAnnounceExtends(other)
}

// ActivityStreamsActivityStreamsApplicationExtends returns true if Application
// extends from the other's type.
func ActivityStreamsActivityStreamsApplicationExtends(other vocab.Type) bool {
	return typeapplication.ActivityStreamsApplicationExtends(other)
}

// ActivityStreamsActivityStreamsArriveExtends returns true if Arrive extends from
// the other's type.
func ActivityStreamsActivityStreamsArriveExtends(other vocab.Type) bool {
	return typearrive.ActivityStreamsArriveExtends(other)
}

// ActivityStreamsActivityStreamsArticleExtends returns true if Article extends
// from the other's type.
func ActivityStreamsActivityStreamsArticleExtends(other vocab.Type) bool {
	return typearticle.ActivityStreamsArticleExtends(other)
}

// ActivityStreamsActivityStreamsAudioExtends returns true if Audio extends from
// the other's type.
func ActivityStreamsActivityStreamsAudioExtends(other vocab.Type) bool {
	return typeaudio.ActivityStreamsAudioExtends(other)
}

// ActivityStreamsActivityStreamsBlockExtends returns true if Block extends from
// the other's type.
func ActivityStreamsActivityStreamsBlockExtends(other vocab.Type) bool {
	return typeblock.ActivityStreamsBlockExtends(other)
}

// ActivityStreamsActivityStreamsCollectionExtends returns true if Collection
// extends from the other's type.
func ActivityStreamsActivityStreamsCollectionExtends(other vocab.Type) bool {
	return typecollection.ActivityStreamsCollectionExtends(other)
}

// ActivityStreamsActivityStreamsCollectionPageExtends returns true if
// CollectionPage extends from the other's type.
func ActivityStreamsActivityStreamsCollectionPageExtends(other vocab.Type) bool {
	return typecollectionpage.ActivityStreamsCollectionPageExtends(other)
}

// ActivityStreamsActivityStreamsCreateExtends returns true if Create extends from
// the other's type.
func ActivityStreamsActivityStreamsCreateExtends(other vocab.Type) bool {
	return typecreate.ActivityStreamsCreateExtends(other)
}

// ActivityStreamsActivityStreamsDeleteExtends returns true if Delete extends from
// the other's type.
func ActivityStreamsActivityStreamsDeleteExtends(other vocab.Type) bool {
	return typedelete.ActivityStreamsDeleteExtends(other)
}

// ActivityStreamsActivityStreamsDislikeExtends returns true if Dislike extends
// from the other's type.
func ActivityStreamsActivityStreamsDislikeExtends(other vocab.Type) bool {
	return typedislike.ActivityStreamsDislikeExtends(other)
}

// ActivityStreamsActivityStreamsDocumentExtends returns true if Document extends
// from the other's type.
func ActivityStreamsActivityStreamsDocumentExtends(other vocab.Type) bool {
	return typedocument.ActivityStreamsDocumentExtends(other)
}

// ActivityStreamsActivityStreamsEventExtends returns true if Event extends from
// the other's type.
func ActivityStreamsActivityStreamsEventExtends(other vocab.Type) bool {
	return typeevent.ActivityStreamsEventExtends(other)
}

// ActivityStreamsActivityStreamsFlagExtends returns true if Flag extends from the
// other's type.
func ActivityStreamsActivityStreamsFlagExtends(other vocab.Type) bool {
	return typeflag.ActivityStreamsFlagExtends(other)
}

// ActivityStreamsActivityStreamsFollowExtends returns true if Follow extends from
// the other's type.
func ActivityStreamsActivityStreamsFollowExtends(other vocab.Type) bool {
	return typefollow.ActivityStreamsFollowExtends(other)
}

// ActivityStreamsActivityStreamsGroupExtends returns true if Group extends from
// the other's type.
func ActivityStreamsActivityStreamsGroupExtends(other vocab.Type) bool {
	return typegroup.ActivityStreamsGroupExtends(other)
}

// ActivityStreamsActivityStreamsIgnoreExtends returns true if Ignore extends from
// the other's type.
func ActivityStreamsActivityStreamsIgnoreExtends(other vocab.Type) bool {
	return typeignore.ActivityStreamsIgnoreExtends(other)
}

// ActivityStreamsActivityStreamsImageExtends returns true if Image extends from
// the other's type.
func ActivityStreamsActivityStreamsImageExtends(other vocab.Type) bool {
	return typeimage.ActivityStreamsImageExtends(other)
}

// ActivityStreamsActivityStreamsIntransitiveActivityExtends returns true if
// IntransitiveActivity extends from the other's type.
func ActivityStreamsActivityStreamsIntransitiveActivityExtends(other vocab.Type) bool {
	return typeintransitiveactivity.ActivityStreamsIntransitiveActivityExtends(other)
}

// ActivityStreamsActivityStreamsInviteExtends returns true if Invite extends from
// the other's type.
func ActivityStreamsActivityStreamsInviteExtends(other vocab.Type) bool {
	return typeinvite.ActivityStreamsInviteExtends(other)
}

// ActivityStreamsActivityStreamsJoinExtends returns true if Join extends from the
// other's type.
func ActivityStreamsActivityStreamsJoinExtends(other vocab.Type) bool {
	return typejoin.ActivityStreamsJoinExtends(other)
}

// ActivityStreamsActivityStreamsLeaveExtends returns true if Leave extends from
// the other's type.
func ActivityStreamsActivityStreamsLeaveExtends(other vocab.Type) bool {
	return typeleave.ActivityStreamsLeaveExtends(other)
}

// ActivityStreamsActivityStreamsLikeExtends returns true if Like extends from the
// other's type.
func ActivityStreamsActivityStreamsLikeExtends(other vocab.Type) bool {
	return typelike.ActivityStreamsLikeExtends(other)
}

// ActivityStreamsActivityStreamsLinkExtends returns true if Link extends from the
// other's type.
func ActivityStreamsActivityStreamsLinkExtends(other vocab.Type) bool {
	return typelink.ActivityStreamsLinkExtends(other)
}

// ActivityStreamsActivityStreamsListenExtends returns true if Listen extends from
// the other's type.
func ActivityStreamsActivityStreamsListenExtends(other vocab.Type) bool {
	return typelisten.ActivityStreamsListenExtends(other)
}

// ActivityStreamsActivityStreamsMentionExtends returns true if Mention extends
// from the other's type.
func ActivityStreamsActivityStreamsMentionExtends(other vocab.Type) bool {
	return typemention.ActivityStreamsMentionExtends(other)
}

// ActivityStreamsActivityStreamsMoveExtends returns true if Move extends from the
// other's type.
func ActivityStreamsActivityStreamsMoveExtends(other vocab.Type) bool {
	return typemove.ActivityStreamsMoveExtends(other)
}

// ActivityStreamsActivityStreamsNoteExtends returns true if Note extends from the
// other's type.
func ActivityStreamsActivityStreamsNoteExtends(other vocab.Type) bool {
	return typenote.ActivityStreamsNoteExtends(other)
}

// ActivityStreamsActivityStreamsObjectExtends returns true if Object extends from
// the other's type.
func ActivityStreamsActivityStreamsObjectExtends(other vocab.Type) bool {
	return typeobject.ActivityStreamsObjectExtends(other)
}

// ActivityStreamsActivityStreamsOfferExtends returns true if Offer extends from
// the other's type.
func ActivityStreamsActivityStreamsOfferExtends(other vocab.Type) bool {
	return typeoffer.ActivityStreamsOfferExtends(other)
}

// ActivityStreamsActivityStreamsOrderedCollectionExtends returns true if
// OrderedCollection extends from the other's type.
func ActivityStreamsActivityStreamsOrderedCollectionExtends(other vocab.Type) bool {
	return typeorderedcollection.ActivityStreamsOrderedCollectionExtends(other)
}

// ActivityStreamsActivityStreamsOrderedCollectionPageExtends returns true if
// OrderedCollectionPage extends from the other's type.
func ActivityStreamsActivityStreamsOrderedCollectionPageExtends(other vocab.Type) bool {
	return typeorderedcollectionpage.ActivityStreamsOrderedCollectionPageExtends(other)
}

// ActivityStreamsActivityStreamsOrganizationExtends returns true if Organization
// extends from the other's type.
func ActivityStreamsActivityStreamsOrganizationExtends(other vocab.Type) bool {
	return typeorganization.ActivityStreamsOrganizationExtends(other)
}

// ActivityStreamsActivityStreamsPageExtends returns true if Page extends from the
// other's type.
func ActivityStreamsActivityStreamsPageExtends(other vocab.Type) bool {
	return typepage.ActivityStreamsPageExtends(other)
}

// ActivityStreamsActivityStreamsPersonExtends returns true if Person extends from
// the other's type.
func ActivityStreamsActivityStreamsPersonExtends(other vocab.Type) bool {
	return typeperson.ActivityStreamsPersonExtends(other)
}

// ActivityStreamsActivityStreamsPlaceExtends returns true if Place extends from
// the other's type.
func ActivityStreamsActivityStreamsPlaceExtends(other vocab.Type) bool {
	return typeplace.ActivityStreamsPlaceExtends(other)
}

// ActivityStreamsActivityStreamsProfileExtends returns true if Profile extends
// from the other's type.
func ActivityStreamsActivityStreamsProfileExtends(other vocab.Type) bool {
	return typeprofile.ActivityStreamsProfileExtends(other)
}

// ActivityStreamsActivityStreamsPublicKeyExtends returns true if PublicKey
// extends from the other's type.
func ActivityStreamsActivityStreamsPublicKeyExtends(other vocab.Type) bool {
	return typepublickey.ActivityStreamsPublicKeyExtends(other)
}

// ActivityStreamsActivityStreamsQuestionExtends returns true if Question extends
// from the other's type.
func ActivityStreamsActivityStreamsQuestionExtends(other vocab.Type) bool {
	return typequestion.ActivityStreamsQuestionExtends(other)
}

// ActivityStreamsActivityStreamsReadExtends returns true if Read extends from the
// other's type.
func ActivityStreamsActivityStreamsReadExtends(other vocab.Type) bool {
	return typeread.ActivityStreamsReadExtends(other)
}

// ActivityStreamsActivityStreamsRejectExtends returns true if Reject extends from
// the other's type.
func ActivityStreamsActivityStreamsRejectExtends(other vocab.Type) bool {
	return typereject.ActivityStreamsRejectExtends(other)
}

// ActivityStreamsActivityStreamsRelationshipExtends returns true if Relationship
// extends from the other's type.
func ActivityStreamsActivityStreamsRelationshipExtends(other vocab.Type) bool {
	return typerelationship.ActivityStreamsRelationshipExtends(other)
}

// ActivityStreamsActivityStreamsRemoveExtends returns true if Remove extends from
// the other's type.
func ActivityStreamsActivityStreamsRemoveExtends(other vocab.Type) bool {
	return typeremove.ActivityStreamsRemoveExtends(other)
}

// ActivityStreamsActivityStreamsServiceExtends returns true if Service extends
// from the other's type.
func ActivityStreamsActivityStreamsServiceExtends(other vocab.Type) bool {
	return typeservice.ActivityStreamsServiceExtends(other)
}

// ActivityStreamsActivityStreamsTentativeAcceptExtends returns true if
// TentativeAccept extends from the other's type.
func ActivityStreamsActivityStreamsTentativeAcceptExtends(other vocab.Type) bool {
	return typetentativeaccept.ActivityStreamsTentativeAcceptExtends(other)
}

// ActivityStreamsActivityStreamsTentativeRejectExtends returns true if
// TentativeReject extends from the other's type.
func ActivityStreamsActivityStreamsTentativeRejectExtends(other vocab.Type) bool {
	return typetentativereject.ActivityStreamsTentativeRejectExtends(other)
}

// ActivityStreamsActivityStreamsTombstoneExtends returns true if Tombstone
// extends from the other's type.
func ActivityStreamsActivityStreamsTombstoneExtends(other vocab.Type) bool {
	return typetombstone.ActivityStreamsTombstoneExtends(other)
}

// ActivityStreamsActivityStreamsTravelExtends returns true if Travel extends from
// the other's type.
func ActivityStreamsActivityStreamsTravelExtends(other vocab.Type) bool {
	return typetravel.ActivityStreamsTravelExtends(other)
}

// ActivityStreamsActivityStreamsUndoExtends returns true if Undo extends from the
// other's type.
func ActivityStreamsActivityStreamsUndoExtends(other vocab.Type) bool {
	return typeundo.ActivityStreamsUndoExtends(other)
}

// ActivityStreamsActivityStreamsUpdateExtends returns true if Update extends from
// the other's type.
func ActivityStreamsActivityStreamsUpdateExtends(other vocab.Type) bool {
	return typeupdate.ActivityStreamsUpdateExtends(other)
}

// ActivityStreamsActivityStreamsVideoExtends returns true if Video extends from
// the other's type.
func ActivityStreamsActivityStreamsVideoExtends(other vocab.Type) bool {
	return typevideo.ActivityStreamsVideoExtends(other)
}

// ActivityStreamsActivityStreamsViewExtends returns true if View extends from the
// other's type.
func ActivityStreamsActivityStreamsViewExtends(other vocab.Type) bool {
	return typeview.ActivityStreamsViewExtends(other)
}
