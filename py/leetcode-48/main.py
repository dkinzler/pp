from typing import List

class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        # rotate image in place
        # one cycle consists of 4 cells
        n = len(matrix)
        b = n-1
        sr, sc = 0, 0
        while b >= 1:
            for i in range(b):
                r1, c1 = sr, sc+i
                r2, c2 = self.rotateCell(r1, c1, n)
                r3, c3 = self.rotateCell(r2, c2, n)
                r4, c4 = self.rotateCell(r3, c3, n)
                matrix[r1][c1], matrix[r2][c2], matrix[r3][c3], matrix[r4][c4] = matrix[r4][c4], matrix[r1][c1], matrix[r2][c2], matrix[r3][c3]


            sr, sc = sr+1, sc+1
            b -= 2

    def rotateCell(self, r: int, c:int, n: int):
        return (c, n-r-1)
