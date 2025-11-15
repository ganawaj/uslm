package uslm

import "encoding/xml"

// Meta represents the metadata section for bills and resolutions.
// This contains machine-processable information that is not published with the document.
type Meta struct {
	XMLName xml.Name `xml:"meta" json:"-"`

	// Dublin Core metadata
	DCTitle     string `xml:"http://purl.org/dc/elements/1.1/ title" json:"dcTitle"`
	DCType      string `xml:"http://purl.org/dc/elements/1.1/ type" json:"dcType"`
	DCCreator   string `xml:"http://purl.org/dc/elements/1.1/ creator" json:"dcCreator,omitempty"`
	DCPublisher string `xml:"http://purl.org/dc/elements/1.1/ publisher" json:"dcPublisher,omitempty"`
	DCFormat    string `xml:"http://purl.org/dc/elements/1.1/ format" json:"dcFormat,omitempty"`
	DCLanguage  string `xml:"http://purl.org/dc/elements/1.1/ language" json:"dcLanguage,omitempty"`
	DCRights    string `xml:"http://purl.org/dc/elements/1.1/ rights" json:"dcRights,omitempty"`

	// Document identifiers
	DocNumber      string   `xml:"docNumber" json:"docNumber"`
	CitableAs      []string `xml:"citableAs" json:"citableAs"`
	DocStage       string   `xml:"docStage" json:"docStage"`
	CurrentChamber string   `xml:"currentChamber,omitempty" json:"currentChamber,omitempty"`

	// Congressional session info
	Congress      string `xml:"congress" json:"congress"`
	Session       string `xml:"session" json:"session"`
	PublicPrivate string `xml:"publicPrivate" json:"publicPrivate"`

	// Processing info
	ProcessedBy   string `xml:"processedBy,omitempty" json:"processedBy,omitempty"`
	ProcessedDate string `xml:"processedDate,omitempty" json:"processedDate,omitempty"`

	// Related documents
	RelatedDocuments []RelatedDocument `xml:"relatedDocument" json:"relatedDocuments,omitempty"`

	// Optional fields
	PopularName string `xml:"popularName,omitempty" json:"popularName,omitempty"`
}

// AmendMeta represents the metadata section for amendment documents.
// Similar to Meta but includes amendment-specific fields.
type AmendMeta struct {
	XMLName xml.Name `xml:"amendMeta" json:"-"`

	// Dublin Core metadata
	DCTitle     string `xml:"http://purl.org/dc/elements/1.1/ title" json:"dcTitle"`
	DCType      string `xml:"http://purl.org/dc/elements/1.1/ type" json:"dcType"`
	DCCreator   string `xml:"http://purl.org/dc/elements/1.1/ creator" json:"dcCreator,omitempty"`
	DCPublisher string `xml:"http://purl.org/dc/elements/1.1/ publisher" json:"dcPublisher,omitempty"`
	DCFormat    string `xml:"http://purl.org/dc/elements/1.1/ format" json:"dcFormat,omitempty"`
	DCLanguage  string `xml:"http://purl.org/dc/elements/1.1/ language" json:"dcLanguage,omitempty"`
	DCRights    string `xml:"http://purl.org/dc/elements/1.1/ rights" json:"dcRights,omitempty"`

	// Document identifiers
	DocNumber      string   `xml:"docNumber" json:"docNumber"`
	CitableAs      []string `xml:"citableAs" json:"citableAs"`
	DocStage       string   `xml:"docStage" json:"docStage"`
	CurrentChamber string   `xml:"currentChamber,omitempty" json:"currentChamber,omitempty"`

	// Amendment-specific
	AmendDegree string `xml:"amendDegree,omitempty" json:"amendDegree,omitempty"`

	// Congressional session info
	Congress      string `xml:"congress" json:"congress"`
	Session       string `xml:"session" json:"session"`
	PublicPrivate string `xml:"publicPrivate" json:"publicPrivate"`

	// Processing info
	ProcessedBy   string `xml:"processedBy,omitempty" json:"processedBy,omitempty"`
	ProcessedDate string `xml:"processedDate,omitempty" json:"processedDate,omitempty"`
}

// RelatedDocument represents a reference to another related document (e.g., committee report).
type RelatedDocument struct {
	XMLName xml.Name `xml:"relatedDocument" json:"-"`
	Role    string   `xml:"role,attr,omitempty" json:"role,omitempty"`
	Href    string   `xml:"href,attr,omitempty" json:"href,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}
