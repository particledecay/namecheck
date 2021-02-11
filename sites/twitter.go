package sites

// Twitter (https://twitter.com)
type Twitter struct {
	Site
}

// Name returns the name of the site to be displayed
func (t Twitter) Name() string {
	return "Twitter"
}

// URL is a path for the site with a template for the username
func (t Twitter) URL() string {
	return "https://twitter.com/%s"
}

// UserAgent returns an optional string to explicitly provide a User-Agent
func (t Twitter) UserAgent() string {
	return ""
}

// Check queries the Site for `username`
func (t Twitter) Check(username string, ch chan *NameResult) {
	ch <- IfElementOnPage(t, username, "//span[contains(text(), 'This account doesn')]")
}
