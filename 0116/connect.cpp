#include "node.h"
#include<queue>
using namespace std;

class Solution {
public:
    Node* connect(Node* root) {
        if (!root) {
            return root;
        }
        Node *p = root;
        queue<Node *> que;
        que.push(p);
        int size = 0;   //每一层的大小

        while(!que.empty()) {
            size = que.size();
            Node *pre = NULL;
            Node *cur = NULL;
            for (int i=0; i< size; i++){
                cur = que.front();
                que.pop();
                if (i > 0) {
                    pre->next = cur;
                }
                pre = cur;
                if (cur->left) {
                    que.push(cur->left);
                }
                if (cur->right) {
                    que.push(cur->right);
                }
            }
        }
        return root;
    }
};