package main

import (
	expences "financetracker/internal/expances"
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("type a command: add, delete, update, list, summary")
		return
	}

	command := os.Args[1]
	os.Args = append(os.Args[:1], os.Args[2:]...)

	description := flag.String("description", "", "description of expence")
	amount := flag.Int("amount", 0, "amount of expence")
	id := flag.Int("id", -1, "id of expence")
	month := flag.Int("month", 0, "month")
	category := flag.String("category", "", "category of expence")
	file := flag.String("file", "expences.csv", "save csv to file")

	flag.Parse()

	if err := expences.Load(); err != nil {
		fmt.Println(err)
		return
	}

	switch command {
	case "add":
		err := expences.Add(expences.NewExpence(*description, *category, *amount))
		if err != nil {
			fmt.Println(err)
			return
		}
	case "delete":
		if err := expences.Delete(*id); err != nil {
			fmt.Println("delete: ", err)
		}
	case "update":
		if err := expences.Update(*id, *description, *amount); err != nil {
			fmt.Println("update: ", err)
		}
	case "summary":
		expences.Summary(*month)
	case "list":
		expences.ViewExpences(*category)
	case "export":
		if err := expences.ExportCSV(*file); err != nil {
			fmt.Println("export: ", err)
			return
		}
	default:
		fmt.Println("unknown command")
	}
	if err := expences.Save(); err != nil {
		fmt.Println(err)
		return
	}
}
