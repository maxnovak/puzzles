
#distance between capital and lowercase letters
$comparisonAmount = "a".ord - "A".ord

p $comparisonAmount

sequence = ""
File.open("day5_input.txt", "r") do |file|
	file.each_line do |data|
		sequence = data.split('')
	end
end

def checkValues(left, right)
	difference = left.ord - right.ord
	if difference.abs == $comparisonAmount
		return []
	end
	return [left] + [right]
end

def doMagic(sequenceOfLetters)
	if sequenceOfLetters.count == 2
		left = sequenceOfLetters[0]
		right = sequenceOfLetters[1]
		result = checkValues(left, right)
		return result
	elsif sequenceOfLetters.count > 2
		left = doMagic(sequenceOfLetters[0..sequenceOfLetters.count/2])
		right = doMagic(sequenceOfLetters[sequenceOfLetters.count/2+1..-1])
		result = checkValues(left.last, right.first)
		return left[0..-2] + result + right[1..-1]
	else
		return sequenceOfLetters
	end
end

p sequence
p doMagic(sequence) == sequence
