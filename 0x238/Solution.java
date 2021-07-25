package _x238;

class Solution{
    public int[] productExceptSelf(int[] nums){

        int n = nums.length;

        int[] L = new int[n];
        int[] R = new int[n];
        int[] ans = new int[n];

//        L[0] = nums[0];
//        R[n-1] = nums[n-1];

        L[0] = 1;
        R[n-1] = 1;
        for (int i = 1,j=n-2;i < n && j>=0;i++,j--){
            L[i] = nums[i-1] * L[i-1];
            R[j] = nums[j+1] * R[j+1];
        }
        for (int i = 0; i < n; i++) {
            ans[i]=L[i]*R[i];
        }

        return ans;
    }
}