def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

def calc(arr)
  s = arr.size
  if s < 2
    return arr
  end
  if arr[s - 1] == arr[s - 2]
    arr.pop
    arr[s-2] = arr[s-2]+ 1
    return calc(arr)
  end

  arr
end

_ = geti
a = getis
c = []
a.each do |ai|
  c << ai
  c = calc(c)
end

puts c.size