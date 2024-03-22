package entities

import "time"

type OrderModels struct {
	ID          uint64              `gorm:"column:id;primaryKey" json:"id"`
	UserID      uint64              `gorm:"column:user_id" json:"user_id"`
	Date        time.Time           `gorm:"column:date" json:"date"`
	TotalPrice  float64             `gorm:"column:total_price" json:"total_price"`
	OrderStatus string              `gorm:"column:order_status" json:"order_status"`
	CreatedAt   time.Time           `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt   time.Time           `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt   *time.Time          `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	User        UserModels          `gorm:"foreignKey:UserID" json:"user"`
	OrderDetail []OrderDetailModels `gorm:"foreignKey:OrderID" json:"order_details"`
}

type OrderDetailModels struct {
	ID         uint64         `gorm:"column:id;primaryKey" json:"id"`
	OrderID    string         `gorm:"column:order_id;type:VARCHAR(255)" json:"order_id"`
	ScheduleID uint64         `gorm:"column:schedule_id" json:"schedule_id"`
	SeatID     uint64         `gorm:"column:seat_id" json:"seat_id"`
	Quantity   int            `gorm:"column:quantity" json:"quantity"`
	Price      float64        `gorm:"column:price" json:"price"`
	CreatedAt  time.Time      `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt  *time.Time     `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	Schedule   ScheduleModels `gorm:"foreignKey:ScheduleID" json:"schedule"`
	Seat       SeatModels     `gorm:"foreignKey:SeatID" json:"seat"`
}

func (OrderModels) TableName() string {
	return "order"
}

func (OrderDetailModels) TableName() string {
	return "order_detail"
}
