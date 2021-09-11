# 그룹1에는 1, 2, 3을 넣고 2 그룹은 2인 두명을 넣을 수 있다.  최대 생성 그룹은 몇인가?

def chack_member_condition(list):
  count = 0
  first_group = []
  last_group = []

  for i in range(len(list)):
    if list[i] == 1 and 1 not in first_group or list[i] == 2 and 2 not in first_group or list[i] ==3 and 3 not in first_group:
        first_group.append(list[i])
    elif list[i] == 2 and 2 > len(last_group):
        last_group.append(list[i])
    else:
        continue

  if 3 == len(first_group):
      count += 1
  if 2 == len(last_group):
      count += 1
    
  return count
  
test = [2, 3 , 1, 2]
print(chack_member_condition(test))
