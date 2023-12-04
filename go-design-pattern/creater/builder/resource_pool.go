package builder

import (
	"errors"
	"fmt"
)

type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

// 建造者模式
type ResourcePoolConfigBuilder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

const (
	defaultMaxTotal = 10
	defaultMaxIdle  = 9
	defaultMinIdle  = 1
)

func (builder *ResourcePoolConfigBuilder) Build() (*ResourcePoolConfig, error) {
	if len(builder.name) == 0 {
		return nil, errors.New("name is empty")
	}
	// 设置默认值
	if builder.minIdle == 0 {
		builder.minIdle = defaultMinIdle
	}

	if builder.maxIdle == 0 {
		builder.maxIdle = defaultMaxIdle
	}

	if builder.maxTotal == 0 {
		builder.maxTotal = defaultMaxTotal
	}

	if builder.maxTotal < builder.maxIdle {
		return nil, fmt.Errorf("max total(%d) cannot < max idle(%d)", builder.maxTotal, builder.maxIdle)
	}

	if builder.minIdle > builder.maxIdle {
		return nil, fmt.Errorf("max idle(%d) cannot < min idle(%d)", builder.maxIdle, builder.minIdle)
	}
	return &ResourcePoolConfig{
		name:     builder.name,
		maxTotal: builder.maxTotal,
		maxIdle:  builder.maxIdle,
		minIdle:  builder.minIdle,
	}, nil
}

func (builder *ResourcePoolConfigBuilder) SetName(name string) (*ResourcePoolConfigBuilder, error) {
	if len(name) == 0 {
		return nil, errors.New("name can't empty")
	}
	builder.name = name
	return builder, nil
}

//常用的赋值方式
type ResourcePoolConfigOption struct {
	maxTotal int
	maxIdle  int
	minIdle  int
}

type OptFun func(option *ResourcePoolConfigOption)

func NewResourcePoolConfig(name string, opts ...OptFun) (*ResourcePoolConfig, error) {
	if name == "" {
		return nil, errors.New("name can't be empty")
	}

	option := &ResourcePoolConfigOption{
		maxTotal: defaultMaxTotal,
		maxIdle:  defaultMaxIdle,
		minIdle:  defaultMinIdle,
	}

	for _, opt := range opts {
		opt(option)
	}

	return &ResourcePoolConfig{
		name:     name,
		maxTotal: option.maxTotal,
		maxIdle:  option.maxIdle,
		minIdle:  option.minIdle,
	}, nil

}
