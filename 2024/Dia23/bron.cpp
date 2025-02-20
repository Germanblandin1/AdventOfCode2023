#include <iostream>
#include <vector>
#include <unordered_set>
#include <algorithm>

using namespace std;

// Función para encontrar la intersección de dos conjuntos
unordered_set<int> intersect(const unordered_set<int>& a, const unordered_set<int>& b) {
    unordered_set<int> result;
    for (int elem : a) {
        if (b.find(elem) != b.end()) {
            result.insert(elem);
        }
    }
    return result;
}

// Función para encontrar todas las cliques máximas
void BronKerbosch(vector<int>& R, unordered_set<int>& P, unordered_set<int>& X, 
                  const vector<vector<int>>& graph, vector<vector<int>>& cliques) {
    if (P.empty() && X.empty()) {
        // R es una clique máxima
        cliques.push_back(R);
        return;
    }

    unordered_set<int> P_copy = P; // Copia para iterar sin modificar P original

    for (int v : P_copy) {
        // Expandir R con el nodo v
        R.push_back(v);
        unordered_set<int> newP = intersect(P, unordered_set<int>(graph[v].begin(), graph[v].end()));
        unordered_set<int> newX = intersect(X, unordered_set<int>(graph[v].begin(), graph[v].end()));

        // Llamada recursiva
        BronKerbosch(R, newP, newX, graph, cliques);

        // Mover v de P a X
        R.pop_back();
        P.erase(v);
        X.insert(v);
    }
}

// Crear un grafo en forma de lista de adyacencia
vector<vector<int>> createGraph(int numNodes, const vector<pair<int, int>>& edges) {
    vector<vector<int>> graph(numNodes);
    for (const auto& edge : edges) {
        int u = edge.first, v = edge.second;
        graph[u].push_back(v);
        graph[v].push_back(u);
    }
    return graph;
}

int main() {
    // Definir las aristas del grafo
    vector<pair<int, int>> edges = {
        {0, 1}, {0, 2}, {1, 2}, {0, 3}, 
        {3, 4}, {4, 5}, {3, 5}
    };

    int numNodes = 6; // Número de nodos (0 a 5)
    vector<vector<int>> graph = createGraph(numNodes, edges);

    // Inicializar R, P, X y la lista de cliques
    vector<int> R;
    unordered_set<int> P, X;
    vector<vector<int>> cliques;

    // P contiene todos los nodos inicialmente
    for (int i = 0; i < numNodes; ++i) {
        P.insert(i);
    }

    // Llamada inicial al algoritmo
    BronKerbosch(R, P, X, graph, cliques);

    // Imprimir las cliques máximas
    cout << "Cliques máximas:" << endl;
    for (const auto& clique : cliques) {
        for (int node : clique) {
            cout << node << " ";
        }
        cout << endl;
    }

    return 0;
}
