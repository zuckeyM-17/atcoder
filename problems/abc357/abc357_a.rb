def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n, m = getis

h = getis

h.each.with_index do |hi, i|
  if hi > m || m == 0
    puts i
    exit
  end

  m -= hi
end

puts n