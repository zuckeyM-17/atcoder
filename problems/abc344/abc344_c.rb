def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti
a = getis
m = geti
b = getis
l = geti
c = getis


candidates = Hash.new(false)

a.each do |i|
  b.each do |j|
    c.each do |k|
      candidates[i + j + k] = true
    end
  end
end

q = geti
x = getis
x.each do |i|
  if candidates[i]
    puts "Yes"
  else
    puts "No"
  end
end