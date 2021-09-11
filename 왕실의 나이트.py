# 행복한 왕실 정원은 체스판과 같은 8n8 좌표 평면이다. 
# 나이트는 말을 타고 있어 L자 형태로만 이동할 수 있으며, 정원 밖으로는 나갈 수 없다. 
# 나이트는 특정한 위치에서 다음과 같은 2가지 경우로 이동할 수 있다.
# 수평으로 두칸 이동한 뒤 수직으로 한칸 이동하기
# 수직으로 두칸 이동한 뒤에 수평으로 한칸 이동하기
# 이때 행 위치는 1 - 8, 열 위치는 a - h 로 한다.

def change_numbers(alpbet):
    ascll = ord(alpbet)

    if ascll > 41  and ascll < 48:
        return 0
    
    return ascll - 40

def move_up_left(x, y):
    y -= 2
    x += 1
    if y > 1:
        return 0
    elif x < 8:
        return 0
    else:
        return 1  

def move_donw_left(x, y):
    y += 2
    x += 1
    if y < 8:
        return 0
    elif x < 8:
        return 0
    else:
        return 1  

def move_up_right(x, y):
    y -= 2
    x -= 1
    if y > 1:
        return 0
    elif x > 1:
        return 0
    else:
        return 1  

def move_down_right(x, y):
    y += 2
    x -= 1
    if y < 8:
        return 0
    elif x > 8:
        return 0
    else:
        return 1  

def move_left_up(x, y):
    y -= 1
    x += 2
    if y > 1:
        return 0
    elif x < 8:
        return 0
    else:
        return 1  

def move_left_down(x, y):
    y += 1
    x += 2
    if y < 8:
        return 0
    elif x < 8:
        return 0
    else:
        return 1  

def move_right_up(x, y):
    y -= 1
    x -= 2
    if y > 1:
        return 0
    elif x > 1:
        return 0
    else:
        return 1  

def move_right_down(x, y):
    y += 1
    x -= 2
    if y < 8:
        return 0
    elif x > 8:
        return 0
    else:
        return 1  

def move_knight(postion):
    y, x = map(int ,postion)    
    true_list = []
    true_list.append(move_up_left(x, y))
    true_list.append(move_donw_left(x, y))
    true_list.append(move_up_right(x, y))
    true_list.append(move_down_right(x, y))
    true_list.append(move_left_down(x, y))
    true_list.append(move_left_up(x, y))
    true_list.append(move_right_up(x, y))
    true_list.append(move_right_down(x, y))
    
    return true_list.count(True)

def main(input):
    number = input[1];
    alpbet = input[0];
    change_number = change_numbers(alpbet);
    
    if change_number == 0:
      return change_number

    postion = [number , change_number]
    result = move_knight(postion)
    return result


input = 'a1'
print(main(input))


# 답안예시
def book_result(input):
    row = int(input[1]);
    col = int(ord(input[0])) - int(ord('a')) + 1;
    
    steps = [
      (-2, -1), (-1, -2), (1, -2), (2 -1),
      (2, 1), (1, 2), (-1, 2), (-2, 1)
    ]

    result = 0
    for step in steps:
        next_row = row + step[0]
        next_col = col + step[1]

        if next_row >= 1 and next_row <= 8 and next_col >= 1 and next_col <= 8:
            result += 1

    return result