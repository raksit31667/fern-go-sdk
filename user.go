// This file was auto-generated by Fern from our API Definition.

package api

type LoginUserRequest struct {
	// The user name for login
	Username *string `json:"-"`
	// The password for login in clear text
	Password *string `json:"-"`
}
