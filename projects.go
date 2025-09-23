package novakeytypes

import "github.com/google/uuid"

type Project struct {
	Id 				uuid.UUID `json:"id"`
	RoleCode  string 		`json:"roleCode"`
}

type Key struct {
	Key				string `json:"key"`
	Value 		string `json:"value"`
}

type SetProjectRequest struct {	
  Id  						uuid.UUID 				`json:"id,omitempty"`
	Name  					string 						`json:"name,omitempty"`
	WorkspaceId  		uuid.UUID					`json:"workspaceId,omitempty"`
	Description  		string 						`json:"description,omitempty"`
	Keys   			  	[]Key 						`json:"keys,omitempty"`
	Signer					AuthEntity 				`json:"signer"`	
}

type SetProjectResponse struct {
	Id 					 		uuid.UUID 	 `json:"id,omitempty"`		
	ErrorResponse
}

type DeleteProjectRequest struct {    
	Id  						uuid.UUID `json:"id"`
	Signer					AuthEntity `json:"signer"`
}

type DeleteProjectResponse struct {
	Id  						uuid.UUID 	 `json:"id,omitempty"`
  ErrorResponse
}