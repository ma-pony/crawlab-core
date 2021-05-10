package server

//import (
//	"context"
//	"encoding/json"
//	"github.com/crawlab-team/crawlab-core/constants"
//	"github.com/crawlab-team/crawlab-core/entity"
//	"github.com/crawlab-team/crawlab-core/errors"
//	"github.com/crawlab-team/crawlab-core/grpc"
//	models2 "github.com/crawlab-team/crawlab-core/models/models"
//	"github.com/crawlab-team/crawlab-core/models/service"
//	"github.com/crawlab-team/crawlab-core/utils"
//	grpc2 "github.com/crawlab-team/crawlab-grpc"
//	"github.com/stretchr/testify/require"
//	"testing"
//	"time"
//)
//
//func TestNodeServer_Register(t *testing.T) {
//	client := grpc.TestServiceWorker.GetClient()
//	nodeKey := "test-node-key"
//	res, err := client.GetNodeClient().Register(context.Background(), &grpc2.Request{
//		NodeKey: nodeKey,
//	})
//	require.Nil(t, err)
//	require.Equal(t, grpc2.ResponseCode_OK, res.Code)
//
//	var node *models2.Node
//	err = json.Unmarshal(res.Data, &node)
//	require.Nil(t, err)
//	require.Equal(t, nodeKey, node.Key)
//	require.Equal(t, constants.NodeStatusRegistered, node.Status)
//	require.False(t, node.Id.IsZero())
//
//	node, err = modelSvc.GetNodeByKey(nodeKey, nil)
//	require.Nil(t, err)
//}
//
//func TestNodeServer_SendHeartbeat(t *testing.T) {
//	grpc.setupTest(t)
//
//	workerClient := grpc.TestServiceWorker.GetClient()
//	workerNodeKey := "worker-node-key"
//	res, err := workerClient.GetNodeClient().Register(context.Background(), &grpc2.Request{
//		NodeKey: workerNodeKey,
//	})
//	require.Nil(t, err)
//	require.Equal(t, grpc2.ResponseCode_OK, res.Code)
//
//	tic := time.Now()
//	res, err = workerClient.GetNodeClient().SendHeartbeat(context.Background(), &grpc2.Request{
//		NodeKey: workerNodeKey,
//	})
//	require.Nil(t, err)
//	var node models2.Node
//	err = json.Unmarshal(res.Data, &node)
//	require.Nil(t, err)
//	require.Equal(t, constants.NodeStatusOnline, node.Status)
//	require.Equal(t, workerNodeKey, node.Key)
//	require.False(t, node.Id.IsZero())
//	toc := node.ActiveTs
//	require.LessOrEqual(t, tic.Unix(), toc.Unix())
//	require.True(t, toc.Sub(tic) < 1*time.Second)
//
//	masterNodeKey := "master-node-key"
//	masterClient := grpc.TestServiceMaster.GetClient()
//	res, err = masterClient.GetNodeClient().SendHeartbeat(context.Background(), &grpc2.Request{
//		NodeKey: masterNodeKey,
//	})
//	require.NotNil(t, err)
//	require.Contains(t, err.Error(), errors.ErrorGrpcNotAllowed.Error())
//}
//
//func TestNodeServer_Ping(t *testing.T) {
//	grpc.setupTest(t)
//
//	masterClient := grpc.TestServiceMaster.GetClient()
//
//	res, err := masterClient.GetNodeClient().Ping(context.Background(), grpc.EmptyRequest)
//	require.Nil(t, err)
//	var nodeInfo entity.NodeInfo
//	err = json.Unmarshal(res.Data, &nodeInfo)
//	require.Nil(t, err)
//	require.NotEmpty(t, nodeInfo.Key)
//
//	workerClient := grpc.TestServiceWorker.GetClient()
//
//	res, err = workerClient.GetNodeClient().Ping(context.Background(), grpc.EmptyRequest)
//	require.NotNil(t, err)
//	require.Contains(t, err.Error(), errors.ErrorGrpcNotAllowed.Error())
//}