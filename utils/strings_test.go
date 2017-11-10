// Copyright 2017 The kubecfg authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsASCIIIdentifier(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "HelloWorld",
			expected: true,
		},
		{
			input:    "Hello World",
			expected: false,
		},
		{
			input:    "helloworld",
			expected: true,
		},
		{
			input:    "hello-world",
			expected: false,
		},
		{
			input:    "hello世界",
			expected: false,
		},
	}
	for _, test := range tests {
		require.EqualValues(t, test.expected, IsASCIIIdentifier(test.input))
	}
}

func TestPadRows(t *testing.T) {
	tests := []struct {
		input    [][]string
		expected string
	}{
		{
			input:    [][]string{},
			expected: ``,
		},
		{
			input: [][]string{
				[]string{"Hello", "World"},
			},
			expected: "Hello World\n",
		},
		{
			input: [][]string{
				[]string{"Hello", "World"},
				[]string{"Hi", "World"},
			},
			expected: `Hello World
Hi    World
`,
		},
		{
			input: [][]string{
				[]string{"Hello"},
				[]string{"Hi", "World"},
			},
			expected: `Hello
Hi    World
`,
		},
		{
			input: [][]string{
				[]string{},
				[]string{"Hi", "World"},
			},
			expected: `
Hi World
`,
		},
		{
			input: [][]string{
				[]string{"Hello", "World"},
				[]string{""},
			},
			expected: `Hello World

`,
		},
		{
			input: [][]string{
				[]string{""},
				[]string{""},
			},
			expected: `

`,
		},
	}
	for _, test := range tests {
		fmt.Println(test.expected)
		padded, err := PadRows(test.input)
		if err != nil {
			t.Error(err)
		}
		require.EqualValues(t, test.expected, padded)
	}
}
