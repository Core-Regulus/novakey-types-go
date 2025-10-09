package novakeytypes

import "github.com/google/uuid"

type Project struct {
	Id 							uuid.UUID 	`json:"id" yaml:"id"`
	RoleCodes  			[]string 		`json:"roleCodes" yaml:"roleCodes"`
	Name  					string 			`json:"name" yaml:"name"`
	WorkspaceId  		uuid.UUID		`json:"workspaceId,omitempty" yaml:"workspaceId"`
	Description  		string 			`json:"description,omitempty" yaml:"description"`
	Keys   			  	[]Key 			`json:"keys,omitempty" yaml:"keys"`

}

type Key struct {
	Key				string `json:"key" yaml:"name"`
	Value 		string `json:"value" yaml:"value"`
}

type SetProjectRequest struct {	
  Id  						uuid.UUID 				`json:"id,omitempty"`
	Name  					string 						`json:"name,omitempty"`
	WorkspaceId  		uuid.UUID					`json:"workspaceId,omitempty"`
	RoleCode  			string 						`json:"roleCode" yaml:"roleCode"`
	Description  		string 						`json:"description,omitempty"`
	Keys   			  	[]Key 						`json:"keys,omitempty"`
	Signer					AuthEntity 				`json:"signer"`	
}

type GetProjectRequest struct {	
  Id  						uuid.UUID 				`json:"id"`		
	Signer					AuthEntity 				`json:"signer"`	
}

type GetProjectResponse struct {
	Id  						uuid.UUID 				`json:"id"`	
	Keys   			  	[]Key 						`json:"keys"`
	Error
}

type SetProjectResponse struct {
	Id 					 		uuid.UUID 	 `json:"id,omitempty"`
	RoleCodes  			[]string 		 `json:"roleCodes" yaml:"roleCodes"`	
	Error
}

type DeleteProjectRequest struct {    
	Id  						uuid.UUID `json:"id"`
	Signer					AuthEntity `json:"signer"`
}

type DeleteProjectResponse struct {
	Id  						uuid.UUID 	 `json:"id,omitempty"`
  Error
}