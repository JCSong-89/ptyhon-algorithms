# 번호 조합 경우의 수
from itertools import combinations   

def main(test_list):
    combinations_list = list(combinations   (test_list, 2))
    return len(combinations_list)

test_list = [1, 3, 2, 3, 2]
print(main(test_list, N))