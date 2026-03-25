package grpc

import (
    "context"
    "time"

    "google.golang.org/grpc"
    
    pb "github.com/kgugunava/flash_sale_engine/shared/proto/order"
    "github.com/kgugunava/flash_sale_engine/pkg/logger"
)

// OrderClient — gRPC клиент к Order Service
type OrderClient struct {
    conn    *grpc.ClientConn
    client  pb.OrderServiceClient
    log     *logger.Logger
    timeout time.Duration
}

// NewOrderClient создаёт новый Order Service клиент
// Подключение устанавливается лениво при первом вызове
func NewOrderClient(address string, log *logger.Logger, timeout time.Duration) *OrderClient {
    conn, _ := NewClientConnection(address)
    
    return &OrderClient{
        conn:    conn,
        client:  pb.NewOrderServiceClient(conn),
        log:     log,
        timeout: timeout,
    }
}

// CreateOrder создаёт новый заказ
func (c *OrderClient) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
    // Добавляем таймаут к контексту ВЫЗОВА (не подключения!)
    ctx, cancel := context.WithTimeout(ctx, c.timeout)
    defer cancel()

    c.log.Debug("Calling OrderService.CreateOrder",
        logger.String("user_id", req.UserId),
        logger.String("item_name", req.ItemName),
        logger.Time("time", req.Time.AsTime()),
    )

    // Вызываем gRPC метод
    // Ошибка подключения вернётся здесь если сервер недоступен
    resp, err := c.client.CreateOrder(ctx, req)
    if err != nil {
        return nil, NewGRPCError(err, "CreateOrder")
    }

    c.log.Debug("OrderService.CreateOrder completed",
        logger.String("order_id", resp.Order.OrderId),
        logger.Bool("status_success", resp.Status.Success),
        logger.String("status_code", resp.Status.Code),
        logger.String("status_message", resp.Status.Message),
        logger.String("error_code", resp.Error.Code),
        logger.String("error_message", resp.Error.Message),
    )

    return resp, nil
}

func (c *OrderClient) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
    // Добавляем таймаут к контексту ВЫЗОВА (не подключения!)
    ctx, cancel := context.WithTimeout(ctx, c.timeout)
    defer cancel()

    c.log.Debug("Calling OrderService.CancelOrder",
        logger.String("order_id", req.OrderId),
    )

    // Вызываем gRPC метод
    // Ошибка подключения вернётся здесь если сервер недоступен
    resp, err := c.client.CancelOrder(ctx, req)
    if err != nil {
        return nil, NewGRPCError(err, "CancelOrder")
    }

    c.log.Debug("OrderService.CancelOrder completed",
        logger.Bool("status_success", resp.Status.Success),
        logger.String("status_code", resp.Status.Code),
        logger.String("status_message", resp.Status.Message),
        logger.String("error_code", resp.Error.Code),
        logger.String("error_message", resp.Error.Message),
    )

    return resp, nil
}

func (c *OrderClient) PayForOrder(ctx context.Context, req *pb.PayForOrderRequest) (*pb.PayForOrderResponse, error) {
    // Добавляем таймаут к контексту ВЫЗОВА (не подключения!)
    ctx, cancel := context.WithTimeout(ctx, c.timeout)
    defer cancel()

    c.log.Debug("Calling OrderService.PayForOrder",
        logger.String("order_id", req.OrderId),
    )

    // Вызываем gRPC метод
    // Ошибка подключения вернётся здесь если сервер недоступен
    resp, err := c.client.PayForOrder(ctx, req)
    if err != nil {
        return nil, NewGRPCError(err, "PayForOrder")
    }

    c.log.Debug("OrderService.PayForOrder completed",
        logger.Bool("status_success", resp.Status.Success),
        logger.String("status_code", resp.Status.Code),
        logger.String("status_message", resp.Status.Message),
        logger.String("error_code", resp.Error.Code),
        logger.String("error_message", resp.Error.Message),
    )

    return resp, nil
}

// GetUserOrdersList получает список заказов пользователя
func (c *OrderClient) GetUserOrdersList(ctx context.Context, req *pb.GetUserOrdersListRequest) (*pb.GetUserOrdersListResponse, error) {
    ctx, cancel := context.WithTimeout(ctx, c.timeout)
    defer cancel()

    c.log.Debug("Calling OrderService.GetUserOrdersList",
        logger.String("user_id", req.UserId),
    )

    resp, err := c.client.GetUserOrdersList(ctx, req)
    if err != nil {
        return nil, NewGRPCError(err, "GetUserOrdersList")
    }

    c.log.Debug("OrderService.GetUserOrdersList completed",
        logger.Bool("status_success", resp.Status.Success),
        logger.String("status_code", resp.Status.Code),
        logger.String("status_message", resp.Status.Message),
        logger.String("error_code", resp.Error.Code),
        logger.String("error_message", resp.Error.Message),
    )

    return resp, nil
}

// ListOrders получает список всех заказов (с фильтрами)
func (c *OrderClient) GetOrdersList(ctx context.Context, req *pb.GetOrdersListRequest) (*pb.GetOrdersListResponse, error) {
    ctx, cancel := context.WithTimeout(ctx, c.timeout)
    defer cancel()

    c.log.Debug("Calling OrderService.GetUserOrdersList")

    resp, err := c.client.GetOrdersList(ctx, req)
    if err != nil {
        return nil, NewGRPCError(err, "ListOrders")
    }

    c.log.Debug("OrderService.GetUserOrdersList completed",
        logger.Bool("status_success", resp.Status.Success),
        logger.String("status_code", resp.Status.Code),
        logger.String("status_message", resp.Status.Message),
        logger.String("error_code", resp.Error.Code),
        logger.String("error_message", resp.Error.Message),
    )

    return resp, nil
}

// Close закрывает gRPC подключение
func (c *OrderClient) Close() error {
    if c.conn != nil {
        return c.conn.Close()
    }
    return nil
}