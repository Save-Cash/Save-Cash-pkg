package testing

import (
	"fmt"
	"testing"
    "github.com/Save-Cash/Save-Cash-pkg/models"
)

func TestGetAllUsers(t *testing.T) {
	data, err := models.GetAllUsers()
	if err != nil {
		t.Errorf("Failed to retrieve users: %v", err)
	} else if len(data) == 0 {
		t.Errorf("No users found")
	} else {
		fmt.Printf("Retrieved users: %v\n", data)
	}
}