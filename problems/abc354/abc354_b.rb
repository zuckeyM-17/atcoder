def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti

s, c = [], []

n.times do |i|
  si, ci = gets.chomp.split(" ")
  s << si
  c << ci.to_i
end
s.sort!
puts s[c.sum % n]