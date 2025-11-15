package uslm

import "encoding/xml"

// Preface represents the preface section for bills and resolutions.
// This contains metadata that IS rendered/published with the document.
type Preface struct {
	XMLName xml.Name `xml:"preface" json:"-"`

	SlugLine         string            `xml:"slugLine,omitempty" json:"slugLine,omitempty"`
	DistributionCode *DistributionCode `xml:"distributionCode" json:"distributionCode,omitempty"`
	Congress         *CongressElement  `xml:"congress" json:"congress,omitempty"`
	Session          *SessionElement   `xml:"session" json:"session,omitempty"`
	DCType           string            `xml:"http://purl.org/dc/elements/1.1/ type" json:"dcType,omitempty"`
	DocNumber        string            `xml:"docNumber,omitempty" json:"docNumber,omitempty"`
	DCTitle          string            `xml:"http://purl.org/dc/elements/1.1/ title" json:"dcTitle,omitempty"`
	CurrentChamber   *CurrentChamber   `xml:"currentChamber" json:"currentChamber,omitempty"`
	Actions          []Action          `xml:"action" json:"actions,omitempty"`
}

// AmendPreface represents the preface section for amendment documents.
type AmendPreface struct {
	XMLName xml.Name `xml:"amendPreface" json:"-"`

	SlugLine       string          `xml:"slugLine,omitempty" json:"slugLine,omitempty"`
	CurrentChamber *CurrentChamber `xml:"currentChamber" json:"currentChamber,omitempty"`
	Actions        []Action        `xml:"action" json:"actions,omitempty"`
}

// DistributionCode represents a distribution code element with display attribute.
type DistributionCode struct {
	XMLName xml.Name `xml:"distributionCode" json:"-"`
	Display string   `xml:"display,attr,omitempty" json:"display,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}

// CongressElement represents the congress element with value attribute.
type CongressElement struct {
	XMLName xml.Name `xml:"congress" json:"-"`
	Value   string   `xml:"value,attr,omitempty" json:"value,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}

// SessionElement represents the session element with value attribute.
type SessionElement struct {
	XMLName xml.Name `xml:"session" json:"-"`
	Value   string   `xml:"value,attr,omitempty" json:"value,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}

// CurrentChamber represents which chamber currently has the document.
type CurrentChamber struct {
	XMLName xml.Name `xml:"currentChamber" json:"-"`
	Value   string   `xml:"value,attr,omitempty" json:"value,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}

// Action represents a legislative action taken on the document.
type Action struct {
	XMLName           xml.Name           `xml:"action" json:"-"`
	ActionStage       string             `xml:"actionStage,attr,omitempty" json:"actionStage,omitempty"`
	Date              *ActionDate        `xml:"date" json:"date,omitempty"`
	ActionDescription *ActionDescription `xml:"actionDescription" json:"actionDescription,omitempty"`
	ActionInstruction string             `xml:"actionInstruction,omitempty" json:"actionInstruction,omitempty"`
}

// ActionDate represents the date of an action.
type ActionDate struct {
	XMLName xml.Name `xml:"date" json:"-"`
	Date    string   `xml:"date,attr,omitempty" json:"date,omitempty"` // ISO format YYYY-MM-DD
	Text    string   `xml:",chardata" json:"text,omitempty"`
	Inline  []Inline `xml:"inline" json:"inline,omitempty"`
}

// ActionDescription describes what happened in an action.
type ActionDescription struct {
	XMLName    xml.Name    `xml:"actionDescription" json:"-"`
	Text       string      `xml:",chardata" json:"text,omitempty"`
	Sponsors   []Sponsor   `xml:"sponsor" json:"sponsors,omitempty"`
	Cosponsors []Cosponsor `xml:"cosponsor" json:"cosponsors,omitempty"`
	Committees []Committee `xml:"committee" json:"committees,omitempty"`
	Inline     []Inline    `xml:"inline" json:"inline,omitempty"`
}

// Sponsor represents the primary sponsor of legislation.
type Sponsor struct {
	XMLName  xml.Name `xml:"sponsor" json:"-"`
	SenateID string   `xml:"senateId,attr,omitempty" json:"senateId,omitempty"`
	HouseID  string   `xml:"houseId,attr,omitempty" json:"houseId,omitempty"`
	Text     string   `xml:",chardata" json:"text,omitempty"`
	Inline   []Inline `xml:"inline" json:"inline,omitempty"`
}

// GetID returns the sponsor's official ID (Senate or House).
func (s *Sponsor) GetID() string {
	if s.SenateID != "" {
		return s.SenateID
	}
	return s.HouseID
}

// GetName returns the sponsor's name text.
func (s *Sponsor) GetName() string {
	return s.Text
}

// Cosponsor represents a cosponsor of legislation.
type Cosponsor struct {
	XMLName  xml.Name `xml:"cosponsor" json:"-"`
	SenateID string   `xml:"senateId,attr,omitempty" json:"senateId,omitempty"`
	HouseID  string   `xml:"houseId,attr,omitempty" json:"houseId,omitempty"`
	Text     string   `xml:",chardata" json:"text,omitempty"`
	Inline   []Inline `xml:"inline" json:"inline,omitempty"`
}

// GetID returns the cosponsor's official ID (Senate or House).
func (c *Cosponsor) GetID() string {
	if c.SenateID != "" {
		return c.SenateID
	}
	return c.HouseID
}

// GetName returns the cosponsor's name text.
func (c *Cosponsor) GetName() string {
	return c.Text
}

// Committee represents a congressional committee.
type Committee struct {
	XMLName     xml.Name `xml:"committee" json:"-"`
	CommitteeID string   `xml:"committeeId,attr,omitempty" json:"committeeId,omitempty"`
	Text        string   `xml:",chardata" json:"text,omitempty"`
}

// GetID returns the committee's official ID.
func (c *Committee) GetID() string {
	return c.CommitteeID
}

// GetName returns the committee's name.
func (c *Committee) GetName() string {
	return c.Text
}
