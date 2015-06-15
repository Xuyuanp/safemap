/*
 * Copyright 2015 Xuyuan Pang
 * Author: Xuyuan Pang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package safemap

import "testing"

func TestSafeMap(t *testing.T) {
	sm := New(10)

	sm.Set("foo", "bar")

	if v := sm.Get("foo"); v != "bar" {
		t.Errorf("bar expected, %s got", v)
	}

	if _, ok := sm.GetOk("foo"); !ok {
		t.Errorf("true expected, %s got", ok)
	}

	if length := sm.Len(); length != 1 {
		t.Errorf("1 expected, %d got", length)
	}

	if v := sm.GetMust("fizz", func() interface{} {
		return "buzz"
	}); v != "buzz" {
		t.Errorf("buzz expected, %s got", v)
	}

	keys, values := sm.All()
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		value := values[i]
		if v := sm.Get(key); v != value {
			t.Errorf("%v expect, %v got", value, v)
		}
	}

	sm.Delete("foo")

	if _, ok := sm.GetOk("foo"); ok {
		t.Errorf("delete failed")
	}
}
