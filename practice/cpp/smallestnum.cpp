//find the smallest number repeating k number of times
#include <iostream>
#include <vector>
#include <limits.h> 
#include <unordered_map>

using namespace std;

int main() {
	int n;
	cin >> n;
	vector<int> id(n);
	for (int i = 0; i < n; i++) cin >> id[i];

	int min;
	cin >> min;

	unordered_map<int, int> numCnt;
	for (int i = 0; i < n; i++){
		if (numCnt.find(id[i]) != numCnt.end()) numCnt[id[i]]++;
		else numCnt[id[i]] = 1;
	}

	int minnum = INT_MAX;
	for (auto v: numCnt){
		if (v.second == min && v.first < minnum) minnum = v.first;
	}	
	cout << minnum << endl; 
	return 0;
}
