package simpletable

import (
	"testing"
)

// All of these really just test that things don't explode - verification of proper output and order is manual.

func TestPrintStringSlice(t *testing.T) {
	table, err := New([]string{"FIRST", "MIDDLE", "LAST"}, [][]string{
		[]string{"Thomas", "A", "Swift"},
		[]string{"John", "F", "Kennedy"},
		[]string{"Susan", "B", "Anthony"},
		[]string{"Anduin", "L", "Wrynn"},
		[]string{"James", "T", "Kirk"},
	})
	if err != nil {
		t.Fatal(err)
	}
	table.Sort(2)
	err = table.Print()
	if err != nil {
		t.Fatal(err)
	}
}

func TestPrintNumberSlice(t *testing.T) {
	table, err := New([]string{"X", "Y", "Z"}, [][]int{
		[]int{-1, -2, 3},
		[]int{0, 1, 2},
		[]int{293847279834, 23, 4},
	})
	if err != nil {
		t.Fatal(err)
	}
	err = table.Print()
	if err != nil {
		t.Fatal(err)
	}
}

func TestPrintStruct(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}
	table, err := New(HeadersForType(person{}), []person{
		person{
			Name: "George",
			Age:  36,
		},
		person{
			Name: "Mark",
			Age:  12,
		},
		person{
			Name: "Johnny",
			Age:  8,
		},
		person{
			Name: "Jane",
			Age:  30,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	table.Sort(0)
	err = table.Print()
	if err != nil {
		t.Fatal(err)
	}
}
