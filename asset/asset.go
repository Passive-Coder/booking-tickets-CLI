package asset

type Conference struct {
	Name           string
	SeatsAvailable uint
}

type UserData struct {
	name            string
	email           string
	numberOfTickets uint
}

var Conferences = []Conference{
	{
		"Go Conference",
		50,
	},
	{
		"Python Conference",
		69,
	},
	{
		"NVIDIA conference",
		1000,
	},
}
