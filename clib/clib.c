#include <stdio.h>


void* Encoder_Interface_init(void* p){
	struct {
		int a;
		int b;
		char* fname;
	} *a = p;
	printf("c Encoder_Interface_init p:%d\n",p);
	printf("Encoder_Interface_init a:%d b:%d fname:%s\n",a->a,a->b,a->fname);
	return (void*)(13);
}
void Encoder_Interface_exit(void* state){
	printf("Encoder_Interface_exit c state:%d\n",state);
}