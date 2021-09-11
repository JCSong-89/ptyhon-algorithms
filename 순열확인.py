def main(string1, string2):
    string.lower()
    string2.lower()

    x = list(string1)
    x.reverse()
    
    str = ''

    for i in x:
        str += i

    if str == string2:
        return True
    return False

print(main('god', 'dogs'))