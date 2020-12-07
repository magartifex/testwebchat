package messages

import (
	"reflect"
	"testing"
)

func TestSetHistory(t *testing.T) {
	type args struct {
		userID string
		text   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal",
			args: args{
				userID: "qwerty123456",
				text:   "Hello, world",
			},
		},
		{
			name: "empty",
			args: args{
				userID: "",
				text:   "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := SetHistory(tt.args.userID, tt.args.text)
			if got.UserID != tt.args.userID {
				t.Errorf("SetHistory() userID got = %v, args %v", got, tt.args)
			}
			if got.Message != tt.args.text {
				t.Errorf("SetHistory() message got = %v, args %v", got, tt.args)
			}

			history.Lock()
			defer history.Unlock()

			var collision bool
			for _, val := range history.messages {
				if reflect.DeepEqual(val, got) {
					collision = true
					break
				}
			}

			if !collision {
				t.Errorf("SetHistory() history got = %v, args %v", got, tt.args)
			}
		})
	}
}
