// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.

package querycoord

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/proto/commonpb"
	"github.com/milvus-io/milvus/internal/proto/internalpb"
	"github.com/milvus-io/milvus/internal/proto/milvuspb"
	"github.com/milvus-io/milvus/internal/proto/querypb"
)

func (qc *QueryCoord) GetComponentStates(ctx context.Context) (*internalpb.ComponentStates, error) {
	serviceComponentInfo := &internalpb.ComponentInfo{
		NodeID:    Params.QueryCoordID,
		StateCode: qc.stateCode.Load().(internalpb.StateCode),
	}

	//subComponentInfos, err := qs.cluster.GetComponentInfos(ctx)
	//if err != nil {
	//	return &internalpb.ComponentStates{
	//		Status: &commonpb.Status{
	//			ErrorCode: commonpb.ErrorCode_UnexpectedError,
	//			Reason:    err.Error(),
	//		},
	//	}, err
	//}
	return &internalpb.ComponentStates{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
		},
		State: serviceComponentInfo,
		//SubcomponentStates: subComponentInfos,
	}, nil
}

func (qc *QueryCoord) GetTimeTickChannel(ctx context.Context) (*milvuspb.StringResponse, error) {
	return &milvuspb.StringResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
			Reason:    "",
		},
		Value: Params.TimeTickChannelName,
	}, nil
}

func (qc *QueryCoord) GetStatisticsChannel(ctx context.Context) (*milvuspb.StringResponse, error) {
	return &milvuspb.StringResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
			Reason:    "",
		},
		Value: Params.StatsChannelName,
	}, nil
}

func (qc *QueryCoord) ShowCollections(ctx context.Context, req *querypb.ShowCollectionsRequest) (*querypb.ShowCollectionsResponse, error) {
	dbID := req.DbID
	log.Debug("show collection start", zap.Int64("dbID", dbID))
	status := &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_Success,
	}
	if qc.stateCode.Load() != internalpb.StateCode_Healthy {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		err := errors.New("query coordinator is not healthy")
		status.Reason = err.Error()
		log.Debug("show collection end with query coordinator not healthy")
		return &querypb.ShowCollectionsResponse{
			Status: status,
		}, err
	}
	collectionIDs := qc.meta.showCollections()
	log.Debug("show collection end", zap.Int64s("collections", collectionIDs))
	return &querypb.ShowCollectionsResponse{
		Status:        status,
		CollectionIDs: collectionIDs,
	}, nil
}

func (qc *QueryCoord) LoadCollection(ctx context.Context, req *querypb.LoadCollectionRequest) (*commonpb.Status, error) {
	collectionID := req.CollectionID
	//schema := req.Schema
	log.Debug("LoadCollectionRequest received", zap.String("role", Params.RoleName), zap.Int64("msgID", req.Base.MsgID), zap.Int64("collectionID", collectionID),
		zap.Stringer("schema", req.Schema))
	status := &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_Success,
	}
	if qc.stateCode.Load() != internalpb.StateCode_Healthy {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		err := errors.New("query coordinator is not healthy")
		status.Reason = err.Error()
		log.Debug("load collection end with query coordinator not healthy")
		return status, err
	}

	loadCollectionTask := &LoadCollectionTask{
		BaseTask: BaseTask{
			ctx:              qc.loopCtx,
			Condition:        NewTaskCondition(ctx),
			triggerCondition: querypb.TriggerCondition_grpcRequest,
		},
		LoadCollectionRequest: req,
		rootCoord:             qc.rootCoordClient,
		dataCoord:             qc.dataCoordClient,
		cluster:               qc.cluster,
		meta:                  qc.meta,
	}
	qc.scheduler.Enqueue([]task{loadCollectionTask})

	err := loadCollectionTask.WaitToFinish()
	if err != nil {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		status.Reason = err.Error()
		return status, err
	}

	log.Debug("LoadCollectionRequest completed", zap.String("role", Params.RoleName), zap.Int64("msgID", req.Base.MsgID), zap.Int64("collectionID", collectionID))
	return status, nil
}

func (qc *QueryCoord) ReleaseCollection(ctx context.Context, req *querypb.ReleaseCollectionRequest) (*commonpb.Status, error) {
	//dbID := req.DbID
	collectionID := req.CollectionID
	log.Debug("ReleaseCollectionRequest received", zap.String("role", Params.RoleName), zap.Int64("msgID", req.Base.MsgID), zap.Int64("collectionID", collectionID))
	status := &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_Success,
	}
	if qc.stateCode.Load() != internalpb.StateCode_Healthy {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		err := errors.New("query coordinator is not healthy")
		status.Reason = err.Error()
		log.Debug("release collection end with query coordinator not healthy")
		return status, err
	}

	hasCollection := qc.meta.hasCollection(collectionID)
	if !hasCollection {
		log.Warn("release collection end, query coordinator don't have the log of", zap.String("collectionID", fmt.Sprintln(collectionID)))
		return status, nil
	}

	releaseCollectionTask := &ReleaseCollectionTask{
		BaseTask: BaseTask{
			ctx:              qc.loopCtx,
			Condition:        NewTaskCondition(qc.loopCtx),
			triggerCondition: querypb.TriggerCondition_grpcRequest,
		},
		ReleaseCollectionRequest: req,
		cluster:                  qc.cluster,
	}
	qc.scheduler.Enqueue([]task{releaseCollectionTask})

	err := releaseCollectionTask.WaitToFinish()
	if err != nil {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		status.Reason = err.Error()
		return status, err
	}

	log.Debug("ReleaseCollectionRequest completed", zap.String("role", Params.RoleName), zap.Int64("msgID", req.Base.MsgID), zap.Int64("collectionID", collectionID))
	qc.meta.printMeta()
	qc.cluster.printMeta()
	return status, nil
}

func (qc *QueryCoord) ShowPartitions(ctx context.Context, req *querypb.ShowPartitionsRequest) (*querypb.ShowPartitionsResponse, error) {
	collectionID := req.CollectionID
	log.Debug("show partitions start, ", zap.Int64("collectionID", collectionID))
	status := &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_Success,
	}
	if qc.stateCode.Load() != internalpb.StateCode_Healthy {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		err := errors.New("query coordinator is not healthy")
		status.Reason = err.Error()
		log.Debug("show partition end with query coordinator not healthy")
		return &querypb.ShowPartitionsResponse{
			Status: status,
		}, err
	}

	partitionIDs, err := qc.meta.showPartitions(collectionID)
	if err != nil {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		status.Reason = err.Error()
		return &querypb.ShowPartitionsResponse{
			Status: status,
		}, err
	}

	log.Debug("show partitions end", zap.Int64("collectionID", collectionID), zap.Int64s("partitionIDs", partitionIDs))

	return &querypb.ShowPartitionsResponse{
		Status:       status,
		PartitionIDs: partitionIDs,
	}, nil
}

func (qc *QueryCoord) LoadPartitions(ctx context.Context, req *querypb.LoadPartitionsRequest) (*commonpb.Status, error) {
	collectionID := req.CollectionID
	partitionIDs := req.PartitionIDs
	log.Debug("LoadPartitionRequest received", zap.String("role", Params.RoleName), zap.Int64("msgID", req.Base.MsgID), zap.Int64("collectionID", collectionID), zap.Int64s("partitionIDs", partitionIDs))
	status := &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_Success,
	}
	if qc.stateCode.Load() != internalpb.StateCode_Healthy {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		err := errors.New("query coordinator is not healthy")
		status.Reason = err.Error()
		log.Debug("load partition end with query coordinator not healthy")
		return status, err
	}

	if len(partitionIDs) == 0 {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		err := errors.New("partitionIDs are empty")
		status.Reason = err.Error()
		log.Debug("LoadPartitionRequest completed", zap.String("role", Params.RoleName), zap.Int64("msgID", req.Base.MsgID), zap.Int64("collectionID", req.CollectionID))
		return status, err
	}

	hasCollection := qc.meta.hasCollection(collectionID)
	if hasCollection {
		partitionIDsToLoad := make([]UniqueID, 0)
		loadCollection, _ := qc.meta.getLoadCollection(collectionID)
		if loadCollection {
			for _, partitionID := range partitionIDs {
				hasReleasePartition := qc.meta.hasReleasePartition(collectionID, partitionID)
				if hasReleasePartition {
					partitionIDsToLoad = append(partitionIDsToLoad, partitionID)
				}
			}
		} else {
			for _, partitionID := range partitionIDs {
				hasPartition := qc.meta.hasPartition(collectionID, partitionID)
				if !hasPartition {
					partitionIDsToLoad = append(partitionIDsToLoad, partitionID)
				}
			}
		}

		if len(partitionIDsToLoad) == 0 {
			log.Debug("LoadPartitionRequest completed", zap.String("role", Params.RoleName), zap.Int64("msgID", req.Base.MsgID), zap.Int64("collectionID", req.CollectionID))
			return status, nil
		}
		req.PartitionIDs = partitionIDsToLoad
	}

	loadPartitionTask := &LoadPartitionTask{
		BaseTask: BaseTask{
			ctx:              qc.loopCtx,
			Condition:        NewTaskCondition(qc.loopCtx),
			triggerCondition: querypb.TriggerCondition_grpcRequest,
		},
		LoadPartitionsRequest: req,
		dataCoord:             qc.dataCoordClient,
		cluster:               qc.cluster,
		meta:                  qc.meta,
	}
	qc.scheduler.Enqueue([]task{loadPartitionTask})

	err := loadPartitionTask.WaitToFinish()
	if err != nil {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		status.Reason = err.Error()
		log.Debug("LoadPartitionRequest completed", zap.String("role", Params.RoleName), zap.Int64("msgID", req.Base.MsgID), zap.Int64("collectionID", req.CollectionID))
		return status, err
	}

	log.Debug("LoadPartitionRequest completed", zap.String("role", Params.RoleName), zap.Int64("msgID", req.Base.MsgID), zap.Int64("collectionID", req.CollectionID))
	return status, nil
}

func (qc *QueryCoord) ReleasePartitions(ctx context.Context, req *querypb.ReleasePartitionsRequest) (*commonpb.Status, error) {
	//dbID := req.DbID
	collectionID := req.CollectionID
	partitionIDs := req.PartitionIDs
	log.Debug("ReleasePartitionRequest received", zap.String("role", Params.RoleName), zap.Int64("msgID", req.Base.MsgID), zap.Int64("collectionID", req.CollectionID), zap.Int64s("partitionIDs", partitionIDs))
	status := &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_Success,
	}
	if qc.stateCode.Load() != internalpb.StateCode_Healthy {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		err := errors.New("query coordinator is not healthy")
		status.Reason = err.Error()
		log.Debug("release partition end with query coordinator not healthy")
		return status, err
	}

	hasCollection := qc.meta.hasCollection(collectionID)
	if !hasCollection {
		log.Warn("release partitions end, query coordinator don't have the log of", zap.String("collectionID", fmt.Sprintln(collectionID)))
		return status, nil
	}

	if len(partitionIDs) == 0 {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		err := errors.New("partitionIDs are empty")
		status.Reason = err.Error()
		log.Debug("releasePartitionsRequest completed", zap.String("role", Params.RoleName), zap.Int64("msgID", req.Base.MsgID), zap.Int64("collectionID", req.CollectionID))
		return status, err
	}

	releasePartitionTask := &ReleasePartitionTask{
		BaseTask: BaseTask{
			ctx:              qc.loopCtx,
			Condition:        NewTaskCondition(qc.loopCtx),
			triggerCondition: querypb.TriggerCondition_grpcRequest,
		},
		ReleasePartitionsRequest: req,
		cluster:                  qc.cluster,
	}
	qc.scheduler.Enqueue([]task{releasePartitionTask})

	err := releasePartitionTask.WaitToFinish()
	if err != nil {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		status.Reason = err.Error()
		return status, err
	}
	log.Debug("ReleasePartitionRequest completed", zap.String("role", Params.RoleName), zap.Int64("msgID", req.Base.MsgID), zap.Int64("collectionID", collectionID), zap.Int64s("partitionIDs", partitionIDs))
	qc.meta.printMeta()
	qc.cluster.printMeta()
	return status, nil
}

func (qc *QueryCoord) CreateQueryChannel(ctx context.Context, req *querypb.CreateQueryChannelRequest) (*querypb.CreateQueryChannelResponse, error) {
	status := &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_Success,
	}
	if qc.stateCode.Load() != internalpb.StateCode_Healthy {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		err := errors.New("query coordinator is not healthy")
		status.Reason = err.Error()
		log.Debug("createQueryChannel end with query coordinator not healthy")
		return &querypb.CreateQueryChannelResponse{
			Status: status,
		}, err
	}

	collectionID := req.CollectionID
	queryChannel, queryResultChannel := qc.meta.GetQueryChannel(collectionID)

	return &querypb.CreateQueryChannelResponse{
		Status:         status,
		RequestChannel: queryChannel,
		ResultChannel:  queryResultChannel,
	}, nil
}

func (qc *QueryCoord) GetPartitionStates(ctx context.Context, req *querypb.GetPartitionStatesRequest) (*querypb.GetPartitionStatesResponse, error) {
	status := &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_Success,
	}
	if qc.stateCode.Load() != internalpb.StateCode_Healthy {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		err := errors.New("query coordinator is not healthy")
		status.Reason = err.Error()
		log.Debug("getPartitionStates end with query coordinator not healthy")
		return &querypb.GetPartitionStatesResponse{
			Status: status,
		}, err
	}

	partitionIDs := req.PartitionIDs
	partitionStates := make([]*querypb.PartitionStates, 0)
	for _, partitionID := range partitionIDs {
		state, err := qc.meta.getPartitionStateByID(partitionID)
		if err != nil {
			status.ErrorCode = commonpb.ErrorCode_UnexpectedError
			status.Reason = err.Error()
			return &querypb.GetPartitionStatesResponse{
				Status: status,
			}, err
		}
		partitionState := &querypb.PartitionStates{
			PartitionID: partitionID,
			State:       state,
		}
		partitionStates = append(partitionStates, partitionState)
	}

	return &querypb.GetPartitionStatesResponse{
		Status:                status,
		PartitionDescriptions: partitionStates,
	}, nil
}

func (qc *QueryCoord) GetSegmentInfo(ctx context.Context, req *querypb.GetSegmentInfoRequest) (*querypb.GetSegmentInfoResponse, error) {
	status := &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_Success,
	}
	if qc.stateCode.Load() != internalpb.StateCode_Healthy {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		err := errors.New("query coordinator is not healthy")
		status.Reason = err.Error()
		log.Debug("getSegmentInfo end with query coordinator not healthy")
		return &querypb.GetSegmentInfoResponse{
			Status: status,
		}, err
	}

	totalMemSize := int64(0)
	totalNumRows := int64(0)
	//TODO::get segment infos from meta
	//segmentIDs := req.SegmentIDs
	//segmentInfos, err := qs.meta.getSegmentInfos(segmentIDs)
	segmentInfos, err := qc.cluster.getSegmentInfo(ctx, req)
	if err != nil {
		status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		status.Reason = err.Error()
		return &querypb.GetSegmentInfoResponse{
			Status: status,
		}, err
	}
	for _, info := range segmentInfos {
		totalNumRows += info.NumRows
		totalMemSize += info.MemSize
	}
	log.Debug("getSegmentInfo", zap.Int64("num rows", totalNumRows), zap.Int64("memory size", totalMemSize))
	return &querypb.GetSegmentInfoResponse{
		Status: status,
		Infos:  segmentInfos,
	}, nil
}
