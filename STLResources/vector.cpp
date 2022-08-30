
#include <iostream>
#include <vector>
using namespace std;

void explainVectors(){
    vector<int> v;
    v.push_back(1);
    v.emplace_back(2);

    vector<pair<int,int>>vec;
    vec.push_back({8,9});
    vec.emplace_back(4,7);

    vector<int> v4(5,100);

    vector<int> v5(5);
    vector<int> v1(5,200);
    vector<int> v2(v1);


    for(vector<int>::iterator it = v.begin(); it!=v.end();it++){
        cout<<*(it)<<" ";
    } 
    cout<<"\n";

    for(auto it1 : vec){
        cout<<it1.first<<" "<<it1.second;
        cout<<"\n";
    }

    vector<int>:: iterator it = v.begin();
    it++;
    cout<<*(it)<<" ";
    cout<<v[0]<<"  "<<v.at(0);
    cout<<v.back()<<" ";
    v.clear();
    cout<<"\n";
    v={10,20,12,23};
     

    //{10,20,12,23}
    v.erase(v.begin()+1);
    for(auto it1 : v){
        cout<<it1<<" ";
        
    }
    //results {10,12,23}
    
    
    v={10,20,12,23,35};
    //{10,20,12,23,35}
    v.erase(v.begin()+2,v.begin()+4);// {10,20,35} [start,end]
     cout<<"\n";
      for(auto it1 : v){
        cout<<it1<<" ";
        
    }

    vector<int> (2,100);
    v.insert(v.begin(),300);//{300,100,100}



}
int main(){
  explainVectors();

    return 0;

}

// r.end()===> points to memory after the last element.
// we have to it-- to move to the last the element.
//v.rend()===> reverse end {10,20,30,40}===>{40,30,20,10}
//v.rbegin()===reverse begin {10,20,30,40} ===>{40,30,20,10} points at 40