namespace go ameidance.paster.core

enum LanguageType {
    PLAIN = 0
    CPP = 1
    JAVA = 2
    PYTHON = 3
    GO = 4
    MD = 5
}

struct PostInfo {
    1: required string Content
    2: required LanguageType Language
    3: required string Nickname
    4: required bool IsDisposable
    5: optional i64 CreateTime
}

struct CommentInfo {
    1: required string Content
    2: required string Nickname
    3: optional i64 CreateTime
}

struct GetPostRequest {
    1: required i64 Id
    2: optional string Password
}

struct GetPostResponse {
    1: optional PostInfo Info
    254: required i32 StatusCode
    255: required string StatusMessage
}

struct SavePostRequest {
    1: required PostInfo Info
    2: optional string Password
}

struct SavePostResponse {
    1: optional i64 Id
    254: required i32 StatusCode
    255: required string StatusMessage
}

struct DeletePostRequest {
    1: required i64 Id
}

struct DeletePostResponse {
    254: required i32 StatusCode
    255: required string StatusMessage
}

struct GetCommentsRequest {
    1: required i64 PostId
    2: optional string Password
}

struct GetCommentsResponse {
    1: optional list<CommentInfo> Info
    254: required i32 StatusCode
    255: required string StatusMessage
}

struct SaveCommentRequest {
    1: required CommentInfo Info
    2: required i64 PostId
    3: optional string Password
}

struct SaveCommentResponse {
    254: required i32 StatusCode
    255: required string StatusMessage
}

service PasterCoreService {
    GetPostResponse GetPost(1: GetPostRequest req)
    SavePostResponse SavePost(1: SavePostRequest req)
    DeletePostResponse DeletePost(1: DeletePostRequest req)
    GetCommentsResponse GetComments(1: GetCommentsRequest req)
    SaveCommentResponse SaveComment(1: SaveCommentRequest req)
}
