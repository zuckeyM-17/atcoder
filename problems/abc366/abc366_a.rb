def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n, t, a = getis

nokori = n - (t + a)

def abs(x)
  x > 0 ? x : -x
end

if abs(t - a) > nokori
  puts "Yes"
else
  puts "No"
end