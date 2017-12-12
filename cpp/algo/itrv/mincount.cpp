#include <iostream>
#include <map>
#include <limits.h> 

using namespace std;

int main() {
	int n;
	cin >> n;
	int tmp;
	map<int, int> mapNum;
	for(int i = 0; i < n; i++){
		cin >> tmp;
		mapNum[tmp]++;
	}
	int min = INT_MAX;
	for(auto el: mapNum){
		if(min > el.second)
			min = el.second;
	}
	cout << min << endl;
	return 0;
} 
