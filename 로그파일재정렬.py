#문자열 앞 부분은 식별자.
# 문자 로그보다 숫자 로그가 앞에 온다.
# 식벼자는 순서영향이 없으나 문자가 동일하면 식별자자 순으로 한다.True
# 순자 로그는 입력 순서대로 한다.

test = ["digl 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"]

def main(test):
    int_loglist = []
    str_loglist =[]

    for i in test:
        temp_list = list(i.split())
        del temp_list[0]
        temp_str = ''

        for j in temp_list:
          temp_str += str(j)
        
        if temp_str.isdecimal() == True:
            int_loglist.append(i)
        else:
            str_loglist.append(i)

    int_loglist.sort()
    str_loglist.sort()
    str_loglist += int_loglist

    return str_loglist      

print(main(test))

## 답안

def solution(test):
    letters, digits = []
    for log in test:
        if log.split()[1].isdigit():
            digits.append(log)
        else:
            letters.append(log)
    letters.sort(key=lambda x: (x.split()[1:], x.split()[0]))
    return letters + digits