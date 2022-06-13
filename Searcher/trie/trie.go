/*
	Author: YuuLeii
	Date: 2022.6.13 20:58
*/

/*
	Tips: 可给Trie结点增加一个count字段记录句子出现的次数
*/
package trie

type Trie struct {
	id       int
	exist    bool
	text     string // 每个结尾的结点存储的是整个句子
	children map[rune]*Trie
}
type Doc struct {
	Text string
	Id   int
}
type Output struct {
	Doc
}
type Input struct {
	Doc
}

/* 初始化树 */
func Constructor() Trie {
	root := new(Trie)
	root.children = make(map[rune]*Trie)
	root.exist = false
	return *root
}

/* 添加数据 */
func (tree *Trie) Insert(word string, id int) {
	word_len := len([]rune(word))
	i := 1
	for _, c := range word {
		if tree.children[c] == nil {
			node := &Trie{}
			node.children = make(map[rune]*Trie)
			if i == word_len {
				node.exist = true
				node.id = id
				node.text = word
			}
			tree.children[c] = node
		}
		tree = tree.children[c]
		i++
	}
}

/* 完全匹配 */
func (tree *Trie) Search(word string) *Output {
	ret := &Output{Doc{Id: 0, Text: ""}}
	for _, c := range word {
		if tree.children[c] == nil {
			return nil
		}
		tree = tree.children[c]
	}
	if tree.exist {
		ret = &Output{Doc{Id: tree.id, Text: tree.text}}
		return ret
	}
	return nil
}

/* 找最后的内容 */
func recursion(data map[rune]*Trie, rets []*Output) []*Output {
	for _, v := range data {
		if v.exist {
			rets = append(rets, &Output{Doc{Id: v.id, Text: v.text}})
		}
		rets = recursion(v.children, rets)
	}
	return rets
}

/* 前缀匹配 */
func (tree *Trie) StartsWith(word string) (ret []*Output) {
	cur := tree
	for _, c := range word {
		if cur.children[c] == nil {
			return nil
		}
		cur = cur.children[c]
	}
	if cur.exist {
		ret = append(ret, &Output{Doc{Id: cur.id, Text: cur.text}})
	}
	if len(cur.children) > 0 {
		ret = recursion(cur.children, ret)
	}
	return ret
}
