#include<stdlib.h>

void* c_malloc(int size){
    void* p = malloc(size);
    return p;
}

void c_free(void* p) {
    free(p);
}
