/*
 * Stack Exchange API
 *
 * Stack Exchange API
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package stackexchange

type QuestionsResponse struct {
	// A list of questions.
	Items          []Question `json:"items,omitempty"`
	HasMore        bool       `json:"has_more,omitempty"`
	QuotaMax       int32      `json:"quota_max,omitempty"`
	QuotaRemaining int32      `json:"quota_remaining,omitempty"`
}
