type Trie struct {
    children [26]*Trie
    isEnd bool
}

type MagicDictionary struct {
    root *Trie
}


func Constructor() MagicDictionary {
    return MagicDictionary{root: &Trie{}}
}


func (this *MagicDictionary) BuildDict(dictionary []string)  {
    for _, word := range dictionary {
        node := this.root
        for _, ch := range word {
            idx := ch - 'a'
            if node.children[idx] == nil {
                node.children[idx] = &Trie{}
            }
            node = node.children[idx]
        }
        node.isEnd = true
    }
}


func (this *MagicDictionary) Search(searchWord string) bool {
    return this.dfs(this.root, searchWord, 0, 0)
}

func (this *MagicDictionary) dfs(node *Trie, word string, i, diff int) bool {
    if i == len(word) {
        return diff == 1 && node.isEnd
    }

    idx := int(word[i] - 'a')

    for c := 0; c < 26; c++ {
        if node.children[c] == nil {
            continue
        }
        newDiff := diff
        if c != idx {
            newDiff++
        }
        if newDiff <= 1 {
            if this.dfs(node.children[c], word, i + 1, newDiff) {
                return true
            }
        }
    }
    return false
}


/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
