package repositories

import (
    "errors"
    "context"
    "github.com/Mafilala/ketero/backend/initializers"
    "github.com/Mafilala/ketero/backend/models"
    "log"
    "strconv"
    "strings"
    "fmt"
)

func CreateOrder(ctx context.Context, order models.Order) (*models.Order, error) {
    query := `INSERT INTO orders (client_id, clothing_type_id, status_id, order_note, due_date) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`
    err := initializers.Pool.QueryRow(ctx, query,
        order.ClientID, order.ClothingTypeID, order.StatusID, order.OrderNote, order.DueDate,
    ).Scan(&order.ID, &order.CreatedAt)

    if err != nil {
        log.Println("Error inserting order:", err)

        
        return nil, err
    }

    return &order, nil
}

func FindOrderByID(ctx context.Context, id int) (*models.Order, error) {
    query := `SELECT id, client_id, clothing_type_id, status_id, order_note, created_at, due_date FROM orders WHERE id = $1`
    var order models.Order
    err := initializers.Pool.QueryRow(ctx, query, id).Scan(
        &order.ID, &order.ClientID, &order.ClothingTypeID, &order.StatusID, &order.OrderNote, &order.CreatedAt, &order.DueDate,
    )
    if err != nil {
        return nil, err
    }
    return &order, nil
}

func DeleteOrder(ctx context.Context, id int) (int, error) {
	var deletedID int
    query := `DELETE FROM orders WHERE id = $1 RETURNING id`
    err := initializers.Pool.QueryRow(ctx, query, id).Scan(&deletedID)

	if err != nil {
		return id, errors.New("unable to delete order")
	}
	return deletedID, nil

}

func FindAllOrders(ctx context.Context, limit, offset int, status string) ([]models.Order, int, error) {
    baseQuery := `SELECT id, client_id, clothing_type_id, status_id, order_note, created_at, due_date FROM orders`
    countQuery := `SELECT COUNT(*) FROM orders`
    var args []interface{}
    var countArgs []interface{}

    var conditions []string

    // Filtering
    if status != "" {
        conditions = append(conditions, "status_id = $1")
        args = append(args, status)
        countArgs = append(countArgs, status)
    } else {
        conditions = append(conditions, "status_id != $1", "status_id != $2")
        args = append(args, "3", "6")
        countArgs = append(countArgs, "3", "6")


    }

    // Apply WHERE if conditions exist
    if len(conditions) > 0 {
        where := " WHERE " + strings.Join(conditions, " AND ")
        baseQuery += where
        countQuery += where
    }

    // Add sorting, pagination
    args = append(args, limit, offset)

    baseQuery += fmt.Sprintf(" ORDER BY due_date DESC LIMIT $%d OFFSET $%d", len(args) - 1, len(args))
        
    // Fetch orders
    log.Println("args", args)
    rows, err := initializers.Pool.Query(ctx, baseQuery, args...)
    if err != nil {
        log.Println("after baseQuery:", err)
        return nil, 0, err
    }
    defer rows.Close()

    var orders []models.Order
    for rows.Next() {
        var o models.Order
        if err := rows.Scan(&o.ID, &o.ClientID, &o.ClothingTypeID, &o.StatusID, &o.OrderNote, &o.CreatedAt, &o.DueDate); err != nil {
            return nil, 0, err
        }
        orders = append(orders, o)
    }

    // Get total count
    var total int
    err = initializers.Pool.QueryRow(ctx, countQuery, countArgs...).Scan(&total) // use only filter args
    if err != nil {
        log.Println("couting:", err)
        return nil, 0, err
    }

    return orders, total, nil
}
func PatchOrder(ctx context.Context, id int, patch *models.Order) (*models.Order, error) {
    query := "UPDATE orders SET "
    args := []interface{}{}
    idx := 1

    if patch.ClientID != nil {
        query += "client_id = $" + strconv.Itoa(idx) + ", "
        args = append(args, *patch.ClientID)
        idx++
    }
    if patch.ClothingTypeID != nil {
        query += "clothing_type_id = $" + strconv.Itoa(idx) + ", "
        args = append(args, *patch.ClothingTypeID)
        idx++
    }
    if patch.StatusID != nil {
        query += "status_id = $" + strconv.Itoa(idx) + ", "
        args = append(args, *patch.StatusID)
        idx++
    }
    if patch.OrderNote != nil {
        query += "order_note = $" + strconv.Itoa(idx) + ", "
        args = append(args, *patch.OrderNote)
        idx++
    }
    if patch.DueDate != nil {
        query += "due_date = $" + strconv.Itoa(idx) + ", "
        args = append(args, *patch.DueDate)
        idx++
    }

    if len(args) == 0 {
        return nil, errors.New("no fields to update")
    }

    // Remove trailing comma and space
    query = query[:len(query)-2]

    query += " WHERE id = $" + strconv.Itoa(idx) + " RETURNING id, client_id, clothing_type_id, status_id, order_note, due_date, created_at"
    args = append(args, id)

    var updated models.Order
    err := initializers.Pool.QueryRow(ctx, query, args...).Scan(
        &updated.ID,
        &updated.ClientID,
        &updated.ClothingTypeID,
        &updated.StatusID,
        &updated.OrderNote,
        &updated.DueDate,
        &updated.CreatedAt,
    )
    if err != nil {
        log.Println("Error patching order:", err)
        return nil, err
    }

    return &updated, nil
}

