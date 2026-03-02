import java.util.Scanner;
import java.util.ArrayList;
public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        ArrayList<String> arr = new ArrayList<>();
        while(sc.hasNext()) arr.add(sc.next());
        if(arr.isEmpty()) return;
        String target = arr.get(arr.size()-1);
        for(int i=0; i<arr.size()-1; i++) {
            if(arr.get(i).equals(target)) break;
        }
    }
}