def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti

one_line = false
n.times do
  *a = getis
  nums = []
  a.each.with_index(1) do |v, idx|
    if v == 1
      nums << idx
    end
  end
  if nums.size > 0
    one_line = true
    puts nums.join(" ")
  end
end

if !one_line
  puts
end