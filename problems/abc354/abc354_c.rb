def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti
cards = []
a, c = Array.new(0), Array.new(0)

n.times do |i|
  ai, ci = getis
  cards << {
    a: ai,
    c: ci,
    i: i + 1
  }
end

cards.sort_by! { |card| [-card[:a], card[:c]] }

remaining_cards = []

min_cost = Float::INFINITY

cards.each do |card|
  if card[:c] < min_cost
    remaining_cards << card
    min_cost = card[:c]
  end
end

puts remaining_cards.size
remaining_cards.sort_by! { |card| card[:i] }
puts remaining_cards.map { |card| card[:i] }.join(" ")



