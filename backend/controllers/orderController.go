package controllers

import (
	"fmt"
	"log"
	"restuarent-manager-go/database"
	"restuarent-manager-go/models"
)

// CreateOrdersTable creates the orders table using raw SQL
func CreateOrdersTable() {
	// Use the CreateDBConnection function from the database package
	db, err := database.CreateDBConnection()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	query := `
	CREATE TABLE IF NOT EXISTS orders (
		id INT AUTO_INCREMENT PRIMARY KEY,
		customer VARCHAR(100),
		amount DECIMAL(10,2),
		order_status VARCHAR(50)
	);`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Error creating orders table:", err)
	}
	fmt.Println("Orders table created successfully.")
}

// InsertOrder inserts a new order into the orders table
func InsertOrder(order models.Order) error {
	// Use the CreateDBConnection function from the database package
	db, err := database.CreateDBConnection()
	if err != nil {
		return fmt.Errorf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	query := `
	INSERT INTO orders (customer, amount, order_status) 
	VALUES (?, ?, ?);`
	result, err := db.Exec(query, order.Customer, order.Amount, order.OrderStatus)
	if err != nil {
		return fmt.Errorf("Error inserting new order: %v", err)
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Error fetching last insert ID: %v", err)
	}
	fmt.Printf("New order inserted with ID: %d\n", lastInsertID)
	order.ID = int(lastInsertID)
	return nil
}

// GetOrders fetches all orders from the database
func GetOrders() ([]models.Order, error) {
    db, err := database.CreateDBConnection()
    if err != nil {
        log.Printf("Error connecting to the database: %v", err)
        return nil, fmt.Errorf("error connecting to the database: %v", err)
    }
    defer db.Close()

	query := `SELECT id, customer, amount, order_status FROM orders`
    rows, err := db.Query(query)
    if err != nil {
        log.Printf("Error executing query: %v", err)
        return nil, fmt.Errorf("error executing query: %v", err)
    }
    defer rows.Close()

    var orders []models.Order

    for rows.Next() {
        var order models.Order
        err := rows.Scan(&order.ID, &order.Customer, &order.Amount, &order.OrderStatus)
        if err != nil {
            log.Printf("Error scanning row: %v", err)
            return nil, fmt.Errorf("error scanning row: %v", err)
        }
        orders = append(orders, order)
    }

    // Check if any errors occurred during row iteration
    if err := rows.Err(); err != nil {
        log.Printf("Error iterating over rows: %v", err)
        return nil, fmt.Errorf("error iterating over rows: %v", err)
    }

    return orders, nil
}
