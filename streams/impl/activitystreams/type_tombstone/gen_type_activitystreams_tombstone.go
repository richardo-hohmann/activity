package typetombstone

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"strings"
)

// A Tombstone represents a content object that has been deleted. It can be used
// in Collections to signify that there used to be an object at this position,
// but it has been deleted.
//
// Example 60 (https://www.w3.org/TR/activitystreams-vocabulary/#ex184b-jsonld):
//   {
//     "name": "Vacation photos 2016",
//     "orderedItems": [
//       {
//         "id": "http://image.example/1",
//         "type": "Image"
//       },
//       {
//         "deleted": "2016-03-17T00:00:00Z",
//         "formerType": "/Image",
//         "id": "http://image.example/2",
//         "type": "Tombstone"
//       },
//       {
//         "id": "http://image.example/3",
//         "type": "Image"
//       }
//     ],
//     "totalItems": 3,
//     "type": "OrderedCollection"
//   }
type ActivityStreamsTombstone struct {
	ActivityStreamsAltitude     vocab.ActivityStreamsAltitudeProperty
	ActivityStreamsAttachment   vocab.ActivityStreamsAttachmentProperty
	ActivityStreamsAttributedTo vocab.ActivityStreamsAttributedToProperty
	ActivityStreamsAudience     vocab.ActivityStreamsAudienceProperty
	ActivityStreamsBcc          vocab.ActivityStreamsBccProperty
	ActivityStreamsBto          vocab.ActivityStreamsBtoProperty
	ActivityStreamsCc           vocab.ActivityStreamsCcProperty
	ActivityStreamsContent      vocab.ActivityStreamsContentProperty
	ActivityStreamsContext      vocab.ActivityStreamsContextProperty
	ActivityStreamsDeleted      vocab.ActivityStreamsDeletedProperty
	ActivityStreamsDuration     vocab.ActivityStreamsDurationProperty
	ActivityStreamsEndTime      vocab.ActivityStreamsEndTimeProperty
	ActivityStreamsFormerType   vocab.ActivityStreamsFormerTypeProperty
	ActivityStreamsGenerator    vocab.ActivityStreamsGeneratorProperty
	ActivityStreamsIcon         vocab.ActivityStreamsIconProperty
	ActivityStreamsId           vocab.ActivityStreamsIdProperty
	ActivityStreamsImage        vocab.ActivityStreamsImageProperty
	ActivityStreamsInReplyTo    vocab.ActivityStreamsInReplyToProperty
	ActivityStreamsLikes        vocab.ActivityStreamsLikesProperty
	ActivityStreamsLocation     vocab.ActivityStreamsLocationProperty
	ActivityStreamsMediaType    vocab.ActivityStreamsMediaTypeProperty
	ActivityStreamsName         vocab.ActivityStreamsNameProperty
	ActivityStreamsObject       vocab.ActivityStreamsObjectProperty
	ActivityStreamsPreview      vocab.ActivityStreamsPreviewProperty
	ActivityStreamsPublished    vocab.ActivityStreamsPublishedProperty
	ActivityStreamsReplies      vocab.ActivityStreamsRepliesProperty
	ActivityStreamsShares       vocab.ActivityStreamsSharesProperty
	ActivityStreamsStartTime    vocab.ActivityStreamsStartTimeProperty
	ActivityStreamsSummary      vocab.ActivityStreamsSummaryProperty
	ActivityStreamsTag          vocab.ActivityStreamsTagProperty
	ActivityStreamsTo           vocab.ActivityStreamsToProperty
	ActivityStreamsType         vocab.ActivityStreamsTypeProperty
	ActivityStreamsUpdated      vocab.ActivityStreamsUpdatedProperty
	ActivityStreamsUrl          vocab.ActivityStreamsUrlProperty
	alias                       string
	unknown                     map[string]interface{}
}

// ActivityStreamsTombstoneExtends returns true if the Tombstone type extends from
// the other type.
func ActivityStreamsTombstoneExtends(other vocab.Type) bool {
	extensions := []string{"Object"}
	for _, ext := range extensions {
		if ext == other.GetName() {
			return true
		}
	}
	return false
}

// DeserializeTombstone creates a Tombstone from a map representation that has
// been unmarshalled from a text or binary format.
func DeserializeTombstone(m map[string]interface{}, aliasMap map[string]string) (*ActivityStreamsTombstone, error) {
	alias := ""
	aliasPrefix := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
		aliasPrefix = a + ":"
	}
	this := &ActivityStreamsTombstone{
		alias:   alias,
		unknown: make(map[string]interface{}),
	}
	if typeValue, ok := m["type"]; !ok {
		return nil, fmt.Errorf("no \"type\" property in map")
	} else if typeString, ok := typeValue.(string); ok {
		typeName := strings.TrimPrefix(typeString, aliasPrefix)
		if typeName != "Tombstone" {
			return nil, fmt.Errorf("\"type\" property is not of %q type: %s", "Tombstone", typeName)
		}
		// Fall through, success in finding a proper Type
	} else if arrType, ok := typeValue.([]interface{}); ok {
		found := false
		for _, elemVal := range arrType {
			if typeString, ok := elemVal.(string); ok && strings.TrimPrefix(typeString, aliasPrefix) == "Tombstone" {
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("could not find a \"type\" property of value %q", "Tombstone")
		}
		// Fall through, success in finding a proper Type
	} else {
		return nil, fmt.Errorf("\"type\" property is unrecognized type: %T", typeValue)
	}
	// Begin: Known property deserialization
	if p, err := mgr.DeserializeAltitudePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsAltitude = p
	}
	if p, err := mgr.DeserializeAttachmentPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsAttachment = p
	}
	if p, err := mgr.DeserializeAttributedToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsAttributedTo = p
	}
	if p, err := mgr.DeserializeAudiencePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsAudience = p
	}
	if p, err := mgr.DeserializeBccPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsBcc = p
	}
	if p, err := mgr.DeserializeBtoPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsBto = p
	}
	if p, err := mgr.DeserializeCcPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsCc = p
	}
	if p, err := mgr.DeserializeContentPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsContent = p
	}
	if p, err := mgr.DeserializeContextPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsContext = p
	}
	if p, err := mgr.DeserializeDeletedPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsDeleted = p
	}
	if p, err := mgr.DeserializeDurationPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsDuration = p
	}
	if p, err := mgr.DeserializeEndTimePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsEndTime = p
	}
	if p, err := mgr.DeserializeFormerTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsFormerType = p
	}
	if p, err := mgr.DeserializeGeneratorPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsGenerator = p
	}
	if p, err := mgr.DeserializeIconPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsIcon = p
	}
	if p, err := mgr.DeserializeIdPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsId = p
	}
	if p, err := mgr.DeserializeImagePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsImage = p
	}
	if p, err := mgr.DeserializeInReplyToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsInReplyTo = p
	}
	if p, err := mgr.DeserializeLikesPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsLikes = p
	}
	if p, err := mgr.DeserializeLocationPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsLocation = p
	}
	if p, err := mgr.DeserializeMediaTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsMediaType = p
	}
	if p, err := mgr.DeserializeNamePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsName = p
	}
	if p, err := mgr.DeserializeObjectPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsObject = p
	}
	if p, err := mgr.DeserializePreviewPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsPreview = p
	}
	if p, err := mgr.DeserializePublishedPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsPublished = p
	}
	if p, err := mgr.DeserializeRepliesPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsReplies = p
	}
	if p, err := mgr.DeserializeSharesPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsShares = p
	}
	if p, err := mgr.DeserializeStartTimePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsStartTime = p
	}
	if p, err := mgr.DeserializeSummaryPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsSummary = p
	}
	if p, err := mgr.DeserializeTagPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsTag = p
	}
	if p, err := mgr.DeserializeToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsTo = p
	}
	if p, err := mgr.DeserializeTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsType = p
	}
	if p, err := mgr.DeserializeUpdatedPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsUpdated = p
	}
	if p, err := mgr.DeserializeUrlPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsUrl = p
	}
	// End: Known property deserialization

	// Begin: Unknown deserialization
	for k, v := range m {
		// Begin: Code that ensures a property name is unknown
		if k == "altitude" {
			continue
		} else if k == "attachment" {
			continue
		} else if k == "attributedTo" {
			continue
		} else if k == "audience" {
			continue
		} else if k == "bcc" {
			continue
		} else if k == "bto" {
			continue
		} else if k == "cc" {
			continue
		} else if k == "content" {
			continue
		} else if k == "contentMap" {
			continue
		} else if k == "context" {
			continue
		} else if k == "deleted" {
			continue
		} else if k == "duration" {
			continue
		} else if k == "endTime" {
			continue
		} else if k == "formerType" {
			continue
		} else if k == "generator" {
			continue
		} else if k == "icon" {
			continue
		} else if k == "id" {
			continue
		} else if k == "image" {
			continue
		} else if k == "inReplyTo" {
			continue
		} else if k == "likes" {
			continue
		} else if k == "location" {
			continue
		} else if k == "mediaType" {
			continue
		} else if k == "name" {
			continue
		} else if k == "nameMap" {
			continue
		} else if k == "object" {
			continue
		} else if k == "preview" {
			continue
		} else if k == "published" {
			continue
		} else if k == "replies" {
			continue
		} else if k == "shares" {
			continue
		} else if k == "startTime" {
			continue
		} else if k == "summary" {
			continue
		} else if k == "summaryMap" {
			continue
		} else if k == "tag" {
			continue
		} else if k == "to" {
			continue
		} else if k == "type" {
			continue
		} else if k == "updated" {
			continue
		} else if k == "url" {
			continue
		} // End: Code that ensures a property name is unknown

		this.unknown[k] = v
	}
	// End: Unknown deserialization

	return this, nil
}

// NewActivityStreamsTombstone creates a new Tombstone type
func NewActivityStreamsTombstone() *ActivityStreamsTombstone {
	typeProp := typePropertyConstructor()
	typeProp.AppendXMLSchemaString("Tombstone")
	return &ActivityStreamsTombstone{
		ActivityStreamsType: typeProp,
		alias:               "",
		unknown:             make(map[string]interface{}, 0),
	}
}

// TombstoneIsDisjointWith returns true if the other provided type is disjoint
// with the Tombstone type.
func TombstoneIsDisjointWith(other vocab.Type) bool {
	disjointWith := []string{"Link", "Mention"}
	for _, disjoint := range disjointWith {
		if disjoint == other.GetName() {
			return true
		}
	}
	return false
}

// TombstoneIsExtendedBy returns true if the other provided type extends from the
// Tombstone type.
func TombstoneIsExtendedBy(other vocab.Type) bool {
	// Shortcut implementation: is not extended by anything.
	return false
}

// GetActivityStreamsAltitude returns the "altitude" property if it exists, and
// nil otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsAltitude() vocab.ActivityStreamsAltitudeProperty {
	return this.ActivityStreamsAltitude
}

// GetActivityStreamsAttachment returns the "attachment" property if it exists,
// and nil otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsAttachment() vocab.ActivityStreamsAttachmentProperty {
	return this.ActivityStreamsAttachment
}

// GetActivityStreamsAttributedTo returns the "attributedTo" property if it
// exists, and nil otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsAttributedTo() vocab.ActivityStreamsAttributedToProperty {
	return this.ActivityStreamsAttributedTo
}

// GetActivityStreamsAudience returns the "audience" property if it exists, and
// nil otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsAudience() vocab.ActivityStreamsAudienceProperty {
	return this.ActivityStreamsAudience
}

// GetActivityStreamsBcc returns the "bcc" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsBcc() vocab.ActivityStreamsBccProperty {
	return this.ActivityStreamsBcc
}

// GetActivityStreamsBto returns the "bto" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsBto() vocab.ActivityStreamsBtoProperty {
	return this.ActivityStreamsBto
}

// GetActivityStreamsCc returns the "cc" property if it exists, and nil otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsCc() vocab.ActivityStreamsCcProperty {
	return this.ActivityStreamsCc
}

// GetActivityStreamsContent returns the "content" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsContent() vocab.ActivityStreamsContentProperty {
	return this.ActivityStreamsContent
}

// GetActivityStreamsContext returns the "context" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsContext() vocab.ActivityStreamsContextProperty {
	return this.ActivityStreamsContext
}

// GetActivityStreamsDeleted returns the "deleted" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsDeleted() vocab.ActivityStreamsDeletedProperty {
	return this.ActivityStreamsDeleted
}

// GetActivityStreamsDuration returns the "duration" property if it exists, and
// nil otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsDuration() vocab.ActivityStreamsDurationProperty {
	return this.ActivityStreamsDuration
}

// GetActivityStreamsEndTime returns the "endTime" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsEndTime() vocab.ActivityStreamsEndTimeProperty {
	return this.ActivityStreamsEndTime
}

// GetActivityStreamsFormerType returns the "formerType" property if it exists,
// and nil otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsFormerType() vocab.ActivityStreamsFormerTypeProperty {
	return this.ActivityStreamsFormerType
}

// GetActivityStreamsGenerator returns the "generator" property if it exists, and
// nil otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsGenerator() vocab.ActivityStreamsGeneratorProperty {
	return this.ActivityStreamsGenerator
}

// GetActivityStreamsIcon returns the "icon" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsIcon() vocab.ActivityStreamsIconProperty {
	return this.ActivityStreamsIcon
}

// GetActivityStreamsId returns the "id" property if it exists, and nil otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsId() vocab.ActivityStreamsIdProperty {
	return this.ActivityStreamsId
}

// GetActivityStreamsImage returns the "image" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsImage() vocab.ActivityStreamsImageProperty {
	return this.ActivityStreamsImage
}

// GetActivityStreamsInReplyTo returns the "inReplyTo" property if it exists, and
// nil otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsInReplyTo() vocab.ActivityStreamsInReplyToProperty {
	return this.ActivityStreamsInReplyTo
}

// GetActivityStreamsLikes returns the "likes" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsLikes() vocab.ActivityStreamsLikesProperty {
	return this.ActivityStreamsLikes
}

// GetActivityStreamsLocation returns the "location" property if it exists, and
// nil otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsLocation() vocab.ActivityStreamsLocationProperty {
	return this.ActivityStreamsLocation
}

// GetActivityStreamsMediaType returns the "mediaType" property if it exists, and
// nil otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsMediaType() vocab.ActivityStreamsMediaTypeProperty {
	return this.ActivityStreamsMediaType
}

// GetActivityStreamsName returns the "name" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsName() vocab.ActivityStreamsNameProperty {
	return this.ActivityStreamsName
}

// GetActivityStreamsObject returns the "object" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsObject() vocab.ActivityStreamsObjectProperty {
	return this.ActivityStreamsObject
}

// GetActivityStreamsPreview returns the "preview" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsPreview() vocab.ActivityStreamsPreviewProperty {
	return this.ActivityStreamsPreview
}

// GetActivityStreamsPublished returns the "published" property if it exists, and
// nil otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsPublished() vocab.ActivityStreamsPublishedProperty {
	return this.ActivityStreamsPublished
}

// GetActivityStreamsReplies returns the "replies" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsReplies() vocab.ActivityStreamsRepliesProperty {
	return this.ActivityStreamsReplies
}

// GetActivityStreamsShares returns the "shares" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsShares() vocab.ActivityStreamsSharesProperty {
	return this.ActivityStreamsShares
}

// GetActivityStreamsStartTime returns the "startTime" property if it exists, and
// nil otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsStartTime() vocab.ActivityStreamsStartTimeProperty {
	return this.ActivityStreamsStartTime
}

// GetActivityStreamsSummary returns the "summary" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsSummary() vocab.ActivityStreamsSummaryProperty {
	return this.ActivityStreamsSummary
}

// GetActivityStreamsTag returns the "tag" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsTag() vocab.ActivityStreamsTagProperty {
	return this.ActivityStreamsTag
}

// GetActivityStreamsTo returns the "to" property if it exists, and nil otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsTo() vocab.ActivityStreamsToProperty {
	return this.ActivityStreamsTo
}

// GetActivityStreamsType returns the "type" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsType() vocab.ActivityStreamsTypeProperty {
	return this.ActivityStreamsType
}

// GetActivityStreamsUpdated returns the "updated" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsUpdated() vocab.ActivityStreamsUpdatedProperty {
	return this.ActivityStreamsUpdated
}

// GetActivityStreamsUrl returns the "url" property if it exists, and nil
// otherwise.
func (this ActivityStreamsTombstone) GetActivityStreamsUrl() vocab.ActivityStreamsUrlProperty {
	return this.ActivityStreamsUrl
}

// GetName returns the name of this type.
func (this ActivityStreamsTombstone) GetName() string {
	return "Tombstone"
}

// GetUnknownProperties returns the unknown properties for the Tombstone type.
// Note that this should not be used by app developers. It is only used to
// help determine which implementation is LessThan the other. Developers who
// are creating a different implementation of this type's interface can use
// this method in their LessThan implementation, but routine ActivityPub
// applications should not use this to bypass the code generation tool.
func (this ActivityStreamsTombstone) GetUnknownProperties() map[string]interface{} {
	return this.unknown
}

// IsExtending returns true if the Tombstone type extends from the other type.
func (this ActivityStreamsTombstone) IsExtending(other vocab.Type) bool {
	return ActivityStreamsTombstoneExtends(other)
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// type and the specific properties that are set. The value in the map is the
// alias used to import the type and its properties.
func (this ActivityStreamsTombstone) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	m = this.helperJSONLDContext(this.ActivityStreamsAltitude, m)
	m = this.helperJSONLDContext(this.ActivityStreamsAttachment, m)
	m = this.helperJSONLDContext(this.ActivityStreamsAttributedTo, m)
	m = this.helperJSONLDContext(this.ActivityStreamsAudience, m)
	m = this.helperJSONLDContext(this.ActivityStreamsBcc, m)
	m = this.helperJSONLDContext(this.ActivityStreamsBto, m)
	m = this.helperJSONLDContext(this.ActivityStreamsCc, m)
	m = this.helperJSONLDContext(this.ActivityStreamsContent, m)
	m = this.helperJSONLDContext(this.ActivityStreamsContext, m)
	m = this.helperJSONLDContext(this.ActivityStreamsDeleted, m)
	m = this.helperJSONLDContext(this.ActivityStreamsDuration, m)
	m = this.helperJSONLDContext(this.ActivityStreamsEndTime, m)
	m = this.helperJSONLDContext(this.ActivityStreamsFormerType, m)
	m = this.helperJSONLDContext(this.ActivityStreamsGenerator, m)
	m = this.helperJSONLDContext(this.ActivityStreamsIcon, m)
	m = this.helperJSONLDContext(this.ActivityStreamsId, m)
	m = this.helperJSONLDContext(this.ActivityStreamsImage, m)
	m = this.helperJSONLDContext(this.ActivityStreamsInReplyTo, m)
	m = this.helperJSONLDContext(this.ActivityStreamsLikes, m)
	m = this.helperJSONLDContext(this.ActivityStreamsLocation, m)
	m = this.helperJSONLDContext(this.ActivityStreamsMediaType, m)
	m = this.helperJSONLDContext(this.ActivityStreamsName, m)
	m = this.helperJSONLDContext(this.ActivityStreamsObject, m)
	m = this.helperJSONLDContext(this.ActivityStreamsPreview, m)
	m = this.helperJSONLDContext(this.ActivityStreamsPublished, m)
	m = this.helperJSONLDContext(this.ActivityStreamsReplies, m)
	m = this.helperJSONLDContext(this.ActivityStreamsShares, m)
	m = this.helperJSONLDContext(this.ActivityStreamsStartTime, m)
	m = this.helperJSONLDContext(this.ActivityStreamsSummary, m)
	m = this.helperJSONLDContext(this.ActivityStreamsTag, m)
	m = this.helperJSONLDContext(this.ActivityStreamsTo, m)
	m = this.helperJSONLDContext(this.ActivityStreamsType, m)
	m = this.helperJSONLDContext(this.ActivityStreamsUpdated, m)
	m = this.helperJSONLDContext(this.ActivityStreamsUrl, m)

	return m
}

// LessThan computes if this Tombstone is lesser, with an arbitrary but stable
// determination.
func (this ActivityStreamsTombstone) LessThan(o vocab.ActivityStreamsTombstone) bool {
	// Begin: Compare known properties
	// Compare property "altitude"
	if lhs, rhs := this.ActivityStreamsAltitude, o.GetActivityStreamsAltitude(); lhs != nil && rhs != nil {
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
	// Compare property "attachment"
	if lhs, rhs := this.ActivityStreamsAttachment, o.GetActivityStreamsAttachment(); lhs != nil && rhs != nil {
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
	// Compare property "attributedTo"
	if lhs, rhs := this.ActivityStreamsAttributedTo, o.GetActivityStreamsAttributedTo(); lhs != nil && rhs != nil {
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
	// Compare property "audience"
	if lhs, rhs := this.ActivityStreamsAudience, o.GetActivityStreamsAudience(); lhs != nil && rhs != nil {
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
	// Compare property "bcc"
	if lhs, rhs := this.ActivityStreamsBcc, o.GetActivityStreamsBcc(); lhs != nil && rhs != nil {
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
	// Compare property "bto"
	if lhs, rhs := this.ActivityStreamsBto, o.GetActivityStreamsBto(); lhs != nil && rhs != nil {
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
	// Compare property "cc"
	if lhs, rhs := this.ActivityStreamsCc, o.GetActivityStreamsCc(); lhs != nil && rhs != nil {
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
	// Compare property "content"
	if lhs, rhs := this.ActivityStreamsContent, o.GetActivityStreamsContent(); lhs != nil && rhs != nil {
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
	// Compare property "context"
	if lhs, rhs := this.ActivityStreamsContext, o.GetActivityStreamsContext(); lhs != nil && rhs != nil {
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
	// Compare property "deleted"
	if lhs, rhs := this.ActivityStreamsDeleted, o.GetActivityStreamsDeleted(); lhs != nil && rhs != nil {
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
	// Compare property "duration"
	if lhs, rhs := this.ActivityStreamsDuration, o.GetActivityStreamsDuration(); lhs != nil && rhs != nil {
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
	// Compare property "endTime"
	if lhs, rhs := this.ActivityStreamsEndTime, o.GetActivityStreamsEndTime(); lhs != nil && rhs != nil {
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
	// Compare property "formerType"
	if lhs, rhs := this.ActivityStreamsFormerType, o.GetActivityStreamsFormerType(); lhs != nil && rhs != nil {
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
	// Compare property "generator"
	if lhs, rhs := this.ActivityStreamsGenerator, o.GetActivityStreamsGenerator(); lhs != nil && rhs != nil {
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
	// Compare property "icon"
	if lhs, rhs := this.ActivityStreamsIcon, o.GetActivityStreamsIcon(); lhs != nil && rhs != nil {
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
	// Compare property "id"
	if lhs, rhs := this.ActivityStreamsId, o.GetActivityStreamsId(); lhs != nil && rhs != nil {
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
	// Compare property "image"
	if lhs, rhs := this.ActivityStreamsImage, o.GetActivityStreamsImage(); lhs != nil && rhs != nil {
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
	// Compare property "inReplyTo"
	if lhs, rhs := this.ActivityStreamsInReplyTo, o.GetActivityStreamsInReplyTo(); lhs != nil && rhs != nil {
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
	// Compare property "likes"
	if lhs, rhs := this.ActivityStreamsLikes, o.GetActivityStreamsLikes(); lhs != nil && rhs != nil {
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
	// Compare property "location"
	if lhs, rhs := this.ActivityStreamsLocation, o.GetActivityStreamsLocation(); lhs != nil && rhs != nil {
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
	// Compare property "mediaType"
	if lhs, rhs := this.ActivityStreamsMediaType, o.GetActivityStreamsMediaType(); lhs != nil && rhs != nil {
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
	// Compare property "name"
	if lhs, rhs := this.ActivityStreamsName, o.GetActivityStreamsName(); lhs != nil && rhs != nil {
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
	// Compare property "object"
	if lhs, rhs := this.ActivityStreamsObject, o.GetActivityStreamsObject(); lhs != nil && rhs != nil {
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
	// Compare property "preview"
	if lhs, rhs := this.ActivityStreamsPreview, o.GetActivityStreamsPreview(); lhs != nil && rhs != nil {
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
	// Compare property "published"
	if lhs, rhs := this.ActivityStreamsPublished, o.GetActivityStreamsPublished(); lhs != nil && rhs != nil {
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
	// Compare property "replies"
	if lhs, rhs := this.ActivityStreamsReplies, o.GetActivityStreamsReplies(); lhs != nil && rhs != nil {
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
	// Compare property "shares"
	if lhs, rhs := this.ActivityStreamsShares, o.GetActivityStreamsShares(); lhs != nil && rhs != nil {
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
	// Compare property "startTime"
	if lhs, rhs := this.ActivityStreamsStartTime, o.GetActivityStreamsStartTime(); lhs != nil && rhs != nil {
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
	// Compare property "summary"
	if lhs, rhs := this.ActivityStreamsSummary, o.GetActivityStreamsSummary(); lhs != nil && rhs != nil {
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
	// Compare property "tag"
	if lhs, rhs := this.ActivityStreamsTag, o.GetActivityStreamsTag(); lhs != nil && rhs != nil {
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
	// Compare property "to"
	if lhs, rhs := this.ActivityStreamsTo, o.GetActivityStreamsTo(); lhs != nil && rhs != nil {
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
	// Compare property "type"
	if lhs, rhs := this.ActivityStreamsType, o.GetActivityStreamsType(); lhs != nil && rhs != nil {
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
	// Compare property "updated"
	if lhs, rhs := this.ActivityStreamsUpdated, o.GetActivityStreamsUpdated(); lhs != nil && rhs != nil {
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
	// Compare property "url"
	if lhs, rhs := this.ActivityStreamsUrl, o.GetActivityStreamsUrl(); lhs != nil && rhs != nil {
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
func (this ActivityStreamsTombstone) Serialize() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	typeName := "Tombstone"
	if len(this.alias) > 0 {
		typeName = this.alias + ":" + "Tombstone"
	}
	m["type"] = typeName
	// Begin: Serialize known properties
	// Maybe serialize property "altitude"
	if this.ActivityStreamsAltitude != nil {
		if i, err := this.ActivityStreamsAltitude.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsAltitude.Name()] = i
		}
	}
	// Maybe serialize property "attachment"
	if this.ActivityStreamsAttachment != nil {
		if i, err := this.ActivityStreamsAttachment.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsAttachment.Name()] = i
		}
	}
	// Maybe serialize property "attributedTo"
	if this.ActivityStreamsAttributedTo != nil {
		if i, err := this.ActivityStreamsAttributedTo.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsAttributedTo.Name()] = i
		}
	}
	// Maybe serialize property "audience"
	if this.ActivityStreamsAudience != nil {
		if i, err := this.ActivityStreamsAudience.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsAudience.Name()] = i
		}
	}
	// Maybe serialize property "bcc"
	if this.ActivityStreamsBcc != nil {
		if i, err := this.ActivityStreamsBcc.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsBcc.Name()] = i
		}
	}
	// Maybe serialize property "bto"
	if this.ActivityStreamsBto != nil {
		if i, err := this.ActivityStreamsBto.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsBto.Name()] = i
		}
	}
	// Maybe serialize property "cc"
	if this.ActivityStreamsCc != nil {
		if i, err := this.ActivityStreamsCc.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsCc.Name()] = i
		}
	}
	// Maybe serialize property "content"
	if this.ActivityStreamsContent != nil {
		if i, err := this.ActivityStreamsContent.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsContent.Name()] = i
		}
	}
	// Maybe serialize property "context"
	if this.ActivityStreamsContext != nil {
		if i, err := this.ActivityStreamsContext.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsContext.Name()] = i
		}
	}
	// Maybe serialize property "deleted"
	if this.ActivityStreamsDeleted != nil {
		if i, err := this.ActivityStreamsDeleted.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsDeleted.Name()] = i
		}
	}
	// Maybe serialize property "duration"
	if this.ActivityStreamsDuration != nil {
		if i, err := this.ActivityStreamsDuration.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsDuration.Name()] = i
		}
	}
	// Maybe serialize property "endTime"
	if this.ActivityStreamsEndTime != nil {
		if i, err := this.ActivityStreamsEndTime.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsEndTime.Name()] = i
		}
	}
	// Maybe serialize property "formerType"
	if this.ActivityStreamsFormerType != nil {
		if i, err := this.ActivityStreamsFormerType.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsFormerType.Name()] = i
		}
	}
	// Maybe serialize property "generator"
	if this.ActivityStreamsGenerator != nil {
		if i, err := this.ActivityStreamsGenerator.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsGenerator.Name()] = i
		}
	}
	// Maybe serialize property "icon"
	if this.ActivityStreamsIcon != nil {
		if i, err := this.ActivityStreamsIcon.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsIcon.Name()] = i
		}
	}
	// Maybe serialize property "id"
	if this.ActivityStreamsId != nil {
		if i, err := this.ActivityStreamsId.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsId.Name()] = i
		}
	}
	// Maybe serialize property "image"
	if this.ActivityStreamsImage != nil {
		if i, err := this.ActivityStreamsImage.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsImage.Name()] = i
		}
	}
	// Maybe serialize property "inReplyTo"
	if this.ActivityStreamsInReplyTo != nil {
		if i, err := this.ActivityStreamsInReplyTo.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsInReplyTo.Name()] = i
		}
	}
	// Maybe serialize property "likes"
	if this.ActivityStreamsLikes != nil {
		if i, err := this.ActivityStreamsLikes.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsLikes.Name()] = i
		}
	}
	// Maybe serialize property "location"
	if this.ActivityStreamsLocation != nil {
		if i, err := this.ActivityStreamsLocation.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsLocation.Name()] = i
		}
	}
	// Maybe serialize property "mediaType"
	if this.ActivityStreamsMediaType != nil {
		if i, err := this.ActivityStreamsMediaType.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsMediaType.Name()] = i
		}
	}
	// Maybe serialize property "name"
	if this.ActivityStreamsName != nil {
		if i, err := this.ActivityStreamsName.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsName.Name()] = i
		}
	}
	// Maybe serialize property "object"
	if this.ActivityStreamsObject != nil {
		if i, err := this.ActivityStreamsObject.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsObject.Name()] = i
		}
	}
	// Maybe serialize property "preview"
	if this.ActivityStreamsPreview != nil {
		if i, err := this.ActivityStreamsPreview.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsPreview.Name()] = i
		}
	}
	// Maybe serialize property "published"
	if this.ActivityStreamsPublished != nil {
		if i, err := this.ActivityStreamsPublished.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsPublished.Name()] = i
		}
	}
	// Maybe serialize property "replies"
	if this.ActivityStreamsReplies != nil {
		if i, err := this.ActivityStreamsReplies.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsReplies.Name()] = i
		}
	}
	// Maybe serialize property "shares"
	if this.ActivityStreamsShares != nil {
		if i, err := this.ActivityStreamsShares.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsShares.Name()] = i
		}
	}
	// Maybe serialize property "startTime"
	if this.ActivityStreamsStartTime != nil {
		if i, err := this.ActivityStreamsStartTime.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsStartTime.Name()] = i
		}
	}
	// Maybe serialize property "summary"
	if this.ActivityStreamsSummary != nil {
		if i, err := this.ActivityStreamsSummary.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsSummary.Name()] = i
		}
	}
	// Maybe serialize property "tag"
	if this.ActivityStreamsTag != nil {
		if i, err := this.ActivityStreamsTag.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsTag.Name()] = i
		}
	}
	// Maybe serialize property "to"
	if this.ActivityStreamsTo != nil {
		if i, err := this.ActivityStreamsTo.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsTo.Name()] = i
		}
	}
	// Maybe serialize property "type"
	if this.ActivityStreamsType != nil {
		if i, err := this.ActivityStreamsType.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsType.Name()] = i
		}
	}
	// Maybe serialize property "updated"
	if this.ActivityStreamsUpdated != nil {
		if i, err := this.ActivityStreamsUpdated.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsUpdated.Name()] = i
		}
	}
	// Maybe serialize property "url"
	if this.ActivityStreamsUrl != nil {
		if i, err := this.ActivityStreamsUrl.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsUrl.Name()] = i
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

// SetActivityStreamsAltitude sets the "altitude" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsAltitude(i vocab.ActivityStreamsAltitudeProperty) {
	this.ActivityStreamsAltitude = i
}

// SetActivityStreamsAttachment sets the "attachment" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsAttachment(i vocab.ActivityStreamsAttachmentProperty) {
	this.ActivityStreamsAttachment = i
}

// SetActivityStreamsAttributedTo sets the "attributedTo" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsAttributedTo(i vocab.ActivityStreamsAttributedToProperty) {
	this.ActivityStreamsAttributedTo = i
}

// SetActivityStreamsAudience sets the "audience" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsAudience(i vocab.ActivityStreamsAudienceProperty) {
	this.ActivityStreamsAudience = i
}

// SetActivityStreamsBcc sets the "bcc" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsBcc(i vocab.ActivityStreamsBccProperty) {
	this.ActivityStreamsBcc = i
}

// SetActivityStreamsBto sets the "bto" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsBto(i vocab.ActivityStreamsBtoProperty) {
	this.ActivityStreamsBto = i
}

// SetActivityStreamsCc sets the "cc" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsCc(i vocab.ActivityStreamsCcProperty) {
	this.ActivityStreamsCc = i
}

// SetActivityStreamsContent sets the "content" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsContent(i vocab.ActivityStreamsContentProperty) {
	this.ActivityStreamsContent = i
}

// SetActivityStreamsContext sets the "context" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsContext(i vocab.ActivityStreamsContextProperty) {
	this.ActivityStreamsContext = i
}

// SetActivityStreamsDeleted sets the "deleted" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsDeleted(i vocab.ActivityStreamsDeletedProperty) {
	this.ActivityStreamsDeleted = i
}

// SetActivityStreamsDuration sets the "duration" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsDuration(i vocab.ActivityStreamsDurationProperty) {
	this.ActivityStreamsDuration = i
}

// SetActivityStreamsEndTime sets the "endTime" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsEndTime(i vocab.ActivityStreamsEndTimeProperty) {
	this.ActivityStreamsEndTime = i
}

// SetActivityStreamsFormerType sets the "formerType" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsFormerType(i vocab.ActivityStreamsFormerTypeProperty) {
	this.ActivityStreamsFormerType = i
}

// SetActivityStreamsGenerator sets the "generator" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsGenerator(i vocab.ActivityStreamsGeneratorProperty) {
	this.ActivityStreamsGenerator = i
}

// SetActivityStreamsIcon sets the "icon" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsIcon(i vocab.ActivityStreamsIconProperty) {
	this.ActivityStreamsIcon = i
}

// SetActivityStreamsId sets the "id" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsId(i vocab.ActivityStreamsIdProperty) {
	this.ActivityStreamsId = i
}

// SetActivityStreamsImage sets the "image" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsImage(i vocab.ActivityStreamsImageProperty) {
	this.ActivityStreamsImage = i
}

// SetActivityStreamsInReplyTo sets the "inReplyTo" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsInReplyTo(i vocab.ActivityStreamsInReplyToProperty) {
	this.ActivityStreamsInReplyTo = i
}

// SetActivityStreamsLikes sets the "likes" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsLikes(i vocab.ActivityStreamsLikesProperty) {
	this.ActivityStreamsLikes = i
}

// SetActivityStreamsLocation sets the "location" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsLocation(i vocab.ActivityStreamsLocationProperty) {
	this.ActivityStreamsLocation = i
}

// SetActivityStreamsMediaType sets the "mediaType" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsMediaType(i vocab.ActivityStreamsMediaTypeProperty) {
	this.ActivityStreamsMediaType = i
}

// SetActivityStreamsName sets the "name" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsName(i vocab.ActivityStreamsNameProperty) {
	this.ActivityStreamsName = i
}

// SetActivityStreamsObject sets the "object" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsObject(i vocab.ActivityStreamsObjectProperty) {
	this.ActivityStreamsObject = i
}

// SetActivityStreamsPreview sets the "preview" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsPreview(i vocab.ActivityStreamsPreviewProperty) {
	this.ActivityStreamsPreview = i
}

// SetActivityStreamsPublished sets the "published" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsPublished(i vocab.ActivityStreamsPublishedProperty) {
	this.ActivityStreamsPublished = i
}

// SetActivityStreamsReplies sets the "replies" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsReplies(i vocab.ActivityStreamsRepliesProperty) {
	this.ActivityStreamsReplies = i
}

// SetActivityStreamsShares sets the "shares" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsShares(i vocab.ActivityStreamsSharesProperty) {
	this.ActivityStreamsShares = i
}

// SetActivityStreamsStartTime sets the "startTime" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsStartTime(i vocab.ActivityStreamsStartTimeProperty) {
	this.ActivityStreamsStartTime = i
}

// SetActivityStreamsSummary sets the "summary" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsSummary(i vocab.ActivityStreamsSummaryProperty) {
	this.ActivityStreamsSummary = i
}

// SetActivityStreamsTag sets the "tag" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsTag(i vocab.ActivityStreamsTagProperty) {
	this.ActivityStreamsTag = i
}

// SetActivityStreamsTo sets the "to" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsTo(i vocab.ActivityStreamsToProperty) {
	this.ActivityStreamsTo = i
}

// SetActivityStreamsType sets the "type" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsType(i vocab.ActivityStreamsTypeProperty) {
	this.ActivityStreamsType = i
}

// SetActivityStreamsUpdated sets the "updated" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsUpdated(i vocab.ActivityStreamsUpdatedProperty) {
	this.ActivityStreamsUpdated = i
}

// SetActivityStreamsUrl sets the "url" property.
func (this *ActivityStreamsTombstone) SetActivityStreamsUrl(i vocab.ActivityStreamsUrlProperty) {
	this.ActivityStreamsUrl = i
}

// VocabularyURI returns the vocabulary's URI as a string.
func (this ActivityStreamsTombstone) VocabularyURI() string {
	return "https://www.w3.org/TR/activitystreams-vocabulary"
}

// helperJSONLDContext obtains the context uris and their aliases from a property,
// if it is not nil.
func (this ActivityStreamsTombstone) helperJSONLDContext(i jsonldContexter, toMerge map[string]string) map[string]string {
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
