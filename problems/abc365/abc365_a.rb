def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

y = geti

if y % 400 == 0
  puts "366"
elsif y % 100 == 0
  puts "365"
elsif y % 4 == 0
  puts "366"
else
  puts "365"
end