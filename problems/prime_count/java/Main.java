import java.util.Scanner;
public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        if(sc.hasNextInt()) {
            int n = sc.nextInt();
            if (n < 2) return;
            boolean[] sieve = new boolean[n+1]; // false = prime
            int count = 0;
            for(int p=2; p<=n; p++) {
                if(!sieve[p]) {
                    count++;
                    if((long)p*p <= n) {
                        for(int i=p*p; i<=n; i+=p) sieve[i] = true;
                    }
                }
            }
        }
    }
}