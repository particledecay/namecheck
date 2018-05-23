package sites

// GitHub (https://github.com)
type GitHub struct {
	Site
}

// Name returns the name of the site to be displayed
func (g GitHub) Name() string {
	return "GitHub"
}

// URL is a path for the site with a template for the username
func (g GitHub) URL() string {
	return "https://github.com/%s"
}

// UserAgent returns an optional string to explicitly provide a User-Agent
func (g GitHub) UserAgent() string {
	return ""
}

// Check queries the Site for `username`
func (g GitHub) Check(username string, ch chan *NameResult) {
	ch <- IfPageNotFound(g, username)
}
