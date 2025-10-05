package novakeytypes

import (
	"github.com/google/uuid"
)

type User struct {
	Id 				uuid.UUID 		`json:"id" yaml:"id"`
	Email  		string 				`json:"email" yaml:"email"`
	Username  string 				`json:"username" yaml:"username"`
	PublicKey string 				`json:"publicKey" yaml:"publicKey"`
	Workspaces []Workspace 	`json:"workspaces,omitempty" yaml:"workspaces,omitempty"`
	Projects   []Project 		`json:"projects,omitempty" yaml:"projects,omitempty"`	
}	

type AuthEntity struct {
	Id 				uuid.UUID `json:"id,omitempty"`
	Username  string 		`json:"username,omitempty"`
	PublicKey string 		`json:"publicKey"`
	Signature string 		`json:"signature"`
	Message   string 		`json:"message"`  
	Timestamp int64  		`json:"timestamp"`  
	Password 	string 		`json:"password,omitempty"`
}

type SetUserRequest struct {	
	AuthEntity
	Email						string   		`json:"email,omitempty"`
	Workspaces   		[]Workspace `json:"workspaces,omitempty"`
	Projects   			[]Project 	`json:"projects,omitempty"`
	Signer					AuthEntity 	`json:"signer"`
}

type SetUserResponse struct {	
	Id								uuid.UUID   `json:"id,omitempty"`	
	Username					string  		`json:"username,omitempty"`
	Password					string   		`json:"password,omitempty"`
	ErrorResponse
}

type DeleteUserRequest struct {
	AuthEntity
}

type DeleteUserResponse struct {
	Id								uuid.UUID   `json:"id,omitempty"`		
  ErrorResponse
}