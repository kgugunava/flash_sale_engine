package grpc

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func NewClientConnection(address string) (*grpc.ClientConn, error) {
    return grpc.Dial(
        address,
        grpc.WithTransportCredentials(insecure.NewCredentials()),
        grpc.WithBlock(),  // блокирует до установки соединения
        grpc.WithTimeout(5*time.Second),  // таймаут 5 секунд
        grpc.WithDefaultCallOptions(
            grpc.MaxCallRecvMsgSize(1024*1024*4), 
            grpc.MaxCallSendMsgSize(1024*1024*4),
        ),
        grpc.WithKeepaliveParams(keepalive.ClientParameters{
            Time:                10 * time.Second,
            Timeout:             time.Second,
            PermitWithoutStream: true,
        }),
        grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
    )
}