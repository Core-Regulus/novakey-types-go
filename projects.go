package novakeytypes

import "github.com/google/uuid"


type SetProjectRequest struct {	
  Id  						uuid.UUID 				`json:"id,omitempty"`
	Name  					string 						`json:"name,omitempty"`
	WorkspaceId  		uuid.UUID					`json:"workspaceId,omitempty"`
	Description  		string 						`json:"description,omitempty"`
	User						AuthEntity 				`json:"user"`	
}

type SetProjectResponse struct {
	Id 					 		uuid.UUID 	 `json:"id,omitempty"`		
	ErrorResponse
}

type DeleteProjectRequest struct {    
	Id  						uuid.UUID `json:"id"`
	User						AuthEntity `json:"user"`
}

type DeleteProjectResponse struct {
	Id  						uuid.UUID 	 `json:"id,omitempty"`
  ErrorResponse
}