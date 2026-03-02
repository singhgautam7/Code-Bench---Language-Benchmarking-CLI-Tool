import java.util.Scanner;
import java.util.ArrayList;
import java.util.Arrays;
public class Main {
    static int[] mergesort(int[] arr) {
        if(arr.length <= 1) return arr;
        int mid = arr.length / 2;
        int[] L = mergesort(Arrays.copyOfRange(arr, 0, mid));
        int[] R = mergesort(Arrays.copyOfRange(arr, mid, arr.length));
        int[] res = new int[arr.length];
        int i=0, j=0, k=0;
        while(i < L.length && j < R.length) {
            if(L[i] < R[j]) res[k++] = L[i++];
            else res[k++] = R[j++];
        }
        while(i < L.length) res[k++] = L[i++];
        while(j < R.length) res[k++] = R[j++];
        return res;
    }
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        ArrayList<Integer> list = new ArrayList<>();
        while(sc.hasNextInt()) list.add(sc.nextInt());
        int[] arr = new int[list.size()];
        for(int i=0; i<list.size(); i++) arr[i] = list.get(i);
        mergesort(arr);
    }
}