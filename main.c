#include <stdio.h>

int main(){
	char nombre[100];

	printf("Hello World\n");
	printf("Nombre: \n");
	fgets(nombre,sizeof(nombre),stdin);
	printf("Hola %s\n",nombre);
	return 0;
}
