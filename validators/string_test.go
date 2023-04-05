package validators

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsStringFieldValid(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		tag      string
		wantErr  bool
		checkErr func(err error) bool
	}{
		{
			name:    "valid string min",
			value:   "abc",
			tag:     "min:3",
			wantErr: false,
		},
		{
			name:    "invalid string min",
			value:   "tinkoff",
			tag:     "min:10",
			wantErr: true,
			checkErr: func(err error) bool {
				return errors.Is(err, ErrLenIsLessThenMin)
			},
		},
		{
			name:    "valid string max",
			value:   "max",
			tag:     "max:10",
			wantErr: false,
		},
		{
			name:    "invalid string max",
			value:   "hihihihihihihihi",
			tag:     "max:10",
			wantErr: true,
			checkErr: func(err error) bool {
				return errors.Is(err, ErrLenIsBiggerThenMax)
			},
		},
		{
			name:    "valid string in",
			value:   "bebra",
			tag:     "in:10,bebra",
			wantErr: false,
		},
		{
			name:    "invalid int in",
			value:   "ab",
			tag:     "in:10,5",
			wantErr: true,
			checkErr: func(err error) bool {
				return errors.Is(err, ErrCantFindStringInArray)
			},
		},
		{
			name:    "valid string len",
			value:   "a",
			tag:     "len:1",
			wantErr: false,
		},
		{
			name:    "invalid string len",
			value:   "a",
			tag:     "len:10",
			wantErr: true,
			checkErr: func(err error) bool {
				return errors.Is(err, ErrLenIsInvalid)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsStringFieldValid(tt.value, tt.tag)
			if tt.wantErr {
				assert.Error(t, err)
				assert.True(t, tt.checkErr(err), "test expect an error, but got wrong error type")
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
