#include <stdlib.h>
#include <stdio.h>

// Arrays are already defined with this interface:
// typedef struct arr_##name {
//   int size;
//   type *arr;
// } arr_##name;
//
// arr_##name alloc_arr_##name(int len) {
//   arr_##name a = {len, len > 0 ? malloc(sizeof(type) * len) : NULL};
//   return a;
// }
//
//
int firstDuplicate(arr_integer a) {
    int maxNumber = pow(10, 5);
    int *dupArray;
    int i = 0;
    int value = 0;
    int smallestValue = maxNumber;
    
    //the trick was off by ones.
    dupArray = (int*)malloc(sizeof(int)*(a.size+1));
    memset(dupArray, -1, sizeof(int)*(a.size+1));
    
    for( i = 0; i < a.size; ++i ){
        value = a.arr[i];
        if( -1 == dupArray[value] ){
            dupArray[value] = 0;
        } else if( 0 == dupArray[value] ){
            dupArray[value] = i;
            if( i < smallestValue ){
                smallestValue = i;
            }
        }
    }
    
    if( smallestValue == maxNumber ){
        smallestValue = -1;
    } else {
        smallestValue = a.arr[smallestValue];
    }
    return smallestValue;
    
}

int main(void){    

    return 0;
}