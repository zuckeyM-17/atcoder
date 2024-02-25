def geti
  gets.chomp.to_i
end

q = geti

queue = []
q.times do
  query, book = gets.chomp.split(" ")

  if query == "1"
    queue.push(book)
  elsif query == "2"
    puts queue.first
  else
    queue.shift
  end
end