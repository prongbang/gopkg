package core

type AccessToken struct {
	Token string `json:"token"`
}

type UserRequestInfo struct {
	Id             string   `json:"-"`
	HasUserRequest bool     `json:"-"`
	Roles          []string `json:"-"`
	ApiInfo        *ApiInfo `json:"-"`
}

type ApiInfo struct {
	PermissionId string `json:"-"`
}

type RequestInfo struct {
	UserRequestInfo *UserRequestInfo
}
