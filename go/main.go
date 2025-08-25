package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type WebhookHandler struct {
	secret []byte
}

type GitHubWebhookPayload struct {
	Action     string                 `json:"action,omitempty"`
	Repository map[string]interface{} `json:"repository,omitempty"`
	Sender     map[string]interface{} `json:"sender,omitempty"`
}

func NewWebhookHandler(secret string) *WebhookHandler {
	return &WebhookHandler{
		secret: []byte(secret),
	}
}

func (wh *WebhookHandler) validateSignature(payload []byte, signature string) bool {
	if !strings.HasPrefix(signature, "sha256=") {
		return false
	}

	expectedMAC := hmac.New(sha256.New, wh.secret)
	expectedMAC.Write(payload)
	expectedSignature := "sha256=" + hex.EncodeToString(expectedMAC.Sum(nil))
	log.Printf("Expected signature: %s", expectedSignature)

	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}

func (wh *WebhookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	signature := r.Header.Get("X-Hub-Signature-256")
	if signature == "" {
		http.Error(w, "Missing signature header", http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if !wh.validateSignature(body, signature) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	eventType := r.Header.Get("X-GitHub-Event")
	if eventType == "" {
		http.Error(w, "Missing event type header", http.StatusBadRequest)
		return
	}

	var payload GitHubWebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Printf("Failed to parse JSON payload: %v", err)
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	log.Printf("Received GitHub webhook event: %s", eventType)

	switch eventType {
	case "push":
		wh.handlePushEvent(payload)
	case "pull_request":
		wh.handlePullRequestEvent(payload)
	case "issues":
		wh.handleIssuesEvent(payload)
	default:
		log.Printf("Unhandled event type: %s", eventType)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received successfully"))
}

func (wh *WebhookHandler) handlePushEvent(payload GitHubWebhookPayload) {
	log.Println("Processing push event")
	if repo, ok := payload.Repository["name"].(string); ok {
		log.Printf("Push to repository: %s", repo)
	}
}

func (wh *WebhookHandler) handlePullRequestEvent(payload GitHubWebhookPayload) {
	log.Printf("Processing pull request event: %s", payload.Action)
	if repo, ok := payload.Repository["name"].(string); ok {
		log.Printf("Pull request in repository: %s", repo)
	}
}

func (wh *WebhookHandler) handleIssuesEvent(payload GitHubWebhookPayload) {
	log.Printf("Processing issues event: %s", payload.Action)
	if repo, ok := payload.Repository["name"].(string); ok {
		log.Printf("Issue in repository: %s", repo)
	}
}

func main() {
	secret := os.Getenv("GITHUB_WEBHOOK_SECRET")
	if secret == "" {
		log.Fatal("GITHUB_WEBHOOK_SECRET environment variable is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := NewWebhookHandler(secret)

	r := mux.NewRouter()
	r.HandleFunc("/webhook", handler.HandleWebhook).Methods("POST")
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	log.Printf("GitHub webhook handler starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
