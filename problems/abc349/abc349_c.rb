def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

s = gets.chomp
t = gets.chomp.downcase

cur = 0
s.chars.each.with_index do |c, i|
  if t[cur] == c
    cur += 1
    if cur == 3
      puts "Yes"
      exit
    end
  end
end

if cur == 2 && t[2] == "x"
  puts "Yes"
  exit
end

puts "No"