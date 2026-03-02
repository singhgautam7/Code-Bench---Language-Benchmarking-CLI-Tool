import java.util.Scanner;
public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        if(sc.hasNextLong()) {
            long n = sc.nextLong(), a=0, b=1, t;
            for(long i=0; i<n; i++) { t=a; a=b; b=t+b; }
        }
    }
}