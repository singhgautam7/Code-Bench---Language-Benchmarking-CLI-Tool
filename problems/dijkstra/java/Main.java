import java.util.*;
public class Main {
    static class Edge { int to, w; Edge(int to, int w){this.to=to; this.w=w;} }
    static class Node implements Comparable<Node> {
        int n, d; Node(int n, int d){this.n=n; this.d=d;}
        public int compareTo(Node o) { return Integer.compare(this.d, o.d); }
    }
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        if(!sc.hasNextInt()) return;
        int V = sc.nextInt(), E = sc.nextInt();
        ArrayList<Edge>[] adj = new ArrayList[V];
        for(int i=0; i<V; i++) adj[i] = new ArrayList<>();
        for(int i=0; i<E; i++) {
            int u = sc.nextInt(), v = sc.nextInt(), w = sc.nextInt();
            adj[u].add(new Edge(v, w));
        }
        int[] dist = new int[V];
        Arrays.fill(dist, 1000000000);
        dist[0] = 0;
        PriorityQueue<Node> pq = new PriorityQueue<>();
        pq.add(new Node(0, 0));
        while(!pq.isEmpty()) {
            Node curr = pq.poll();
            if(curr.d > dist[curr.n]) continue;
            for(Edge e : adj[curr.n]) {
                if(dist[curr.n] + e.w < dist[e.to]) {
                    dist[e.to] = dist[curr.n] + e.w;
                    pq.add(new Node(e.to, dist[e.to]));
                }
            }
        }
    }
}