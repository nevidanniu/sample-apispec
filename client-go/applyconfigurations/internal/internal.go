/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package internal

import (
	"fmt"
	"sync"

	typed "sigs.k8s.io/structured-merge-diff/v4/typed"
)

func Parser() *typed.Parser {
	parserOnce.Do(func() {
		var err error
		parser, err = typed.NewParser(schemaYAML)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse schema: %v", err))
		}
	})
	return parser
}

var parserOnce sync.Once
var parser *typed.Parser
var schemaYAML = typed.YAMLObject(`types:
- name: __untyped_atomic_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
- name: __untyped_deduced_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_deduced_
    elementRelationship: separable
`)
