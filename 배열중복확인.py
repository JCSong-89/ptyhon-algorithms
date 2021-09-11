#문자열 중복 확인
from collections import Counter

def main(strings):
    count = Counter(list(strings))
    values = list(count.values())

    if values[-1] > 1:
        return False
    return True

a = 'abcdbcdcd'
print(main(a))