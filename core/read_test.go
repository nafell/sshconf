package core

import "testing"

func TestSplitBlocks(t *testing.T) {
	type want struct{
		lines int
		blocks int
		firstBlockLines int
	}

	tests:= []struct {
		name string
		args string
		want want
	}{
		{
			name: "normal (two entries)",
			args: `Host example1
  HostName example.com
  Port 22
  User joe
  PreferredAuthentications publickey
  IdentityFile /Users/joe/.ssh/aaa_rsa
  IdentitiesOnly yes
Host example2
  HostName example.com
  Port 10022
  User joe
  PreferredAuthentications publickey
  IdentityFile /Users/joe/.ssh/aaa_rsa
  IdentitiesOnly yes`,
			want: want{lines:14, blocks:2, firstBlockLines:7},
		},
		{
			name: "normal (one entry)",
			args: `Host example1
  HostName example.com
  Port 22
  User joe
  PreferredAuthentications publickey
  IdentityFile /Users/joe/.ssh/aaa_rsa
  IdentitiesOnly yes`,
			want: want{lines:7, blocks:1, firstBlockLines:7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SplitEntryBlocks(tt.args)
			if (err != nil) {
				t.Errorf("%v", err)
			}
			if len(got.Lines) != tt.want.lines {
				t.Errorf("SplitHostBlocks() lines = %v, want %v", len(got.Lines), tt.want.lines)
			}
			if len(got.Blocks) != tt.want.blocks {
				t.Errorf("SplitHostBlocks() blocks = %v, want %v", len(got.Blocks), tt.want.blocks)
			}
		})
	}
}
