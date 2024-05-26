def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

a, b = getis

if a == b
  puts "-1"
  exit
end

if (a == 1 && b == 2) || (a == 2 && b == 1)
  puts "3"
  exit
end

if (a== 1 && b == 3) || (a == 3 && b == 1)
  puts "2"
  exit
end

if (a == 2 && b == 3) || (a == 3 && b == 2)
  puts "1"
  exit
end