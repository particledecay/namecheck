package sites

// Instagram (https://instagram.com)
type Instagram struct {
	Site
}

// Name returns the name of the site to be displayed
func (i Instagram) Name() string {
	return "Instagram"
}

// URL is a path for the site with a template for the username
func (i Instagram) URL() string {
	return "https://www.instagram.com/%s/"
}

// UserAgent returns an optional string to explicitly provide a User-Agent
func (i Instagram) UserAgent() string {
	return "User-Agent: Mozilla/5.0 (Linux; Android 4.4; Nexus 5 Build/_BuildID_) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36"
}

// Check queries the Site for `username`
func (i Instagram) Check(username string, ch chan *NameResult) {
	ch <- IfPageNotFound(i, username)
}
