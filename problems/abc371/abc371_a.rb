
sab, sac, sbc = gets.chomp.split(" ")

if sab == '>'
  if sac == '>'
    if sbc == '>'
      puts 'B'
    else
      puts 'C'
    end
  else
    puts 'A'
  end
else
  if sac == '>'
    puts 'A'
  else
    if sbc == '>'
      puts 'C'
    else
      puts 'B'
    end
  end
end