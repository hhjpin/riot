// Copyright 2016 ego authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"fmt"
	"os"

	"github.com/hhjpin/riot"
	"github.com/hhjpin/riot/types"
)

var (
	// searcher 是协程安全的
	searcher = riot.Engine{}
)

func initEngine() {
	var path = "./riot-index"

	searcher.Init(types.EngineOpts{
		// Using: 1,
		PinYin: true,
		IndexerOpts: &types.IndexerOpts{
			IndexType: types.DocIdsIndex,
		},
		UseStore:    true,
		StoreFolder: path,
		GseDict:     "../../data/dict/dictionary.txt",
		// StopTokenFile:           "../../riot/data/dict/stop_tokens.txt",
	})
	defer searcher.Close()
	os.MkdirAll(path, 0777)

	text := "在路上, in the way"
	index1 := types.DocData{Content: text, Fields: "在路上"}
	index2 := types.DocData{Content: text}
	index3 := types.DocData{Content: "In the way."}

	searcher.Index("10", index1)
	searcher.Index("11", index2)
	searcher.Index("12", index3)

	// 等待索引刷新完毕
	searcher.Flush()
}

func main() {
	initEngine()

	sea := searcher.Search(types.SearchReq{
		Text: "zl",
		RankOpts: &types.RankOpts{
			OutputOffset: 0,
			MaxOutputs:   100,
		}})

	fmt.Println("search response: ", sea, "; docs = ", sea.Docs)
}
