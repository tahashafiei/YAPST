package main

// Edit these values to personalize your site.
var siteConfig = SiteConfig{
	Name:        "Your Name",
	GitHubURL:   "https://github.com/yourusername",
	LinkedInURL: "https://www.linkedin.com/in/yourusername",
	ResumeFile:  "resume.pdf", // filename inside static/; leave empty to hide the link
}

type SiteConfig struct {
	Name        string
	GitHubURL   string
	LinkedInURL string
	ResumeFile  string
}
