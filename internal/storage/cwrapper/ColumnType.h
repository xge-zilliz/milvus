#pragma once
enum ColumnType : int {
  NONE = 0,
  BOOL = 1,
  INT8 = 2,
  INT16 = 3,
  INT32 = 4,
  INT64 = 5,
  FLOAT = 10,
  DOUBLE = 11,
  STRING = 20,
  VECTOR_BINARY = 100,
  VECTOR_FLOAT = 101
};

enum ErrorCode : int {
  SUCCESS = 0,
  UNEXPECTED_ERROR = 1,
  CONNECT_FAILED = 2,
  PERMISSION_DENIED = 3,
  COLLECTION_NOT_EXISTS = 4,
  ILLEGAL_ARGUMENT = 5,
  ILLEGAL_DIMENSION = 7,
  ILLEGAL_INDEX_TYPE = 8,
  ILLEGAL_COLLECTION_NAME = 9,
  ILLEGAL_TOPK = 10,
  ILLEGAL_ROWRECORD = 11,
  ILLEGAL_VECTOR_ID = 12,
  ILLEGAL_SEARCH_RESULT = 13,
  FILE_NOT_FOUND = 14,
  META_FAILED = 15,
  CACHE_FAILED = 16,
  CANNOT_CREATE_FOLDER = 17,
  CANNOT_CREATE_FILE = 18,
  CANNOT_DELETE_FOLDER = 19,
  CANNOT_DELETE_FILE = 20,
  BUILD_INDEX_ERROR = 21,
  ILLEGAL_NLIST = 22,
  ILLEGAL_METRIC_TYPE = 23,
  OUT_OF_MEMORY = 24,
  DD_REQUEST_RACE = 1000
};