def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n, m = getis

a = getis
b = getis

c = a.map { |ai| { v: ai, t: "a" } }.concat(b.map { |bi| { v: bi, t: "b" } })

c.sort_by! { |ci| ci[:v] }

c.each.with_index do |ci, i|
  next if i == 0

  if ci[:t] == "a" && c[i - 1][:t] == c[i][:t]
    puts "Yes"
    exit
  end
end

puts "No"

