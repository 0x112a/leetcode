public class ListNode{
    int Val;
    ListNode next;
    ListNode(){};
    ListNode(int val){
        this.Val = val;
    }
    ListNode(int val, ListNode next){
        this.Val = val;
        this.next = next;
    }
}

class Solution{
    public ListNode mergeTwoLists(ListNode l1,ListNode l2){
        if (l1 == null && l2 == null){
            return null;
        }
        if (l1 == null){
            return l2;
        }
        if (l2 == null){
            return l1;
        }

        ListNode ansHead = new ListNode();
        ListNode p;

        if(l1.val > l2.val){
            ansHead.next = l2;
            p = l2;
            l2 = l2.next;
        }else{
            ansHead.next = l1;
            p = l1;
            l1 = l1.next;

        }

        while(l1 != null || l2 != null){
            if (l1 == nill){
                p.next = l2;
                break;
            }
            if (l2 == null){
                p.next = l1;
                break
            }

            if(l1.val > l2.val){
                p.next = l2;
                p = p.next;
                l2 = l2.next;
            }else{
                p.next = l1;
                p = p.next;
                l1 = l1.next;
            }
        }

        return ansHead.next;
    }
}
