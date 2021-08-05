package po

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/ameidance/paster_core/model/dto/kitex_gen/ameidance/paster/core"
	"github.com/ameidance/paster_core/util"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// Post [...]
type Post struct {
	ID           int64     `gorm:"primaryKey;column:id;type:bigint;not null"`               // 文本自增 ID
	Content      string    `gorm:"column:content;type:text;not null"`                       // 文本内容
	Lang         int16     `gorm:"column:lang;type:smallint;not null;default:0"`            // 文本语言
	Passwd       string    `gorm:"column:passwd;type:varchar(255)"`                         // 密码
	Nickname     string    `gorm:"column:nickname;type:varchar(20);not null"`               // 文本作者昵称
	IsDisposable bool      `gorm:"column:is_disposable;type:tinyint(1);not null;default:0"` // 是否阅后即焚
	CreateTime   time.Time `gorm:"column:create_time;type:timestamp"`                       // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;type:timestamp"`                       // 更新时间
}

type _PostMgr struct {
	*_BaseMgr
}

// PostMgr open func
func PostMgr(db *gorm.DB) *_PostMgr {
	if db == nil {
		panic(fmt.Errorf("PostMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PostMgr{_BaseMgr: &_BaseMgr{DB: db.Table("post"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PostMgr) GetTableName() string {
	return "post"
}

// Get 获取
func (obj *_PostMgr) Get() (result Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PostMgr) Gets() (results []*Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 文本自增 ID
func (obj *_PostMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithContent content获取 文本内容
func (obj *_PostMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithLang lang获取 文本语言
func (obj *_PostMgr) WithLang(lang int16) Option {
	return optionFunc(func(o *options) { o.query["lang"] = lang })
}

// WithPasswd passwd获取 密码
func (obj *_PostMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithNickname nickname获取 文本作者昵称
func (obj *_PostMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithIsDisposable is_disposable获取 是否阅后即焚
func (obj *_PostMgr) WithIsDisposable(isDisposable bool) Option {
	return optionFunc(func(o *options) { o.query["is_disposable"] = isDisposable })
}

// WithCreateTime create_time获取 创建时间
func (obj *_PostMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_PostMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_PostMgr) GetByOption(opts ...Option) (result Post, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PostMgr) GetByOptions(opts ...Option) (results []*Post, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 文本自增 ID
func (obj *_PostMgr) GetFromID(id int64) (result Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 文本自增 ID
func (obj *_PostMgr) GetBatchFromID(ids []int64) (results []*Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 文本内容
func (obj *_PostMgr) GetFromContent(content string) (results []*Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 文本内容
func (obj *_PostMgr) GetBatchFromContent(contents []string) (results []*Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromLang 通过lang获取内容 文本语言
func (obj *_PostMgr) GetFromLang(lang int16) (results []*Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`lang` = ?", lang).Find(&results).Error

	return
}

// GetBatchFromLang 批量查找 文本语言
func (obj *_PostMgr) GetBatchFromLang(langs []int16) (results []*Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`lang` IN (?)", langs).Find(&results).Error

	return
}

// GetFromPasswd 通过passwd获取内容 密码
func (obj *_PostMgr) GetFromPasswd(passwd string) (results []*Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`passwd` = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量查找 密码
func (obj *_PostMgr) GetBatchFromPasswd(passwds []string) (results []*Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`passwd` IN (?)", passwds).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容 文本作者昵称
func (obj *_PostMgr) GetFromNickname(nickname string) (results []*Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`nickname` = ?", nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量查找 文本作者昵称
func (obj *_PostMgr) GetBatchFromNickname(nicknames []string) (results []*Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`nickname` IN (?)", nicknames).Find(&results).Error

	return
}

// GetFromIsDisposable 通过is_disposable获取内容 是否阅后即焚
func (obj *_PostMgr) GetFromIsDisposable(isDisposable bool) (results []*Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`is_disposable` = ?", isDisposable).Find(&results).Error

	return
}

// GetBatchFromIsDisposable 批量查找 是否阅后即焚
func (obj *_PostMgr) GetBatchFromIsDisposable(isDisposables []bool) (results []*Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`is_disposable` IN (?)", isDisposables).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_PostMgr) GetFromCreateTime(createTime time.Time) (results []*Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_PostMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_PostMgr) GetFromUpdateTime(updateTime time.Time) (results []*Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_PostMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_PostMgr) FetchByPrimaryKey(id int64) (result Post, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Post{}).Where("`id` = ?", id).Find(&result).Error

	return
}

func (post *Post) ConvertToDTO(password string) (*core.PostInfo, error) {
	if post == nil {
		return nil, nil
	}

	encryptedData, err := base64.StdEncoding.DecodeString(post.Content)
	if err != nil {
		klog.Errorf("[Post -> ConvertToDTO] base64 decode failed. err:%v", err)
		return nil, err
	}
	decryptedData := encryptedData
	// decrypt when password exists
	if len(password) > 0 {
		decryptedData, err = util.AesDecrypt(encryptedData, util.GetAesKeyFromString(password))
		if err != nil {
			klog.Errorf("[Post -> ConvertToDTO] aes decrypt failed. err:%v", err)
			return nil, err
		}
	}

	return &core.PostInfo{
		Content:      string(decryptedData),
		Language:     core.LanguageType(post.Lang),
		Nickname:     post.Nickname,
		IsDisposable: post.IsDisposable,
		CreateTime:   thrift.Int64Ptr(post.CreateTime.Unix()),
	}, nil
}

func (post *Post) ConvertFromDTO(info *core.PostInfo, password string) error {
	if info == nil {
		return nil
	}

	post.Lang = int16(info.GetLanguage())
	post.Nickname = info.GetNickname()
	post.IsDisposable = info.GetIsDisposable()
	post.CreateTime = time.Now()
	post.UpdateTime = time.Now()

	encryptedData := []byte(info.GetContent())
	// encrypt when password exists
	if len(password) > 0 {
		post.Passwd = util.GetMd5String([]byte(password))
		var err error
		encryptedData, err = util.AesEncrypt([]byte(info.GetContent()), util.GetAesKeyFromString(password))
		if err != nil {
			klog.Errorf("[Post -> ConvertFromDTO] aes encrypt failed. err:%v", err)
			return err
		}
	}
	post.Content = base64.StdEncoding.EncodeToString(encryptedData)

	return nil
}
