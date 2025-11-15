// Example program demonstrating USLM package usage.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/usgpo/uslm/pkg/uslm"
)

func main() {
	// Parse a sample bill
	billPath := filepath.Join("..", "..", "..", "bill-version-samples-september-2024", "BILLS-114s32cds.xml")
	billData, err := os.ReadFile(billPath)
	if err != nil {
		fmt.Printf("Error reading bill: %v\n", err)
		return
	}

	fmt.Println("=== Parsing Bill ===")
	bill, err := uslm.ParseBill(billData)
	if err != nil {
		fmt.Printf("Error parsing bill: %v\n", err)
		return
	}

	// Display basic information
	fmt.Printf("Document Number: %s\n", bill.GetDocumentNumber())
	fmt.Printf("Document Type: %s\n", bill.GetDocumentType())
	fmt.Printf("Congress: %s\n", bill.GetCongress())
	fmt.Printf("Session: %s\n", bill.GetSession())
	fmt.Printf("Stage: %s\n", bill.GetStage())
	fmt.Printf("Chamber: %s\n", bill.GetChamber())
	fmt.Printf("Public: %v\n", bill.IsPublic())
	fmt.Println()

	// Display title
	fmt.Println("=== Title ===")
	fmt.Printf("%s\n\n", bill.GetTitle())

	// Display citations
	fmt.Println("=== Citations ===")
	for _, citation := range bill.GetCitations() {
		fmt.Printf("  - %s\n", citation)
	}
	fmt.Println()

	// Display sponsors
	fmt.Println("=== Sponsors ===")
	for _, sponsor := range bill.GetSponsors() {
		fmt.Printf("  Primary: %s (ID: %s)\n", sponsor.GetName(), sponsor.GetID())
	}
	fmt.Println()

	// Display cosponsors
	fmt.Println("=== Cosponsors ===")
	for _, cosponsor := range bill.GetCosponsors() {
		fmt.Printf("  - %s (ID: %s)\n", cosponsor.GetName(), cosponsor.GetID())
	}
	fmt.Println()

	// Display committees
	fmt.Println("=== Committees ===")
	for _, committee := range bill.GetCommittees() {
		fmt.Printf("  - %s (ID: %s)\n", committee.GetName(), committee.GetID())
	}
	fmt.Println()

	// Display actions
	fmt.Println("=== Actions ===")
	for i, action := range bill.GetActions() {
		if action.Date != nil {
			fmt.Printf("  %d. Date: %s\n", i+1, action.Date.Date)
		}
		if action.ActionDescription != nil {
			fmt.Printf("     Description: %s\n", action.ActionDescription.Text)
		}
	}
	fmt.Println()

	// Display sections
	fmt.Println("=== Sections ===")
	for _, section := range bill.GetSections() {
		fmt.Printf("  Section %s: %s\n", section.GetNumValue(), section.GetHeading())
	}
	fmt.Println()

	// Display metadata
	fmt.Println("=== Metadata ===")
	fmt.Printf("Creator: %s\n", bill.GetCreator())
	fmt.Printf("Publisher: %s\n", bill.GetPublisher())
	fmt.Printf("Language: %s\n", bill.GetLanguage())
	fmt.Printf("Processed By: %s\n", bill.GetProcessedBy())
	fmt.Printf("Processed Date: %s\n", bill.GetProcessedDate())
	fmt.Println()

	// Demonstrate JSON serialization
	fmt.Println("=== JSON Serialization Demo ===")
	jsonData, err := uslm.ToJSON(bill)
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
		return
	}

	// Show a snippet of the JSON
	var jsonObj map[string]interface{}
	json.Unmarshal(jsonData, &jsonObj)
	if meta, ok := jsonObj["meta"].(map[string]interface{}); ok {
		fmt.Printf("JSON Meta Section:\n")
		fmt.Printf("  dcTitle: %s\n", meta["dcTitle"])
		fmt.Printf("  docNumber: %s\n", meta["docNumber"])
		fmt.Printf("  congress: %s\n", meta["congress"])
		fmt.Printf("  session: %s\n", meta["session"])
	}
	fmt.Println()

	// Round-trip demonstration
	fmt.Println("=== Round-Trip Demo ===")
	bill2, err := uslm.BillFromJSON(jsonData)
	if err != nil {
		fmt.Printf("Error parsing from JSON: %v\n", err)
		return
	}
	fmt.Printf("Original Document Number: %s\n", bill.GetDocumentNumber())
	fmt.Printf("After Round-Trip: %s\n", bill2.GetDocumentNumber())
	fmt.Printf("Match: %v\n", bill.GetDocumentNumber() == bill2.GetDocumentNumber())
	fmt.Println()

	// Parse a resolution
	fmt.Println("=== Parsing Resolution ===")
	resPath := filepath.Join("..", "..", "..", "bill-version-samples-september-2024", "BILLS-116sres100ats.xml")
	resData, err := os.ReadFile(resPath)
	if err != nil {
		fmt.Printf("Error reading resolution: %v\n", err)
		return
	}

	resolution, err := uslm.ParseResolution(resData)
	if err != nil {
		fmt.Printf("Error parsing resolution: %v\n", err)
		return
	}

	fmt.Printf("Document Number: %s\n", resolution.GetDocumentNumber())
	fmt.Printf("Document Type: %s\n", resolution.GetDocumentType())
	fmt.Printf("Congress: %s\n", resolution.GetCongress())
	fmt.Printf("Number of Cosponsors: %d\n", len(resolution.GetCosponsors()))
	fmt.Println()

	// Parse an amendment
	fmt.Println("=== Parsing Engrossed Amendment ===")
	amendPath := filepath.Join("..", "..", "..", "bill-version-samples-september-2024", "BILLS-116hr1865eas.xml")
	amendData, err := os.ReadFile(amendPath)
	if err != nil {
		fmt.Printf("Error reading amendment: %v\n", err)
		return
	}

	amendment, err := uslm.ParseEngrossedAmendment(amendData)
	if err != nil {
		fmt.Printf("Error parsing amendment: %v\n", err)
		return
	}

	fmt.Printf("Document Number: %s\n", amendment.GetDocumentNumber())
	fmt.Printf("Document Type: %s\n", amendment.GetDocumentType())
	fmt.Printf("Congress: %s\n", amendment.GetCongress())
	fmt.Printf("Amendment Degree: %s\n", amendment.GetAmendmentDegree())
	fmt.Println()

	// Demonstrate polymorphism with interfaces
	fmt.Println("=== Interface Polymorphism Demo ===")
	docs := []uslm.LegislativeDocument{bill, resolution, amendment}
	for i, doc := range docs {
		fmt.Printf("%d. %s %s - %s\n",
			i+1,
			doc.GetDocumentType(),
			doc.GetDocumentNumber(),
			doc.GetChamber())
	}
}
