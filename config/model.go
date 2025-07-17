package config

var ModelReverseMap = map[string]string{}
var ModelMap = map[string]string{
	"claude-4.0-sonnet":       "claude2",
	"claude-4.0-sonnet-think": "claude37sonnetthinking",
	"deepseek-r1":             "r1",
	"o4-mini":                 "o4mini",
	"gpt-4o":                  "gpt4o",
	"gemini-2.5-pro-06-05":    "gemini2flash",
	"grok4":                   "grok4",
	"gpt-4.1":                 "gpt41",
	// "claude-4.0-opus":         "claude40opus",
	"claude-4.0-opus-think": "claude40opusthinking",
	"o3":                    "o3",
	"o3-pro":                "o3pro",
}

// Get returns the value for the given key from the ModelMap.
// If the key doesn't exist, it returns the provided default value.
func ModelMapGet(key string, defaultValue string) string {
	if value, exists := ModelMap[key]; exists {
		return value
	}
	return defaultValue
}

// GetReverse returns the value for the given key from the ModelReverseMap.
// If the key doesn't exist, it returns the provided default value.
func ModelReverseMapGet(key string, defaultValue string) string {
	if value, exists := ModelReverseMap[key]; exists {
		return value
	}
	return defaultValue
}

var ResponseModles []map[string]string

func init() {
	for k, v := range ModelMap {
		ModelReverseMap[v] = k
		model := map[string]string{
			"id": k,
		}
		// 互联网搜索
		modelSearch := map[string]string{
			"id": k + "-search",
		}
		// 学术搜索
		modelScholar := map[string]string{
			"id": k + "-scholar",
		}
		// 社区搜索
		modelSocial := map[string]string{
			"id": k + "-social",
		}
		// Edgar数据库搜索
		modelEdgar := map[string]string{
			"id": k + "-edgar",
		}
		// 全部搜索
		modelAll := map[string]string{
			"id": k + "-all",
		}
		ResponseModles = append(ResponseModles, model, modelSearch, modelScholar, modelSocial, modelEdgar, modelAll)
	}
}
