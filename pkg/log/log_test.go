// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package log

import (
	"io/ioutil"
)

//func Test_LogNewEmpty(t *testing.T) {
//	path := mustTempFile()
//	defer os.Remove(path)
//
//	l, err := NewLog(path)
//	if err != nil {
//		t.Fatalf("failed to create log: %s", err)
//	}
//	fi, err := l.FirstIndex()
//	if err != nil {
//		t.Fatalf("failed to get first index: %s", err)
//	}
//	if fi != 0 {
//		t.Fatalf("got non-zero value for first index of empty log: %d", fi)
//	}
//
//	li, err := l.LastIndex()
//	if err != nil {
//		t.Fatalf("failed to get last index: %s", err)
//	}
//	if li != 0 {
//		t.Fatalf("got non-zero value for last index of empty log: %d", li)
//	}
//
//	lci, err := l.LastCommandIndex()
//	if err != nil {
//		t.Fatalf("failed to get last command index: %s", err)
//	}
//	if lci != 0 {
//		t.Fatalf("got wrong value for last command index of not empty log: %d", lci)
//	}
//
//}
//
//func Test_LogNewExistNotEmpty(t *testing.T) {
//	path := mustTempFile()
//	defer os.Remove(path)
//
//	// JSON some entries directory to the BoltDB Raft store.
//	bs, err := raftboltdb.NewBoltStore(path)
//	if err != nil {
//		t.Fatalf("failed to create bolt store: %s", err)
//	}
//	for i := 4; i > 0; i-- {
//		if err := bs.StoreLog(&raft.Log{
//			Index: uint64(i),
//		}); err != nil {
//			t.Fatalf("failed to write entry to raft log: %s", err)
//		}
//	}
//	if err := bs.Close(); err != nil {
//		t.Fatalf("failed to close bolt db: %s", err)
//	}
//
//	l, err := NewLog(path)
//	if err != nil {
//		t.Fatalf("failed to create new log: %s", err)
//	}
//
//	fi, err := l.FirstIndex()
//	if err != nil {
//		t.Fatalf("failed to get first index: %s", err)
//	}
//	if fi != 1 {
//		t.Fatalf("got wrong value for first index of empty log: %d", fi)
//	}
//
//	li, err := l.LastIndex()
//	if err != nil {
//		t.Fatalf("failed to get last index: %s", err)
//	}
//	if li != 4 {
//		t.Fatalf("got wrong value for last index of not empty log: %d", li)
//	}
//
//	lci, err := l.LastCommandIndex()
//	if err != nil {
//		t.Fatalf("failed to get last command index: %s", err)
//	}
//	if lci != 4 {
//		t.Fatalf("got wrong value for last command index of not empty log: %d", lci)
//	}
//
//	if err := l.Close(); err != nil {
//		t.Fatalf("failed to close log: %s", err)
//	}
//
//	// Delete an entry, recheck index functionality.
//	bs, err = raftboltdb.NewBoltStore(path)
//	if err != nil {
//		t.Fatalf("failed to re-open bolt store: %s", err)
//	}
//	if err := bs.DeleteRange(1, 1); err != nil {
//		t.Fatalf("failed to delete range: %s", err)
//	}
//	if err := bs.Close(); err != nil {
//		t.Fatalf("failed to close bolt db: %s", err)
//	}
//
//	l, err = NewLog(path)
//	if err != nil {
//		t.Fatalf("failed to create new log: %s", err)
//	}
//
//	fi, err = l.FirstIndex()
//	if err != nil {
//		t.Fatalf("failed to get first index: %s", err)
//	}
//	if fi != 2 {
//		t.Fatalf("got wrong value for first index of empty log: %d", fi)
//	}
//
//	li, err = l.LastIndex()
//	if err != nil {
//		t.Fatalf("failed to get last index: %s", err)
//	}
//	if li != 4 {
//		t.Fatalf("got wrong value for last index of empty log: %d", li)
//	}
//
//	fi, li, err = l.Indexes()
//	if err != nil {
//		t.Fatalf("failed to get indexes: %s", err)
//	}
//	if fi != 2 {
//		t.Fatalf("got wrong value for first index of empty log: %d", fi)
//	}
//	if li != 4 {
//		t.Fatalf("got wrong value for last index of empty log: %d", li)
//	}
//
//	if err := l.Close(); err != nil {
//		t.Fatalf("failed to close log: %s", err)
//	}
//}
//
//func Test_LogLastCommandIndexNotExist(t *testing.T) {
//	path := mustTempFile()
//	defer os.Remove(path)
//
//	// JSON some entries directory to the BoltDB Raft store.
//	bs, err := raftboltdb.NewBoltStore(path)
//	if err != nil {
//		t.Fatalf("failed to create bolt store: %s", err)
//	}
//	for i := 4; i > 0; i-- {
//		if err := bs.StoreLog(&raft.Log{
//			Index: uint64(i),
//			Type:  raft.LogNoop,
//		}); err != nil {
//			t.Fatalf("failed to write entry to raft log: %s", err)
//		}
//	}
//	if err := bs.Close(); err != nil {
//		t.Fatalf("failed to close bolt db: %s", err)
//	}
//
//	l, err := NewLog(path)
//	if err != nil {
//		t.Fatalf("failed to create new log: %s", err)
//	}
//
//	fi, err := l.FirstIndex()
//	if err != nil {
//		t.Fatalf("failed to get first index: %s", err)
//	}
//	if fi != 1 {
//		t.Fatalf("got wrong value for first index of empty log: %d", fi)
//	}
//
//	li, err := l.LastIndex()
//	if err != nil {
//		t.Fatalf("failed to get last index: %s", err)
//	}
//	if li != 4 {
//		t.Fatalf("got wrong for last index of not empty log: %d", li)
//	}
//
//	lci, err := l.LastCommandIndex()
//	if err != nil {
//		t.Fatalf("failed to get last command index: %s", err)
//	}
//	if lci != 0 {
//		t.Fatalf("got wrong value for last command index of not empty log: %d", lci)
//	}
//
//	if err := l.Close(); err != nil {
//		t.Fatalf("failed to close log: %s", err)
//	}
//
//	// Delete first log.
//	bs, err = raftboltdb.NewBoltStore(path)
//	if err != nil {
//		t.Fatalf("failed to re-open bolt store: %s", err)
//	}
//	if err := bs.DeleteRange(1, 1); err != nil {
//		t.Fatalf("failed to delete range: %s", err)
//	}
//	if err := bs.Close(); err != nil {
//		t.Fatalf("failed to close bolt db: %s", err)
//	}
//
//	l, err = NewLog(path)
//	if err != nil {
//		t.Fatalf("failed to create new log: %s", err)
//	}
//
//	lci, err = l.LastCommandIndex()
//	if err != nil {
//		t.Fatalf("failed to get last command index: %s", err)
//	}
//	if lci != 0 {
//		t.Fatalf("got wrong value for last command index of not empty log: %d", lci)
//	}
//}

// mustTempFile returns a path to a temporary file in directory dir. It is up to the
// caller to remove the file once it is no longer needed.
func mustTempFile() string {
	tmpfile, err := ioutil.TempFile("", "casbin-mesh-db-test")
	if err != nil {
		panic(err.Error())
	}
	tmpfile.Close()
	return tmpfile.Name()
}
