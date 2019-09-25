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

// NewActivityStreamsAccept creates a new ActivityStreamsAccept
func NewActivityStreamsAccept() vocab.ActivityStreamsAccept {
	return typeaccept.NewActivityStreamsAccept()
}

// NewActivityStreamsActivity creates a new ActivityStreamsActivity
func NewActivityStreamsActivity() vocab.ActivityStreamsActivity {
	return typeactivity.NewActivityStreamsActivity()
}

// NewActivityStreamsAdd creates a new ActivityStreamsAdd
func NewActivityStreamsAdd() vocab.ActivityStreamsAdd {
	return typeadd.NewActivityStreamsAdd()
}

// NewActivityStreamsAnnounce creates a new ActivityStreamsAnnounce
func NewActivityStreamsAnnounce() vocab.ActivityStreamsAnnounce {
	return typeannounce.NewActivityStreamsAnnounce()
}

// NewActivityStreamsApplication creates a new ActivityStreamsApplication
func NewActivityStreamsApplication() vocab.ActivityStreamsApplication {
	return typeapplication.NewActivityStreamsApplication()
}

// NewActivityStreamsArrive creates a new ActivityStreamsArrive
func NewActivityStreamsArrive() vocab.ActivityStreamsArrive {
	return typearrive.NewActivityStreamsArrive()
}

// NewActivityStreamsArticle creates a new ActivityStreamsArticle
func NewActivityStreamsArticle() vocab.ActivityStreamsArticle {
	return typearticle.NewActivityStreamsArticle()
}

// NewActivityStreamsAudio creates a new ActivityStreamsAudio
func NewActivityStreamsAudio() vocab.ActivityStreamsAudio {
	return typeaudio.NewActivityStreamsAudio()
}

// NewActivityStreamsBlock creates a new ActivityStreamsBlock
func NewActivityStreamsBlock() vocab.ActivityStreamsBlock {
	return typeblock.NewActivityStreamsBlock()
}

// NewActivityStreamsCollection creates a new ActivityStreamsCollection
func NewActivityStreamsCollection() vocab.ActivityStreamsCollection {
	return typecollection.NewActivityStreamsCollection()
}

// NewActivityStreamsCollectionPage creates a new ActivityStreamsCollectionPage
func NewActivityStreamsCollectionPage() vocab.ActivityStreamsCollectionPage {
	return typecollectionpage.NewActivityStreamsCollectionPage()
}

// NewActivityStreamsCreate creates a new ActivityStreamsCreate
func NewActivityStreamsCreate() vocab.ActivityStreamsCreate {
	return typecreate.NewActivityStreamsCreate()
}

// NewActivityStreamsDelete creates a new ActivityStreamsDelete
func NewActivityStreamsDelete() vocab.ActivityStreamsDelete {
	return typedelete.NewActivityStreamsDelete()
}

// NewActivityStreamsDislike creates a new ActivityStreamsDislike
func NewActivityStreamsDislike() vocab.ActivityStreamsDislike {
	return typedislike.NewActivityStreamsDislike()
}

// NewActivityStreamsDocument creates a new ActivityStreamsDocument
func NewActivityStreamsDocument() vocab.ActivityStreamsDocument {
	return typedocument.NewActivityStreamsDocument()
}

// NewActivityStreamsEvent creates a new ActivityStreamsEvent
func NewActivityStreamsEvent() vocab.ActivityStreamsEvent {
	return typeevent.NewActivityStreamsEvent()
}

// NewActivityStreamsFlag creates a new ActivityStreamsFlag
func NewActivityStreamsFlag() vocab.ActivityStreamsFlag {
	return typeflag.NewActivityStreamsFlag()
}

// NewActivityStreamsFollow creates a new ActivityStreamsFollow
func NewActivityStreamsFollow() vocab.ActivityStreamsFollow {
	return typefollow.NewActivityStreamsFollow()
}

// NewActivityStreamsGroup creates a new ActivityStreamsGroup
func NewActivityStreamsGroup() vocab.ActivityStreamsGroup {
	return typegroup.NewActivityStreamsGroup()
}

// NewActivityStreamsIgnore creates a new ActivityStreamsIgnore
func NewActivityStreamsIgnore() vocab.ActivityStreamsIgnore {
	return typeignore.NewActivityStreamsIgnore()
}

// NewActivityStreamsImage creates a new ActivityStreamsImage
func NewActivityStreamsImage() vocab.ActivityStreamsImage {
	return typeimage.NewActivityStreamsImage()
}

// NewActivityStreamsIntransitiveActivity creates a new
// ActivityStreamsIntransitiveActivity
func NewActivityStreamsIntransitiveActivity() vocab.ActivityStreamsIntransitiveActivity {
	return typeintransitiveactivity.NewActivityStreamsIntransitiveActivity()
}

// NewActivityStreamsInvite creates a new ActivityStreamsInvite
func NewActivityStreamsInvite() vocab.ActivityStreamsInvite {
	return typeinvite.NewActivityStreamsInvite()
}

// NewActivityStreamsJoin creates a new ActivityStreamsJoin
func NewActivityStreamsJoin() vocab.ActivityStreamsJoin {
	return typejoin.NewActivityStreamsJoin()
}

// NewActivityStreamsLeave creates a new ActivityStreamsLeave
func NewActivityStreamsLeave() vocab.ActivityStreamsLeave {
	return typeleave.NewActivityStreamsLeave()
}

// NewActivityStreamsLike creates a new ActivityStreamsLike
func NewActivityStreamsLike() vocab.ActivityStreamsLike {
	return typelike.NewActivityStreamsLike()
}

// NewActivityStreamsLink creates a new ActivityStreamsLink
func NewActivityStreamsLink() vocab.ActivityStreamsLink {
	return typelink.NewActivityStreamsLink()
}

// NewActivityStreamsListen creates a new ActivityStreamsListen
func NewActivityStreamsListen() vocab.ActivityStreamsListen {
	return typelisten.NewActivityStreamsListen()
}

// NewActivityStreamsMention creates a new ActivityStreamsMention
func NewActivityStreamsMention() vocab.ActivityStreamsMention {
	return typemention.NewActivityStreamsMention()
}

// NewActivityStreamsMove creates a new ActivityStreamsMove
func NewActivityStreamsMove() vocab.ActivityStreamsMove {
	return typemove.NewActivityStreamsMove()
}

// NewActivityStreamsNote creates a new ActivityStreamsNote
func NewActivityStreamsNote() vocab.ActivityStreamsNote {
	return typenote.NewActivityStreamsNote()
}

// NewActivityStreamsObject creates a new ActivityStreamsObject
func NewActivityStreamsObject() vocab.ActivityStreamsObject {
	return typeobject.NewActivityStreamsObject()
}

// NewActivityStreamsOffer creates a new ActivityStreamsOffer
func NewActivityStreamsOffer() vocab.ActivityStreamsOffer {
	return typeoffer.NewActivityStreamsOffer()
}

// NewActivityStreamsOrderedCollection creates a new
// ActivityStreamsOrderedCollection
func NewActivityStreamsOrderedCollection() vocab.ActivityStreamsOrderedCollection {
	return typeorderedcollection.NewActivityStreamsOrderedCollection()
}

// NewActivityStreamsOrderedCollectionPage creates a new
// ActivityStreamsOrderedCollectionPage
func NewActivityStreamsOrderedCollectionPage() vocab.ActivityStreamsOrderedCollectionPage {
	return typeorderedcollectionpage.NewActivityStreamsOrderedCollectionPage()
}

// NewActivityStreamsOrganization creates a new ActivityStreamsOrganization
func NewActivityStreamsOrganization() vocab.ActivityStreamsOrganization {
	return typeorganization.NewActivityStreamsOrganization()
}

// NewActivityStreamsPage creates a new ActivityStreamsPage
func NewActivityStreamsPage() vocab.ActivityStreamsPage {
	return typepage.NewActivityStreamsPage()
}

// NewActivityStreamsPerson creates a new ActivityStreamsPerson
func NewActivityStreamsPerson() vocab.ActivityStreamsPerson {
	return typeperson.NewActivityStreamsPerson()
}

// NewActivityStreamsPlace creates a new ActivityStreamsPlace
func NewActivityStreamsPlace() vocab.ActivityStreamsPlace {
	return typeplace.NewActivityStreamsPlace()
}

// NewActivityStreamsProfile creates a new ActivityStreamsProfile
func NewActivityStreamsProfile() vocab.ActivityStreamsProfile {
	return typeprofile.NewActivityStreamsProfile()
}

// NewActivityStreamsQuestion creates a new ActivityStreamsQuestion
func NewActivityStreamsQuestion() vocab.ActivityStreamsQuestion {
	return typequestion.NewActivityStreamsQuestion()
}

// NewActivityStreamsRead creates a new ActivityStreamsRead
func NewActivityStreamsRead() vocab.ActivityStreamsRead {
	return typeread.NewActivityStreamsRead()
}

// NewActivityStreamsReject creates a new ActivityStreamsReject
func NewActivityStreamsReject() vocab.ActivityStreamsReject {
	return typereject.NewActivityStreamsReject()
}

// NewActivityStreamsRelationship creates a new ActivityStreamsRelationship
func NewActivityStreamsRelationship() vocab.ActivityStreamsRelationship {
	return typerelationship.NewActivityStreamsRelationship()
}

// NewActivityStreamsRemove creates a new ActivityStreamsRemove
func NewActivityStreamsRemove() vocab.ActivityStreamsRemove {
	return typeremove.NewActivityStreamsRemove()
}

// NewActivityStreamsService creates a new ActivityStreamsService
func NewActivityStreamsService() vocab.ActivityStreamsService {
	return typeservice.NewActivityStreamsService()
}

// NewActivityStreamsTentativeAccept creates a new ActivityStreamsTentativeAccept
func NewActivityStreamsTentativeAccept() vocab.ActivityStreamsTentativeAccept {
	return typetentativeaccept.NewActivityStreamsTentativeAccept()
}

// NewActivityStreamsTentativeReject creates a new ActivityStreamsTentativeReject
func NewActivityStreamsTentativeReject() vocab.ActivityStreamsTentativeReject {
	return typetentativereject.NewActivityStreamsTentativeReject()
}

// NewActivityStreamsTombstone creates a new ActivityStreamsTombstone
func NewActivityStreamsTombstone() vocab.ActivityStreamsTombstone {
	return typetombstone.NewActivityStreamsTombstone()
}

// NewActivityStreamsTravel creates a new ActivityStreamsTravel
func NewActivityStreamsTravel() vocab.ActivityStreamsTravel {
	return typetravel.NewActivityStreamsTravel()
}

// NewActivityStreamsUndo creates a new ActivityStreamsUndo
func NewActivityStreamsUndo() vocab.ActivityStreamsUndo {
	return typeundo.NewActivityStreamsUndo()
}

// NewActivityStreamsUpdate creates a new ActivityStreamsUpdate
func NewActivityStreamsUpdate() vocab.ActivityStreamsUpdate {
	return typeupdate.NewActivityStreamsUpdate()
}

// NewActivityStreamsVideo creates a new ActivityStreamsVideo
func NewActivityStreamsVideo() vocab.ActivityStreamsVideo {
	return typevideo.NewActivityStreamsVideo()
}

// NewActivityStreamsView creates a new ActivityStreamsView
func NewActivityStreamsView() vocab.ActivityStreamsView {
	return typeview.NewActivityStreamsView()
}
