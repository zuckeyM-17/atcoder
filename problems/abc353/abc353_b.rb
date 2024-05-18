def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n, k = getis
a = getis

ans = 0
seats = 0
a.each do |ai|
  if seats + ai > k
    ans += 1
    seats = 0
  end
  seats += ai
end

puts ans + 1