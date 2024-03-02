def getis
  gets.chomp.split(" ").map(&:to_i)
end

c1, c2, c3 = getis

if c1 + c2*2 + c3*3 != 1029
  puts "No"
  exit
end

puts "Yes"