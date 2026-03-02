import java.math.BigInteger;
import java.util.Scanner;
public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        if(sc.hasNextInt()) {
            int n = sc.nextInt();
            BigInteger s = BigInteger.ONE;
            for(int i=1; i<=n; i++) s = s.multiply(BigInteger.valueOf(i));
        }
    }
}