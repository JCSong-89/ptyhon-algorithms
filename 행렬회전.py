from collections import deque

def main(matrix):
    new_matrix = []
    frist_list = []
    last_list = []
    q = deque()
    frist_list += matrix[0]
    del frist_list[-1]
    last_list += matrix[-1] 
    del last_list[0]
    left_list = []
    right_list = []
    
    for i in range(len(matrix)):
           right_list.append(matrix[i][0])
           left_list.append(matrix[i][-1])

    del right_list[0]
    del left_list[-1]

    temp = []

    q.append(frist_list)
    frist_list = right_list[::-1]
    q.append(left_list)
    left_list = q.popleft()
    q.append(last_list)
    temp = q.popleft()
    last_list = temp[::-1]
    right_list = q.popleft()
    center =[]

    for i in range(len(matrix)):
        if i != 0 and i != len(matrix) -1:
            center += matrix[i]
            del center[0]
            del center[-1]
    
    frist_list.append(int(left_list[0]))
    q.appendleft(int(right_list[-1]))

    for i in last_list:
        q.append(int(i))

    del left_list[0]
    del right_list[-1]
    
    new_matrix.append(frist_list)
    swtich = True

    i =0
    j = 0
    y = 0
    row_count =0
    col_count= 0

    while swtich == True:

        if i == 0:
            array_x = []
            array_x.append(right_list[i])
            array_x.append(center[j])
            new_matrix.append(array_x)
            row_count += 1
            col_count += 2
            i += 1
            j += 1           
        elif len(right_list) == row_count and len(right_list) < 1:
            array_x = []
            array_x.append(right_list[i])
            array_x.append(center[j])
            new_matrix.append(array_x)
            row_count = 0
            col_count += 2
            i += 1
            j + 1
        elif len(frist_list) -1 == col_count :
            new_matrix[-1].append(left_list[y])
            col_count = 0
            y += 1
        elif new_matrix[-1][-1] == left_list[-1]:
            swtich = False
            break
        else:
            new_matrix[-1].append(center[j])
            j += 1
            col_count += 1
            row_count += 1


    new_matrix.append(list(q))
    
    return new_matrix









      


test = [[1,2 ,3],[4,5,6],[7,8,9]];
print(main(test))