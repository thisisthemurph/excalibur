// Package hateoas hypermedia as the engine of application state
package hateoas

// Hateoas interface
type Hateoas interface {
	Make() H
}

// H the hateoas object
type H struct {
	Links []L `json:"links,onemitempty"`
}

// L a hateoas link
type L struct {
	Href string `json:"href,omitempty"`
	Rel  string `json:"rel,onemitempty"`
	Type string `json:"type,onemitempty"`
}

// Make returns an empty Hateoas object
func Make() H {
	return H{}
}

// WithLink chain method to add a link to the given Hateoas object
func (h H) WithLink(link L) H {
	links := append(h.Links, link)
	return H{Links: links}
}

// WithGetLink chain method to add a link with type `GET` to the given Hateoas object
func (h H) WithGetLink(link L) H {
	return h.WithLink(L{Href: link.Href, Rel: link.Rel, Type: "GET"})
}

// WithPutLink chain method to add a link with type `PUT` to the given Hateoas object
func (h H) WithPutLink(link L) H {
	return h.WithLink(L{Href: link.Href, Rel: link.Rel, Type: "PUT"})
}

// WithPostLink chain method to add a link with type `POST` to the given Hateoas object
func (h H) WithPostLink(link L) H {
	return h.WithLink(L{Href: link.Href, Rel: link.Rel, Type: "POST"})
}

// WithDeleteLink chain method to add a link with type `DELETE` to the given Hateoas object
func (h H) WithDeleteLink(link L) H {
	return h.WithLink(L{Href: link.Href, Rel: link.Rel, Type: "DELETE"})
}
