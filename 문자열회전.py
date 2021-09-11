def mian(string1, string2):
    sort_string2 = string2.sorted()
    sort_string1 = string1.sorted()

    if sort_string1 != sort_string2:
        return False
    
    q_string1 = []
    q_string2 = []
    n =  len(string1)

    for i in range(n):
        if i != n:
          q_string1.append([string1[i], string1[i+1]])
          q_string2.append([string2[i], string2[i+1]])
        else:
          q_string1.append([string1[-1], string1[0]])
          q_string2.append([string2[-1], string2[0]])

    q_string1.sort()
    q_string2.sort()

    for i in range(n):
      if q_string1[i] != q_string2[i]:
          return False

    return True    