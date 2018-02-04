package pub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	postContentTypeHeader     = "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\""
	responseContentTypeHeader = "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\""
	getAcceptHeader           = "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\""
	contentTypeHeader         = "Content-Type"
	acceptHeader              = "Accept"
	publicActivityPub         = "https://www.w3.org/ns/activitystreams#Public"
	publicJsonLD              = "Public"
	publicJsonLDAS            = "as:Public"
)

var alternatives = []string{"application/activity+json"}

func trimAll(s []string) []string {
	var r []string
	for _, e := range s {
		r = append(r, strings.Trim(e, " "))
	}
	return r
}

func headerEqualsOneOf(header string, acceptable []string) bool {
	sanitizedHeader := strings.Join(trimAll(strings.Split(header, ";")), ";")
	for _, v := range acceptable {
		// Remove any number of whitespace after ;'s
		sanitizedV := strings.Join(trimAll(strings.Split(v, ";")), ";")
		if sanitizedHeader == sanitizedV {
			return true
		}
	}
	return false
}

func isActivityPubPost(r *http.Request) bool {
	return r.Method == "POST" && headerEqualsOneOf(r.Header.Get(contentTypeHeader), []string{postContentTypeHeader, contentTypeHeader})
}

func isActivityPubGet(r *http.Request) bool {
	return r.Method == "GET" && headerEqualsOneOf(r.Header.Get(acceptHeader), []string{getAcceptHeader, contentTypeHeader})
}

// isPublic determines if a target is the Public collection as defined in the
// spec, including JSON-LD compliant collections.
func isPublic(s string) bool {
	return s == publicActivityPub || s == publicJsonLD || s == publicJsonLDAS
}

// dereference makes an HTTP GET request to an IRI in order to obtain the
// ActivityStream representation.
func dereference(c *http.Client, u url.URL, agent string) ([]byte, error) {
	// TODO: (section 7.1) The server MUST dereference the collection (with the user's credentials)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(acceptHeader, getAcceptHeader)
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("Date", time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05")+" GMT")
	req.Header.Add("User-Agent", fmt.Sprintf("%s (go-fed ActivityPub)", agent))
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Request to %s failed (%d): %s", u.String(), resp.StatusCode, resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}

// postToOutbox will attempt to send a POST request to the given URL with the
// body set to the provided bytes.
func postToOutbox(c *http.Client, b []byte, to url.URL, agent string) error {
	byteCopy := make([]byte, 0, len(b))
	copy(b, byteCopy)
	buf := bytes.NewBuffer(byteCopy)
	req, err := http.NewRequest("POST", to.String(), buf)
	if err != nil {
		return err
	}
	req.Header.Add(contentTypeHeader, postContentTypeHeader)
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("Date", time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05")+" GMT")
	req.Header.Add("User-Agent", fmt.Sprintf("%s (go-fed ActivityPub)", agent))
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Request to %s failed (%d): %s", to.String(), resp.StatusCode, resp.Status)
	}
	return nil
}

// wrapInCreate will automatically wrap the provided object in a Create
// activity. This will copy over the 'to', 'bto', 'cc', 'bcc', and 'audience'
// properties. It will also copy over the published time if present.
func (f *federator) wrapInCreate(o vocab.ObjectType, actor url.URL) *vocab.Create {
	c := &vocab.Create{}
	c.AddObject(o)
	c.AddActorIRI(actor)
	if o.IsPublished() {
		c.SetPublished(o.GetPublished())
	}
	for i := 0; i < o.ToLen(); i++ {
		if o.IsToObject(i) {
			c.AddToObject(o.GetToObject(i))
		} else if o.IsToLink(i) {
			c.AddToLink(o.GetToLink(i))
		} else if o.IsToIRI(i) {
			c.AddToIRI(o.GetToIRI(i))
		}
	}
	for i := 0; i < o.BtoLen(); i++ {
		if o.IsBtoObject(i) {
			c.AddBtoObject(o.GetBtoObject(i))
		} else if o.IsBtoLink(i) {
			c.AddBtoLink(o.GetBtoLink(i))
		} else if o.IsBtoIRI(i) {
			c.AddBtoIRI(o.GetBtoIRI(i))
		}
	}
	for i := 0; i < o.CcLen(); i++ {
		if o.IsCcObject(i) {
			c.AddCcObject(o.GetCcObject(i))
		} else if o.IsCcLink(i) {
			c.AddCcLink(o.GetCcLink(i))
		} else if o.IsCcIRI(i) {
			c.AddCcIRI(o.GetCcIRI(i))
		}
	}
	for i := 0; i < o.BccLen(); i++ {
		if o.IsBccObject(i) {
			c.AddBccObject(o.GetBccObject(i))
		} else if o.IsBccLink(i) {
			c.AddBccLink(o.GetBccLink(i))
		} else if o.IsBccIRI(i) {
			c.AddBccIRI(o.GetBccIRI(i))
		}
	}
	for i := 0; i < o.AudienceLen(); i++ {
		if o.IsAudienceObject(i) {
			c.AddAudienceObject(o.GetAudienceObject(i))
		} else if o.IsAudienceLink(i) {
			c.AddAudienceLink(o.GetAudienceLink(i))
		} else if o.IsAudienceIRI(i) {
			c.AddAudienceIRI(o.GetAudienceIRI(i))
		}
	}
	return c
}

// TODO: (Section 7) HTTP caching mechanisms [RFC7234] SHOULD be respected when appropriate, both when receiving responses from other servers as well as sending responses to other servers.

// prepare takes a DeliverableObject and returns a list of the proper recipient
// target URIs. Additionally, the DeliverableObject will have any hidden
// hidden recipients ("bto" and "bcc") stripped from it.
func (c *federator) prepare(o DeliverableObject) ([]url.URL, error) {
	// Get inboxes of recipients
	var r []url.URL
	r = append(r, getToIRIs(o)...)
	r = append(r, getBToIRIs(o)...)
	r = append(r, getCcIRIs(o)...)
	r = append(r, getBccIRIs(o)...)
	r = append(r, getAudienceIRIs(o)...)
	// TODO: Support delivery to shared inbox
	// 1. When an object is being delivered to the originating actor's
	// followers, a server MAY reduce the number of receiving actors
	// delivered to by identifying all followers which share the same
	// sharedInbox who would otherwise be individual recipients and instead
	// deliver objects to said sharedInbox.
	// 2. If an object is addressed to the Public special collection, a
	// server MAY deliver that object to all known sharedInbox endpoints on
	// the network.
	r = filterURLs(r, isPublic)
	receiverActors, err := c.resolveInboxes(r, 0, c.MaxDeliveryDepth)
	if err != nil {
		return nil, err
	}
	targets := getInboxes(receiverActors)
	// Get inboxes of sender(s)
	senderActors, err := c.resolveInboxes(getActorsAttributedToURI(o), 0, c.MaxDeliveryDepth)
	if err != nil {
		return nil, err
	}
	ignore := getInboxes(senderActors)
	// Post-processing
	r = dedupeIRIs(targets, ignore)
	stripHiddenRecipients(o)
	return r, nil
}

// resolveInboxes takes a list of Actor id URIs and returns them as concrete
// instances of ActorObject. It applies recursively when it encounters a target
// that is a Collection or OrderedCollection.
func (c *federator) resolveInboxes(r []url.URL, depth int, max int) ([]ActorObject, error) {
	if depth >= max {
		return nil, nil
	}
	a := make([]ActorObject, 0, len(r))
	for _, u := range r {
		// Do not retry here -- if a dereference fails, then fail the
		// entire delivery.
		resp, err := dereference(c.Client, u, c.Agent)
		if err != nil {
			return nil, err
		}
		var m map[string]interface{}
		if err = json.Unmarshal(resp, &m); err != nil {
			return nil, err
		}
		var actor ActorObject
		var co *streams.Collection
		var oc *streams.OrderedCollection
		var cp *streams.CollectionPage
		var ocp *streams.OrderedCollectionPage
		// Set AT MOST one of: co, oc, cp, ocp
		// If none of these are set, attempt to use actor
		if err = toActorCollectionResolver(&actor, &co, &oc, &cp, &ocp).Deserialize(m); err != nil {
			return nil, err
		}
		// If a recipient is a Collection or OrderedCollection, then the
		// server MUST dereference the collection. Note that this also
		// applies to CollectionPage and OrderedCollectionPage.
		var uris []url.URL
		if co != nil {
			uris := getURIsInItemer(co.Raw())
			actors, err := c.resolveInboxes(uris, depth+1, max)
			if err != nil {
				return nil, err
			}
			a = append(a, actors...)
		} else if oc != nil {
			uris := getURIsInOrderedItemer(oc.Raw())
			actors, err := c.resolveInboxes(uris, depth+1, max)
			if err != nil {
				return nil, err
			}
			a = append(a, actors...)
		} else if cp != nil {
			cb := func(c vocab.CollectionPageType) error {
				uris = getURIsInItemer(c)
				return nil
			}
			err := doForCollectionPage(c.Client, c.Agent, cb, cp.Raw())
			if err != nil {
				return nil, err
			}
			actors, err := c.resolveInboxes(uris, depth+1, max)
			if err != nil {
				return nil, err
			}
			a = append(a, actors...)
		} else if ocp != nil {
			cb := func(c vocab.OrderedCollectionPageType) error {
				uris = getURIsInOrderedItemer(c)
				return nil
			}
			err := doForOrderedCollectionPage(c.Client, c.Agent, cb, ocp.Raw())
			if err != nil {
				return nil, err
			}
			actors, err := c.resolveInboxes(uris, depth+1, max)
			if err != nil {
				return nil, err
			}
			a = append(a, actors...)
		} else if actor != nil {
			a = append(a, actor)
		}
	}
	return a, nil
}

// getInboxes extracts the 'inbox' IRIs from actors.
func getInboxes(a []ActorObject) []url.URL {
	var u []url.URL
	for _, actor := range a {
		if actor.HasInbox() {
			u = append(u, actor.GetInbox())
		}
	}
	return u
}

// getActorAttributedToURI attempts to find the URIs for the "actor" and
// "attributedTo" originators on the object.
func getActorsAttributedToURI(a ActorObject) []url.URL {
	var u []url.URL
	for i := 0; i < a.AttributedToLen(); i++ {
		if a.IsAttributedToObject(i) {
			obj := a.GetAttributedToObject(i)
			if obj.HasId() {
				u = append(u, obj.GetId())
			}
		} else if a.IsAttributedToLink(i) {
			l := a.GetAttributedToLink(i)
			if l.HasHref() {
				u = append(u, l.GetHref())
			}
		} else if a.IsAttributedToIRI(i) {
			u = append(u, a.GetAttributedToIRI(i))
		}
	}
	for i := 0; i < a.ActorLen(); i++ {
		if a.IsActorObject(i) {
			obj := a.GetActorObject(i)
			if obj.HasId() {
				u = append(u, obj.GetId())
			}
		} else if a.IsActorLink(i) {
			l := a.GetActorLink(i)
			if l.HasHref() {
				u = append(u, l.GetHref())
			}
		} else if a.IsActorIRI(i) {
			u = append(u, a.GetActorIRI(i))
		}
	}
	return u
}

// stripHiddenRecipients removes "bto" and "bcc" from the DeliverableObject.
// Note that this requirement of the specification is under "Section 6: Client
// to Server Interactions", the Social API, and not the Federative API.
func stripHiddenRecipients(o DeliverableObject) {
	for o.BtoLen() > 0 {
		if o.IsBtoObject(0) {
			o.RemoveBtoObject(0)
		} else if o.IsBtoLink(0) {
			o.RemoveBtoLink(0)
		} else if o.IsBtoIRI(0) {
			o.RemoveBtoIRI(0)
		}
	}
	for o.BccLen() > 0 {
		if o.IsBccObject(0) {
			o.RemoveBccObject(0)
		} else if o.IsBccLink(0) {
			o.RemoveBccLink(0)
		} else if o.IsBtoIRI(0) {
			o.RemoveBccIRI(0)
		}
	}
}

// dedupeIRIs will deduplicate final inbox IRIs. The ignore list is applied to
// the final list
func dedupeIRIs(recipients, ignored []url.URL) (out []url.URL) {
	ignoredMap := make(map[url.URL]bool, len(ignored))
	for _, elem := range ignored {
		ignoredMap[elem] = true
	}
	outMap := make(map[url.URL]bool, len(recipients))
	for k, _ := range outMap {
		if !ignoredMap[k] {
			out = append(out, k)
		}
	}
	return
}

// dedupeOrderedItems will deduplicate the 'orderedItems' within an ordered
// collection type. Deduplication happens by simply examining the 'id'.
func (f *federator) dedupeOrderedItems(oc vocab.OrderedCollectionType) (vocab.OrderedCollectionType, error) {
	i := 0
	seen := make(map[string]bool, oc.OrderedItemsLen())
	for i < oc.OrderedItemsLen() {
		var id string
		var removeFn func(int)
		if oc.IsOrderedItemsObject(i) {
			removeFn = oc.RemoveOrderedItemsObject
			iri := oc.GetOrderedItemsObject(i).GetId()
			pIri := &iri
			id = pIri.String()
		} else if oc.IsOrderedItemsLink(i) {
			removeFn = oc.RemoveOrderedItemsLink
			iri := oc.GetOrderedItemsLink(i).GetId()
			pIri := &iri
			id = pIri.String()
		} else if oc.IsOrderedItemsIRI(i) {
			removeFn = oc.RemoveOrderedItemsIRI
			b, err := dereference(f.Client, oc.GetOrderedItemsIRI(i), f.Agent)
			var m map[string]interface{}
			if err := json.Unmarshal(b, &m); err != nil {
				return oc, err
			}
			var iri url.URL
			var hasIri bool
			if err = toIdResolver(&hasIri, &iri).Deserialize(m); err != nil {
				return oc, err
			}
			pIri := &iri
			id = pIri.String()
		}
		if seen[id] {
			removeFn(i)
		} else {
			seen[id] = true
			i++
		}
	}
	return oc, nil
}

// filterURLs removes urls whose strings match the provided filter
func filterURLs(u []url.URL, fn func(s string) bool) []url.URL {
	i := 0
	for i < len(u) {
		if fn(u[i].String()) {
			u = append(u[:i], u[i+1:]...)
		} else {
			i++
		}
	}
	return u
}

func getToIRIs(o DeliverableObject) []url.URL {
	var r []url.URL
	for i := 0; i < o.ToLen(); i++ {
		if o.IsToObject(i) {
			obj := o.GetToObject(i)
			if obj.HasId() {
				r = append(r, obj.GetId())
			}
		} else if o.IsToLink(i) {
			l := o.GetToLink(i)
			if l.HasHref() {
				r = append(r, l.GetHref())
			}
		} else if o.IsToIRI(i) {
			r = append(r, o.GetToIRI(i))
		}
	}
	return r
}

func getBToIRIs(o DeliverableObject) []url.URL {
	var r []url.URL
	for i := 0; i < o.BtoLen(); i++ {
		if o.IsBtoObject(i) {
			obj := o.GetBtoObject(i)
			if obj.HasId() {
				r = append(r, obj.GetId())
			}
		} else if o.IsBtoLink(i) {
			l := o.GetBtoLink(i)
			if l.HasHref() {
				r = append(r, l.GetHref())
			}
		} else if o.IsBtoIRI(i) {
			r = append(r, o.GetBtoIRI(i))
		}
	}
	return r
}

func getCcIRIs(o DeliverableObject) []url.URL {
	var r []url.URL
	for i := 0; i < o.CcLen(); i++ {
		if o.IsCcObject(i) {
			obj := o.GetCcObject(i)
			if obj.HasId() {
				r = append(r, obj.GetId())
			}
		} else if o.IsCcLink(i) {
			l := o.GetCcLink(i)
			if l.HasHref() {
				r = append(r, l.GetHref())
			}
		} else if o.IsCcIRI(i) {
			r = append(r, o.GetCcIRI(i))
		}
	}
	return r
}

func getBccIRIs(o DeliverableObject) []url.URL {
	var r []url.URL
	for i := 0; i < o.BccLen(); i++ {
		if o.IsBccObject(i) {
			obj := o.GetBccObject(i)
			if obj.HasId() {
				r = append(r, obj.GetId())
			}
		} else if o.IsBccLink(i) {
			l := o.GetBccLink(i)
			if l.HasHref() {
				r = append(r, l.GetHref())
			}
		} else if o.IsBccIRI(i) {
			r = append(r, o.GetBccIRI(i))
		}
	}
	return r
}

func getAudienceIRIs(o DeliverableObject) []url.URL {
	var r []url.URL
	for i := 0; i < o.AudienceLen(); i++ {
		if o.IsAudienceObject(i) {
			obj := o.GetAudienceObject(i)
			if obj.HasId() {
				r = append(r, obj.GetId())
			}
		} else if o.IsAudienceLink(i) {
			l := o.GetAudienceLink(i)
			if l.HasHref() {
				r = append(r, l.GetHref())
			}
		} else if o.IsAudienceIRI(i) {
			r = append(r, o.GetAudienceIRI(i))
		}
	}
	return r
}

// doForCollectionPage applies a function over a collection and its subsequent
// pages recursively. It returns the first non-nil error it encounters.
func doForCollectionPage(h *http.Client, agent string, cb func(c vocab.CollectionPageType) error, c vocab.CollectionPageType) error {
	err := cb(c)
	if err != nil {
		return err
	}
	if c.IsNextCollectionPage() {
		// Handle this one weird trick that other peers HATE federating
		// with.
		return doForCollectionPage(h, agent, cb, c.GetNextCollectionPage())
	} else if c.IsNextLink() {
		l := c.GetNextLink()
		if l.HasHref() {
			u := l.GetHref()
			resp, err := dereference(h, u, agent)
			if err != nil {
				return err
			}
			var m map[string]interface{}
			err = json.Unmarshal(resp, &m)
			if err != nil {
				return err
			}
			next, err := toCollectionPage(m)
			if err != nil {
				return err
			}
			if next != nil {
				return doForCollectionPage(h, agent, cb, next.Raw())
			}
		}
	} else if c.IsNextIRI() {
		u := c.GetNextIRI()
		resp, err := dereference(h, u, agent)
		if err != nil {
			return err
		}
		var m map[string]interface{}
		err = json.Unmarshal(resp, &m)
		if err != nil {
			return err
		}
		next, err := toCollectionPage(m)
		if err != nil {
			return err
		}
		if next != nil {
			return doForCollectionPage(h, agent, cb, next.Raw())
		}
	}
	return nil
}

// doForOrderedCollectionPage applies a function over a collection and its
// subsequent pages recursively. It returns the first non-nil error it
// encounters.
func doForOrderedCollectionPage(h *http.Client, agent string, cb func(c vocab.OrderedCollectionPageType) error, c vocab.OrderedCollectionPageType) error {
	err := cb(c)
	if err != nil {
		return err
	}
	if c.IsNextOrderedCollectionPage() {
		// Handle this one weird trick that other peers HATE federating
		// with.
		return doForOrderedCollectionPage(h, agent, cb, c.GetNextOrderedCollectionPage())
	} else if c.IsNextLink() {
		l := c.GetNextLink()
		if l.HasHref() {
			u := l.GetHref()
			resp, err := dereference(h, u, agent)
			if err != nil {
				return err
			}
			var m map[string]interface{}
			err = json.Unmarshal(resp, &m)
			if err != nil {
				return err
			}
			next, err := toOrderedCollectionPage(m)
			if err != nil {
				return err
			}
			if next != nil {
				return doForOrderedCollectionPage(h, agent, cb, next.Raw())
			}
		}
	} else if c.IsNextIRI() {
		u := c.GetNextIRI()
		resp, err := dereference(h, u, agent)
		if err != nil {
			return err
		}
		var m map[string]interface{}
		err = json.Unmarshal(resp, &m)
		if err != nil {
			return err
		}
		next, err := toOrderedCollectionPage(m)
		if err != nil {
			return err
		}
		if next != nil {
			return doForOrderedCollectionPage(h, agent, cb, next.Raw())
		}
	}
	return nil
}

type itemer interface {
	ItemsLen() (l int)
	IsItemsObject(index int) (ok bool)
	GetItemsObject(index int) (v vocab.ObjectType)
	IsItemsLink(index int) (ok bool)
	GetItemsLink(index int) (v vocab.LinkType)
	IsItemsIRI(index int) (ok bool)
	GetItemsIRI(index int) (v url.URL)
}

// getURIsInItemer will extract 'items' from the provided Collection or
// CollectionPage.
func getURIsInItemer(i itemer) []url.URL {
	var u []url.URL
	for j := 0; j < i.ItemsLen(); j++ {
		if i.IsItemsObject(j) {
			obj := i.GetItemsObject(j)
			if obj.HasId() {
				u = append(u, obj.GetId())
			}
		} else if i.IsItemsLink(j) {
			l := i.GetItemsLink(j)
			if l.HasHref() {
				u = append(u, l.GetHref())
			}
		} else if i.IsItemsIRI(j) {
			u = append(u, i.GetItemsIRI(j))
		}
	}
	return u
}

type orderedItemer interface {
	OrderedItemsLen() (l int)
	IsOrderedItemsObject(index int) (ok bool)
	GetOrderedItemsObject(index int) (v vocab.ObjectType)
	IsOrderedItemsLink(index int) (ok bool)
	GetOrderedItemsLink(index int) (v vocab.LinkType)
	IsOrderedItemsIRI(index int) (ok bool)
	GetOrderedItemsIRI(index int) (v url.URL)
}

// getURIsInOrderedItemer will extract 'items' from the provided
// OrderedCollection or OrderedCollectionPage.
func getURIsInOrderedItemer(i orderedItemer) []url.URL {
	var u []url.URL
	for j := 0; j < i.OrderedItemsLen(); j++ {
		if i.IsOrderedItemsObject(j) {
			obj := i.GetOrderedItemsObject(j)
			if obj.HasId() {
				u = append(u, obj.GetId())
			}
		} else if i.IsOrderedItemsLink(j) {
			l := i.GetOrderedItemsLink(j)
			if l.HasHref() {
				u = append(u, l.GetHref())
			}
		} else if i.IsOrderedItemsIRI(j) {
			u = append(u, i.GetOrderedItemsIRI(j))
		}
	}
	return u
}

// TODO: Move this to vocab package.
var activityTypes = []string{"Accept", "Add", "Announce", "Arrive", "Block", "Create", "Delete", "Dislike", "Flag", "Follow", "Ignore", "Invite", "Join", "Leave", "Like", "Listen", "Move", "Offer", "Question", "Reject", "Read", "Remove", "TentativeReject", "TentativeAccept", "Travel", "Undo", "Update", "View"}

func isActivityType(t Typer) bool {
	hasType := make(map[string]bool, 1)
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			hasType[s] = true
		}
	}
	for _, t := range activityTypes {
		if hasType[t] {
			return true
		}
	}
	return false
}
