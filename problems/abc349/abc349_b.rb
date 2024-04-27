def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

s = gets.chomp

h = Hash.new(0)
s.chars.each do |c|
  h[c] += 1
end

h2 = Hash.new(0)
h.values.each do |v|
  h2[v] += 1
end

h2.each do |k, v|
  if v != 2
    puts "No"
    exit
  end
end

puts "Yes"