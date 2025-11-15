// Package uslm provides Go types for parsing and generating United States Legislative Markup (USLM) XML documents.
// These types support round-trip XML parsing and JSON serialization without data loss.
package uslm

// LegislativeDocument is the common interface for all legislative document types.
// This includes bills, resolutions, amendments, public laws, etc.
type LegislativeDocument interface {
	// GetDocumentNumber returns the document number (e.g., "32" for S. 32)
	GetDocumentNumber() string

	// GetDocumentType returns the type of document (e.g., "Senate Bill", "House Resolution")
	GetDocumentType() string

	// GetCongress returns the congress number (e.g., "114", "116")
	GetCongress() string

	// GetSession returns the session number (e.g., "1", "2")
	GetSession() string

	// GetTitle returns the full title of the document
	GetTitle() string

	// GetStage returns the document stage (e.g., "Introduced in Senate", "Committee Discharged")
	GetStage() string

	// GetChamber returns the current chamber ("SENATE" or "HOUSE")
	GetChamber() string

	// IsPublic returns true if this is a public (not private) document
	IsPublic() bool

	// GetCitations returns all citable forms of this document
	GetCitations() []string
}

// SponsoredDocument represents documents that can have sponsors and cosponsors.
// This typically includes bills and resolutions but not all document types.
type SponsoredDocument interface {
	// GetSponsors returns all primary sponsors of the document
	GetSponsors() []Sponsor

	// GetCosponsors returns all cosponsors of the document
	GetCosponsors() []Cosponsor
}

// ActionDocument represents documents that have legislative actions.
// Actions describe the history and current status of the document.
type ActionDocument interface {
	// GetActions returns all actions taken on this document
	GetActions() []Action
}

// CommitteeDocument represents documents that reference committees.
type CommitteeDocument interface {
	// GetCommittees returns all committees referenced in this document
	GetCommittees() []Committee
}

// HierarchicalDocument represents documents with hierarchical structure (sections, paragraphs, etc.)
type HierarchicalDocument interface {
	// GetSections returns all top-level sections in the document
	GetSections() []Section
}

// MetadataDocument provides access to Dublin Core and processing metadata.
type MetadataDocument interface {
	// GetCreator returns the document creator (e.g., "United States Senate")
	GetCreator() string

	// GetPublisher returns the publisher
	GetPublisher() string

	// GetLanguage returns the document language code
	GetLanguage() string

	// GetRights returns the copyright/rights statement
	GetRights() string

	// GetProcessedBy returns the tool that processed this document
	GetProcessedBy() string

	// GetProcessedDate returns when the document was processed
	GetProcessedDate() string
}

// AmendmentDocument represents amendment-specific functionality.
type AmendmentDocument interface {
	LegislativeDocument

	// GetAmendmentDegree returns the degree of amendment (e.g., "first", "second")
	GetAmendmentDegree() string
}

// Identifiable represents elements that have identifiers.
type Identifiable interface {
	// GetID returns the unique ID of this element
	GetID() string

	// GetIdentifier returns the logical URL identifier (e.g., "/us/bill/114/s/32/s1")
	GetIdentifier() string
}

// Numbered represents elements that have a designation number.
type Numbered interface {
	// GetNum returns the element's number/designation
	GetNum() string

	// GetNumValue returns the normalized numeric value
	GetNumValue() string
}

// Headed represents elements that have headings.
type Headed interface {
	// GetHeading returns the element's heading text
	GetHeading() string
}

// ContentContainer represents elements that contain legislative text content.
type ContentContainer interface {
	// GetContent returns the main content text
	GetContent() string

	// GetChapeau returns the lead-in text (if any)
	GetChapeau() string
}
