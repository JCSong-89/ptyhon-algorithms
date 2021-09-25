import collections
import re

def mostCommonWord(self, paragraph, banned):
    words = [word for word in re.sub(r'[^/w]', ' ', paragrapqh). lower().split() if word not in banned]
    count = collections.Counter(words)

    return count.most_common(1)[0][0]

def main(paragraph, banned):
    return mostCommonWord(paragraph, banned)

#정규식에서 /w는 단어 문자를 뜻하며, ^는 not을 의미한다. paragrapqh 문자열을 받아 문자가 아닌 것을 공백으로 치환한다.
#banned에 포함되지 않은 문자만 데이터 클렌징을 한다.
#most_common(1)은 가장 흔하게 나오는 단어를 추출하며 키값때문에 [0][0]로 단어를 추출한다.
