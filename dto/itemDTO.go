package dto

type CreateItemInput struct {
	Name        string `json:"Name" binding:"required,min=2"`
	Price       uint   `json:"Price" binding:"required,min=1,max=999999"`
	Description string `json:"Description"`
}

type UpdateItemInput struct {
	Name        *string `json:"Name" binding:"omitempty,min=2"`
	Price       *uint   `json:"Price" binding:"omitempty,min=1,max=999999"`
	Description *string `json:"Description"`
	SoldOut     *bool   `json:"SoldOut"`
}
