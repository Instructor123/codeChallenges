#include <stdlib.h>
#include <stdio.h>
#include <string.h>

char firstNotRepeatingCharacter(char *s){
    
    int arraySize = strlen(s);
    int alphabet[26] = {0};
    int alphabetOccurance[26] = {0};
    int ASCII_OFFSET = 97;
    int i = 0;
    int j = 0;
    int value = 0;
    int position = 28;
    int temp = 0;
    char retValue = '_';

    for( i = 0; i < arraySize; ++i ){
        value = s[i] - ASCII_OFFSET;
        alphabet[value] += 1;
        if( '\0' == alphabetOccurance[j] ){
            alphabetOccurance[j] = s[i];
            j++;
        }
    }

    //Don't know what test case this fails, but it's frustrating. There must be a mistake
    //in the alphabetOccurance portion.
    // for( i = 0; i < j; ++i){
    //     temp = alphabetOccurance[i] - ASCII_OFFSET;
    //     if( 1 == alphabet[temp] ){
    //         retValue = alphabetOccurance[i];
    //         break;
    //     }
    // }
    
    for( i = 0; i < arraySize; ++i ){
        value = s[i] - ASCII_OFFSET;
        if( alphabet[value] == 1 ){
            retValue = s[i];
            break;
        }
    }


    printf("%c\n", retValue);

}

int main(void){

    char *name = (char*)malloc(sizeof(char)); 
    name[0] = 'a';
    // name[1] = 'b';
    // name[2] = 'a';
    // name[3] = 'c';
    // name[4] = 'a';
    // name[5] = 'b';
    // name[6] = 'a';
    // name[7] = 'a';
    // name[8] = 'b';
    // name[9] = 'a';
    // name[10] = 'c';
    // name[11] = 'a';
    // name[12] = 'b';
    // name[13] = 'a';

    firstNotRepeatingCharacter(name);

    return 0;
}