package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

// Define Employee model
type Employee struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Company string `json:"company"`
	Health  string `json:"health"`
}

// Sample data
var employees = []Employee{
	{ID: 1, Name: "Rohit", Age: 30, Company: "Lupin", Health: "Fit"},
	{ID: 2, Name: "Aarti", Age: 28, Company: "JK Tyres", Health: "Needs checkup"},
}

// Define GraphQL Schema
func createSchema() graphql.Schema {
	employeeType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Employee",
		Fields: graphql.Fields{
			"id":      &graphql.Field{Type: graphql.Int},
			"name":    &graphql.Field{Type: graphql.String},
			"age":     &graphql.Field{Type: graphql.Int},
			"company": &graphql.Field{Type: graphql.String},
			"health":  &graphql.Field{Type: graphql.String},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"employees": &graphql.Field{
				Type: graphql.NewList(employeeType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return employees, nil
				},
			},
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"addEmployee": &graphql.Field{
				Type: employeeType,
				Args: graphql.FieldConfigArgument{
					"name":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"age":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					"company": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"health":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					newEmp := Employee{
						ID:      len(employees) + 1,
						Name:    p.Args["name"].(string),
						Age:     p.Args["age"].(int),
						Company: p.Args["company"].(string),
						Health:  p.Args["health"].(string),
					}
					employees = append(employees, newEmp)
					return newEmp, nil
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	return schema
}

// Define GraphQL endpoint with http server
func main() {
	schema := createSchema()

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Query string `json:"query"`
		}
		json.NewDecoder(r.Body).Decode(&params)

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: params.Query,
		})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("GraphQL server started at http://localhost:8080/graphql")
	http.ListenAndServe(":8080", nil)
}
