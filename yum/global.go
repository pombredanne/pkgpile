package yum

type Package struct {
	Version      Version `xml:"version"`
	Architecture string  `xml:"arch,attr"`
	Pkgid        string  `xml:"pkgid,attr"`
	Name         string  `xml:"name,attr"`
}

type Version struct {
	Epoch   int    `xml:"epoch,attr"`
	Version string `xml:"ver,attr"`
	Release string `xml:"rel,attr"`
}

type Checksum struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
	Pkgid string `xml:"pkgid,attr,omitempty"`
}

type Location struct {
	Href string `xml:"href,attr"`
}

type File struct {
	Type  string `xml:"type,attr,omitempty"`
	Value string `xml:",chardata"`
}