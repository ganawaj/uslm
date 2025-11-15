package uslm

import "encoding/xml"

// Bill represents a bill document (Senate or House bill).
type Bill struct {
	XMLName xml.Name `xml:"bill" json:"-"`

	// XML namespace declarations (important for round-trip preservation)
	XMLNS           string `xml:"xmlns,attr" json:"xmlns"`
	XMLNSDC         string `xml:"xmlns dc,attr" json:"xmlnsDC,omitempty"`
	XMLNSHTML       string `xml:"xmlns html,attr" json:"xmlnsHTML,omitempty"`
	XMLNSUSLM       string `xml:"xmlns uslm,attr" json:"xmlnsUSLM,omitempty"`
	XMLNSXSI        string `xml:"xmlns xsi,attr" json:"xmlnsXSI,omitempty"`
	XSISchemaLocation string `xml:"xsi schemaLocation,attr" json:"xsiSchemaLocation,omitempty"`
	XMLLang         string `xml:"xml lang,attr" json:"xmlLang,omitempty"`

	// Document sections
	Meta    *Meta    `xml:"meta" json:"meta"`
	Preface *Preface `xml:"preface" json:"preface,omitempty"`
	Main    *Main    `xml:"main" json:"main,omitempty"`

	// End marker
	EndMarker string `xml:"endMarker,omitempty" json:"endMarker,omitempty"`
}

// Ensure Bill implements all relevant interfaces
var (
	_ LegislativeDocument = (*Bill)(nil)
	_ SponsoredDocument   = (*Bill)(nil)
	_ ActionDocument      = (*Bill)(nil)
	_ CommitteeDocument   = (*Bill)(nil)
	_ HierarchicalDocument = (*Bill)(nil)
	_ MetadataDocument    = (*Bill)(nil)
)

// GetDocumentNumber returns the bill number.
func (b *Bill) GetDocumentNumber() string {
	if b.Meta != nil {
		return b.Meta.DocNumber
	}
	return ""
}

// GetDocumentType returns the document type.
func (b *Bill) GetDocumentType() string {
	if b.Meta != nil {
		return b.Meta.DCType
	}
	return ""
}

// GetCongress returns the congress number.
func (b *Bill) GetCongress() string {
	if b.Meta != nil {
		return b.Meta.Congress
	}
	return ""
}

// GetSession returns the session number.
func (b *Bill) GetSession() string {
	if b.Meta != nil {
		return b.Meta.Session
	}
	return ""
}

// GetTitle returns the document title.
func (b *Bill) GetTitle() string {
	if b.Meta != nil {
		return b.Meta.DCTitle
	}
	return ""
}

// GetStage returns the document stage.
func (b *Bill) GetStage() string {
	if b.Meta != nil {
		return b.Meta.DocStage
	}
	return ""
}

// GetChamber returns the current chamber.
func (b *Bill) GetChamber() string {
	if b.Meta != nil {
		return b.Meta.CurrentChamber
	}
	return ""
}

// IsPublic returns true if this is a public bill.
func (b *Bill) IsPublic() bool {
	if b.Meta != nil {
		return b.Meta.PublicPrivate == "public"
	}
	return false
}

// GetCitations returns all citable forms.
func (b *Bill) GetCitations() []string {
	if b.Meta != nil {
		return b.Meta.CitableAs
	}
	return nil
}

// GetSponsors returns all primary sponsors.
func (b *Bill) GetSponsors() []Sponsor {
	var sponsors []Sponsor
	if b.Preface != nil {
		for _, action := range b.Preface.Actions {
			if action.ActionDescription != nil {
				sponsors = append(sponsors, action.ActionDescription.Sponsors...)
			}
		}
	}
	return sponsors
}

// GetCosponsors returns all cosponsors.
func (b *Bill) GetCosponsors() []Cosponsor {
	var cosponsors []Cosponsor
	if b.Preface != nil {
		for _, action := range b.Preface.Actions {
			if action.ActionDescription != nil {
				cosponsors = append(cosponsors, action.ActionDescription.Cosponsors...)
			}
		}
	}
	return cosponsors
}

// GetActions returns all legislative actions.
func (b *Bill) GetActions() []Action {
	if b.Preface != nil {
		return b.Preface.Actions
	}
	return nil
}

// GetCommittees returns all referenced committees.
func (b *Bill) GetCommittees() []Committee {
	var committees []Committee
	if b.Preface != nil {
		for _, action := range b.Preface.Actions {
			if action.ActionDescription != nil {
				committees = append(committees, action.ActionDescription.Committees...)
			}
		}
	}
	return committees
}

// GetSections returns all top-level sections.
func (b *Bill) GetSections() []Section {
	if b.Main != nil {
		return b.Main.Sections
	}
	return nil
}

// GetCreator returns the document creator.
func (b *Bill) GetCreator() string {
	if b.Meta != nil {
		return b.Meta.DCCreator
	}
	return ""
}

// GetPublisher returns the publisher.
func (b *Bill) GetPublisher() string {
	if b.Meta != nil {
		return b.Meta.DCPublisher
	}
	return ""
}

// GetLanguage returns the language code.
func (b *Bill) GetLanguage() string {
	if b.Meta != nil {
		return b.Meta.DCLanguage
	}
	return ""
}

// GetRights returns the rights statement.
func (b *Bill) GetRights() string {
	if b.Meta != nil {
		return b.Meta.DCRights
	}
	return ""
}

// GetProcessedBy returns the processing tool.
func (b *Bill) GetProcessedBy() string {
	if b.Meta != nil {
		return b.Meta.ProcessedBy
	}
	return ""
}

// GetProcessedDate returns the processing date.
func (b *Bill) GetProcessedDate() string {
	if b.Meta != nil {
		return b.Meta.ProcessedDate
	}
	return ""
}

// Resolution represents a resolution document (simple, joint, or concurrent).
type Resolution struct {
	XMLName xml.Name `xml:"resolution" json:"-"`

	// XML namespace declarations
	XMLNS           string `xml:"xmlns,attr" json:"xmlns"`
	XMLNSDC         string `xml:"xmlns dc,attr" json:"xmlnsDC,omitempty"`
	XMLNSHTML       string `xml:"xmlns html,attr" json:"xmlnsHTML,omitempty"`
	XMLNSUSLM       string `xml:"xmlns uslm,attr" json:"xmlnsUSLM,omitempty"`
	XMLNSXSI        string `xml:"xmlns xsi,attr" json:"xmlnsXSI,omitempty"`
	XSISchemaLocation string `xml:"xsi schemaLocation,attr" json:"xsiSchemaLocation,omitempty"`
	XMLLang         string `xml:"xml lang,attr" json:"xmlLang,omitempty"`

	// Document sections
	Meta    *Meta    `xml:"meta" json:"meta"`
	Preface *Preface `xml:"preface" json:"preface,omitempty"`
	Main    *Main    `xml:"main" json:"main,omitempty"`

	// End marker
	EndMarker string `xml:"endMarker,omitempty" json:"endMarker,omitempty"`
}

// Ensure Resolution implements all relevant interfaces
var (
	_ LegislativeDocument = (*Resolution)(nil)
	_ SponsoredDocument   = (*Resolution)(nil)
	_ ActionDocument      = (*Resolution)(nil)
	_ CommitteeDocument   = (*Resolution)(nil)
	_ HierarchicalDocument = (*Resolution)(nil)
	_ MetadataDocument    = (*Resolution)(nil)
)

// GetDocumentNumber returns the resolution number.
func (r *Resolution) GetDocumentNumber() string {
	if r.Meta != nil {
		return r.Meta.DocNumber
	}
	return ""
}

// GetDocumentType returns the document type.
func (r *Resolution) GetDocumentType() string {
	if r.Meta != nil {
		return r.Meta.DCType
	}
	return ""
}

// GetCongress returns the congress number.
func (r *Resolution) GetCongress() string {
	if r.Meta != nil {
		return r.Meta.Congress
	}
	return ""
}

// GetSession returns the session number.
func (r *Resolution) GetSession() string {
	if r.Meta != nil {
		return r.Meta.Session
	}
	return ""
}

// GetTitle returns the document title.
func (r *Resolution) GetTitle() string {
	if r.Meta != nil {
		return r.Meta.DCTitle
	}
	return ""
}

// GetStage returns the document stage.
func (r *Resolution) GetStage() string {
	if r.Meta != nil {
		return r.Meta.DocStage
	}
	return ""
}

// GetChamber returns the current chamber.
func (r *Resolution) GetChamber() string {
	if r.Meta != nil {
		return r.Meta.CurrentChamber
	}
	return ""
}

// IsPublic returns true if this is a public resolution.
func (r *Resolution) IsPublic() bool {
	if r.Meta != nil {
		return r.Meta.PublicPrivate == "public"
	}
	return false
}

// GetCitations returns all citable forms.
func (r *Resolution) GetCitations() []string {
	if r.Meta != nil {
		return r.Meta.CitableAs
	}
	return nil
}

// GetSponsors returns all primary sponsors.
func (r *Resolution) GetSponsors() []Sponsor {
	var sponsors []Sponsor
	if r.Preface != nil {
		for _, action := range r.Preface.Actions {
			if action.ActionDescription != nil {
				sponsors = append(sponsors, action.ActionDescription.Sponsors...)
			}
		}
	}
	return sponsors
}

// GetCosponsors returns all cosponsors.
func (r *Resolution) GetCosponsors() []Cosponsor {
	var cosponsors []Cosponsor
	if r.Preface != nil {
		for _, action := range r.Preface.Actions {
			if action.ActionDescription != nil {
				cosponsors = append(cosponsors, action.ActionDescription.Cosponsors...)
			}
		}
	}
	return cosponsors
}

// GetActions returns all legislative actions.
func (r *Resolution) GetActions() []Action {
	if r.Preface != nil {
		return r.Preface.Actions
	}
	return nil
}

// GetCommittees returns all referenced committees.
func (r *Resolution) GetCommittees() []Committee {
	var committees []Committee
	if r.Preface != nil {
		for _, action := range r.Preface.Actions {
			if action.ActionDescription != nil {
				committees = append(committees, action.ActionDescription.Committees...)
			}
		}
	}
	return committees
}

// GetSections returns all top-level sections.
func (r *Resolution) GetSections() []Section {
	if r.Main != nil {
		return r.Main.Sections
	}
	return nil
}

// GetCreator returns the document creator.
func (r *Resolution) GetCreator() string {
	if r.Meta != nil {
		return r.Meta.DCCreator
	}
	return ""
}

// GetPublisher returns the publisher.
func (r *Resolution) GetPublisher() string {
	if r.Meta != nil {
		return r.Meta.DCPublisher
	}
	return ""
}

// GetLanguage returns the language code.
func (r *Resolution) GetLanguage() string {
	if r.Meta != nil {
		return r.Meta.DCLanguage
	}
	return ""
}

// GetRights returns the rights statement.
func (r *Resolution) GetRights() string {
	if r.Meta != nil {
		return r.Meta.DCRights
	}
	return ""
}

// GetProcessedBy returns the processing tool.
func (r *Resolution) GetProcessedBy() string {
	if r.Meta != nil {
		return r.Meta.ProcessedBy
	}
	return ""
}

// GetProcessedDate returns the processing date.
func (r *Resolution) GetProcessedDate() string {
	if r.Meta != nil {
		return r.Meta.ProcessedDate
	}
	return ""
}

// EngrossedAmendment represents an engrossed amendment document.
type EngrossedAmendment struct {
	XMLName xml.Name `xml:"engrossedAmendment" json:"-"`

	// XML namespace declarations
	XMLNS           string `xml:"xmlns,attr" json:"xmlns"`
	XMLNSDC         string `xml:"xmlns dc,attr" json:"xmlnsDC,omitempty"`
	XMLNSHTML       string `xml:"xmlns html,attr" json:"xmlnsHTML,omitempty"`
	XMLNSUSLM       string `xml:"xmlns uslm,attr" json:"xmlnsUSLM,omitempty"`
	XMLNSXSI        string `xml:"xmlns xsi,attr" json:"xmlnsXSI,omitempty"`
	StyleType       string `xml:"styleType,attr,omitempty" json:"styleType,omitempty"`
	XSISchemaLocation string `xml:"xsi schemaLocation,attr" json:"xsiSchemaLocation,omitempty"`
	XMLLang         string `xml:"xml lang,attr" json:"xmlLang,omitempty"`

	// Document sections
	AmendMeta    *AmendMeta    `xml:"amendMeta" json:"amendMeta"`
	AmendPreface *AmendPreface `xml:"amendPreface" json:"amendPreface,omitempty"`
	AmendMain    *AmendMain    `xml:"amendMain" json:"amendMain,omitempty"`

	// Signatures (can appear after amendMain)
	Signatures *Signatures `xml:"signatures" json:"signatures,omitempty"`

	// Endorsement (can appear after signatures)
	Endorsement *Endorsement `xml:"endorsement" json:"endorsement,omitempty"`
}

// Ensure EngrossedAmendment implements all relevant interfaces
var (
	_ LegislativeDocument = (*EngrossedAmendment)(nil)
	_ AmendmentDocument   = (*EngrossedAmendment)(nil)
	_ ActionDocument      = (*EngrossedAmendment)(nil)
	_ MetadataDocument    = (*EngrossedAmendment)(nil)
)

// GetDocumentNumber returns the amendment document number.
func (e *EngrossedAmendment) GetDocumentNumber() string {
	if e.AmendMeta != nil {
		return e.AmendMeta.DocNumber
	}
	return ""
}

// GetDocumentType returns the document type.
func (e *EngrossedAmendment) GetDocumentType() string {
	if e.AmendMeta != nil {
		return e.AmendMeta.DCType
	}
	return ""
}

// GetCongress returns the congress number.
func (e *EngrossedAmendment) GetCongress() string {
	if e.AmendMeta != nil {
		return e.AmendMeta.Congress
	}
	return ""
}

// GetSession returns the session number.
func (e *EngrossedAmendment) GetSession() string {
	if e.AmendMeta != nil {
		return e.AmendMeta.Session
	}
	return ""
}

// GetTitle returns the document title.
func (e *EngrossedAmendment) GetTitle() string {
	if e.AmendMeta != nil {
		return e.AmendMeta.DCTitle
	}
	return ""
}

// GetStage returns the document stage.
func (e *EngrossedAmendment) GetStage() string {
	if e.AmendMeta != nil {
		return e.AmendMeta.DocStage
	}
	return ""
}

// GetChamber returns the current chamber.
func (e *EngrossedAmendment) GetChamber() string {
	if e.AmendMeta != nil {
		return e.AmendMeta.CurrentChamber
	}
	return ""
}

// IsPublic returns true if this is a public amendment.
func (e *EngrossedAmendment) IsPublic() bool {
	if e.AmendMeta != nil {
		return e.AmendMeta.PublicPrivate == "public"
	}
	return false
}

// GetCitations returns all citable forms.
func (e *EngrossedAmendment) GetCitations() []string {
	if e.AmendMeta != nil {
		return e.AmendMeta.CitableAs
	}
	return nil
}

// GetAmendmentDegree returns the degree of amendment.
func (e *EngrossedAmendment) GetAmendmentDegree() string {
	if e.AmendMeta != nil {
		return e.AmendMeta.AmendDegree
	}
	return ""
}

// GetActions returns all legislative actions.
func (e *EngrossedAmendment) GetActions() []Action {
	if e.AmendPreface != nil {
		return e.AmendPreface.Actions
	}
	return nil
}

// GetCreator returns the document creator.
func (e *EngrossedAmendment) GetCreator() string {
	if e.AmendMeta != nil {
		return e.AmendMeta.DCCreator
	}
	return ""
}

// GetPublisher returns the publisher.
func (e *EngrossedAmendment) GetPublisher() string {
	if e.AmendMeta != nil {
		return e.AmendMeta.DCPublisher
	}
	return ""
}

// GetLanguage returns the language code.
func (e *EngrossedAmendment) GetLanguage() string {
	if e.AmendMeta != nil {
		return e.AmendMeta.DCLanguage
	}
	return ""
}

// GetRights returns the rights statement.
func (e *EngrossedAmendment) GetRights() string {
	if e.AmendMeta != nil {
		return e.AmendMeta.DCRights
	}
	return ""
}

// GetProcessedBy returns the processing tool.
func (e *EngrossedAmendment) GetProcessedBy() string {
	if e.AmendMeta != nil {
		return e.AmendMeta.ProcessedBy
	}
	return ""
}

// GetProcessedDate returns the processing date.
func (e *EngrossedAmendment) GetProcessedDate() string {
	if e.AmendMeta != nil {
		return e.AmendMeta.ProcessedDate
	}
	return ""
}

// Amendment represents a generic amendment document.
type Amendment struct {
	XMLName xml.Name `xml:"amendment" json:"-"`

	// XML namespace declarations
	XMLNS           string `xml:"xmlns,attr" json:"xmlns"`
	XMLNSDC         string `xml:"xmlns dc,attr" json:"xmlnsDC,omitempty"`
	XMLNSHTML       string `xml:"xmlns html,attr" json:"xmlnsHTML,omitempty"`
	XMLNSUSLM       string `xml:"xmlns uslm,attr" json:"xmlnsUSLM,omitempty"`
	XMLNSXSI        string `xml:"xmlns xsi,attr" json:"xmlnsXSI,omitempty"`
	XSISchemaLocation string `xml:"xsi schemaLocation,attr" json:"xsiSchemaLocation,omitempty"`
	XMLLang         string `xml:"xml lang,attr" json:"xmlLang,omitempty"`

	// Document sections
	AmendMeta    *AmendMeta    `xml:"amendMeta" json:"amendMeta"`
	AmendPreface *AmendPreface `xml:"amendPreface" json:"amendPreface,omitempty"`
	AmendMain    *AmendMain    `xml:"amendMain" json:"amendMain,omitempty"`
}

// Ensure Amendment implements all relevant interfaces
var (
	_ LegislativeDocument = (*Amendment)(nil)
	_ AmendmentDocument   = (*Amendment)(nil)
	_ ActionDocument      = (*Amendment)(nil)
	_ MetadataDocument    = (*Amendment)(nil)
)

// GetDocumentNumber returns the amendment document number.
func (a *Amendment) GetDocumentNumber() string {
	if a.AmendMeta != nil {
		return a.AmendMeta.DocNumber
	}
	return ""
}

// GetDocumentType returns the document type.
func (a *Amendment) GetDocumentType() string {
	if a.AmendMeta != nil {
		return a.AmendMeta.DCType
	}
	return ""
}

// GetCongress returns the congress number.
func (a *Amendment) GetCongress() string {
	if a.AmendMeta != nil {
		return a.AmendMeta.Congress
	}
	return ""
}

// GetSession returns the session number.
func (a *Amendment) GetSession() string {
	if a.AmendMeta != nil {
		return a.AmendMeta.Session
	}
	return ""
}

// GetTitle returns the document title.
func (a *Amendment) GetTitle() string {
	if a.AmendMeta != nil {
		return a.AmendMeta.DCTitle
	}
	return ""
}

// GetStage returns the document stage.
func (a *Amendment) GetStage() string {
	if a.AmendMeta != nil {
		return a.AmendMeta.DocStage
	}
	return ""
}

// GetChamber returns the current chamber.
func (a *Amendment) GetChamber() string {
	if a.AmendMeta != nil {
		return a.AmendMeta.CurrentChamber
	}
	return ""
}

// IsPublic returns true if this is a public amendment.
func (a *Amendment) IsPublic() bool {
	if a.AmendMeta != nil {
		return a.AmendMeta.PublicPrivate == "public"
	}
	return false
}

// GetCitations returns all citable forms.
func (a *Amendment) GetCitations() []string {
	if a.AmendMeta != nil {
		return a.AmendMeta.CitableAs
	}
	return nil
}

// GetAmendmentDegree returns the degree of amendment.
func (a *Amendment) GetAmendmentDegree() string {
	if a.AmendMeta != nil {
		return a.AmendMeta.AmendDegree
	}
	return ""
}

// GetActions returns all legislative actions.
func (a *Amendment) GetActions() []Action {
	if a.AmendPreface != nil {
		return a.AmendPreface.Actions
	}
	return nil
}

// GetCreator returns the document creator.
func (a *Amendment) GetCreator() string {
	if a.AmendMeta != nil {
		return a.AmendMeta.DCCreator
	}
	return ""
}

// GetPublisher returns the publisher.
func (a *Amendment) GetPublisher() string {
	if a.AmendMeta != nil {
		return a.AmendMeta.DCPublisher
	}
	return ""
}

// GetLanguage returns the language code.
func (a *Amendment) GetLanguage() string {
	if a.AmendMeta != nil {
		return a.AmendMeta.DCLanguage
	}
	return ""
}

// GetRights returns the rights statement.
func (a *Amendment) GetRights() string {
	if a.AmendMeta != nil {
		return a.AmendMeta.DCRights
	}
	return ""
}

// GetProcessedBy returns the processing tool.
func (a *Amendment) GetProcessedBy() string {
	if a.AmendMeta != nil {
		return a.AmendMeta.ProcessedBy
	}
	return ""
}

// GetProcessedDate returns the processing date.
func (a *Amendment) GetProcessedDate() string {
	if a.AmendMeta != nil {
		return a.AmendMeta.ProcessedDate
	}
	return ""
}
