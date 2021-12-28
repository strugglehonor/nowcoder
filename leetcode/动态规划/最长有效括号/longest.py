class Solution:
    def longestValidParentheses(self, s: str) -> int:
        # 维护栈中的栈底元素为最后一个没有被匹配的右括号的索引
        stack = [-1] 
        n = len(s)
        ans = 0
        for i in range(n):
            if s[i] == "(":
                stack.append(i)
            else:
                stack.pop()
                if len(stack) == 0:
                    stack.append(i)
                else:
                    ans = max(ans, i-stack[len(stack)-1])
        return ans