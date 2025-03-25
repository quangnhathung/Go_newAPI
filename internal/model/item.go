package model

type Vocab struct {
    ID          uint        `gorm:"primaryKey" json:"id"`
    Vocabulary  string      `gorm:"not null" json:"vocabulary"`
    Image       string      `json:"image,omitempty"`
    AudioUS     string      `json:"audio_us,omitempty"`
    AudioGB     string      `json:"audio_gb,omitempty"`
    Categories  []Category  `gorm:"many2many:vocab_categories;constraint:OnDelete:CASCADE;" json:"categories,omitempty"`
}

func (Vocab) TableName() string {
    return "vocabs"
}
