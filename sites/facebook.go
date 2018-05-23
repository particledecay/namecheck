package sites

// Facebook (https://www.facebook.com)
type Facebook struct {
	Site
}

// Name returns the name of the site to be displayed
func (f Facebook) Name() string {
	return "Facebook"
}

// URL is a path for the site with a template for the username
func (f Facebook) URL() string {
	return "https://www.facebook.com/%s"
}

// UserAgent returns an optional string to explicitly provide a User-Agent
func (f Facebook) UserAgent() string {
	return ""
}

// Check queries the Site for `username`
func (f Facebook) Check(username string, ch chan *NameResult) {
	ch <- IfPageNotFound(f, username)
}
