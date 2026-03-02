import java.util.Scanner;
public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        if(sc.hasNextLong()) {
            long n = sc.nextLong(), s=0;
            for(long i=1; i<=n; i++) s+=i;
        }
    }
}