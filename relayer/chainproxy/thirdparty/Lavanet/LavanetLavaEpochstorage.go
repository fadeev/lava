
package lavanet_thirdparty

import (
	"context"
	"encoding/json"

	// add protobuf here as pb_pkg
	"github.com/lavanet/lava/utils"
	"google.golang.org/protobuf/proto"
)

type implementedLavanetLavaEpochstorage struct {
	pb_pkg.UnimplementedQueryServer
	cb func(ctx context.Context, method string, reqBody []byte) ([]byte, error)
}

// this line is used by grpc_scaffolder #implementedLavanetLavaEpochstorage



func (is *implementedLavanetLavaEpochstorage) EpochDetails(ctx context.Context, req *pb_pkg.QueryGetEpochDetailsRequest) (*pb_pkg.QueryGetEpochDetailsResponse, error) {
	reqMarshaled, err := json.Marshal(req)
	if err != nil {
		return nil, utils.LavaFormatError("Failed to proto.Marshal(req)", err, nil)
	}
	res, err := is.cb(ctx, "lavanet.lava.epochstorage.Query.EpochDetails", reqMarshaled)
	if err != nil {
		return nil, utils.LavaFormatError("Failed to SendRelay cb", err, nil)
	}
	result := &pb_pkg.QueryGetEpochDetailsResponse{}
	err = proto.Unmarshal(res, result)
	if err != nil {
		return nil, utils.LavaFormatError("Failed to proto.Unmarshal", err, nil)
	}
	return result, nil
}
// this line is used by grpc_scaffolder #Method



func (is *implementedLavanetLavaEpochstorage) FixatedParams(ctx context.Context, req *pb_pkg.QueryGetFixatedParamsRequest) (*pb_pkg.QueryGetFixatedParamsResponse, error) {
	reqMarshaled, err := json.Marshal(req)
	if err != nil {
		return nil, utils.LavaFormatError("Failed to proto.Marshal(req)", err, nil)
	}
	res, err := is.cb(ctx, "lavanet.lava.epochstorage.Query.FixatedParams", reqMarshaled)
	if err != nil {
		return nil, utils.LavaFormatError("Failed to SendRelay cb", err, nil)
	}
	result := &pb_pkg.QueryGetFixatedParamsResponse{}
	err = proto.Unmarshal(res, result)
	if err != nil {
		return nil, utils.LavaFormatError("Failed to proto.Unmarshal", err, nil)
	}
	return result, nil
}
// this line is used by grpc_scaffolder #Method



func (is *implementedLavanetLavaEpochstorage) FixatedParamsAll(ctx context.Context, req *pb_pkg.QueryAllFixatedParamsRequest) (*pb_pkg.QueryAllFixatedParamsResponse, error) {
	reqMarshaled, err := json.Marshal(req)
	if err != nil {
		return nil, utils.LavaFormatError("Failed to proto.Marshal(req)", err, nil)
	}
	res, err := is.cb(ctx, "lavanet.lava.epochstorage.Query.FixatedParamsAll", reqMarshaled)
	if err != nil {
		return nil, utils.LavaFormatError("Failed to SendRelay cb", err, nil)
	}
	result := &pb_pkg.QueryAllFixatedParamsResponse{}
	err = proto.Unmarshal(res, result)
	if err != nil {
		return nil, utils.LavaFormatError("Failed to proto.Unmarshal", err, nil)
	}
	return result, nil
}
// this line is used by grpc_scaffolder #Method



func (is *implementedLavanetLavaEpochstorage) StakeStorage(ctx context.Context, req *pb_pkg.QueryGetStakeStorageRequest) (*pb_pkg.QueryGetStakeStorageResponse, error) {
	reqMarshaled, err := json.Marshal(req)
	if err != nil {
		return nil, utils.LavaFormatError("Failed to proto.Marshal(req)", err, nil)
	}
	res, err := is.cb(ctx, "lavanet.lava.epochstorage.Query.StakeStorage", reqMarshaled)
	if err != nil {
		return nil, utils.LavaFormatError("Failed to SendRelay cb", err, nil)
	}
	result := &pb_pkg.QueryGetStakeStorageResponse{}
	err = proto.Unmarshal(res, result)
	if err != nil {
		return nil, utils.LavaFormatError("Failed to proto.Unmarshal", err, nil)
	}
	return result, nil
}
// this line is used by grpc_scaffolder #Method



func (is *implementedLavanetLavaEpochstorage) StakeStorageAll(ctx context.Context, req *pb_pkg.QueryAllStakeStorageRequest) (*pb_pkg.QueryAllStakeStorageResponse, error) {
	reqMarshaled, err := json.Marshal(req)
	if err != nil {
		return nil, utils.LavaFormatError("Failed to proto.Marshal(req)", err, nil)
	}
	res, err := is.cb(ctx, "lavanet.lava.epochstorage.Query.StakeStorageAll", reqMarshaled)
	if err != nil {
		return nil, utils.LavaFormatError("Failed to SendRelay cb", err, nil)
	}
	result := &pb_pkg.QueryAllStakeStorageResponse{}
	err = proto.Unmarshal(res, result)
	if err != nil {
		return nil, utils.LavaFormatError("Failed to proto.Unmarshal", err, nil)
	}
	return result, nil
}
// this line is used by grpc_scaffolder #Method

// this line is used by grpc_scaffolder #Methods