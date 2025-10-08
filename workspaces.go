package novakeytypes

import (	
	"github.com/google/uuid"
)

type Workspace struct {
	Id 				uuid.UUID 	`json:"id" yaml:"id"`
	Name  		string 			`json:"name" yaml:"name"`
	Project 	Project 		`json:"project" yaml:"project"`
	RoleCodes []string 			`json:"roleCodes" yaml:"roleCodes"`
}

type SetWorkspaceRequest struct {	
  Id  						uuid.UUID 				`json:"id,omitempty"`
	Name  					string 						`json:"name"`
	RoleCodes  			[]string 					`json:"roleCodes" yaml:"roleCodes"`
	Signer					AuthEntity 				`json:"signer"`	
}


type SetWorkspaceResponse struct {
	Id 					 		uuid.UUID 	 `json:"id,omitempty"`
	RoleCodes  			[]string 			`json:"roleCodes" yaml:"roleCodes"`
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