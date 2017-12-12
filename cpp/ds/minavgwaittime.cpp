// find min avg wait time
#include <iostream>
#include <utility>
#include <vector>
#include <queue>

using namespace std;

long calcMinAvgWaitTime(vector<pair<int, int>> &orders);

int main(){
    int n;
    cin >> n;

    vector<pair<int, int>> orders(n);
    for (int i = 0; i < n; i++){
        cin >> orders[i].first >> orders[i].second;
    }
    auto res = calcMinAvgWaitTime(orders);
    cout << res << endl;
    return 0;
}

long calcMinAvgWaitTime(vector<pair<int, int>> &orders) {
    if (orders.size() == 0) return 0;

    auto cmp = [](pair<int, int> left, pair<int, int> right){
        return left.first + left.second > right.first + right.second; };
    priority_queue<pair<int, int>, vector<pair<int, int> >, decltype(cmp)> tstore(cmp);
    
    for (auto el : orders) tstore.push(el);
    long res = 0;
    long starttime = 0; // start timepoint for each job
    while(!tstore.empty()){
        auto tmp = tstore.top();
        res += starttime + tmp.second - tmp.first; // current wait time
        starttime += tmp.second; 
        tstore.pop();
    }
    return res/orders.size();
}