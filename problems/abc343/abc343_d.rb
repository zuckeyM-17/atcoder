def getis
  gets.chomp.split(" ").map(&:to_i)
end

n, t = getis

l = Hash.new(0)
l[0] = n
m = Hash.new(0)
d = 1
t.times do
  a, b = getis

  l[m[a]] = l[m[a]] - 1
  if l[m[a]] == 0
    d = d - 1
    l.delete(m[a])
  end
  m[a] = m[a] + b
  if l[m[a]] == 0
    d = d + 1
  end
  l[m[a]] = l[m[a]] + 1

  puts d
end