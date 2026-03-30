package grpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/kgugunava/flash_sale_engine/orders_service/internal/service"
	"github.com/kgugunava/flash_sale_engine/pkg/logger"
	pb_order "github.com/kgugunava/flash_sale_engine/shared/proto/order"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	server  *grpc.Server
	port    int
	log     *logger.Logger
	handler *OrderHandler
}

// NewGRPCServer создает новый gRPC сервер
func NewGRPCServer(port int, log *logger.Logger, ordersService service.OrderServiceInterface) *GRPCServer {
	// Создаем gRPC сервер с опциями
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			// Можно добавить middleware для логирования и трейсинга
		),
	)

	// Создаем обработчик заказов
	handler := NewOrderHandler(log, ordersService)

	// Регистрируем сервис
	pb_order.RegisterOrderServiceServer(grpcServer, handler)

	return &GRPCServer{
		server:  grpcServer,
		port:    port,
		log:     log,
		handler: handler,
	}
}

// Start запускает gRPC сервер
func (s *GRPCServer) Start(ctx context.Context) error {
	// Создаем listener на указанном порту
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		s.log.Error("Failed to listen on port",
			logger.Any("error", err.Error()),
			logger.Int("port", s.port),
		)
		return err
	}

	s.log.Info("gRPC server starting",
		logger.Int("port", s.port),
	)

	// Канал для передачи ошибок из горутины Serve()
	errChan := make(chan error, 1)

	// Запускаем сервер в отдельной горутине
	go func() {
		if err := s.server.Serve(listener); err != nil {
			s.log.Error("gRPC server error",
				logger.Any("error", err.Error()),
			)
			errChan <- err
		}
	}()

	// Слушаем сигналы для graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	s.log.Info("gRPC server started successfully",
		logger.Int("port", s.port),
	)

	// Ждем сигнала завершения или ошибки
	select {
	case <-ctx.Done():
		s.log.Info("Context cancelled, shutting down gRPC server")
		s.Shutdown()
	case sig := <-sigChan:
		s.log.Info("Received signal, shutting down gRPC server",
			logger.Any("signal", sig.String()),
		)
		s.Shutdown()
	case err := <-errChan:
		s.log.Error("gRPC server failed",
			logger.Any("error", err.Error()),
		)
		return err
	}

	return nil
}

// Shutdown корректно завершает работу сервера
func (s *GRPCServer) Shutdown() {
	s.log.Info("Stopping gRPC server...")
	s.server.GracefulStop()
	s.log.Info("gRPC server stopped")
}

// Stop останавливает сервер немедленно (принудительно)
func (s *GRPCServer) Stop() {
	s.log.Info("Force stopping gRPC server...")
	s.server.Stop()
	s.log.Info("gRPC server force stopped")
}
