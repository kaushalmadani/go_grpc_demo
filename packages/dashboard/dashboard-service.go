package dashboard

import (
	"Settings/models"
	"Settings/output"
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"os"
)

type DashboardService interface {
	Subscribe(ctx context.Context, dashboardId string, stream output.DashboardService_SubscribeServer) (models.Status, error)
}

type dashboardService struct {
}

func NewDashboardService() DashboardService { return &dashboardService{} }

func (w *dashboardService) Subscribe(_ context.Context, dashboardId string, stream output.DashboardService_SubscribeServer) (models.Status, error) {
	fmt.Println("Dashboard Id", dashboardId)
	fmt.Println("sending data:")
	if stream == nil {
		fmt.Println("Stream is empty")
		return models.Success, nil
	} else {
		fmt.Println("Sending data")
		err := stream.Send(&output.SubscribeResponse{
			Status: string(models.Success),
			Error:  "",
		})
		if err != nil {
			fmt.Println("Error while sending data through stream")
			return "", err
		}
	}
	return models.Success, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "dashboard", log.DefaultTimestampUTC)
}
