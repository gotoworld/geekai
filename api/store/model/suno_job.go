package model

import "time"

type SunoJob struct {
	Id           uint `gorm:"primarykey;column:id"`
	UserId       int
	Title        string
	Type         string
	TaskId       string
	ReferenceId  string // 续写的任务id
	Tags         string // 歌曲风格和标签
	Instrumental bool   // 是否生成纯音乐
	ExtendSecs   int    // 续写秒数
	SongId       int    // 续写的歌曲id
	Prompt       string // 提示词
	ThumbImgURL  string // 缩略图 URL
	CoverImgURL  string // 封面图 URL
	AudioURL     string // 音频 URL
	ModelName    string // 模型名称
	Progress     int    // 任务进度
	Duration     int    // 银屏时长，秒
	Publish      bool   // 是否发布
	ErrMsg       string // 错误信息
	RawData      string // 原始数据 json
	Power        int    // 消耗算力
	CreatedAt    time.Time
}

func (SunoJob) TableName() string {
	return "chatgpt_suno_jobs"
}
