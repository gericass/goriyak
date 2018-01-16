package domain

import (
	"github.com/gericass/goriyak/model/public"
	pb "github.com/gericass/goriyak/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// post MiningResult to other admin node
func broadcastMiningResult(r *pb.MiningResult) error {

	admins, err := public.GetAdminKey()
	if err != nil {
		return err
	}
	var sendTo string
	for _, vk := range admins.Keys {
		flag := true
		for _, vc := range r.Check {
			if vk == vc {
				flag = false
			}
		}
		if flag {
			sendTo = vk
		}
	}

	admin, err := public.GetAdmin(sendTo)
	conn, err := grpc.Dial(admin.IP+":50051", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := pb.NewAdminClient(conn)

	if _, err := c.PostBlock(context.Background(), r); err != nil {
		return err
	}
	return nil

	return nil
}

func checkActiveMiningResult(r *pb.MiningResult) (bool, error) {
	admins, err := public.GetAdminKey()
	if err != nil {
		return false, err
	}
	lack := (len(admins.Keys) * 2 / 3) - len(r.Sign)
	remaining := len(admins.Keys) - len(r.Check)
	if lack > remaining {
		return false, nil
	}
	return true, nil
}

// TODO implements
func generateBlockByMiningResult(r *pb.MiningResult) (*pb.Block, error) {
	return nil, nil
}

// TODO implements
func confirmHashing(b *pb.Block) bool {

	return true
}

// TODO implements
func updatedMiningResult(r *pb.MiningResult) (*pb.MiningResult, error) {
	block, err := generateBlockByMiningResult(r)
	if err != nil {
		return nil, err
	}

	hashingResult := confirmHashing(block)
	if hashingResult {
		return nil, nil
	}

	return nil, nil
}

func MiningController(miningResult *pb.MiningResult) (*pb.Status, error) {
	if res, _ := public.GetBlock(miningResult.BlockId); res != nil {
		return &pb.Status{Message: "Block already exists"}, nil
	}

	ex, err := checkActiveMiningResult(miningResult)
	if err != nil {
		return &pb.Status{Message: "Server error"}, err
	}
	if ex {

		res, err := updatedMiningResult(miningResult)
		if err != nil {
			return &pb.Status{Message: "Mining failed"}, nil
		}

		if err := broadcastMiningResult(res); err != nil {
			return &pb.Status{Message: "Server error"}, err
		}
	}

	return &pb.Status{Message: "mining result received"}, nil

}
