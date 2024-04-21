package dto

type CreateItemInput struct {
	Name        string `json:"Name" binding:"required,min=2"`
	Price       uint   `json:"Price" binding:"required,min=1,max=999999"`
	Description string `json:"Description"`
}
