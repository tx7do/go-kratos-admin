package service

import "testing"

func TestMenuListToQueryString(t *testing.T) {
	type args struct {
		menus []uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Single menu ID",
			args: args{menus: []uint32{1}},
			want: `{"id__in":"[\"1\"]","status":"ON","type__not":"BUTTON"}`,
		},
		{
			name: "Multiple menu IDs",
			args: args{menus: []uint32{1, 2, 3}},
			want: `{"id__in":"[\"1\", \"2\", \"3\"]","status":"ON","type__not":"BUTTON"}`,
		},
		{
			name: "No menu IDs",
			args: args{menus: []uint32{}},
			want: `{"id__in":"[]","status":"ON","type__not":"BUTTON"}`,
		},
		{
			name: "Large menu IDs",
			args: args{menus: []uint32{1234567890, 987654321}},
			want: `{"id__in":"[\"1234567890\", \"987654321\"]","status":"ON","type__not":"BUTTON"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RouterService{}
			if got := s.menuListToQueryString(tt.args.menus, false); got != tt.want {
				t.Errorf("menuListToQueryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
