def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n, k = getis
s = gets.chomp

h = Hash.new(0)

s.each_char do |c|
  h[c] += 1
end

# nの階乗を求める
def fact(n)
  return 1 if n == 0
  return n * fact(n-1)
end

f = fact(n)
h.each do |_, v|
  f /= fact(v)
end

# 与えられた文字列の文字をすべて使って、文字列の並び替えを列挙
require 'set'

def permute(str)
  return [str] if str.length <= 1

  chars = str.chars
  result = Set.new

  chars.each_with_index do |char, index|
    remaining = chars[0...index] + chars[index+1..-1]
    permute(remaining.join).each do |perm|
      result << char + perm
    end
  end

  result.to_a
end

def is_palindrome?(str)
  str == str.reverse
end

def generate_palindromes(s, k, current_str = "", palindromes = Set.new)
  if current_str.length == k
    palindromes.add(current_str) if is_palindrome?(current_str)
    return
  end

  s.each_char.with_index do |char, index|
    remaining = s[0...index] + s[index+1..-1]
    generate_palindromes(remaining, k, current_str + char, palindromes)
  end

  palindromes
end

def find_palindromes(s, k)
  palindromes = generate_palindromes(s, k)
  palindromes.to_a
end

p permute(s).size
p find_palindromes(s, k)