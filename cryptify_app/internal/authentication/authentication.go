package authentication

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/badoux/checkmail"
	"github.com/matcornic/hermes/v2"
)

// User represents a user in the system
type User struct {
	ID                int
	Email             string
	PasswordHash      string
	IsVerified        bool
	VerificationToken string
}

// AuthService handles user authentication and email verification
type AuthService struct {
	users           map[int]User
	usersByEmail    map[string]int
	usersMutex      sync.Mutex
	verificationMap map[string]int // Map token to user ID for email verification
	hermes          *hermes.Hermes // Email templating engine
}

// NewAuthService creates a new instance of AuthService
func NewAuthService() *AuthService {
	return &AuthService{
		users:           make(map[int]User),
		usersByEmail:    make(map[string]int),
		verificationMap: make(map[string]int),
		hermes:          hermes.New(),
	}
}

// RegisterUser registers a new user and sends a verification email
func (a *AuthService) RegisterUser(email, password string) (int, error) {
	if err := checkmail.ValidateFormat(email); err != nil {
		return 0, fmt.Errorf("invalid email format: %v", err)
	}

	a.usersMutex.Lock()
	defer a.usersMutex.Unlock()

	// Check if email is already registered
	if _, exists := a.usersByEmail[email]; exists {
		return 0, fmt.Errorf("email already registered")
	}

	// Generate a random verification token
	token := generateVerificationToken()

	// Create a new user
	userID := len(a.users) + 1
	newUser := User{
		ID:                userID,
		Email:             email,
		PasswordHash:      password, // In a real scenario, you would hash the password
		VerificationToken: token,
	}

	// Add the user to the maps
	a.users[userID] = newUser
	a.usersByEmail[email] = userID
	a.verificationMap[token] = userID

	// Send verification email
	a.sendVerificationEmail(email, token)

	return userID, nil
}

// VerifyEmail verifies the user's email based on the token
func (a *AuthService) VerifyEmail(token string) error {
	a.usersMutex.Lock()
	defer a.usersMutex.Unlock()

	userID, exists := a.verificationMap[token]
	if !exists {
		return fmt.Errorf("invalid verification token")
	}

	// Mark the user as verified
	a.users[userID].IsVerified = true

	// Remove the token from the verification map
	delete(a.verificationMap, token)

	return nil
}

// sendVerificationEmail sends a verification email to the user
func (a *AuthService) sendVerificationEmail(email, token string) {
	// In a real-world scenario, you would use a proper email sending mechanism here
	// This example uses a simple fmt.Println to simulate sending an email

	emailBody, err := a.generateVerificationEmail(email, token)
	if err != nil {
		fmt.Printf("Error generating email: %v\n", err)
		return
	}

	fmt.Printf("Verification email sent to %s:\n%s\n", email, emailBody)
}

// generateVerificationEmail generates the content for the verification email
func (a *AuthService) generateVerificationEmail(email, token string) (string, error) {
	emailTemplate := hermes.Email{
		Body: hermes.Body{
			Name:   "User",
			Intros: []string{"Welcome to the Crypto Exchange App!"},
			Actions: []hermes.Action{
				{
					Instructions: "To verify your email, please click here:",
					Button: hermes.Button{
						Color: "#22BC66",
						Text:  "Verify Email",
						Link:  fmt.Sprintf("http://localhost:8080/verify?token=%s", token),
					},
				},
			},
		},
	}

	emailBody, err := a.hermes.GenerateHTML(emailTemplate)
	if err != nil {
		return "", fmt.Errorf("error generating email body: %v", err)
	}

	return emailBody, nil
}

// generateVerificationToken generates a random verification token
func generateVerificationToken() string {
	const tokenLength = 32
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())

	tokenBytes := make([]byte, tokenLength)
	for i := range tokenBytes {
		tokenBytes[i] = letters[rand.Intn(len(letters))]
	}

	return string(tokenBytes)
}
