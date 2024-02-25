def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

q = geti

stack = []
q.times do
  query, book = gets.chomp.split(" ")
  if query == "1"
    stack.push(book)
  elsif query == "2"
    puts stack.last
  elsif query == "3"
    stack.pop
  end
end
