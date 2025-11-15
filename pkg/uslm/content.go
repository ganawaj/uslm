package uslm

import "encoding/xml"

// Main represents the main content section of a bill or resolution.
type Main struct {
	XMLName   xml.Name   `xml:"main" json:"-"`
	StyleType string     `xml:"styleType,attr,omitempty" json:"styleType,omitempty"`
	LongTitle *LongTitle `xml:"longTitle" json:"longTitle,omitempty"`
	EnactingFormula *EnactingFormula `xml:"enactingFormula" json:"enactingFormula,omitempty"`
	TOC       *TOC       `xml:"toc" json:"toc,omitempty"`
	Preamble  *Preamble  `xml:"preamble" json:"preamble,omitempty"`
	Sections  []Section  `xml:"section" json:"sections,omitempty"`
	Titles    []Title    `xml:"title" json:"titles,omitempty"`
	EndMarker string     `xml:"endMarker,omitempty" json:"endMarker,omitempty"`
}

// AmendMain represents the main content section of an amendment document.
type AmendMain struct {
	XMLName                       xml.Name               `xml:"amendMain" json:"-"`
	AmendmentInstructionLineNumbering string             `xml:"amendmentInstructionLineNumbering,attr,omitempty" json:"amendmentInstructionLineNumbering,omitempty"`
	ResolvingClause               *ResolvingClause       `xml:"resolvingClause" json:"resolvingClause,omitempty"`
	Sections                      []Section              `xml:"section" json:"sections,omitempty"`
	DocTitle                      string                 `xml:"docTitle,omitempty" json:"docTitle,omitempty"`
	AmendmentInstructions         []AmendmentInstruction `xml:"amendmentInstruction" json:"amendmentInstructions,omitempty"`
	Signatures                    *Signatures            `xml:"signatures" json:"signatures,omitempty"`
	Endorsement                   *Endorsement           `xml:"endorsement" json:"endorsement,omitempty"`
}

// LongTitle represents the long title section containing doc title and official title.
type LongTitle struct {
	XMLName       xml.Name `xml:"longTitle" json:"-"`
	DocTitle      string   `xml:"docTitle" json:"docTitle,omitempty"`
	OfficialTitle string   `xml:"officialTitle" json:"officialTitle,omitempty"`
}

// EnactingFormula represents the enacting formula (e.g., "Be it enacted...").
type EnactingFormula struct {
	XMLName xml.Name `xml:"enactingFormula" json:"-"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
	I       []Italic `xml:"i" json:"i,omitempty"`
}

// TOC represents the table of contents.
type TOC struct {
	XMLName       xml.Name        `xml:"toc" json:"-"`
	ReferenceItem []ReferenceItem `xml:"referenceItem" json:"referenceItems,omitempty"`
}

// ReferenceItem represents an item in the table of contents.
type ReferenceItem struct {
	XMLName    xml.Name `xml:"referenceItem" json:"-"`
	Role       string   `xml:"role,attr,omitempty" json:"role,omitempty"`
	Designator string   `xml:"designator,omitempty" json:"designator,omitempty"`
	Label      string   `xml:"label,omitempty" json:"label,omitempty"`
}

// Preamble represents the preamble section (for resolutions with recitals).
type Preamble struct {
	XMLName         xml.Name         `xml:"preamble" json:"-"`
	Recitals        []Recital        `xml:"recital" json:"recitals,omitempty"`
	ResolvingClause *ResolvingClause `xml:"resolvingClause" json:"resolvingClause,omitempty"`
}

// Recital represents a "whereas" clause in a resolution preamble.
type Recital struct {
	XMLName    xml.Name    `xml:"recital" json:"-"`
	Text       string      `xml:",chardata" json:"text,omitempty"`
	P          []P         `xml:"p" json:"p,omitempty"`
	Paragraphs []Paragraph `xml:"paragraph" json:"paragraphs,omitempty"`
}

// ResolvingClause represents the resolving clause (e.g., "Resolved, ").
type ResolvingClause struct {
	XMLName xml.Name `xml:"resolvingClause" json:"-"`
	Class   string   `xml:"class,attr,omitempty" json:"class,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
	I       []Italic `xml:"i" json:"i,omitempty"`
}

// Section represents a section of legislative content.
type Section struct {
	XMLName       xml.Name       `xml:"section" json:"-"`
	ID            string         `xml:"id,attr,omitempty" json:"id,omitempty"`
	Identifier    string         `xml:"identifier,attr,omitempty" json:"identifier,omitempty"`
	Role          string         `xml:"role,attr,omitempty" json:"role,omitempty"`
	Class         string         `xml:"class,attr,omitempty" json:"class,omitempty"`
	Num           *Num           `xml:"num" json:"num,omitempty"`
	Heading       *Heading       `xml:"heading" json:"heading,omitempty"`
	Chapeau       *Chapeau       `xml:"chapeau" json:"chapeau,omitempty"`
	Content       *Content       `xml:"content" json:"content,omitempty"`
	Paragraphs    []Paragraph    `xml:"paragraph" json:"paragraphs,omitempty"`
	Subsections   []Subsection   `xml:"subsection" json:"subsections,omitempty"`
}

// GetID returns the section's unique ID.
func (s *Section) GetID() string {
	return s.ID
}

// GetIdentifier returns the section's logical identifier.
func (s *Section) GetIdentifier() string {
	return s.Identifier
}

// GetNum returns the section's number text.
func (s *Section) GetNum() string {
	if s.Num != nil {
		return s.Num.Text
	}
	return ""
}

// GetNumValue returns the section's normalized number value.
func (s *Section) GetNumValue() string {
	if s.Num != nil {
		return s.Num.Value
	}
	return ""
}

// GetHeading returns the section's heading text.
func (s *Section) GetHeading() string {
	if s.Heading != nil {
		return s.Heading.GetText()
	}
	return ""
}

// GetContent returns the section's content text.
func (s *Section) GetContent() string {
	if s.Content != nil {
		return s.Content.Text
	}
	return ""
}

// GetChapeau returns the section's chapeau text.
func (s *Section) GetChapeau() string {
	if s.Chapeau != nil {
		return s.Chapeau.Text
	}
	return ""
}

// Title represents a title division (in large bills).
type Title struct {
	XMLName  xml.Name  `xml:"title" json:"-"`
	ID       string    `xml:"id,attr,omitempty" json:"id,omitempty"`
	Num      *Num      `xml:"num" json:"num,omitempty"`
	Heading  *Heading  `xml:"heading" json:"heading,omitempty"`
	Sections []Section `xml:"section" json:"sections,omitempty"`
}

// Subsection represents a subsection (e.g., (a), (b), (c)).
type Subsection struct {
	XMLName    xml.Name    `xml:"subsection" json:"-"`
	ID         string      `xml:"id,attr,omitempty" json:"id,omitempty"`
	Identifier string      `xml:"identifier,attr,omitempty" json:"identifier,omitempty"`
	Class      string      `xml:"class,attr,omitempty" json:"class,omitempty"`
	Num        *Num        `xml:"num" json:"num,omitempty"`
	Heading    *Heading    `xml:"heading" json:"heading,omitempty"`
	Chapeau    *Chapeau    `xml:"chapeau" json:"chapeau,omitempty"`
	Content    *Content    `xml:"content" json:"content,omitempty"`
	Paragraphs []Paragraph `xml:"paragraph" json:"paragraphs,omitempty"`
}

// Paragraph represents a paragraph (e.g., (1), (2), (3)).
type Paragraph struct {
	XMLName       xml.Name       `xml:"paragraph" json:"-"`
	ID            string         `xml:"id,attr,omitempty" json:"id,omitempty"`
	Identifier    string         `xml:"identifier,attr,omitempty" json:"identifier,omitempty"`
	Class         string         `xml:"class,attr,omitempty" json:"class,omitempty"`
	Role          string         `xml:"role,attr,omitempty" json:"role,omitempty"`
	Num           *Num           `xml:"num" json:"num,omitempty"`
	Heading       *Heading       `xml:"heading" json:"heading,omitempty"`
	Chapeau       *Chapeau       `xml:"chapeau" json:"chapeau,omitempty"`
	Content       *Content       `xml:"content" json:"content,omitempty"`
	Subparagraphs []Subparagraph `xml:"subparagraph" json:"subparagraphs,omitempty"`
}

// Subparagraph represents a subparagraph (e.g., (A), (B), (C)).
type Subparagraph struct {
	XMLName    xml.Name `xml:"subparagraph" json:"-"`
	ID         string   `xml:"id,attr,omitempty" json:"id,omitempty"`
	Identifier string   `xml:"identifier,attr,omitempty" json:"identifier,omitempty"`
	Class      string   `xml:"class,attr,omitempty" json:"class,omitempty"`
	Num        *Num     `xml:"num" json:"num,omitempty"`
	Chapeau    *Chapeau `xml:"chapeau" json:"chapeau,omitempty"`
	Content    *Content `xml:"content" json:"content,omitempty"`
	Clauses    []Clause `xml:"clause" json:"clauses,omitempty"`
}

// Clause represents a clause (e.g., (i), (ii), (iii)).
type Clause struct {
	XMLName    xml.Name    `xml:"clause" json:"-"`
	ID         string      `xml:"id,attr,omitempty" json:"id,omitempty"`
	Identifier string      `xml:"identifier,attr,omitempty" json:"identifier,omitempty"`
	Class      string      `xml:"class,attr,omitempty" json:"class,omitempty"`
	Num        *Num        `xml:"num" json:"num,omitempty"`
	Content    *Content    `xml:"content" json:"content,omitempty"`
	Subclauses []Subclause `xml:"subclause" json:"subclauses,omitempty"`
}

// Subclause represents a subclause (e.g., (I), (II), (III)).
type Subclause struct {
	XMLName    xml.Name `xml:"subclause" json:"-"`
	ID         string   `xml:"id,attr,omitempty" json:"id,omitempty"`
	Identifier string   `xml:"identifier,attr,omitempty" json:"identifier,omitempty"`
	Class      string   `xml:"class,attr,omitempty" json:"class,omitempty"`
	Num        *Num     `xml:"num" json:"num,omitempty"`
	Content    *Content `xml:"content" json:"content,omitempty"`
}

// AmendmentInstruction represents an instruction for how to amend existing law.
type AmendmentInstruction struct {
	XMLName xml.Name `xml:"amendmentInstruction" json:"-"`
	Num     *Num     `xml:"num" json:"num,omitempty"`
	Content *Content `xml:"content" json:"content,omitempty"`
}

// Signatures represents the signatures block in amendment documents.
type Signatures struct {
	XMLName   xml.Name    `xml:"signatures" json:"-"`
	Signature []Signature `xml:"signature" json:"signatures,omitempty"`
}

// Signature represents an individual signature.
type Signature struct {
	XMLName  xml.Name `xml:"signature" json:"-"`
	Notation *Notation `xml:"notation" json:"notation,omitempty"`
	Role     string   `xml:"role,omitempty" json:"role,omitempty"`
	Text     string   `xml:",chardata" json:"text,omitempty"`
}

// Notation represents a notation within a signature (e.g., "Attest:").
type Notation struct {
	XMLName xml.Name `xml:"notation" json:"-"`
	Type    string   `xml:"type,attr,omitempty" json:"type,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}

// Endorsement represents the endorsement block at the end of amendment documents.
type Endorsement struct {
	XMLName     xml.Name         `xml:"endorsement" json:"-"`
	Orientation string           `xml:"orientation,attr,omitempty" json:"orientation,omitempty"`
	Congress    *CongressElement `xml:"congress" json:"congress,omitempty"`
	Session     *SessionElement  `xml:"session" json:"session,omitempty"`
	DCType      string           `xml:"http://purl.org/dc/elements/1.1/ type" json:"dcType,omitempty"`
	DocNumber   string           `xml:"docNumber,omitempty" json:"docNumber,omitempty"`
	DocTitle    string           `xml:"docTitle,omitempty" json:"docTitle,omitempty"`
}
