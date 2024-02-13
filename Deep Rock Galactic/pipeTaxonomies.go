package main

import "fmt"

// Pipe represents a pipe in the taxonomy
type Pipe struct {
	Name     string
	Category string
}

// Taxonomy represents the entire taxonomy of pipes
type Taxonomy struct {
	Categories map[string][]string
}

// NewTaxonomy creates a new instance of Taxonomy
func NewTaxonomy() *Taxonomy {
	return &Taxonomy{
		Categories: make(map[string][]string),
	}
}

// AddPipe adds a new pipe to the taxonomy
func (t *Taxonomy) AddPipe(pipe Pipe) {
	t.Categories[pipe.Category] = append(t.Categories[pipe.Category], pipe.Name)
}

// PrintTaxonomy prints the taxonomy
func (t *Taxonomy) PrintTaxonomy() {
	for category, pipes := range t.Categories {
		fmt.Println("Category:", category)
		fmt.Println("Pipes:", pipes)
	}
}

func main() {
	// Create a new taxonomy
	taxonomy := NewTaxonomy()

	// Add pipes to the taxonomy
	taxonomy.AddPipe(Pipe{Name: "Main Line", Category: "Main Pipes"})
	taxonomy.AddPipe(Pipe{Name: "Branch Line", Category: "Branch Pipes"})
	taxonomy.AddPipe(Pipe{Name: "Vertical Line", Category: "Vertical Pipes"})
	taxonomy.AddPipe(Pipe{Name: "Supply Line", Category: "Supply Pipes"})
	taxonomy.AddPipe(Pipe{Name: "Extractor Line", Category: "Extractor Pipes"})
	taxonomy.AddPipe(Pipe{Name: "Power Line", Category: "Power Pipes"})

	// Print the taxonomy
	taxonomy.PrintTaxonomy()
}
