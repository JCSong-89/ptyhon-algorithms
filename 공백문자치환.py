
def main(string):
    list_string = string.split()
    new_str = ''
    for i in range(len(list_string)):
        if i == 0:
            new_str += list_string[i]
        else:
            new_str += '%20' + list_string[i]
    return new_str


print(main('안녕 하세요'))