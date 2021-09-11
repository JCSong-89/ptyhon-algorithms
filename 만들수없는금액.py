# 각 동전으로 만들 수 없는 최소의 금액을 제출

def switch_number(test_list):
    test_list.sort();
    new_list = []
    big_sum = test_list[0] + test_list[1]
    length = len(test_list)
    coins = [2, 3, 4, 5, 6 ,7, 8, 9]

    if 1 not in test_list:
        return 1
    
    for i in range(length):
      big_sum = test_list[i]
      for j in range(length - 1):
          if i == j +1:
              continue
          sum_number = test_list[i] + test_list[j + 1]
          big_sum += test_list[j + 1] 

          if sum_number not in new_list:
              new_list.append(sum_number)

          if test_list[i] not in new_list:
              new_list.append(test_list[i])

          if big_sum not in new_list:
              new_list.append(big_sum)

    new_list.sort()
    n = len(new_list)

    for i in range(n):
        if  i + 1 != n and new_list[i] + 1 != new_list[i + 1]:
            return new_list[i] + 1
        if i + 1 == n:
            return new_list[-1] + 1

test = [3, 2, 1,1, 9]
print(switch_number(test))