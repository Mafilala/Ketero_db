package models

import "time"


type Measure struct {
    ID    int    `json:"id" binding:"required"`
    Name  string `json:"name" binding:"required"`
}

type MeasureValue struct {
    MeasureID    int    `json:"measure_id" binding:"required"`
    Value int `json:"value" binding:"required"`
}

type ClothingType struct {
    ID    int    `json:"id" binding:"required"`
    Name  string `json:"name" binding:"required"`
}

type Clothing struct {
    ID    int    `json:"id" binding:"required"`
    Name  string `json:"name" binding:"required"`
}

type Status struct {
    ID    int    `json:"id" binding:"required"`
    Name  string `json:"name" binding:"required"`
}

type Client struct {
    ID    int    `json:"id" binding:"required"`
    Name  string `json:"full_name" binding:"required"`
    Phone  string `json:"phone_number" binding:"required"`
}

type ClothingTypePart struct {
    Clothing_type_id    int    `json:"clothing_type_id" binding:"required"`
    Clothing_id    int    `json:"clothing_id" binding:"required"`
}

type ClothingMeasures struct {
    Clothing_id    int    `json:"clothing_id" binding:"required"`
    Measure_id    int    `json:"measure_id" binding:"required"`
}

type Order struct {
    ID             int        `json:"id"`
    ClientID       *int       `json:"client_id,omitempty"`
    ClothingTypeID *int       `json:"clothing_type_id,omitempty"`
    StatusID       *int       `json:"status_id,omitempty"`
    OrderNote      *string    `json:"order_note,omitempty"`
    DueDate        *time.Time `json:"due_date,omitempty"`
    CreatedAt      time.Time  `json:"created_at"`
}

type OrderMeasure struct {
    OrderID    int `json:"order_id" binding:"required"`
    MeasureID  int `json:"measure_id" binding:"required"`
    Measure    *int `json:"measure" binding:"required"`
    ClothingID int `json:"clothing_id" binding:"required"`
}

type PriceDetail struct {
	OrderID int     `json:"order_id" binding:"required"`
	Price   float64 `json:"price" binding:"required"`
	Paid    float64 `json:"paid" binding:"required"`
}


type OrderDetail struct {
	OrderID int `json:"order_id" binding:"required"`
	Style   *int `json:"style" binding:"required"`
	Fabric  *int `json:"fabric" binding:"required"`
	Color   *int `json:"color" binding:"required"`
}

