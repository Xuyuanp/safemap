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

import "sync"

// Map is goroutine-safe map.
type Map struct {
	mu sync.RWMutex
	m  map[interface{}]interface{}
}

// New creates a new Map with the capacity.
func New(capacity int) *Map {
	return &Map{
		m: make(map[interface{}]interface{}, capacity),
	}
}

// Set sets a k-v paire.
func (m *Map) Set(key, value interface{}) {
	m.withLockContext(func() {
		m.m[key] = value
	})
}

// Get returns value binded with the key, nil if no value binded.
func (m *Map) Get(key interface{}) (value interface{}) {
	value, _ = m.GetOk(key)
	return
}

// GetOk returns value binded with the key and true, or nil, fale if no value binded.
func (m *Map) GetOk(key interface{}) (value interface{}, ok bool) {
	m.withRLockContext(func() {
		value, ok = m.m[key]
	})
	return
}

// GetMust returns value binded with the key,
// if no value binded, creates a new value by useing the newFunc and adds the value in map.
func (m *Map) GetMust(key interface{}, newFunc func() interface{}) (value interface{}) {
	m.withLockContext(func() {
		var ok bool
		value, ok = m.m[key]
		if !ok {
			value = newFunc()
			m.m[key] = value
		}
	})
	return
}

// Delete deletes the key in map.
func (m *Map) Delete(key interface{}) {
	m.withLockContext(func() {
		delete(m.m, key)
	})
}

// Len returns the map length.
func (m *Map) Len() (length int) {
	m.withRLockContext(func() {
		length = len(m.m)
	})
	return length
}

// Reset clears the map.
func (m *Map) Reset() {
	m.withLockContext(func() {
		m.m = make(map[interface{}]interface{})
	})
}

// All returns all keys and values.
func (m *Map) All() (keys, values []interface{}) {
	m.withRLockContext(func() {
		keys = make([]interface{}, len(m.m))
		values = make([]interface{}, len(m.m))
		i := 0
		for key, value := range m.m {
			keys[i] = key
			values[i] = value
			i++
		}
	})
	return
}

func (m *Map) withLockContext(f func()) {
	m.mu.Lock()
	defer m.mu.Unlock()
	f()
}

func (m *Map) withRLockContext(f func()) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	f()
}
