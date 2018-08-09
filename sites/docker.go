package sites

// Docker (https://hub.docker.com)
type Docker struct {
	Site
}

// Name returns the name of the site to be displayed
func (g Docker) Name() string {
	return "Docker"
}

// URL is a path for the site with a template for the username
func (g Docker) URL() string {
	return "https://hub.docker.com/u/%s"
}

// UserAgent returns an optional string to explicitly provide a User-Agent
func (g Docker) UserAgent() string {
	return ""
}

// Check queries the Site for `username`
func (g Docker) Check(username string, ch chan *NameResult) {
	ch <- IfElementOnPage(g, username, "//h1[contains(@class, 'NotFound404')][text()=404]")
}
