import collections

global anagrams

def groupAnagrams(strs):

    for str in strs:
        str.sort()
        list(anagrams).append(str)
    print(wrods)

    # 애너


def main(strs):
    groupAnagrams(strs)


strs = ['eat', 'tea', 'tan', 'ate']
main(strs) 