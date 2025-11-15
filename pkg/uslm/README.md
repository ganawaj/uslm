# USLM Go Package

Go structs for parsing and generating United States Legislative Markup (USLM) XML documents.

## Features

- **Complete XML/JSON round-trip support**: Parse XML to Go structs, marshal to JSON, unmarshal back to Go, marshal back to XML without data loss
- **Common interfaces**: Unified interfaces for accessing document properties across different document types (bills, resolutions, amendments)
- **Type-safe**: Strong typing with proper struct tags for both XML and JSON
- **Simple and readable**: Verbose but easy-to-follow code structure

## Supported Document Types

- **Bill** - Senate and House bills
- **Resolution** - Simple, joint, and concurrent resolutions
- **Amendment** - Amendment documents
- **EngrossedAmendment** - Engrossed amendment documents

## Installation

```bash
go get github.com/usgpo/uslm/pkg/uslm
```

## Quick Start

### Parsing Documents

```go
package main

import (
    "fmt"
    "os"
    "github.com/usgpo/uslm/pkg/uslm"
)

func main() {
    // Read XML file
    data, err := os.ReadFile("bill.xml")
    if err != nil {
        panic(err)
    }

    // Auto-detect and parse document type
    doc, err := uslm.ParseDocument(data)
    if err != nil {
        panic(err)
    }

    // Access common properties via interface
    fmt.Printf("Document Number: %s\n", doc.GetDocumentNumber())
    fmt.Printf("Congress: %s\n", doc.GetCongress())
    fmt.Printf("Session: %s\n", doc.GetSession())
    fmt.Printf("Title: %s\n", doc.GetTitle())
    fmt.Printf("Stage: %s\n", doc.GetStage())
    fmt.Printf("Chamber: %s\n", doc.GetChamber())
}
```

### Type-Specific Parsing

```go
// Parse specifically as a Bill
bill, err := uslm.ParseBill(data)
if err != nil {
    panic(err)
}

// Access Bill-specific features
sponsors := bill.GetSponsors()
for _, sponsor := range sponsors {
    fmt.Printf("Sponsor: %s (ID: %s)\n", sponsor.GetName(), sponsor.GetID())
}

cosponsors := bill.GetCosponsors()
for _, cosponsor := range cosponsors {
    fmt.Printf("Cosponsor: %s (ID: %s)\n", cosponsor.GetName(), cosponsor.GetID())
}

committees := bill.GetCommittees()
for _, committee := range committees {
    fmt.Printf("Committee: %s (ID: %s)\n", committee.GetName(), committee.GetID())
}

sections := bill.GetSections()
for _, section := range sections {
    fmt.Printf("Section %s: %s\n", section.GetNumValue(), section.GetHeading())
}
```

### JSON Serialization

```go
// Convert to JSON
jsonData, err := uslm.ToJSON(bill)
if err != nil {
    panic(err)
}

// Parse back from JSON
bill2, err := uslm.BillFromJSON(jsonData)
if err != nil {
    panic(err)
}
```

### Working with Interfaces

```go
// All document types implement common interfaces
func processDocument(doc uslm.LegislativeDocument) {
    fmt.Printf("Processing: %s\n", doc.GetTitle())
    fmt.Printf("Congress: %s, Session: %s\n", doc.GetCongress(), doc.GetSession())
    fmt.Printf("Stage: %s\n", doc.GetStage())
}

// Check for specific capabilities
if sponsored, ok := doc.(uslm.SponsoredDocument); ok {
    sponsors := sponsored.GetSponsors()
    fmt.Printf("Has %d sponsors\n", len(sponsors))
}

if actionDoc, ok := doc.(uslm.ActionDocument); ok {
    actions := actionDoc.GetActions()
    fmt.Printf("Has %d actions\n", len(actions))
}

if metaDoc, ok := doc.(uslm.MetadataDocument); ok {
    fmt.Printf("Creator: %s\n", metaDoc.GetCreator())
    fmt.Printf("Publisher: %s\n", metaDoc.GetPublisher())
}
```

## Available Interfaces

### LegislativeDocument
Common interface for all legislative documents:
- `GetDocumentNumber()` - Document number
- `GetDocumentType()` - Type (e.g., "Senate Bill")
- `GetCongress()` - Congress number
- `GetSession()` - Session number
- `GetTitle()` - Full title
- `GetStage()` - Document stage
- `GetChamber()` - Current chamber
- `IsPublic()` - Public or private
- `GetCitations()` - Citable forms

### SponsoredDocument
For documents with sponsors:
- `GetSponsors()` - Primary sponsors
- `GetCosponsors()` - Cosponsors

### ActionDocument
For documents with legislative actions:
- `GetActions()` - All actions

### CommitteeDocument
For documents referencing committees:
- `GetCommittees()` - Referenced committees

### HierarchicalDocument
For documents with sectional structure:
- `GetSections()` - Top-level sections

### MetadataDocument
For accessing Dublin Core metadata:
- `GetCreator()` - Document creator
- `GetPublisher()` - Publisher
- `GetLanguage()` - Language code
- `GetRights()` - Rights statement
- `GetProcessedBy()` - Processing tool
- `GetProcessedDate()` - Processing date

### AmendmentDocument
For amendment-specific features:
- `GetAmendmentDegree()` - Degree of amendment

## Structure Overview

```
pkg/uslm/
├── interfaces.go    - Common interfaces
├── common.go        - Shared types (Inline, Content, etc.)
├── metadata.go      - Meta and AmendMeta structs
├── preface.go       - Preface elements (Actions, Sponsors, etc.)
├── content.go       - Main content (Sections, Paragraphs, etc.)
├── documents.go     - Root document types (Bill, Resolution, etc.)
├── parser.go        - Parsing and marshaling helpers
└── parser_test.go   - Tests
```

## Testing

```bash
cd pkg/uslm
go test -v
```

## License

Same as the USLM schema - public domain per Title 17 Section 105 of the United States Code.
