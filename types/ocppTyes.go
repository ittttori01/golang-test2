package types

import (
	"bytes"
	"encoding/gob"
)

// OCPP type
type MessageQueName struct{
	IncomeRequest string `json:"IncomeRequest"`
	IncomeResponse string `json:"IncomeResponse"`
	Outcome string `json:"Outcome"`
    OutcomeRequest string `json:"OutcomeRequest"`	
    OutcomeResponse string `json:"OutcomeResponse"`
}

type MessageQueueRouteKey struct{
	IncomeRequest  string
	IncomeResponse string
	OutcomeRequest string
	OutcomeResponse string
}

type MessageQueueExchange struct{
    Incomes string
	Outcomes string
}

type Person struct {
	Name string
	Age int
}

func NewPerson(name string, age int) *Person {
	
	p := &Person{}
	p.Name = name
	p.Age = age
	
	return p 
	

}

func (person *Person) Copy() *Person {
	bytes := bytes.Buffer{}
	newPerson := &Person{}
	encoder := gob.NewEncoder(&bytes)
	decoder := gob.NewDecoder(&bytes)
	
	_ = encoder.Encode(&person)
	_ = decoder.Decode(&newPerson)

	return newPerson
}



 
 
 
 
 
 
 
 
 
 
