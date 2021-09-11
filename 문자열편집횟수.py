def main(string1, string2):
    string1 = string1.lower()
    string2 = string2.lower()

    array_x = []
    array_y = []

    for i in string1:
        array_x.append(i)
    for i in string2:
        array_y.append(i)

    xn = len(array_x)
    yn = len(array_y)

    count = 0

    if xn >= yn:
        for i in range(yn):
            if array_y[i] != array_x[i]:
                count += 1
    else:
        for i in range(xn):      
            if array_x[i] != array_y[i]:
                count += 1

    if count > 1:
        return False
    else:
        return True

print(main('pale', 'brak'))
        