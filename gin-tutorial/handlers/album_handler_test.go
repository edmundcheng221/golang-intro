package album_handler_test

import (
	album_handler "EdmundsBankai/golang-intro/gin-tutorial/handlers"
	album_model "EdmundsBankai/golang-intro/gin-tutorial/models"
	"fmt"
	"testing"
)

func TestValidation(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name     string
		input    album_model.Album
		expected bool
	}

	testCases := []TestCase{
		// true:  panics
		// false: no panic
		{
			Name: "Valid",
			input: album_model.Album{
				ID:     "1",
				Title:  "title1",
				Artist: "artist1",
				Price:  1,
			},
			expected: false,
		},
		{
			Name: "Missing id field",
			input: album_model.Album{
				Title:  "title2",
				Artist: "artist2",
				Price:  2,
			},
			expected: true,
		},
		{
			Name: "Missing Title field",
			input: album_model.Album{
				ID:     "3",
				Artist: "artist3",
				Price:  3,
			},
			expected: true,
		},
		{
			Name: "Missing Artist field",
			input: album_model.Album{
				ID:    "4",
				Title: "title4",
				Price: 4,
			},
			expected: true,
		},
		{
			Name: "Missing Price field",
			input: album_model.Album{
				ID:     "5",
				Title:  "title5",
				Artist: "artist5",
			},
			expected: true,
		},
		{
			Name:     "Missing all fields",
			input:    album_model.Album{},
			expected: true,
		},
	}

	for i := range testCases {
		testCase := testCases[i]
		t.Run(fmt.Sprintf("%d - %s", i, testCase.Name), func(t *testing.T) {
			t.Parallel()
			defer func() {
				if r := recover(); r == nil { // no panic occured
					if testCase.expected == true {
						t.Errorf("The function did not panic even tho we expected it to")
					}
				} else { // recovered, meaning that there was a panic
					if testCase.expected == false {
						t.Errorf("The function did a panic even tho we expected it not to")
					}
				}
			}()

			album_handler.ValidateAlbumFields(testCase.input)
		})
	}
}
