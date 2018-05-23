package sites

// Reddit (https://reddit.com)
type Reddit struct {
	Site
}

// Name returns the name of the site to be displayed
func (r Reddit) Name() string {
	return "Reddit"
}

// URL is a path for the site with a template for the username
func (r Reddit) URL() string {
	return "https://www.reddit.com/user/%s"
}

// UserAgent returns an optional string to explicitly provide a User-Agent
func (r Reddit) UserAgent() string {
	return ""
}

// Check queries the Site for `username`
func (r Reddit) Check(username string, ch chan *NameResult) {
	ch <- IfElementOnPage(r, username, "//h1[text()='page not found']")
}
