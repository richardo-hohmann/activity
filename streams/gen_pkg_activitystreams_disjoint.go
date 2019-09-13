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

// ActivityStreamsAcceptIsDisjointWith returns true if Accept is disjoint with the
// other's type.
func ActivityStreamsAcceptIsDisjointWith(other vocab.Type) bool {
	return typeaccept.AcceptIsDisjointWith(other)
}

// ActivityStreamsActivityIsDisjointWith returns true if Activity is disjoint with
// the other's type.
func ActivityStreamsActivityIsDisjointWith(other vocab.Type) bool {
	return typeactivity.ActivityIsDisjointWith(other)
}

// ActivityStreamsAddIsDisjointWith returns true if Add is disjoint with the
// other's type.
func ActivityStreamsAddIsDisjointWith(other vocab.Type) bool {
	return typeadd.AddIsDisjointWith(other)
}

// ActivityStreamsAnnounceIsDisjointWith returns true if Announce is disjoint with
// the other's type.
func ActivityStreamsAnnounceIsDisjointWith(other vocab.Type) bool {
	return typeannounce.AnnounceIsDisjointWith(other)
}

// ActivityStreamsApplicationIsDisjointWith returns true if Application is
// disjoint with the other's type.
func ActivityStreamsApplicationIsDisjointWith(other vocab.Type) bool {
	return typeapplication.ApplicationIsDisjointWith(other)
}

// ActivityStreamsArriveIsDisjointWith returns true if Arrive is disjoint with the
// other's type.
func ActivityStreamsArriveIsDisjointWith(other vocab.Type) bool {
	return typearrive.ArriveIsDisjointWith(other)
}

// ActivityStreamsArticleIsDisjointWith returns true if Article is disjoint with
// the other's type.
func ActivityStreamsArticleIsDisjointWith(other vocab.Type) bool {
	return typearticle.ArticleIsDisjointWith(other)
}

// ActivityStreamsAudioIsDisjointWith returns true if Audio is disjoint with the
// other's type.
func ActivityStreamsAudioIsDisjointWith(other vocab.Type) bool {
	return typeaudio.AudioIsDisjointWith(other)
}

// ActivityStreamsBlockIsDisjointWith returns true if Block is disjoint with the
// other's type.
func ActivityStreamsBlockIsDisjointWith(other vocab.Type) bool {
	return typeblock.BlockIsDisjointWith(other)
}

// ActivityStreamsCollectionIsDisjointWith returns true if Collection is disjoint
// with the other's type.
func ActivityStreamsCollectionIsDisjointWith(other vocab.Type) bool {
	return typecollection.CollectionIsDisjointWith(other)
}

// ActivityStreamsCollectionPageIsDisjointWith returns true if CollectionPage is
// disjoint with the other's type.
func ActivityStreamsCollectionPageIsDisjointWith(other vocab.Type) bool {
	return typecollectionpage.CollectionPageIsDisjointWith(other)
}

// ActivityStreamsCreateIsDisjointWith returns true if Create is disjoint with the
// other's type.
func ActivityStreamsCreateIsDisjointWith(other vocab.Type) bool {
	return typecreate.CreateIsDisjointWith(other)
}

// ActivityStreamsDeleteIsDisjointWith returns true if Delete is disjoint with the
// other's type.
func ActivityStreamsDeleteIsDisjointWith(other vocab.Type) bool {
	return typedelete.DeleteIsDisjointWith(other)
}

// ActivityStreamsDislikeIsDisjointWith returns true if Dislike is disjoint with
// the other's type.
func ActivityStreamsDislikeIsDisjointWith(other vocab.Type) bool {
	return typedislike.DislikeIsDisjointWith(other)
}

// ActivityStreamsDocumentIsDisjointWith returns true if Document is disjoint with
// the other's type.
func ActivityStreamsDocumentIsDisjointWith(other vocab.Type) bool {
	return typedocument.DocumentIsDisjointWith(other)
}

// ActivityStreamsEventIsDisjointWith returns true if Event is disjoint with the
// other's type.
func ActivityStreamsEventIsDisjointWith(other vocab.Type) bool {
	return typeevent.EventIsDisjointWith(other)
}

// ActivityStreamsFlagIsDisjointWith returns true if Flag is disjoint with the
// other's type.
func ActivityStreamsFlagIsDisjointWith(other vocab.Type) bool {
	return typeflag.FlagIsDisjointWith(other)
}

// ActivityStreamsFollowIsDisjointWith returns true if Follow is disjoint with the
// other's type.
func ActivityStreamsFollowIsDisjointWith(other vocab.Type) bool {
	return typefollow.FollowIsDisjointWith(other)
}

// ActivityStreamsGroupIsDisjointWith returns true if Group is disjoint with the
// other's type.
func ActivityStreamsGroupIsDisjointWith(other vocab.Type) bool {
	return typegroup.GroupIsDisjointWith(other)
}

// ActivityStreamsIgnoreIsDisjointWith returns true if Ignore is disjoint with the
// other's type.
func ActivityStreamsIgnoreIsDisjointWith(other vocab.Type) bool {
	return typeignore.IgnoreIsDisjointWith(other)
}

// ActivityStreamsImageIsDisjointWith returns true if Image is disjoint with the
// other's type.
func ActivityStreamsImageIsDisjointWith(other vocab.Type) bool {
	return typeimage.ImageIsDisjointWith(other)
}

// ActivityStreamsIntransitiveActivityIsDisjointWith returns true if
// IntransitiveActivity is disjoint with the other's type.
func ActivityStreamsIntransitiveActivityIsDisjointWith(other vocab.Type) bool {
	return typeintransitiveactivity.IntransitiveActivityIsDisjointWith(other)
}

// ActivityStreamsInviteIsDisjointWith returns true if Invite is disjoint with the
// other's type.
func ActivityStreamsInviteIsDisjointWith(other vocab.Type) bool {
	return typeinvite.InviteIsDisjointWith(other)
}

// ActivityStreamsJoinIsDisjointWith returns true if Join is disjoint with the
// other's type.
func ActivityStreamsJoinIsDisjointWith(other vocab.Type) bool {
	return typejoin.JoinIsDisjointWith(other)
}

// ActivityStreamsLeaveIsDisjointWith returns true if Leave is disjoint with the
// other's type.
func ActivityStreamsLeaveIsDisjointWith(other vocab.Type) bool {
	return typeleave.LeaveIsDisjointWith(other)
}

// ActivityStreamsLikeIsDisjointWith returns true if Like is disjoint with the
// other's type.
func ActivityStreamsLikeIsDisjointWith(other vocab.Type) bool {
	return typelike.LikeIsDisjointWith(other)
}

// ActivityStreamsLinkIsDisjointWith returns true if Link is disjoint with the
// other's type.
func ActivityStreamsLinkIsDisjointWith(other vocab.Type) bool {
	return typelink.LinkIsDisjointWith(other)
}

// ActivityStreamsListenIsDisjointWith returns true if Listen is disjoint with the
// other's type.
func ActivityStreamsListenIsDisjointWith(other vocab.Type) bool {
	return typelisten.ListenIsDisjointWith(other)
}

// ActivityStreamsMentionIsDisjointWith returns true if Mention is disjoint with
// the other's type.
func ActivityStreamsMentionIsDisjointWith(other vocab.Type) bool {
	return typemention.MentionIsDisjointWith(other)
}

// ActivityStreamsMoveIsDisjointWith returns true if Move is disjoint with the
// other's type.
func ActivityStreamsMoveIsDisjointWith(other vocab.Type) bool {
	return typemove.MoveIsDisjointWith(other)
}

// ActivityStreamsNoteIsDisjointWith returns true if Note is disjoint with the
// other's type.
func ActivityStreamsNoteIsDisjointWith(other vocab.Type) bool {
	return typenote.NoteIsDisjointWith(other)
}

// ActivityStreamsObjectIsDisjointWith returns true if Object is disjoint with the
// other's type.
func ActivityStreamsObjectIsDisjointWith(other vocab.Type) bool {
	return typeobject.ObjectIsDisjointWith(other)
}

// ActivityStreamsOfferIsDisjointWith returns true if Offer is disjoint with the
// other's type.
func ActivityStreamsOfferIsDisjointWith(other vocab.Type) bool {
	return typeoffer.OfferIsDisjointWith(other)
}

// ActivityStreamsOrderedCollectionIsDisjointWith returns true if
// OrderedCollection is disjoint with the other's type.
func ActivityStreamsOrderedCollectionIsDisjointWith(other vocab.Type) bool {
	return typeorderedcollection.OrderedCollectionIsDisjointWith(other)
}

// ActivityStreamsOrderedCollectionPageIsDisjointWith returns true if
// OrderedCollectionPage is disjoint with the other's type.
func ActivityStreamsOrderedCollectionPageIsDisjointWith(other vocab.Type) bool {
	return typeorderedcollectionpage.OrderedCollectionPageIsDisjointWith(other)
}

// ActivityStreamsOrganizationIsDisjointWith returns true if Organization is
// disjoint with the other's type.
func ActivityStreamsOrganizationIsDisjointWith(other vocab.Type) bool {
	return typeorganization.OrganizationIsDisjointWith(other)
}

// ActivityStreamsPageIsDisjointWith returns true if Page is disjoint with the
// other's type.
func ActivityStreamsPageIsDisjointWith(other vocab.Type) bool {
	return typepage.PageIsDisjointWith(other)
}

// ActivityStreamsPersonIsDisjointWith returns true if Person is disjoint with the
// other's type.
func ActivityStreamsPersonIsDisjointWith(other vocab.Type) bool {
	return typeperson.PersonIsDisjointWith(other)
}

// ActivityStreamsPlaceIsDisjointWith returns true if Place is disjoint with the
// other's type.
func ActivityStreamsPlaceIsDisjointWith(other vocab.Type) bool {
	return typeplace.PlaceIsDisjointWith(other)
}

// ActivityStreamsProfileIsDisjointWith returns true if Profile is disjoint with
// the other's type.
func ActivityStreamsProfileIsDisjointWith(other vocab.Type) bool {
	return typeprofile.ProfileIsDisjointWith(other)
}

// ActivityStreamsPublicKeyIsDisjointWith returns true if PublicKey is disjoint
// with the other's type.
func ActivityStreamsPublicKeyIsDisjointWith(other vocab.Type) bool {
	return typepublickey.PublicKeyIsDisjointWith(other)
}

// ActivityStreamsQuestionIsDisjointWith returns true if Question is disjoint with
// the other's type.
func ActivityStreamsQuestionIsDisjointWith(other vocab.Type) bool {
	return typequestion.QuestionIsDisjointWith(other)
}

// ActivityStreamsReadIsDisjointWith returns true if Read is disjoint with the
// other's type.
func ActivityStreamsReadIsDisjointWith(other vocab.Type) bool {
	return typeread.ReadIsDisjointWith(other)
}

// ActivityStreamsRejectIsDisjointWith returns true if Reject is disjoint with the
// other's type.
func ActivityStreamsRejectIsDisjointWith(other vocab.Type) bool {
	return typereject.RejectIsDisjointWith(other)
}

// ActivityStreamsRelationshipIsDisjointWith returns true if Relationship is
// disjoint with the other's type.
func ActivityStreamsRelationshipIsDisjointWith(other vocab.Type) bool {
	return typerelationship.RelationshipIsDisjointWith(other)
}

// ActivityStreamsRemoveIsDisjointWith returns true if Remove is disjoint with the
// other's type.
func ActivityStreamsRemoveIsDisjointWith(other vocab.Type) bool {
	return typeremove.RemoveIsDisjointWith(other)
}

// ActivityStreamsServiceIsDisjointWith returns true if Service is disjoint with
// the other's type.
func ActivityStreamsServiceIsDisjointWith(other vocab.Type) bool {
	return typeservice.ServiceIsDisjointWith(other)
}

// ActivityStreamsTentativeAcceptIsDisjointWith returns true if TentativeAccept is
// disjoint with the other's type.
func ActivityStreamsTentativeAcceptIsDisjointWith(other vocab.Type) bool {
	return typetentativeaccept.TentativeAcceptIsDisjointWith(other)
}

// ActivityStreamsTentativeRejectIsDisjointWith returns true if TentativeReject is
// disjoint with the other's type.
func ActivityStreamsTentativeRejectIsDisjointWith(other vocab.Type) bool {
	return typetentativereject.TentativeRejectIsDisjointWith(other)
}

// ActivityStreamsTombstoneIsDisjointWith returns true if Tombstone is disjoint
// with the other's type.
func ActivityStreamsTombstoneIsDisjointWith(other vocab.Type) bool {
	return typetombstone.TombstoneIsDisjointWith(other)
}

// ActivityStreamsTravelIsDisjointWith returns true if Travel is disjoint with the
// other's type.
func ActivityStreamsTravelIsDisjointWith(other vocab.Type) bool {
	return typetravel.TravelIsDisjointWith(other)
}

// ActivityStreamsUndoIsDisjointWith returns true if Undo is disjoint with the
// other's type.
func ActivityStreamsUndoIsDisjointWith(other vocab.Type) bool {
	return typeundo.UndoIsDisjointWith(other)
}

// ActivityStreamsUpdateIsDisjointWith returns true if Update is disjoint with the
// other's type.
func ActivityStreamsUpdateIsDisjointWith(other vocab.Type) bool {
	return typeupdate.UpdateIsDisjointWith(other)
}

// ActivityStreamsVideoIsDisjointWith returns true if Video is disjoint with the
// other's type.
func ActivityStreamsVideoIsDisjointWith(other vocab.Type) bool {
	return typevideo.VideoIsDisjointWith(other)
}

// ActivityStreamsViewIsDisjointWith returns true if View is disjoint with the
// other's type.
func ActivityStreamsViewIsDisjointWith(other vocab.Type) bool {
	return typeview.ViewIsDisjointWith(other)
}
