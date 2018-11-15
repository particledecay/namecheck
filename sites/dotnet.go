package sites

// DomainDotNet (https://www.whois.com)
type DomainDotNet struct {
	Site
}

// Name returns the name of the site to be displayed
func (ddc DomainDotNet) Name() string {
	return "Domain (.net)"
}

// URL is a path for the site with a template for the username
func (ddc DomainDotNet) URL() string {
	return "https://www.whois.com/whois/%s.net"
}

// UserAgent returns an optional string to explicitly provide a User-Agent
func (ddc DomainDotNet) UserAgent() string {
	return ""
}

// Check queries the Site for `username`
func (ddc DomainDotNet) Check(username string, ch chan *NameResult) {
	ch <- IfElementOnPage(ddc, username, "//div[@id='availableBlk']/div[contains(text(), 'not been registered')]")
}
