import java.util.Scanner;
import java.util.HashMap;
public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        HashMap<String, Integer> counts = new HashMap<>();
        while(sc.hasNext()) {
            String w = sc.next();
            counts.put(w, counts.getOrDefault(w, 0) + 1);
        }
    }
}