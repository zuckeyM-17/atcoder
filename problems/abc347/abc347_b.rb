def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

s = gets.chomp

c = Hash.new(false)
ln = s.size
ln.times do |i|
  (ln - i).times do |j|
    c[s[i..i+j]] = true
  end
end

puts c.size