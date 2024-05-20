package log

import "testing"

func TestLevel_Key(t *testing.T) {
	if InfoLevel.Key() != LevelKey {
		t.Errorf("want: %s, got: %s", LevelKey, InfoLevel.Key())
	}
}

func TestLevel_String(t *testing.T) {
	tests := []struct {
		name string
		l    Level
		want string
	}{
		{
			name: "DEBUG",
			l:    DebugLevel,
			want: "DEBUG",
		},
		{
			name: "INFO",
			l:    InfoLevel,
			want: "INFO",
		},
		{
			name: "WARN",
			l:    WarnLevel,
			want: "WARN",
		},
		{
			name: "ERROR",
			l:    ErrorLevel,
			want: "ERROR",
		},
		{
			name: "FATAL",
			l:    FatalLevel,
			want: "FATAL",
		},
		{
			name: "other",
			l:    10,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseLevel(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Level
	}{
		{
			name: "DEBUG",
			want: DebugLevel,
			s:    "DEBUG",
		},
		{
			name: "INFO",
			want: InfoLevel,
			s:    "INFO",
		},
		{
			name: "WARN",
			want: WarnLevel,
			s:    "WARN",
		},
		{
			name: "ERROR",
			want: ErrorLevel,
			s:    "ERROR",
		},
		{
			name: "FATAL",
			want: FatalLevel,
			s:    "FATAL",
		},
		{
			name: "other",
			want: InfoLevel,
			s:    "other",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLevel(tt.s); got != tt.want {
				t.Errorf("ParseLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
