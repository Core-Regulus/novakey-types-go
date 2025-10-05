package novakeytypes

import (	
	"github.com/google/uuid"
)

type Workspace struct {
	Id 				uuid.UUID 	`json:"id" yaml:"id"`
	Name  		string 			`json:"name" yaml:"name"`
	Project 	Project 		`json:"project,omitempty" yaml:"project,omitempty"`
	RoleCode  string 			`json:"roleCode" yaml:"roleCode"`
}

type SetWorkspaceRequest struct {	
  Id  						uuid.UUID 				`json:"id,omitempty"`
	Name  					string 						`json:"name"`
	Signer					AuthEntity 				`json:"signer"`	
}

type SetWorkspaceResponse struct {
	Id 					 		uuid.UUID 	 `json:"id,omitempty"`		
	ErrorResponse
}

type DeleteWorkspaceRequest struct {    
	Id  						uuid.UUID `json:"id"`
	Signer					AuthEntity `json:"signer"`
}

type DeleteWorkspaceResponse struct {
	Id  						uuid.UUID 	 `json:"id,omitempty"`
  ErrorResponse
}