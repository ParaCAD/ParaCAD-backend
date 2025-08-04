package auth

import (
	"testing"
	"time"
)

func TestAuth(t *testing.T) {
	shortAuth := New("secret", 1*time.Second)
	if shortAuth == nil {
		t.Fatal("Expected non-nil Auth instance")
	}

	shortJWT, err := shortAuth.CreateToken("testUserID", "testUser", RoleUser)
	if err != nil {
		t.Fatalf("Expected no error creating token, got: %v", err)
	}
	if shortJWT == "" {
		t.Fatal("Expected non-empty JWT token")
	}

	time.Sleep(2 * time.Second)

	_, _, err = shortAuth.VerifyToken(shortJWT)
	if err == nil {
		t.Fatal("Expected error verifying expired token, got nil")
	}

	longAuth := New("longSecret", 10*time.Minute)
	if longAuth == nil {
		t.Fatal("Expected non-nil Auth instance")
	}

	longJWT, err := longAuth.CreateToken("testUserID", "testUser", RoleAdmin)
	if err != nil {
		t.Fatalf("Expected no error creating token, got: %v", err)
	}
	if longJWT == "" {
		t.Fatal("Expected non-empty JWT token")
	}

	userID, role, err := longAuth.VerifyToken(longJWT)
	if err != nil {
		t.Fatalf("Expected no error verifying token, got: %v", err)
	}
	if userID != "testUserID" {
		t.Fatalf("Expected userID 'testUserID', got: %s", userID)
	}
	if role != RoleAdmin {
		t.Fatalf("Expected role 'RoleAdmin', got: %s", role)
	}
}
