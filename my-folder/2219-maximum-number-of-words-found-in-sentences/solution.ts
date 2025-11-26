function mostWordsFound(sentences: string[]): number {
    let max = 0;
    for (let i = 0; i < sentences.length; i++) {
        let l = sentences[i].split(' ').length;
        if (l > max) max = l;
    }
    return max
};
