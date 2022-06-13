package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	type test struct {
		Input
		Output
	}

	tests := []*test{
		{Input{Doc{"喜欢空气", 2}}, Output{Doc{"喜欢空气", 2}}},
		{Input{Doc{"喜欢空气的自由", 0}}, Output{Doc{"喜欢空气的自由", 0}}},
		{Input{Doc{"喜欢自由的空气", 1}}, Output{Doc{"喜欢自由的空气", 1}}},
		{Input{Doc{"自由空气", 3}}, Output{Doc{"自由空气", 3}}},
	}
	// t.Log(tests)

	tree := Constructor()
	for _, test := range tests {
		tree.Insert(test.Input.Doc.Text, test.Input.Doc.Id)
	}
	for _, test := range tests {
		otpt := tree.Search(test.Output.Doc.Text)
		if otpt == nil {
			t.Log("cannot find: ", test)
			continue
		}
		if otpt.Doc.Id != test.Input.Id || otpt.Doc.Text != test.Input.Text {
			t.Error("搜索的结果和插入的数据不一致")
		}
	}

	// res := tree.Search("喜欢空气")
	// t.Log(res)

	// res = tree.Search("喜欢空气的自由")
	// t.Log(res)

	// res = tree.Search("喜欢自由的空气")
	// t.Log(res)

	otpts := tree.StartsWith("喜欢")
	// for _, op := range otpts {
	// 	t.Log(op.Id, op.Text)
	// }
	if len(otpts) != 3 {
		t.Error("搜索结果应该为3")
	}

}
