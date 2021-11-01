#include <iostream>
#include <vector>

using namespace std;

class Node
{
public:
    int Val;
    Node * pNext;
};

Node* RemoveVal2(Node* pHead, int target)
{
    if (nullptr == pHead)
    {
        return nullptr;
    }

    while (nullptr != pHead && pHead->Val == target)
    {
        pHead = pHead->pNext;
    }

    if (nullptr == pHead)
    {
        return nullptr;
    }

    Node *ret = pHead;
    while (pHead->pNext != nullptr)
    {
        if (pHead->pNext->Val == target)
        {
            Node *tmp = pHead->pNext;
            pHead->pNext = pHead->pNext->pNext;
            delete tmp;
            tmp = nullptr;
        }
        else 
        {
            pHead = pHead->pNext;
        }
    }

    return ret;
} 

Node * RemoveVal(Node *pHead, int target)
{
    if (nullptr == pHead)
    {
        return nullptr;
    }
    
    Node tmp;
    tmp.pNext = pHead;

    Node *first = &tmp;
    while(first->pNext != nullptr)
    {
        if (first->pNext->Val == target)
        {
            Node *t = first;
            first->pNext = first->pNext->pNext;
            delete t;
            t = nullptr;
        }
        else
        {
            first = first->pNext;
        }
    }

    return tmp.pNext;
}

int main(int argc, char **argv)
{
    Node *pNode1 = new Node;
    Node *pNode2 = new Node;
    Node *pNode3 = new Node;
    Node *pNode4 = new Node;
    Node *pNode5 = new Node;
    pNode5->Val = 1;
    pNode5->pNext = nullptr;
    pNode4->Val = 2;
    pNode4->pNext = pNode5;
    pNode3->Val = 3;
    pNode3->pNext = pNode4;
    pNode2->Val = 4;
    pNode2->pNext = pNode3;
    pNode1->Val = 5;
    pNode1->pNext = pNode2;

    Node *ret = RemoveVal(pNode1, 5);

    while (nullptr != ret)
    {
        printf("%d\n", ret->Val);
        ret = ret->pNext;
    }
    return 1;
}