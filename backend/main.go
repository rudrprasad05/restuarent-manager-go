// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"restuarent-manager-go/controllers"
// 	"restuarent-manager-go/middleware"
// 	"restuarent-manager-go/models"
// 	"strconv"
// 	"sync"

// )

// // Define a struct to represent the data model (e.g., Item)
// type Item struct {
//     ID    int    `json:"id"`
//     Name  string `json:"name"`
//     Price float64 `json:"price"`
// }

// // Create a slice to act as an in-memory database
// var items = []Item{}
// var nextID = 1
// var mu sync.Mutex // Mutex for thread-safe access

// // Get all items (GET /items)
// func getItems(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(items)
// }

// // Get a specific item by ID (GET /items/{id})
// func getItem(w http.ResponseWriter, r *http.Request) {
// 	idString, err := strconv.Atoi(r.PathValue("id"))
//     w.Header().Set("Content-Type", "application/json")

//     if err != nil {
//         http.Error(w, "Invalid item ID", http.StatusBadRequest)
//         return
//     }
//     for _, item := range items {
//         if item.ID == idString {
//             json.NewEncoder(w).Encode(item)
//             return
//         }
//     }
//     http.Error(w, "Item not found", http.StatusNotFound)
// }

// // Create a new item (POST /items)
// func createItem(w http.ResponseWriter, r *http.Request) {
//     var newItem Item
//     err := json.NewDecoder(r.Body).Decode(&newItem)
//     if err != nil {
//         http.Error(w, "Invalid request body", http.StatusBadRequest)
//         return
//     }
//     mu.Lock()
//     newItem.ID = nextID
//     nextID++
//     items = append(items, newItem)
//     mu.Unlock()
//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(newItem)
// }

// // Update an item by ID (PUT /items/{id})
// func updateItem(w http.ResponseWriter, r *http.Request) {

//     id, err := strconv.Atoi(r.URL.Path[len("/items/"):])
//     if err != nil {
//         http.Error(w, "Invalid item ID", http.StatusBadRequest)
//         return
//     }

//     var updatedItem Item
//     err = json.NewDecoder(r.Body).Decode(&updatedItem)
//     if err != nil {
//         http.Error(w, "Invalid request body", http.StatusBadRequest)
//         return
//     }

//     mu.Lock()
//     defer mu.Unlock()

//     for i, item := range items {
//         if item.ID == id {
//             updatedItem.ID = id
//             items[i] = updatedItem
//             w.Header().Set("Content-Type", "application/json")
//             json.NewEncoder(w).Encode(updatedItem)
//             return
//         }
//     }
//     http.Error(w, "Item not found", http.StatusNotFound)
// }

// // Delete an item by ID (DELETE /items/{id})
// func deleteItem(w http.ResponseWriter, r *http.Request) {
//     id, err := strconv.Atoi(r.URL.Path[len("/items/"):])
//     if err != nil {
//         http.Error(w, "Invalid item ID", http.StatusBadRequest)
//         return
//     }

//     mu.Lock()
//     defer mu.Unlock()

//     for i, item := range items {
//         if item.ID == id {
//             items = append(items[:i], items[i+1:]...) // Remove item from the slice
//             w.WriteHeader(http.StatusNoContent)      // Return 204 No Content
//             return
//         }

//     }
//     http.Error(w, "Item not found", http.StatusNotFound)
// }

// func main() {
//     router := gin.Default()

//     // new `GET /users` route associated with our `getUsers` function
//     router.GET("/users", getUsers)
//     mux := http.NewServeMux()
// 	mux.HandleFunc("/orders", routes.GetOrdersHandler)

// 	// Apply the CORS middleware
// 	handlerWithCORS := middleware.CORS(mux)

// 	// Start the server
// 	log.Println("Starting server on :8080")
// 	if err := http.ListenAndServe(":8080", handlerWithCORS); err != nil {
// 		log.Fatal("Failed to start server:", err)
// 	}
// }

// func insert(){
//     newOrder := models.Order{
// 		Customer:    "John Doe",
// 		Amount:      99.99,
// 		OrderStatus: "Pending",
// 	}
// 	err := controllers.InsertOrder(newOrder)
//     log.Println(err)
// }

package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct{
    ID          string  `json:"id"`
    Item        string  `json:"item"`
    Completed   bool    `json:"completed"`

}

var todos = []todo{
    {ID: "1", Item: "item1", Completed: false},
    {ID: "2", Item: "item2", Completed: true},
    {ID: "3", Item: "item3", Completed: false},
}

func getTodos(context *gin.Context){
    context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context){
    var newTodo todo
    log.Println("started")

    if err := context.BindJSON(&newTodo); err != nil {
        log.Println("error")
        return
    }
    log.Println("ended")

    todos = append(todos, newTodo)
    context.IndentedJSON(http.StatusCreated, newTodo)

}

func getTodoById(id string)(*todo, error){
    for i, t := range todos {
        if t.ID == id {
            return &todos[i], nil
        } 
    }

    return nil, errors.New("Todo not found")
}

func getTodo(context *gin.Context){
    id := context.Param("id")
    todo, err := getTodoById(id)

    if err != nil{
        context.IndentedJSON(http.StatusNotFound, gin.H{"message" : "not found todo"})
        return
    }

    context.IndentedJSON(http.StatusOK, todo)
}

func patchTodo(context *gin.Context){
    id := context.Param("id")
    todo, err := getTodoById(id)

    if err != nil{
        context.IndentedJSON(http.StatusNotFound, gin.H{"message" : "not found todo"})
        return
    }

    todo.Completed = !todo.Completed

    context.IndentedJSON(http.StatusOK, todo)
}

func main(){
    router := gin.Default()
    router.GET("/todos", getTodos)
    router.POST("/todos", addTodo)
    router.GET("/todos/:id", getTodo)
    router.PATCH("/todos/:id", patchTodo)

    router.Run("localhost:9090")
}