package sites

// DomainDotOrg (https://www.whois.com)
type DomainDotOrg struct {
	Site
}

// Name returns the name of the site to be displayed
func (ddc DomainDotOrg) Name() string {
	return "Domain (.org)"
}

// URL is a path for the site with a template for the username
func (ddc DomainDotOrg) URL() string {
	return "https://www.whois.com/whois/%s.org"
}

// UserAgent returns an optional string to explicitly provide a User-Agent
func (ddc DomainDotOrg) UserAgent() string {
	return ""
}

// Check queries the Site for `username`
func (ddc DomainDotOrg) Check(username string, ch chan *NameResult) {
	ch <- IfElementOnPage(ddc, username, "//div[@id='availableBlk']/div[contains(text(), 'not been registered')]")
}
