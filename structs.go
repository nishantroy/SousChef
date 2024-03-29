package mas

// User defines the user profile and their weekly plan
type User struct {
	Name           string   `json:"name"`
	Days           []Day    `json:"weekly_plan"`
	Diet           string   `json:"diet"`
	Exclusions     []string `json:"exclusions"`
	WeeklyPlanDate string   `json:"weekly_plan_date"`
}

// Recipe defines the fields for a recipe
type Recipe struct {
	Ingredients  []ingredient  `json:"extendedIngredients"`
	ID           int           `json:"id"`
	Title        string        `json:"title"`
	CookTime     int           `json:"readyInMinutes"`
	Image        string        `json:"image"`
	Instructions []instruction `json:"analyzedInstructions"`
	Vegetarian   bool          `json:"vegetarian"`
	Vegan        bool          `json:"vegan"`
	GlutenFree   bool          `json:"glutenFree"`
	DairyFree    bool          `json:"dairyFree"`
	Cheap        bool          `json:"cheap"`
	LowFodmap    bool          `json:"lowFodmap"`
	Ketogenic    bool          `json:"ketogenic"`
	Whole30      bool          `json:"whole30"`
	Servings     int           `json:"servings"`
	Nutrition    nutrition     `json:"nutrition"`
}

// RecipeChanges gives the user options to change a recipe
type RecipeChanges struct {
	Recipes []struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		CookTime int    `json:"readyInMinutes"`
		Image    string `json:"image"`
	} `json:"results"`
	BaseURL      string `json:"baseUri"`
	Offset       int    `json:"offset"`
	Number       int    `json:"number"`
	TotalResults int    `json:"totalResults"`
}

// instruction defines all steps of a recipe
type instruction struct {
	Steps []step `json:"steps"`
}

// step defines each step of a recipe
type step struct {
	Number int    `json:"number"`
	Step   string `json:"step"`
}

// ingredient defines the fields for an ingredient of a recipe
type ingredient struct {
	ID             int     `json:"id"`
	Category       string  `json:"aisle"`
	Name           string  `json:"name"`
	Amount         float32 `json:"amount"`
	Unit           string  `json:"unitShort"`
	FullDescriptor string  `json:"originalString"`
}

// nutrition contains nutrient details and caloric breakdown
type nutrition struct {
	Nutrients        []nutrient         `json:"nutrients"`
	CaloricBreakdown map[string]float32 `json:"caloricBreakdown"`
}

// nutrient defines the fields for a nutrient of a recipe
type nutrient struct {
	Title   string  `json:"title"`
	Amount  float32 `json:"amount"`
	Unit    string  `json:"unit"`
	Percent float32 `json:"percentOfDailyNeeds"`
}

// WeekPlan holds a week's worth of recipes
type WeekPlan struct {
	Days []Day
}

// Day holds 3 meals for a day
type Day struct {
	Breakfast Meal
	Lunch     Meal
	Dinner    Meal
}

// Meal defines the fields for a meal in a meal plan: the recipe ID, name, cook time and image URL
type Meal struct {
	ID       int
	Name     string
	CookTime int
	Image    string
}

// GroceryItem is each entry in the shopping list
type GroceryItem struct {
	UnitMap map[string]float32
	Done    bool
}
