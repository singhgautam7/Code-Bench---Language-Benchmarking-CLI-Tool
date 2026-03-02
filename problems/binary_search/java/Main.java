import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.StringTokenizer;

public class Main {
    /**
     * Perform binary search on a sorted array.
     * Returns the index of target, or -1 if not found.
     */
    static int binarySearch(int[] arr, int target) {
        int low = 0, high = arr.length - 1;
        while (low <= high) {
            int mid = low + (high - low) / 2;
            if (arr[mid] == target) return mid;
            else if (arr[mid] < target) low = mid + 1;
            else high = mid - 1;
        }
        return -1;
    }

    public static void main(String[] args) throws Exception {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder sb = new StringBuilder();
        String line;
        while ((line = br.readLine()) != null) {
            sb.append(line).append(" ");
        }

        StringTokenizer st = new StringTokenizer(sb.toString());
        int count = st.countTokens();
        if (count < 2) {
            System.err.println("Error: expected array and target on stdin");
            System.exit(1);
        }

        // All tokens: last is target, rest are array
        int[] tokens = new int[count];
        for (int i = 0; i < count; i++) {
            tokens[i] = Integer.parseInt(st.nextToken());
        }

        int target = tokens[count - 1];
        int[] arr = new int[count - 1];
        System.arraycopy(tokens, 0, arr, 0, count - 1);

        int result = binarySearch(arr, target);
        System.out.println(result);
    }
}
