/*
 * Todolist RESTful API
 *
 * OpenAPI for Todolist RESTful API
 *
 * API version: 1
 * Contact: haris@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Todolist struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Priority float64 `json:"priority,omitempty"`
	Tags []string `json:"tags,omitempty"`
}
