// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameStep = "step"

// Step 步骤表
type Step struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TaskID       int32     `gorm:"column:task_id;not null" json:"taskId"`
	StepName     string    `gorm:"column:step_name;not null" json:"stepName"`
	Type         int32     `gorm:"column:type;not null;comment:类型：wiki、chat..." json:"type"`                             // 类型：wiki、chat...
	InputPrompt  *string   `gorm:"column:input_prompt;comment:预先设置的输入，告诉AI你是那种角色" json:"inputPrompt"`                    // 预先设置的输入，告诉AI你是那种角色
	Status       int32     `gorm:"column:status;not null;default:1;comment:状态：1创建中、2已完成、-1已删除" json:"status"`            // 状态：1创建中、2已完成、-1已删除
	ModelIds     string    `gorm:"column:model_ids;not null;comment:支持的模型ids" json:"modelIds"`                           // 支持的模型ids
	CreateTime   time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updateTime"` // 更新时间
	CreateUserID string    `gorm:"column:create_user_id;not null;default:系统管理员;comment:创建人" json:"createUserId"`         // 创建人
	UpdateUserID string    `gorm:"column:update_user_id;not null;default:系统管理员;comment:创建人" json:"updateUserId"`         // 创建人
	Extra        *string   `gorm:"column:extra" json:"extra"`
}

// TableName Step's table name
func (*Step) TableName() string {
	return TableNameStep
}
