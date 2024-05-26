def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n, t = getis
a = getis

rows = n.times.map do |i|
  []
end

cols = n.times.map do |i|
  []
end

crosses = 2.times.map do |i|
  []
end

a.each_with_index do |ai, i|
  rows[(ai - 1) / n] << ai
  if rows[(ai - 1) / n].size == n
    puts "#{i + 1}"
    exit
  end

  cols[(ai - 1) % n] << ai
  if cols[(ai - 1) % n].size == n
    puts "#{i + 1}"
    exit
  end

  if (ai - 1) / n == (ai - 1) % n
    crosses[0] << ai
    if crosses[0].size == n
      puts "#{i + 1}"
      exit
    end
  end

  if (ai - 1) / n == n - 1 - (ai - 1) % n
    crosses[1] << ai
    if crosses[1].size == n
      puts "#{i + 1}"
      exit
    end
  end
end

puts "-1"