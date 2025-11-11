function reverseVowels(s: string): string {
    if (s.length < 2)
        return s;
    
    const VOWELS = new Set(['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U']);
    let arr = s.split("");
    let left = 0;
    let right = s.length-1;

    while (left < right) {
        while (left < right && !VOWELS.has(arr[left]))
            left++;
        
        while (left < right && !VOWELS.has(arr[right]))
            right--;
        
        if (left < right) {
            [arr[left], arr[right]] = [arr[right], arr[left]];
            left++;
            right--;
        }
    }

    return arr.join("");
};

function isVowel(l: string): boolean {
    return "aeiou".includes(l) || "AEIOU".includes(l);
}
