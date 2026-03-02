import java.util.Scanner;
import java.util.HashMap;
public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        if(sc.hasNextInt()) {
            int n = sc.nextInt();
            HashMap<String, Integer> h = new HashMap<>(n);
            for(int i=0; i<n; i++) h.put(String.valueOf(i), i);
        }
    }
}