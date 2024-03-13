package chapter9

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func Solve() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		errs := ValidateEmployee(emp)
		fmt.Printf("record %d: %+v\n", count, emp)
		for _, err := range errs {
			if errors.Is(err, ErrInvalidID) {
				fmt.Println("error: There is an Employee with an Invalid ID")
			}
			var errFieldEmpty ErrFieldEmpty
			if errors.As(err, &errFieldEmpty) {
				fmt.Println("error: ", errFieldEmpty.Message, " in the field ", errFieldEmpty.Field)
			}
		}
	}
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

type ErrFieldEmpty struct {
	Message string
	Field   string
}

func (e ErrFieldEmpty) Error() string {
	return e.Message + " " + e.Field
}

func NewErrFieldEmpty(message string, field string) ErrFieldEmpty {
	return ErrFieldEmpty{Message: message, Field: field}
}

var (
	validID = regexp.MustCompile(`\w{4}-\d{3}`)
)

var ErrInvalidID = errors.New("invalid ID")

func ValidateEmployee(e Employee) (err []error) {
	if len(e.ID) == 0 {
		err = append(err, ErrFieldEmpty{"missing", "ID"})
	} else if !validID.MatchString(e.ID) {
		err = append(err, ErrInvalidID)
	}
	if len(e.FirstName) == 0 {
		err = append(err, NewErrFieldEmpty("missing", "FirstName"))
	}
	if len(e.LastName) == 0 {
		err = append(err, NewErrFieldEmpty("missing", "LastName"))
	}
	if len(e.Title) == 0 {
		err = append(err, NewErrFieldEmpty("missing", "Title"))
	}
	return err
}
