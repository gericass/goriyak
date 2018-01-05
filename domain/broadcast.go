package domain

import (
	"time"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/gericass/goriyak/model/public"
	pb "github.com/gericass/goriyak/proto"
)

// MulticastTransactionAdmin : multicast transaction from admin to other admin
func MulticastTransactionAdmin(r *pb.Transaction, currentTime time.Time) error {
	admins, err := getAdminIPs()
	if err != nil {
		return err
	}
	timeProto, err := ptypes.TimestampProto(currentTime)
	if err != nil {
		return err
	}
	tr := &pb.Transaction{
		Name:          r.Name,
		SendNodeId:    r.SendNodeId,
		ReceiveNodeId: r.ReceiveNodeId,
		Amount:        r.Amount,
		Status:        "approved",
		CreatedAt:     timeProto,
	}
	for _, v := range admins {
		err := broadcastTransaction(v.IP, tr)
		if err != nil {
			return err
		}
	}
	return nil
}

// MulticastTransactionClient : multicast transaction from client to other admin
func MulticastTransactionClient(r *pb.TransactionRequest, currentTime time.Time) error {
	admins, err := getAdminIPs()
	if err != nil {
		return err
	}
	timeProto, err := ptypes.TimestampProto(currentTime)
	if err != nil {
		return err
	}
	tr := &pb.Transaction{
		Name:          r.Name,
		SendNodeId:    r.SendNodeId,
		ReceiveNodeId: r.ReceiveNodeId,
		Amount:        r.Amount,
		Status:        "approved",
		CreatedAt:     timeProto,
	}
	for _, v := range admins {
		err := broadcastTransaction(v.IP, tr)
		if err != nil {
			return err
		}
	}
	return nil
}

func getAdminIPs() ([]*public.Admin, error) {
	keys, err := public.GetAdminKey()
	if err != nil {
		return []*public.Admin{}, nil
	}
	admins := make([]*public.Admin, 0)
	for _, v := range keys.Keys {
		admin, err := public.GetAdmin(v)
		if err != nil {
			return []*public.Admin{}, err
		}
		admins = append(admins, admin)
	}
	return admins, nil
}

func broadcastTransaction(ip string, tr *pb.Transaction) error {
	conn, err := grpc.Dial(ip+":50051", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := pb.NewAdminClient(conn)

	if _, err := c.PostTransactionFromServer(context.Background(), tr); err != nil {
		return err
	}
	return nil
}
