package sites

// DomainDotCom (https://www.whois.com)
type DomainDotCom struct {
	Site
}

// Name returns the name of the site to be displayed
func (ddc DomainDotCom) Name() string {
	return "Domain (.com)"
}

// URL is a path for the site with a template for the username
func (ddc DomainDotCom) URL() string {
	return "https://www.whois.com/whois/%s.com"
}

// UserAgent returns an optional string to explicitly provide a User-Agent
func (ddc DomainDotCom) UserAgent() string {
	return ""
}

// Check queries the Site for `username`
func (ddc DomainDotCom) Check(username string, ch chan *NameResult) {
	ch <- IfElementOnPage(ddc, username, "//div[@id='availableBlk']/div[contains(text(), 'not been registered')]")
}
