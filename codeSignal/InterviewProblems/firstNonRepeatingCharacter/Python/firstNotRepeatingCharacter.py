#!/usr/bin/python3

def firstNotRepeatingCharacter(s):
    for x in range(len(s)):
        #count should work, but apparently they are being tricky somehow.
        '''
        if 1 == s.count(s[x]):
            return s[x]
        '''
        if s.index(s[x]) == s.rindex(s[x]):
            return s[x]
    return "_"


if __name__ == "__main__":
    print(firstNotRepeatingCharacter("abacabad"))