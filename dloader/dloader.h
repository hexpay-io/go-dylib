void close(void *h);

void * dloader_open(char *path, char **err);
void * dloader_lookup(void* h, char *symbol, char **err);