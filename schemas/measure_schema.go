package schemas
import "time"
import "github.com/Mafilala/ketero/backend/models"
type CreateMeasureRequest struct {
    Name  string `json:"name" binding:"required"`
}

type GetMeasureResponse struct {
    Id int `json:"id" binding;"required"`
    Name  string `json:"name" binding:"required"`
}

type UpdateClothingRequest struct {
    Name  string `json:"name" binding:"required"`
}

type UpdateStatus struct {
    Name  string `json:"name" binding:"required"`
}

type UpdateClothingTypeRequest struct {
    Name  string `json:"name" binding:"required"`
}



type CreateClothingTypeRequest struct {
    Name string `json:"name" binding:"required"` 
}

type CreateClothingRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateStatusRequest struct {
	Name string `json:"name" binding:"required"`
}



type CreateClientRequest struct {
    Name string `json:"full_name" binding:"required"`
    Phone  string `json:"phone_number" binding:"required"`
}

type CreatClothingTypePartRequest struct {
    Clothing_type_id    int    `json:"clothing_type_id" binding:"required"`
    Clothing_id    int    `json:"clothing_id" binding:"required"`
}

type CreatClothingMeasureRequest struct {
    Clothing_id    int    `json:"clothing_id" binding:"required"`
    Measure_id    int    `json:"measure_id" binding:"required"`
}

type CreateOrderRequest struct {
    ClientID       *int    `json:"client_id" binding:"required"`
    ClothingTypeID *int    `json:"clothing_type_id" binding:"required"`
    StatusID       *int    `json:"status_id"`
    OrderNote      *string `json:"order_note"`
    DueDate        *time.Time `json:"due_date" binding:"required"`
}

type CreateOrderMeasureRequest struct {
    OrderID    int `json:"order_id" binding:"required"`
    MeasureID  int `json:"measure_id" binding:"required"`
    Measure    int `json:"measure" binding:"required"`
    ClothingID int `json:"clothing_id" binding:"required"`
}

type Measure struct {
    MeasureID int `json:"measure_id" binding:"required"`
    Value int `json:"value" binding:"required"`
}

type UpdateOrderMeasureRequest struct {
    ClothingID int `json:"clothing_id" binding:"required"`
    Measures []models.MeasureValue `json: "measures" binding: "required"`
}

func (r CreateOrderMeasureRequest) ToModel() models.OrderMeasure {
	return models.OrderMeasure{
		OrderID:    r.OrderID,
		MeasureID:  r.MeasureID,
		Measure:    r.Measure,
		ClothingID: r.ClothingID,
	}
}

type CreatePriceDetailRequest struct {
	OrderID int     `json:"order_id" binding:"required"`
	Price   float64 `json:"price" binding:"required"`
	Paid    float64 `json:"paid" binding:"required"`
}

func (r CreatePriceDetailRequest) ToModel() models.PriceDetail {
	return models.PriceDetail{
		OrderID: r.OrderID,
		Price:   r.Price,
		Paid:    r.Paid,
	}
}


type UpdatePriceDetailRequest struct {
    Price *float64 `json:"price"`
    Paid  *float64 `json:"paid"`
}

func (r *UpdatePriceDetailRequest) ToModel(orderID int, existing models.PriceDetail) models.PriceDetail {
    if r.Price != nil {
        existing.Price = *r.Price
    }
    if r.Paid != nil {
        existing.Paid = *r.Paid
    }
    existing.OrderID = orderID
    return existing
}

type CreateOrderDetailRequest struct {
	OrderID int `json:"order_id" binding:"required"`
	Style   int `json:"style" binding:"required"`
	Fabric  int `json:"fabric" binding:"required"`
	Color   int `json:"color" binding:"required"`
}

type UpdateOrderDetailRequest struct {
	Style  *int `json:"style,omitempty"`
	Fabric *int `json:"fabric,omitempty"`
	Color  *int `json:"color,omitempty"`
}

type PatchOrderRequest struct {
    ClientID       *int       `json:"client_id"`
    ClothingTypeID *int       `json:"clothing_type_id"`
    StatusID       *int       `json:"status_id"`
    OrderNote      *string    `json:"order_note"`
    DueDate        *time.Time `json:"due_date"`
}
