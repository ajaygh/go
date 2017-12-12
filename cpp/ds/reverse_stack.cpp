#include <iostream>
#include <stack>

using namespace std;


void printStack(stack<int> s){
    while(!s.empty()){
        cout << s.top() <<" ";
        s.pop();
    }
    cout << endl;
}

void putAtBottom(stack<int> &s, int top){
    if(s.empty()){
        s.push(top);
        return;
    }
    auto ttop = s.top();
    s.pop();
    putAtBottom(s, top);
    s.push(ttop);
}

void revStack(stack<int> &sin){
    if (sin.size() < 2) return;
    auto top = sin.top();
    sin.pop();
    revStack(sin);
    putAtBottom(sin, top);
    cout << top <<endl;
    printStack(sin);
}

int main(){
    stack<int> s{{1,2,3,4}};
    printStack(s);
    revStack(s);
   
    cout  <<"reversed stack is " << endl;
    printStack(s);
}