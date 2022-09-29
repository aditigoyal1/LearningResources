#include<bits/stdc++.h>
using namespace std;

int main(){
    int a[]={1,4,5,6,9,9};
    int n = sizeof(a)/sizeof(a[0]);
    //lower_bound return the iterator pointing to tht element or just immediate greater element.
    cout<<lower_bound(a,a+ n,4) - a<<"\n";
    cout<<lower_bound(a,a+n,7) -a<<"\n";
    cout<<lower_bound(a,a+n,10)-a<<"\n"; 

    //upper bound always return the iterator pointing to the next greater element.
     cout<<upper_bound(a,a+ n,4) - a<<"\n";
    cout<<upper_bound(a,a+n,7) -a<<"\n";
    cout<<upper_bound(a,a+n,10)-a<<"\n"; 


    //For vector the syntax is:
    // int ind = upper_bound(v.begin(),v.end(),v)-v.begin();

    //find the first occurence of x in a sorted array. If it does not exists, print -1.
    int a1[]={1,4,4,4,4,9,9,10,11};
       n = sizeof(a1)/sizeof(a1[0]);
       int x=4;
    int ind = lower_bound(a, a+n,x);
    if(ind!=n && a1[ind]==x){
        cout<<x;
    }else{
        cout<<-1;
    }
}