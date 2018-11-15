package sites

// DomainDotIo (https://www.whois.com)
type DomainDotIo struct {
	Site
}

// Name returns the name of the site to be displayed
func (ddc DomainDotIo) Name() string {
	return "Domain (.io)"
}

// URL is a path for the site with a template for the username
func (ddc DomainDotIo) URL() string {
	return "https://www.whois.com/whois/%s.io"
}

// UserAgent returns an optional string to explicitly provide a User-Agent
func (ddc DomainDotIo) UserAgent() string {
	return ""
}

// Check queries the Site for `username`
func (ddc DomainDotIo) Check(username string, ch chan *NameResult) {
	ch <- IfElementOnPage(ddc, username, "//div[@id='availableBlk']//tr[1]/td[contains(text(), 'is available')]")
}
