#include <iostream>
#include <vector>

using namespace std;

int minSum(vector<vector<int>> price);

int main() {
	int n;
	cin >> n;
	vector<vector<int>> price(n, vector<int>(3));
	for(int i = 0; i < n; i++)
		cin >> price[i][0] >> price[i][1] >> price[i][2];
	cout << minSum(price) << endl;	
	return 0;
} 

int minSum(vector<vector<int>> price){
	for(int i = 1; i < price.size(); i++){
		price[i][0] += min(price[i-1][1], price[i-1][2]);
		price[i][1] += min(price[i-1][0], price[i-1][2]);
		price[i][2] += min(price[i-1][0], price[i-1][1]);
	}

	int n = price.size();
	auto min = price[n-1][0];
	if (min > price[n-1][1])
		min = price[n-1][1];
	if (min > price[n-1][2])
		min = price[n-1][2];
	return min;
}
