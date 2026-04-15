package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"unicode"
)

type searchRule struct {
	Route  string
	Weight int
	Roots  []string
}

var searchRules = []searchRule{
	{
		Route:  "/accruals#cold",
		Weight: 16,
		Roots:  []string{"холод", "хвс", "водоканал", "счетчик", "счетч", "куб"},
	},
	{
		Route:  "/accruals#hot",
		Weight: 16,
		Roots:  []string{"горяч", "гвс", "бойлер", "нагрев", "генерац", "гкал"},
	},
	{
		Route:  "/accruals#heat",
		Weight: 16,
		Roots:  []string{"отопл", "тепл", "батар", "радиатор", "обогрев", "гкал"},
	},
	{
		Route:  "/faq",
		Weight: 11,
		Roots:  []string{"вопрос", "ответ", "почему", "зачем", "как", "кто", "что", "когда", "сосед", "затоп", "льгот", "оспор", "помощ"},
	},
	{
		Route:  "/complaints",
		Weight: 10,
		Roots:  []string{"жалоб", "претенз", "обращен", "отказ", "наруш", "гжи", "бездеи", "неправ"},
	},
	{
		Route:  "/contacts",
		Weight: 9,
		Roots:  []string{"контакт", "телефон", "адрес", "ук", "управля", "авар", "диспетч", "найти"},
	},
	{
		Route:  "/tariffs",
		Weight: 8,
		Roots:  []string{"тариф", "стоим", "цен", "руб", "поставщик", "норматив"},
	},
	{
		Route:  "/consumer/epd",
		Weight: 10,
		Roots:  []string{"епд", "единый", "платеж", "документ", "расшифров", "квитанц"},
	},
	{
		Route:  "/consumer",
		Weight: 7,
		Roots:  []string{"потребител", "справочник", "права", "жкх", "жку"},
	},
	{
		Route:  "/accruals",
		Weight: 4,
		Roots:  []string{"начисл", "калькул", "расчет", "рассчит", "платеж", "сумм"},
	},
}

var stopTokens = map[string]struct{}{
	"и":   {},
	"в":   {},
	"во":  {},
	"на":  {},
	"по":  {},
	"к":   {},
	"ко":  {},
	"с":   {},
	"со":  {},
	"о":   {},
	"об":  {},
	"от":  {},
	"до":  {},
	"за":  {},
	"у":   {},
	"из":  {},
	"для": {},
}

func SearchHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on search page")

	query := req.URL.Query().Get("q")
	if strings.TrimSpace(query) == "" {
		http.Redirect(w, req, "/", http.StatusFound)
		return
	}

	route := resolveSearchRoute(query)
	if route == "" {
		route = "/"
	}

	http.Redirect(w, req, route, http.StatusFound)
}

func resolveSearchRoute(query string) string {
	tokens := tokenize(query)
	if len(tokens) == 0 {
		return ""
	}

	tariffIntentPos := firstTokenPrefixPosition(tokens, []string{"тариф", "стоим", "цен", "руб", "норматив", "поставщик"})

	hasWater := hasTokenPrefix(tokens, "вод") || hasTokenPrefix(tokens, "вода")
	hotPos := firstTokenPrefixPosition(tokens, []string{"горяч", "гвс"})
	coldPos := firstTokenPrefixPosition(tokens, []string{"холод", "хвс"})
	calcIntentPos := firstTokenPrefixPosition(tokens, []string{"калькул", "рассчит", "расчет", "начисл", "провер"})
	heatPos := firstTokenPrefixPosition(tokens, []string{"отопл", "тепл", "батар", "радиатор"})
	sewerPos := firstTokenPrefixPosition(tokens, []string{"водоотвед", "канализ", "сток"})
	gasPos := firstTokenPrefixPosition(tokens, []string{"газ", "газоснаб"})
	electricityPos := firstTokenPrefixPosition(tokens, []string{"электро", "электр", "свет", "квт"})
	wastePos := firstTokenPrefixPosition(tokens, []string{"тко", "мусор", "отход", "тверьспец", "спецавто"})

	if tariffIntentPos >= 0 {
		bestUtilityPos := minPositive(
			minPositive(hotPos, coldPos),
			minPositive(
				minPositive(minPositive(heatPos, sewerPos), minPositive(gasPos, electricityPos)),
				wastePos,
			),
		)
		if bestUtilityPos == -1 || tariffIntentPos <= bestUtilityPos {
			topic := bestTariffTopic([]topicMatch{
				{position: coldPos, topic: "cold"},
				{position: hotPos, topic: "hot"},
				{position: heatPos, topic: "heat"},
				{position: sewerPos, topic: "sewer"},
				{position: gasPos, topic: "gas"},
				{position: electricityPos, topic: "electricity"},
				{position: wastePos, topic: "waste"},
			})
			if topic != "" {
				return "/tariffs#" + topic
			}
			return "/tariffs"
		}
	}

	if hasWater && hotPos >= 0 {
		if calcIntentPos == -1 {
			return "/tariffs#hot"
		}
		return "/accruals#hot"
	}
	if hasWater && coldPos >= 0 {
		return "/accruals#cold"
	}

	bestRoute := ""
	bestScore := 0
	for _, rule := range searchRules {
		score := scoreRule(rule, tokens)
		if score > bestScore {
			bestScore = score
			bestRoute = rule.Route
		}
	}

	if bestScore == 0 {
		return ""
	}
	return bestRoute
}

func tokenize(source string) []string {
	normalized := normalizeText(source)
	if normalized == "" {
		return nil
	}

	parts := strings.Fields(normalized)
	seen := make(map[string]struct{}, len(parts))
	tokens := make([]string, 0, len(parts))

	for _, token := range parts {
		if len(token) < 2 {
			continue
		}
		if _, stop := stopTokens[token]; stop {
			continue
		}
		if _, ok := seen[token]; ok {
			continue
		}
		seen[token] = struct{}{}
		tokens = append(tokens, token)
	}

	return tokens
}

func normalizeText(source string) string {
	source = strings.ToLower(strings.TrimSpace(source))
	source = strings.ReplaceAll(source, "ё", "е")

	var builder strings.Builder
	builder.Grow(len(source))

	for _, r := range source {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) || r == '-' {
			builder.WriteRune(r)
			continue
		}
		builder.WriteRune(' ')
	}

	return strings.Join(strings.Fields(builder.String()), " ")
}

func scoreRule(rule searchRule, tokens []string) int {
	score := 0
	matchedRoots := 0

	for _, root := range rule.Roots {
		matchPos := -1
		for idx, token := range tokens {
			if tokenMatchesRoot(token, root) {
				matchPos = idx
				break
			}
		}

		if matchPos >= 0 {
			matchedRoots++
			positionBonus := 0
			if matchPos < 5 {
				positionBonus = 5 - matchPos
			}
			score += rule.Weight + positionBonus
		}
	}

	if matchedRoots == 0 {
		return 0
	}

	return score + matchedRoots*2
}

func tokenMatchesRoot(token, root string) bool {
	if strings.HasPrefix(token, root) {
		return true
	}

	return len(token) >= 4 && strings.HasPrefix(root, token)
}

func hasTokenPrefix(tokens []string, prefix string) bool {
	for _, token := range tokens {
		if strings.HasPrefix(token, prefix) {
			return true
		}
	}
	return false
}

func firstTokenPrefixPosition(tokens []string, prefixes []string) int {
	for i, token := range tokens {
		for _, prefix := range prefixes {
			if strings.HasPrefix(token, prefix) {
				return i
			}
		}
	}
	return -1
}

func minPositive(a, b int) int {
	if a < 0 {
		return b
	}
	if b < 0 {
		return a
	}
	if a < b {
		return a
	}
	return b
}

type topicMatch struct {
	position int
	topic    string
}

func bestTariffTopic(matches []topicMatch) string {
	bestPos := -1
	bestTopic := ""
	for _, item := range matches {
		if item.position < 0 {
			continue
		}
		if bestPos == -1 || item.position < bestPos {
			bestPos = item.position
			bestTopic = item.topic
		}
	}
	return bestTopic
}

