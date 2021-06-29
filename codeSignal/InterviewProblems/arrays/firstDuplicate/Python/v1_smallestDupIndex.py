#!/usr/bin/python3

def firstDuplicate(a):
    cont = True
    checkIter = 0
    nextIter = 0
    checkValue = 0
    nextValue = 0
    myNumbers = dict()
    
    checkValue = a[checkIter]
    
    while cont:
        checkIter += 1
        if checkIter < len(a):
            nextValue = a[checkIter]
            if (checkValue == nextValue):
                if checkValue not in myNumbers:
                    myNumbers[checkValue] = checkIter
                else:
                    nextIter += 1
                    checkIter = nextIter
                    checkValue = a[checkIter]        
        else:
            if nextIter != len(a)-1:
                nextIter += 1
                checkIter = nextIter
                checkValue = a[checkIter]
            else:
                cont = False
    
    retValue = -1
    
    if myNumbers:
        smallestValue = pow(10, 5)
        smallestIndex = 0
        for x in myNumbers.keys():
            if myNumbers[x] < smallestValue:
                smallestValue = myNumbers[x]
        retValue = a[smallestValue]
    
    return retValue
    

if __name__ == "__main__":
    temp = [2, 1, 3, 5, 3, 2]
    firstDuplicate(temp)