#include <iostream>
#include <vector>

using namespace std;

void prt(vector<int> &d, vector<bool> &s){
    for(int i = 0; i < d.size();i++){
        if (s[i]){
            cout << d[i] <<" ";
        }
    }
    cout << endl;
}

void printAll(vector<int> &d, vector<bool> &s, int pos){
    if(pos == d.size()-1){
        prt(d, s);
        s[pos] = true;
        prt(d,s);
        s[pos] = false;
        return;
    }
    printAll(d, s, pos+1);
    s[pos] = true;
    printAll(d, s, pos+1);
    s[pos] = false;
}

int main(){
    vector<int> d{{1,2,3,4,5, 6,7}};
    vector<bool> status(d.size());
    printAll(d, status, 0);
    return 0;
}