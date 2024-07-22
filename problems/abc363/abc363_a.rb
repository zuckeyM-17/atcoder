def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

r = geti

if r < 100
  puts 100 - r
elsif r < 200
  puts 200-r
elsif r < 300
  puts 300-r
else
  puts 400-r
end