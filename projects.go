package novakeytypes

import "github.com/google/uuid"

type Project struct {
	Id 				uuid.UUID `json:"id"`
	RoleCode  string 		`json:"roleCode"`
}

type Key struct {
	Key				string `json:"key" yaml:"name"`
	Value 		string `json:"value" yaml:"value"`
}

type SetProjectRequest struct {	
  Id  						uuid.UUID 				`json:"id,omitempty" yaml:"id"`
	Name  					string 						`json:"name,omitempty" yaml:"name"`
	WorkspaceId  		uuid.UUID					`json:"workspaceId,omitempty" yaml:"workspaceId"`
	Description  		string 						`json:"description,omitempty" yaml:"description"`
	Keys   			  	[]Key 						`json:"keys,omitempty" yaml:"keys"`
	Signer					AuthEntity 				`json:"signer"`	
}

type GetProjectRequest struct {	
  Id  						uuid.UUID 				`json:"id"`		
	Signer					AuthEntity 				`json:"signer"`	
}

type GetProjectResponse struct {
	Id  						uuid.UUID 				`json:"id"`
	Keys   			  	[]Key 						`json:"keys"`
	ErrorResponse
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