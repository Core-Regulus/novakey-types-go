package novakeytypes

import "github.com/google/uuid"

type Project struct {
	Id 							uuid.UUID 		`json:"id" yaml:"id"`
	RoleCode  			string 				`json:"roleCode" yaml:"roleCode"`
	Name  					string 				`json:"name" yaml:"name"`
	KeyPass  				string 				`json:"keyPass" yaml:"keyPass"`
	WorkspaceId  		uuid.UUID			`json:"workspaceId,omitempty" yaml:"workspaceId"`
	Description  		string 				`json:"description,omitempty" yaml:"description"`
	Keys   			  	[]Key					`json:"keys,omitempty" yaml:"keys"`
	Users						[]UserEntry 	`json:"users,omitempty" yaml:"users"`
}

type Key struct {
	Key				string `json:"key" yaml:"key"`
	Value 		string `json:"value" yaml:"value"`
}

type UserEntry struct {
	Key				string `json:"key" yaml:"name"`
	RoleName 	string `json:"roleName" yaml:"roleName"`
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
	RoleCode  			string 						`json:"roleCode"`
	Name  					string 						`json:"name"`
	WorkspaceId  		uuid.UUID					`json:"workspaceId"`
	Description  		string 						`json:"description"`
	Keys   			  	[]Key 						`json:"keys"`
	Error
}

type SetProjectResponse struct {
	Id 					 		uuid.UUID 	`json:"id,omitempty"`
	RoleCode  			string 		 	`json:"roleCode" yaml:"roleCode"`
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

func (p *Project) GetKey(keyName string) (string, bool) {
	for _, key := range p.Keys {
		if key.Key == keyName {
			return key.Value, true
		}
	}
	return "", false
}