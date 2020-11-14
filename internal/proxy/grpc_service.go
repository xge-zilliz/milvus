package proxy

import (
	"context"
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/zilliztech/milvus-distributed/internal/msgstream"
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/schemapb"
	"github.com/zilliztech/milvus-distributed/internal/proto/servicepb"
)

func (p *Proxy) Insert(ctx context.Context, in *servicepb.RowBatch) (*servicepb.IntegerRangeResponse, error) {
	it := &InsertTask{
		BaseInsertTask: BaseInsertTask{
			BaseMsg: msgstream.BaseMsg{
				HashValues: in.HashKeys,
			},
			InsertRequest: internalpb.InsertRequest{
				MsgType:        internalpb.MsgType_kInsert,
				CollectionName: in.CollectionName,
				PartitionTag:   in.PartitionTag,
				RowData:        in.RowData,
			},
		},
		done:                  make(chan error),
		manipulationMsgStream: p.manipulationMsgStream,
	}

	it.ctx, it.cancel = context.WithCancel(ctx)
	// TODO: req_id, segment_id, channel_id, proxy_id, timestamps, row_ids

	defer it.cancel()

	fn := func() error {
		select {
		case <-ctx.Done():
			return errors.New("insert timeout")
		default:
			return p.taskSch.DdQueue.Enqueue(it)
		}
	}
	err := fn()

	if err != nil {
		return &servicepb.IntegerRangeResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			},
		}, nil
	}

	err = it.WaitToFinish()
	if err != nil {
		return &servicepb.IntegerRangeResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			},
		}, nil
	}

	return it.result, nil
}

func (p *Proxy) CreateCollection(ctx context.Context, req *schemapb.CollectionSchema) (*commonpb.Status, error) {
	cct := &CreateCollectionTask{
		CreateCollectionRequest: internalpb.CreateCollectionRequest{
			MsgType: internalpb.MsgType_kCreateCollection,
			Schema:  &commonpb.Blob{},
			// TODO: req_id, timestamp, proxy_id
		},
		masterClient: p.masterClient,
		done:         make(chan error),
	}
	schemaBytes, _ := proto.Marshal(req)
	cct.CreateCollectionRequest.Schema.Value = schemaBytes
	cct.ctx, cct.cancel = context.WithCancel(ctx)
	defer cct.cancel()

	fn := func() error {
		select {
		case <-ctx.Done():
			return errors.New("create collection timeout")
		default:
			return p.taskSch.DdQueue.Enqueue(cct)
		}
	}
	err := fn()
	if err != nil {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    err.Error(),
		}, err
	}

	err = cct.WaitToFinish()
	if err != nil {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    err.Error(),
		}, err
	}

	return cct.result, nil
}

func (p *Proxy) Search(ctx context.Context, req *servicepb.Query) (*servicepb.QueryResult, error) {
	qt := &QueryTask{
		SearchRequest: internalpb.SearchRequest{
			MsgType: internalpb.MsgType_kSearch,
			Query:   &commonpb.Blob{},
			// TODO: req_id, proxy_id, timestamp, result_channel_id
		},
		queryMsgStream: p.queryMsgStream,
		done:           make(chan error),
		resultBuf:      make(chan []*internalpb.SearchResult),
	}
	qt.ctx, qt.cancel = context.WithCancel(ctx)
	queryBytes, _ := proto.Marshal(req)
	qt.SearchRequest.Query.Value = queryBytes
	defer qt.cancel()

	fn := func() error {
		select {
		case <-ctx.Done():
			return errors.New("create collection timeout")
		default:
			return p.taskSch.DdQueue.Enqueue(qt)
		}
	}
	err := fn()
	if err != nil {
		return &servicepb.QueryResult{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			},
		}, err
	}

	err = qt.WaitToFinish()
	if err != nil {
		return &servicepb.QueryResult{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			},
		}, err
	}

	return qt.result, nil
}

func (p *Proxy) DropCollection(ctx context.Context, req *servicepb.CollectionName) (*commonpb.Status, error) {
	dct := &DropCollectionTask{
		DropCollectionRequest: internalpb.DropCollectionRequest{
			MsgType: internalpb.MsgType_kDropCollection,
			// TODO: req_id, timestamp, proxy_id
			CollectionName: req,
		},
		masterClient: p.masterClient,
		done:         make(chan error),
	}
	dct.ctx, dct.cancel = context.WithCancel(ctx)
	defer dct.cancel()

	fn := func() error {
		select {
		case <-ctx.Done():
			return errors.New("create collection timeout")
		default:
			return p.taskSch.DdQueue.Enqueue(dct)
		}
	}
	err := fn()
	if err != nil {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    err.Error(),
		}, err
	}

	err = dct.WaitToFinish()
	if err != nil {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    err.Error(),
		}, err
	}

	return dct.result, nil
}

func (p *Proxy) HasCollection(ctx context.Context, req *servicepb.CollectionName) (*servicepb.BoolResponse, error) {
	hct := &HasCollectionTask{
		HasCollectionRequest: internalpb.HasCollectionRequest{
			MsgType: internalpb.MsgType_kHasCollection,
			// TODO: req_id, timestamp, proxy_id
			CollectionName: req,
		},
		masterClient: p.masterClient,
		done:         make(chan error),
	}
	hct.ctx, hct.cancel = context.WithCancel(ctx)
	defer hct.cancel()

	fn := func() error {
		select {
		case <-ctx.Done():
			return errors.New("create collection timeout")
		default:
			return p.taskSch.DdQueue.Enqueue(hct)
		}
	}
	err := fn()
	if err != nil {
		return &servicepb.BoolResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			},
		}, err
	}

	err = hct.WaitToFinish()
	if err != nil {
		return &servicepb.BoolResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			},
		}, err
	}

	return hct.result, nil
}

func (p *Proxy) DescribeCollection(ctx context.Context, req *servicepb.CollectionName) (*servicepb.CollectionDescription, error) {
	dct := &DescribeCollectionTask{
		DescribeCollectionRequest: internalpb.DescribeCollectionRequest{
			MsgType: internalpb.MsgType_kDescribeCollection,
			// TODO: req_id, timestamp, proxy_id
			CollectionName: req,
		},
		masterClient: p.masterClient,
		done:         make(chan error),
	}
	dct.ctx, dct.cancel = context.WithCancel(ctx)
	defer dct.cancel()

	fn := func() error {
		select {
		case <-ctx.Done():
			return errors.New("create collection timeout")
		default:
			return p.taskSch.DdQueue.Enqueue(dct)
		}
	}
	err := fn()
	if err != nil {
		return &servicepb.CollectionDescription{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			},
		}, err
	}

	err = dct.WaitToFinish()
	if err != nil {
		return &servicepb.CollectionDescription{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			},
		}, err
	}

	return dct.result, nil
}

func (p *Proxy) ShowCollections(ctx context.Context, req *commonpb.Empty) (*servicepb.StringListResponse, error) {
	sct := &ShowCollectionsTask{
		ShowCollectionRequest: internalpb.ShowCollectionRequest{
			MsgType: internalpb.MsgType_kDescribeCollection,
			// TODO: req_id, timestamp, proxy_id
		},
		masterClient: p.masterClient,
		done:         make(chan error),
	}
	sct.ctx, sct.cancel = context.WithCancel(ctx)
	defer sct.cancel()

	fn := func() error {
		select {
		case <-ctx.Done():
			return errors.New("create collection timeout")
		default:
			return p.taskSch.DdQueue.Enqueue(sct)
		}
	}
	err := fn()
	if err != nil {
		return &servicepb.StringListResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			},
		}, err
	}

	err = sct.WaitToFinish()
	if err != nil {
		return &servicepb.StringListResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			},
		}, err
	}

	return sct.result, nil
}

func (p *Proxy) CreatePartition(ctx context.Context, in *servicepb.PartitionName) (*commonpb.Status, error) {
	return &commonpb.Status{
		ErrorCode: 0,
		Reason:    "",
	}, nil
}

func (p *Proxy) DropPartition(ctx context.Context, in *servicepb.PartitionName) (*commonpb.Status, error) {
	return &commonpb.Status{
		ErrorCode: 0,
		Reason:    "",
	}, nil
}

func (p *Proxy) HasPartition(ctx context.Context, in *servicepb.PartitionName) (*servicepb.BoolResponse, error) {
	return &servicepb.BoolResponse{
		Status: &commonpb.Status{
			ErrorCode: 0,
			Reason:    "",
		},
		Value: true,
	}, nil
}

func (p *Proxy) DescribePartition(ctx context.Context, in *servicepb.PartitionName) (*servicepb.PartitionDescription, error) {
	return &servicepb.PartitionDescription{
		Status: &commonpb.Status{
			ErrorCode: 0,
			Reason:    "",
		},
	}, nil
}

func (p *Proxy) ShowPartitions(ctx context.Context, req *servicepb.CollectionName) (*servicepb.StringListResponse, error) {
	return &servicepb.StringListResponse{
		Status: &commonpb.Status{
			ErrorCode: 0,
			Reason:    "",
		},
	}, nil
}