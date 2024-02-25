def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n, x = getis
a = gets.chomp

queue = [x-1]
a[x-1] = '@'

while queue.size > 0
  pos = queue.shift

  if pos-1 >= 0 && a[pos-1] == '.'
    a[pos-1] = '@'
    queue.push(pos-1)
  end

  if pos+1 < n && a[pos+1] == '.'
    a[pos+1] = '@'
    queue.push(pos+1)
  end
end

puts a
