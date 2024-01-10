package movie

type InvoiceInput struct {
	CustomerID uint64      `json:"customer_id" binding:"required"`
	Subject    string      `json:"subject" binding:"required"`
	IssueDate  string      `json:"issue_date" binding:"required" time_format:"2006-01-02"`
	DueDate    string      `json:"due_date" binding:"required" time_format:"2006-01-02"`
	Items      []ItemInput `json:"items" binding:"required"`
}

type ItemInput struct {
	ItemID   uint64  `json:"item_id" binding:"required"`
	Quantity int32   `json:"quantity" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
}

type MovieInput struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Rating      float32 `json:"rating" binding:"required"`
	Image       string  `form:"image" json:"image"`
}
