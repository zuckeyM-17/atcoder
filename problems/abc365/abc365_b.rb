def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti
a = getis

i = 0

a.map! do |x| 
  i+=1
  {idx: i, val: x}
end

a.sort_by! { |x| x[:val] }
a.reverse!

puts a[1][:idx]