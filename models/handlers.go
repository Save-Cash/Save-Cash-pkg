package models

import (
	"context"
	"fmt"
	"github.com/google/uuid"
    "github.com/Save-Cash/Save-Cash-pkg/db"
)

func CreateUser(user User) (*User, error) {
	id := uuid.New()
	query := `INSERT INTO users (id, name, email) VALUES ($1, $2, $3) RETURNING id, name, email`
	row := db.GetDB().QueryRow(context.Background(), query, id, user.Name, user.Email)

	var createdUser User
	if err := row.Scan(&createdUser.ID, &createdUser.Name, &createdUser.Email); err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return &createdUser, nil
}

func GetUserByID(id uuid.UUID) (*User, error) {
	query := `SELECT id, name, email FROM users WHERE id = $1`
	row := db.GetDB().QueryRow(context.Background(), query, id)

	var user User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	return &user, nil
}

func GetAllUsers() ([]User, error) {
	query := `SELECT id, name, email FROM users`
	rows, err := db.GetDB().Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateTransactionCategory(category TransactionCategory) (*TransactionCategory, error) {
	id := uuid.New()
	query := `INSERT INTO transaction_categories (id, name, type) VALUES ($1, $2, $3) RETURNING id, name, type`
	row := db.GetDB().QueryRow(context.Background(), query, id, category.Name, category.Type)

	var createdCategory TransactionCategory
	if err := row.Scan(&createdCategory.ID, &createdCategory.Name, &createdCategory.Type); err != nil {
		return nil, fmt.Errorf("failed to create category: %v", err)
	}

	return &createdCategory, nil
}

func GetTransactionCategoryByID(id uuid.UUID) (*TransactionCategory, error) {
	query := `SELECT id, name, type FROM transaction_categories WHERE id = $1`
	row := db.GetDB().QueryRow(context.Background(), query, id)

	var category TransactionCategory
	if err := row.Scan(&category.ID, &category.Name, &category.Type); err != nil {
		return nil, fmt.Errorf("category not found: %v", err)
	}

	return &category, nil
}

func GetAllTransactionCategories() ([]TransactionCategory, error) {
	query := `SELECT id, name, type FROM transaction_categories`
	rows, err := db.GetDB().Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %v", err)
	}
	defer rows.Close()

	var categories []TransactionCategory
	for rows.Next() {
		var category TransactionCategory
		if err := rows.Scan(&category.ID, &category.Name, &category.Type); err != nil {
			return nil, fmt.Errorf("failed to scan category: %v", err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func CreateFinancialTransaction(transaction FinancialTransaction) (*FinancialTransaction, error) {
	id := uuid.New()
	query := `INSERT INTO financial_transactions (id, user_id, category_id, amount, description, date) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, user_id, category_id, amount, description, date`
	row := db.GetDB().QueryRow(context.Background(), query, id, transaction.UserID, transaction.CategoryID, transaction.Amount, transaction.Description, transaction.Date)

	var createdTransaction FinancialTransaction
	if err := row.Scan(&createdTransaction.ID, &createdTransaction.UserID, &createdTransaction.CategoryID, &createdTransaction.Amount, &createdTransaction.Description, &createdTransaction.Date); err != nil {
		return nil, fmt.Errorf("failed to create financial transaction: %v", err)
	}

	return &createdTransaction, nil
}

func GetFinancialTransactionByID(id uuid.UUID) (*FinancialTransaction, error) {
	query := `SELECT id, user_id, category_id, amount, description, date FROM financial_transactions WHERE id = $1`
	row := db.GetDB().QueryRow(context.Background(), query, id)

	var transaction FinancialTransaction
	if err := row.Scan(&transaction.ID, &transaction.UserID, &transaction.CategoryID, &transaction.Amount, &transaction.Description, &transaction.Date); err != nil {
		return nil, fmt.Errorf("transaction not found: %v", err)
	}

	return &transaction, nil
}

func GetFinancialTransactionsByUserID(userID uuid.UUID) ([]FinancialTransaction, error) {
	query := `SELECT id, user_id, category_id, amount, description, date FROM financial_transactions WHERE user_id = $1`
	rows, err := db.GetDB().Query(context.Background(), query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %v", err)
	}
	defer rows.Close()

	var transactions []FinancialTransaction
	for rows.Next() {
		var transaction FinancialTransaction
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.CategoryID, &transaction.Amount, &transaction.Description, &transaction.Date); err != nil {
			return nil, fmt.Errorf("failed to scan transaction: %v", err)
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
