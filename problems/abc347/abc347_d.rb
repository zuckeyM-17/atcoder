def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

def popcount_2_bs(max, p)
  (1..max).to_a.combination(p).map do |a|
    (1..max).map do |i|
      a.include?(i) ? "1" : "0"
    end.join("")
  end
end

a, b, C = getis
arr = (1..60).to_a

cX, cY = popcount_2_bs(60, a), popcount_2_bs(60, b)
puts cX.first
