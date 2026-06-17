type WordDictionary struct {
    children [26]*WordDictionary
    isEnd bool
}


func Constructor() WordDictionary {
    return WordDictionary{}
}


func (this *WordDictionary) AddWord(word string)  {
    node := this
    for _, ch := range word {
        idx := ch - 'a'
        if node.children[idx] == nil {
            node.children[idx] = &WordDictionary{}
        }
        node = node.children[idx]
    }
    node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
    return this.dfs(word, 0)
}

func (this *WordDictionary) dfs(word string, i int) bool {
    if i == len(word) {
        return this.isEnd
    }
    ch := word[i]
    if ch == '.' {
        for _, child := range this.children {
            if child != nil && child.dfs(word, i+1) {
                return true
            }
        }
        return false
    }

    idx := ch - 'a'
    if this.children[idx] == nil {
        return false
    }
    return this.children[idx].dfs(word, i+1)
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
