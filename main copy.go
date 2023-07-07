// package main

// import (
// 	"fmt"

// 	"example.com/jwt/initializers"
// 	"example.com/jwt/types"
// )


// func init(){
// 	initializers.LoadEnvFiles()
// 	initializers.ConnectToDB()
// 	initializers.SyncDatabase()
// 	// initializers.ConnectRabbitMq()
// 	// rabbitMQ.RabbitMQConsumer()
// }

// func main(){

// 	// r := gin.Default()

// 	// r.GET("/", func(c *gin.Context) {
// 	// 	c.JSON(200, gin.H{
// 	// 		"message": "pong",
// 	// 	})
// 	// })

// 	// r.Run() 


// 	//깊은복사 
// 	a := types.NewPerson("강동원",32)
// 	b := a.Copy()
// 	b.Name="원빈"
// 	fmt.Printf("깊은복사\na주소:%v\nb주소:%v\n",&a.Name, &b.Name)
// 	fmt.Printf("%s\n%s\n",a.Name,b.Name)

// 	a = types.NewPerson("강동원",30)
// 	b = a
// 	b.Name = "원빈"
// 	fmt.Printf("얕은복사\na주소:%v\nb주소:%v\n",&a.Name, &b.Name)
// 	fmt.Printf("%s\n%s\n",a.Name,b.Name)

// } 