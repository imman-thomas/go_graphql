package main

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

const var graphql_url : String = "https://api.test.com/graphql"

func main() {
	client := graphql.NewClient(graphql_url)
	queryData(client)
}

// func queryData(client *graphql.Client) {
// 	mutation := `
// 	mutation($email: String!,$password: String!)
// {

// login(data:{
//   email: $email,
//   password: $password
// }){
//   token
//   organisation{
// 	name
//   }

// }
// }
// 	`
// 	request := graphql.NewRequest(mutation)
// 	request.Var("email", "test@holidayinn.com")
// 	request.Var("password", "aePeig1o")

// 	request.Header.Set("Cache-Control", "no-cache")

// 	var response MutationLoginResponse
// 	err := client.Run(context.Background(), request, &response)
// 	if err != nil {
// 		panic(err)
// 	}

// 	log.Println(response)

// 	log.Println(response.Login.Token)
// }

// type MutationLoginResponse struct {
// 	Login struct {
// 		Token        string
// 		Organisation struct {
// 			Name string
// 		}
// 	}
// }

func queryData(client *graphql.Client) {

	query := `
	query{
		getOrganisation{
			name 
		}
	}
 
	`
	const token = "sample_token"
	request := graphql.NewRequest(query)
	request.Header.Set("Cache-Control", "no-cache")

	request.Header.Set("Authorization", token)

	var resp QueryOrganisationResponse
	err := client.Run(context.Background(), request, &resp)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

	for _, arr := range resp.GetOrganisation {
		fmt.Println(arr)
	}

}

type QueryOrganisationResponse struct {
	GetOrganisation []struct {
		Name string
	}
}
