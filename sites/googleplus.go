package sites

// GooglePlus (https://plus.google.com)
type GooglePlus struct {
	Site
}

// Name returns the name of the site to be displayed
func (g GooglePlus) Name() string {
	return "Google+"
}

// URL is a path for the site with a template for the username
func (g GooglePlus) URL() string {
	return "https://plus.google.com/+%s"
}

// UserAgent returns an optional string to explicitly provide a User-Agent
func (g GooglePlus) UserAgent() string {
	return ""
}

// Check queries the Site for `username`
func (g GooglePlus) Check(username string, ch chan *NameResult) {
	ch <- IfPageNotFound(g, username)
}
