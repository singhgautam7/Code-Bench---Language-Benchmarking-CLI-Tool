import java.util.Scanner;
import java.util.HashSet;
public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        if(!sc.hasNext()) return;
        String s = sc.next();
        HashSet<Character> chars = new HashSet<>();
        int l=0, res=0;
        for(int r=0; r<s.length(); r++) {
            while(chars.contains(s.charAt(r))) { chars.remove(s.charAt(l)); l++; }
            chars.add(s.charAt(r));
            res = Math.max(res, r-l+1);
        }
    }
}