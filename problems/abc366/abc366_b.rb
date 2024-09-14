def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti

s = []

n.times do
  s << gets.chomp.split("")
end

M = s.map { |a| a.count }.max

s.each do |si|
  cnt = si.count
  (M - cnt).times do
    si << "*"
  end
end

t = []

M.times do |i|
  t << []
  n.times do |j|
    t[i] << s[j][i]
  end

  t[i].reverse!
  puts t[i].join.gsub(/\*+$/, '')
end