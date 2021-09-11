from collections import Counter

def main(string):
    dic = Counter(string)
    new_string = ''
    for i in dic.keys():
      new_string += i + str(dic.get(i))
        
    if len(new_string) > len(string):
        return string
    return new_string

print(main('ac'))