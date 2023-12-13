// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameStepLog = "step_log"

// StepLog step生成的日志
type StepLog struct {
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TaskID          int32     `gorm:"column:task_id;not null" json:"taskId"`
	TaskLogID       int32     `gorm:"column:task_log_id;not null" json:"taskLogId"`
	StepID          int32     `gorm:"column:step_id;not null" json:"stepId"`
	AdditionalInput *string   `gorm:"column:additional_input;comment:这一步骤里的附加输入" json:"additionalInput"`                    // 这一步骤里的附加输入
	Input           string    `gorm:"column:input;not null;comment:这一步骤里的输入" json:"input"`                                  // 这一步骤里的输入
	Output          string    `gorm:"column:output;not null;comment:这一步骤里的输出" json:"output"`                                // 这一步骤里的输出
	PluginID        int32     `gorm:"column:plugin_id;not null;comment:使用的插件id" json:"pluginId"`                            // 使用的插件id
	CreateTime      time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateTime      time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updateTime"` // 更新时间
	CreateUserID    string    `gorm:"column:create_user_id;not null;default:系统管理员;comment:创建人" json:"createUserId"`         // 创建人
	UpdateUserID    string    `gorm:"column:update_user_id;not null;default:系统管理员;comment:创建人" json:"updateUserId"`         // 创建人
	Extra           *string   `gorm:"column:extra" json:"extra"`
	Status          int32     `gorm:"column:status;not null;comment:状态:1创建中、2已完成、-1已删除" json:"status"` // 状态:1创建中、2已完成、-1已删除
}

// TableName StepLog's table name
func (*StepLog) TableName() string {
	return TableNameStepLog
}
