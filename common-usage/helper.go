package main

type Payload struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

type Response struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

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
