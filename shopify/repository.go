package shopify

import (
	"context"
	"encoding/json"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/Shridhar2104/logilo/shopify/pb"
)

// Order defines the structure of an order in the system.
type Order struct {
	ID        string    `json:"id"`
	AccountID string    `json:"account_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Phase     string    `json:"phase"`
}

// Repository defines the methods to interact with the order storage.
type Repository interface {
	Close()
	CreateOrder(ctx context.Context, order *Order) error
	GetOrderByID(ctx context.Context, orderID string) (*Order, error)
	GetOrdersByIDs(ctx context.Context, orderIDs []string) ([]*Order, error)
	UpdateOrder(ctx context.Context, order *Order) (*Order, error)
	DeleteOrder(ctx context.Context, orderID string) error
	FetchOrders(ctx context.Context, request *pb.FetchOrdersRequest) (*pb.FetchOrdersResponse, error)
	ListOrders(ctx context.Context, skip, take int) ([]*Order, error)
	SearchOrders(ctx context.Context, query string, skip, take int) ([]*Order, error)
}

// elasticSearchRepository is an implementation of Repository using Elasticsearch.
type elasticSearchRepository struct {
	client *elastic.Client
}

// NewElasticSearchRepository initializes a new Elasticsearch repository.
func NewElasticSearchRepository(url string) (*elasticSearchRepository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	return &elasticSearchRepository{client: client}, nil
}

// CreateOrder adds a new order to Elasticsearch.
func (r *elasticSearchRepository) CreateOrder(ctx context.Context, order *Order) error {
	_, err := r.client.Index().
		Index("orders").
		Id(order.ID).
		BodyJson(order).
		Do(ctx)
	return err
}

// GetOrderByID retrieves an order by its ID.
func (r *elasticSearchRepository) GetOrderByID(ctx context.Context, orderID string) (*Order, error) {
	res, err := r.client.Get().
		Index("orders").
		Id(orderID).
		Do(ctx)
	if elastic.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var order Order
	err = json.Unmarshal(res.Source, &order)
	return &order, err
}

// GetOrdersByIDs retrieves multiple orders by their IDs
func (r *elasticSearchRepository) GetOrdersByIDs(ctx context.Context, orderIDs []string) ([]*Order, error) {
	// If no IDs are provided, return an empty slice
	if len(orderIDs) == 0 {
		return []*Order{}, nil
	}

	// Create a multi-get request
	multiGet := r.client.MultiGet()

	// Add each order ID to the multi-get request
	for _, orderID := range orderIDs {
		multiGet = multiGet.Add(elastic.NewMultiGetItem().
			Index("orders").
			Id(orderID))
	}

	// Execute the multi-get request
	res, err := multiGet.Do(ctx)
	if err != nil {
		return nil, err
	}

	// Slice to store retrieved orders
	orders := make([]*Order, 0, len(orderIDs))

	// Process each retrieved item
	for _, item := range res.Docs {
		// Skip items that weren't found or have errors
		if !item.Found || item.Source == nil {
			continue
		}

		// Unmarshal the individual order
		var order Order
		if err := json.Unmarshal(item.Source, &order); err != nil {
			// Log the error or handle it as needed
			// For this implementation, we'll skip the order if unmarshaling fails
			continue
		}

		orders = append(orders, &order)
	}

	return orders, nil
}

// UpdateOrder updates an existing order.
func (r *elasticSearchRepository) UpdateOrder(ctx context.Context, order *Order) (*Order, error) {
	_, err := r.client.Update().
		Index("orders").
		Id(order.ID).
		Doc(order).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// DeleteOrder removes an order by its ID.
func (r *elasticSearchRepository) DeleteOrder(ctx context.Context, orderID string) error {
	_, err := r.client.Delete().
		Index("orders").
		Id(orderID).
		Do(ctx)
	return err
}

// ListOrders retrieves a paginated list of orders.
func (r *elasticSearchRepository) ListOrders(ctx context.Context, skip, take int) ([]*Order, error) {
	res, err := r.client.Search().
		Index("orders").
		Query(elastic.NewMatchAllQuery()).
		From(skip).
		Size(take).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	orders := make([]*Order, len(res.Hits.Hits))
	for i, hit := range res.Hits.Hits {
		var order Order
		err := json.Unmarshal(hit.Source, &order)
		if err != nil {
			return nil, err
		}
		orders[i] = &order
	}
	return orders, nil
}

// SearchOrders searches for orders by a query string.
func (r *elasticSearchRepository) SearchOrders(ctx context.Context, query string, skip, take int) ([]*Order, error) {
	multiMatchQuery := elastic.NewMultiMatchQuery(query, "id", "account_id").
		Type("best_fields")

	res, err := r.client.Search().
		Index("orders").
		Query(multiMatchQuery).
		From(skip).
		Size(take).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	orders := make([]*Order, len(res.Hits.Hits))
	for i, hit := range res.Hits.Hits {
		var order Order
		err := json.Unmarshal(hit.Source, &order)
		if err != nil {
			return nil, err
		}
		orders[i] = &order
	}
	return orders, nil
}

// Close gracefully closes the Elasticsearch client.
func (r *elasticSearchRepository) Close() {
	r.client.Stop()
}
