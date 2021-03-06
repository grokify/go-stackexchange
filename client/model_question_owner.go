/*
 * Stack Exchange API
 *
 * Stack Exchange API
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package stackoverflow

type QuestionOwner struct {
	Reputation   int32  `json:"reputation,omitempty"`
	UserId       int32  `json:"user_id,omitempty"`
	UserType     string `json:"user_type,omitempty"`
	ProfileImage string `json:"profile_image,omitempty"`
	DisplayName  string `json:"display_name,omitempty"`
	Link         string `json:"link,omitempty"`
}
