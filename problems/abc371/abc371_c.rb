def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

class Graph
  attr_accessor :edges, :adj_matrix

  def initialize(n)
    @n = n
    @edges = []
    @adj_matrix = Array.new(n+1) { Array.new(n+1, 0) }
  end

  def add_edge(u, v)
    @edges << [u, v]
    @adj_matrix[u][v] = 1
    @adj_matrix[v][u] = 1
  end
end

n = geti
mg = geti

g = Graph.new(n)

mg.times do
  u, v = getis
  g.add_edge(u, v)
end

mh = geti
h = Graph.new(n)

mh.times do
  a, b = getis
  h.add_edge(a, b)
end

cost = Hash.new { |h, k| h[k] = Hash.new(0) }
(1...n-1).each do |i|
  a = getis
  cnt = 0
  (i+1...n).each do |j|
    cost[i][j] = a[cnt]
    cnt += 1
  end
end



def min_cost(g, h, cost, n)
  total_cost = 0

  (1...(n-1)).each do |i|
    ((i+1)...n).each do |j|
      if g.adj_matrix[i][j] != h.adj_matrix[i][j]
        total_cost += cost[i][j]
      end
    end
  end

  total_cost
end

puts min_cost(g, h, cost, n)