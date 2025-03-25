package model

type Category struct {
    ID     uint     `gorm:"primaryKey" json:"id"`
    Name   string   `gorm:"unique;not null" json:"name"`
    Vocabs []Vocab  `gorm:"many2many:vocab_categories;constraint:OnDelete:CASCADE;" json:"vocabs,omitempty"`
}

func (Category) TableName() string {
    return "categories"
}