import java.util.Scanner;
public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        if(sc.hasNext()) {
            String c = sc.next(); int n = sc.nextInt();
            StringBuilder sb = new StringBuilder(n);
            for(int i=0; i<n; i++) sb.append(c);
        }
    }
}