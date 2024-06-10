def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

s = gets.chomp

l, u = 0, 0

lowercases = ('a'..'z').to_a

s.each_char do |c|
  if lowercases.include?(c)
    l += 1
  else
    u += 1
  end
end

if l > u
  puts s.downcase
else
  puts s.upcase
end