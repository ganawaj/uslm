package uslm

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// ParseBill parses XML data into a Bill struct.
func ParseBill(data []byte) (*Bill, error) {
	var bill Bill
	if err := xml.Unmarshal(data, &bill); err != nil {
		return nil, fmt.Errorf("failed to parse bill: %w", err)
	}
	return &bill, nil
}

// ParseResolution parses XML data into a Resolution struct.
func ParseResolution(data []byte) (*Resolution, error) {
	var resolution Resolution
	if err := xml.Unmarshal(data, &resolution); err != nil {
		return nil, fmt.Errorf("failed to parse resolution: %w", err)
	}
	return &resolution, nil
}

// ParseEngrossedAmendment parses XML data into an EngrossedAmendment struct.
func ParseEngrossedAmendment(data []byte) (*EngrossedAmendment, error) {
	var amendment EngrossedAmendment
	if err := xml.Unmarshal(data, &amendment); err != nil {
		return nil, fmt.Errorf("failed to parse engrossed amendment: %w", err)
	}
	return &amendment, nil
}

// ParseAmendment parses XML data into an Amendment struct.
func ParseAmendment(data []byte) (*Amendment, error) {
	var amendment Amendment
	if err := xml.Unmarshal(data, &amendment); err != nil {
		return nil, fmt.Errorf("failed to parse amendment: %w", err)
	}
	return &amendment, nil
}

// DocumentType represents the type of USLM document.
type DocumentType string

const (
	DocumentTypeBill               DocumentType = "bill"
	DocumentTypeResolution         DocumentType = "resolution"
	DocumentTypeAmendment          DocumentType = "amendment"
	DocumentTypeEngrossedAmendment DocumentType = "engrossedAmendment"
	DocumentTypeUnknown            DocumentType = "unknown"
)

// DetectDocumentType examines XML data to determine the document type.
func DetectDocumentType(data []byte) DocumentType {
	// Simple detection based on root element
	content := string(data)

	if strings.Contains(content, "<bill ") || strings.HasPrefix(strings.TrimSpace(content), "<?xml") && strings.Contains(content, "<bill") {
		return DocumentTypeBill
	}
	if strings.Contains(content, "<resolution ") || strings.Contains(content, "<resolution>") {
		return DocumentTypeResolution
	}
	if strings.Contains(content, "<engrossedAmendment ") || strings.Contains(content, "<engrossedAmendment>") {
		return DocumentTypeEngrossedAmendment
	}
	if strings.Contains(content, "<amendment ") || strings.Contains(content, "<amendment>") {
		return DocumentTypeAmendment
	}

	return DocumentTypeUnknown
}

// ParseDocument automatically detects and parses the document type.
// Returns a LegislativeDocument interface that can be type-asserted to the specific type.
func ParseDocument(data []byte) (LegislativeDocument, error) {
	docType := DetectDocumentType(data)

	switch docType {
	case DocumentTypeBill:
		return ParseBill(data)
	case DocumentTypeResolution:
		return ParseResolution(data)
	case DocumentTypeEngrossedAmendment:
		return ParseEngrossedAmendment(data)
	case DocumentTypeAmendment:
		return ParseAmendment(data)
	default:
		return nil, fmt.Errorf("unknown document type")
	}
}

// ParseDocumentFromReader parses a document from an io.Reader.
func ParseDocumentFromReader(r io.Reader) (LegislativeDocument, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read document: %w", err)
	}
	return ParseDocument(data)
}

// MarshalBillToXML marshals a Bill to XML with proper formatting.
func MarshalBillToXML(bill *Bill) ([]byte, error) {
	data, err := xml.MarshalIndent(bill, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal bill to XML: %w", err)
	}
	// Add XML declaration
	return append([]byte(xml.Header), data...), nil
}

// MarshalResolutionToXML marshals a Resolution to XML with proper formatting.
func MarshalResolutionToXML(resolution *Resolution) ([]byte, error) {
	data, err := xml.MarshalIndent(resolution, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resolution to XML: %w", err)
	}
	return append([]byte(xml.Header), data...), nil
}

// MarshalEngrossedAmendmentToXML marshals an EngrossedAmendment to XML.
func MarshalEngrossedAmendmentToXML(amendment *EngrossedAmendment) ([]byte, error) {
	data, err := xml.MarshalIndent(amendment, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal engrossed amendment to XML: %w", err)
	}
	return append([]byte(xml.Header), data...), nil
}

// MarshalAmendmentToXML marshals an Amendment to XML.
func MarshalAmendmentToXML(amendment *Amendment) ([]byte, error) {
	data, err := xml.MarshalIndent(amendment, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal amendment to XML: %w", err)
	}
	return append([]byte(xml.Header), data...), nil
}

// ToJSON converts any USLM document to JSON.
func ToJSON(doc interface{}) ([]byte, error) {
	return json.MarshalIndent(doc, "", "  ")
}

// BillFromJSON parses JSON data into a Bill struct.
func BillFromJSON(data []byte) (*Bill, error) {
	var bill Bill
	if err := json.Unmarshal(data, &bill); err != nil {
		return nil, fmt.Errorf("failed to parse bill from JSON: %w", err)
	}
	return &bill, nil
}

// ResolutionFromJSON parses JSON data into a Resolution struct.
func ResolutionFromJSON(data []byte) (*Resolution, error) {
	var resolution Resolution
	if err := json.Unmarshal(data, &resolution); err != nil {
		return nil, fmt.Errorf("failed to parse resolution from JSON: %w", err)
	}
	return &resolution, nil
}

// EngrossedAmendmentFromJSON parses JSON data into an EngrossedAmendment struct.
func EngrossedAmendmentFromJSON(data []byte) (*EngrossedAmendment, error) {
	var amendment EngrossedAmendment
	if err := json.Unmarshal(data, &amendment); err != nil {
		return nil, fmt.Errorf("failed to parse engrossed amendment from JSON: %w", err)
	}
	return &amendment, nil
}

// AmendmentFromJSON parses JSON data into an Amendment struct.
func AmendmentFromJSON(data []byte) (*Amendment, error) {
	var amendment Amendment
	if err := json.Unmarshal(data, &amendment); err != nil {
		return nil, fmt.Errorf("failed to parse amendment from JSON: %w", err)
	}
	return &amendment, nil
}
