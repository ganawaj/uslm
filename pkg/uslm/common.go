package uslm

import "encoding/xml"

// Common XML namespace constants used throughout USLM documents.
const (
	NamespaceUSLM    = "http://schemas.gpo.gov/xml/uslm"
	NamespaceDC      = "http://purl.org/dc/elements/1.1/"
	NamespaceHTML    = "http://www.w3.org/1999/xhtml"
	NamespaceXSI     = "http://www.w3.org/2001/XMLSchema-instance"
	NamespaceMathML  = "http://www.w3.org/1998/Math/MathML"
	NamespaceSenate  = "https://www.senate.gov/schemas"
)

// MixedContent represents XML content that can contain both text and child elements.
// This is used for elements that have inline markup mixed with text.
type MixedContent struct {
	Text     string    `xml:",chardata" json:"text,omitempty"`
	Inline   []Inline  `xml:"inline" json:"inline,omitempty"`
	I        []Italic  `xml:"i" json:"i,omitempty"`
	B        []Bold    `xml:"b" json:"b,omitempty"`
	Sup      []Sup     `xml:"sup" json:"sup,omitempty"`
	Sub      []Sub     `xml:"sub" json:"sub,omitempty"`
	Term     []Term    `xml:"term" json:"term,omitempty"`
	Ref      []Ref     `xml:"ref" json:"ref,omitempty"`
	P        []P       `xml:"p" json:"p,omitempty"`
}

// Inline represents a generic inline element with optional class and content.
type Inline struct {
	XMLName xml.Name `xml:"inline" json:"-"`
	Class   string   `xml:"class,attr,omitempty" json:"class,omitempty"`
	Role    string   `xml:"role,attr,omitempty" json:"role,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}

// Italic represents italic text (<i> element).
type Italic struct {
	XMLName xml.Name `xml:"i" json:"-"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}

// Bold represents bold text (<b> element).
type Bold struct {
	XMLName xml.Name `xml:"b" json:"-"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}

// Sup represents superscript text.
type Sup struct {
	XMLName xml.Name `xml:"sup" json:"-"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}

// Sub represents subscript text.
type Sub struct {
	XMLName xml.Name `xml:"sub" json:"-"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}

// Term represents a defined term.
type Term struct {
	XMLName xml.Name `xml:"term" json:"-"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}

// Ref represents a reference/hyperlink to other content.
type Ref struct {
	XMLName xml.Name `xml:"ref" json:"-"`
	Href    string   `xml:"href,attr,omitempty" json:"href,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
	InnerRef *Ref    `xml:"ref" json:"innerRef,omitempty"` // Nested refs can occur
}

// P represents a paragraph element within mixed content.
type P struct {
	XMLName xml.Name `xml:"p" json:"-"`
	Class   string   `xml:"class,attr,omitempty" json:"class,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}

// ShortTitle represents a short title citation within content.
type ShortTitle struct {
	XMLName xml.Name `xml:"shortTitle" json:"-"`
	Role    string   `xml:"role,attr,omitempty" json:"role,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}

// QuotedText represents quoted text in content.
type QuotedText struct {
	XMLName xml.Name `xml:"quotedText" json:"-"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}

// AmendingAction represents an amendment action type (delete, insert, amend, etc.).
type AmendingAction struct {
	XMLName xml.Name `xml:"amendingAction" json:"-"`
	Type    string   `xml:"type,attr,omitempty" json:"type,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}

// Num represents a designation number (e.g., "SECTION 1.", "(a)", "(1)").
type Num struct {
	XMLName xml.Name `xml:"num" json:"-"`
	Value   string   `xml:"value,attr,omitempty" json:"value,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
}

// Heading represents a heading for a section or other structural element.
type Heading struct {
	XMLName xml.Name `xml:"heading" json:"-"`
	Class   string   `xml:"class,attr,omitempty" json:"class,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
	Inline  []Inline `xml:"inline" json:"inline,omitempty"`
}

// GetText returns the text content of the heading.
func (h *Heading) GetText() string {
	return h.Text
}

// Content represents the main content of a legislative element.
type Content struct {
	XMLName        xml.Name          `xml:"content" json:"-"`
	Class          string            `xml:"class,attr,omitempty" json:"class,omitempty"`
	Text           string            `xml:",chardata" json:"text,omitempty"`
	Inline         []Inline          `xml:"inline" json:"inline,omitempty"`
	I              []Italic          `xml:"i" json:"i,omitempty"`
	Ref            []Ref             `xml:"ref" json:"ref,omitempty"`
	ShortTitle     []ShortTitle      `xml:"shortTitle" json:"shortTitle,omitempty"`
	QuotedText     []QuotedText      `xml:"quotedText" json:"quotedText,omitempty"`
	AmendingAction []AmendingAction  `xml:"amendingAction" json:"amendingAction,omitempty"`
	QuotedContent  []QuotedContent   `xml:"quotedContent" json:"quotedContent,omitempty"`
	AmendmentContent []AmendmentContent `xml:"amendmentContent" json:"amendmentContent,omitempty"`
}

// Chapeau represents introductory text (lead-in) before nested elements.
type Chapeau struct {
	XMLName        xml.Name         `xml:"chapeau" json:"-"`
	Class          string           `xml:"class,attr,omitempty" json:"class,omitempty"`
	Text           string           `xml:",chardata" json:"text,omitempty"`
	Inline         []Inline         `xml:"inline" json:"inline,omitempty"`
	Ref            []Ref            `xml:"ref" json:"ref,omitempty"`
	AmendingAction []AmendingAction `xml:"amendingAction" json:"amendingAction,omitempty"`
}

// QuotedContent represents quoted legislative content (for amending existing law).
type QuotedContent struct {
	XMLName    xml.Name    `xml:"quotedContent" json:"-"`
	ID         string      `xml:"id,attr,omitempty" json:"id,omitempty"`
	StyleType  string      `xml:"styleType,attr,omitempty" json:"styleType,omitempty"`
	Paragraph  []Paragraph `xml:"paragraph" json:"paragraph,omitempty"`
	Subsection []Subsection `xml:"subsection" json:"subsection,omitempty"`
	Section    []Section   `xml:"section" json:"section,omitempty"`
}

// AmendmentContent represents content being added or modified by an amendment.
type AmendmentContent struct {
	XMLName   xml.Name  `xml:"amendmentContent" json:"-"`
	Class     string    `xml:"class,attr,omitempty" json:"class,omitempty"`
	Changed   string    `xml:"changed,attr,omitempty" json:"changed,omitempty"`
	StyleType string    `xml:"styleType,attr,omitempty" json:"styleType,omitempty"`
	Section   []Section `xml:"section" json:"section,omitempty"`
}
