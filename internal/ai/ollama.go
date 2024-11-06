package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"github.com/AashishKumar-3002/FealtyX/internal/models"
)

type GenerateRequest struct { 
    Model string `json:"model"` 
    Prompt string `json:"prompt"` 
    Stream bool `json:"stream"` 
    } 
type GenerateResponse struct {
     Model string `json:"model"` 
     CreatedAt string `json:"created_at"` 
     Response string `json:"response"` 
     Done bool `json:"done"` 
     Context []int `json:"context"` 
     TotalDuration int64 `json:"total_duration"` 
     LoadDuration int64 `json:"load_duration"` 
     PromptEvalCount int `json:"prompt_eval_count"` 
     PromptEvalDuration int64 `json:"prompt_eval_duration"` 
     EvalCount int `json:"eval_count"` 
     EvalDuration int64 `json:"eval_duration"` }


func GenerateStudentSummary(student models.Student) (string, error) {
    prompt := fmt.Sprintf("Generate a brief summary for a student named %s, who is %d years old and has the email %s.", student.Name, student.Age, student.Email)
    const ollamaURL = "http://localhost:12345/api/generate"
    requestBody := GenerateRequest{ 
        Model: "llama3.2:1b", 
        Prompt: prompt, 
        Stream: false, 
    } 
    jsonStr, err := json.Marshal(requestBody)
    if err != nil { 
        return "", fmt.Errorf("error marshalling request body: %v", err) 
    }

    req, err := http.NewRequest("POST", ollamaURL, bytes.NewBuffer(jsonStr))
    if err != nil { 
        return "", fmt.Errorf("error creating request: %v", err) 
    } 
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{} 
    resp, err := client.Do(req) 
    if err != nil { 
        return "", fmt.Errorf("error making request: %v", err) 
    } 
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("error reading response: %v", err)
    }

    var generateResponse GenerateResponse 
    err = json.Unmarshal(body, &generateResponse) 
    if err != nil { 
        return "", fmt.Errorf("error unmarshalling response: %v", err) 
    } 
    return generateResponse.Response, nil

}