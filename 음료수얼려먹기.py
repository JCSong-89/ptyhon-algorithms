# N X M 크기의 얼음 틀이 있다. 구멍이 뚫려 있는 부분은 0, 칸막이가 존재하는 부분은 1로 표시된다. 구멍이 뚫려 있는 부분끼리 상, 하, 좌 , 우로
# 붙어있는 경우 서로 연결되어 있는 것으로 간주한다. 이때 얼음 틀의 모양이 주어졌을 때 생성되는 총 아이스크림 개수를 구하는 프로그램을 작성하시오

test = [(0, 0, 1, 1, 0),(0,0,0,1,1),(1,1,1,1,1), (0,0,0,0,0)]

def chack_zero(list):
    count = 0
    for i in range(len(list)):
        for j in range(len(list[i])):
            if list[i][j] == 1:
                continue
            elif list[i][j] == 0 and i != 0 and list[i - 1][j] == 0 or list[i][j] == 0 and j != 0 and list[i] [j - 1] == 0:
                continue
            else:
                count += 1
    return count   

print(chack_zero(test))

# 답안예시

def dfs(x, y):
    if x <= - 1 or x >= n or y <= - 1 or y >= m:
        return False
    if graph[x][y] == 0:
          graph[x][y] = 1
          dfs(x, -1, y)
          dfs(x, y -1)
          dfs(x + 1, y)
          dfs(x, y + 1)
          return True
    return False

result = 0
for i in range(len(test)):
    for j in range(len(test[i])):
        if dfs(i, j) == True:
              result += 1

