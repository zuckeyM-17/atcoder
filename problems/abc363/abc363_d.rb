def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti

i = 10
number_of_digits = 1

if n <= 10
  puts n - 1
  exit
end

if n < 19
  puts n
  exit
end

n -= 19
number_of_digits = 3
i = 90

while n - i > 0
  n -= i
  if number_of_digits.even?
    i *= 10
  end
  number_of_digits += 1
end

def nth_k_digit_number(k, n)
  start_number = if k == 1
    0
  else
    10**(k - 1)
  end
  start_number + n - 1
end

if number_of_digits % 2 == 0
  c = nth_k_digit_number(number_of_digits / 2, n)
  puts c.to_s + c.to_s.reverse
else
  c = nth_k_digit_number(number_of_digits / 2+1, n)
  puts c.to_s + c.to_s[0..-2].reverse
end
