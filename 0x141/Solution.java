class ListNode {
     int val;
     ListNode next;
     ListNode(int x) {
         val = x;
         next = null;
     }
}
public class Solution {
    public boolean hasCycle(ListNode head) {
        if (head==null || head.next == null){
            return false;
        }

        ListNode slowP = head;
        ListNode fastP = head.next;

        for(;fastP != null;){
            if(fastP == slowP || fastP.next == slowP){
                return true;
            }else if(fastP.next != null){
                fastP = fastP.next.next;
            }else{
                fastP = fastP.next;
            }
            slowP= slowP.next;
        }
        return false;

    }
}
