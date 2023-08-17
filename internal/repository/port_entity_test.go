package repository

import (
	"testing"
	"time"

	"github.com/dkhrunov/hexagonal-architecture/internal/domain"
	"github.com/stretchr/testify/require"
)

func Test_portStoreToDomain(t *testing.T) {
	type args struct {
		p *PortEntity
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Port
		wantErr bool
	}{
		{
			name: "should return error when store port is nil",
			args: args{
				p: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return domain port when store port is not nil",
			args: args{
				p: newTestStorePort(t),
			},
			want:    newTestDomainPort(t),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := PortEntityToDomain(tt.args.p)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}

const testString = "test"

func newTestStorePort(t *testing.T) *PortEntity {
	t.Helper()
	return &PortEntity{
		ID:          testString,
		Name:        testString,
		Code:        testString,
		City:        testString,
		Country:     testString,
		Alias:       []string{testString},
		Regions:     []string{testString},
		Coordinates: [2]float64{1.0, 2.0},
		Province:    testString,
		Timezone:    testString,
		Unlocs:      []string{testString},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func newTestDomainPort(t *testing.T) *domain.Port {
	t.Helper()
	port, err := domain.NewPort(testString, testString, testString, testString, testString,
		[]string{testString}, []string{testString}, [2]float64{1.0, 2.0}, testString, testString, []string{testString})
	require.NoError(t, err)
	return port
}
