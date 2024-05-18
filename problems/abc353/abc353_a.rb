def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti
h = getis

a = h.first

h.each_with_index do |hi, i|
  if hi > a
    puts "#{i+1}"
    exit
  end
end

puts "#{-1}"