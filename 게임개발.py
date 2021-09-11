# 케릭터가 움직이는 N x M 크기의 직사각형으로 칸은 육지 또는 바다이다.
# 케릭터는 동서남북 중 한 곳을 바라본다.
# 맵의 각 칸은 (A,B)로 나타낼 수 있고 A는 북쪽으로 떨어진 칸의 개수, B는 서쪽으로부터 떨어진 칸의 개수이다.
# 캐릭터는 좌우상하 움직일 수 있으며, 바다로 되어 있는 공간에는 갈 수 없다.
# 케릭터의 움직임을 설정하는 매뉴얼은 다음과 같다.
# 현재 위치에서 현재 방향을 기준으로 왼쪽방향(반 시계방향으로 90도 회전 방향)부터 차례대로 갈 곳을 정한다.
# 케릭터의 바로 왼쪽 방향에 아직 가보지 않은 칸이 존재한다면, 왼존 방향으로 회전한 다음 한 칸 전진한다. 만약 존재 하지 않는 다면 다시 회전 한다.
# 만약 네 방향 모두 이미 가본 칸이거나 바다로 되어 있는 칸인 경우에는 바라보는 방향을 유치한 채로 한칸 뒤로 가고  1 단계로 돌아간다. 이때 뒤쪽 방향이 바다인 칸이라
# 뒤로 갈 수 없는 경우에는 움직임을 멈춘다. 
# 매뉴얼에 따라 케릭터를 이동시킨 뒤에 캐릭터가 방문한 칸의 수를 출력하는 프로그램을 만드시오.
# 첫째 줄은  맵 크기
# 둘째 줄은 캐릭터 위치와 바라보는 곳
# 셋째 줄부터는 육지인지 바다인지 보여준다. 맵의 외곽은 항상 바다이다.  0은 육지 1은 바다
# 처음 케릭터는 항상 육지에 있는다. 

def set_map(map_size):
    map_list = map_size[2:]

    return map_list

def change_cycle_numbers(z):
    if z == 0:
       return (-1, 0)
    elif z == 3:
       return (0, -1)
    elif z == 2:
        return (1, 0)
    else: 
        return (0, 1)

def change_back_potion(x, y, z):
    if z == 0:
       return (x -1, y)
    elif z == 3:
       return (x, y -1)
    elif z == 2:
        return (x + 1, y)
    else: 
        return (x, y + 1) 

def view_moving_point(x, y , z):
    dx, dy = map(int, change_cycle_numbers(z)) 

    return [x+ dx, y + dy]

def chack_moving_point(x, y ,z , map_list, next_postion, count, mark_cash_list, chack_cash_list):
    next_point = map_list[next_postion[0]][next_postion[1]]

    if next_point == 1 and count  != 4:
      chack_cash_list.append(1)
      count += 1
      turn_view = view_moving_point(x, y, z)
      chack_moving_point(x, y , z, map_list, turn_view, count, mark_cash_list, chack_cash_list)
    elif next_postion in mark_cash_list and count != 4:
      chack_cash_list.append(2)
      count += 1
      turn_view = view_moving_point(x, y , z)
      chack_moving_point(x, y , z, map_list, turn_view, count, mark_cash_list, chack_cash_list)
    elif count == 4 and chack_cash_list.count(1) == 4:
       return list(mark_cash_list)
    elif count == 4 and chack_cash_list.count(1) != 4:
       back_x, back_y = map(int, change_back_potion(x, y, z))
       count = 0
       chack_moving_point(x, y , z, map_list, turn_view, count, mark_cash_list, chack_cash_list)
    elif count == 4:
        return []
    else:
        count = 0
        chack_cash_list = []
        mark_cash_list.append((x, y))
        next_postion = view_moving_point(x, y ,z)  
        chack_moving_point(x, y , z, map_list, next_postion, count, mark_cash_list, chack_cash_list)

def main(input, mark_cash_list):
  x, y, z = map(int, input[1])
  chack_cash_list = []
  count = 0
  map_list = set_map(input)
  next_postion = view_moving_point(x, y ,z)
  result = chack_moving_point(x, y, z, map_list, next_postion, count, mark_cash_list, chack_cash_list)

  if result == None:
    return 0
  return len(result)

mark_cash_list = []
input = [(4, 4), (2, 2, 0), (1, 1 ,1 ,1), (1, 0, 0 , 1), (1, 1, 0, 1), (1, 1, 1, 1)]
print(main(input, mark_cash_list))

#답안 

def result():
    n, m = map(int, input().split())
    d= [ [0] * m for _ in range(n) ]
    x, y , direction = map(int, input().split())
    d[x][y] = 1

    array = []
    for i in range(n):
        array.append(list(map(int, input().split()))) 

    dx = [-1 , 0 , 1, 0]
    dy = [0, 1, 0 -1]

    def turn_left():
        global direction
        direction -= 1
        if direction == -1:
              direction = 3
        
  count = 1
  trun_time = 0
  while True:
      trun_left()
      nx = x + dx[direction]
      ny = y + dy[direction]

      if d[nx][ny] == 0 and array[nx][ny] == 0:
          d[nx][ny] = 1
          x = nx
          y = ny
          count += 1
          turn_time = 0
          continue
      else:
          turn_time += 1
      if turn_time == 4:
          nx = x -dx[direction]
          ny = y - dy[direction]

          if array[nx][ny] ==0:
              x = nx
              y = ny
          else:
              break
          turn_time = 0

      