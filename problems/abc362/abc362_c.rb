def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti

ls = []
rs = []

n.times do
  l, r = getis
  ls.push(l)
  rs.push(r)
end

