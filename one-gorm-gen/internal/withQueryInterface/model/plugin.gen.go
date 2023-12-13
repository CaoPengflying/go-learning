// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePlugin = "plugin"

// Plugin 插件表
type Plugin struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	PluginName   string    `gorm:"column:plugin_name;not null" json:"plugin_name"`
	Status       int32     `gorm:"column:status;not null;default:1;comment:表明插件目前的状态:1创建中、2已完成、-1已删除" json:"status"`      // 表明插件目前的状态:1创建中、2已完成、-1已删除
	StepID       int32     `gorm:"column:step_id;not null;comment:插件可以在哪一步骤使用" json:"step_id"`                            // 插件可以在哪一步骤使用
	InputPrompt  *string   `gorm:"column:input_prompt;comment:预先设置的输入，告诉AI你是那种角色" json:"input_prompt"`                    // 预先设置的输入，告诉AI你是那种角色
	Type         int32     `gorm:"column:type;not null;comment:插件类型" json:"type"`                                         // 插件类型
	ModelIds     string    `gorm:"column:model_ids;not null;comment:模型ids" json:"model_ids"`                              // 模型ids
	CreateTime   time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
	CreateUserID string    `gorm:"column:create_user_id;not null;default:系统管理员;comment:创建人" json:"create_user_id"`        // 创建人
	UpdateUserID string    `gorm:"column:update_user_id;not null;default:系统管理员;comment:创建人" json:"update_user_id"`        // 创建人
	Extra        string    `gorm:"column:extra;not null;comment:额外信息" json:"extra"`                                       // 额外信息
}

// TableName Plugin's table name
func (*Plugin) TableName() string {
	return TableNamePlugin
}