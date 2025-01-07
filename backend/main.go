package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Parse form data from the request body
	parsedFormData, err := url.ParseQuery(request.Body)
	if err != nil {
		log.Printf("Error parsing form data: %s", err)
		return events.APIGatewayProxyResponse{StatusCode: 400, Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		}}, err
	}

	destination := parsedFormData.Get("destination")
	budget := parsedFormData.Get("budget")
	activityLevel := parsedFormData.Get("activityLevel")
	activityTypes := parsedFormData.Get("activityTypes")
	resturantTypes := parsedFormData.Get("resturantTypes")
	radiusMiles := parsedFormData.Get("radiusMiles")
	numDays := parsedFormData.Get("numDays")

	apiKey := os.Getenv("api_key")

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 400, Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		}}, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"

	activitySchema := &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"name": {
				Type: genai.TypeString,
			},
			"description": {
				Type: genai.TypeString,
			},
		},
		Required: []string{"name", "description"},
	}

	dayAgendaSchema := &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"day": {
				Type: genai.TypeNumber,
			},
			"hotel": {
				Type: genai.TypeString,
			},
			"breakfast": {
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"name": {
						Type: genai.TypeString,
					},
					"description": {
						Type: genai.TypeString,
					},
				},
				Required: []string{"name", "description"},
			},
			"morning_activities": {
				Type:  genai.TypeArray,
				Items: activitySchema,
			},
			"lunch": {
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"name": {
						Type: genai.TypeString,
					},
					"description": {
						Type: genai.TypeString,
					},
				},
				Required: []string{"name", "description"},
			},
			"afternoon_activities": {
				Type:  genai.TypeArray,
				Items: activitySchema,
			},
			"dinner": {
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"name": {
						Type: genai.TypeString,
					},
					"description": {
						Type: genai.TypeString,
					},
				},
				Required: []string{"name", "description"},
			},
			"night_activities": {
				Type:  genai.TypeArray,
				Items: activitySchema,
			},
		},
		Required: []string{"day", "hotel", "breakfast", "morning_activities", "lunch", "afternoon_activities", "dinner", "night_activities"},
	}
	model.ResponseSchema = &genai.Schema{
		Type:  genai.TypeArray,
		Items: dayAgendaSchema,
	}

	session := model.StartChat()
	session.History = []*genai.Content{}

	prompt := fmt.Sprintf(os.Getenv("prompt"), destination, budget, activityLevel, activityTypes, resturantTypes, radiusMiles, numDays)
	resp, err := session.SendMessage(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 400, Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		}}, err
	}

	if txt, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       string(txt),
			Headers: map[string]string{
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Credentials": "true",
			},
		}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 400, Headers: map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Credentials": "true",
	}}, err
}

func main() {
	lambda.Start(handler)
}