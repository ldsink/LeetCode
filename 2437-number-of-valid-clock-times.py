class Solution:
    def countTime(self, time: str) -> int:
        def f(s: str, m: int):
            result = 0
            for i in range(m):
                a = s[0] == "?" or int(s[0]) == i // 10
                b = s[1] == "?" or int(s[1]) == i % 10
                result += a and b
            return result

        return f(time[:2], 24) * f(time[3:], 60)
