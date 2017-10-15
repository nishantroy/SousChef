package mas

// USER INFO
type User struct {
	Name string `json:"name"`
	//Meals interface{} `json:"meals"`
	WeeklyPlan map[string]DailyPlan `json:"weekly_plan"`
	Diet       []string             `json:"diet"`
	Exclusions []string             `json:"exclusions"`
}

type DailyPlan struct {
	Breakfast Meal               `json:"breakfast"`
	Lunch     Meal               `json:"lunch"`
	Dinner    Meal               `json:"dinner"`
	Nutrition map[string]float32 `json:"nutrients"`
}

type Meal struct {
	RecipeID    int    `json:"recipe_id"`
	RecipeTitle string `json:"recipe_title"`
	RecipeImage string `json:"recipe_image"`
	CookTime    int    `json:"ready_in_minutes"`
}

// RECIPE INFO
type Recipe struct {
	Ingredients  []Ingredient `json:"extendedIngredients"`
	ID           int          `json:"id"`
	Title        int          `json:"title"`
	CookTime     int          `json:"readyInMinutes"`
	Image        string       `json:"image"`
	Instructions string       `json:"instructions"`
	Vegetarian   bool         `json:"vegetarian"`
	Vegan        bool         `json:"vegan"`
	GlutenFree   bool         `json:"glutenFree"`
	DairyFree    bool         `json:"dairyFree"`
	Cheap        bool         `json:"cheap"`
	LowFodmap    bool         `json:"lowFodmap"`
	Ketogenic    bool         `json:"ketogenic"`
	Whole30      bool         `json:"whole30"`
	Servings     bool         `json:"servings"`
}

type Ingredient struct {
	ID             int     `json:"id"`
	Category       string  `json:"aisle"`
	Name           string  `json:"name"`
	Amount         float32 `json:"amount"`
	Unit           string  `json:"unitShort"`
	FullDescriptor string  `json:"originalString"`
}