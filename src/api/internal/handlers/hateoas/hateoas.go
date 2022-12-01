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
