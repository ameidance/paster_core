# paster_core

Paster 服务端核心模块，使用字节跳动开源的微服务 RPC 框架 [KiteX](https://github.com/cloudwego/kitex)
，以 [Apache Thrift](https://github.com/apache/thrift) 作为通信协议。

[Sequence Diagram](https://mermaid-js.github.io/mermaid-live-editor/edit/#eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG4gICAgcGFydGljaXBhbnQgZmVcbiAgICBwYXJ0aWNpcGFudCBmYWNhZGVcbiAgICBwYXJ0aWNpcGFudCBjb3JlXG4gICAgcGFydGljaXBhbnQgcmVkaXNcbiAgICBwYXJ0aWNpcGFudCBteXNxbFxuXG4gICAgZmUtPj5mZTogL1xuICAgIGFjdGl2YXRlIGZlXG4gICAgZGVhY3RpdmF0ZSBmZVxuXG4gICAgZmUtPj5mYWNhZGU6IFBPU1QgL3Bvc3Qvc2F2ZS9cbiAgICBhY3RpdmF0ZSBmZVxuICAgIGFjdGl2YXRlIGZhY2FkZVxuICAgIGZhY2FkZS0-PnJlZGlzOiBnZXRcbiAgICBhY3RpdmF0ZSByZWRpc1xuICAgIHJlZGlzLS0-PmZhY2FkZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSByZWRpc1xuICAgIGZhY2FkZS0-PmNvcmU6IFNhdmVQb3N0XG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogd3JpdGVcbiAgICBhY3RpdmF0ZSBteXNxbFxuICAgIG15c3FsLS0-PmNvcmU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgbXlzcWxcbiAgICBjb3JlLS0-PmZhY2FkZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBjb3JlXG4gICAgZmFjYWRlLT4-cmVkaXM6IHNldFxuICAgIGFjdGl2YXRlIHJlZGlzXG4gICAgcmVkaXMtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIHJlZGlzXG4gICAgZmFjYWRlLS0-PmZlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGZhY2FkZVxuICAgIGRlYWN0aXZhdGUgZmVcblxuICAgIGZlLT4-ZmU6IC88aWQ-XG4gICAgYWN0aXZhdGUgZmVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogR0VUIC9jb21tZW50L2dldC9cbiAgICBhY3RpdmF0ZSBmZVxuICAgIGFjdGl2YXRlIGZhY2FkZVxuICAgIGZhY2FkZS0-PmNvcmU6IEdldENvbW1lbnRzXG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogcmVhZFxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBhY3RpdmF0ZSByZWRpc1xuICAgIGZhY2FkZS0tPj5mZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBmYWNhZGVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogUE9TVCAvcG9zdC9nZXQvXG4gICAgYWN0aXZhdGUgZmVcbiAgICBhY3RpdmF0ZSBmYWNhZGVcbiAgICBmYWNhZGUtPj5jb3JlOiBHZXRQb3N0XG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogcmVhZFxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBmYWNhZGUtPj5jb3JlOiAoRGVsZXRlUG9zdClcbiAgICBhY3RpdmF0ZSBjb3JlXG4gICAgY29yZS0tPj5mYWNhZGU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgY29yZVxuICAgIGZhY2FkZS0tPj5mZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBmYWNhZGVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogUE9TVCAvY29tbWVudC9zYXZlL1xuICAgIGFjdGl2YXRlIGZlXG4gICAgYWN0aXZhdGUgZmFjYWRlXG4gICAgZmFjYWRlLT4-cmVkaXM6IGdldFxuICAgIGFjdGl2YXRlIHJlZGlzXG4gICAgcmVkaXMtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIHJlZGlzXG4gICAgZmFjYWRlLT4-Y29yZTogU2F2ZUNvbW1lbnRcbiAgICBhY3RpdmF0ZSBjb3JlXG4gICAgY29yZS0-Pm15c3FsOiB3cml0ZVxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBmYWNhZGUtPj5yZWRpczogc2V0XG4gICAgYWN0aXZhdGUgcmVkaXNcbiAgICByZWRpcy0tPj5mYWNhZGU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgcmVkaXNcbiAgICBmYWNhZGUtLT4-ZmU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgZmFjYWRlXG4gICAgZGVhY3RpdmF0ZSBmZVxuIiwibWVybWFpZCI6IntcbiAgXCJ0aGVtZVwiOiBcImRlZmF1bHRcIlxufSIsInVwZGF0ZUVkaXRvciI6ZmFsc2UsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjpmYWxzZX0)

![](https://mermaid.ink/svg/eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG4gICAgcGFydGljaXBhbnQgZmVcbiAgICBwYXJ0aWNpcGFudCBmYWNhZGVcbiAgICBwYXJ0aWNpcGFudCBjb3JlXG4gICAgcGFydGljaXBhbnQgcmVkaXNcbiAgICBwYXJ0aWNpcGFudCBteXNxbFxuXG4gICAgZmUtPj5mZTogL1xuICAgIGFjdGl2YXRlIGZlXG4gICAgZGVhY3RpdmF0ZSBmZVxuXG4gICAgZmUtPj5mYWNhZGU6IFBPU1QgL3Bvc3Qvc2F2ZS9cbiAgICBhY3RpdmF0ZSBmZVxuICAgIGFjdGl2YXRlIGZhY2FkZVxuICAgIGZhY2FkZS0-PnJlZGlzOiBnZXRcbiAgICBhY3RpdmF0ZSByZWRpc1xuICAgIHJlZGlzLS0-PmZhY2FkZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSByZWRpc1xuICAgIGZhY2FkZS0-PmNvcmU6IFNhdmVQb3N0XG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogd3JpdGVcbiAgICBhY3RpdmF0ZSBteXNxbFxuICAgIG15c3FsLS0-PmNvcmU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgbXlzcWxcbiAgICBjb3JlLS0-PmZhY2FkZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBjb3JlXG4gICAgZmFjYWRlLT4-cmVkaXM6IHNldFxuICAgIGFjdGl2YXRlIHJlZGlzXG4gICAgcmVkaXMtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIHJlZGlzXG4gICAgZmFjYWRlLS0-PmZlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGZhY2FkZVxuICAgIGRlYWN0aXZhdGUgZmVcblxuICAgIGZlLT4-ZmU6IC88aWQ-XG4gICAgYWN0aXZhdGUgZmVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogR0VUIC9jb21tZW50L2dldC9cbiAgICBhY3RpdmF0ZSBmZVxuICAgIGFjdGl2YXRlIGZhY2FkZVxuICAgIGZhY2FkZS0-PmNvcmU6IEdldENvbW1lbnRzXG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogcmVhZFxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBhY3RpdmF0ZSByZWRpc1xuICAgIGZhY2FkZS0tPj5mZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBmYWNhZGVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogUE9TVCAvcG9zdC9nZXQvXG4gICAgYWN0aXZhdGUgZmVcbiAgICBhY3RpdmF0ZSBmYWNhZGVcbiAgICBmYWNhZGUtPj5jb3JlOiBHZXRQb3N0XG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogcmVhZFxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBmYWNhZGUtPj5jb3JlOiAoRGVsZXRlUG9zdClcbiAgICBhY3RpdmF0ZSBjb3JlXG4gICAgY29yZS0tPj5mYWNhZGU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgY29yZVxuICAgIGZhY2FkZS0tPj5mZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBmYWNhZGVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogUE9TVCAvY29tbWVudC9zYXZlL1xuICAgIGFjdGl2YXRlIGZlXG4gICAgYWN0aXZhdGUgZmFjYWRlXG4gICAgZmFjYWRlLT4-cmVkaXM6IGdldFxuICAgIGFjdGl2YXRlIHJlZGlzXG4gICAgcmVkaXMtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIHJlZGlzXG4gICAgZmFjYWRlLT4-Y29yZTogU2F2ZUNvbW1lbnRcbiAgICBhY3RpdmF0ZSBjb3JlXG4gICAgY29yZS0-Pm15c3FsOiB3cml0ZVxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBmYWNhZGUtPj5yZWRpczogc2V0XG4gICAgYWN0aXZhdGUgcmVkaXNcbiAgICByZWRpcy0tPj5mYWNhZGU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgcmVkaXNcbiAgICBmYWNhZGUtLT4-ZmU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgZmFjYWRlXG4gICAgZGVhY3RpdmF0ZSBmZVxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjpmYWxzZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOmZhbHNlfQ)

```thrift
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
```

[ER Diagram](https://mermaid-js.github.io/mermaid-live-editor/edit/#eyJjb2RlIjoiZXJEaWFncmFtXG4gIHBvc3QgfHwtLW97IGNvbW1lbnQgOiBoYXNcblxuICBwb3N0IHtcbiAgICBiaWdpbnQgaWRcbiAgICB0ZXh0IGNvbnRlbnRcbiAgICBzbWFsbGludCBsYW5nXG4gICAgdmFyY2hhciBwYXNzd2RcbiAgICB2YXJjaGFyIG5pY2tuYW1lXG4gICAgdGlueWludCBpc19kaXNwb3NhYmxlXG4gICAgdGltZXN0YW1wIGNyZWF0ZV90aW1lXG4gICAgdGltZXN0YW1wIHVwZGF0ZV90aW1lXG4gIH1cblxuICBjb21tZW50IHtcbiAgICBiaWdpbnQgaWRcbiAgICBiaWdpbnQgcG9zdF9pZFxuICAgIHRleHQgY29udGVudFxuICAgIHZhcmNoYXIgbmlja25hbWVcbiAgICB0aW1lc3RhbXAgY3JlYXRlX3RpbWVcbiAgICB0aW1lc3RhbXAgdXBkYXRlX3RpbWVcbiAgfVxuIiwibWVybWFpZCI6IntcbiAgXCJ0aGVtZVwiOiBcImRlZmF1bHRcIlxufSIsInVwZGF0ZUVkaXRvciI6ZmFsc2UsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjpmYWxzZX0)

![](https://mermaid.ink/svg/eyJjb2RlIjoiZXJEaWFncmFtXG4gIHBvc3QgfHwtLW97IGNvbW1lbnQgOiBoYXNcblxuICBwb3N0IHtcbiAgICBiaWdpbnQgaWRcbiAgICB0ZXh0IGNvbnRlbnRcbiAgICBzbWFsbGludCBsYW5nXG4gICAgdmFyY2hhciBwYXNzd2RcbiAgICB2YXJjaGFyIG5pY2tuYW1lXG4gICAgdGlueWludCBpc19kaXNwb3NhYmxlXG4gICAgdGltZXN0YW1wIGNyZWF0ZV90aW1lXG4gICAgdGltZXN0YW1wIHVwZGF0ZV90aW1lXG4gIH1cblxuICBjb21tZW50IHtcbiAgICBiaWdpbnQgaWRcbiAgICBiaWdpbnQgcG9zdF9pZFxuICAgIHRleHQgY29udGVudFxuICAgIHZhcmNoYXIgbmlja25hbWVcbiAgICB0aW1lc3RhbXAgY3JlYXRlX3RpbWVcbiAgICB0aW1lc3RhbXAgdXBkYXRlX3RpbWVcbiAgfVxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjpmYWxzZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOmZhbHNlfQ)

```sql
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '评论自增 ID',
  `post_id` bigint NOT NULL COMMENT '文本自增 ID',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '评论内容',
  `nickname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '评论人昵称',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_fk_post_id` (`post_id`),
  CONSTRAINT `idx_fk_post_id` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '文本自增 ID',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文本内容',
  `lang` smallint NOT NULL DEFAULT '0' COMMENT '文本语言',
  `passwd` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '密码',
  `nickname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文本作者昵称',
  `is_disposable` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否阅后即焚',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
```
