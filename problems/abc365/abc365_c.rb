def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n, m = getis
a = getis

if a.sum <= m
  puts "infinite"
  exit
end

def max_subsidy(m, a)
  low = 0
  high = a.max

  while low <= high
    mid = (low + high) / 2
    total_cost = a.reduce(0) { |sum, cost| sum + [mid, cost].min }

    if total_cost <= m
      low = mid + 1
    else
      high = mid - 1
    end
  end

  return high
end

puts max_subsidy(m, a)