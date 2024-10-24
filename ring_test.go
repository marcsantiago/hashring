package go_hashring

import (
	"math"
	"strconv"
	"testing"
)

func TestHashRing_AddNode(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		args struct {
			keys []string
		}
		want int
	}{
		{
			name: "Test AddNode",
			args: struct{ keys []string }{keys: []string{"1", "2", "3", "4"}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ring := New(nil)
			ring.AddNode(tt.args.keys...)
			if len(ring.keys) != tt.want {
				t.Fatalf("AddNode() = %v, want %v", len(ring.keys), tt.want)
			}
		})
	}
}

func TestHashRing_GetTargetNode(t *testing.T) {
	t.Parallel()
	type fields struct {
		nodes []string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		iteration int
		want      string
		wantErr   bool
	}{
		{
			name:      "Test GetTargetNode",
			fields:    fields{nodes: []string{"1", "2", "3", "4"}},
			args:      args{key: "id_d5d25b3b-5acc-49fb-8cc7-0798ceeece69"},
			iteration: 1_000_000,
			want:      "2",
			wantErr:   false,
		},
		{
			name:      "Test GetTargetNode",
			fields:    fields{nodes: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}},
			args:      args{key: "id_ced5c816-f8a8-4f6e-bcc7-61472f099857"},
			iteration: 1_000_000,
			want:      "8",
			wantErr:   false,
		},

		{
			name:      "Test GetTargetNode",
			fields:    fields{nodes: nil},
			args:      args{key: "id_ced5c816-f8a8-4f6e-bcc7-61472f099857"},
			iteration: 1_000_000,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.fields.nodes)
			for i := 0; i < tt.iteration; i++ {
				got, err := h.GetTargetNode(tt.args.key)
				if (err != nil) != tt.wantErr {
					t.Fatalf("GetTargetNode() error = %v, wantErr %v", err, tt.wantErr)
				}
				if got != tt.want {
					t.Fatalf("GetTargetNode() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestHashRing_Distribution(t *testing.T) {
	t.Parallel()
	replicas := []string{"1", "2", "3", "4"}
	nodes := New(replicas)

	iter := 1_000_000

	distributionMap := make(map[string]int)
	for i := 0; i < iter; i++ {
		targetID := "id_" + strconv.Itoa(i)
		node, _ := nodes.GetTargetNode(targetID)
		distributionMap[node]++
	}

	if len(distributionMap) != len(replicas) {
		t.Fatalf("GetTargetNode() got = %v, want %v", len(distributionMap), len(replicas))
	}

	tolerance := 0.1
	expected := 0.25
	for _, node := range replicas {
		count := distributionMap[node]
		percentage := float64(count) / float64(iter)
		if !WithinTolerance(expected, percentage, tolerance) {
			t.Fatalf("GetTargetNode() got = %v, want %v, tolerance %v, got %v", percentage, 0.25, tolerance, percentage-expected)
		}
	}
}

// goos: darwin
// goarch: arm64
// pkg: github.com/marcsantiago/go-hashring
// cpu: Apple M2
// BenchmarkHashRing_GetTargetNode-8   	 4668703	       256.5 ns/op	       0 B/op	       0 allocs/op
// BenchmarkHashRing_GetTargetNode-8   	 4722736	       253.8 ns/op	       0 B/op	       0 allocs/op
// BenchmarkHashRing_GetTargetNode-8   	 4702320	       254.4 ns/op	       0 B/op	       0 allocs/op
func BenchmarkHashRing_GetTargetNode(b *testing.B) {
	replicas := []string{"1", "2", "3", "4"}
	nodes := New(replicas)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = nodes.GetTargetNode("id")
	}
}

func WithinTolerance(expected, got, tolerance float64) bool {
	if expected == got {
		return true
	}
	d := math.Abs(expected - got)
	if got == 0 {
		return d < tolerance
	}
	return (d / math.Abs(got)) < tolerance
}
