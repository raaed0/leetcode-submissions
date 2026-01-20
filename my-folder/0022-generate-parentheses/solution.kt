class Solution {
    fun generateParenthesis(n: Int): List<String> {
        val res = mutableListOf<String>()

        fun valid(s: String): Boolean {
            var open = 0
            for (c in s) {
                if (c == '(') open++ else open--
                if (open < 0) return false
            }
            return open == 0
        }

        fun dfs(s: String) {
            if (s.length == n * 2) {
                if (valid(s)) {
                    res.add(s)
                }
                return
            }

            dfs(s + "(")
            dfs(s + ")")
        }

        dfs("")
        return res
    }
}
