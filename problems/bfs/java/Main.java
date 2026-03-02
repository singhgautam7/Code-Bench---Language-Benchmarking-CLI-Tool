import java.util.*;
public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        if(sc.hasNextInt()) {
            int V = sc.nextInt(), E = sc.nextInt();
            ArrayList<Integer>[] adj = new ArrayList[V];
            for(int i=0; i<V; i++) adj[i] = new ArrayList<>();
            for(int i=0; i<E; i++) {
                int u = sc.nextInt(), v = sc.nextInt(), w = sc.nextInt();
                adj[u].add(v);
            }
            Queue<Integer> q = new LinkedList<>();
            q.add(0);
            boolean[] vis = new boolean[V];
            vis[0] = true;
            while(!q.isEmpty()) {
                int curr = q.poll();
                for(int nxt : adj[curr]) {
                    if(!vis[nxt]) { vis[nxt]=true; q.add(nxt); }
                }
            }
        }
    }
}