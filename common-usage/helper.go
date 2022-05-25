package main

func generatePayloads() []Payload {
	payloads := []Payload{
		{
			Name: "name1",
			Job:  "job1",
		}, {
			Name: "name2",
			Job:  "job2",
		},
	}

	return payloads
}

func generatePayload() Payload {
	return Payload{
		Name: "name1",
		Job:  "job1",
	}
}
