#include <iostream>
#include <vector>

typedif vector<char> vc;
typedif vector<vector<char>> vcc;

#define for(i, n) for(int i = 0; i < n; i++)

using namespace std;

int main() {
	int l, b;
	cin >> l >> b;
	vcc rec(l, vc(b));
	for(i, l){
		for(j, b){
			cin >> rec[i][j];
		}
	}

	cout << maxPeri(rec, 0, 0, l, b) << endl;
	return 0;
} 

long maxPeri(vcc& rec, int x, int y, int l, int b){
	//check marshy area
	if(chkMarshy(rec, x, y, l, b)){
		
	}
}

bool chkMarshy(vcc& rec, int x, int y, int l, int b){
	//check marshy area
	for(int j = y; j < l; j++){
		if(rec[x][j] == 'X' || rec[x+b-1] == 'X'){
			return true;
		}
	}
	for(int i = x; j < b; i++){
		if(rec[i][y] == 'X' || rec[i+l-1] == 'X'){
			return true;
		}
	}
}
