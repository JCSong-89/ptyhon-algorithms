def main(test):
    q = []
    n = len(test)
    m = len(test[0])
    for i in range(n):
        if 0 in test[i]:
            for j in range(m):
                if 0 == test[i][j]:
                    q.append([i, j])

    for i in q:
        x, y = map(int, i)
        test[x] = []
        test[x] = [0] * m
        for j in range(n):
            test[j][y] = 0
    
    return test


