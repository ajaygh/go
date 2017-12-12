#include <iostream>
#include <vector>
#include <stack> 

using namespace std;

int maxArea(vector<int> &hist);

int main() {
	int n; 
	cin >> n;
	vector<int> arr(n);
	for(int i = 0; i < n; i++)
		cin >> arr[i];
	cout << maxArea(arr) << endl;	
	return 0;
}

int maxArea(vector<int> &hist){
	stack<int> stkHeight, stkPos;
	long area = 0, tmpArea = 0;
	if (hist.size() == 1)
		return int(hist[0]);
	
	stkHeight.push(hist[0]);
	stkPos.push(0);
	for(int i = 1; i < hist.size(); i++){
		if(hist[i] > stkHeight.top()){
			stkHeight.push(hist[i]);
			stkPos.push(i);
		}else if(hist[i] < stkHeight.top()){
			int pos = 0;
			while(hist[i] < stkHeight.top()){
				int ht = stkHeight.top();
				pos = stkPos.top();

				tmpArea = long(ht* (i-pos));
				if (area < tmpArea)
					area = tmpArea;
				stkHeight.pop();

				if(stkHeight.empty()){
					stkHeight.push(hist[i]);
					break;
				}else{
					stkPos.pop();
				}
			 }
		}
	}
	while(!stkHeight.empty()){
		tmpArea = long(stkHeight.top()* (hist.size()-stkPos.top()));
		stkHeight.pop();
		stkPos.pop();
		if (area < tmpArea)
			area = tmpArea;
	}
	return area;
}
