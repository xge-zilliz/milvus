#include <dog_segment/SegmentNaive.h>
#include <random>
#include <algorithm>
#include <numeric>
#include <thread>
#include <queue>

#include <knowhere/index/vector_index/adapter/VectorAdapter.h>
#include <knowhere/index/vector_index/VecIndexFactory.h>


namespace milvus::dog_segment {
int
TestABI() {
    return 42;
}

std::unique_ptr<SegmentBase>
CreateSegment(SchemaPtr schema, IndexMetaPtr remote_index_meta) {

    if (remote_index_meta == nullptr) {
        auto index_meta = std::make_shared<IndexMeta>(schema);
        auto dim = schema->operator[]("fakevec").get_dim();
        // TODO: this is merge of query conf and insert conf
        // TODO: should be splitted into multiple configs
        auto conf = milvus::knowhere::Config{
                {milvus::knowhere::meta::DIM,           dim},
                {milvus::knowhere::IndexParams::nlist,  100},
                {milvus::knowhere::IndexParams::nprobe, 4},
                {milvus::knowhere::IndexParams::m,      4},
                {milvus::knowhere::IndexParams::nbits,  8},
                {milvus::knowhere::Metric::TYPE,        milvus::knowhere::Metric::L2},
                {milvus::knowhere::meta::DEVICEID,      0},
        };
        index_meta->AddEntry("fakeindex", "fakevec", knowhere::IndexEnum::INDEX_FAISS_IVFPQ,
                             knowhere::IndexMode::MODE_CPU, conf);
        remote_index_meta = index_meta;
    }
    auto segment = std::make_unique<SegmentNaive>(schema, remote_index_meta);
    return segment;
}

SegmentNaive::Record::Record(const Schema &schema) : uids_(1), timestamps_(1) {
    for (auto &field : schema) {
        if (field.is_vector()) {
            assert(field.get_data_type() == DataType::VECTOR_FLOAT);
            entity_vec_.emplace_back(std::make_shared<ConcurrentVector<float>>(field.get_dim()));
        } else {
            assert(field.get_data_type() == DataType::INT32);
            entity_vec_.emplace_back(std::make_shared<ConcurrentVector<int32_t, true>>());
        }
    }
}

int64_t
SegmentNaive::PreInsert(int64_t size) {
    auto reserved_begin = record_.reserved.fetch_add(size);
    return reserved_begin;
}

int64_t
SegmentNaive::PreDelete(int64_t size) {
    auto reserved_begin = deleted_record_.reserved.fetch_add(size);
    return reserved_begin;
}

auto SegmentNaive::get_deleted_bitmap(int64_t del_barrier, Timestamp query_timestamp,
                                      int64_t insert_barrier) -> std::shared_ptr<DeletedRecord::TmpBitmap> {
    auto old = deleted_record_.get_lru_entry();
    if (old->del_barrier == del_barrier) {
        return old;
    }
    auto current = std::make_shared<DeletedRecord::TmpBitmap>(*old);
    auto &vec = current->bitmap;

    if (del_barrier < old->del_barrier) {
        for (auto del_index = del_barrier; del_index < old->del_barrier; ++del_index) {
            // get uid in delete logs
            auto uid = deleted_record_.uids_[del_index];
            // map uid to corrensponding offsets, select the max one, which should be the target
            // the max one should be closest to query_timestamp, so the delete log should refer to it
            int64_t the_offset = -1;
            auto[iter_b, iter_e] = uid2offset_.equal_range(uid);
            for (auto iter = iter_b; iter != iter_e; ++iter) {
                auto offset = iter->second;
                if (record_.timestamps_[offset] < query_timestamp) {
                    assert(offset < vec.size());
                    the_offset = std::max(the_offset, offset);
                }
            }
            // if not found, skip
            if (the_offset == -1) {
                continue;
            }
            // otherwise, clear the flag
            vec[the_offset] = false;
        }
        return current;
    } else {
        vec.resize(insert_barrier);
        for (auto del_index = old->del_barrier; del_index < del_barrier; ++del_index) {
            // get uid in delete logs
            auto uid = deleted_record_.uids_[del_index];
            // map uid to corrensponding offsets, select the max one, which should be the target
            // the max one should be closest to query_timestamp, so the delete log should refer to it
            int64_t the_offset = -1;
            auto[iter_b, iter_e] = uid2offset_.equal_range(uid);
            for (auto iter = iter_b; iter != iter_e; ++iter) {
                auto offset = iter->second;
                if (offset >= insert_barrier) {
                    continue;
                }
                if (offset >= vec.size()) {
                    continue;
                }
                if (record_.timestamps_[offset] < query_timestamp) {
                    assert(offset < vec.size());
                    the_offset = std::max(the_offset, offset);
                }
            }

            // if not found, skip
            if (the_offset == -1) {
                continue;
            }

            // otherwise, set the flag
            vec[the_offset] = true;
        }
        this->deleted_record_.insert_lru_entry(current);
    }
    return current;
}

Status
SegmentNaive::Insert(int64_t reserved_begin, int64_t size, const int64_t *uids_raw, const Timestamp *timestamps_raw,
                     const DogDataChunk &entities_raw) {
    assert(entities_raw.count == size);
    assert(entities_raw.sizeof_per_row == schema_->get_total_sizeof());
    auto raw_data = reinterpret_cast<const char *>(entities_raw.raw_data);
    //    std::vector<char> entities(raw_data, raw_data + size * len_per_row);

    auto len_per_row = entities_raw.sizeof_per_row;
    std::vector<std::tuple<Timestamp, idx_t, int64_t>> ordering;
    ordering.resize(size);
    // #pragma omp parallel for
    for (int i = 0; i < size; ++i) {
        ordering[i] = std::make_tuple(timestamps_raw[i], uids_raw[i], i);
    }
    std::sort(ordering.begin(), ordering.end());
    auto sizeof_infos = schema_->get_sizeof_infos();
    std::vector<int> offset_infos(schema_->size() + 1, 0);
    std::partial_sum(sizeof_infos.begin(), sizeof_infos.end(), offset_infos.begin() + 1);
    std::vector<std::vector<char>> entities(schema_->size());

    for (int fid = 0; fid < schema_->size(); ++fid) {
        auto len = sizeof_infos[fid];
        entities[fid].resize(len * size);
    }

    std::vector<idx_t> uids(size);
    std::vector<Timestamp> timestamps(size);
    // #pragma omp parallel for
    for (int index = 0; index < size; ++index) {
        auto[t, uid, order_index] = ordering[index];
        timestamps[index] = t;
        uids[index] = uid;
        for (int fid = 0; fid < schema_->size(); ++fid) {
            auto len = sizeof_infos[fid];
            auto offset = offset_infos[fid];
            auto src = raw_data + offset + order_index * len_per_row;
            auto dst = entities[fid].data() + index * len;
            memcpy(dst, src, len);
        }
    }

    record_.timestamps_.set_data(reserved_begin, timestamps.data(), size);
    record_.uids_.set_data(reserved_begin, uids.data(), size);
    for (int fid = 0; fid < schema_->size(); ++fid) {
        record_.entity_vec_[fid]->set_data_raw(reserved_begin, entities[fid].data(), size);
    }

    for (int i = 0; i < uids.size(); ++i) {
        auto uid = uids[i];
        // NOTE: this must be the last step, cannot be put above
        uid2offset_.insert(std::make_pair(uid, reserved_begin + i));
    }

    record_.ack_responder_.AddSegment(reserved_begin, reserved_begin + size);
    return Status::OK();

    //    std::thread go(executor, std::move(uids), std::move(timestamps), std::move(entities));
    //    go.detach();
    //    const auto& schema = *schema_;
    //    auto record_ptr = GetMutableRecord();
    //    assert(record_ptr);
    //    auto& record = *record_ptr;
    //    auto data_chunk = ColumnBasedDataChunk::from(row_values, schema);
    //
    //    // TODO: use shared_lock for better concurrency
    //    std::lock_guard lck(mutex_);
    //    assert(state_ == SegmentState::Open);
    //    auto ack_id = ack_count_.load();
    //    record.uids_.grow_by(primary_keys, primary_keys + size);
    //    for (int64_t i = 0; i < size; ++i) {
    //        auto key = primary_keys[i];
    //        auto internal_index = i + ack_id;
    //        internal_indexes_[key] = internal_index;
    //    }
    //    record.timestamps_.grow_by(timestamps, timestamps + size);
    //    for (int fid = 0; fid < schema.size(); ++fid) {
    //        auto field = schema[fid];
    //        auto total_len = field.get_sizeof() * size / sizeof(float);
    //        auto source_vec = data_chunk.entity_vecs[fid];
    //        record.entity_vecs_[fid].grow_by(source_vec.data(), source_vec.data() + total_len);
    //    }
    //
    //    // finish insert
    //    ack_count_ += size;
    //    return Status::OK();
}

Status
SegmentNaive::Delete(int64_t reserved_begin, int64_t size, const int64_t *uids_raw,
                     const Timestamp *timestamps_raw) {
    std::vector<std::tuple<Timestamp, idx_t>> ordering;
    ordering.resize(size);
    // #pragma omp parallel for
    for (int i = 0; i < size; ++i) {
        ordering[i] = std::make_tuple(timestamps_raw[i], uids_raw[i]);
    }
    std::sort(ordering.begin(), ordering.end());
    std::vector<idx_t> uids(size);
    std::vector<Timestamp> timestamps(size);
    // #pragma omp parallel for
    for (int index = 0; index < size; ++index) {
        auto[t, uid] = ordering[index];
        timestamps[index] = t;
        uids[index] = uid;
    }
    deleted_record_.timestamps_.set_data(reserved_begin, timestamps.data(), size);
    deleted_record_.uids_.set_data(reserved_begin, uids.data(), size);
    deleted_record_.ack_responder_.AddSegment(reserved_begin, reserved_begin + size);
    return Status::OK();
    //    for (int i = 0; i < size; ++i) {
    //        auto key = primary_keys[i];
    //        auto time = timestamps[i];
    //        delete_logs_.insert(std::make_pair(key, time));
    //    }
    //    return Status::OK();
}

// TODO: remove mock

Status
SegmentNaive::QueryImpl(const query::QueryPtr &query, Timestamp timestamp, QueryResult &result) {
//    assert(query);

    throw std::runtime_error("unimplemnted");
}

template<typename RecordType>
int64_t get_barrier(const RecordType &record, Timestamp timestamp) {
    auto &vec = record.timestamps_;
    int64_t beg = 0;
    int64_t end = record.ack_responder_.GetAck();
    while (beg < end) {
        auto mid = (beg + end) / 2;
        if (vec[mid] < timestamp) {
            beg = mid + 1;
        } else {
            end = mid;
        }
    }
    return beg;
}

Status
SegmentNaive::Query(query::QueryPtr query_info, Timestamp timestamp, QueryResult &result) {
    // TODO: enable delete
    // TODO: enable index

    if (query_info == nullptr) {
        query_info = std::make_shared<query::Query>();
        query_info->field_name = "fakevec";
        query_info->topK = 10;
        query_info->num_queries = 1;

        auto dim = schema_->operator[]("fakevec").get_dim();
        std::default_random_engine e(42);
        std::uniform_real_distribution<> dis(0.0, 1.0);
        query_info->query_raw_data.resize(query_info->num_queries * dim);
        for (auto &x: query_info->query_raw_data) {
            x = dis(e);
        }
    }

    auto &field = schema_->operator[](query_info->field_name);
    assert(field.get_data_type() == DataType::VECTOR_FLOAT);

    auto dim = field.get_dim();
    auto topK = query_info->topK;
    auto num_queries = query_info->num_queries;

    auto barrier = get_barrier(record_, timestamp);
    auto del_barrier = get_barrier(deleted_record_, timestamp);
    auto bitmap_holder = get_deleted_bitmap(del_barrier, timestamp, barrier);

    if (!bitmap_holder) {
        throw std::runtime_error("fuck");
    }

    auto bitmap = &bitmap_holder->bitmap;

    if (topK > barrier) {
        topK = barrier;
    }

    auto get_L2_distance = [dim](const float *a, const float *b) {
        float L2_distance = 0;
        for (auto i = 0; i < dim; ++i) {
            auto d = a[i] - b[i];
            L2_distance += d * d;
        }
        return L2_distance;
    };

    std::vector<std::priority_queue<std::pair<float, int>>> records(num_queries);
    // TODO: optimize
    auto vec_ptr = std::static_pointer_cast<ConcurrentVector<float>>(record_.entity_vec_[0]);
    for (int64_t i = 0; i < barrier; ++i) {
        if (i < bitmap->size() && bitmap->at(i)) {
            continue;
        }
        auto element = vec_ptr->get_element(i);
        for (auto query_id = 0; query_id < num_queries; ++query_id) {
            auto query_blob = query_info->query_raw_data.data() + query_id * dim;
            auto dis = get_L2_distance(query_blob, element);
            auto &record = records[query_id];
            if (record.size() < topK) {
                record.emplace(dis, i);
            } else if (record.top().first > dis) {
                record.emplace(dis, i);
                record.pop();
            }
        }
    }


    result.num_queries_ = num_queries;
    result.topK_ = topK;
    auto row_num = topK * num_queries;
    result.row_num_ = topK * num_queries;

    result.result_ids_.resize(row_num);
    result.result_distances_.resize(row_num);

    for (int q_id = 0; q_id < num_queries; ++q_id) {
        // reverse
        for (int i = 0; i < topK; ++i) {
            auto dst_id = topK - 1 - i + q_id * topK;
            auto[dis, offset] = records[q_id].top();
            records[q_id].pop();
            result.result_ids_[dst_id] = record_.uids_[offset];
            result.result_distances_[dst_id] = dis;
        }
    }

    return Status::OK();
//     find end of binary
//        throw std::runtime_error("unimplemented");
//        auto record_ptr = GetMutableRecord();
//        if (record_ptr) {
//            return QueryImpl(*record_ptr, query, timestamp, result);
//        } else {
//            assert(ready_immutable_);
//            return QueryImpl(*record_immutable_, query, timestamp, result);
//        }
}

Status
SegmentNaive::Close() {
    if (this->record_.reserved != this->record_.ack_responder_.GetAck()) {
        std::runtime_error("insert not ready");
    }
    if (this->deleted_record_.reserved != this->record_.ack_responder_.GetAck()) {
        std::runtime_error("delete not ready");
    }
    state_ = SegmentState::Closed;
    return Status::OK();
}

template<typename Type>
knowhere::IndexPtr SegmentNaive::BuildVecIndexImpl(const IndexMeta::Entry &entry) {
    auto offset_opt = schema_->get_offset(entry.field_name);
    assert(offset_opt.has_value());
    auto offset = offset_opt.value();
    auto field = (*schema_)[offset];
    auto dim = field.get_dim();

    auto indexing = knowhere::VecIndexFactory::GetInstance().CreateVecIndex(entry.type, entry.mode);
    auto chunk_size = record_.uids_.chunk_size();

    auto &uids = record_.uids_;
    auto entities = record_.get_vec_entity<float>(offset);

    std::vector<knowhere::DatasetPtr> datasets;
    for (int chunk_id = 0; chunk_id < uids.chunk_size(); ++chunk_id) {
        auto &uids_chunk = uids.get_chunk(chunk_id);
        auto &entities_chunk = entities->get_chunk(chunk_id);
        int64_t count = chunk_id == uids.chunk_size() - 1 ? record_.reserved - chunk_id * DefaultElementPerChunk
                                                          : DefaultElementPerChunk;
        datasets.push_back(knowhere::GenDatasetWithIds(count, dim, entities_chunk.data(), uids_chunk.data()));
    }
    for (auto &ds: datasets) {
        indexing->Train(ds, entry.config);
    }
    for (auto &ds: datasets) {
        indexing->Add(ds, entry.config);
    }
    return indexing;
}

Status
SegmentNaive::BuildIndex() {
    for (auto&[index_name, entry]: index_meta_->get_entries()) {
        assert(entry.index_name == index_name);
        const auto &field = (*schema_)[entry.field_name];

        if (field.is_vector()) {
            assert(field.get_data_type() == engine::DataType::VECTOR_FLOAT);
            auto index_ptr = BuildVecIndexImpl<float>(entry);
            indexings_[index_name] = index_ptr;
        } else {
            throw std::runtime_error("unimplemented");
        }
    }
    return Status::OK();
}

}  // namespace milvus::dog_segment