// Copyright 2019 The LevelDB-Go and Pebble Authors. All rights reserved. Use
// of this source code is governed by a BSD-style license that can be found in
// the LICENSE file.

// +build !linux

package storage

func (f *syncingFile) init() {
	f.syncTo = f.syncToGeneric
}

func (f *syncingFile) syncData() error {
	return f.File.Sync()
}

func (f *syncingFile) syncToGeneric(_ int64) error {
	return f.Sync()
}
