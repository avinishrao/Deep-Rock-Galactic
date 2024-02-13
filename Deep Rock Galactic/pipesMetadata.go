package main

import "fmt"

// ResourceMetadata represents the metadata of a resource in Deep Rock Galactic
type ResourceMetadata struct {
	Name     string
	Type     string
	Location string
	Owner    string
}

// NewResourceMetadata creates a new instance of ResourceMetadata
func NewResourceMetadata(name, resourceType, location, owner string) *ResourceMetadata {
	return &ResourceMetadata{
		Name:     name,
		Type:     resourceType,
		Location: location,
		Owner:    owner,
	}
}

// PrintMetadata prints the metadata of a resource
func (m *ResourceMetadata) PrintMetadata() {
	fmt.Println("Resource Name:", m.Name)
	fmt.Println("Resource Type:", m.Type)
	fmt.Println("Location:", m.Location)
	fmt.Println("Owner:", m.Owner)
}

func main() {
	// Create metadata for a resource in Deep Rock Galactic
	resourceMetadata := NewResourceMetadata("Small Pipe", "Oil Pipe", "Cave System A", "Miner123")

	// Print the metadata
	resourceMetadata.PrintMetadata()
}
