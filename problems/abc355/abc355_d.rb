def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti

lrs = []
n.times do
  li, ri = getis
  lrs << [li, ri]
end

class FenwickTree
  def initialize(size)
    @size = size
    @tree = Array.new(size + 1, 0)
  end

  def update(index, value)
    while index <= @size
      @tree[index] += value
      index += index & -index
    end
  end

  def query(index)
    sum = 0
    while index > 0
      sum += @tree[index]
      index -= index & -index
    end
    sum
  end
end

def count_overlapping_intervals(intervals)
  points = intervals.flat_map { |interval| [interval[0], interval[1]] }.uniq
  points.sort!
  point_index = {}
  points.each_with_index do |point, index|
    point_index[point] = index + 1
  end

  bit = FenwickTree.new(points.size)
  events = []

  intervals.each do |interval|
    events << [point_index[interval[0]], :start]
    events << [point_index[interval[1]], :end]
  end

  events.sort_by! { |event| [event[0], event[1] == :start ? 0 : 1] }

  overlap_count = 0
  active_intervals = 0

  events.each do |event|
    if event[1] == :start
      overlap_count += active_intervals
      active_intervals += 1
      bit.update(event[0], 1)
    else
      active_intervals -= 1
      bit.update(event[0], -1)
    end
  end

  overlap_count
end

puts count_overlapping_intervals(lrs)
