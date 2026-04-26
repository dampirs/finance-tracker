package expences

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

var storage []*Expence = []*Expence{}
var categories map[string]bool = map[string]bool{
	"food":          true,
	"transport":     true,
	"health":        true,
	"entertainment": true,
	"other":         true,
}

func validateCategories(category string) bool {
	return categories[category]
}

func Add(e *Expence) error {
	if !validateCategories(e.Category) {
		return fmt.Errorf("invalid category: %s\nAvailable: food, transport, health, entertainment, other", e.Category)
	}

	storage = append(storage, e)
	return nil
}

func Delete(id int) error {
	if id < 0 {
		return fmt.Errorf("delete: invalid id")
	}
	for i, e := range storage {
		if e.Id == id {
			storage = append(storage[:i], storage[i+1:]...)
			break
		}
	}
	return nil
}

func Update(id int, description string, amout int) error {
	if id < 0 {
		return fmt.Errorf("delete: invalid id")
	}

	for i, e := range storage {
		if e.Id == id {
			storage[i].Description = description
			storage[i].Amount = amout
			break
		}
	}
	return nil
}

func Summary(month int) {
	total := 0
	for _, e := range storage {
		if month == 0 || int(e.Time.Month()) == month {
			total += e.Amount
		}
	}
	if month == 0 {
		fmt.Printf("Total expences: $%d\n", total)
	} else {
		fmt.Printf("Total expences for %s: $%d", time.Month(month), total)
	}
}

func ViewExpences(category string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintln(w, "ID\tDate\tCategory\tDescription\tAmount")
	for _, e := range storage {
		if e.Category == category || category == "" {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%d\n", e.Id, e.Time.Format("2006-01-02"), e.Category, e.Description, e.Amount)
		}
	}
	w.Flush()
}

func ExportCSV(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("exportCSV: %w", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write([]string{"id", "date", "description", "amount", "category"})

	for _, e := range storage {
		w.Write([]string{
			fmt.Sprintf("%d", e.Id),
			e.Time.Format("2006-01-02"),
			e.Description,
			fmt.Sprintf("%d", e.Amount),
			e.Category,
		})
	}

	return nil
}
