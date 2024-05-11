def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

h, w = getis

s = []
h.times do
  s << gets.chomp.split("")
end

def check_around(s, i, j)
  p1 = if i == 0
    [i+1]
  elsif i == h-1
    [i-1]
  else
    [i-1, i+1]
  end

  p2 = if j == 0
    [j+1]
  elsif j == w-1
    [j-1]
  else
    [j-1, j+1]
  end

  p1.each do |ii|
    p2.each do |jj|
      return true if s[i + ii][j + jj] == "#"
    end
  end
end

h.times do |i|
  w.times do |j|
    next if s[i][j] == "#"
    s[i][j] = "x" if check_around(s, i, j)
  end
end

max = 0
candidates = []
h.times do |i|

end
