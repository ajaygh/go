#include <iostream>
#include <stack>
#include <vector>

using namespace std;

long sumAllSubstr(string);
int sumPos(string str, int pos);

long rem = 1000000007;

int main() {
	string n;
	cin >> n;
	cout << sumAllSubstr(n) << endl;	
	return 0;
} 

long sumAllSubstr(string nstr){
	//store all digits into stack

	int total = 0;
	for(int i = 0; i < nstr.length(); i++){
		total += sumPos(nstr, i);
		total %= rem;
	}
	return total;
}

int sumPos(string str, int pos){
	int curr = 0;
	int total = 0;
	for (int j = pos; j < str.length(); j++){
		curr = curr*10 + str[j]-'0';
		curr %= rem;
		total += curr;
		total %= rem;
	}	
	return total;
}

