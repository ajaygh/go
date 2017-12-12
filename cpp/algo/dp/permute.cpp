#include <iostream>
#include <vector>

using namespace std;

void permute(string s);
vector<string> permutePos(vector<string> &prev, char ch);

int main() {
	string in;
	cin >> in;
	permute(in);
	cout << endl;

	return 0;
} 

void permute(string s){
	vector<string> res;
	//base case
	string tmp;
	tmp.push_back(s[0]);
	res.push_back(tmp);

	for(int i = 1; i < s.length(); i++){
		res = permutePos(res, s[i]);
	}
	for(auto str: res) cout << str <<" ";
}

vector<string> permutePos(vector<string> &prev, char ch){
	vector<string> res;

	for(auto str: prev){
		for(int i = 0; i < str.length()+1; i++){
			string tmp;
		   	for(int j = 0; j < i; j++)
				tmp.push_back(str[j]);
			tmp.push_back(ch);
			for(int j = i; j < str.length();j++)
				tmp.push_back(str[j]);
			res.push_back(tmp);
		}
	}
	return res;
}
