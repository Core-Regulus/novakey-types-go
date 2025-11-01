package novakeytypes

import (	
	"github.com/google/uuid"
)

type Workspace struct {
	Id 				uuid.UUID 	`json:"id" yaml:"id"`
	Name  		string 			`json:"name" yaml:"name"`
	Project 	Project 		`json:"project" yaml:"project"`
	RoleCode  string 			`json:"roleCode" yaml:"roleCode"`
}

type SetWorkspaceRequest struct {	
  Id  						uuid.UUID 				`json:"id,omitempty"`
	Name  					string 						`json:"name"`
	Description  		string 						`json:"description" yaml:"description"`
	RoleCode  			string 						`json:"roleCode" yaml:"roleCode"`
	Signer					AuthEntity 				`json:"signer"`	
}

type SetWorkspaceResponse struct {
	Id 					 		uuid.UUID 	 `json:"id,omitempty"`
	RoleCode  			string 			 `json:"roleCode" yaml:"roleCode"`
	Error
}

type GetWorkspaceRequest struct {	
  Id  						uuid.UUID 		`json:"id,omitempty"`	
	Signer					AuthEntity 		`json:"signer"`	
}

type GetWorkspaceResponse struct {
	Id 					 		uuid.UUID 	 `json:"id,omitempty"`
	Name  					string 			 `json:"name"`
	Description  		string 			 `json:"description" yaml:"description"`
	RoleCode  			string 			 `json:"roleCode" yaml:"roleCode"`
	Error
}


type DeleteWorkspaceRequest struct {    
	Id  						uuid.UUID `json:"id"`
	Signer					AuthEntity `json:"signer"`
}

type DeleteWorkspaceResponse struct {
	Id  						uuid.UUID 	 `json:"id,omitempty"`
  Error
}