#include "node.h"
#include<queue>
#include<vector>
using namespace std;

/*
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};
*/
class Solution {
public:
    Node* connect(Node* root) {
        if (!root) {
            return root;
        }
        Node *p = root;
        Node *q = NULL;
        queue<Node *> que;
        que.push(p);
        int size = 0;   //每一层的大小

        while(!que.empty()) {
            size = que.size();
            vector<Node *> vec; //用于存每一层的节点
            for (int i=0; i< size; i++){
                q = que.front();
                que.pop();
                if (q->left) {
                    que.push(q->left);
                }
                if (q->right) {
                    que.push(q->right);
                }
                vec.push_back(q);
            }
            for (int i=0; i< vec.size()-1; i++){
                vec[i]->next=vec[i+1];
            }
        }
        return root;
    }
};