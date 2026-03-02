import java.util.*;
public class Main {
    static ArrayList<Integer>[] adj;
    static boolean[] vis;
    static void dfs(int u) {
        vis[u] = true;
        for(int v : adj[u]) { if(!vis[v]) dfs(v); }
    }
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        if(sc.hasNextInt()) {
            int V = sc.nextInt(), E = sc.nextInt();
            adj = new ArrayList[V];
            for(int i=0; i<V; i++) adj[i] = new ArrayList<>();
            for(int i=0; i<E; i++) {
                int u = sc.nextInt(), v = sc.nextInt(), w = sc.nextInt();
                adj[u].add(v);
            }
            vis = new boolean[V];
            dfs(0);
        }
    }
}