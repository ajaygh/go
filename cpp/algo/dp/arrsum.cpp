#include <iostream>
#include <vector>

using namespace std;

long sumAll(vector<int> &va);

int main() {
	int n;
	cin >> n;
	vector<int> va(n);
	for (int i = 0; i < n; i++) cin >> va[i];
	auto sum = sumAll(va);
	cout << sum << endl;	
	return 0;
} 

long sumAll(vector<int> &va){
	long ressum = 0;
	int rem = 1000000007;
	vector<int> res(va.size(),0);
	res[0] = va[0]; //base case
	ressum += res[0];
	for (int i = 1; i < va.size(); i++){
		res[i] = i*va[i];
		for(int j = i-1; j >= 0; j--){
			long tmp = 0;
		    for(int k = j; k<= i; k++) tmp += va[k];
	   		res[i] += tmp*(i-j+1);	   
			res[i] %= rem;
			for(int k = 0; k < j; k++){
				res[i] += va[k];
				res[i] %= rem;
			}
		}
		ressum += res[i];
		ressum %= rem;
	}
	return ressum;
}
