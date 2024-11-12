package testing

import (
	"testing"
	"github.com/google/uuid"
	"github.com/Save-Cash/Save-Cash-pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	user := models.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	createdUser, err := models.CreateUser(user)

	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}

	assert.NotNil(t, createdUser)
	assert.Equal(t, "John Doe", createdUser.Name)
	assert.Equal(t, "john.doe@example.com", createdUser.Email)
	assert.True(t, createdUser.ID != uuid.Nil, "User ID should not be empty")

	fetchedUser, err := models.GetUserByID(createdUser.ID)
	if err != nil {
		t.Errorf("Error fetching user: %v", err)
	}
	assert.Equal(t, createdUser.ID, fetchedUser.ID)
}

func TestGetAllUsers(t *testing.T) {
	users, err := models.GetAllUsers()

	if err != nil {
		t.Errorf("Error retrieving users: %v", err)
	}

	if len(users) == 0 {
		t.Errorf("No users found")
	}

	t.Logf("Retrieved users: %v", users)
}

func TestGetUserByID(t *testing.T) {
	user := models.User{
		Name:  "Jane Doe",
		Email: "jane.doe@example.com",
	}

	createdUser, err := models.CreateUser(user)
	if err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	fetchedUser, err := models.GetUserByID(createdUser.ID)

	if err != nil {
		t.Errorf("Error retrieving user by ID: %v", err)
	}
	if fetchedUser.ID != createdUser.ID {
		t.Errorf("Expected user ID %v, but got %v", createdUser.ID, fetchedUser.ID)
	}
}

func TestCreateTransactionCategory(t *testing.T) {
	category := models.TransactionCategory{
		Name: "Groceries",
		Type: "Expense",
	}

	createdCategory, err := models.CreateTransactionCategory(category)
	if err != nil {
		t.Errorf("Error creating category: %v", err)
	}

	assert.NotNil(t, createdCategory)
	assert.Equal(t, "Groceries", createdCategory.Name)
	assert.Equal(t, "Expense", createdCategory.Type)
}

func TestGetTransactionCategoryByID(t *testing.T) {
	category := models.TransactionCategory{
		Name: "Entertainment",
		Type: "Expense",
	}

	createdCategory, err := models.CreateTransactionCategory(category)
	if err != nil {
		t.Fatalf("Error creating category: %v", err)
	}

	fetchedCategory, err := models.GetTransactionCategoryByID(createdCategory.ID)

	if err != nil {
		t.Errorf("Error retrieving category by ID: %v", err)
	}
	if fetchedCategory.ID != createdCategory.ID {
		t.Errorf("Expected category ID %v, but got %v", createdCategory.ID, fetchedCategory.ID)
	}
}
