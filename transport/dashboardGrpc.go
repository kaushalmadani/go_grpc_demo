package transport

import (
	"Settings/models"
	_ "Settings/models"
	"Settings/output"
	"Settings/packages/dashboard"
	_ "github.com/go-kit/kit/transport/grpc"
	"log"
	"time"
)

type DashboardServer struct {
	DashService dashboard.DashboardService
}

func (s *DashboardServer) Subscribe(req *output.SubscribeRequest,
	stream output.DashboardService_SubscribeServer) error {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second * 2)
		err := stream.Send(&output.SubscribeResponse{Status: string(models.Success), Data: int64(i)})
		if err != nil {
			log.Println("Error sending metric message ", err)
		}
	}
	return nil
}
