package v1

import "time"

type ObjectMeta struct {
	ID        uint64    `json:"id,omitempty" gorm:"primaryKey;AUTO_INCREMENT;column:id"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdateAt  time.Time `json:"updateAt" gorm:"column:updateAt"`
}

type ObjectMeatAccess interface {
	GetObjectMeta() Object
}

type Object interface {
	GetID() uint64
	SetID(id uint64)
	GetCreateAt() time.Time
	SetCreateAt(t time.Time)
	GetUpdateAt() time.Time
	SetUpdateAt(t time.Time)
}

func (meta *ObjectMeta) GetObjectMeta() Object { return meta }

func (meta *ObjectMeta) GetID() uint64           { return meta.ID }
func (meta *ObjectMeta) SetID(id uint64)         { meta.ID = id }
func (meta *ObjectMeta) GetCreateAt() time.Time  { return meta.CreatedAt }
func (meta *ObjectMeta) SetCreateAt(t time.Time) { meta.CreatedAt = t }
func (meta *ObjectMeta) GetUpdateAt() time.Time  { return meta.UpdateAt }
func (meta *ObjectMeta) SetUpdateAt(t time.Time) { meta.UpdateAt = t }
