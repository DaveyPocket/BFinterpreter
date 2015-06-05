#include <stdio.h>


int forwardTrack(int n, int count, unsigned char* loc);
int backTrack(int n, int count, unsigned char* loc);

int main(){
	//char* fDir;
	//printf("Program directory: ");
	//scanf("%s", fDir);
	FILE* fName = fopen("PROGRAM DIRECTORY HERE :DD", "r");
	printf("opened");
	unsigned char prog[500000];
	char temp;
	int i = 0;
	while(fscanf(fName, "%c", &temp) != EOF){
		switch (temp){
			case '>':
			case '<':
			case '.':
			case ',':
			case '+':
			case '-':
			case '[':
			case ']':
				prog[i] = temp;	
				i++;
				break;
		}
	}
	fclose(fName);
	prog[i] = 0;
	
	static unsigned char mem[30000];
	unsigned int mPtr = 0;

	for (i = 0; prog[i] != 0; i++){
		switch (prog[i]){
			case '>':
				mPtr++;
				break;
			case '<':
				mPtr--;
				break;
			case '.':
				printf("%c", mem[mPtr]);
				break;
			case ',':
				scanf("%c", &mem[mPtr]);
				break;
			case '+':
				mem[mPtr]++;
				break;
			case '-':
				mem[mPtr]--;
				break;
			case '[':
				if (mem[mPtr] == 0){
					i += forwardTrack(i + 1, 0, prog);
				}	
				break;
			case ']':
				if (mem[mPtr] != 0){
					i -= backTrack(i - 1, 0, prog);
				}
				break;
		}
	}

	return 0;
}


int backTrack(int n, int count, unsigned char* loc){
	if (count < 0){
		return 0;
	}
	if (loc[n] == 91){		// [
		count--;
	}else if (loc[n] == 93){		// ]
		count++;
	}
	return backTrack(n - 1, count, loc) + 1;
	
}

int forwardTrack(int n, int count, unsigned char* loc){
	if (count < 0){
		return 0;
	}
	if (loc[n] == 91){		// [
		count++;
	}else if (loc[n] == 93){		// ]
		count--;
	}
	return forwardTrack(n + 1, count, loc) + 1;
	
}

