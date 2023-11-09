package tops

import (
	"testing"
)

func TestTops(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want string
	}{
		{
			name: "When msg is empty, return empty",
			msg:  "",
			want: "",
		},
		{
			name: "When msg has 1 character, return empty",
			msg:  "a",
			want: "",
		},
		{
			name: "When msg has 2 characters long, return 1 top",
			msg:  "12",
			want: "2",
		},
		{
			name: "When msg is long enough for 4 tops, return the expected 4 tops",
			msg:  "abcdefghijklmnopqrstuvwxyz1234",
			want: "3pgb",
		},
		{
			name: "When msg is long enough for 5 tops, return the expected 5 tops",
			msg:  "abcdefghijklmnopqrstuvwxyz1236789ABCDEFGHIJKLMN",
			want: "M3pgb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tops(tt.msg); got != tt.want {
				t.Errorf("Tops() = %v, want %v", got, tt.want)
			}
		})
	}
}
