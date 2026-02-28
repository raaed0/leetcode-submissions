use std::collections::{HashMap, HashSet, VecDeque};

struct Edge {
    To: String,
    Weight: f64,
}

impl Solution {
    pub fn calc_equation(equations: Vec<Vec<String>>, values: Vec<f64>, queries: Vec<Vec<String>>) -> Vec<f64> {
        let mut graph: HashMap<String, Vec<Edge>> = HashMap::new();
        for (i, eq) in equations.iter().enumerate() {
            graph.entry(eq[0].clone()).or_insert_with(Vec::new).push(Edge {
                To: eq[1].clone(),
                Weight: values[i]
            });
            graph.entry(eq[1].clone()).or_insert_with(Vec::new).push(Edge {
                To: eq[0].clone(),
                Weight: 1.0 / values[i]
            });
        }
        
        let mut res = vec![-1.0; queries.len()];
        for i in 0..queries.len() {
            res[i] = bfs(&graph, &queries[i][0], &queries[i][1]);
        }

        return res;
    }
}

fn bfs(graph: &HashMap<String, Vec<Edge>>, source: &str, dst: &str) -> f64 {
    if !graph.contains_key(source) || !graph.contains_key(dst) {
        return -1.0;
    }

    let mut queue: VecDeque<(&str, f64)> = VecDeque::new();
    let mut visited: HashSet<&str> = HashSet::new();
    queue.push_back((source, 1.0));
    visited.insert(source);


    while let Some((node, weight)) = queue.pop_front() {
        if node == dst {
            return weight;
        }
        if let Some(edges) = graph.get(node) {
            for edge in edges {
                if !visited.contains(edge.To.as_str()) {
                    visited.insert(&edge.To);
                    queue.push_back((&edge.To, weight*edge.Weight));
                }
            }
        }
    }
    
    -1.0
}
