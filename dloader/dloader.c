#include <dlfcn.h>
#include <stdio.h>


void close(void *h) {
    dlclose(h);
}

void * dloader_open(char *path, char **err) {
    void *h = dlopen(path, RTLD_LAZY);

    if (h == NULL)
        *err = dlerror();

    return h;
}

void * dloader_lookup(void* h, char *symbol, char **err) {
    void *f = dlsym(h, symbol);

    if (f == NULL)
        *err = dlerror();

    return f;
}