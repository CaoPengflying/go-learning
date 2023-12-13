// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTaskStep = "task_step"

// TaskStep 任务流-步骤表，标明该步骤应该在任务流的哪一步
type TaskStep struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TaskID       int32     `gorm:"column:task_id;not null" json:"taskId"`
	StepID       int32     `gorm:"column:step_id;not null" json:"stepId"`
	Sort         int32     `gorm:"column:sort;not null;comment:用来标记同一task中step的顺序" json:"sort"`                          // 用来标记同一task中step的顺序
	CreateTime   time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updateTime"` // 更新时间
	CreateUserID string    `gorm:"column:create_user_id;not null;default:系统管理员;comment:创建人" json:"createUserId"`         // 创建人
	UpdateUserID string    `gorm:"column:update_user_id;not null;default:系统管理员;comment:创建人" json:"updateUserId"`         // 创建人
	Extra        *string   `gorm:"column:extra" json:"extra"`
	PreStep      int32     `gorm:"column:pre_step;not null;comment:这一步的输入应该取哪一步的输出" json:"preStep"`  // 这一步的输入应该取哪一步的输出
	IsAutoNext   bool      `gorm:"column:is_auto_next;not null;comment:是否自动执行下一步" json:"isAutoNext"` // 是否自动执行下一步
}

// TableName TaskStep's table name
func (*TaskStep) TableName() string {
	return TableNameTaskStep
}
