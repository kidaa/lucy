/* Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lucy

import "git-wip-us.apache.org/repos/asf/lucy-clownfish.git/runtime/go/clownfish"
import "testing"
import "reflect"

func TestStuff(t *testing.T) {
	NewSchema()
}

func TestOpenIndexer(t *testing.T) {
	_, err := OpenIndexer(&OpenIndexerArgs{Index: "notalucyindex"})
	if _, ok := err.(clownfish.Err); !ok {
		t.Error("Didn't catch exception opening indexer")
	}
}

func TestRegex(t *testing.T) {
	tokenizer := NewRegexTokenizer("\\S+")
	var expected []interface{} = []interface{}{"foo", "bar", "baz"}
	got := tokenizer.Split("foo bar baz")
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
