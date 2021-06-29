#!/usr/bin/python3

def firstDuplicate(a):
    myDict = dict()

    for x in range(len(a)):
        value = a[x]
        if value in myDict:
            myDict[value].append(x)
        else:
            myDict[value] = list()
    
    myMin = pow(10,5)
    for x in myDict.keys():
        if myDict[x]:
            temp = min(myDict[x])
            if myMin > temp:
                myMin = temp
    
    if myMin == pow(10,5):
        return -1
    else:
        return a[myMin]

    

if __name__ == "__main__":
    myList = [2,2]
    firstDuplicate(myList)