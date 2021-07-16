package po

import (
    "context"
    "encoding/base64"
    "fmt"
    "time"

    "github.com/ameidance/paster_core/model/dto/kitex_gen/ameidance/paster/core"
    "github.com/ameidance/paster_core/util"
    "github.com/apache/thrift/lib/go/thrift"
    "github.com/bytedance/gopkg/util/logger"
    "gorm.io/gorm"
)

// Comment [...]
type Comment struct {
    ID         int64     `gorm:"primaryKey;column:id;type:bigint;not null"`                // 评论自增 ID
    PostID     int64     `gorm:"index:idx_fk_post_id;column:post_id;type:bigint;not null"` // 文本自增 ID
    Post       Post      `gorm:"joinForeignKey:post_id;foreignKey:id"`
    Content    string    `gorm:"column:content;type:text;not null"`         // 评论内容
    Nickname   string    `gorm:"column:nickname;type:varchar(20);not null"` // 评论人昵称
    CreateTime time.Time `gorm:"column:create_time;type:timestamp"`         // 创建时间
    UpdateTime time.Time `gorm:"column:update_time;type:timestamp"`         // 更新时间
}

type _CommentMgr struct {
    *_BaseMgr
}

// CommentMgr open func
func CommentMgr(db *gorm.DB) *_CommentMgr {
    if db == nil {
        panic(fmt.Errorf("CommentMgr need init by db"))
    }
    ctx, cancel := context.WithCancel(context.Background())
    return &_CommentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("comment"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CommentMgr) GetTableName() string {
    return "comment"
}

// Get 获取
func (obj *_CommentMgr) Get() (result Comment, err error) {
    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Find(&result).Error
    if err == nil && obj.isRelated {
        if err = obj.New().Table("post").Where("id = ?", result.PostID).Find(&result.Post).Error; err != nil { //
            if err != gorm.ErrRecordNotFound { // 非 没找到
                return
            }
        }
    }

    return
}

// Gets 获取批量结果
func (obj *_CommentMgr) Gets() (results []*Comment, err error) {
    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Find(&results).Error
    if err == nil && obj.isRelated {
        for i := 0; i < len(results); i++ {
            if err = obj.New().Table("post").Where("id = ?", results[i].PostID).Find(&results[i].Post).Error; err != nil { //
                if err != gorm.ErrRecordNotFound { // 非 没找到
                    return
                }
            }
        }
    }
    return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 评论自增 ID
func (obj *_CommentMgr) WithID(id int64) Option {
    return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPostID post_id获取 文本自增 ID
func (obj *_CommentMgr) WithPostID(postID int64) Option {
    return optionFunc(func(o *options) { o.query["post_id"] = postID })
}

// WithContent content获取 评论内容
func (obj *_CommentMgr) WithContent(content string) Option {
    return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithNickname nickname获取 评论人昵称
func (obj *_CommentMgr) WithNickname(nickname string) Option {
    return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithCreateTime create_time获取 创建时间
func (obj *_CommentMgr) WithCreateTime(createTime time.Time) Option {
    return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_CommentMgr) WithUpdateTime(updateTime time.Time) Option {
    return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_CommentMgr) GetByOption(opts ...Option) (result Comment, err error) {
    options := options{
        query: make(map[string]interface{}, len(opts)),
    }
    for _, o := range opts {
        o.apply(&options)
    }

    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where(options.query).Find(&result).Error
    if err == nil && obj.isRelated {
        if err = obj.New().Table("post").Where("id = ?", result.PostID).Find(&result.Post).Error; err != nil { //
            if err != gorm.ErrRecordNotFound { // 非 没找到
                return
            }
        }
    }

    return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CommentMgr) GetByOptions(opts ...Option) (results []*Comment, err error) {
    options := options{
        query: make(map[string]interface{}, len(opts)),
    }
    for _, o := range opts {
        o.apply(&options)
    }

    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where(options.query).Find(&results).Error
    if err == nil && obj.isRelated {
        for i := 0; i < len(results); i++ {
            if err = obj.New().Table("post").Where("id = ?", results[i].PostID).Find(&results[i].Post).Error; err != nil { //
                if err != gorm.ErrRecordNotFound { // 非 没找到
                    return
                }
            }
        }
    }
    return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 评论自增 ID
func (obj *_CommentMgr) GetFromID(id int64) (result Comment, err error) {
    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`id` = ?", id).Find(&result).Error
    if err == nil && obj.isRelated {
        if err = obj.New().Table("post").Where("id = ?", result.PostID).Find(&result.Post).Error; err != nil { //
            if err != gorm.ErrRecordNotFound { // 非 没找到
                return
            }
        }
    }

    return
}

// GetBatchFromID 批量查找 评论自增 ID
func (obj *_CommentMgr) GetBatchFromID(ids []int64) (results []*Comment, err error) {
    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`id` IN (?)", ids).Find(&results).Error
    if err == nil && obj.isRelated {
        for i := 0; i < len(results); i++ {
            if err = obj.New().Table("post").Where("id = ?", results[i].PostID).Find(&results[i].Post).Error; err != nil { //
                if err != gorm.ErrRecordNotFound { // 非 没找到
                    return
                }
            }
        }
    }
    return
}

// GetFromPostID 通过post_id获取内容 文本自增 ID
func (obj *_CommentMgr) GetFromPostID(postID int64) (results []*Comment, err error) {
    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`post_id` = ?", postID).Find(&results).Error
    if err == nil && obj.isRelated {
        for i := 0; i < len(results); i++ {
            if err = obj.New().Table("post").Where("id = ?", results[i].PostID).Find(&results[i].Post).Error; err != nil { //
                if err != gorm.ErrRecordNotFound { // 非 没找到
                    return
                }
            }
        }
    }
    return
}

// GetBatchFromPostID 批量查找 文本自增 ID
func (obj *_CommentMgr) GetBatchFromPostID(postIDs []int64) (results []*Comment, err error) {
    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`post_id` IN (?)", postIDs).Find(&results).Error
    if err == nil && obj.isRelated {
        for i := 0; i < len(results); i++ {
            if err = obj.New().Table("post").Where("id = ?", results[i].PostID).Find(&results[i].Post).Error; err != nil { //
                if err != gorm.ErrRecordNotFound { // 非 没找到
                    return
                }
            }
        }
    }
    return
}

// GetFromContent 通过content获取内容 评论内容
func (obj *_CommentMgr) GetFromContent(content string) (results []*Comment, err error) {
    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`content` = ?", content).Find(&results).Error
    if err == nil && obj.isRelated {
        for i := 0; i < len(results); i++ {
            if err = obj.New().Table("post").Where("id = ?", results[i].PostID).Find(&results[i].Post).Error; err != nil { //
                if err != gorm.ErrRecordNotFound { // 非 没找到
                    return
                }
            }
        }
    }
    return
}

// GetBatchFromContent 批量查找 评论内容
func (obj *_CommentMgr) GetBatchFromContent(contents []string) (results []*Comment, err error) {
    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`content` IN (?)", contents).Find(&results).Error
    if err == nil && obj.isRelated {
        for i := 0; i < len(results); i++ {
            if err = obj.New().Table("post").Where("id = ?", results[i].PostID).Find(&results[i].Post).Error; err != nil { //
                if err != gorm.ErrRecordNotFound { // 非 没找到
                    return
                }
            }
        }
    }
    return
}

// GetFromNickname 通过nickname获取内容 评论人昵称
func (obj *_CommentMgr) GetFromNickname(nickname string) (results []*Comment, err error) {
    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`nickname` = ?", nickname).Find(&results).Error
    if err == nil && obj.isRelated {
        for i := 0; i < len(results); i++ {
            if err = obj.New().Table("post").Where("id = ?", results[i].PostID).Find(&results[i].Post).Error; err != nil { //
                if err != gorm.ErrRecordNotFound { // 非 没找到
                    return
                }
            }
        }
    }
    return
}

// GetBatchFromNickname 批量查找 评论人昵称
func (obj *_CommentMgr) GetBatchFromNickname(nicknames []string) (results []*Comment, err error) {
    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`nickname` IN (?)", nicknames).Find(&results).Error
    if err == nil && obj.isRelated {
        for i := 0; i < len(results); i++ {
            if err = obj.New().Table("post").Where("id = ?", results[i].PostID).Find(&results[i].Post).Error; err != nil { //
                if err != gorm.ErrRecordNotFound { // 非 没找到
                    return
                }
            }
        }
    }
    return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_CommentMgr) GetFromCreateTime(createTime time.Time) (results []*Comment, err error) {
    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`create_time` = ?", createTime).Find(&results).Error
    if err == nil && obj.isRelated {
        for i := 0; i < len(results); i++ {
            if err = obj.New().Table("post").Where("id = ?", results[i].PostID).Find(&results[i].Post).Error; err != nil { //
                if err != gorm.ErrRecordNotFound { // 非 没找到
                    return
                }
            }
        }
    }
    return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_CommentMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Comment, err error) {
    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
    if err == nil && obj.isRelated {
        for i := 0; i < len(results); i++ {
            if err = obj.New().Table("post").Where("id = ?", results[i].PostID).Find(&results[i].Post).Error; err != nil { //
                if err != gorm.ErrRecordNotFound { // 非 没找到
                    return
                }
            }
        }
    }
    return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_CommentMgr) GetFromUpdateTime(updateTime time.Time) (results []*Comment, err error) {
    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`update_time` = ?", updateTime).Find(&results).Error
    if err == nil && obj.isRelated {
        for i := 0; i < len(results); i++ {
            if err = obj.New().Table("post").Where("id = ?", results[i].PostID).Find(&results[i].Post).Error; err != nil { //
                if err != gorm.ErrRecordNotFound { // 非 没找到
                    return
                }
            }
        }
    }
    return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_CommentMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Comment, err error) {
    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
    if err == nil && obj.isRelated {
        for i := 0; i < len(results); i++ {
            if err = obj.New().Table("post").Where("id = ?", results[i].PostID).Find(&results[i].Post).Error; err != nil { //
                if err != gorm.ErrRecordNotFound { // 非 没找到
                    return
                }
            }
        }
    }
    return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_CommentMgr) FetchByPrimaryKey(id int64) (result Comment, err error) {
    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`id` = ?", id).Find(&result).Error
    if err == nil && obj.isRelated {
        if err = obj.New().Table("post").Where("id = ?", result.PostID).Find(&result.Post).Error; err != nil { //
            if err != gorm.ErrRecordNotFound { // 非 没找到
                return
            }
        }
    }

    return
}

// FetchIndexByIDxFkPostID  获取多个内容
func (obj *_CommentMgr) FetchIndexByIDxFkPostID(postID int64) (results []*Comment, err error) {
    err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`post_id` = ?", postID).Find(&results).Error
    if err == nil && obj.isRelated {
        for i := 0; i < len(results); i++ {
            if err = obj.New().Table("post").Where("id = ?", results[i].PostID).Find(&results[i].Post).Error; err != nil { //
                if err != gorm.ErrRecordNotFound { // 非 没找到
                    return
                }
            }
        }
    }
    return
}

type Comments []*Comment

func (comments Comments) ConvertToDTO(ctx context.Context, password string) ([]*core.CommentInfo, error) {
    if len(comments) == 0 {
        return nil, nil
    }

    result := make([]*core.CommentInfo, 0)
    for _, comment := range comments {
        encryptedData, err := base64.StdEncoding.DecodeString(comment.Content)
        if err != nil {
            logger.CtxErrorf(ctx, "[Comments -> ConvertToDTO] base64 decode failed. err:%v", err)
            return nil, err
        }
        decryptedData := encryptedData
        // decrypt when password exists
        if len(password) > 0 {
            decryptedData, err = util.AesDecrypt(encryptedData, util.GetAesKeyFromString(password))
            if err != nil {
                logger.CtxErrorf(ctx, "[Comments -> ConvertToDTO] aes decrypt failed. err:%v", err)
                return nil, err
            }
        }
        result = append(result, &core.CommentInfo{
            Content:    string(decryptedData),
            Nickname:   comment.Nickname,
            CreateTime: thrift.Int64Ptr(comment.CreateTime.Unix()),
        })
    }

    return result, nil
}

func (comment *Comment) ConvertFromDTO(ctx context.Context, info *core.CommentInfo, postId int64, password string) error {
    if info == nil {
        return nil
    }

    comment.PostID = postId
    comment.Nickname = info.GetNickname()
    comment.CreateTime = time.Now()
    comment.UpdateTime = time.Now()

    encryptedData := []byte(info.GetContent())
    // encrypt when password exists
    if len(password) > 0 {
        var err error
        encryptedData, err = util.AesEncrypt([]byte(info.GetContent()), util.GetAesKeyFromString(password))
        if err != nil {
            logger.CtxErrorf(ctx, "[Comment -> ConvertFromDTO] aes encrypt failed. err:%v", err)
            return err
        }
    }
    comment.Content = base64.StdEncoding.EncodeToString(encryptedData)

    return nil
}
