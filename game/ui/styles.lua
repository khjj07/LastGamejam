return {
	default = {
		font_height = 42,
		spacing = 1, -- pixels between letters
		scale = 1, -- scale of character
		waving = false, -- if true - do waving by sinus
		color = "#f5e3bc",
		speed = 0.04,
		appear = 0.05,
		shaking = 0, -- shaking power. Set to 0 to disable
		sound = "letter",
		can_skip = true,
		shake_on_write = 0 -- when letter appear, shake dialogue screen
	},
	alert1={
		shaking = 5,
		font_height=20,
		scale=2.5,
		color = "#FF0000"
	},
	alert2={
		shaking = 5,
		font_height=10,
		scale=2
	}
}