def main(string):
    new_string = ''
    string = string.split()
    
    for i in string:
        new_string += i

    new_string = new_string.lower()
    reverse_list = list(new_string)
    reverse_list.reverse()
    reverse_string = ''

    for i in reverse_list:
        reverse_string += i

    if reverse_string == new_string:
        return True
    else: 
        return False

print(main('taco cat'))
