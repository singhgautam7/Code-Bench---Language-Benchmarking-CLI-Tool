import java.util.Scanner;
import java.util.ArrayList;

public class Main {
    static void quicksort(int[] arr, int left, int right) {
        if (left >= right) return;
        int pivot = arr[left + (right - left) / 2];
        int i = left, j = right;
        while (i <= j) {
            while (arr[i] < pivot) i++;
            while (arr[j] > pivot) j--;
            if (i <= j) {
                int temp = arr[i];
                arr[i] = arr[j];
                arr[j] = temp;
                i++; j--;
            }
        }
        quicksort(arr, left, j);
        quicksort(arr, i, right);
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        ArrayList<Integer> list = new ArrayList<>();
        while(sc.hasNextInt()) list.add(sc.nextInt());

        int[] arr = new int[list.size()];
        for(int i=0; i<list.size(); i++) arr[i] = list.get(i);

        if (arr.length > 0) quicksort(arr, 0, arr.length - 1);
    }
}
