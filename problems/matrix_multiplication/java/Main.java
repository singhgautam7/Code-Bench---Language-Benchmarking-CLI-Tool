import java.util.Scanner;
public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        if(!sc.hasNextInt()) return;
        int n = sc.nextInt();
        int[][] A = new int[n][n];
        int[][] B = new int[n][n];
        int[][] C = new int[n][n];
        for(int i=0; i<n; i++) for(int j=0; j<n; j++) A[i][j] = sc.nextInt();
        for(int i=0; i<n; i++) for(int j=0; j<n; j++) B[i][j] = sc.nextInt();
        for(int i=0; i<n; i++) {
            for(int k=0; k<n; k++) {
                for(int j=0; j<n; j++) C[i][j] += A[i][k] * B[k][j];
            }
        }
    }
}