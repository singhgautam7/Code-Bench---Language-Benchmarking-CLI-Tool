import java.util.Scanner;
public class Main {
    static int gcd(int a, int b) {
        while(b != 0) { int t = b; b = a % b; a = t; }
        return a;
    }
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        if(sc.hasNextInt()) {
            int n = sc.nextInt();
            for(int i=1; i<=n; i++) gcd(1000000007, i);
        }
    }
}