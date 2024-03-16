def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n, h, w = getis

ss = []

n.times do
  a, b = getis
  ss << [a, b]
end

def can_fill?(h, w, rects, layout = Array.new(h) { Array.new(w, false) }, x = 0, y = 0)
  return true if x >= h
  next_x, next_y = y < w - 1 ? [x, y + 1] : [x + 1, 0]

  return can_fill?(h, w, rects, layout, next_x, next_y) if layout[x][y]

  rects.each_with_index do |(rect_h, rect_w), i|
    [[rect_h, rect_w], [rect_w, rect_h]].uniq.each do |rh, rw|
      next if x + rh > h || y + rw > w
      next if !can_place?(layout, x, y, rh, rw)

      place_rect(layout, x, y, rh, rw, true)

      return true if can_fill?(h, w, rects[0...i] + rects[i+1..-1], layout, next_x, next_y)

      place_rect(layout, x, y, rh, rw, false)
    end
  end

  false
end

def can_place?(layout, x, y, rh, rw)
  (x...x+rh).all? { |xi| (y...y+rw).all? { |yi| !layout[xi][yi] } }
end

def place_rect(layout, x, y, rh, rw, value)
  (x...x+rh).each { |xi| (y...y+rw).each { |yi| layout[xi][yi] = value } }
end

puts can_fill?(h, w, ss) ? "Yes" : "No"