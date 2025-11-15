package uslm

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseBill(t *testing.T) {
	// Read sample bill
	data, err := os.ReadFile(filepath.Join("..", "..", "bill-version-samples-september-2024", "BILLS-114s32cds.xml"))
	if err != nil {
		t.Fatalf("failed to read sample bill: %v", err)
	}

	bill, err := ParseBill(data)
	if err != nil {
		t.Fatalf("failed to parse bill: %v", err)
	}

	// Verify basic metadata
	if bill.GetDocumentNumber() != "32" {
		t.Errorf("expected doc number '32', got '%s'", bill.GetDocumentNumber())
	}
	if bill.GetCongress() != "114" {
		t.Errorf("expected congress '114', got '%s'", bill.GetCongress())
	}
	if bill.GetSession() != "1" {
		t.Errorf("expected session '1', got '%s'", bill.GetSession())
	}
	if bill.GetDocumentType() != "Senate Bill" {
		t.Errorf("expected type 'Senate Bill', got '%s'", bill.GetDocumentType())
	}
	if bill.GetStage() != "Committee Discharged Senate" {
		t.Errorf("expected stage 'Committee Discharged Senate', got '%s'", bill.GetStage())
	}
	if bill.GetChamber() != "SENATE" {
		t.Errorf("expected chamber 'SENATE', got '%s'", bill.GetChamber())
	}
	if !bill.IsPublic() {
		t.Error("expected bill to be public")
	}

	// Verify citations
	citations := bill.GetCitations()
	if len(citations) != 3 {
		t.Errorf("expected 3 citations, got %d", len(citations))
	}

	// Verify sponsors
	sponsors := bill.GetSponsors()
	if len(sponsors) != 1 {
		t.Errorf("expected 1 sponsor, got %d", len(sponsors))
	} else if sponsors[0].GetID() != "S221" {
		t.Errorf("expected sponsor ID 'S221', got '%s'", sponsors[0].GetID())
	}

	// Verify cosponsors
	cosponsors := bill.GetCosponsors()
	if len(cosponsors) != 5 {
		t.Errorf("expected 5 cosponsors, got %d", len(cosponsors))
	}

	// Verify committees
	committees := bill.GetCommittees()
	if len(committees) != 2 {
		t.Errorf("expected 2 committees, got %d", len(committees))
	}

	// Verify sections
	sections := bill.GetSections()
	if len(sections) != 3 {
		t.Errorf("expected 3 sections, got %d", len(sections))
	}
	if len(sections) > 0 && sections[0].GetNumValue() != "1" {
		t.Errorf("expected section 1 value '1', got '%s'", sections[0].GetNumValue())
	}

	// Verify actions
	actions := bill.GetActions()
	if len(actions) != 2 {
		t.Errorf("expected 2 actions, got %d", len(actions))
	}

	// Test round-trip: Bill -> JSON -> Bill
	jsonData, err := ToJSON(bill)
	if err != nil {
		t.Fatalf("failed to marshal bill to JSON: %v", err)
	}

	bill2, err := BillFromJSON(jsonData)
	if err != nil {
		t.Fatalf("failed to parse bill from JSON: %v", err)
	}

	// Verify data preserved
	if bill2.GetDocumentNumber() != bill.GetDocumentNumber() {
		t.Errorf("doc number not preserved: got '%s', want '%s'", bill2.GetDocumentNumber(), bill.GetDocumentNumber())
	}
	if bill2.GetCongress() != bill.GetCongress() {
		t.Errorf("congress not preserved: got '%s', want '%s'", bill2.GetCongress(), bill.GetCongress())
	}
}

func TestParseResolution(t *testing.T) {
	// Read sample resolution
	data, err := os.ReadFile(filepath.Join("..", "..", "bill-version-samples-september-2024", "BILLS-116sres100ats.xml"))
	if err != nil {
		t.Fatalf("failed to read sample resolution: %v", err)
	}

	resolution, err := ParseResolution(data)
	if err != nil {
		t.Fatalf("failed to parse resolution: %v", err)
	}

	// Verify basic metadata
	if resolution.GetDocumentNumber() != "100" {
		t.Errorf("expected doc number '100', got '%s'", resolution.GetDocumentNumber())
	}
	if resolution.GetCongress() != "116" {
		t.Errorf("expected congress '116', got '%s'", resolution.GetCongress())
	}
	if resolution.GetDocumentType() != "Senate Simple Resolution" {
		t.Errorf("expected type 'Senate Simple Resolution', got '%s'", resolution.GetDocumentType())
	}

	// Verify sponsors
	sponsors := resolution.GetSponsors()
	if len(sponsors) != 1 {
		t.Errorf("expected 1 sponsor, got %d", len(sponsors))
	}

	// Verify cosponsors (this resolution has many)
	cosponsors := resolution.GetCosponsors()
	if len(cosponsors) < 30 {
		t.Errorf("expected at least 30 cosponsors, got %d", len(cosponsors))
	}

	// Test round-trip: Resolution -> JSON -> Resolution
	jsonData, err := ToJSON(resolution)
	if err != nil {
		t.Fatalf("failed to marshal resolution to JSON: %v", err)
	}

	resolution2, err := ResolutionFromJSON(jsonData)
	if err != nil {
		t.Fatalf("failed to parse resolution from JSON: %v", err)
	}

	if resolution2.GetDocumentNumber() != resolution.GetDocumentNumber() {
		t.Errorf("doc number not preserved: got '%s', want '%s'", resolution2.GetDocumentNumber(), resolution.GetDocumentNumber())
	}
}

func TestParseEngrossedAmendment(t *testing.T) {
	// Read sample engrossed amendment
	data, err := os.ReadFile(filepath.Join("..", "..", "bill-version-samples-september-2024", "BILLS-116hr1865eas.xml"))
	if err != nil {
		t.Fatalf("failed to read sample engrossed amendment: %v", err)
	}

	amendment, err := ParseEngrossedAmendment(data)
	if err != nil {
		t.Fatalf("failed to parse engrossed amendment: %v", err)
	}

	// Verify basic metadata
	if amendment.GetDocumentNumber() != "1865" {
		t.Errorf("expected doc number '1865', got '%s'", amendment.GetDocumentNumber())
	}
	if amendment.GetCongress() != "116" {
		t.Errorf("expected congress '116', got '%s'", amendment.GetCongress())
	}
	if amendment.GetAmendmentDegree() != "first" {
		t.Errorf("expected amendment degree 'first', got '%s'", amendment.GetAmendmentDegree())
	}
	if amendment.GetDocumentType() != "Engrossed Amendment Senate" {
		t.Errorf("expected type 'Engrossed Amendment Senate', got '%s'", amendment.GetDocumentType())
	}
	if amendment.GetChamber() != "SENATE" {
		t.Errorf("expected chamber 'SENATE', got '%s'", amendment.GetChamber())
	}

	// Test round-trip
	jsonData, err := ToJSON(amendment)
	if err != nil {
		t.Fatalf("failed to marshal amendment to JSON: %v", err)
	}

	amendment2, err := EngrossedAmendmentFromJSON(jsonData)
	if err != nil {
		t.Fatalf("failed to parse amendment from JSON: %v", err)
	}

	if amendment2.GetDocumentNumber() != amendment.GetDocumentNumber() {
		t.Errorf("doc number not preserved: got '%s', want '%s'", amendment2.GetDocumentNumber(), amendment.GetDocumentNumber())
	}
}

func TestDetectDocumentType(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		expected DocumentType
	}{
		{
			name:     "bill",
			file:     "BILLS-114s32cds.xml",
			expected: DocumentTypeBill,
		},
		{
			name:     "resolution",
			file:     "BILLS-116sres100ats.xml",
			expected: DocumentTypeResolution,
		},
		{
			name:     "engrossed amendment",
			file:     "BILLS-116hr1865eas.xml",
			expected: DocumentTypeEngrossedAmendment,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := os.ReadFile(filepath.Join("..", "..", "bill-version-samples-september-2024", tt.file))
			if err != nil {
				t.Fatalf("failed to read file: %v", err)
			}

			docType := DetectDocumentType(data)
			if docType != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, docType)
			}
		})
	}
}

func TestParseDocument(t *testing.T) {
	// Test automatic parsing of a bill
	data, err := os.ReadFile(filepath.Join("..", "..", "bill-version-samples-september-2024", "BILLS-114s32cds.xml"))
	if err != nil {
		t.Fatalf("failed to read sample bill: %v", err)
	}

	doc, err := ParseDocument(data)
	if err != nil {
		t.Fatalf("failed to parse document: %v", err)
	}

	// Verify it implements the interface
	if doc.GetDocumentNumber() != "32" {
		t.Errorf("expected doc number '32', got '%s'", doc.GetDocumentNumber())
	}

	// Type assert to Bill
	bill, ok := doc.(*Bill)
	if !ok {
		t.Fatal("expected document to be a *Bill")
	}

	// Verify we can access Bill-specific features
	sections := bill.GetSections()
	if len(sections) == 0 {
		t.Error("expected bill to have sections")
	}
}

func TestInterfacePolymorphism(t *testing.T) {
	// Load different document types
	billData, _ := os.ReadFile(filepath.Join("..", "..", "bill-version-samples-september-2024", "BILLS-114s32cds.xml"))
	resData, _ := os.ReadFile(filepath.Join("..", "..", "bill-version-samples-september-2024", "BILLS-116sres100ats.xml"))
	amendData, _ := os.ReadFile(filepath.Join("..", "..", "bill-version-samples-september-2024", "BILLS-116hr1865eas.xml"))

	bill, _ := ParseBill(billData)
	res, _ := ParseResolution(resData)
	amend, _ := ParseEngrossedAmendment(amendData)

	// Test that they all implement LegislativeDocument
	docs := []LegislativeDocument{bill, res, amend}

	for _, doc := range docs {
		// These should all work regardless of document type
		_ = doc.GetDocumentNumber()
		_ = doc.GetCongress()
		_ = doc.GetSession()
		_ = doc.GetTitle()
		_ = doc.GetStage()
		_ = doc.GetChamber()
		_ = doc.IsPublic()
		_ = doc.GetCitations()
	}

	// Test SponsoredDocument interface
	sponsoredDocs := []SponsoredDocument{bill, res}
	for _, doc := range sponsoredDocs {
		_ = doc.GetSponsors()
		_ = doc.GetCosponsors()
	}

	// Test ActionDocument interface
	actionDocs := []ActionDocument{bill, res, amend}
	for _, doc := range actionDocs {
		_ = doc.GetActions()
	}
}

func TestMetadataDocumentInterface(t *testing.T) {
	data, _ := os.ReadFile(filepath.Join("..", "..", "bill-version-samples-september-2024", "BILLS-114s32cds.xml"))
	bill, _ := ParseBill(data)

	// Test MetadataDocument interface
	var metaDoc MetadataDocument = bill

	creator := metaDoc.GetCreator()
	if creator != "United States Senate" {
		t.Errorf("expected creator 'United States Senate', got '%s'", creator)
	}

	publisher := metaDoc.GetPublisher()
	if publisher != "United States Government Publishing Office" {
		t.Errorf("expected GPO publisher, got '%s'", publisher)
	}

	language := metaDoc.GetLanguage()
	if language != "EN" {
		t.Errorf("expected language 'EN', got '%s'", language)
	}

	processedBy := metaDoc.GetProcessedBy()
	if processedBy == "" {
		t.Error("expected processedBy to be set")
	}

	processedDate := metaDoc.GetProcessedDate()
	if processedDate == "" {
		t.Error("expected processedDate to be set")
	}
}
