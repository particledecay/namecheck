package sites

// Fortnite (https://www.epicgames.com/fortnite)
type Fortnite struct {
	Site
}

// Name returns the name of the site to be displayed
func (f Fortnite) Name() string {
	return "Fortnite"
}

// URL is a path for the site with a template for the username
func (f Fortnite) URL() string {
	return "https://www.stormshield.one/pvp/stats/%s"
}

// UserAgent returns an optional string to explicitly provide a User-Agent
func (f Fortnite) UserAgent() string {
	return ""
}

// Check queries the Site for `username`
func (f Fortnite) Check(username string, ch chan *NameResult) {
	ch <- IfElementOnPage(f, username, "//div[contains(@class, 'profile__title')]//h2/following-sibling::div[1]/div[contains(text(), 'Not found')]")
}
