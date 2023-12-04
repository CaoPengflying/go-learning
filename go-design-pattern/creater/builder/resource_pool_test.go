package builder

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestResourcePoolConfigBuilder_Build(t *testing.T) {
	tests := []struct {
		name    string
		builder *ResourcePoolConfigBuilder
		want    *ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "name empty",
			builder: &ResourcePoolConfigBuilder{
				name:     "",
				maxTotal: 0,
				maxIdle:  0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "name test",
			builder: &ResourcePoolConfigBuilder{
				name:     "test",
				maxTotal: 0,
			},
			want:    &ResourcePoolConfig{name: "test", maxTotal: 10, maxIdle: 9, minIdle: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.builder.Build()
			require.Equalf(t, tt.wantErr, err != nil, "build() error = %v,wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}

}

func TestConfigOption(t *testing.T) {
	opts := []OptFun{
		func(option *ResourcePoolConfigOption) {
			option.minIdle = 5
		},
		func(option *ResourcePoolConfigOption) {
			option.maxIdle = 20
		},
		func(option *ResourcePoolConfigOption) {
			option.maxTotal = 30
		},
	}
	poolconfig, _ := NewResourcePoolConfig("123", opts...)
	t.Log(poolconfig)
}
