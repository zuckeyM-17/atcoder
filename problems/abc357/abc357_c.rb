def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti

def level_up(b)
  n = []
  bs = b.split("\n")
  bs.each do |l|
    n << l*3
  end
  bs.each do |l|
    n << l + '.' * bs.size + l
  end
  bs.each do |l|
    n << l*3
  end
  n.join("\n")
end

l0 = '#'
l1 = level_up(l0)
l2 = level_up(l1)
l3 = level_up(l2)
l4 = level_up(l3)
l5 = level_up(l4)
l6 = level_up(l5)

if n == 0
  puts '#'
elsif n == 1
  puts l1
elsif n == 2
  puts l2
elsif n == 3
  puts l3
elsif n == 4
  puts l4
elsif n == 5
  puts l5
elsif n == 6
  puts l6
else
  puts l6
end