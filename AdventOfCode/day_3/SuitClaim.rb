class SuitClaim
	def initialize(id, startX, startY, width, height)
		@id = id
		@startX = startX.to_i
		@startY = startY.to_i
		@width = width.to_i
		@height = height.to_i
	end

	def id
		@id
	end

	def startX
		@startX
	end

	def startY
		@startY
	end

	def width
		@width
	end

	def height
		@height
	end
end