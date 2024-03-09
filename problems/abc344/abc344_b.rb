def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

a = readlines.map do |i|
  i.chomp.to_i
end
a.reverse.each do |i|
  puts i
end