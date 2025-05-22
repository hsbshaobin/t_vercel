package model

type GeneralCodeBO struct {
	AppId string `json:"appId" binding:"required,min=6"`
	Count int    `json:"count" binding:"required,min=1,max=500"`
}

type CheckCodeBO struct {
	BrowserId  string `json:"browserId" binding:"required,min=6"`
	AppId      string `json:"appId" binding:"required,min=6"`
	InviteCode string `json:"inviteCode" binding:"required,min=6"`
}
