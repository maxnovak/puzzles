class GuardEvent
	def initialize(date, time, event)
		@date = date
		@time = time
		@event = event
	end

	def date
		@date
	end

	def time
		@time
	end

	def event
		@event
	end
end