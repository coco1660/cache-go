namespace go cache_service

struct BaseResp {
    1: i64 code
    2: string msg
}

struct CacheItem {
    1: string key
    2: string data
    3: i64 life_span
    4: i64 created_on
    5: i64 accessed_on
    6: i64 access_count
}

struct NewCacheRequest {
    1: required string name
}

struct NewCacheResponse {
    1: bool success
    2: BaseResp base
}

struct SetRequest {
    1: required string key
    2: required string value
    3: optional i64 life_span
}

struct SetResponse {
    1: BaseResp base
}

struct GetRequest {
    1: required string cache
    2: required string key
}

struct GetResponse {
    1: optional CacheItem item
    2: BaseResp base
}

struct DeleteRequest {
    1: required string cache
    2: required string key
}

struct DeleteResponse {
    1: bool deleted
    2: BaseResp base
}

struct ExistsRequest {
    1: required string cache
    2: required string key
}

struct ExistsResponse {
    1: bool exists
    2: BaseResp base
}

service CacheService {
    NewCacheResponse New(1: NewCacheRequest req)
    GetResponse Get(1: GetRequest req)
    SetResponse Set(1: SetRequest req)
    DeleteResponse Delete(1: DeleteRequest req)
    ExistsResponse Exists(1: ExistsRequest req)
}
