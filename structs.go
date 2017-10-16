package mas

// User defines the user profile and their weekly plan
type User struct {
	Name       string               `json:"name"`
	WeeklyPlan map[string]DailyPlan `json:"weekly_plan"`
	Diet       []string             `json:"diet"`
	Exclusions []string             `json:"exclusions"`
}

// DailyPlan defines the meal plan for one day for a user
type DailyPlan struct {
	Breakfast Meal               `json:"breakfast"`
	Lunch     Meal               `json:"lunch"`
	Dinner    Meal               `json:"dinner"`
	Nutrition map[string]float32 `json:"nutrients"`
}

// Meal defines the fields of a meal for a user's meal plan
type Meal struct {
	RecipeID    int    `json:"recipe_id"`
	RecipeTitle string `json:"recipe_title"`
	RecipeImage string `json:"recipe_image"`
	CookTime    int    `json:"ready_in_minutes"`
}

// Recipe defines the fields for a recipe
type Recipe struct {
	Ingredients  []Ingredient `json:"extendedIngredients"`
	ID           int          `json:"id"`
	Title        string       `json:"title"`
	CookTime     int          `json:"readyInMinutes"`
	Image        string       `json:"image"`
	Instructions []string     `json:"instructions"`
	Vegetarian   bool         `json:"vegetarian"`
	Vegan        bool         `json:"vegan"`
	GlutenFree   bool         `json:"glutenFree"`
	DairyFree    bool         `json:"dairyFree"`
	Cheap        bool         `json:"cheap"`
	LowFodmap    bool         `json:"lowFodmap"`
	Ketogenic    bool         `json:"ketogenic"`
	Whole30      bool         `json:"whole30"`
	Servings     int          `json:"servings"`
}

// Ingredient defines the fields for an ingredient of a recipe
type Ingredient struct {
	ID             int     `json:"id"`
	Category       string  `json:"aisle"`
	Name           string  `json:"name"`
	Amount         float32 `json:"amount"`
	Unit           string  `json:"unitShort"`
	FullDescriptor string  `json:"originalString"`
}
